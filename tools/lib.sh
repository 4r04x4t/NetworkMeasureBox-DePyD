# shellcheck shell=bash
# shellcheck disable=SC2034 # the variables set here are read by the scripts that source this file
# Common part of the gate scripts (sourced, never run). See docs/DEVELOPMENT.md.
#
# Rules for every script: set -euo pipefail; a step's status is the exit code of the step itself
# (never a pipe with tail, tee or grep in front of it); file lists are NUL-separated; every variable
# is quoted; no eval; temporary folders are removed by a trap.

nmb_die() {
	printf 'gate: %s\n' "$*" >&2
	exit 2
}

# nmb_init: refusals, repository root, knobs, toolchain, result folder. Every script calls it first.
nmb_init() {
	if [ "$(id -u)" -eq 0 ]; then
		nmb_die "refuses to run as root (fault KF-0001: a test run as root overwrote the real settings)"
	fi
	local here
	here=$(dirname -- "${BASH_SOURCE[0]}")
	NMB_ROOT=$(git -C "$here" rev-parse --show-toplevel) || nmb_die "not inside a git work tree"
	cd -- "$NMB_ROOT" || nmb_die "cannot enter the repository root"
	nmb_knobs
	nmb_toolchain
	nmb_outdir
	NMB_FAILED=0
	NMB_STEPS_RUN=0
	NMB_BIN="$NMB_OUT/bin"
	mkdir -p -- "$NMB_BIN"
}

