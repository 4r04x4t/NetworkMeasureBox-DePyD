package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
)

type goEnvJSON struct {
	GOROOT     string
	GOVERSION  string
	GOMODCACHE string
}

func loadGoEnv(goCmd goRunner, dir string) (goEnvJSON, error) {
	out, err := goCmd(dir, "env", "-json", "GOROOT", "GOVERSION", "GOMODCACHE")
	if err != nil {
		return goEnvJSON{}, err
	}
	var e goEnvJSON
	if err := json.Unmarshal(out, &e); err != nil {
		return e, fmt.Errorf("go env: %w", err)
	}
	if e.GOROOT == "" || e.GOVERSION == "" || e.GOMODCACHE == "" {
		return e, fmt.Errorf("go env: GOROOT, GOVERSION or GOMODCACHE empty")
	}
	return e, nil
}

// modEditJSON is the part of `go mod edit -json` depcheck reads.
type modEditJSON struct {
	Module    struct{ Path string }
	Toolchain string
	Require   []Require
	Replace   []struct {
		Old struct{ Path string }
		New struct{ Path, Version string }
	}
}

// loadModFile reads one go.mod through `go mod edit -json` (no network, no build list).
func loadModFile(goCmd goRunner, repo, gomod string) (ModFile, error) {
	if filepath.Base(gomod) != "go.mod" || filepath.IsAbs(gomod) || strings.HasPrefix(filepath.Clean(gomod), "..") {
		return ModFile{}, fmt.Errorf("%s: want a go.mod path inside the repository", gomod)
	}
	dir := filepath.Join(repo, filepath.Dir(gomod))
	out, err := goCmd(dir, "mod", "edit", "-json")
	if err != nil {
		return ModFile{}, err
	}
	return parseModEdit(gomod, dir, out)
}

func parseModEdit(gomod, dir string, out []byte) (ModFile, error) {
	var j modEditJSON
	if err := json.Unmarshal(out, &j); err != nil {
		return ModFile{}, fmt.Errorf("%s: go mod edit -json: %w", gomod, err)
	}
	m := ModFile{File: gomod, Dir: dir, Module: j.Module.Path, Toolchain: j.Toolchain, Require: j.Require,
		Replace: map[string]string{}}
	for _, r := range j.Replace {
		target := r.New.Path
		if r.New.Version != "" {
			target += "@" + r.New.Version
		}
		m.Replace[r.Old.Path] = target
	}
	return m, nil
}

type downloadJSON struct {
	Dir   string
	Error string
}

// downloadDir returns the module cache folder of path@version (from the cache when offline).
func downloadDir(goCmd goRunner, repo, path, version string) (string, error) {
	out, err := goCmd(repo, "mod", "download", "-json", path+"@"+version)
	if err != nil {
		return "", err
	}
	var d downloadJSON
	if err := json.Unmarshal(out, &d); err != nil {
		return "", fmt.Errorf("go mod download: %w", err)
	}
	if d.Error != "" || d.Dir == "" {
		return "", fmt.Errorf("go mod download %s@%s: %s", path, version, d.Error)
	}
	return d.Dir, nil
}

// toolBuild reads `go version -m` of an installed tool.
func toolBuild(goCmd goRunner, repo, binary string) (ToolBuild, error) {
	out, err := goCmd(repo, "version", "-m", binary)
	if err != nil {
		return ToolBuild{}, err
	}
	return parseVersionM(string(out))
}

func parseVersionM(out string) (ToolBuild, error) {
	var tb ToolBuild
	lines := strings.Split(out, "\n")
	if _, v, ok := strings.Cut(lines[0], ": "); ok {
		tb.GoVersion = strings.TrimSpace(v)
	}
	for _, l := range lines[1:] {
		f := strings.Split(strings.TrimSpace(l), "\t")
		switch {
		case len(f) >= 2 && f[0] == "path":
			tb.Package = f[1]
		case len(f) >= 3 && f[0] == "mod":
			tb.Module, tb.Version = f[1], f[2]
		}
	}
	if tb.GoVersion == "" || tb.Package == "" || tb.Module == "" {
		return tb, fmt.Errorf("go version -m: no build information")
	}
	return tb, nil
}
