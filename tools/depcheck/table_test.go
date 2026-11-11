package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const goodTable = "# Deps\n\n" +
	"| Name | Kind | Module or package | Version | Licence (SPDX) | Origin | Use | Licence file SHA-256 | Added by | Checked |\n" +
	"| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |\n" +
	"| Go | toolchain | `go` | `go1.27.4` | BSD-3-Clause | https://go.dev/dl/ | built in | `abababababababababababababababababababababababababababababababab` | F-00 | 2026-11-11 |\n" +
	"| git | system | git | 2.43.0 | GPL-2.0-only | https://git-scm.com | build only | - | F-00 | 2026-11-11 |\n" +
	"\nNotes.\n"

func TestParseGoodTable(t *testing.T) {
	rows, err := ParseTable(goodTable)
	if err != nil || len(rows) != 2 {
		t.Fatalf("rows %v, err %v", rows, err)
	}
	if rows[0].Module != "go" || rows[0].Version != "go1.27.4" || !rows[0].permitted || rows[1].permitted || rows[1].Line != 6 {
		t.Errorf("rows %+v", rows)
	}
}

func TestUnknownLicenceFails(t *testing.T) {
	doc := strings.Replace(goodTable, "GPL-2.0-only", "Proprietary-1.0", 1)
	if _, err := ParseTable(doc); err == nil || !strings.Contains(err.Error(), "not a known SPDX id") {
		t.Errorf("err %v", err)
	}
}

