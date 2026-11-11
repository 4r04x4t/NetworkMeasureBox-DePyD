# Licence of NetworkMeasureBox-DePyD: proposal

Package LIC-01, synthesis of the three candidate lists in this folder (`candidates-net.md`, `candidates-vpn.md`,
`candidates-build.md`) after their independent checks, 2026-11-11. Work plan section 7 point 9: the licence of
the whole follows from the components, and a copyleft part that is built in decides more than a part that is only
run. Roadmap 6B.3A: every component carries its licence, origin and version. This is a proposal for the owner. It
is not legal advice, and nothing here is decided until the owner answers section 4.

Terms used: **built in** = compiled into `nmbd` (or, later, into an NMB program of the C code base, 6B.8);
**run** = a separate program that NMB starts or talks to over a command line, a socket or a file; **shipped** = a
file inside an NMB image or package.

## 1. The components that decide the licence

Only built-in parts can force a licence on NMB's own code. A program that is only run never does, whatever its
licence. It brings duties for the image that ships it (notices, complete source or a written offer), and those
duties are the same under every option in section 2.

### 1.1 Built-in copyleft: the parts that decide

| Candidate | Licence | Where recorded | What it would force | Permissive way to get the same function |
| :--- | :--- | :--- | :--- | :--- |
| OONI engine `github.com/ooni/probe-engine` v0.37.0, planned as built in (roadmap 4.7 and section 7 matrix) | GPL-3.0-or-later; also GPL-3.0-or-later among its 89 built modules: `ooni/minivpn`, `yawning/utls`, obfs4's `internal/x25519ell2`; DB-IP data CC-BY-4.0 | vpn O-1 | `nmbd` as a whole under GPL-3.0-or-later (or AGPL-3.0, which can be combined with it), with complete source for every copy. For boxes handed over with the image (a GPL-3.0 "User Product"): installation information, so the user can install a modified `nmbd` even though updates are signed (owner decision 6) | `ooniprobe` as a run program (vpn O-2), or own tests written from the OONI specifications without OONI code |
| zouppp v0.3.1, souppp v1.5.0 (user-space PPPoE) | GPL-3.0-only, conservative reading | net 12, 13 | as OONI | Rejected anyway: the user-space data path capped the line (KF-0166). Own Go control plane on the kernel's PPPoE (net 15), or `pppd` run (net 11) |
| C code base only (6B.8): linking `libnftnl` or `libnftables` | GPL-2.0-or-later; `libnftables` GPL-2.0-only AND GPL-2.0-or-later | net 23, vpn V-8 | that C program under GPL-2.0 as a whole | `google/nftables` in Go (Apache-2.0), or run `nft` |
| C code base only: `readline` (if an NMB C program ever links it) | GPL-3.0-or-later | net 23 | that C program under GPL-3.0 | `libedit` (BSD-3-Clause), or no line editor |
| C code base only: `libnl-tiny` linked statically (roadmap 4.2 C plan) | LGPL-2.1 (OpenWrt metadata, not checked) | build 27 | not the licence of the program, but a static link must let the user relink (object files or source of NMB's program) | dynamic linking, or own netlink code |

### 1.2 Built-in permissive parts that still narrow the choice

| Candidate | Licence | Where recorded | Effect |
| :--- | :--- | :--- | :--- |
| `github.com/google/nftables` v0.3.0 | Apache-2.0 | net 8, vpn V-8 | Apache-2.0 cannot be combined into a GPL-2.0-only whole, so **GPL-2.0-only for `nmbd` is out**. No MIT library of the same scope exists |
| `github.com/vishvananda/netlink` v1.3.1, `tailscale.com/tsnet` (gvisor, btree and others), `wireguard-go` (gvisor), `jackpal/go-nat-pmp` | Apache-2.0 | net 3, 20; vpn V-2, V-3 | same; `vishvananda/netlink` has an MIT replacement (`jsimonetti/rtnetlink/v2`, net 22), tsnet has none |
| Every other planned Go part (x/net, x/sys, the mdlayher netlink stack, miekg/dns, insomniacslk/dhcp, wgctrl, go-qrcode, modernc or ncruces SQLite, Go fonts, the Telegram bot on `net/http`) | MIT, BSD-2/3-Clause, ISC, MIT-0, SQLite blessing | net, vpn | no effect: notices only. `depcheck` must still learn `MIT-0`, the SQLite blessing, `BSD-2-Clause-Patent` and `BSD-3-Clause-LBNL` before such rows can go into `DEPENDENCIES.md` |
| iperf3's `libiperf.a` if the C build links it (roadmap 4.8) | BSD-3-Clause-LBNL (with MIT and BSD parts) | vpn S-2 | permissive; LBNL may use published enhancements |

### 1.3 Code-origin questions that matter as much as dependencies

| Question | Where recorded | Effect |
| :--- | :--- | :--- |
| The Python Cloudflare engine calls itself a "cloudflare-speed-cli reproduction" (`nmb/speed/cloudflare.py` lines 1 and 423). That project is GPL-3.0 | vpn S-4 | If code was copied, GPL-3.0 code sits in the MIT Python repository, and a line-by-line port would carry it into the native build. Native M-06 is written from Cloudflare's MIT engine (vpn S-5) and own measurements only, whatever licence is chosen |
| `components/open-rmbt/` of the Python repository is RTR's Apache-2.0 C client without its `LICENSE`, origin or change notices | vpn S-3, KF-0283 | A fault of the Python repository. The native build does not repeat it (Quality Bar 6) |

### 1.4 Not deciding: run or shipped only

The Ookla CLI (proprietary EULA, never shipped by NMB; vpn S-1), `ooniprobe`, iperf3, OpenVPN, `wg`/`wg-quick`,
`tailscaled`, `nft`, `ip`, `ethtool`, `pppd` and its GPL PPPoE plugin, `miniupnpd` (its nftables build is a
combined GPL-2.0-or-later work with libnftnl; net 27), dnsmasq, BusyBox, the base system, the kernel and the
firmware files (build 18-46). Each image NMB ships carries their licence texts, the acknowledgement texts of
`pppd`, and the complete corresponding source of every GPL/LGPL part (or a written offer). GPL-3.0 parts of the base
(GRUB on x86-64, libgcc, `dosfstools`, parts of util-linux, gpg if shipped) bring the installation-information
point for handed-over boxes **under every option**, so signed updates need a way for the box owner to install
their own build in any case. The proprietary firmware (Raspberry Pi, Broadcom, Cypress, Realtek) is shipped
unmodified with its texts; Broadcom and Cypress add export-control clauses.

## 2. Licence options for NetworkMeasureBox-DePyD as a whole

The Python version today: its `README.md` (badge at line 5, section "License" at 485-487) and `README.hu.md` (655-656)
state the MIT License, but the repository has no `LICENSE` file and no copyright line (KF-0283). The owner wrote
that code and published it as MIT; copies already given out stay MIT whatever the native build chooses. The owner,
as author, can put the native code under any licence and can still take Python code into it.

| | A. MIT | B. Apache-2.0 | C. GPL-3.0-or-later | D. AGPL-3.0-or-later |
| :--- | :--- | :--- | :--- | :--- |
| **What others may do** | anything, including closed forks and closed firmware with `nmbd` inside | as MIT, with an explicit patent licence from every contributor (it ends for someone who sues over patents) | use, change, share; every copy of `nmbd` or a changed version comes with its complete source under the same licence | as C, and whoever runs a changed `nmbd` for users over a network (the Web UI, the bot, the API, a collector) must offer those users its source |
| **What it requires** | keep the copyright line and licence text | licence text, any `NOTICE` file, changed files marked | source with every binary (image, package, box); for handed-over boxes the installation information (a user key or a documented way past the signed-update check); patent licence | as C, plus a source link in the Web UI of every running box |
| **Built-in candidates it rules in** | every permissive part of 1.2, Apache-2.0 ones included (the binary carries both notices) | same | everything of A, **plus OONI built in** (O-1) and other GPL-3.0/LGPL code | as C |
| **Built-in candidates it rules out** | OONI built in, zouppp/souppp, any GPL/LGPL code built in; in the C build, linking libnftnl/libnftables or readline | as A; also an NMB C program can never be combined with GPL-2.0-only code (libnftables, code from netifd, odhcp6c or the kernel) | GPL-2.0-only code built in (libnftables, code taken from GPL-2.0-only OpenWrt programs or the kernel). Ookla and every run program stay possible | as C |
| **OONI** | run `ooniprobe` (O-2) or own tests from the OONI specifications | same | may be built in as the roadmap planned | same |
| **The Python version (MIT)** | one licence for both; code and specs move freely in both directions | Python code can go into the native code (MIT notice kept); native code going back into the Python repository stays Apache-2.0 unless it is the owner's own code, which the owner can relicense | Python code can go in; native code could only go back into the MIT Python repository by the owner relicensing their own code, and not at all for code from others or from OONI | as C |
| **OpenWrt package** (`PKG_LICENSE`, `PKG_LICENSE_FILES:=LICENSE`) | accepted in feeds; nothing more | accepted; `NOTICE` kept in the package | accepted; source is the public repository at the tag; an image for a router handed over adds the installation information | accepted; the Web UI source link must work on routers too |
| **Debian package** (DEP-5 `debian/copyright` listing every built-in Go module, `Static-Built-Using`) | DFSG-free, `main` | DFSG-free, `main` | DFSG-free, `main` | DFSG-free, `main` |
| **Image (6A)** | the image duties of 1.4 only | same | 1.4, plus source of `nmbd` itself | as C |

GPL-2.0-only is not an option: the Apache-2.0 parts of 1.2 cannot be combined into it, and `google/nftables` has no
replacement. The Ookla EULA is a question of its own (vpn S-1) that no option answers.

## 3. Recommendation

**Option A, MIT**, with these rules:

1. **OONI is not built in.** The OONI tests run as the separate `ooniprobe` program (vpn O-2), installed on the box
   from OONI's own or the distribution's packages, so NMB ships no GPL engine and owes no source for it. It is
   started in the strict sandbox of 6B.5, which already lists `ooniprobe` as third-party code. If OONI must run
   inside `nmbd` later, the tests are written from the OONI specifications without OONI code.
2. **The `depcheck` rule stays as it is**: only MIT, BSD-2/3-Clause, Apache-2.0, ISC, 0BSD and MPL-2.0 are built in.
   The list learns the other permissive ids (`MIT-0`, SQLite blessing, `BSD-2-Clause-Patent`, `BSD-3-Clause-LBNL`)
   when a row needs them.
3. **The C code base never links GPL libraries** (libnftnl, libnftables, readline). It links `libnl-tiny` only
   dynamically, or uses its own netlink code.
4. **No code is translated from GPL projects**, directly or through the Python code (cloudflare-speed-cli, vpn S-4).
5. **The image duties of 1.4 are met whatever the licence**: a licences folder, an about page, the `pppd`
   acknowledgements, a source archive or offer for each GPL/LGPL part, and a way for the box owner to install
   their own build next to the signed updates.

Reasons:

- **Nothing that is needed forces copyleft.** The only built-in copyleft candidate is the OONI engine, and it has a
  working run alternative. Every other planned built-in part is permissive (section 1.2).
- **The OONI engine is a poor part to build in anyway**: 89 built modules in a graph of 628, 10 known
  vulnerabilities in its dependencies with govulncheck v1.9.0, one without a fix (GO-2026-4479 in pion/dtls/v2),
  and its README says issues and pull requests there are ignored. As a separate program it can be updated,
  sandboxed and left out on small routers without touching `nmbd`.
- **One licence for the whole project.** The Python version is MIT and the owner's README promises "free for
  personal, community, and commercial use". MIT keeps that promise, and code, specifications and test vectors can
  move between the Python version, the Go code and the C code (6B.8) without a licence step.
- **The widest compatibility.** MIT code can be combined with every licence in these lists, GPL-2.0-only included.
  That matters for the C code base and for OpenWrt, whose own programs are GPL-2.0, LGPL-2.1 and ISC. Apache-2.0
  (B) cannot be combined with GPL-2.0-only.
- **The least work for users, packagers and router vendors**: a copyright line and a licence text. The signed-update
  design is not bound by NMB's own licence.

**B (Apache-2.0)** is the close second, if the owner wants an explicit patent licence from contributors and changed
files to be marked; the cost is the GPL-2.0-only incompatibility and one more licence beside the Python version's.
**C (GPL-3.0-or-later)** is the choice only if keeping every changed `nmbd` open matters more to the owner than the
points above, or if OONI must be built in; it then also requires the user-key path for signed updates for `nmbd`
itself. **D** adds a network source duty for every changed box and has no gain for NMB's use.

## 4. The decision the owner must take

**Question 1: Which licence should the native NetworkMeasureBox (NetworkMeasureBox-DePyD) have?**

- **A. MIT, the same as the Python version (recommended).** Anyone may use, change and sell it; they keep your
  copyright line and the licence text. OONI tests then run as the separate `ooniprobe` program, not built into
  `nmbd`.
- **B. Apache-2.0.** As A, with an explicit patent licence and marked changes; it cannot be mixed with GPL-2.0-only
  code.
- **C. GPL-3.0-or-later.** Anyone who passes the program on, changed or not, must give its source; boxes handed
  over must let their owner install their own build despite the signed updates. OONI could then be built in.
- **D. AGPL-3.0-or-later.** As C, and whoever runs a changed box for other users must also give them the source.

**Question 2 (with A or B): is it all right that the OONI tests run as the separate `ooniprobe` program instead of
inside `nmbd`?** Recommended: yes.

**Question 3: what name goes in the copyright line of the `LICENSE` file?** Every option needs one. Quality Bar 8
keeps personal data out of the code, so the proposal is `NetworkMeasureBox contributors`, or the GitHub account
name if you prefer. The same line also fixes the missing `LICENSE` file of the Python repository (KF-0283).

## 5. Follow-ups outside this folder (not done by LIC-01)

- `docs/DEPENDENCIES.md` of the native repository: the shellcheck row says `GPL-3.0-only`; the source headers at
  v0.11.0 and Debian say GPL-3.0-or-later (build 9).
- `docs/images/openwrt-25.12.6-x86-64.md` line 194 of the native repository repeats the `pppd` "advertising clause";
  the source has acknowledgement-retention clauses and no advertising duty (build 29).
- `tools/depcheck` learns `MIT-0`, the SQLite blessing, `BSD-2-Clause-Patent` and `BSD-3-Clause-LBNL` when needed.
- The Python repository (closed by the owner on 2026-11-10; for the owner to decide): a `LICENSE` file, the RTR
  `LICENSE` and change notices in `components/open-rmbt/` (KF-0283), and whether `cloudflare.py` contains code from
  cloudflare-speed-cli (vpn S-4).
- Ookla: someone reads the EULA in a browser and keeps a dated copy outside the repository before decision 3 is
  revisited (vpn S-1).
- After the owner's answer: `README.md`, `LICENSE` and `CHANGELOG.md` of the native repository, and the roadmap 4.7
  line that plans the OONI engine as built in.
