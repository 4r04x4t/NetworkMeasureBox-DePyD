package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ModFile is what depcheck needs of one go.mod.
type ModFile struct {
	File      string // the go.mod, relative to the repository
	Dir       string // its folder, absolute
	Module    string
	Toolchain string
	Require   []Require
	Replace   map[string]string // module path -> local folder (relative to Dir) or "path@version"
}

// Require is one requirement of a go.mod.
type Require struct {
	Path    string
	Version string
}

// ToolBuild is what `go version -m` says about an installed gate tool.
type ToolBuild struct {
	GoVersion string
	Package   string
	Module    string
	Version   string
}

// Env holds everything the check reads; main fills it from the go command, tests by hand.
type Env struct {
	Rows       []Row
	ModFiles   []ModFile
	MainModule string
	GoRoot     string
	GoVersion  string
	Roots      []string // folders licence files may come from (module cache, GOROOT, repository)
	ModuleDir  func(path, version string) (string, error)
	Tool       func(name, version string) (ToolBuild, error)
}

// Check returns every problem found; none means the list is complete and correct.
func Check(e Env) []string {
	var p []string
	toolchain := checkToolchain(e, &p)
	used := checkRequires(e, &p)
	for _, r := range e.Rows {
		if r.Use == useBuiltIn && !r.permitted {
			p = append(p, fmt.Sprintf("line %d: %s: licence %s is not allowed for parts built into nmbd; the owner decides (Q4)", r.Line, r.Name, r.Licence))
		}
		switch r.Kind {
		case kindModule:
			if !used[r.Module+"@"+r.Version] {
				p = append(p, fmt.Sprintf("line %d: %s %s is in no go.mod (stale row)", r.Line, r.Module, r.Version))
			}
		case kindToolchainPart:
			checkToolchainPart(e, r, toolchain, &p)
		case kindGateTool:
			checkGateTool(e, r, toolchain, &p)
		}
	}
	return append(p, duplicates(e.Rows)...)
}

func checkToolchain(e Env, p *[]string) string {
	var rows []Row
	for _, r := range e.Rows {
		if r.Kind == kindToolchain {
			rows = append(rows, r)
		}
	}
	if len(rows) != 1 {
		*p = append(*p, fmt.Sprintf("want exactly one %s row, found %d", kindToolchain, len(rows)))
		return ""
	}
	r := rows[0]
	for _, m := range e.ModFiles {
		if m.Toolchain != r.Version {
			*p = append(*p, fmt.Sprintf("%s: toolchain %q differs from the %s row %q", m.File, m.Toolchain, kindToolchain, r.Version))
		}
	}
	if e.GoVersion != r.Version {
		*p = append(*p, fmt.Sprintf("the go command runs %q, the %s row says %q", e.GoVersion, kindToolchain, r.Version))
	}
	licenceMatches(e, r, e.GoRoot, p)
	return r.Version
}

// checkRequires checks every requirement of every go.mod against the module rows and returns
// the set of "path@version" in use.
func checkRequires(e Env, p *[]string) map[string]bool {
	rows := map[string][]Row{}
	for _, r := range e.Rows {
		if r.Kind == kindModule {
			rows[r.Module] = append(rows[r.Module], r)
		}
	}
	used := map[string]bool{}
	for _, m := range e.ModFiles {
		for _, req := range m.Require {
			used[req.Path+"@"+req.Version] = true
			row, ok := findVersion(rows[req.Path], req.Version)
			switch {
			case !ok && len(rows[req.Path]) > 0:
				*p = append(*p, fmt.Sprintf("%s: %s %s: the list has another version (line %d)", m.File, req.Path, req.Version, rows[req.Path][0].Line))
			case !ok:
				*p = append(*p, fmt.Sprintf("%s: %s %s is not in docs/DEPENDENCIES.md", m.File, req.Path, req.Version))
			default:
				checkModuleRow(e, m, req, row, p)
			}
		}
	}
	return used
}

func findVersion(rows []Row, version string) (Row, bool) {
	for _, r := range rows {
		if r.Version == version {
			return r, true
		}
	}
	return Row{}, false
}

func checkModuleRow(e Env, m ModFile, req Require, r Row, p *[]string) {
	if m.Module == e.MainModule && r.Use == useBuildOnly {
		*p = append(*p, fmt.Sprintf("line %d: %s is required by the product module, so it cannot be %q", r.Line, r.Module, useBuildOnly))
	}
	dir, err := moduleDir(e, m, req)
	if err != nil {
		*p = append(*p, fmt.Sprintf("line %d: %s: %v", r.Line, r.Module, err))
		return
	}
	licenceMatches(e, r, dir, p)
}

