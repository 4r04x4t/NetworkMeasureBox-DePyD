#!/usr/bin/env bash
# Fetches what the gate needs, through GOPROXY with checksum verification only: the toolchain of
# go.mod, the modules of every go.mod, the gate tools at the versions of docs/DEPENDENCIES.md
# (built with that toolchain into NMB_TOOLS_DIR, outside every repository) and the gosec driver.
# Usage: tools/prepare.sh [STEP...]   steps: toolchain modules tools gosecvet
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
nmb_init

step_toolchain() {
	go version
	go env GOROOT GOVERSION
}

step_modules() {
	local gomod
	while IFS= read -r gomod; do
		(cd -- "$(dirname -- "$gomod")" && go mod download && go mod verify)
	done < <(nmb_gomods)
}

# installed_ok BIN PKG VER: the binary was built by the fixed toolchain from PKG at VER.
installed_ok() {
	local info
	[ -x "$1" ] || return 1
	info=$(go version -m "$1") || return 1
	[ "$(sed -n '1s/^[^:]*: //p' <<<"$info")" = "$NMB_TOOLCHAIN" ] || return 1
	grep -qxF "$(printf '\tpath\t%s' "$2")" <<<"$info" || return 1
	grep -qE "^[[:space:]]mod[[:space:]][^[:space:]]+[[:space:]]$3([[:space:]]|$)" <<<"$info"
}

step_tools() {
	local name bin dir
	for name in "${!NMB_TOOL_PKG[@]}"; do
		bin=${NMB_TOOL_BIN[$name]}
		dir=$(dirname -- "$bin")
		if installed_ok "$bin" "${NMB_TOOL_PKG[$name]}" "${NMB_TOOL_VER[$name]}"; then
			printf '%s: installed, %s %s built with %s\n' "$name" "${NMB_TOOL_PKG[$name]}" "${NMB_TOOL_VER[$name]}" "$NMB_TOOLCHAIN"
			continue
		fi
		mkdir -p -- "$dir"
		(cd -- "$dir" && GOBIN="$dir" go install -trimpath "${NMB_TOOL_PKG[$name]}@${NMB_TOOL_VER[$name]}")
		installed_ok "$bin" "${NMB_TOOL_PKG[$name]}" "${NMB_TOOL_VER[$name]}" || {
			printf '%s: installed binary is not %s@%s built with %s\n' "$name" "${NMB_TOOL_PKG[$name]}" "${NMB_TOOL_VER[$name]}" "$NMB_TOOLCHAIN"
			return 1
		}
		printf '%s: built\n' "$name"
	done
}

step_gosecvet() {
	mkdir -p -- "$(dirname -- "$NMB_GOSECVET")"
	(cd tools/gosecvet && go build -trimpath -o "$NMB_GOSECVET" .)
	go version -m "$NMB_GOSECVET" | sed -n '1,4p'
}

step_record() {
	local name
	for name in "${!NMB_TOOL_BIN[@]}" gosecvet; do
		go version -m "$(nmb_tool "$name")"
	done >"$NMB_OUT/tools.txt"
	printf 'recorded in tools.txt\n'
}

nmb_load_tools
steps=("$@")
[ "${#steps[@]}" -gt 0 ] || steps=(toolchain modules tools gosecvet record)
for s in "${steps[@]}"; do
	case "$s" in
	toolchain | modules | tools | gosecvet | record) nmb_run_step "prepare:$s" "PC, network (proxy and checksum DB)" "step_$s" ;;
	*) nmb_die "unknown step $s" ;;
	esac
done
nmb_finish