# nmb_knobs validates the environment knobs; a bad value is an error, never a default.
nmb_knobs() {
	NMB_GATE_JOBS=${NMB_GATE_JOBS:-2}
	[[ "$NMB_GATE_JOBS" =~ ^[1-4]$ ]] || nmb_die "NMB_GATE_JOBS must be 1 to 4"
	NMB_FUZZTIME=${NMB_FUZZTIME:-30s}
	[[ "$NMB_FUZZTIME" =~ ^[1-9][0-9]*[sm]$ ]] || nmb_die "NMB_FUZZTIME must look like 30s or 2m"
	NMB_TOOLS_DIR=${NMB_TOOLS_DIR:-$HOME/.cache/nmb/gate-tools}
	[[ "$NMB_TOOLS_DIR" =~ ^/[A-Za-z0-9._/+-]+$ ]] || nmb_die "NMB_TOOLS_DIR must be an absolute path of plain characters"
	case "$NMB_TOOLS_DIR/" in
	"$NMB_ROOT"/*) nmb_die "NMB_TOOLS_DIR must lie outside the repository" ;;
	esac
	export NMB_GATE_JOBS NMB_FUZZTIME NMB_TOOLS_DIR
}

# nmb_toolchain fixes the Go toolchain to the toolchain line of go.mod and refuses settings that
# switch off the checksum database.
nmb_toolchain() {
	NMB_TOOLCHAIN=$(sed -n 's/^toolchain \(go1\.[0-9]*\.[0-9]*\)$/\1/p' go.mod)
	[ -n "$NMB_TOOLCHAIN" ] || nmb_die "go.mod has no toolchain line of the form go1.N.M"
	export GOTOOLCHAIN="$NMB_TOOLCHAIN" GOWORK=off CGO_ENABLED=0
	local vals
	vals=$(go env GOVERSION GOSUMDB GONOSUMDB GOINSECURE GOPRIVATE GOFLAGS) || nmb_die "go env failed"
	local -a v
	mapfile -t v <<<"$vals"
	[ "${v[0]:-}" = "$NMB_TOOLCHAIN" ] || nmb_die "go runs ${v[0]:-?}, go.mod says $NMB_TOOLCHAIN"
	[ "${v[1]:-}" != "off" ] || nmb_die "GOSUMDB=off is refused: the checksum database makes downloads trustworthy"
	[ -z "${v[2]:-}${v[3]:-}${v[4]:-}${GONOSUMCHECK:-}" ] || nmb_die "GONOSUMDB, GOINSECURE, GOPRIVATE and GONOSUMCHECK must be empty"
	case " ${v[5]:-} " in
	*-insecure*) nmb_die "GOFLAGS with -insecure is refused" ;;
	esac
	export GOFLAGS=-mod=readonly
}

# nmb_outdir sets NMB_OUT: the folder of this gate run (from gate.sh) or a new one.
nmb_outdir() {
	if [ -n "${NMB_GATE_OUT:-}" ]; then
		case "$NMB_GATE_OUT" in
		"$NMB_ROOT"/build/gate/*) ;;
		*) nmb_die "NMB_GATE_OUT must lie under build/gate of this repository" ;;
		esac
		NMB_OUT=$NMB_GATE_OUT
	else
		NMB_OUT="$NMB_ROOT/build/gate/$(date -u +%Y%m%dT%H%M%SZ)-$(basename -- "$0" .sh)-$$"
	fi
	mkdir -p -- "$NMB_OUT"
	export NMB_GATE_OUT="$NMB_OUT"
}

# nmb_run_step NAME WHERE COMMAND...: runs the command with its output in NAME.log and records
# NAME, WHERE, the exit code and the seconds in summary.tsv. The command runs in a subshell with
# errexit on; it is not called as an if condition, because bash ignores errexit inside a function
# called there, so a step would go on after a failed command.
nmb_run_step() {
	local name=$1 where=$2 log start rc
	shift 2
	log="$NMB_OUT/${name//[^A-Za-z0-9._-]/_}.log"
	start=$SECONDS
	set +e
	(
		set -euo pipefail
		"$@"
	) >"$log" 2>&1
	rc=$?
	set -e
	printf '%s\t%s\t%s\t%s\n' "$name" "$where" "$rc" "$((SECONDS - start))" >>"$NMB_OUT/summary.tsv"
	NMB_STEPS_RUN=$((NMB_STEPS_RUN + 1))
	if [ "$rc" -eq 0 ]; then
		printf '== %-28s ok      %4ss  %s\n' "$name" "$((SECONDS - start))" "$where"
	else
		NMB_FAILED=$((NMB_FAILED + 1))
		printf '== %-28s FAILED  rc %s  %s  (log: %s)\n' "$name" "$rc" "$where" "$log"
		tail -n 40 -- "$log"
	fi
}

# nmb_finish: exit 1 if any step of this script failed, 2 if no step ran.
nmb_finish() {
	if [ "$NMB_STEPS_RUN" -eq 0 ]; then
		nmb_die "no step ran"
	fi
	if [ "$NMB_FAILED" -ne 0 ]; then
		printf 'gate: %s of %s steps failed in %s\n' "$NMB_FAILED" "$NMB_STEPS_RUN" "$(basename -- "$0")"
		exit 1
	fi
}

# nmb_files PATTERN...: NUL-separated list of existing tracked and untracked (not ignored) files.
nmb_files() {
	local f list="$NMB_OUT/.files.$$"
	local -a all
	git ls-files -z --cached --others --exclude-standard -- "$@" >"$list" || return 1
	mapfile -d '' all <"$list"
	rm -f -- "$list"
	for f in "${all[@]}"; do
		if [ -f "$f" ] || [ -L "$f" ]; then
			printf '%s\0' "$f"
		fi
	done
}

# nmb_file_array VAR PATTERN...: fills the array VAR with nmb_files; fails when git fails.
nmb_file_array() {
	local -n nmb_out_array=$1
	shift
	local tmp="$NMB_OUT/.list.$$.$RANDOM"
	if ! nmb_files "$@" >"$tmp"; then
		rm -f -- "$tmp"
		return 1
	fi
	mapfile -d '' nmb_out_array <"$tmp"
	rm -f -- "$tmp"
}

# nmb_gomods: the go.mod files of the repository, one per line (paths without spaces).
nmb_gomods() {
	git ls-files --cached --others --exclude-standard -- 'go.mod' '*/go.mod'
}

# nmb_build_helper NAME: builds tools/NAME into NMB_BIN.
nmb_build_helper() {
	go build -trimpath -o "$NMB_BIN/$1" "./tools/$1"
}

