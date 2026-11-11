// Command nmbd is the native NetworkMeasureBox program.
//
// This is the skeleton of package F-00: it prints its usage and its build information and does
// nothing else. It opens no file or socket and changes no setting (roadmap Phase 0, "--help").
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"nmb/internal/buildinfo"
)

// Exit codes; docs/cli/nmbd.md lists the same.
const (
	exitOK      = 0
	exitFailure = 1
	exitUsage   = 2
)

// usage is the --help text. docs/cli/nmbd.md carries the same text; a test keeps them equal.
const usage = `Usage: nmbd --help | --version

nmbd is the native NetworkMeasureBox program. This build is the skeleton:
there is no daemon yet and it does nothing besides printing this text and
its build information. It changes nothing and needs no root.

Options:
  -h, --help   print this text and exit
  --version    print the build information and exit

Exit status: 0 for --help and --version, 2 for any other use,
1 when the output cannot be written or the build information is missing.
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr, debug.ReadBuildInfo))
}

// run handles one invocation and returns the exit code. It only writes to stdout and stderr.
func run(args []string, stdout, stderr io.Writer, read buildinfo.ReadFunc) int {
	fs := flag.NewFlagSet("nmbd", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	help := fs.Bool("help", false, "print the usage and exit")
	short := fs.Bool("h", false, "print the usage and exit")
	version := fs.Bool("version", false, "print the build information and exit")
	if err := fs.Parse(args); err != nil {
		return usageError(stderr, err)
	}
	switch {
	case fs.NArg() > 0:
		return usageError(stderr, fmt.Errorf("unexpected argument %q", fs.Arg(0)))
	case *help || *short:
		return write(stdout, stderr, usage)
	case *version:
		return printVersion(stdout, stderr, read)
	default:
		return usageError(stderr, errors.New("no option given"))
	}
}

func usageError(stderr io.Writer, err error) int {
	_, _ = fmt.Fprintf(stderr, "nmbd: %v\n\n%s", err, usage)
	return exitUsage
}

func printVersion(stdout, stderr io.Writer, read buildinfo.ReadFunc) int {
	fields, err := buildinfo.Read(read)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "nmbd: %v\n", err)
		return exitFailure
	}
	text := ""
	for _, f := range fields {
		text += f.Key + ": " + f.Value + "\n"
	}
	return write(stdout, stderr, text)
}

func write(stdout, stderr io.Writer, text string) int {
	if _, err := io.WriteString(stdout, text); err != nil {
		_, _ = fmt.Fprintf(stderr, "nmbd: writing output: %v\n", err)
		return exitFailure
	}
	return exitOK
}
