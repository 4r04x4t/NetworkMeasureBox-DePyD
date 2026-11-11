package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

// gitRunner runs git with the arguments (a list, never a shell line) and returns its stdout.
type gitRunner func(stdin []byte, args ...string) ([]byte, error)

// execGit returns a runner for the repository in dir.
func execGit(dir string) gitRunner {
	return func(stdin []byte, args ...string) ([]byte, error) {
		cmd := exec.Command("git", append([]string{"-C", dir}, args...)...) // #nosec G204 -- fixed program, argument list
		cmd.Stdin = bytes.NewReader(stdin)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		out, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
		}
		return out, nil
	}
}

// ScanHistory scans every blob reachable from any ref, every commit message, and checks the
// author and committer addresses of every commit.
func (s *Scanner) ScanHistory(git gitRunner, accepted map[string]bool) ([]Hit, error) {
	hits, err := s.scanBlobs(git)
	if err != nil {
		return nil, err
	}
	msgHits, err := s.scanMessages(git)
	if err != nil {
		return nil, err
	}
	idHits, err := checkIdentities(git, accepted)
	if err != nil {
		return nil, err
	}
	return append(append(hits, msgHits...), idHits...), nil
}

func (s *Scanner) scanBlobs(git gitRunner) ([]Hit, error) {
	list, err := git(nil, "rev-list", "--objects", "--all")
	if err != nil {
		return nil, err
	}
	names := map[string]string{}
	var input bytes.Buffer
	for _, line := range strings.Split(strings.TrimSpace(string(list)), "\n") {
		sha, name, _ := strings.Cut(line, " ")
		if sha == "" || names[sha] != "" {
			continue
		}
		names[sha] = name + "@" + sha[:min(12, len(sha))]
		input.WriteString(sha + "\n")
	}
	out, err := git(input.Bytes(), "cat-file", "--batch")
	if err != nil {
		return nil, err
	}
	return s.scanBatch(out, names)
}

// scanBatch walks the output of git cat-file --batch: "<sha> <type> <size>\n<content>\n".
func (s *Scanner) scanBatch(out []byte, names map[string]string) ([]Hit, error) {
	r := bufio.NewReader(bytes.NewReader(out))
	var hits []Hit
	for {
		header, err := r.ReadString('\n')
		if err == io.EOF && header == "" {
			return hits, nil
		}
		fields := strings.Fields(header)
		if err != nil || len(fields) != 3 {
			return nil, fmt.Errorf("git cat-file: unexpected header %q", header)
		}
		size, err := strconv.Atoi(fields[2])
		if err != nil || size < 0 || size > maxFileSize {
			return nil, fmt.Errorf("git cat-file: bad size in %q", header)
		}
		content := make([]byte, size+1)
		if _, err := io.ReadFull(r, content); err != nil {
			return nil, fmt.Errorf("git cat-file: short object %s: %w", fields[0], err)
		}
		if fields[1] != "blob" {
			continue
		}
		h, err := s.Scan(names[fields[0]], bytes.NewReader(content[:size]))
		if err != nil {
			return nil, err
		}
		hits = append(hits, h...)
	}
}

func (s *Scanner) scanMessages(git gitRunner) ([]Hit, error) {
	out, err := git(nil, "log", "--all", "--format=%H%x00%B%x00")
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(out), "\x00")
	var hits []Hit
	for i := 0; i+1 < len(parts); i += 2 {
		sha := strings.TrimSpace(parts[i])
		h, err := s.Scan("commit message "+sha[:min(12, len(sha))], strings.NewReader(parts[i+1]))
		if err != nil {
			return nil, err
		}
		hits = append(hits, h...)
	}
	return hits, nil
}

// checkIdentities fails every author or committer address that is neither a GitHub noreply
// address nor accepted in the private list. Addresses are reported masked.
func checkIdentities(git gitRunner, accepted map[string]bool) ([]Hit, error) {
	out, err := git(nil, "log", "--all", "--format=%H%x00%ae%x00%ce")
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	var hits []Hit
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		f := strings.Split(line, "\x00")
		if len(f) != 3 {
			continue
		}
		for _, addr := range f[1:] {
			lower := strings.ToLower(addr)
			if seen[lower] || accepted[lower] || strings.HasSuffix(lower, "@users.noreply.github.com") {
				continue
			}
			seen[lower] = true
			hits = append(hits, Hit{Path: "commit identity " + f[0][:min(12, len(f[0]))], Line: 1, Col: 1,
				Category: catEmail, value: []byte(addr)})
		}
	}
	return hits, nil
}
