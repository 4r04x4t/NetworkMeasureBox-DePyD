#!/usr/bin/env bash
# Runs every Fuzz* target of every package for NMB_FUZZTIME (default 30s) on amd64, offline and
# hermetic like tools/test.sh. A failing input is left in the package's testdata/fuzz/<Name>/ for
# the implementer to commit as a regression seed (gate skill, point 4).
# Usage: tools/fuzz.sh [PACKAGE...]   (default: every package of the module)
set -euo pipefail
# shellcheck source=tools/lib.sh
. "$(dirname -- "${BASH_SOURCE[0]}")/lib.sh"
nmb_init
nmb_go_offline "$NMB_ROOT/tools/fuzz.sh" "$@"

fuzz_one() {
	nmb_hermetic go test -run '^$' -fuzz "^$2\$" -fuzztime "$NMB_FUZZTIME" -parallel "$NMB_GATE_JOBS" "$1"
}

# targets PKG: the names of the package's fuzz targets.
targets() {
	local out
	out=$(nmb_hermetic go test -list '^Fuzz' -run '^$' "$1") || return 1
	grep -E '^Fuzz[A-Za-z0-9_]*$' <<<"$out" || true
}

nmb_hermetic_init || nmb_die "go env failed"
pkgs=("$@")
if [ "${#pkgs[@]}" -eq 0 ]; then
	list=$(go list ./...)
	mapfile -t pkgs <<<"$list"
fi
count=0
for pkg in "${pkgs[@]}"; do
	names=$(targets "$pkg") || nmb_die "cannot list the fuzz targets of $pkg"
	while IFS= read -r name; do
		[ -n "$name" ] || continue
		count=$((count + 1))
		nmb_run_step "fuzz:${pkg#nmb/}.$name" "PC amd64, $NMB_WHERE_NET, $NMB_FUZZTIME" fuzz_one "$pkg" "$name"
	done <<<"$names"
done
if [ "$count" -eq 0 ]; then
	nmb_run_step "fuzz:none" "PC amd64" printf 'no fuzz targets in %s packages\n' "${#pkgs[@]}"
fi
printf 'fuzz: %s targets run for %s each\n' "$count" "$NMB_FUZZTIME"
nmb_finish
