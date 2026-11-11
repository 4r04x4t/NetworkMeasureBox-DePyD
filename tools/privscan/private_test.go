package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// privateValueForTest is a made-up value that stands for one of the owner's.
const privateValueForTest = "Zebra-Fjord-Quokka"

func TestPrivateValueNeverPrinted(t *testing.T) {
	repo, outside := t.TempDir(), t.TempDir()
	list := writeFile(t, outside, "private.txt", "# owner's values\n\n"+privateValueForTest+"\n")
	target := writeFile(t, repo, "notes.txt", "nothing here\nsee "+strings.ToLower(privateValueForTest)+" twice\n")
	e, out, errOut := testEnv(map[string]string{privateEnv: list})
	code := run([]string{"-repo", repo, "tree", target}, e)
	all := out.String() + errOut.String()
	if code != exitHits || !strings.Contains(all, "notes.txt:2:5: private: private value #3") {
		t.Errorf("exit %d, output %q", code, all)
	}
	if strings.Contains(strings.ToLower(all), strings.ToLower(privateValueForTest[:6])) {
		t.Errorf("output repeats the private value: %q", all)
	}
}

func TestPrivateListOutsideRepoOnly(t *testing.T) {
	repo := t.TempDir()
	inside := writeFile(t, repo, "private.txt", privateValueForTest+"\n")
	if _, err := LoadPrivate(inside, repo); err == nil || !strings.Contains(err.Error(), "inside the repository") {
		t.Errorf("list inside the repository: err %v", err)
	}
	outside := t.TempDir()
	link := filepath.Join(outside, "link.txt")
	if err := os.Symlink(inside, link); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadPrivate(link, repo); err == nil {
		t.Errorf("symbolic link into the repository accepted")
	}
	if _, err := LoadPrivate(outside, repo); err == nil {
		t.Errorf("a folder accepted as the list")
	}
	good := writeFile(t, outside, "private.txt", privateValueForTest+"\n")
	if list, err := LoadPrivate(good, repo); err != nil || len(list.values) != 1 {
		t.Errorf("good list: %v %v", list, err)
	}
}

func TestShortPrivateValueRefused(t *testing.T) {
	list := writeFile(t, t.TempDir(), "private.txt", privateValueForTest+"\nabc\n")
	_, err := LoadPrivate(list, t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "line 2") || strings.Contains(err.Error(), "abc") {
		t.Errorf("err %v", err)
	}
}

func TestRequirePrivate(t *testing.T) {
	path := writeFile(t, t.TempDir(), "f", "clean\n")
	e, _, _ := testEnv(nil)
	if code := run([]string{"-require-private", "tree", path}, e); code != exitError {
		t.Errorf("exit %d", code)
	}
	e, out, _ := testEnv(nil)
	if code := run([]string{"tree", path}, e); code != exitClean || !strings.Contains(out.String(), "private list: not used") {
		t.Errorf("exit %d, output %q", code, out)
	}
}

func TestUsageErrors(t *testing.T) {
	for _, args := range [][]string{{}, {"nosuchmode"}, {"dist"}, {"-bogus", "tree"}, {"-allow", "/nonexistent/allow", "tree"}} {
		e, _, _ := testEnv(nil)
		if code := run(args, e); code != exitError {
			t.Errorf("%q: exit %d", args, code)
		}
	}
}

func TestDistModeScansBinaries(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "linux-arm64"), 0o700); err != nil {
		t.Fatal(err)
	}
	writeFile(t, dir, "linux-arm64/nmbd", "\x7fELF\x00"+j("/ho", "me/", "builder", "/x")+"\x00")
	writeFile(t, dir, "clean", "\x7fELF\x00\x00")
	e, out, _ := testEnv(nil)
	if code := run([]string{"dist", dir}, e); code != exitHits || !strings.Contains(out.String(), "2 inputs, 1 hits") {
		t.Errorf("exit %d, output %q", code, out)
	}
}
