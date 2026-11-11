package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const licenceText = "Permission is granted.\n"

func licenceHash() string {
	sum := sha256.Sum256([]byte(licenceText))
	return hex.EncodeToString(sum[:])
}

func mkdirWithLicence(t *testing.T, dir, name string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, name), []byte(licenceText), 0o600); err != nil {
		t.Fatal(err)
	}
}

// fixture is a clean environment: one toolchain, one product module requiring one module, one
// gate tool and the race runtime, all with licence files in temp folders.
type fixture struct {
	env   Env
	cache string
}

func newFixture(t *testing.T) *fixture {
	t.Helper()
	root := t.TempDir()
	goroot, cache, repo := filepath.Join(root, "goroot"), filepath.Join(root, "cache"), filepath.Join(root, "repo")
	mkdirWithLicence(t, goroot, "LICENSE")
	mkdirWithLicence(t, filepath.Join(goroot, "src", "runtime", "race"), "README")
	mkdirWithLicence(t, filepath.Join(cache, "example.com/lib@v1.2.3"), "LICENSE.txt")
	mkdirWithLicence(t, filepath.Join(cache, "example.com/tool@v0.1.0"), "LICENSE")
	h := licenceHash()
	rows := []Row{
		{Line: 3, Name: "Go", Kind: kindToolchain, Module: "go", Version: "go1.27.4", Licence: "BSD-3-Clause", Use: useBuiltIn, Hash: h, permitted: true},
		{Line: 4, Name: "race", Kind: kindToolchainPart, Module: "runtime/race", Version: "go1.27.4", Licence: "Apache-2.0 WITH LLVM-exception", Use: useTestOnly, Hash: noHash},
		{Line: 5, Name: "lib", Kind: kindModule, Module: "example.com/lib", Version: "v1.2.3", Licence: "MIT", Use: useBuiltIn, Hash: h, permitted: true},
		{Line: 6, Name: "tool", Kind: kindGateTool, Module: "example.com/tool/cmd/tool", Version: "v0.1.0", Licence: "MIT", Use: useBuildOnly, Hash: h, permitted: true},
		{Line: 7, Name: "git", Kind: kindSystem, Module: "git", Version: "2.43.0", Licence: "GPL-2.0-only", Use: useBuildOnly, Hash: noHash},
	}
	env := Env{
		Rows: rows, MainModule: "nmb", GoRoot: goroot, GoVersion: "go1.27.4", Roots: []string{cache, goroot, repo},
		ModFiles: []ModFile{{File: "go.mod", Dir: repo, Module: "nmb", Toolchain: "go1.27.4",
			Require: []Require{{"example.com/lib", "v1.2.3"}}, Replace: map[string]string{}}},
		ModuleDir: func(path, version string) (string, error) {
			dir := filepath.Join(cache, path+"@"+version)
			if _, err := os.Stat(dir); err != nil {
				return "", errors.New("not in the module cache")
			}
			return dir, nil
		},
		Tool: func(name, version string) (ToolBuild, error) {
			return ToolBuild{GoVersion: "go1.27.4", Package: "example.com/tool/cmd/tool", Module: "example.com/tool", Version: "v0.1.0"}, nil
		},
	}
	return &fixture{env: env, cache: cache}
}

func expectProblem(t *testing.T, env Env, want string) {
	t.Helper()
	problems := Check(env)
	if len(problems) != 1 || !strings.Contains(problems[0], want) {
		t.Errorf("problems %q, want one containing %q", problems, want)
	}
}

func TestCleanPasses(t *testing.T) {
	if p := Check(newFixture(t).env); len(p) != 0 {
		t.Errorf("problems: %q", p)
	}
}

func TestUnlistedModuleFails(t *testing.T) {
	f := newFixture(t)
	f.env.ModFiles[0].Require = append(f.env.ModFiles[0].Require, Require{"example.com/other", "v0.0.1"})
	expectProblem(t, f.env, "example.com/other v0.0.1 is not in docs/DEPENDENCIES.md")
}

