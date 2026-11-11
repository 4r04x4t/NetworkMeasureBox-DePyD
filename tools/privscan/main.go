// Command privscan finds personal data and secrets: e-mail addresses, public IP addresses outside
// the documentation and private ranges, MAC addresses with a real vendor prefix, key- and
// token-like strings, developer home folders, and the values of the owner's private list.
//
// Usage:
//
//	privscan [flags] tree [-0] [FILE...]   files named as arguments, or NUL-separated on stdin with -0
//	privscan [flags] dist DIR...           every regular file below the folders, scanned as bytes
//	privscan [flags] history               every blob, commit message and identity of all refs
//
// Flags: -allow FILE (reviewed allow list), -repo DIR (repository root, default "."),
// -require-private (fail when NMB_PRIVATE_VALUES is not set). The private list is read from the
// file named by NMB_PRIVATE_VALUES; it must lie outside the repository and its values are never
// printed. Hits are printed masked. Exit status: 0 clean, 1 hits, 2 usage or read error.
package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	exitClean = 0
	exitHits  = 1
	exitError = 2

	privateEnv = "NMB_PRIVATE_VALUES"
)

// env is what run needs from the process, passed in so tests control it.
type env struct {
	stdin   io.Reader
	stdout  io.Writer
	stderr  io.Writer
	getenv  func(string) string
	gitFrom func(dir string) gitRunner
}

func main() {
	os.Exit(run(os.Args[1:], env{stdin: os.Stdin, stdout: os.Stdout, stderr: os.Stderr,
		getenv: os.Getenv, gitFrom: execGit}))
}

type options struct {
	allow          string
	repo           string
	requirePrivate bool
	nul            bool
	mode           string
	args           []string
}

func parseArgs(args []string) (options, error) {
	var o options
	fl := flag.NewFlagSet("privscan", flag.ContinueOnError)
	fl.SetOutput(io.Discard)
	fl.StringVar(&o.allow, "allow", "", "allow list")
	fl.StringVar(&o.repo, "repo", ".", "repository root")
	fl.BoolVar(&o.requirePrivate, "require-private", false, "fail without a private list")
	if err := fl.Parse(args); err != nil {
		return o, err
	}
	if fl.NArg() == 0 {
		return o, errors.New("no mode given (tree, dist or history)")
	}
	o.mode, o.args = fl.Arg(0), fl.Args()[1:]
	if o.mode == "tree" && len(o.args) > 0 && o.args[0] == "-0" {
		o.nul, o.args = true, o.args[1:]
	}
	return o, nil
}

func run(args []string, e env) int {
	o, err := parseArgs(args)
	if err != nil {
		_, _ = fmt.Fprintf(e.stderr, "privscan: %v\n", err)
		return exitError
	}
	sc, accepted, err := newScannerFor(o, e)
	if err != nil {
		_, _ = fmt.Fprintf(e.stderr, "privscan: %v\n", err)
		return exitError
	}
	hits, files, err := scanMode(sc, o, e, accepted)
	if err != nil {
		_, _ = fmt.Fprintf(e.stderr, "privscan: %v\n", err)
		return exitError
	}
	for _, h := range hits {
		_, _ = fmt.Fprintln(e.stdout, h.String())
	}
	_, _ = fmt.Fprintf(e.stdout, "privscan %s: %d inputs, %d hits\n", o.mode, files, len(hits))
	if len(hits) > 0 {
		return exitHits
	}
	return exitClean
}

func newScannerFor(o options, e env) (*Scanner, map[string]bool, error) {
	allow := map[string]map[string]bool{}
	if o.allow != "" {
		data, err := os.ReadFile(filepath.Clean(o.allow))
		if err != nil {
			return nil, nil, fmt.Errorf("allow list: %w", err)
		}
		if allow, err = LoadAllow(data); err != nil {
			return nil, nil, err
		}
	}
	path := e.getenv(privateEnv)
	if path == "" {
		if o.requirePrivate {
			return nil, nil, errors.New(privateEnv + " is not set and -require-private was given")
		}
		_, _ = fmt.Fprintln(e.stdout, "private list: not used ("+privateEnv+" not set)")
		return NewScanner(allow, nil), nil, nil
	}
	list, err := LoadPrivate(path, o.repo)
	if err != nil {
		return nil, nil, err
	}
	_, _ = fmt.Fprintf(e.stdout, "private list: %d values\n", len(list.values))
	return NewScanner(allow, list.values), list.identities, nil
}

func scanMode(sc *Scanner, o options, e env, accepted map[string]bool) ([]Hit, int, error) {
	switch o.mode {
	case "tree":
		files := o.args
		if o.nul {
			var err error
			if files, err = readNULList(e.stdin); err != nil {
				return nil, 0, err
			}
		}
		hits, err := scanFiles(sc, files)
		return hits, len(files), err
	case "dist":
		return scanDirs(sc, o.args)
	case "history":
		hits, err := sc.ScanHistory(e.gitFrom(o.repo), accepted)
		return hits, 1, err
	}
	return nil, 0, fmt.Errorf("unknown mode %q", o.mode)
}

func readNULList(r io.Reader) ([]string, error) {
	data, err := io.ReadAll(bufio.NewReader(r))
	if err != nil {
		return nil, fmt.Errorf("reading the file list: %w", err)
	}
	var files []string
	for _, name := range bytes.Split(data, []byte{0}) {
		if len(name) > 0 {
			files = append(files, string(name))
		}
	}
	return files, nil
}

// scanFiles scans each file; a symbolic link is scanned as its target text, a folder is an error.
func scanFiles(sc *Scanner, files []string) ([]Hit, error) {
	var hits []Hit
	for _, name := range files {
		h, err := scanFile(sc, name)
		if err != nil {
			return nil, err
		}
		hits = append(hits, h...)
	}
	return hits, nil
}

func scanFile(sc *Scanner, name string) ([]Hit, error) {
	info, err := os.Lstat(name)
	if err != nil {
		return nil, err
	}
	switch {
	case info.Mode()&fs.ModeSymlink != 0:
		target, err := os.Readlink(name)
		if err != nil {
			return nil, err
		}
		return sc.Scan(name, strings.NewReader(target))
	case !info.Mode().IsRegular():
		return nil, fmt.Errorf("%s: not a regular file", name)
	case info.Size() > maxFileSize:
		return nil, fmt.Errorf("%s: larger than %d bytes", name, maxFileSize)
	}
	f, err := os.Open(filepath.Clean(name))
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	return sc.Scan(name, f)
}

func scanDirs(sc *Scanner, dirs []string) ([]Hit, int, error) {
	if len(dirs) == 0 {
		return nil, 0, errors.New("dist: no folder given")
	}
	var files []string
	for _, dir := range dirs {
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() {
				files = append(files, path)
			}
			return err
		})
		if err != nil {
			return nil, 0, err
		}
	}
	hits, err := scanFiles(sc, files)
	return hits, len(files), err
}