# nmb_load_tools: the gate tools from docs/DEPENDENCIES.md (the only place their versions are
# written) into NMB_TOOL_PKG, NMB_TOOL_VER and NMB_TOOL_BIN, keyed by name.
nmb_load_tools() {
	declare -gA NMB_TOOL_PKG=() NMB_TOOL_VER=() NMB_TOOL_BIN=()
	local list name pkg ver
	nmb_build_helper depcheck || nmb_die "cannot build tools/depcheck"
	list=$("$NMB_BIN/depcheck" -list docs/DEPENDENCIES.md gate-tools) || nmb_die "cannot read the gate tools"
	while IFS=$'\t' read -r name pkg ver; do
		[ -n "$name" ] || continue
		NMB_TOOL_PKG[$name]=$pkg
		NMB_TOOL_VER[$name]=$ver
		NMB_TOOL_BIN[$name]="$NMB_TOOLS_DIR/$NMB_TOOLCHAIN/$name@$ver/$name"
	done <<<"$list"
	NMB_GOSECVET=$(nmb_gosecvet_path)
}

# nmb_gosecvet_path: the driver binary, keyed by the hash of its module files.
nmb_gosecvet_path() {
	local sum
	sum=$(cat tools/gosecvet/go.mod tools/gosecvet/go.sum tools/gosecvet/main.go | sha256sum)
	printf '%s\n' "$NMB_TOOLS_DIR/$NMB_TOOLCHAIN/gosecvet@${sum:0:12}/gosecvet"
}

# nmb_tool NAME: the path of an installed gate tool; fails when it is missing.
nmb_tool() {
	local bin=${NMB_TOOL_BIN[$1]:-}
	if [ "$1" = gosecvet ]; then
		bin=$NMB_GOSECVET
	fi
	if [ -z "$bin" ] || [ ! -x "$bin" ]; then
		printf 'gate tool %s is not installed; run tools/prepare.sh\n' "$1" >&2
		return 1
	fi
	printf '%s\n' "$bin"
}

# nmb_hermetic COMMAND...: runs the command with a fresh empty HOME and TMPDIR. Go's own folders
# are kept explicitly; the go command's work files and its configuration go to folders of their
# own, so HOME and TMPDIR stay empty. Call nmb_hermetic_init first.
nmb_hermetic_init() {
	NMB_GOCACHE=$(go env GOCACHE) || return 1
	NMB_GOMODCACHE=$(go env GOMODCACHE) || return 1
	NMB_GOPATH=$(go env GOPATH) || return 1
}

nmb_hermetic() {
	local t rc=0
	t=$(mktemp -d "$NMB_OUT/hermetic.XXXXXX") || return 1
	mkdir -p -- "$t/home" "$t/tmp" "$t/gotmp" "$t/config" || return 1
	GOCACHE=$NMB_GOCACHE GOMODCACHE=$NMB_GOMODCACHE GOPATH=$NMB_GOPATH GOTMPDIR="$t/gotmp" \
		XDG_CONFIG_HOME="$t/config" HOME="$t/home" TMPDIR="$t/tmp" "$@" || rc=$?
	rm -rf -- "$t"
	return "$rc"
}

# Offline runs: a new network namespace with only loopback, inside a second user namespace mapped
# back to the real user, so tests run without network, without root and without capabilities.
nmb_ip_path() {
	local p
	for p in "$(command -v ip || true)" /usr/sbin/ip /sbin/ip; do
		if [ -n "$p" ] && [ -x "$p" ]; then
			printf '%s\n' "$p"
			return 0
		fi
	done
	return 1
}

nmb_offline() {
	local ip
	ip=$(nmb_ip_path) || return 1
	# shellcheck disable=SC2016 # the inner shell expands its own arguments
	unshare -rn sh -c 'ip=$1 u=$2 g=$3; shift 3; "$ip" link set lo up && exec unshare --user --map-user="$u" --map-group="$g" -- "$@"' \
		sh "$ip" "$(id -u)" "$(id -g)" "$@"
}

# nmb_go_offline SCRIPT ARGS...: re-runs the calling script inside the offline namespace, or says
# NOT OFFLINE when the namespace cannot be made. Sets NMB_WHERE_NET for the step labels.
nmb_go_offline() {
	if [ "${NMB_IN_OFFLINE:-}" = 1 ]; then
		NMB_WHERE_NET="offline"
		export GOPROXY=off
		return 0
	fi
	if nmb_offline true >/dev/null 2>&1; then
		local rc=0
		export NMB_IN_OFFLINE=1
		nmb_offline "$@" || rc=$?
		exit "$rc"
	fi
	NMB_WHERE_NET="NOT OFFLINE (no unprivileged network namespace)"
	printf 'gate: NOT OFFLINE: unshare -rn failed, the steps run with the network\n'
}