func TestUnlistedLocalReplacementFails(t *testing.T) {
	f := newFixture(t)
	m := &f.env.ModFiles[0]
	m.Require = append(m.Require, Require{"example.com/local", "v0.0.0"})
	m.Replace["example.com/local"] = "./third_party/local"
	expectProblem(t, f.env, "example.com/local v0.0.0 is not in docs/DEPENDENCIES.md")
}

func TestWrongVersionFails(t *testing.T) {
	f := newFixture(t)
	f.env.ModFiles[0].Require[0].Version = "v1.2.4"
	problems := Check(f.env)
	if len(problems) != 2 || !strings.Contains(problems[0], "the list has another version") || !strings.Contains(problems[1], "stale row") {
		t.Errorf("problems %q", problems)
	}
}

func TestStaleRowFails(t *testing.T) {
	f := newFixture(t)
	f.env.ModFiles[0].Require = nil
	expectProblem(t, f.env, "example.com/lib v1.2.3 is in no go.mod (stale row)")
}

func TestCopyleftBuiltInFails(t *testing.T) {
	f := newFixture(t)
	f.env.Rows[2].Licence, f.env.Rows[2].permitted = "GPL-3.0-only", false
	expectProblem(t, f.env, "licence GPL-3.0-only is not allowed for parts built into nmbd")
}

func TestProductModuleCannotBeBuildOnly(t *testing.T) {
	f := newFixture(t)
	f.env.Rows[2].Use = useBuildOnly
	expectProblem(t, f.env, "required by the product module")
}

func TestLicenceHashMismatchFails(t *testing.T) {
	f := newFixture(t)
	path := filepath.Join(f.cache, "example.com/lib@v1.2.3", "LICENSE.txt")
	if err := os.WriteFile(path, []byte(licenceText+"changed\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	expectProblem(t, f.env, "licence file LICENSE.txt has SHA-256")
}

func TestMissingLicenceFileFails(t *testing.T) {
	f := newFixture(t)
	if err := os.Remove(filepath.Join(f.cache, "example.com/tool@v0.1.0", "LICENSE")); err != nil {
		t.Fatal(err)
	}
	expectProblem(t, f.env, "no licence file")
}

func TestLicenceOutsideRootsFails(t *testing.T) {
	f := newFixture(t)
	f.env.Roots = f.env.Roots[1:]
	problems := Check(f.env)
	if len(problems) != 2 || !strings.Contains(problems[0], "outside the module cache") || !strings.Contains(problems[1], "outside the module cache") {
		t.Errorf("problems %q", problems)
	}
}

func TestToolchainRowMismatchFails(t *testing.T) {
	f := newFixture(t)
	f.env.ModFiles[0].Toolchain = "go1.27.3"
	expectProblem(t, f.env, `go.mod: toolchain "go1.27.3" differs`)
	f = newFixture(t)
	f.env.GoVersion = "go1.22.2"
	expectProblem(t, f.env, `the go command runs "go1.22.2"`)
	f = newFixture(t)
	f.env.Rows[1].Version = "go1.27.3"
	expectProblem(t, f.env, "differs from the toolchain")
	f = newFixture(t)
	f.env.Rows = append(f.env.Rows, f.env.Rows[0])
	if p := Check(f.env); len(p) == 0 || !strings.Contains(p[0], "exactly one toolchain row") {
		t.Errorf("two toolchain rows: %q", p)
	}
}

func TestGateToolVersionMismatchFails(t *testing.T) {
	f := newFixture(t)
	f.env.Tool = func(string, string) (ToolBuild, error) {
		return ToolBuild{GoVersion: "go1.22.2", Package: "example.com/tool/cmd/tool", Module: "example.com/tool", Version: "v0.1.0"}, nil
	}
	expectProblem(t, f.env, "built with go1.22.2")
	f.env.Tool = func(string, string) (ToolBuild, error) { return ToolBuild{}, errors.New("not installed") }
	expectProblem(t, f.env, "run tools/prepare.sh")
}

func TestDuplicateRowFails(t *testing.T) {
	f := newFixture(t)
	f.env.Rows = append(f.env.Rows, f.env.Rows[4])
	expectProblem(t, f.env, "duplicate of line 7")
}
