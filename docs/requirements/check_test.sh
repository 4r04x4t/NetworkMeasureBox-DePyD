#!/bin/sh
# Test of check.sh: the real register must pass, and every planted fault in a copy must fail.
# Usage, from the worktree root:   sh docs/requirements/check_test.sh [register-dir]
set -u
R=${1:-docs/requirements}
CHECK="$R/check.sh"
[ -f "$CHECK" ] || { echo "check_test: no $CHECK" >&2; exit 2; }
T=$(mktemp -d "${TMPDIR:-/tmp}/nmbregtest.XXXXXX") || exit 2
trap 'rm -rf "$T"' EXIT INT TERM
failures=0

sh "$CHECK" "$R" > "$T/real.out" 2>&1
rc=$?
if [ "$rc" -eq 0 ]; then echo "ok   real register passes"; else echo "FAIL real register: exit $rc"; cat "$T/real.out"; failures=$((failures + 1)); fi

# plant NAME FILE SED-SCRIPT: copy the register, apply the sed script to FILE, expect exit 1
plant() {
	name=$1; file=$2; script=$3
	rm -rf "$T/reg"; mkdir "$T/reg"; cp -R "$R/." "$T/reg/"
	cp "$T/reg/$file" "$T/before"
	sed "$script" "$T/before" > "$T/reg/$file"
	if cmp -s "$T/before" "$T/reg/$file"; then echo "FAIL $name: the plant changed nothing"; failures=$((failures + 1)); return; fi
	sh "$CHECK" "$T/reg" > "$T/out" 2>&1
	rc=$?
	if [ "$rc" -eq 1 ]; then echo "ok   $name: exit 1 ($(grep -m1 '^FAIL' "$T/out"))"
	else echo "FAIL $name: expected exit 1, got $rc"; cat "$T/out"; failures=$((failures + 1)); fi
}

plant "id defined twice" parts/changelog.md '/^| REQ-CL-0001 |/{p;}'
plant "id defined twice across parts" parts/testing.md 's/^| REQ-TS-0001 |/| REQ-CL-0001 |/'
plant "coverage gap" parts/changelog.md '/^| 2 | blank |/d'
plant "coverage overlap" parts/roadmap-1-507.md 's/^| 9-10 |/| 8-10 |/'
plant "coverage stops early" parts/testing.md '/^| 400-401 |/d; /^| 400 |/d; /^| 401 |/d'
plant "id in no coverage row" parts/testing.md '/^## Coverage/,$s/REQ-TS-0001\([^0-9]\)/context: planted\1/g'
plant "coverage names an undefined id" parts/testing.md 's/^\(| 3-4 | [^|]*| \)REQ-TS-0001 |/\1REQ-TS-0001, REQ-TS-9999 |/'
plant "requirement row without package" parts/testing.md '/^| REQ-TS-0001 |/s/| P-13 |/|  |/'
plant "requirement row with unknown package" parts/testing.md '/^| REQ-TS-0001 |/s/| P-13 |/| Z-99 |/'
plant "requirement row with five cells" parts/testing.md '/^| REQ-TS-0001 |/s/ | P-13 | / P-13 | /'
plant "owner changed, BY_PACKAGE.md not regenerated" parts/testing.md '/^| REQ-TS-0001 |/s/| P-13 |/| P-12 |/'
plant "BY_PACKAGE.md loses an id" BY_PACKAGE.md 's/REQ-TS-0001, //'
plant "BY_PACKAGE.md drops a work plan package" BY_PACKAGE.md 's/^### C-10 /### C-1O /'
plant "Python file missing" parts/python-files.md '/^| `\.gitignore`@ad3758c |/d'
plant "Python file listed twice" parts/python-files.md '/^| `\.gitignore`@ad3758c |/{p;}'
plant "Python file lines missing" parts/python-files.md 's/^| `\.gitignore`@ad3758c | 1-86 |/| `.gitignore`@ad3758c | 1-80 |/'
plant "Python file that does not exist" parts/python-files.md '/^| `\.gitignore`@ad3758c |/{p;s/\.gitignore/.gitignored/;}'
plant "roadmap section not mapped" PACKAGE_CHANGES.md '/^| 116 | 116-/d'
plant "roadmap section mapped twice" PACKAGE_CHANGES.md '/^| 116 | 116-/{p;}'
plant "TESTING.md row not mapped" PACKAGE_CHANGES.md '/^## M2\./,/^## M3\./{/^| 400 |/d;}'
plant "Python file not mapped" PACKAGE_CHANGES.md '/^| `setup.sh` |/d'
plant "Python file mapped to unknown package" PACKAGE_CHANGES.md 's/^| `setup.sh` | P-11 |/| `setup.sh` | Q-11 |/'

if [ "$failures" -eq 0 ]; then echo "check_test: OK"; exit 0; fi
echo "check_test: $failures FAILED"
exit 1