func TestMalformedTableFails(t *testing.T) {
	header := "| Name | Kind | Module or package | Version | Licence (SPDX) | Origin | Use | Licence file SHA-256 | Added by | Checked |"
	cases := map[string]string{
		"no table":         "# nothing\n",
		"missing cell":     strings.Replace(goodTable, "| F-00 | 2026-11-11 |\n\nNotes", "| 2026-11-11 |\n\nNotes", 1),
		"renamed column":   strings.Replace(goodTable, "| Use |", "| Usage |", 1),
		"no separator":     strings.Replace(goodTable, "| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |\n", "", 1),
		"second table":     goodTable + goodTable,
		"bad kind":         strings.Replace(goodTable, "| system |", "| library |", 1),
		"bad version":      strings.Replace(goodTable, "`go1.27.4`", "`1.27`", 1),
		"bad hash":         strings.Replace(goodTable, strings.Repeat("ab", 32), "abc", 1),
		"hash for system":  strings.Replace(goodTable, "| build only | - |", "| build only | "+strings.Repeat("ab", 32)+" |", 1),
		"bad use":          strings.Replace(goodTable, "| build only |", "| sometimes |", 1),
		"bad date":         strings.Replace(goodTable, "| F-00 | 2026-11-11 |\n\n", "| F-00 | yesterday |\n\n", 1),
		"bad package id":   strings.Replace(goodTable, "| F-00 | 2026-11-11 |\n| git", "| me | 2026-11-11 |\n| git", 1),
		"plain http":       strings.Replace(goodTable, "https://git-scm.com", "http://git-scm.com", 1),
		"no rows":          header + "\n| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |\n",
		"module with pipe": strings.Replace(goodTable, "| `go` |", "| `go` | x |", 1),
	}
	for name, doc := range cases {
		if _, err := ParseTable(doc); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}

// FuzzParseDependencies: any document is parsed or refused, never a panic, and every accepted
// row has a known kind, use and licence.
func FuzzParseDependencies(f *testing.F) {
	f.Add(goodTable)
	f.Add("| a | b |\n| --- | --- |\n")
	f.Add(strings.Replace(goodTable, "built in", "test only", 1))
	f.Fuzz(func(t *testing.T, doc string) {
		rows, err := ParseTable(doc)
		if err != nil {
			return
		}
		for _, r := range rows {
			if kindVersion(r.Kind, newValidators()) == nil || (r.Use != useBuiltIn && r.Use != useTestOnly && r.Use != useBuildOnly) {
				t.Fatalf("accepted row %+v", r)
			}
			if _, ok := licences()[r.Licence]; !ok {
				t.Fatalf("accepted licence %q", r.Licence)
			}
		}
	})
}

// fakeGo answers the go commands depcheck runs with recorded output.
func fakeGo(t *testing.T, goroot string, modJSON string) goRunner {
	return func(dir string, args ...string) ([]byte, error) {
		switch strings.Join(args, " ") {
		case "env -json GOROOT GOVERSION GOMODCACHE":
			return []byte(`{"GOROOT":"` + goroot + `","GOVERSION":"go1.27.4","GOMODCACHE":"` + t.TempDir() + `"}`), nil
		case "mod edit -json":
			return []byte(modJSON), nil
		}
		return nil, errors.New("unexpected go " + strings.Join(args, " "))
	}
}

func TestRunWithRecordedGoOutput(t *testing.T) {
	goroot := t.TempDir()
	if err := os.WriteFile(filepath.Join(goroot, "LICENSE"), []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	list := filepath.Join(t.TempDir(), "deps.md")
	doc := strings.Replace(goodTable, strings.Repeat("ab", 32), "2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881", 1)
	if err := os.WriteFile(list, []byte(doc), 0o600); err != nil {
		t.Fatal(err)
	}
	clean := `{"Module":{"Path":"nmb"},"Go":"1.27.0","Toolchain":"go1.27.4"}`
	var out, errOut bytes.Buffer
	if code := run([]string{"-list", list, "check", "go.mod"}, fakeGo(t, goroot, clean), &out, &errOut); code != exitOK {
		t.Errorf("clean: exit %d\n%s%s", code, out.String(), errOut.String())
	}
	unlisted := `{"Module":{"Path":"nmb"},"Toolchain":"go1.27.4","Require":[{"Path":"example.com/x","Version":"v1.0.0"}],` +
		`"Replace":[{"Old":{"Path":"example.com/x"},"New":{"Path":"./x"}}]}`
	out.Reset()
	if code := run([]string{"-list", list, "check", "go.mod"}, fakeGo(t, goroot, unlisted), &out, &errOut); code != exitProblems ||
		!strings.Contains(out.String(), "example.com/x v1.0.0 is not in docs/DEPENDENCIES.md") {
		t.Errorf("unlisted: exit %d\n%s", code, out.String())
	}
}

func TestGateToolsListing(t *testing.T) {
	list := filepath.Join(t.TempDir(), "deps.md")
	row := "| gocyclo | gate-tool | `example.com/gocyclo/cmd/gocyclo` | `v0.6.0` | BSD-3-Clause | https://example.com | build only | `" +
		strings.Repeat("cd", 32) + "` | F-00 | 2026-11-11 |\n"
	doc := strings.Replace(goodTable, "\n\nNotes", "\n"+strings.TrimSuffix(row, "\n")+"\n\nNotes", 1)
	if err := os.WriteFile(list, []byte(doc), 0o600); err != nil {
		t.Fatal(err)
	}
	var out, errOut bytes.Buffer
	code := run([]string{"-list", list, "gate-tools"}, nil, &out, &errOut)
	if code != exitOK || out.String() != "gocyclo\texample.com/gocyclo/cmd/gocyclo\tv0.6.0\n" {
		t.Errorf("exit %d, output %q %q", code, out.String(), errOut.String())
	}
}

func TestParseVersionM(t *testing.T) {
	out := "/x/staticcheck: go1.27.4\n\tpath\thonnef.co/go/tools/cmd/staticcheck\n\tmod\thonnef.co/go/tools\tv0.8.2\th1:x=\n\tbuild\tCGO_ENABLED=0\n"
	tb, err := parseVersionM(out)
	if err != nil || tb != (ToolBuild{"go1.27.4", "honnef.co/go/tools/cmd/staticcheck", "honnef.co/go/tools", "v0.8.2"}) {
		t.Errorf("%+v %v", tb, err)
	}
	if _, err := parseVersionM("/x/y: go1.27.4\n"); err == nil {
		t.Errorf("no build information accepted")
	}
}
