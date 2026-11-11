#!/usr/bin/env bash
# Proves that the gate can fail (brief F-00 4.7): on throw-away copies of the repository (never on
# this work tree), each planted fault must make its step fail, and the clean copy must pass.
# Usage: tools/gate-selftest.sh [CASE...]
# Cases: clean gofmt nesting gosec depcheck email testfail testwrites root summary matrix pipe
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
nmb_init

WORK=$(mktemp -d)
trap 'rm -rf -- "$WORK"' EXIT

# copy_repo NAME: a copy of every tracked and untracked (not ignored) file, as a new git
# repository with one commit; prints its folder.
copy_repo() {
	local dest="$WORK/$1" list="$WORK/$1.files"
	nmb_files . >"$list"
	mkdir -p -- "$dest"
	xargs -0 -a "$list" cp --parents -P -t "$dest" --
	GIT_CONFIG_GLOBAL=/dev/null GIT_CONFIG_NOSYSTEM=1 git -C "$dest" init -q -b main
	git -C "$dest" add -A
	GIT_CONFIG_GLOBAL=/dev/null GIT_CONFIG_NOSYSTEM=1 git -C "$dest" -c user.name=selftest \
		-c user.email=selftest@example.com commit -q -m "gate self-test copy"
	printf '%s\n' "$dest"
}

# expect WANT MARKER DIR COMMAND...: runs COMMAND in the copy DIR; WANT is 0, 2 or "fail"
# (any non-zero code); the output must contain MARKER.
expect() {
	local want=$1 marker=$2 dir=$3 rc=0 ok=1
	shift 3
	local out
	out="$dir.$(date +%s%N).out"
	(cd -- "$dir" && env -u NMB_GATE_OUT -u NMB_IN_OFFLINE "$@") >"$out" 2>&1 || rc=$?
	case "$want" in
	fail) [ "$rc" -ne 0 ] || ok=0 ;;
	*) [ "$rc" -eq "$want" ] || ok=0 ;;
	esac
	grep -qF -- "$marker" "$out" || ok=0
	printf 'case %s: rc %s (want %s), marker "%s" %s\n' "$*" "$rc" "$want" "$marker" \
		"$(if grep -qF -- "$marker" "$out"; then echo found; else echo missing; fi)"
	if [ "$ok" -ne 1 ]; then
		tail -n 30 -- "$out"
		return 1
	fi
}

plant() {
	mkdir -p -- "$(dirname -- "$1/$2")"
	cat >"$1/$2"
}

case_clean() {
	local d
	d=$(copy_repo clean)
	plant "$d" internal/selftestprobe/probe.go <<'GO'
// Package selftestprobe exists only in the gate self-test copy.
package selftestprobe
GO
	plant "$d" internal/selftestprobe/probe_test.go <<'GO'
package selftestprobe

import (
	"net"
	"os"
	"testing"
)

func TestHermetic(t *testing.T) {
	for _, dir := range []string{os.Getenv("HOME"), os.TempDir()} {
		if entries, err := os.ReadDir(dir); err != nil || len(entries) != 0 {
			t.Errorf("%s: %d entries, %v", dir, len(entries), err)
		}
	}
	ifs, err := net.Interfaces()
	if err != nil || len(ifs) != 1 || ifs[0].Name != "lo" {
		t.Errorf("interfaces %v, %v: want loopback only", ifs, err)
	}
	if os.Getuid() == 0 {
		t.Errorf("tests run as root")
	}
}
GO
	expect 0 "test:B" "$d" tools/test.sh B
	expect 0 "analyse:scriptrules" "$d" tools/analyse.sh gofmt cleancode scriptrules gosec
	expect 0 "scan:tree" "$d" tools/scan.sh tree
	expect 0 "deps:depcheck" "$d" tools/deps.sh
	expect 0 "build:sizes" "$d" tools/build-matrix.sh linux-arm64
}

case_gofmt() {
	local d
	d=$(copy_repo gofmt)
	printf 'package buildinfo\nfunc  planted( ) {}\n' >"$d/internal/buildinfo/planted.go"
	expect fail "not gofmt-formatted" "$d" tools/analyse.sh gofmt
}

