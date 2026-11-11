#!/usr/bin/env bash
# The scan for personal data and secrets (roadmap Quality Bar 8) with tools/privscan.
# Usage: tools/scan.sh [STEP...]
#   tree     every tracked and untracked, not ignored file (the default with dist)
#   dist     every built file under dist/ (run tools/build-matrix.sh first)
#   history  every blob, commit message and commit identity of all refs; run before a push
# The owner's private list is read from the file named by NMB_PRIVATE_VALUES (outside the
# repository, never committed); its values are never printed.
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
nmb_init

scan() {
	"$NMB_BIN/privscan" -allow tools/privscan/allow.txt -repo "$NMB_ROOT" "$@"
}

step_tree() {
	local list="$NMB_OUT/scan-files.bin"
	nmb_files . >"$list"
	scan tree -0 <"$list"
}

step_dist() {
	[ -d dist ] || {
		printf 'no dist/ folder: run tools/build-matrix.sh first\n'
		return 1
	}
	scan dist dist
}

step_history() {
	scan history
}

nmb_build_helper privscan || nmb_die "cannot build tools/privscan"
steps=("$@")
[ "${#steps[@]}" -gt 0 ] || steps=(tree dist)
for s in "${steps[@]}"; do
	case "$s" in
	tree | dist | history) nmb_run_step "scan:$s" "PC" "step_$s" ;;
	*) nmb_die "unknown step $s" ;;
	esac
done
nmb_finish
