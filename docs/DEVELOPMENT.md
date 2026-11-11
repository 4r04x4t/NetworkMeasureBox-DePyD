# Development

How the native build is developed and checked. The plans it follows live in the Python repository
(`docs/NATIVE_WORK_PLAN.md`, `docs/DEPYTHONIZATION_ROADMAP.md`, `docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md`).

## Environment (recorded 2026-11-11, package F-00)

| What | Value |
| :--- | :--- |
| Development PC | Windows with WSL2, Ubuntu 24.04 inside (since 2026-11-10; Garuda Linux before), Linux 6.18 WSL2 kernel |
| Its limits | 4 threads, about 1.8 GB of RAM for WSL2 (3.9 GB in the PC), 1 GB swap; no root for the gate |
| Target box | Raspberry Pi 4 with Debian 13 (arm64); the box is not used by F-00 |
| Go launcher | the system `golang-go` 1.22.2 (Ubuntu package); it only starts the fixed toolchain |
| Go toolchain | `go1.27.4`, fixed by the `toolchain` line of `go.mod` (see below) |
| Gate tools | staticcheck v0.8.2, govulncheck v1.9.0, gocyclo v0.6.0, gosec v2.30.0 rules through `tools/gosecvet` |
| System programs | qemu-user-static 8.2.2, util-linux 2.39.3 (`unshare`), iproute2 6.1.0 (`ip`), git 2.43.0, gcc 13.3.0 (race detector only), shellcheck 0.9.0, tzdata 2026e |

Every dependency, with its licence and how it is used, is in [`DEPENDENCIES.md`](DEPENDENCIES.md); the gate fails
when that list and the code disagree.

## The toolchain

`go.mod` says `go 1.27.0` and `toolchain go1.27.4` (the newest patched release on go.dev's release list on
2026-11-11; go1.27.4 includes security fixes). The gate scripts read the `toolchain` line and set `GOTOOLCHAIN`
to exactly that value (not `auto`), then refuse to run when `go env GOVERSION` differs. The launcher downloads the
toolchain once through the Go proxy, checked against the checksum database; with `GOTOOLCHAIN=local` the
system Go 1.22 refuses the module ("go.mod requires go >= 1.27.0"), so the old Go cannot build it by mistake.

The module path is `nmb` (no account or host name, Quality Bar point 8). Such a path cannot be fetched with
`go get` from the network; nothing outside imports this module (everything is `cmd/` or `internal/`).

The analysers are built with the fixed toolchain into
`${NMB_TOOLS_DIR:-~/.cache/nmb/gate-tools}/<toolchain>/<tool>@<version>/<tool>`, outside every repository and never
committed; their versions are read from `DEPENDENCIES.md` and nowhere else. Binaries in `~/go/bin` are not used:
they were built with Go 1.22 and cannot read code for a newer Go.

## Running the gate

```sh
tools/gate.sh            # the whole gate; exit 0 passed, 1 a step failed, 2 usage or environment error
```

It runs, always all of them and in this order:

| Script | Does | Network |
| :--- | :--- | :--- |
| `tools/prepare.sh` | toolchain, module downloads and `go mod verify`, gate tools, the gosec driver | proxy and checksum database |
| `tools/deps.sh` | `go mod tidy -diff` of every `go.mod`; `tools/depcheck` against `DEPENDENCIES.md` | none (`GOPROXY=off`) |
| `tools/build-matrix.sh` | `CGO_ENABLED=0` builds for linux/amd64, arm64, arm v7, mips and mipsle (softfloat) into `dist/`; checks the build settings in each binary; records the sizes | none |
| `tools/analyse.sh` | gofmt, go vet (amd64, arm64, arm, mips), staticcheck (amd64 and 32-bit arm), gosec, govulncheck (source and the built binaries), gocyclo over 15, `tools/cleancode`, shellcheck, the script rules, the list of `#nosec` annotations, the same checks for `tools/gosecvet` | govulncheck only (vuln.go.dev) |
| `tools/test.sh` | runs A to G: amd64 with the race detector; amd64 without cgo, shuffled, in `Pacific/Chatham`; arm64 and arm v7 test binaries under `qemu-aarch64-static` and `qemu-arm-static`, in UTC and shuffled in `Pacific/Chatham`; mips and mipsle test binaries compiled only | none (namespace) |
| `tools/fuzz.sh` | every `Fuzz*` target for `NMB_FUZZTIME` | none (namespace) |
| `tools/scan.sh tree dist` | `tools/privscan` over every tracked and untracked (not ignored) file and over `dist/` | none |
| `tools/gate-selftest.sh` | plants one fault per check in a throw-away copy and requires the check to fail, and the clean copy to pass | none |

