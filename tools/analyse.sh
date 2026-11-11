#!/usr/bin/env bash
# Static analysis of the work plan's gate (2.4): gofmt, go vet, staticcheck, gosec (through the
# tools/gosecvet driver), govulncheck (source and the built binaries), gocyclo, the clean-code
# limits, shellcheck and the rules of the gate scripts themselves. Warnings fail.
# Usage: tools/analyse.sh [STEP...]
# Steps: gofmt vet staticcheck gosec govulncheck gocyclo cleancode shellcheck scriptrules nosec gosecvet-module
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
nmb_init

readonly CYCLO_LIMIT=15

step_gofmt() {
	local gofmt out
	local -a files
	gofmt="$(go env GOROOT)/bin/gofmt"
	nmb_file_array files '*.go'
	out=$("$gofmt" -l -- "${files[@]}")
	if [ -n "$out" ]; then
		printf 'not gofmt-formatted:\n%s\n' "$out"
		return 1
	fi
	printf 'gofmt: %s files formatted\n' "${#files[@]}"
}

step_vet() {
	local target
	for target in amd64 arm64 arm mips; do
		printf 'go vet GOARCH=%s\n' "$target"
		GOARCH=$target GOARM=7 GOMIPS=softfloat go vet ./...
	done
}

step_staticcheck() {
	local sc
	sc=$(nmb_tool staticcheck)
	"$sc" -version
	printf 'staticcheck GOARCH=amd64\n'
	GOARCH=amd64 "$sc" -checks all ./...
	printf 'staticcheck GOARCH=arm GOARM=7 (32-bit)\n'
	GOARCH=arm GOARM=7 "$sc" -checks all ./...
}

step_gosec() {
	"$(nmb_tool gosecvet)" -test ./...
}

step_govulncheck() {
	local gv t
	gv=$(nmb_tool govulncheck)
	"$gv" -version
	"$gv" -test ./...
	for t in dist/*/nmbd; do
		[ -f "$t" ] || {
			printf 'no binaries in dist/: run tools/build-matrix.sh first\n'
			return 1
		}
		printf 'govulncheck -mode binary %s\n' "$t"
		"$gv" -mode binary "$t"
	done
}

step_gocyclo() {
	local -a files
	nmb_file_array files '*.go'
	"$(nmb_tool gocyclo)" -over "$CYCLO_LIMIT" "${files[@]}"
	printf 'gocyclo: no function above %s in %s files\n' "$CYCLO_LIMIT" "${#files[@]}"
}

step_cleancode() {
	local -a files
	nmb_file_array files '*.go' '*.sh'
	nmb_build_helper cleancode
	"$NMB_BIN/cleancode" "${files[@]}"
	printf 'cleancode: %s files within the limits\n' "${#files[@]}"
}

step_shellcheck() {
	local -a files
	nmb_file_array files '*.sh'
	shellcheck --version | sed -n 2p
	shellcheck -x --severity=style -- "${files[@]}"
}

# step_scriptrules: every gate script uses strict mode; no status is taken behind a pipe to tail,
# tee or grep; no eval.
step_scriptrules() {
	local f bad=0
	local -a files
	nmb_file_array files 'tools/*.sh'
	for f in "${files[@]}"; do
		if [ "$f" != tools/lib.sh ] && ! grep -qx 'set -euo pipefail' -- "$f"; then
			printf '%s: no "set -euo pipefail" line\n' "$f"
			bad=1
		fi
		if grep -nE '\|[[:space:]]*(tail|tee)([[:space:]]|$)|^[[:space:]]*eval[[:space:]]|[;&|][[:space:]]*eval[[:space:]]' -- "$f"; then
			printf '%s: a pipe to tail or tee, or the eval builtin (see the rules in tools/lib.sh)\n' "$f"
			bad=1
		fi
	done
	[ "$bad" -eq 0 ]
	printf 'script rules: %s scripts\n' "${#files[@]}"
}

# step_nosec lists every #nosec annotation so the verifier sees it; each needs a reason after "--".
step_nosec() {
	local f line bad=0
	local -a files
	nmb_file_array files '*.go'
	for f in "${files[@]}"; do
		while IFS= read -r line; do
			printf 'nosec: %s:%s\n' "$f" "$line"
			case "$line" in
			*"#nosec G"[0-9]*" -- "?*) ;;
			*) bad=1 ;;
			esac
		done < <(grep -n '#nosec' -- "$f" || true)
	done
	[ "$bad" -eq 0 ] || {
		printf 'a #nosec without a rule id and a reason after --\n'
		return 1
	}
}

step_gosecvet-module() {
	local gofmt out
	gofmt="$(go env GOROOT)/bin/gofmt"
	out=$("$gofmt" -l tools/gosecvet)
	[ -z "$out" ] || {
		printf 'not gofmt-formatted: %s\n' "$out"
		return 1
	}
	(cd tools/gosecvet && go vet ./... && "$(nmb_tool staticcheck)" -checks all ./... && "$(nmb_tool gosecvet)" ./...)
}

all=(gofmt vet staticcheck gosec govulncheck gocyclo cleancode shellcheck scriptrules nosec gosecvet-module)
nmb_load_tools
steps=("$@")
[ "${#steps[@]}" -gt 0 ] || steps=("${all[@]}")
for s in "${steps[@]}"; do
	case " ${all[*]} " in
	*" $s "*) ;;
	*) nmb_die "unknown step $s" ;;
	esac
	where="PC amd64"
	[ "$s" != govulncheck ] || where="PC amd64, network (vuln.go.dev)"
	nmb_run_step "analyse:$s" "$where" "step_$s"
done
nmb_finish
