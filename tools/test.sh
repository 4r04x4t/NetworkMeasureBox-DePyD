#!/usr/bin/env bash
# The test runs of the gate (work plan 2.4, brief F-00 4.3.1), hermetic: no network (a new network
# namespace with loopback only, as the real user without capabilities), fresh empty HOME and
# TMPDIR for every run, no test cache, and the repository must be unchanged afterwards.
# Usage: tools/test.sh [RUN...]
#   A  amd64 with the race detector (needs cgo: test binaries only), TZ=UTC
#   B  amd64 with CGO_ENABLED=0, shuffled, TZ=Pacific/Chatham
#   C  arm64 under qemu-aarch64-static, TZ=UTC        D  the same shuffled, TZ=Pacific/Chatham
#   E  arm v7 under qemu-arm-static, TZ=UTC           F  the same shuffled, TZ=Pacific/Chatham
#   G  mips and mipsle (softfloat): test binaries compiled, not run
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
nmb_init
nmb_go_offline "$NMB_ROOT/tools/test.sh" "$@"

readonly OTHER_TZ=Pacific/Chatham

# repo_state: the status of the tracked files, and name, size and modification time of every
# untracked file, ignored ones included (build/, dist/), except this gate run's own folder.
repo_state() {
	local list="$NMB_OUT/.untracked.$$" rel=${NMB_OUT#"$NMB_ROOT"/} f
	local -a all keep=()
	git status --porcelain=v1 -uall || return 1
	git ls-files -z --others >"$list" || return 1
	mapfile -d '' all <"$list"
	rm -f -- "$list"
	for f in "${all[@]}"; do
		case "$f" in
		"$rel"/*) ;;
		*) keep+=("$f") ;;
		esac
	done
	[ "${#keep[@]}" -eq 0 ] || stat --printf '%n\t%s\t%.9Y\n' -- "${keep[@]}"
}

# unchanged COMMAND...: runs the command and fails when the repository changed meanwhile, ignored
# paths included. The command runs in a subshell with errexit on, never as a condition (bash
# would ignore errexit in every function it calls).
unchanged() {
	local before after rc
	before=$(repo_state) || return 1
	set +e
	(
		set -euo pipefail
		"$@"
	)
	rc=$?
	set -e
	after=$(repo_state) || return 1
	if [ "$before" != "$after" ]; then
		printf 'the tests changed the repository:\n%s\n' "$(diff <(printf '%s\n' "$before") <(printf '%s\n' "$after") || true)"
		return 1
	fi
	return "$rc"
}

run_A() { CGO_ENABLED=1 TZ=UTC nmb_hermetic go test -race -count=1 -p "$NMB_GATE_JOBS" ./...; }
run_B() { CGO_ENABLED=0 TZ=$OTHER_TZ nmb_hermetic go test -count=1 -shuffle=on -p "$NMB_GATE_JOBS" ./...; }

# test_packages ARCH: "import-path<TAB>folder" of every package with tests for that architecture.
test_packages() {
	GOARCH=$1 GOARM=7 GOMIPS=softfloat go list -f '{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{"\t"}}{{.Dir}}{{end}}' ./...
}

# cross_run ARCH QEMU TZ [FLAG...]: compiles each package's test binary and runs it under QEMU
# user mode from the package folder (QEMU empty: compile only).
cross_run() {
	local arch=$1 qemu=$2 tz=$3 pkg dir bin list n=0
	shift 3
	if [ -n "$qemu" ]; then
		local name=$qemu
		qemu=$(command -v "$name") || {
			printf '%s is not installed\n' "$name"
			return 1
		}
	fi
	bin=$(mktemp -d "$NMB_OUT/testbin.XXXXXX") || return 1
	list=$(test_packages "$arch") || return 1
	while IFS=$'\t' read -r pkg dir; do
		[ -n "$pkg" ] || continue
		GOARCH=$arch GOARM=7 GOMIPS=softfloat go test -c -o "$bin/${pkg//\//-}.test" "$pkg" || return 1
		n=$((n + 1))
		[ -n "$qemu" ] || continue
		printf '%s %s\n' "$(basename -- "$qemu")" "$pkg"
		(cd -- "$dir" && TZ=$tz nmb_hermetic "$qemu" "$bin/${pkg//\//-}.test" -test.count=1 "$@") || return 1
	done <<<"$list"
	rm -rf -- "$bin"
	printf '%s: %s test binaries\n' "$arch" "$n"
	[ "$n" -gt 0 ]
}

run_C() { cross_run arm64 qemu-aarch64-static UTC; }
run_D() { cross_run arm64 qemu-aarch64-static "$OTHER_TZ" -test.shuffle=on; }
run_E() { cross_run arm qemu-arm-static UTC; }
run_F() { cross_run arm qemu-arm-static "$OTHER_TZ" -test.shuffle=on; }
run_G() {
	cross_run mips "" UTC || return 1
	cross_run mipsle "" UTC
}

where_of() {
	case "$1" in
	A | B) printf 'PC amd64, %s' "$NMB_WHERE_NET" ;;
	C | D) printf 'QEMU user mode arm64 on the PC, %s' "$NMB_WHERE_NET" ;;
	E | F) printf 'QEMU user mode arm v7 on the PC, %s' "$NMB_WHERE_NET" ;;
	G) printf 'PC, compile only' ;;
	esac
}

nmb_hermetic_init || nmb_die "go env failed"
runs=("$@")
[ "${#runs[@]}" -gt 0 ] || runs=(A B C D E F G)
for r in "${runs[@]}"; do
	case "$r" in
	[A-G]) nmb_run_step "test:$r" "$(where_of "$r")" unchanged "run_$r" ;;
	*) nmb_die "unknown run $r" ;;
	esac
done
nmb_finish
