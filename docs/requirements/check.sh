#!/bin/sh
# Check of the F-02 requirements register (POSIX sh and awk).
#
# Usage, from the worktree root:   sh docs/requirements/check.sh [register-dir]
#   register-dir  defaults to docs/requirements
#   NMB_SRC       the Python repository (default: $HOME/src/NetworkMeasureBox), read only, at commit ad3758c
#
# Fails (exit 1) when:
#   1. a requirement id is defined twice, or a requirement row is malformed (not six cells, bad id);
#   2. a requirement row has no package, or a package that BY_PACKAGE.md does not list;
#   3. a coverage table has a gap or an overlap, does not start or end where its source does, or a
#      defined id is in no coverage row, or a coverage row names an id that is not defined, or a line that a
#      requirement cites from its part's own source lies in a coverage row that does not name the id;
#   4. a Python file of the source commit (without docs/ and the top-level *.md) is missing from
#      parts/python-files.md, a listed file does not exist, or a line of a file is listed twice or not at all;
#   5. BY_PACKAGE.md disagrees with the parts (owner and Also-in lists, counts), or a work plan package
#      has no section in it;
#   6. PACKAGE_CHANGES.md does not map every roadmap section, every TESTING.md row and every Python
#      file to exactly one valid package, or an M2 row's owners column or package differs from what the
#      owners of its REQ-TS rows in parts/testing.md give, or an M1 row's "Rows by owner" differs from the
#      owners of the REQ-RM rows that cite a line of the section (both added in round 2);
#   7. a document of the source commit (a top-level *.md or a file under docs/) is neither the source of a
#      coverage table nor named in OUT_OF_SCOPE below.
# Exit 2: the register or the source cannot be read.
set -u
R=${1:-docs/requirements}
SRC=${NMB_SRC:-$HOME/src/NetworkMeasureBox}
COMMIT=ad3758c
# Documents with no coverage table of their own, each with its reason in the register README ("Sources").
# None since round 3: parts/other-docs.md covers the last four (the NMHH comparison and three Hungarian translations).
OUT_OF_SCOPE=""
P="$R/parts"
PARTS="roadmap-1-507.md roadmap-508-1034.md porting-plan.md workplan-architecture-devsetup.md testing.md readme.md changelog.md contributing.md other-docs.md"

for f in $PARTS python-files.md; do
	[ -f "$P/$f" ] || { echo "check: missing part $P/$f" >&2; exit 2; }
done
for f in BY_PACKAGE.md PACKAGE_CHANGES.md; do
	[ -f "$R/$f" ] || { echo "check: missing $R/$f" >&2; exit 2; }
done
git -C "$SRC" rev-parse -q --verify "$COMMIT^{commit}" >/dev/null 2>&1 || {
	echo "check: Python repository $SRC has no commit $COMMIT (set NMB_SRC)" >&2; exit 2; }

T=$(mktemp -d "${TMPDIR:-/tmp}/nmbreg.XXXXXX") || exit 2
trap 'rm -rf "$T"' EXIT INT TERM
fail=0
bad() { echo "FAIL: $*"; fail=1; }

# Shared awk functions: package keys of a Package or Also-in cell.
cat > "$T/lib.awk" <<'EOF'
function trim(s) { sub(/^[ \t]+/, "", s); sub(/[ \t]+$/, "", s); return s }
# keys of one item, separated by spaces
function keys(item,   a, p, n1, n2, k, out) {
	item = trim(item)
	if (item ~ /^(all|every) (.* )?packages?([^a-z]|$)/) return "ALL"
	sub(/^NEW: /, "", item)
	if (item ~ /^S-02(\.\.|…)S-0n/) return "S-02…S-0n"
	if (item ~ /^U-02(\.\.|…)U-0n/) return "U-02…U-0n"
	if (item ~ /^[A-Z]-[0-9][0-9](\.\.|…)[A-Z]-[0-9][0-9]$/ && substr(item, 1, 1) == substr(item, length(item) - 3, 1)) {
		p = substr(item, 1, 1); n1 = substr(item, 3, 2) + 0; n2 = substr(item, length(item) - 1, 2) + 0
		out = ""
		for (k = n1; k <= n2; k++) out = out (out == "" ? "" : " ") sprintf("%s-%02d", p, k)
		return out
	}
	split(item, a, " ")
	return a[1]
}
# split a cell on commas outside parentheses into it[1..n]; returns n (items "-" and "" dropped)
function items(s, it,   i, ch, d, cur, n) {
	d = 0; cur = ""; n = 0
	for (i = 1; i <= length(s); i++) {
		ch = substr(s, i, 1)
		if (ch == "(") d++
		if (ch == ")") d--
		if (ch == "," && d == 0) { cur = trim(cur); if (cur != "" && cur != "-") it[++n] = cur; cur = "" }
		else cur = cur ch
	}
	cur = trim(cur); if (cur != "" && cur != "-") it[++n] = cur
	return n
}
EOF

