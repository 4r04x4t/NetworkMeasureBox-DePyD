# Dependencies and licences

Every dependency of the native build, with its licence, origin, version and how it is used (work plan section 7
point 9, roadmap 6B.3A). The licence of the whole is not decided yet (see `README.md`); until it is, nothing is
taken whose licence would rule out an open licence.

The table below is read by `tools/depcheck` (run by `tools/deps.sh` in every gate) and must stay one table with
exactly these columns. The gate fails when:

- a `go.mod` of the repository requires a module that has no `module` row with the same path and version, or a
  `module` row is in no `go.mod` (stale);
- a licence is not one of the SPDX ids `depcheck` knows, or a part `built in` to `nmbd` has a licence other than
  MIT, BSD-2-Clause, BSD-3-Clause, Apache-2.0, ISC, 0BSD or MPL-2.0 (anything else waits for the owner, question Q4
  of brief F-00);
- the licence file (`LICENSE`, `LICENSE.txt`, `LICENSE.md`, `LICENCE` or `COPYING`) of a module, gate tool or the
  toolchain is missing or its SHA-256 differs from the row (a changed licence text is noticed);
- the `toolchain` row differs from the `toolchain` line of any `go.mod`, or from the go command that runs;
- an installed gate tool is not the listed version or was not built with the listed toolchain.

`tools/prepare.sh` installs the gate tools from their rows; there is no second copy of their versions.

Columns: **Kind** is `toolchain`, `toolchain-part` (shipped inside the toolchain), `module` (in the require list
of a `go.mod`), `gate-tool` (installed by `tools/prepare.sh`) or `system` (a program of the development PC that the
gate runs). **Use** is `built in` (compiled into `nmbd`), `test only` (only in test binaries) or `build only` (runs
on the development PC; nothing of it reaches `nmbd`). **Licence file SHA-256** is `-` only for `system` and
`toolchain-part` rows, whose licence files are not checked by the gate (the gate prints a note for each). The
licence column carries SPDX ids only; no author names are copied from licence texts. Versions of `system` rows are
as found on the development PC on the date in **Checked**; the gate does not check them.

| Name | Kind | Module or package | Version | Licence (SPDX) | Origin | Use | Licence file SHA-256 | Added by | Checked |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| Go toolchain and standard library | toolchain | `go` | `go1.27.4` | BSD-3-Clause | https://go.dev/dl/ | built in | `911f8f5782931320f5b8d1160a76365b83aea6447ee6c04fa6d5591467db9dad` | F-00 | 2026-11-11 |
| Race detector runtime (ThreadSanitizer) | toolchain-part | `runtime/race` | `go1.27.4` | Apache-2.0 WITH LLVM-exception | https://github.com/llvm/llvm-project | test only | - | F-00 | 2026-11-11 |
| gosec rules | module | `github.com/securego/gosec/v2` | `v2.30.0` | Apache-2.0 | https://github.com/securego/gosec | build only | `2be0a81f30d2b73cb3b5182b969175245042158735fcc8af133929212c772e78` | F-00 | 2026-11-11 |
| Go analysis framework | module | `golang.org/x/tools` | `v0.52.0` | BSD-3-Clause | https://go.googlesource.com/tools | build only | `911f8f5782931320f5b8d1160a76365b83aea6447ee6c04fa6d5591467db9dad` | F-00 | 2026-11-11 |
| Go module support | module | `golang.org/x/mod` | `v0.42.0` | BSD-3-Clause | https://go.googlesource.com/mod | build only | `911f8f5782931320f5b8d1160a76365b83aea6447ee6c04fa6d5591467db9dad` | F-00 | 2026-11-11 |
| Go sync extras | module | `golang.org/x/sync` | `v0.24.0` | BSD-3-Clause | https://go.googlesource.com/sync | build only | `911f8f5782931320f5b8d1160a76365b83aea6447ee6c04fa6d5591467db9dad` | F-00 | 2026-11-11 |
| Password strength estimator (used by a gosec rule) | module | `github.com/ccojocar/zxcvbn-go` | `v1.0.4` | MIT | https://github.com/ccojocar/zxcvbn-go | build only | `cf0492a938bfb94aeff5261fe47da430ae8742ae0363e11894f99f0bb6ddbdf5` | F-00 | 2026-11-11 |
| staticcheck | gate-tool | `honnef.co/go/tools/cmd/staticcheck` | `v0.8.2` | MIT | https://staticcheck.dev | build only | `f221780eca6d16dc9dce29ce46cd089ce1ca1178998b31ebbcc89ff822c6072b` | F-00 | 2026-11-11 |
| govulncheck | gate-tool | `golang.org/x/vuln/cmd/govulncheck` | `v1.9.0` | BSD-3-Clause | https://go.googlesource.com/vuln | build only | `911f8f5782931320f5b8d1160a76365b83aea6447ee6c04fa6d5591467db9dad` | F-00 | 2026-11-11 |
| gocyclo | gate-tool | `github.com/fzipp/gocyclo/cmd/gocyclo` | `v0.6.0` | BSD-3-Clause | https://github.com/fzipp/gocyclo | build only | `3fee689d85fe208d5e9a809fd8a2790e2012e054ff8961661792311b01924c91` | F-00 | 2026-11-11 |
| QEMU user-mode emulators (qemu-aarch64-static, qemu-arm-static) | system | qemu-user-static | 8.2.2 | GPL-2.0-only | https://www.qemu.org | test only | - | F-00 | 2026-11-11 |
| unshare (offline test runs) | system | util-linux | 2.39.3 | GPL-2.0-or-later | https://github.com/util-linux/util-linux | test only | - | F-00 | 2026-11-11 |
| ip (loopback inside the test namespace) | system | iproute2 | 6.1.0 | GPL-2.0-only | https://wiki.linuxfoundation.org/networking/iproute2 | test only | - | F-00 | 2026-11-11 |
| git (file lists, history scan) | system | git | 2.43.0 | GPL-2.0-only | https://git-scm.com | build only | - | F-00 | 2026-11-11 |
| C compiler (only for the race detector run) | system | gcc | 13.3.0 | GPL-3.0-or-later WITH GCC-exception-3.1 | https://gcc.gnu.org | test only | - | F-00 | 2026-11-11 |
| shellcheck (gate scripts) | system | shellcheck | 0.9.0 | GPL-3.0-only | https://www.shellcheck.net | build only | - | F-00 | 2026-11-11 |
| Time zone data (test run in another time zone) | system | tzdata | 2026e | LicenseRef-Public-Domain | https://www.iana.org/time-zones | test only | - | F-00 | 2026-11-11 |

Notes:

- The Go standard library and runtime are compiled into `nmbd`, so the toolchain row is `built in`.
- The gosec rules run through the 10-line driver `tools/gosecvet` (its own module, so its dependencies stay out of
  the product `go.mod`): the `gosec` command of v2.23.0 and later imports AI vendor SDKs whose build does not fit
  the development PC (brief F-00, section 3.3). The five `module` rows are that driver's `go.mod`; its `go.sum`
  pins them by hash. The module graph of gosec names further modules that provide no package to the driver;
  they are not built and not listed (`go mod tidy -diff` in `tools/deps.sh` keeps the require lists exact).
- staticcheck, govulncheck and gocyclo are installed with `go install <package>@<version>`, each with its own
  released dependency set; their dependencies are part of the tool and not listed separately.
- The product module `nmb` has no requirement yet: `nmbd` is the standard library only.
