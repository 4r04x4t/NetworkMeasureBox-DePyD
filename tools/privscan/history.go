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

// historyStats counts what a history scan read, for its summary line.
type historyStats struct{ refs, commits, blobs int }

func (st historyStats) String() string {
	return fmt.Sprintf("%d refs, %d commits, %d blobs", st.refs, st.commits, st.blobs)
}

// ScanHistory scans every blob reachable from any ref, every commit message, and checks the
// author and committer addresses of every commit.
func (s *Scanner) ScanHistory(git gitRunner, accepted map[string]bool) ([]Hit, historyStats, error) {
	var st historyStats
	refs, err := git(nil, "for-each-ref", "--format=%(refname)")
	if err != nil {
		return nil, st, err
	}
	st.refs = len(strings.Fields(string(refs)))
	hits, blobs, err := s.scanBlobs(git)
	if err != nil {
		return nil, st, err
	}
	msgHits, commits, err := s.scanMessages(git)
	if err != nil {
		return nil, st, err
	}
	idHits, err := checkIdentities(git, accepted)
	if err != nil {
		return nil, st, err
	}
	st.blobs, st.commits = blobs, commits
	return append(append(hits, msgHits...), idHits...), st, nil
}

// scanBlobs scans every blob reachable from any ref and returns the hits and the blob count.
func (s *Scanner) scanBlobs(git gitRunner) ([]Hit, int, error) {
	list, err := git(nil, "rev-list", "--objects", "--all")
	if err != nil {
		return nil, 0, err
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
		return nil, 0, err
	}
	return s.scanBatch(out, names)
}

// scanBatch walks the output of git cat-file --batch: "<sha> <type> <size>\n<content>\n".
func (s *Scanner) scanBatch(out []byte, names map[string]string) ([]Hit, int, error) {
	r := bufio.NewReader(bytes.NewReader(out))
	var hits []Hit
	blobs := 0
	for {
		header, err := r.ReadString('\n')
		if err == io.EOF && header == "" {
			return hits, blobs, nil
		}
		fields := strings.Fields(header)
		if err != nil || len(fields) != 3 {
			return nil, 0, fmt.Errorf("git cat-file: unexpected header %q", header)
		}
		size, err := strconv.Atoi(fields[2])
		if err != nil || size < 0 || size > maxFileSize {
			return nil, 0, fmt.Errorf("git cat-file: bad size in %q", header)
		}
		content := make([]byte, size+1)
		if _, err := io.ReadFull(r, content); err != nil {
			return nil, 0, fmt.Errorf("git cat-file: short object %s: %w", fields[0], err)
		}
		if fields[1] != "blob" {
			continue
		}
		h, err := s.Scan(names[fields[0]], bytes.NewReader(content[:size]))
		if err != nil {
			return nil, 0, err
		}
		hits = append(hits, h...)
		blobs++
	}
}

// scanMessages scans every commit message of all refs and returns the hits and the commit count.
func (s *Scanner) scanMessages(git gitRunner) ([]Hit, int, error) {
	out, err := git(nil, "log", "--all", "--format=%H%x00%B%x00")
	if err != nil {
		return nil, 0, err
	}
	parts := strings.Split(string(out), "\x00")
	var hits []Hit
	commits := 0
	for i := 0; i+1 < len(parts); i += 2 {
		sha := strings.TrimSpace(parts[i])
		h, err := s.Scan("commit message "+sha[:min(12, len(sha))], strings.NewReader(parts[i+1]))
		if err != nil {
			return nil, 0, err
		}
		hits = append(hits, h...)
		commits++
	}
	return hits, commits, nil
}

// checkIdentities fails every author or committer address that is neither a GitHub noreply
// address nor accepted in the private list, and every author or committer name that is neither
// the account name its own noreply address already shows nor accepted there. Values are masked.
func checkIdentities(git gitRunner, accepted map[string]bool) ([]Hit, error) {
	out, err := git(nil, "log", "--all", "--format=%H%x00%an%x00%ae%x00%cn%x00%ce")
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	var hits []Hit
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		f := strings.Split(line, "\x00")
		if len(f) != 5 {
			continue
		}
		where := "commit identity " + f[0][:min(12, len(f[0]))]
		for _, id := range []identity{{name: f[1], addr: f[2]}, {name: f[3], addr: f[4]}} {
			hits = append(hits, newHits(seen, where, id, id.problems(accepted))...)
		}
	}
	return hits, nil
}

// newHits returns the hits of id not reported before, each at where; seen remembers them.
func newHits(seen map[string]bool, where string, id identity, found []Hit) []Hit {
	var out []Hit
	for _, h := range found {
		key := h.Category + "\x00" + string(h.value) + "\x00" + id.addr
		if !seen[key] {
			seen[key] = true
			h.Path = where
			out = append(out, h)
		}
	}
	return out
}

// noreplyDomain is GitHub's address domain for accounts that keep their address private.
const noreplyDomain = "@users.noreply.github.com"

// identity is the name and address of one author or committer.
type identity struct{ name, addr string }

// problems returns the hits of one identity (Path is set by the caller).
func (id identity) problems(accepted map[string]bool) []Hit {
	var hits []Hit
	lower := strings.ToLower(id.addr)
	if !accepted[lower] && !strings.HasSuffix(lower, noreplyDomain) {
		hits = append(hits, Hit{Line: 1, Col: 1, Category: catEmail, value: []byte(id.addr)})
	}
	name := strings.ToLower(strings.TrimSpace(id.name))
	if name != "" && !accepted[name] && name != noreplyLogin(lower) {
		hits = append(hits, Hit{Line: 1, Col: 1, Category: catName, value: []byte(id.name)})
	}
	return hits
}

// noreplyLogin returns the account name of "<id>+<login>@users.noreply.github.com" or
// "<login>@users.noreply.github.com", or "" for any other address.
func noreplyLogin(addr string) string {
	local, ok := strings.CutSuffix(addr, noreplyDomain)
	if !ok {
		return ""
	}
	if _, login, found := strings.Cut(local, "+"); found {
		return login
	}
	return local
}