# ---- 1, 2: requirement rows -> T/req (id part line owner) and T/pairs (id own|also key)
cat > "$T/rows.awk" <<'EOF'
/^\| REQ-/ {
	line = $0; gsub(/\\\|/, "\001", line)
	n = split(line, c, "|")
	id = trim(c[2])
	if (id !~ /^REQ-(RM|PP|WP|AR|DV|TS|RE|CL|CT|NH)-[0-9][0-9][0-9][0-9]$/) { print "BAD " FILENAME ":" FNR " bad id cell: " id; next }
	if (n != 8) { print "BAD " FILENAME ":" FNR " " id " has " n - 2 " cells, not 6"; next }
	if (id in seen) { print "BAD " id " defined twice: " seen[id] " and " FILENAME ":" FNR; next }
	seen[id] = FILENAME ":" FNR
	pk = trim(c[5])
	if (pk == "" || pk == "-") { print "BAD " FILENAME ":" FNR " " id " has no package"; next }
	split(keys(pk), o, " "); own = o[1]
	print "REQ " id " " FILENAME " " FNR " " own
	s = c[4]
	while (match(s, /[A-Za-z0-9_.\/-]+@ad3758c:[0-9]+(-[0-9]+)?/)) {
		ref = substr(s, RSTART, RLENGTH); s = substr(s, RSTART + RLENGTH)
		f = ref; sub(/@.*$/, "", f); r = ref; sub(/^.*:/, "", r)
		a = r; b = r; if (index(r, "-")) { a = substr(r, 1, index(r, "-") - 1); b = substr(r, index(r, "-") + 1) }
		print "SRC " id " " FILENAME " " f " " a + 0 " " b + 0
	}
	print "PAIR " id " own " own
	split("", done)
	m = items(trim(c[6]), it)
	for (i = 1; i <= m; i++) {
		k = split(keys(it[i]), kk, " ")
		for (j = 1; j <= k; j++) if (kk[j] != own && !(kk[j] in done)) { done[kk[j]] = 1; print "PAIR " id " also " kk[j] }
	}
}
EOF
(cd "$P" && awk -f "$T/lib.awk" -f "$T/rows.awk" $PARTS) > "$T/rows" || exit 2
if grep -q '^BAD ' "$T/rows"; then grep '^BAD ' "$T/rows" | sed 's/^BAD /FAIL: /'; fail=1; fi
grep '^REQ ' "$T/rows" | cut -d' ' -f2- > "$T/req"
grep '^SRC ' "$T/rows" > "$T/srcs"
grep '^PAIR ' "$T/rows" | cut -d' ' -f2- | sort > "$T/pairs"
nreq=$(wc -l < "$T/req" | tr -d ' ')

# ---- BY_PACKAGE.md: headings and lists
awk '
/^## / { key = ""; mode = ""; next }
/^### / { key = $2; print "HEAD " key; mode = ""; next }
key != "" && /^Owns [0-9]+:/ { mode = "own"; want = $2 + 0; sub(/:$/, "", want); print "COUNT " key " own " want + 0; next }
key != "" && /^Also in [0-9]+:/ { mode = "also"; want = $3 + 0; print "COUNT " key " also " want + 0; next }
mode != "" && /^REQ-/ {
	s = $0
	while (match(s, /REQ-[A-Z][A-Z]-[0-9][0-9][0-9][0-9]/)) { print "PAIR " substr(s, RSTART, RLENGTH) " " mode " " key; s = substr(s, RSTART + RLENGTH) }
	next
}
{ mode = "" }' "$R/BY_PACKAGE.md" > "$T/bp"
grep '^HEAD ' "$T/bp" | cut -d' ' -f2 | sort > "$T/heads"
sort -u "$T/heads" > "$T/keys"
dups=$(uniq -d "$T/heads"); [ -z "$dups" ] || bad "BY_PACKAGE.md has a package twice: $dups"
grep '^PAIR ' "$T/bp" | cut -d' ' -f2- | sort > "$T/bppairs"

