#!/usr/bin/env bash
# Checks the dependency and licence list docs/DEPENDENCIES.md: every go.mod is tidy (its require
# list is exactly what is built) and tools/depcheck finds no unlisted module, stale row, unknown or
# disallowed licence, changed licence text, toolchain mismatch or gate tool of another version.
# Usage: tools/deps.sh [STEP...]   steps: tidy depcheck   (offline: module cache only)
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
nmb_init
export GOPROXY=off

step_tidy() {
	local gomod
	local -a mods
	mapfile -t mods < <(nmb_gomods)
	[ "${#mods[@]}" -gt 0 ] || return 1
	for gomod in "${mods[@]}"; do
		printf 'go mod tidy -diff in %s\n' "$(dirname -- "$gomod")"
		(cd -- "$(dirname -- "$gomod")" && go mod tidy -diff)
	done
}

step_depcheck() {
	local -a mods
	mapfile -t mods < <(nmb_gomods)
	[ "${#mods[@]}" -gt 0 ] || return 1
	"$NMB_BIN/depcheck" -list docs/DEPENDENCIES.md -tools "$NMB_TOOLS_DIR" check "${mods[@]}"
}

nmb_load_tools
steps=("$@")
[ "${#steps[@]}" -gt 0 ] || steps=(tidy depcheck)
for s in "${steps[@]}"; do
	case "$s" in
	tidy | depcheck) nmb_run_step "deps:$s" "PC, GOPROXY=off (module cache)" "step_$s" ;;
	*) nmb_die "unknown step $s" ;;
	esac
done
nmb_finish
