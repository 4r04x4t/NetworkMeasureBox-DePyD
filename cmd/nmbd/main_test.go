package main

import (
	"bytes"
	"errors"
	"os"
	"runtime/debug"
	"strings"
	"testing"
)

const (
	referenceFile = "../../docs/cli/nmbd.md"
	helpBegin     = "<!-- nmbd-help-begin -->\n```text\n"
	helpEnd       = "```\n<!-- nmbd-help-end -->"
)

func fakeInfo() (*debug.BuildInfo, bool) {
	return &debug.BuildInfo{
		GoVersion: "go1.27.4",
		Main:      debug.Module{Path: "nmb", Version: "v0.0.0-20261111000000-000000000000"},
		Settings:  []debug.BuildSetting{{Key: "GOOS", Value: "linux"}, {Key: "GOARCH", Value: "amd64"}},
	}, true
}

func call(args ...string) (code int, stdout, stderr string) {
	var out, errOut bytes.Buffer
	code = run(args, &out, &errOut, fakeInfo)
	return code, out.String(), errOut.String()
}

// F00-R09: --help prints the usage to stdout and exits 0.
func TestHelpPrintsUsageExit0(t *testing.T) {
	for _, arg := range []string{"--help", "-h", "-help"} {
		code, out, errOut := call(arg)
		if code != 0 || out != usage || errOut != "" {
			t.Errorf("%s: code %d, stdout %q, stderr %q", arg, code, out, errOut)
		}
	}
}

// F00-R14: the command reference holds exactly the --help output.
func TestHelpMatchesReference(t *testing.T) {
	doc, err := os.ReadFile(referenceFile)
	if err != nil {
		t.Fatalf("reading the reference: %v", err)
	}
	text := string(doc)
	start := strings.Index(text, helpBegin)
	end := strings.Index(text, helpEnd)
	if start < 0 || end < start {
		t.Fatalf("%s has no usage block between the markers", referenceFile)
	}
	block := text[start+len(helpBegin) : end]
	_, out, _ := call("--help")
	if block != out {
		t.Errorf("reference block differs from --help:\nreference:\n%s\n--help:\n%s", block, out)
	}
}

func TestVersionLines(t *testing.T) {
	code, out, errOut := call("--version")
	want := "version: v0.0.0-20261111000000-000000000000\nmodule: nmb\ngo: go1.27.4\nos: linux\n" +
		"arch: amd64\ncgo: not recorded\nrevision: not recorded\ncommit-time: not recorded\nmodified: not recorded\n"
	if code != 0 || out != want || errOut != "" {
		t.Errorf("code %d, stdout %q, stderr %q", code, out, errOut)
	}
}

func TestVersionWithoutBuildInfoFails(t *testing.T) {
	var out, errOut bytes.Buffer
	code := run([]string{"--version"}, &out, &errOut, func() (*debug.BuildInfo, bool) { return nil, false })
	if code != 1 || out.Len() != 0 || !strings.Contains(errOut.String(), "no build information") {
		t.Errorf("code %d, stdout %q, stderr %q", code, out.String(), errOut.String())
	}
}

func TestNoArgsUsageExit2(t *testing.T) {
	code, out, errOut := call()
	if code != 2 || out != "" || !strings.HasSuffix(errOut, usage) {
		t.Errorf("code %d, stdout %q, stderr %q", code, out, errOut)
	}
}

func TestUnknownOptionExit2(t *testing.T) {
	for _, args := range [][]string{{"--bogus"}, {"--help", "extra"}, {"start"}, {"--help=maybe"}} {
		code, out, errOut := call(args...)
		if code != 2 || out != "" || !strings.HasSuffix(errOut, usage) {
			t.Errorf("%q: code %d, stdout %q, stderr %q", args, code, out, errOut)
		}
	}
}

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) { return 0, errors.New("closed") }

func TestOutputErrorExit1(t *testing.T) {
	var errOut bytes.Buffer
	if code := run([]string{"--help"}, failingWriter{}, &errOut, fakeInfo); code != 1 {
		t.Errorf("code %d, want 1 (stderr %q)", code, errOut.String())
	}
}

// F00-R09 (roadmap Phase 0): --help changes nothing: HOME, TMPDIR and the working folder stay empty.
func TestHelpChangesNothing(t *testing.T) {
	home, tmp, cwd := t.TempDir(), t.TempDir(), t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("TMPDIR", tmp)
	t.Chdir(cwd)
	before := os.Environ()
	if code, _, _ := call("--help"); code != 0 {
		t.Fatalf("--help exit %d", code)
	}
	for _, dir := range []string{home, tmp, cwd} {
		entries, err := os.ReadDir(dir)
		if err != nil || len(entries) != 0 {
			t.Errorf("%s: %d entries, err %v", dir, len(entries), err)
		}
	}
	if strings.Join(os.Environ(), "\n") != strings.Join(before, "\n") {
		t.Errorf("environment changed")
	}
}

// FuzzRun: any arguments give exit 0 or 2 and never panic; output goes to one stream only.
func FuzzRun(f *testing.F) {
	for _, seed := range []string{"", "--help", "-h", "--version", "--bogus", "--help\x00x", "-version=false"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, joined string) {
		var args []string
		if joined != "" {
			args = strings.Split(joined, "\x00")
		}
		code, out, errOut := call(args...)
		switch code {
		case 0:
			if errOut != "" || out == "" {
				t.Fatalf("%q: exit 0 with stdout %q, stderr %q", args, out, errOut)
			}
		case 2:
			if out != "" || !strings.HasSuffix(errOut, usage) {
				t.Fatalf("%q: exit 2 with stdout %q, stderr %q", args, out, errOut)
			}
		default:
			t.Fatalf("%q: exit %d", args, code)
		}
	})
}