# 2: every owner and Also-in key is a package of BY_PACKAGE.md
awk 'NR == FNR { k[$1] = 1; next } !($3 in k) { print "FAIL: " $1 " names package " $3 " (" $2 "), which BY_PACKAGE.md does not list" }' \
	"$T/keys" "$T/pairs" > "$T/o"; [ -s "$T/o" ] && { cat "$T/o"; fail=1; }

# 5: BY_PACKAGE.md agrees with the parts
if ! cmp -s "$T/pairs" "$T/bppairs"; then
	bad "BY_PACKAGE.md disagrees with the parts (< parts, > BY_PACKAGE.md):"
	diff "$T/pairs" "$T/bppairs" | grep '^[<>]' | head -20
fi
awk '$1 == "COUNT" { want[$2 " " $3] = $4; next } $1 == "PAIR" { got[$4 " " $3]++ }
END { for (k in want) if (want[k] != got[k] + 0) print "FAIL: BY_PACKAGE.md " k " says " want[k] " ids, lists " got[k] + 0 }' "$T/bp" > "$T/o"
[ -s "$T/o" ] && { cat "$T/o"; fail=1; }
for k in F-00 F-01 F-02 F-03 F-04 F-05 F-06 F-07 C-01 C-02 C-03 C-04 C-05 C-06 C-07 C-08 C-09 C-10 C-11 \
	M-01 M-02 M-03 M-04 M-05 M-06 M-07 M-08 M-09 M-10 M-11 M-12 M-13 M-14 \
	N-01 N-02 N-03 N-04 N-05 N-06 N-07 N-08 N-09 N-10 N-11 N-12 N-13 N-14 N-15 N-16 N-17 N-18 \
	R-01 R-02 R-03 R-04 R-05 S-01 'S-02…S-0n' U-01 'U-02…U-0n' S-10 S-11 S-12 S-13 \
	P-01 P-02 P-03 P-04 P-05 P-06 P-07 P-08 P-09 P-10 P-11 P-12 P-13 \
	I-01 I-02 I-03 I-04 I-05 I-06 I-07 I-08 I-09 I-10; do
	grep -qxF -- "$k" "$T/keys" || bad "work plan package $k has no section in BY_PACKAGE.md"
done

# ---- 3: coverage tables
# manifest: part table source start end(N = last line of the source) mode
cat > "$T/manifest" <<'EOF'
roadmap-1-507.md 1 docs/DEPYTHONIZATION_ROADMAP.md 1 507 contig
roadmap-508-1034.md 1 docs/DEPYTHONIZATION_ROADMAP.md 508 N contig
porting-plan.md 1 docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md 1 N contig
workplan-architecture-devsetup.md 1 docs/NATIVE_WORK_PLAN.md 1 N contig
workplan-architecture-devsetup.md 2 docs/ARCHITECTURE.md 1 N contig
workplan-architecture-devsetup.md 3 docs/DEV_SETUP_WINDOWS.md 1 N contig
testing.md 1 docs/TESTING.md 1 N contig
readme.md 1 README.md 1 N contig
readme.md 2 README.hu.md 1 N sparse
changelog.md 1 CHANGELOG.md 1 N contig
contributing.md 1 CONTRIBUTING.md 1 N contig
contributing.md 2 CONTRIBUTING.hu.md 1 N contig
other-docs.md 1 docs/NMHH_SZELESSAVNET_COMPARISON.md 1 N contig
other-docs.md 2 docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md 1 N contig
other-docs.md 3 docs/DEV_SETUP_WINDOWS.hu.md 1 N contig
other-docs.md 4 docs/TESTING.hu.md 1 N contig
EOF
: > "$T/manifest2"
while read -r part tab src start end mode; do
	n=$(git -C "$SRC" show "$COMMIT:$src" 2>/dev/null | awk 'END { print NR }')
	[ -n "$n" ] && [ "$n" -gt 0 ] || { echo "check: cannot read $src at $COMMIT" >&2; exit 2; }
	[ "$end" = N ] && end=$n
	[ "$end" -le "$n" ] || bad "manifest: $part table $tab ends at $end, source $src has $n lines"
	echo "$part $tab $src $start $end $mode" >> "$T/manifest2"
