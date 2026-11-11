package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func testEnv(getenv map[string]string) (env, *bytes.Buffer, *bytes.Buffer) {
	var out, errOut bytes.Buffer
	return env{stdin: strings.NewReader(""), stdout: &out, stderr: &errOut,
		getenv: func(k string) string { return getenv[k] }, gitFrom: execGit}, &out, &errOut
}

// TestOwnSourcesClean: the scanner's own package, with its allow list, has no hit.
func TestOwnSourcesClean(t *testing.T) {
	files, err := filepath.Glob("*")
	if err != nil || len(files) < 5 {
		t.Fatalf("glob: %v %v", files, err)
	}
	e, out, errOut := testEnv(nil)
	code := run(append([]string{"-allow", "allow.txt", "tree"}, files...), e)
	if code != exitClean {
		t.Errorf("exit %d\n%s%s", code, out, errOut)
	}
}

func TestTreeModeHitExit1(t *testing.T) {
	dir := t.TempDir()
	clean := writeFile(t, dir, "clean.txt", "user@example.com 192.0.2.1\n")
	dirty := writeFile(t, dir, "dirty.txt", "line one\n"+j("a.person", "@", "mailhost", ".net")+"\n")
	e, out, _ := testEnv(nil)
	e.stdin = strings.NewReader(clean + "\x00" + dirty + "\x00")
	if code := run([]string{"tree", "-0"}, e); code != exitHits || !strings.Contains(out.String(), "dirty.txt:2:1: email: a.…(") {
		t.Errorf("exit %d, output %q", code, out)
	}
	e, _, _ = testEnv(nil)
	if code := run([]string{"tree", clean}, e); code != exitClean {
		t.Errorf("clean file: exit %d", code)
	}
}

func TestOutputMasked(t *testing.T) {
	secret := j("gh", "p_", strings.Repeat("Zq8kLw2mR", 4))
	path := writeFile(t, t.TempDir(), "f", "token: "+secret+"\n")
	e, out, errOut := testEnv(nil)
	code := run([]string{"tree", path}, e)
	if code != exitHits || strings.Contains(out.String()+errOut.String(), secret[2:]) {
		t.Errorf("exit %d, value not masked: %q", code, out)
	}
	if !strings.Contains(out.String(), "key: gh…(40)") {
		t.Errorf("masked form missing: %q", out)
	}
}

func TestLongLineScanned(t *testing.T) {
	line := strings.Repeat("x", 200_000) + " " + j("11.22", ".33.44") + "\n"
	hits := scanText(t, NewScanner(nil, nil), "first\n"+line)
	if len(hits) != 1 || hits[0].Line != 2 || hits[0].Col != 200_002 {
		t.Errorf("hits: %v", hits)
	}
}

func TestBinaryInputScanned(t *testing.T) {
	data := []byte("\x00\xff\x7fELF" + j("/ho", "me/", "builder", "/go") + "\x00\x01")
	hits, err := NewScanner(nil, nil).Scan("bin", bytes.NewReader(data))
	if err != nil || categories(hits) != catPath || !strings.Contains(hits[0].String(), "/h…(") {
		t.Errorf("hits %v, err %v", hits, err)
	}
}

func TestAllowEntryExactOnly(t *testing.T) {
	allow, err := LoadAllow([]byte(j("ipv4 ", "11.22", ".33.44", " a public resolver for a test\n")))
	if err != nil {
		t.Fatal(err)
	}
	sc := NewScanner(allow, nil)
	if hits := scanText(t, sc, j("11.22", ".33.44")); len(hits) != 0 {
		t.Errorf("allowed value reported: %v", hits)
	}
	if hits := scanText(t, sc, j("11.22", ".33.45 ", "11.22", ".33.4")); len(hits) != 2 {
		t.Errorf("other values must stay hits: %v", hits)
	}
}

func TestAllowEntryNeedsReason(t *testing.T) {
	for _, entry := range []string{"ipv4 100.64.0.0", "ipv4", "nosuchcategory 100.64.0.0 reason"} {
		if _, err := LoadAllow([]byte("# comment\n\n" + entry + "\n")); err == nil {
			t.Errorf("%q accepted", entry)
		}
	}
	if _, err := LoadAllow([]byte("# comment\nipv4 100.64.0.0 the range in the docs\n")); err != nil {
		t.Errorf("good entry refused: %v", err)
	}
}

// FuzzScan: arbitrary input never panics, hit columns lie inside their line, output is masked.
func FuzzScan(f *testing.F) {
	f.Add([]byte(j("a@b", ".net 11.22", ".33.44 fe80::", "0200:5eff:fe00:5301\n")))
	f.Add([]byte("::ffff:192.0.2.1 02:00:00:00:00:01 /home/nmb/\x00\xff"))
	f.Add([]byte(j(`password = "`, "Zq8#Lw2!Rt5x", `"`)))
	sc := NewScanner(nil, []privateValue{{lower: []byte("zz-private"), line: 1}})
	f.Fuzz(func(t *testing.T, data []byte) {
		hits, err := sc.Scan("f", bytes.NewReader(data))
		if err != nil {
			t.Fatalf("Scan: %v", err)
		}
		lines := bytes.Split(data, []byte("\n"))
		for _, h := range hits {
			if h.Line < 1 || h.Line > len(lines) || h.Col < 1 || h.Col > len(lines[h.Line-1])+1 {
				t.Fatalf("hit outside the input: %+v", h)
			}
			if len(h.value) > maskedPrefix+2 && strings.Contains(h.String(), string(h.value)) {
				t.Fatalf("unmasked output %q", h.String())
			}
		}
	})
}
