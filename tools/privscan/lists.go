package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// minPrivateLen is the shortest private value accepted; shorter ones would match everywhere.
const minPrivateLen = 4

// identityPrefix marks a private-list line whose address is accepted as a commit identity.
const identityPrefix = "identity-ok "

// privateValue is one value of the owner's private list. Only its line number is ever printed.
type privateValue struct {
	lower []byte
	line  int
}

// privateList is the parsed private list.
type privateList struct {
	values     []privateValue
	identities map[string]bool // lower-case addresses accepted as commit identities
}

// LoadAllow parses the allow list: lines "<category> <exact value> <reason>"; blank lines and
// lines starting with # are comments. An entry without a reason or with an unknown category is
// an error.
func LoadAllow(data []byte) (map[string]map[string]bool, error) {
	known := map[string]bool{catEmail: true, catIPv4: true, catIPv6: true, catMAC: true, catKey: true, catPath: true}
	allow := map[string]map[string]bool{}
	for i, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 3 {
			return nil, fmt.Errorf("allow list line %d: want <category> <value> <reason>", i+1)
		}
		if !known[fields[0]] {
			return nil, fmt.Errorf("allow list line %d: unknown category %q", i+1, fields[0])
		}
		if allow[fields[0]] == nil {
			allow[fields[0]] = map[string]bool{}
		}
		allow[fields[0]][fields[1]] = true
	}
	return allow, nil
}

// LoadPrivate reads the owner's private list. It must be a regular file outside the repository;
// each value must have at least minPrivateLen characters. Errors never contain a value.
func LoadPrivate(path, repoRoot string) (privateList, error) {
	list := privateList{identities: map[string]bool{}}
	data, err := readOutside(path, repoRoot)
	if err != nil {
		return list, err
	}
	sc := bufio.NewScanner(bytes.NewReader(data))
	sc.Buffer(make([]byte, 0, 4096), 1<<20)
	for n := 1; sc.Scan(); n++ {
		value := strings.TrimSpace(sc.Text())
		if value == "" || strings.HasPrefix(value, "#") {
			continue
		}
		if rest, ok := strings.CutPrefix(value, identityPrefix); ok {
			value = strings.TrimSpace(rest)
			list.identities[strings.ToLower(value)] = true
		}
		if len(value) < minPrivateLen {
			return list, fmt.Errorf("private list line %d: shorter than %d characters", n, minPrivateLen)
		}
		list.values = append(list.values, privateValue{lower: bytes.ToLower([]byte(value)), line: n})
	}
	if err := sc.Err(); err != nil {
		return list, fmt.Errorf("reading the private list: %w", err)
	}
	return list, nil
}

func readOutside(path, repoRoot string) ([]byte, error) {
	resolved, err := filepath.EvalSymlinks(path)
	if err != nil {
		return nil, fmt.Errorf("private list: %w", err)
	}
	abs, err := filepath.Abs(resolved)
	if err != nil {
		return nil, fmt.Errorf("private list: %w", err)
	}
	root, err := filepath.EvalSymlinks(repoRoot)
	if err != nil {
		return nil, fmt.Errorf("repository root: %w", err)
	}
	root, err = filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("repository root: %w", err)
	}
	if rel, err := filepath.Rel(root, abs); err == nil && rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return nil, fmt.Errorf("private list: refused, it lies inside the repository")
	}
	info, err := os.Stat(abs)
	if err != nil {
		return nil, fmt.Errorf("private list: %w", err)
	}
	if !info.Mode().IsRegular() || info.Size() > maxFileSize {
		return nil, fmt.Errorf("private list: not a regular file of at most %d bytes", maxFileSize)
	}
	return os.ReadFile(filepath.Clean(abs))
}
