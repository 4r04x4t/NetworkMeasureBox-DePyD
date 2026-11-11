#!/usr/bin/env bash
# Builds every package and the nmbd binary for the five targets of the work plan (2.4) with
# CGO_ENABLED=0, checks the recorded build settings of each binary and records its size.
# Usage: tools/build-matrix.sh [TARGET...]   targets: linux-amd64 linux-arm64 linux-armv7 linux-mips linux-mipsle
# Output: dist/<target>/nmbd (git-ignored) and sizes.tsv in the result folder.
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
nmb_init

# target_env TARGET: prints the GOARCH and the variant setting of a target.
target_env() {
	case "$1" in
	linux-amd64) printf 'amd64\t\n' ;;
	linux-arm64) printf 'arm64\t\n' ;;
	linux-armv7) printf 'arm\tGOARM=7\n' ;;
	linux-mips) printf 'mips\tGOMIPS=softfloat\n' ;;
	linux-mipsle) printf 'mipsle\tGOMIPS=softfloat\n' ;;
	*) return 1 ;;
	esac
}

step_build() {
	local target=$1 arch variant out info size
	IFS=$'\t' read -r arch variant <<<"$(target_env "$target")"
	local -a env=(CGO_ENABLED=0 GOOS=linux "GOARCH=$arch")
	if [ -n "$variant" ]; then
		env+=("$variant")
	fi
	out="dist/$target/nmbd"
	mkdir -p -- "dist/$target"
	env "${env[@]}" go build -p "$NMB_GATE_JOBS" ./...
	env "${env[@]}" go build -p "$NMB_GATE_JOBS" -trimpath -ldflags='-s -w' -o "$out" ./cmd/nmbd
	info=$(go version -m "$out")
	printf '%s\n' "$info"
	local want
	for want in CGO_ENABLED=0 GOOS=linux "GOARCH=$arch" ${variant:+"$variant"} -trimpath=true; do
		grep -qE "^[[:space:]]build[[:space:]]$want\$" <<<"$info" || {
			printf 'binary %s lacks the build setting %s\n' "$out" "$want"
			return 1
		}
	done
	size=$(stat -c %s -- "$out")
	[ "$size" -gt 0 ]
	printf '%s\t%s\n' "$target" "$size" >>"$NMB_OUT/sizes.tsv"
	printf 'size %s: %s bytes\n' "$target" "$size"
}

step_sizes() {
	local n
	n=$(wc -l <"$NMB_OUT/sizes.tsv")
	cat -- "$NMB_OUT/sizes.tsv"
	[ "$n" -eq "$1" ] || {
		printf 'want %s sizes, found %s\n' "$1" "$n"
		return 1
	}
}

targets=("$@")
[ "${#targets[@]}" -gt 0 ] || targets=(linux-amd64 linux-arm64 linux-armv7 linux-mips linux-mipsle)
for t in "${targets[@]}"; do
	target_env "$t" >/dev/null || nmb_die "unknown target $t"
done
: >"$NMB_OUT/sizes.tsv"
for t in "${targets[@]}"; do
	nmb_run_step "build:$t" "PC (cross build, CGO_ENABLED=0)" step_build "$t"
done
nmb_run_step "build:sizes" "PC" step_sizes "${#targets[@]}"
nmb_finish