func moduleDir(e Env, m ModFile, req Require) (string, error) {
	local, ok := m.Replace[req.Path]
	switch {
	case !ok:
		return e.ModuleDir(req.Path, req.Version)
	case filepath.IsAbs(local):
		return local, nil
	case strings.HasPrefix(local, "./") || strings.HasPrefix(local, "../"):
		return filepath.Join(m.Dir, local), nil
	}
	return "", fmt.Errorf("replaced by the module %s; depcheck does not handle module replacements yet", local)
}

// licenceNames are the names a module's main licence file may have, in this order.
func licenceNames() []string {
	return []string{"LICENSE", "LICENSE.txt", "LICENSE.md", "LICENCE", "COPYING"}
}

// readLicence returns the name and content of the main licence file in dir. The folder must lie
// under one of the roots, and os.Root keeps every lookup inside it.
func readLicence(dir string, roots []string) (string, []byte, error) {
	if !under(dir, roots) {
		return "", nil, errors.New("module folder outside the module cache, GOROOT and the repository")
	}
	root, err := os.OpenRoot(dir)
	if err != nil {
		return "", nil, err
	}
	defer func() { _ = root.Close() }()
	for _, name := range licenceNames() {
		if info, err := root.Stat(name); err != nil || !info.Mode().IsRegular() {
			continue
		}
		data, err := root.ReadFile(name)
		return name, data, err
	}
	return "", nil, fmt.Errorf("no licence file (%s) in the module", strings.Join(licenceNames(), ", "))
}

// licenceMatches checks that the licence file in dir hashes to the row.
func licenceMatches(e Env, r Row, dir string, p *[]string) {
	name, data, err := readLicence(dir, e.Roots)
	if err != nil {
		*p = append(*p, fmt.Sprintf("line %d: %s: %v", r.Line, r.Name, err))
		return
	}
	sum := sha256.Sum256(data)
	if got := hex.EncodeToString(sum[:]); got != r.Hash {
		*p = append(*p, fmt.Sprintf("line %d: %s: licence file %s has SHA-256 %s, the list says %s", r.Line, r.Name, name, got, r.Hash))
	}
}

func under(path string, roots []string) bool {
	for _, root := range roots {
		rel, err := filepath.Rel(root, path)
		if err == nil && rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) && !filepath.IsAbs(rel) {
			return true
		}
	}
	return false
}

func checkToolchainPart(e Env, r Row, toolchain string, p *[]string) {
	if r.Version != toolchain {
		*p = append(*p, fmt.Sprintf("line %d: %s: version %s differs from the toolchain %s", r.Line, r.Name, r.Version, toolchain))
	}
	if !insideToolchain(e.GoRoot, r.Module) {
		*p = append(*p, fmt.Sprintf("line %d: %s: %s is not part of the toolchain", r.Line, r.Name, r.Module))
	}
}

// insideToolchain reports whether pkg is a folder under GOROOT/src; os.Root keeps the lookup
// inside GOROOT whatever the row says.
func insideToolchain(goroot, pkg string) bool {
	root, err := os.OpenRoot(filepath.Join(goroot, "src"))
	if err != nil {
		return false
	}
	defer func() { _ = root.Close() }()
	info, err := root.Stat(filepath.FromSlash(pkg))
	return err == nil && info.IsDir()
}

func checkGateTool(e Env, r Row, toolchain string, p *[]string) {
	tb, err := e.Tool(r.Name, r.Version)
	if err != nil {
		*p = append(*p, fmt.Sprintf("line %d: %s: %v (run tools/prepare.sh)", r.Line, r.Name, err))
		return
	}
	if tb.GoVersion != toolchain || tb.Package != r.Module || tb.Version != r.Version {
		*p = append(*p, fmt.Sprintf("line %d: %s: installed %s %s built with %s, the list says %s %s with %s",
			r.Line, r.Name, tb.Package, tb.Version, tb.GoVersion, r.Module, r.Version, toolchain))
		return
	}
	dir, err := e.ModuleDir(tb.Module, tb.Version)
	if err != nil {
		*p = append(*p, fmt.Sprintf("line %d: %s: %v", r.Line, r.Name, err))
		return
	}
	licenceMatches(e, r, dir, p)
}

func duplicates(rows []Row) []string {
	seen := map[string]int{}
	var p []string
	for _, r := range rows {
		key := r.Kind + " " + r.Module + " " + r.Version
		if first, dup := seen[key]; dup {
			p = append(p, fmt.Sprintf("line %d: duplicate of line %d", r.Line, first))
			continue
		}
		seen[key] = r.Line
	}
	return p
}
