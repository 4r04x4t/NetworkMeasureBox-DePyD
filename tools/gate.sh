#!/usr/bin/env bash
# The whole gate of a package (work plan 2.4): runs every step, always all of them, in this order:
# prepare, deps, build matrix, analysis, tests, fuzzing, scans of the tree and of dist/, and the
# gate's own self-test; then prints what ran where with its exit code.
# Exit status: 0 all passed, 1 a step failed, 2 a usage or environment error (refusal to run as
# root, wrong toolchain, bad knob, a script that could not start).
# Knobs: NMB_GATE_JOBS (1-4, default 2), NMB_FUZZTIME (default 30s), NMB_TOOLS_DIR,
# NMB_PRIVATE_VALUES (see docs/DEVELOPMENT.md).
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
[ "$#" -eq 0 ] || nmb_die "usage: tools/gate.sh (no arguments; run the single scripts for parts)"
nmb_init

{
	printf 'date (UTC)\t%s\n' "$(date -u +%Y-%m-%dT%H:%M:%SZ)"
	printf 'commit\t%s\n' "$(git rev-parse HEAD)"
	printf 'tree\t%s\n' "$(if [ -z "$(git status --porcelain)" ]; then echo clean; else echo 'uncommitted changes'; fi)"
	printf 'machine\t%s, %s threads, %s MiB RAM\n' "$(uname -m)" "$(nproc)" "$(awk '/MemTotal/ {print int($2/1024)}' /proc/meminfo)"
	printf 'go\t%s\n' "$(go version)"
	printf 'jobs\t%s\nfuzz time\t%s\n' "$NMB_GATE_JOBS" "$NMB_FUZZTIME"
	printf 'private list\t%s\n' "$(if [ -n "${NMB_PRIVATE_VALUES:-}" ]; then echo used; else echo 'not used'; fi)"
} >"$NMB_OUT/environment.tsv"
cat -- "$NMB_OUT/environment.tsv"

worst=0
script() {
	local rc=0
	printf '\n## %s\n' "$*"
	"$NMB_ROOT/tools/$1" "${@:2}" || rc=$?
	if [ "$rc" -eq 2 ]; then
		worst=2
	elif [ "$rc" -ne 0 ] && [ "$worst" -eq 0 ]; then
		worst=1
	fi
	printf '%s\t%s\t%s\t-\n' "script:$1 ${*:2}" "PC" "$rc" >>"$NMB_OUT/summary.tsv"
}

script prepare.sh
script deps.sh
script build-matrix.sh
script analyse.sh
script test.sh
script fuzz.sh
script scan.sh tree dist
script gate-selftest.sh

printf '\n## Summary (%s)\n' "$NMB_OUT/summary.tsv"
failed=0
while IFS=$'\t' read -r name where rc secs; do
	mark=ok
	if [ "$rc" != 0 ]; then
		mark=FAILED
		failed=$((failed + 1))
	fi
	printf '%-7s rc %-3s %5ss  %-44s %s\n' "$mark" "$rc" "$secs" "$name" "$where"
done <"$NMB_OUT/summary.tsv"
if [ -f "$NMB_OUT/sizes.tsv" ]; then
	printf '\nBinary sizes (bytes, stripped, -trimpath):\n'
	cat -- "$NMB_OUT/sizes.tsv"
fi
printf '\nNot done by this gate: clock shifting (static binaries; fake clocks come with F-06); ICMP in the\n'
printf 'namespace (unprivileged ping is off there); mips and mipsle tests are compiled, not run; the\n'
printf 'history scan runs before a push (tools/scan.sh history), not per package.\n'
if [ "$worst" -eq 0 ] && [ "$failed" -ne 0 ]; then
	worst=1
fi
printf '\ngate: exit %s (%s failed entries)\n' "$worst" "$failed"
exit "$worst"