Each script can run alone and takes step names, for example `tools/analyse.sh gofmt cleancode`,
`tools/test.sh B`, `tools/build-matrix.sh linux-arm64`, `tools/gate-selftest.sh root`. Results go to
`build/gate/<UTC time>-.../` (git-ignored): one log per step, `summary.tsv` (step, where it ran, exit code,
seconds), `sizes.tsv`, `tools.txt` (`go version -m` of every analyser) and `environment.tsv`.

Knobs (each one checked; a bad value stops the gate with exit 2):

| Variable | Default | Meaning |
| :--- | :--- | :--- |
| `NMB_GATE_JOBS` | 2 | parallel Go builds and fuzz workers, 1 to 4 (memory: about 1.8 GB in WSL2) |
| `NMB_FUZZTIME` | 30s | time per fuzz target, for example `30s` or `2m` |
| `NMB_TOOLS_DIR` | `~/.cache/nmb/gate-tools` | where the gate tools are installed; absolute, outside the repository |
| `NMB_PRIVATE_VALUES` | not set | a file of the owner's own values (names, addresses, keys) that must appear nowhere; outside the repository, never committed; its values are never printed |

The gate refuses to run as root (fault KF-0001), and refuses `GOSUMDB=off`, `GONOSUMDB`, `GONOSUMCHECK`,
`GOINSECURE`, `GOPRIVATE` and `GOFLAGS` with `-insecure`.

### What the step labels mean

- `offline`: the step ran in a new network namespace with only loopback, as the real user without capabilities
  (`unshare -rn`, then a second user namespace mapped back to the real user). `NOT OFFLINE` means the namespace
  could not be made and the step ran with the network; the gate says so and never hides it.
- `QEMU user mode`: the test binary ran under `qemu-aarch64-static` or `qemu-arm-static` on the PC, not on ARM
  hardware.
- `compile only`: the test binaries were built, not run (mips and mipsle).

Every test run has a fresh empty `HOME` and `TMPDIR`, no test cache (`-count=1`), and fails when the repository
changed during it.

### What the gate does not do

- It cannot shift the clock: Go binaries are static, so `LD_PRELOAD` tools do not apply; tests use the fake clock
  of F-06 instead.
- Unprivileged ICMP is off inside the namespace; tests that need ping get fakes from F-06.
- It does not run QEMU system guests (no package of F-00 needs one); later packages take the lock
  `~/nmb-work/qemu.lock` with `flock` and give a guest at most 512 MB.
- The history scan (`tools/scan.sh history`: every blob, commit message and author or committer address of all
  refs) runs before a push, not in every package gate. Commit addresses must be GitHub noreply addresses or be
  accepted in the private list with a line `identity-ok <address>`.

## Clean-code limits

`tools/cleancode` fails a function with more than 50 lines in its body, a file with more than 500 lines (Go and
shell) and nesting deeper than 3 levels (the function body is level 0; `if`, `for`, `switch`, `select`, a bare
block and a function literal open a level; `else if` and `case` do not). A longer function or file needs a
stated reason in a directive (`//nmb:allow-long <reason>`, `//nmb:allow-long-file <reason>`, or
`# nmb:allow-long-file <reason>` in a script's first five lines), and the gate lists every such exception. Nesting
has no exception. `gocyclo` fails a cyclomatic complexity above 15.

## Personal data

`tools/privscan` finds e-mail addresses outside the reserved domains, public IPv4 and IPv6 addresses outside the
documentation and private ranges (the shared range 100.64.0.0/10 counts as a hit), MAC addresses that are
universally administered, key- and token-like strings, developer home folders, and the values of the private
list. Tests and examples use reserved values only: 192.0.2.0/24, 198.51.100.0/24, 203.0.113.0/24, 2001:db8::/32,
example.com, and MAC addresses starting with `02:00:00`. A value that must stay (for example a public resolver)
goes into `tools/privscan/allow.txt` with its reason, reviewed by the package's verifier.