case_nesting() {
	local d
	d=$(copy_repo nesting)
	plant "$d" internal/buildinfo/planted.go <<'GO'
package buildinfo

func planted(x int) int {
	if x > 0 {
		if x > 1 {
			if x > 2 {
				if x > 3 {
					return 4
				}
			}
		}
	}
	return 0
}
GO
	expect fail "nesting planted" "$d" tools/analyse.sh cleancode
}

case_gosec() {
	local d
	d=$(copy_repo gosec)
	plant "$d" internal/buildinfo/planted.go <<'GO'
package buildinfo

import (
	"os"
	"os/exec"
)

// Planted runs a shell with an argument and reads a file named by the caller.
func Planted(arg, path string) ([]byte, error) {
	if err := exec.Command("sh", "-c", arg).Run(); err != nil {
		return nil, err
	}
	return os.ReadFile(path)
}
GO
	expect fail "G204" "$d" tools/analyse.sh gosec
}

case_depcheck() {
	local d
	d=$(copy_repo depcheck)
	plant "$d" third_party/plantedmod/go.mod <<<'module example.com/plantedmod'
	plant "$d" third_party/plantedmod/LICENSE <<<'Permission is granted.'
	plant "$d" third_party/plantedmod/planted.go <<<'package plantedmod'
	printf '\nrequire example.com/plantedmod v0.0.0\n\nreplace example.com/plantedmod => ./third_party/plantedmod\n' >>"$d/go.mod"
	expect fail "example.com/plantedmod v0.0.0 is not in docs/DEPENDENCIES.md" "$d" tools/deps.sh depcheck
}

case_email() {
	local d
	d=$(copy_repo email)
	printf 'contact: %s%s%s\n' 'jane.roe' '@' 'mailhost.net' >"$d/docs/planted.txt"
	expect fail "docs/planted.txt:1:10: email:" "$d" tools/scan.sh tree
}

case_testfail() {
	local d
	d=$(copy_repo testfail)
	printf 'package buildinfo\n\nimport "testing"\n\nfunc TestPlanted(t *testing.T) { t.Fatal("planted") }\n' \
		>"$d/internal/buildinfo/planted_test.go"
	expect fail "FAIL" "$d" tools/test.sh B
}

case_testwrites() {
	local d
	d=$(copy_repo testwrites)
	plant "$d" internal/buildinfo/planted_test.go <<'GO'
package buildinfo

import (
	"os"
	"testing"
)

func TestPlantedWrite(t *testing.T) {
	if err := os.WriteFile("planted-by-test.txt", []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
}
GO
	expect fail "the tests changed the repository" "$d" tools/test.sh B
}

case_root() {
	local d
	d=$(copy_repo root)
	unshare -r true || {
		printf 'cannot make a user namespace to test the refusal\n'
		return 1
	}
	expect 2 "refuses to run as root" "$d" unshare -r tools/test.sh B
}

case_summary() {
	local d
	d=$(copy_repo summary)
	# shellcheck disable=SC2016 # expanded by the inner bash
	expect 1 "1 of 2 steps failed" "$d" bash -c 'set -euo pipefail; . tools/lib.sh; nmb_init
		nmb_run_step planted-fail here false; nmb_run_step planted-pass here true; nmb_finish'
}

case_matrix() {
	local d
	d=$(copy_repo matrix)
	printf '//go:build amd64\n\npackage buildinfo\n\nfunc onlyAMD64() int { return 1 }\n' >"$d/internal/buildinfo/planted_amd64.go"
	printf 'package buildinfo\n\nvar _ = onlyAMD64()\n' >"$d/internal/buildinfo/planted.go"
	expect fail "undefined: onlyAMD64" "$d" tools/build-matrix.sh linux-arm64
}

case_pipe() {
	local d
	d=$(copy_repo pipe)
	printf 'go version %s tail -n 1\n' '|' >>"$d/tools/build-matrix.sh"
	expect fail "a pipe to tail or tee" "$d" tools/analyse.sh scriptrules
}

all=(clean gofmt nesting gosec depcheck email testfail testwrites root summary matrix pipe)
cases=("$@")
[ "${#cases[@]}" -gt 0 ] || cases=("${all[@]}")
for c in "${cases[@]}"; do
	case " ${all[*]} " in
	*" $c "*) nmb_run_step "selftest:$c" "PC, copy of the repository" "case_$c" ;;
	*) nmb_die "unknown case $c" ;;
	esac
done
nmb_finish