done < "$T/manifest"
(cd "$P" && awk '
function emit(id) { print "ID " id; print "CID " rowkey " " id }
FNR == 1 { cov = 0; t = 0 }
/^## / { cov = ($0 ~ /^## Coverage/); next }
cov && /^\| Lines \|/ { t++; next }
cov && /^\| [0-9]+(-[0-9]+)? \|/ {
	split($0, c, "|"); r = c[2]; gsub(/ /, "", r)
	a = r; b = r; if (index(r, "-")) { a = substr(r, 1, index(r, "-") - 1); b = substr(r, index(r, "-") + 1) }
	print "ROW " FILENAME " " t " " a + 0 " " b + 0 " " FNR
	rowkey = FILENAME " " t " " a + 0 " " b + 0
	# ids: full ids, and the short forms "REQ-XX-0001 to -0004", "REQ-XX-0001, -0003", "... to REQ-XX-0009"
	s = $0; pre = ""; last = 0; gap = ""
	while (match(s, /((,|to|and)[ ]*)?(REQ-[A-Z][A-Z]-)?-?[0-9][0-9][0-9][0-9]/)) {
		tok = substr(s, RSTART, RLENGTH); before = substr(s, 1, RSTART - 1); s = substr(s, RSTART + RLENGTH)
		conn = ""; if (tok ~ /^(,|to|and)/) { conn = tok; sub(/[ ]*(REQ-[A-Z][A-Z]-)?-?[0-9]+$/, "", conn) }
		num = tok; sub(/^.*[^0-9]/, "", num)
		adj = (pre != "" && before ~ /^[ ]*$/)
		if (tok ~ /REQ-[A-Z][A-Z]-/) {
			p = tok; sub(/^.*REQ-/, "REQ-", p); p = substr(p, 1, 7)
			print "LIT " p num " " FILENAME ":" FNR
			if (conn == "to" && adj && p == pre) for (k = last + 1; k < num + 0; k++) emit(sprintf("%s%04d", pre, k))
			emit(p num); pre = p; last = num + 0
		} else if (conn != "" && adj) {
			if (conn == "to") for (k = last + 1; k <= num + 0; k++) emit(sprintf("%s%04d", pre, k))
			else emit(sprintf("%s%04d", pre, num + 0))
			last = num + 0
		} else pre = ""
	}
}' $PARTS) > "$T/cov" || exit 2
awk '
NR == FNR { key = $1 " " $2; src[key] = $3; st[key] = $4; en[key] = $5; md[key] = $6; order[++m] = key; next }
$1 == "ROW" {
	key = $2 " " $3
	if (!(key in src)) { print "FAIL: " $2 ":" $6 " coverage row in table " $3 ", which the manifest does not know"; next }
	a = $4; b = $5; seen[key]++
	if (b < a) print "FAIL: " $2 ":" $6 " coverage range " a "-" b " runs backwards"
	if (md[key] == "contig") {
		want = (key in prev) ? prev[key] + 1 : st[key]
		if (a > want) print "FAIL: " $2 " table " $3 ": gap, lines " want "-" a - 1 " of " src[key] " are in no row"
		if (a < want) print "FAIL: " $2 " table " $3 ": overlap at line " a " of " src[key] " (row " $6 ")"
	} else {
		if ((key in prev) && a <= prev[key]) print "FAIL: " $2 " table " $3 ": overlap or wrong order at line " a " (row " $6 ")"
		if (a < 1 || b > en[key]) print "FAIL: " $2 " table " $3 ": line " a "-" b " outside " src[key]
	}
	if (!(key in prev) || b > prev[key]) prev[key] = b
}
END {
	for (i = 1; i <= m; i++) {
		key = order[i]
		if (!seen[key]) { print "FAIL: " key ": coverage table missing"; continue }
		if (md[key] == "contig" && prev[key] != en[key]) print "FAIL: " key ": coverage ends at " prev[key] ", " src[key] " range ends at " en[key]
	}
}' "$T/manifest2" "$T/cov" > "$T/o"; [ -s "$T/o" ] && { cat "$T/o"; fail=1; }
cut -d' ' -f1 "$T/req" | sort -u > "$T/defined"
grep '^ID ' "$T/cov" | cut -d' ' -f2 | sort -u > "$T/covered"
comm -23 "$T/defined" "$T/covered" > "$T/o"; [ -s "$T/o" ] && { bad "ids in no coverage row: $(tr '\n' ' ' < "$T/o")"; }
grep '^LIT ' "$T/cov" | sort -k2 | awk 'NR == FNR { d[$1] = 1; next } !($2 in d) { print "FAIL: coverage row " $3 " names undefined id " $2 }' \
	"$T/defined" - > "$T/o"; [ -s "$T/o" ] && { cat "$T/o"; fail=1; }
# 3b: a line a requirement cites from its own part's source lies in a coverage row that names the id
awk '
FILENAME == ARGV[1] { tsrc[$1 " " $2] = $3; tabs[$1] = tabs[$1] " " $2; next }
$1 == "ROW" { k = $2 " " $3; for (x = $4; x <= $5; x++) { row[k, x] = $4 "-" $5; has[k, x] = 1 }; next }
$1 == "CID" { k = $2 " " $3; for (x = $4; x <= $5; x++) named[k, x, $6] = 1; next }
$1 == "SRC" {
	id = $2; part = $3; f = $4; a = $5; b = $6; nt = split(tabs[part], tl, " ")
	for (i = 1; i <= nt; i++) {
		k = part " " tl[i]; if (tsrc[k] != f) continue
		for (x = a; x <= b; x++) if (has[k, x] && !named[k, x, id]) {
			print "FAIL: " id " cites " f ":" x ", but the coverage row " row[k, x] " of " part " does not name it"; break }
	}
}' "$T/manifest2" "$T/cov" "$T/srcs" > "$T/o"; [ -s "$T/o" ] && { cat "$T/o"; fail=1; }

# ---- 4: Python files
git -C "$SRC" ls-tree -r --name-only "$COMMIT" | grep -v '^docs/' | grep -v '^[^/]*\.md$' > "$T/files" || exit 2
# line counts of the text files (git grep counts an unterminated last line too); binary and empty files count 0
git -C "$SRC" grep -I -c '' "$COMMIT" > "$T/gg" || exit 2
awk 'NR == FNR { s = $0; sub(/^[^:]*:/, "", s); n = s; sub(/^.*:/, "", n); p = s; sub(/:[0-9]+$/, "", p); c[p] = n; next }
{ printf "%s\t%s\n", (($0 in c) ? c[$0] : 0), $0 }' "$T/gg" "$T/files" > "$T/nlines"
awk -F'|' '
/^## / { sect = ($0 ~ /^## Files and packages/); next }
sect && /^\| `/ {
	path = $2; sub(/^[^`]*`/, "", path); sub(/`.*$/, "", path)
	r = $3; gsub(/^ +| +$/, "", r); pk = $4; gsub(/^ +/, "", pk); sub(/ .*$/, "", pk)
	print path "\t" r "\t" pk "\t" FNR
}' "$P/python-files.md" > "$T/pyrows"
awk -F'\t' '
FILENAME == ARGV[1] { n[$2] = $1; next }
FILENAME == ARGV[2] { k[$1] = 1; next }
{
	f = $1; r = $2; pk = $3; ln = $4
	if (!(f in n)) { print "FAIL: python-files.md:" ln " lists " f ", which is not a file of the commit"; next }
	if (!(pk in k)) print "FAIL: python-files.md:" ln " " f " has package " pk ", which BY_PACKAGE.md does not list"
	rows[f]++
	if (r == "binary" || r ~ /^0 \(empty\)$/) {
		special[f]++
		if (n[f] != 0) print "FAIL: python-files.md:" ln " " f " marked " r " but is a text file of " n[f] " lines"
		next
	}
	m = split(r, part, ",")
	for (i = 1; i <= m; i++) {
		x = part[i]; gsub(/ /, "", x)
		a = x; b = x; if (index(x, "-")) { a = substr(x, 1, index(x, "-") - 1); b = substr(x, index(x, "-") + 1) }
		if (a !~ /^[0-9]+$/ || b !~ /^[0-9]+$/ || b + 0 < a + 0) { print "FAIL: python-files.md:" ln " " f " bad range " x; continue }
		if (b + 0 > n[f]) print "FAIL: python-files.md:" ln " " f " range " x " beyond its " n[f] " lines"
		for (j = a + 0; j <= b + 0; j++) { if ((f, j) in mark) dupl[f] = dupl[f] " " j; else { mark[f, j] = 1; cnt[f]++ } }
	}
}
END {
	for (f in n) {
		if (!(f in rows)) { print "FAIL: python file " f " is missing from python-files.md"; continue }
		if (f in special) { if (rows[f] != 1) print "FAIL: " f " (binary or empty) is listed " rows[f] " times"; continue }
		if (f in dupl) print "FAIL: " f " lists lines twice:" substr(dupl[f], 1, 60)
		if (cnt[f] + 0 != n[f]) print "FAIL: " f " covers " cnt[f] + 0 " of its " n[f] " lines"
	}
}' "$T/nlines" "$T/keys" "$T/pyrows" > "$T/o"; [ -s "$T/o" ] && { cat "$T/o"; fail=1; }

