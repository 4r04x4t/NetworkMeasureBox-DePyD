package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

// funcWithBody returns a Go file whose function f has exactly n lines between its braces.
func funcWithBody(doc string, n int) string {
	var b strings.Builder
	b.WriteString("package p\n\n" + doc + "func f() {\n")
	for i := 0; i < n; i++ {
		b.WriteString("\t_ = 1\n")
	}
	b.WriteString("}\n")
	return b.String()
}

func check(t *testing.T, src string) Result {
	t.Helper()
	res, err := CheckGo("p.go", []byte(src))
	if err != nil {
		t.Fatalf("CheckGo: %v", err)
	}
	return res
}

func rules(fs []Finding) string {
	var out []string
	for _, f := range fs {
		out = append(out, f.Rule)
	}
	return strings.Join(out, ",")
}

func TestFunctionLength50Passes(t *testing.T) {
	if res := check(t, funcWithBody("", 50)); len(res.Violations) != 0 {
		t.Errorf("violations: %v", res.Violations)
	}
}

func TestFunctionLength51Fails(t *testing.T) {
	res := check(t, funcWithBody("", 51))
	if len(res.Violations) != 1 || res.Violations[0].Value != 51 || !strings.HasPrefix(res.Violations[0].Rule, "function length f") {
		t.Errorf("violations: %v", res.Violations)
	}
}

func TestPackageLevelFuncLitLength(t *testing.T) {
	src := strings.Replace(funcWithBody("", 51), "func f() {", "var f = func() {", 1)
	if res := check(t, src); len(res.Violations) != 1 {
		t.Errorf("violations: %v", res.Violations)
	}
}

func fileOfLines(n int) string {
	return "package p\n" + strings.Repeat("\n", n-1)
}

func TestFileLength500Passes(t *testing.T) {
	if res := check(t, fileOfLines(500)); len(res.Violations) != 0 {
		t.Errorf("violations: %v", res.Violations)
	}
	if res := CheckText("x.sh", []byte(strings.Repeat("x\n", 500))); len(res.Violations) != 0 {
		t.Errorf("shell violations: %v", res.Violations)
	}
}

func TestFileLength501Fails(t *testing.T) {
	if res := check(t, fileOfLines(501)); rules(res.Violations) != "file length" {
		t.Errorf("violations: %v", res.Violations)
	}
	if res := CheckText("x.sh", []byte(strings.Repeat("x\n", 501))); rules(res.Violations) != "file length" {
		t.Errorf("shell violations: %v", res.Violations)
	}
}

func TestAllowLongFileListed(t *testing.T) {
	res := check(t, "//nmb:allow-long-file generated table\n"+fileOfLines(600))
	if len(res.Violations) != 0 || len(res.Exceptions) != 1 || res.Exceptions[0].Reason != "generated table" {
		t.Errorf("result: %+v", res)
	}
	sh := "#!/bin/sh\n# nmb:allow-long-file one long table\n" + strings.Repeat("x\n", 600)
	if res := CheckText("x.sh", []byte(sh)); len(res.Violations) != 0 || len(res.Exceptions) != 1 {
		t.Errorf("shell result: %+v", res)
	}
}

func TestAllowLongNeedsReason(t *testing.T) {
	for _, doc := range []string{"//nmb:allow-long\n", "//nmb:allow-long   \n"} {
		res := check(t, funcWithBody(doc, 60))
		if len(res.Violations) != 1 || len(res.Exceptions) != 0 || !strings.Contains(res.Violations[0].Rule, "without a reason") {
			t.Errorf("%q: %+v", doc, res)
		}
	}
	if res := check(t, "//nmb:allow-long-file\n"+fileOfLines(600)); len(res.Violations) != 1 {
		t.Errorf("file directive without reason: %+v", res)
	}
}

func TestAllowLongListed(t *testing.T) {
	res := check(t, funcWithBody("// f is long.\n//nmb:allow-long one table of all rules\n", 60))
	if len(res.Violations) != 0 || len(res.Exceptions) != 1 || res.Exceptions[0].Reason != "one table of all rules" {
		t.Errorf("result: %+v", res)
	}
	var out bytes.Buffer
	read := func(string) ([]byte, error) { return []byte(funcWithBody("//nmb:allow-long reason\n", 60)), nil }
	if code := run([]string{"p.go"}, read, &out, &out); code != 0 || !strings.Contains(out.String(), "exception: p.go:4:") {
		t.Errorf("run: code %d, output %q", code, out.String())
	}
}

func TestOtherDirectiveIsNoException(t *testing.T) {
	if res := check(t, funcWithBody("//nmb:allow-longer reason\n", 60)); len(res.Violations) != 1 {
		t.Errorf("result: %+v", res)
	}
}

func TestParseErrorIsExit2(t *testing.T) {
	var out, errOut bytes.Buffer
	read := func(string) ([]byte, error) { return []byte("package p\nfunc {"), nil }
	if code := run([]string{"bad.go"}, read, &out, &errOut); code != 2 || !strings.Contains(errOut.String(), "parsing bad.go") {
		t.Errorf("code %d, stderr %q", code, errOut.String())
	}
	missing := func(string) ([]byte, error) { return nil, errors.New("no such file") }
	if code := run([]string{"gone.go"}, missing, &out, &errOut); code != 2 {
		t.Errorf("unreadable file: code %d", code)
	}
	if code := run(nil, missing, &out, &errOut); code != 2 {
		t.Errorf("no arguments: code %d", code)
	}
}

func TestViolationIsExit1(t *testing.T) {
	var out bytes.Buffer
	read := func(string) ([]byte, error) { return []byte(funcWithBody("", 51)), nil }
	if code := run([]string{"p.go"}, read, &out, &out); code != 1 || !strings.Contains(out.String(), "p.go:3: function length f: 51 (limit 50)") {
		t.Errorf("code %d, output %q", code, out.String())
	}
}

// FuzzCheckSource: arbitrary bytes never panic; unparsable input is an error, never a pass.
func FuzzCheckSource(f *testing.F) {
	f.Add([]byte(funcWithBody("//nmb:allow-long r\n", 3)))
	f.Add([]byte(nested(5)))
	f.Add([]byte("package p\nvar f = func() { for { select {} } }\n"))
	f.Add([]byte("package"))
	f.Fuzz(func(t *testing.T, src []byte) {
		res, err := CheckGo("f.go", src)
		if err != nil && (len(res.Violations) != 0 || len(res.Exceptions) != 0) {
			t.Fatalf("error with findings: %v", err)
		}
		_ = CheckText("f.sh", src)
	})
}
