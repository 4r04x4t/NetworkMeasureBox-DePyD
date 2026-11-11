#!/usr/bin/env bash
# The requirements register of F-02 in the gate: docs/requirements/check.sh (the register is
# consistent with itself and with the Python repository at its source commit) and
# docs/requirements/check_test.sh (the check passes the real register and fails every planted fault).
# Usage: tools/register.sh [STEP...]
#   check  sh docs/requirements/check.sh          test  sh docs/requirements/check_test.sh
# Knob: NMB_SRC, the Python repository the register is read against (default
# $HOME/src/NetworkMeasureBox; read only, through git, at the commit the check names).
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
nmb_init

NMB_SRC=${NMB_SRC:-$HOME/src/NetworkMeasureBox}
[[ "$NMB_SRC" =~ ^/[A-Za-z0-9._/+-]+$ ]] || nmb_die "NMB_SRC must be an absolute path of plain characters"
case "$NMB_SRC/" in
"$NMB_ROOT"/*) nmb_die "NMB_SRC must lie outside this repository" ;;
esac
export NMB_SRC

step_check() {
	sh docs/requirements/check.sh
}

step_test() {
	sh docs/requirements/check_test.sh
}

steps=("$@")
[ "${#steps[@]}" -gt 0 ] || steps=(check test)
for s in "${steps[@]}"; do
	case "$s" in
	check | test) nmb_run_step "register:$s" "PC, Python repository at the register's source commit" "step_$s" ;;
	*) nmb_die "unknown step $s" ;;
	esac
done
nmb_finish