# ---- 6: PACKAGE_CHANGES.md mappings
git -C "$SRC" show "$COMMIT:docs/DEPYTHONIZATION_ROADMAP.md" | awk '/^```/ { f = !f; next } !f && /^###? / { print NR }' > "$T/rmheads"
git -C "$SRC" show "$COMMIT:docs/TESTING.md" | awk '
/^## / { lim = ($0 ~ /^## Known limits/) }
/^\|/ { if (!tab) { tab = 1; next } if ($0 ~ /^\|[ ]*:?-+/) next; print NR; next }
{ tab = 0; if (lim && ($0 ~ /^- / || $0 ~ /^[0-9]+\. /)) print NR }' > "$T/tsrows"
awk -F'|' '
/^## / { s = ""; if ($0 ~ /^## M1\./) s = "M1"; if ($0 ~ /^## M2\./) s = "M2"; if ($0 ~ /^## M3\./) s = "M3"; next }
s == "M1" && /^\| [0-9]+ \|/ { x = $2; gsub(/ /, "", x); pk = $5; gsub(/ /, "", pk); print "M1\t" x "\t" pk }
s == "M2" && /^\| [0-9]+ \|/ { x = $2; gsub(/ /, "", x); pk = $3; gsub(/ /, "", pk); print "M2\t" x "\t" pk }
s == "M3" && /^\| `/ { x = $2; sub(/^[^`]*`/, "", x); sub(/`.*$/, "", x); pk = $3; gsub(/ /, "", pk); print "M3\t" x "\t" pk }
' "$R/PACKAGE_CHANGES.md" > "$T/maps"
awk -F'\t' '
FILENAME == ARGV[1] { k[$1] = 1; next }
FILENAME == ARGV[2] { want["M1\t" $1] = "roadmap heading at line " $1; next }
FILENAME == ARGV[3] { want["M2\t" $1] = "TESTING.md row at line " $1; next }
FILENAME == ARGV[4] { want["M3\t" $2] = "Python file " $2; next }
{
	key = $1 "\t" $2
	if (!(key in want)) { print "FAIL: PACKAGE_CHANGES.md " $1 " maps " $2 ", which is no section, row or file of the source"; next }
	if (++got[key] == 2) print "FAIL: PACKAGE_CHANGES.md maps the " want[key] " twice"
	if (!($3 in k)) print "FAIL: PACKAGE_CHANGES.md maps the " want[key] " to " $3 ", which BY_PACKAGE.md does not list"
}
END { for (key in want) if (!(key in got)) print "FAIL: PACKAGE_CHANGES.md does not map the " want[key] }
' "$T/keys" "$T/rmheads" "$T/tsrows" "$T/nlines" "$T/maps" > "$T/o"; [ -s "$T/o" ] && { sort "$T/o"; fail=1; }
# 6b: each M2 row's owners column and package agree with the owners of its REQ-TS rows in parts/testing.md
# (owners: "<package> <count>", most first, then by name; package: the most frequent owner, on a tie the first id's owner)
awk '
FILENAME == ARGV[1] { if ($2 == "testing.md") { id = $1; sub(/^REQ-/, "", id); own[id] = $4 }; next }
/^## / { s = ($0 ~ /^## M2\./); next }
s && /^\| [0-9]+ \|/ {
	split($0, c, "|"); ln = c[2]; gsub(/ /, "", ln); pk = c[3]; gsub(/ /, "", pk)
	ow = c[5]; gsub(/^ +| +$/, "", ow)
	n = split(c[4], ids, ","); split("", cnt); nk = 0; bad = ""
	for (i = 1; i <= n; i++) {
		x = ids[i]; gsub(/ /, "", x)
		if (!(x in own)) { bad = x; break }
		if (!(own[x] in cnt)) key[++nk] = own[x]
		cnt[own[x]]++
	}
	if (bad != "") { print "FAIL: PACKAGE_CHANGES.md M2 line " ln " names " bad ", which parts/testing.md does not define"; next }
	for (i = 1; i <= nk; i++) for (j = i + 1; j <= nk; j++)
		if (cnt[key[j]] > cnt[key[i]] || (cnt[key[j]] == cnt[key[i]] && key[j] < key[i])) { t = key[i]; key[i] = key[j]; key[j] = t }
	want = ""; for (i = 1; i <= nk; i++) want = want (i > 1 ? ", " : "") key[i] " " cnt[key[i]]
	x = ids[1]; gsub(/ /, "", x)
	wpk = (nk > 1 && cnt[key[1]] == cnt[key[2]]) ? own[x] : key[1]
	if (ow != want) print "FAIL: PACKAGE_CHANGES.md M2 line " ln " gives owners \"" ow "\", the parts give \"" want "\""
	if (pk != wpk) print "FAIL: PACKAGE_CHANGES.md M2 line " ln " maps to " pk ", its rule gives " wpk
}' "$T/req" "$R/PACKAGE_CHANGES.md" > "$T/o"; [ -s "$T/o" ] && { cat "$T/o"; fail=1; }
# 6c: each M1 row's "Rows by owner" column agrees with the owners of the REQ-RM rows that cite a line of the section
awk '
FILENAME == ARGV[1] { own[$1] = $4; next }
FILENAME == ARGV[2] { if ($4 == "docs/DEPYTHONIZATION_ROADMAP.md" && $2 ~ /^REQ-RM-/) { nr[$2]++; ra[$2, nr[$2]] = $5; rb[$2, nr[$2]] = $6 }; next }
/^## / { s = ($0 ~ /^## M1\./); next }
s && /^\| [0-9]+ \|/ {
	split($0, c, "|"); ln = c[2]; gsub(/ /, "", ln); r = c[3]; gsub(/ /, "", r)
	a = r; b = r; if (index(r, "-")) { a = substr(r, 1, index(r, "-") - 1); b = substr(r, index(r, "-") + 1) }
	a += 0; b += 0; have = c[7]; gsub(/^ +| +$/, "", have)
	split("", cnt); nk = 0
	for (id in nr) for (i = 1; i <= nr[id]; i++) if (ra[id, i] <= b && rb[id, i] >= a) {
		if (!(own[id] in cnt)) key[++nk] = own[id]
		cnt[own[id]]++; break
	}
	for (i = 1; i <= nk; i++) for (j = i + 1; j <= nk; j++)
		if (cnt[key[j]] > cnt[key[i]] || (cnt[key[j]] == cnt[key[i]] && key[j] < key[i])) { t = key[i]; key[i] = key[j]; key[j] = t }
	want = ""; for (i = 1; i <= nk; i++) want = want (i > 1 ? ", " : "") key[i] " " cnt[key[i]]
	if (want == "") want = "none"
	if (have != want) print "FAIL: PACKAGE_CHANGES.md M1 line " ln " gives rows by owner \"" substr(have, 1, 60) "\", the parts give \"" substr(want, 1, 60) "\""
}' "$T/req" "$T/srcs" "$R/PACKAGE_CHANGES.md" > "$T/o"; [ -s "$T/o" ] && { cat "$T/o"; fail=1; }

# ---- 7: every document of the commit has a coverage table or is named out of scope
cut -d' ' -f3 "$T/manifest2" | sort -u > "$T/docsrc"
for d in $OUT_OF_SCOPE; do echo "$d"; done >> "$T/docsrc"
git -C "$SRC" ls-tree -r --name-only "$COMMIT" | grep -E '^docs/|^[^/]*\.md$' | while read -r d; do
	grep -qxF -- "$d" "$T/docsrc" || echo "FAIL: document $d of the commit has no coverage table and is not named out of scope"
done > "$T/o"; [ -s "$T/o" ] && { cat "$T/o"; fail=1; }

nfiles=$(wc -l < "$T/files" | tr -d ' ')
if [ "$fail" -eq 0 ]; then
	echo "check: OK: $nreq requirements, $(wc -l < "$T/manifest2" | tr -d ' ') coverage tables, $nfiles Python files," \
		"$(wc -l < "$T/rmheads" | tr -d ' ') roadmap sections, $(wc -l < "$T/tsrows" | tr -d ' ') TESTING.md rows"
	exit 0
fi
echo "check: FAILED"
exit 1
