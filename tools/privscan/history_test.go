package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// gitRepo makes a temporary repository whose commits use the given identity.
func gitRepo(t *testing.T, identity string) (string, func(args ...string)) {
	t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed")
	}
	dir := t.TempDir()
	t.Setenv("HOME", t.TempDir())
	t.Setenv("GIT_CONFIG_NOSYSTEM", "1")
	t.Setenv("GIT_CONFIG_GLOBAL", os.DevNull)
	for _, k := range []string{"GIT_AUTHOR_NAME", "GIT_COMMITTER_NAME"} {
		t.Setenv(k, "Test")
	}
	for _, k := range []string{"GIT_AUTHOR_EMAIL", "GIT_COMMITTER_EMAIL"} {
		t.Setenv(k, identity)
	}
	git := func(args ...string) {
		t.Helper()
		cmd := exec.Command("git", append([]string{"-C", dir}, args...)...) // #nosec G204 -- test helper, fixed program, argument list
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	git("init", "-q", "-b", "main")
	return dir, git
}

func noreplyIdentity() string { return j("12345+tester", "@", "users.noreply", ".github.com") }

// F00-R22: a value added and later removed is still found in the history.
func TestHistoryModeFindsDeletedFile(t *testing.T) {
	dir, git := gitRepo(t, noreplyIdentity())
	writeFile(t, dir, "notes.txt", "contact "+j("a.person", "@", "mailhost", ".net")+"\n")
	git("add", "notes.txt")
	git("commit", "-q", "-m", "add notes")
	git("rm", "-q", "notes.txt")
	git("commit", "-q", "-m", "remove notes")
	e, out, errOut := testEnv(nil)
	code := run([]string{"-repo", dir, "history"}, e)
	if code != exitHits || !strings.Contains(out.String(), "notes.txt@") || strings.Count(out.String(), "email:") != 1 {
		t.Errorf("exit %d\n%s%s", code, out, errOut)
	}
}

func TestHistoryCleanRepoPasses(t *testing.T) {
	dir, git := gitRepo(t, noreplyIdentity())
	writeFile(t, dir, "a.txt", "user@example.com 192.0.2.1\n")
	git("add", "a.txt")
	git("commit", "-q", "-m", "clean\n\nCo-Authored-By: Bot <"+j("noreply", "@", "vendor", ".com")+">")
	e, out, errOut := testEnv(nil)
	if code := run([]string{"-repo", dir, "history"}, e); code != exitClean {
		t.Errorf("exit %d\n%s%s", code, out, errOut)
	}
}

func TestHistoryMessageScanned(t *testing.T) {
	dir, git := gitRepo(t, noreplyIdentity())
	git("commit", "-q", "--allow-empty", "-m", "deployed to "+j("11.22", ".33.44"))
	e, out, _ := testEnv(nil)
	if code := run([]string{"-repo", dir, "history"}, e); code != exitHits || !strings.Contains(out.String(), "commit message") {
		t.Errorf("exit %d, output %q", code, out)
	}
}

func TestHistoryIdentityNotNoreplyFails(t *testing.T) {
	identity := "developer@example.com"
	dir, git := gitRepo(t, identity)
	git("commit", "-q", "--allow-empty", "-m", "first")
	e, out, _ := testEnv(nil)
	if code := run([]string{"-repo", dir, "history"}, e); code != exitHits || !strings.Contains(out.String(), "commit identity") ||
		strings.Contains(out.String(), identity) {
		t.Errorf("exit %d, output %q", code, out)
	}
	list := writeFile(t, t.TempDir(), "private.txt", identityPrefix+identity+"\n")
	e, out, _ = testEnv(map[string]string{privateEnv: list})
	if code := run([]string{"-repo", dir, "history"}, e); code != exitClean {
		t.Errorf("accepted identity: exit %d, output %q", code, out)
	}
}
