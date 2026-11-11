// Command cleancode enforces the clean-code limits of the roadmap where a tool can: functions of
// at most 50 lines, files of at most 500 lines and nesting of at most 3 levels.
//
// Usage: cleancode FILE...
//
// Go files (*.go) get every rule; other files (shell scripts) get the file-length rule. A longer
// function or file needs a stated reason in a directive ("//nmb:allow-long <reason>" in the doc
// comment of the function, "//nmb:allow-long-file <reason>" before the package clause, or
// "# nmb:allow-long-file <reason>" in the first five lines of a script); every exception is
// printed. Nesting has no exception. Exit status: 0 clean, 1 violations, 2 usage or read/parse error.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	exitClean     = 0
	exitViolation = 1
	exitError     = 2
)

func main() {
	os.Exit(run(os.Args[1:], os.ReadFile, os.Stdout, os.Stderr))
}

// run checks every file and returns the exit status.
func run(paths []string, read func(string) ([]byte, error), stdout, stderr io.Writer) int {
	if len(paths) == 0 {
		_, _ = fmt.Fprintln(stderr, "usage: cleancode FILE...")
		return exitError
	}
	status := exitClean
	for _, path := range paths {
		res, err := checkPath(path, read)
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "cleancode: %v\n", err)
			status = exitError
			continue
		}
		for _, e := range res.Exceptions {
			_, _ = fmt.Fprintln(stdout, "exception: "+e.String())
		}
		for _, v := range res.Violations {
			_, _ = fmt.Fprintln(stdout, v.String())
			if status == exitClean {
				status = exitViolation
			}
		}
	}
	return status
}

func checkPath(path string, read func(string) ([]byte, error)) (Result, error) {
	src, err := read(path)
	if err != nil {
		return Result{}, fmt.Errorf("reading %s: %w", path, err)
	}
	if strings.HasSuffix(path, ".go") {
		return CheckGo(path, src)
	}
	return CheckText(path, src), nil
}
