// Command depcheck checks the dependency and licence list docs/DEPENDENCIES.md against the code.
//
// Usage:
//
//	depcheck -list FILE [-tools DIR] check GOMOD...   check the list against these go.mod files
//	depcheck -list FILE gate-tools                     print "name<TAB>package<TAB>version" per gate tool
//
// check fails when a go.mod requires a module that is not listed or listed with another version,
// a module row is in no go.mod, a licence is unknown or not allowed for a part built into nmbd,
// a licence file is missing or its SHA-256 differs, the toolchain row differs from a go.mod or
// from the running go command, or an installed gate tool is not the listed version built with the
// listed toolchain. Gate tools live in DIR/<toolchain>/<name>@<version>/<name> (tools/lib.sh).
// Exit status: 0 complete and correct, 1 problems found, 2 usage or read error.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	exitOK       = 0
	exitProblems = 1
	exitError    = 2
)

// goRunner runs the go command in dir with the argument list and returns its stdout.
type goRunner func(dir string, args ...string) ([]byte, error)

func execGo(dir string, args ...string) ([]byte, error) {
	cmd := exec.Command("go", args...) // #nosec G204 -- fixed program, argument list, no shell
	cmd.Dir = dir
	var stderr strings.Builder
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("go %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return out, nil
}

func main() {
	os.Exit(run(os.Args[1:], execGo, os.Stdout, os.Stderr))
}

func run(args []string, goCmd goRunner, stdout, stderr io.Writer) int {
	fl := flag.NewFlagSet("depcheck", flag.ContinueOnError)
	fl.SetOutput(io.Discard)
	list := fl.String("list", "", "dependency list (Markdown)")
	tools := fl.String("tools", "", "folder of the gate tools")
	if err := fl.Parse(args); err != nil || *list == "" || fl.NArg() == 0 {
		_, _ = fmt.Fprintln(stderr, "usage: depcheck -list FILE [-tools DIR] check GOMOD... | gate-tools")
		return exitError
	}
	doc, err := os.ReadFile(filepath.Clean(*list))
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "depcheck: %v\n", err)
		return exitError
	}
	rows, err := ParseTable(string(doc))
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "depcheck: %s: %v\n", *list, err)
		return exitProblems
	}
	switch fl.Arg(0) {
	case "gate-tools":
		for _, r := range rows {
			if r.Kind == kindGateTool {
				_, _ = fmt.Fprintf(stdout, "%s\t%s\t%s\n", r.Name, r.Module, r.Version)
			}
		}
		return exitOK
	case "check":
		return runCheck(rows, fl.Args()[1:], *tools, goCmd, stdout, stderr)
	}
	_, _ = fmt.Fprintf(stderr, "depcheck: unknown command %q\n", fl.Arg(0))
	return exitError
}

func runCheck(rows []Row, gomods []string, tools string, goCmd goRunner, stdout, stderr io.Writer) int {
	if len(gomods) == 0 || (tools != "" && !filepath.IsAbs(tools)) {
		_, _ = fmt.Fprintln(stderr, "depcheck: check needs go.mod files and an absolute -tools folder")
		return exitError
	}
	env, err := loadEnv(rows, gomods, tools, goCmd)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "depcheck: %v\n", err)
		return exitError
	}
	problems := Check(env)
	for _, p := range problems {
		_, _ = fmt.Fprintln(stdout, p)
	}
	for _, r := range rows {
		if r.Hash == noHash {
			_, _ = fmt.Fprintf(stdout, "note: line %d: %s: licence file not checked (kind %s)\n", r.Line, r.Name, r.Kind)
		}
	}
	_, _ = fmt.Fprintf(stdout, "depcheck: %d rows, %d go.mod files, %d problems\n", len(rows), len(gomods), len(problems))
	if len(problems) > 0 {
		return exitProblems
	}
	return exitOK
}

func loadEnv(rows []Row, gomods []string, tools string, goCmd goRunner) (Env, error) {
	repo, err := filepath.Abs(".")
	if err != nil {
		return Env{}, err
	}
	goEnv, err := loadGoEnv(goCmd, repo)
	if err != nil {
		return Env{}, err
	}
	env := Env{Rows: rows, GoRoot: goEnv.GOROOT, GoVersion: goEnv.GOVERSION,
		Roots:     []string{goEnv.GOMODCACHE, goEnv.GOROOT, repo},
		ModuleDir: func(path, version string) (string, error) { return downloadDir(goCmd, repo, path, version) },
		Tool: func(name, version string) (ToolBuild, error) {
			return toolBuild(goCmd, repo, filepath.Join(tools, goEnv.GOVERSION, name+"@"+version, name))
		},
	}
	if tools == "" {
		env.Tool = func(string, string) (ToolBuild, error) { return ToolBuild{}, errors.New("no -tools folder given") }
	}
	for _, gomod := range gomods {
		m, err := loadModFile(goCmd, repo, gomod)
		if err != nil {
			return Env{}, err
		}
		if gomod == "go.mod" {
			env.MainModule = m.Module
		}
		env.ModFiles = append(env.ModFiles, m)
	}
	return env, nil
}
