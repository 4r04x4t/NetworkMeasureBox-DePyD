# Requirements register: package changes

Proposed additions, splits, merges and moves of the work packages, gathered from the parts and de-duplicated, and the
three mappings the work plan asks for (`docs/NATIVE_WORK_PLAN.md@ad3758c:217-218`: every Python file, every roadmap
section and every `TESTING.md` row in exactly one package). Nothing here is decided: the owner (F-01) or the analyst of
the package named accepts or rejects each proposal. Ids are permanent; a proposed id that is rejected is not reused.

Sections A to G are written by hand; sections M1 to M3 are generated from the parts and checked by `check.sh`.

## A. Numbers for the open series

Work plan section 3.5 leaves `S-02…S-0n` (API by area) and `U-02…U-0n` (pages) without numbers. The requirement rows
name the series with an area or page label (`S-02…S-0n (API: status)`, `U-02…U-0n (page: speed)`); only
`parts/python-files.md` uses numbers. The numbers below are proposed; until they are accepted the rows keep their labels.

**API areas.** `parts/python-files.md` (section A) and `parts/changelog.md` (section 3) propose the same eight numbers:

| Id | Area | Label in the requirement rows |
| :--- | :--- | :--- |
| S-02 | Status | `API: status` |
| S-03 | Speed and measuring | `API: speed` |
| S-04 | Network | `API: network` |
| S-05 | Wi-Fi | `API: Wi-Fi` |
| S-06 | VPN and remote access | `API: VPN` |
| S-07 | DNS blocker | `API: blocker` |
| S-08 | System and settings (also `/api/config`) | `API: system` |
| S-09 | Exports and public results | `API: exports` |

S-0n stops at S-09 because S-10 is `nmb-cli`.

**Pages.** Two proposals differ: `parts/changelog.md` proposes nine coarse pages (U-02 home ... U-10 results, logout and
public pages); `parts/python-files.md` proposes one package per tab of `nmb/web/assets/dashboard.html@ad3758c:87-108`
(U-02 to U-21) plus three size splits (U-22 to U-24, section B). The register uses the second, because the file table
already cuts `dashboard.html` and `dashboard.js` along it and it keeps each page within the size rule. The changelog
numbering is not used anywhere and is withdrawn by this merge.

| Label in the requirement rows | Proposed page package |
| :--- | :--- |
| home, Home, Dashboard | U-02 Dashboard page |
| Ping | U-03 Ping telemetry page |
| Speed, speed | U-04 Speed results page; the changelog part's "speed" also means *Measurement & Servers*, U-14: the analyst checks each row |
| Schedule | U-14 Measurement and servers page (the schedule sub-tab moves there, python-files C.11/A.2) |
| iPerf3 | U-05 iPerf3 page |
| Traceroute; Port check | U-06 Traceroute page; U-11 Port neutrality page |
| Blocker | U-07 DNS blocker page |
| VPN, Remote | U-09 Remote access and VPN page (DynDNS, proxy and Tailscale parts: U-22 if accepted) |
| Devices; network | U-10 Network explorer page; the changelog part's "network" means network and router: U-16 |
| Router, router; Ports & Drivers, ports and drivers | U-16 Router page |
| Wi-Fi | U-23 Wi-Fi page if accepted, otherwise U-16 |
| System, system; Settings, settings | U-17 System page (backup and stored data: U-24 if accepted) |
| Tasks | U-21 Tasks and load page |
| AI diagnostics | U-19 AI diagnostics page |
| Developer console | U-12 Logs page (the tab is "Developer & System Logs") |
| tools | several: U-10, U-13, U-06, U-21, U-20 (the changelog part's tools page) |
| pages: Ping, Speed (, Outages); pages: Traceroute, Port check, RTR | one row for several pages: each page package takes it, the first named owns it |
| SLA | U-16 Router page: the ISP profile and the live SLA audit card are `nmb/web/assets/dashboard.html@ad3758c:3704-3780`, a range of U-16 in `parts/python-files.md` (round 2) |
| Exports | U-04 Speed results page, which holds the speed, broadband, outage and event CSV buttons (`dashboard.html@ad3758c:237-240`); the ping CSV button is on U-03 (`:193`) and takes the row there too; the exports themselves are R-02 and the API S-09 (round 2) |
| Results, results and logout | U-01: the logout page (`nmb/web/assets/logout.html`, `dashboard.js@ad3758c:8300-8331`) is U-01 in `parts/python-files.md`; the public `/results` page is served by S-09 (`nmb/web/results_page.py`) and its look and texts are U-01's (round 2) |
| Support | U-01: the Support button, donation dialog and footer are `dashboard.html@ad3758c:1-109, 5221-5239`, U-01 in `parts/python-files.md` (round 2) |
| Outages | U-03 Ping telemetry page: the 24-hour latency and outage telemetry is in the ping tab (`dashboard.html@ad3758c:186-226`); the outages CSV button is on U-04 (round 2) |

**Bare `S-02` and `U-02`.** Some rows (mostly in the roadmap parts and REQ-WP-0134, which copies work plan line 215) write
`S-02` or `U-02` where the series is meant. They are counted under the numbers S-02 (owns 3, Also in 18) and U-02 (owns 2,
Also in 50) in `BY_PACKAGE.md`. The S-02 and U-02 analysts must hand rows that belong to another area or page to it.

## B. New packages (de-duplicated)

Where two parts proposed the same id for different packages, the merge kept the id that rows already use and gave the
other proposal the next free number (section F). Letter suffixes (`M-12a`, `S-11b` ...) used in some parts are not ids:
a split that is accepted gets the next free number of its series.

| Id | Name | Proposed in | Owns / Also in | Reason |
| :--- | :--- | :--- | :--- | :--- |
| X2-01 | Step-2 router tiers (holding package) | roadmap-1-507 (1), porting-plan (1), workplan (4), testing (3), readme | 141 / 44 | Work plan section 1 puts the C11 tier and all router work in step 2, and section 3 has no package for it; without an owner these rows would fail the check. Every row that concerns only the C code base or only router targets is owned here, with the step-1 package under *Also in* (round 1, section G). Stays one holding package in step 1 (its rows go into `spec/` through F-07); split when step 2 is planned. |
| T-01 | First router backends, TL-WR1043ND v1 | roadmap-508-1034 (2) | 3 / 29 | Step 2 writes code (ath79 switch, ath9k, procd/fw4, flash layout, C11/musl build, Go-to-C port, optional UPX) that no package owns; I-10 only analyses images. **Child of X2-01** (porting-plan 1); later routers get T-02 ... in the owner's order. |
| P-14 | Maker's firmware round trip | roadmap-508-1034 (1) | 9 / 10 | Back to stock firmware (6B.3): per-target write method, optional fetch from the maker, recovery path, round-trip test on each router. Step 2 and per-device; keeping it in P-02 would block P-02's step-1 acceptance. Generic safety rules stay in P-02. |
| P-15 | Endurance and hostile-network campaign | testing (1) | 3 / 0 | The 24-72 h run, the 24 h low-memory run and the hostile-network suite need days of box or VM time; inside P-13 they would make the release proof too big. P-15 runs before P-13 and hands it its results. |
| P-16 | SSH access (keys only) | changelog (1), there as "P-14" | 2 / 0 | No package owns the box's SSH server setting (keys-only switch, check that sshd takes it, automatic switch back). Alternative: merge into P-05 (maintenance mode also gives keys-only access). |
| P-17 | Guided live scenarios | python-files (D, there as "P-14"), testing (2, there as "P-07b") | 0 / 0 | `nmb/selftest/` is 2,837 lines; P-07 keeps the live self-test groups (`live.py`, testing's "P-07a"), P-17 takes `scenario*.py` with their restore safety. Rows stay in P-07 until accepted. |
| S-14 | Bot commands: measuring, reports, mode, VPN | python-files (D) | 0 / 0 | `nmb/bot/` is about 4,100 lines; S-14 = `commands.py@ad3758c:249-1061`. |
| S-15 | Bot commands: DynDNS, blocker, healing, tasks, Wi-Fi, backup, drivers, reboot | python-files (D) | 0 / 0 | S-15 = `commands.py@ad3758c:1062-1725`. **Other proposal** (workplan 2): S-11a commands and menus, S-11b alert monitor (REQ-AR-0072 to -0078). The two cuts are compatible: if the alert monitor is split too, it gets the next free number; the S-11 analyst decides. |
| N-19 | Blocker lists: gravity update, allow/deny/regex lists, adlists, local DNS records | python-files (D) | 0 / 1 | N-12 is about 2,600 lines with its CLI; N-19 = `netd/adblock.py@ad3758c:1003-1515`. |
| N-20 | Box resolver and DNS fallback (optional) | changelog (2), there as "N-19" | 0 / 1 | Rows of changelog lines 210, 224, 254 are the box's own resolver management, not a measurement; if M-12 is too big, `netd/dns.py` is cut into M-12 (audit, lookups) and N-20 (fallback, restore). |
| M-15 | Scheduler: cron and interval parser, next run, retry timing | python-files (D) | 0 / 0 | M-04 is about 1,930 lines; M-15 = `speed/scheduler.py` (also used by the web server). |
| M-16 | RMBT protocol client | python-files (D) | 0 / 0 | M-08 is about 1,450 lines of Python and 2,900 of C; M-16 = `speed/rmbt/` and `components/open-rmbt/`. |
| M-17 | Shared HTTP transfer core (optional) | roadmap-1-507 (4), there as "M-15"; readme | 0 / 0 | Multi-stream worker pool and per-socket interface binding written and tested once for M-06, M-07, M-10 and M-04 (REQ-RM-0026, -0027, and REQ-RE-0193 per the readme part). If not taken, the rows stay with M-06 and M-04. |
| C-12 | Translation data conversion | python-files (D) | 0 / 0 | About 27,000 lines of strings; C-06 keeps the key system and the build checks, C-12 moves every string into the F-04 key files. |
| U-22 | DynDNS, proxy and Tailscale page | python-files (D) | 0 / 0 | Split from U-09 (HTML 827 + JS 1,573 lines). |
| U-23 | Wi-Fi page: access point and Wi-Fi client | python-files (D) | 0 / 0 | Split from U-16. |
| U-24 | Backup, restore and stored data page | python-files (D) | 0 / 0 | Split from U-17. |

**Children of X2-01** proposed by the porting-plan part (not ids yet; they get numbers when step 2 is planned): one
package per router target (T-01 first); Asuswrt/Merlin integration (REQ-PP-0042 to -0054); C tier storage and budgets
(REQ-PP-0055 to -0066, -0175 to -0184); missing parts from the maker's firmware, 7H (REQ-PP-0106 to -0115, next to I-10);
remote access, recovery and brick risk per target (REQ-PP-0144, -0150, -0162, -0171); MikroTik and MediaTek families
(REQ-PP-0021 to -0026). The roadmap part 1-507 proposes the split per component instead (probe, netd, web, RMBT, iPerf3,
VPN, driver finder, updater, storage guard, USB tethering on OpenWrt, RouterOS container). Both are step-2 decisions.

## C. Splits and size notes without a new id yet

| Package | Proposal | From | Note |
| :--- | :--- | :--- | :--- |
| M-12 | Split into DNS resolver matrix and hijack check (4.20, REQ-RM-0116 to -0119) and STUN, NAT type, WAN topology and public address (4.21, REQ-RM-0120 to -0123) | roadmap-1-507 (2), optional | **Conflict:** python-files C.2-C.4 moves more code *into* M-12 (public addresses from `interfaces.py`, NAT/STUN/DNS interception from `llm_bundle.py`), and changelog proposes N-20 out of it. Taken together M-12 needs a split; its analyst decides the cut. |
| N-10 | Split into detection, uplink, table 152, data limit, fallback (REQ-RM-0165 to -0171, -0176, -0177) and router mode on USB with LAN-port sharing (REQ-RM-0172, -0173) | roadmap-1-507 (3), optional | Only if `netd/usb_uplink.py` plus tests is above the size rule. |
| F-00 | Split into repository skeleton, toolchain and licence list, and gate scripts and work status (REQ-WP-0017 to -0028, -0034 to -0036, -0043 to -0053, -0172; REQ-DV-0003 to -0005, -0022, -0023) | workplan (3), optional | F-00 owns 125 rows. |
| P-04, P-02 | Check the size rule; e.g. P-04 service profiles and P-04 third-party/upload sandbox | roadmap-508-1034 (3) | P-02 owns 64 rows, P-04 29. |
| R-04, C-01 | Shrink by moving line ranges (CPU and thermal to C-09, PHY to N-17, Wi-Fi PHY to M-14; boot-partition settings copy to P-03) | python-files (D) | No new id. |
| P-13 | Owns 113 rows | (merge) | With P-15 accepted it loses the long runs; the release checklist rows stay. |
| S-11 | Owns 124 rows | (merge) | See S-14, S-15. |

## D. Moves and one owner per file

1. **One owner per Python file.** The work plan names six files for two packages (`probe/engine.py`, `report/metrics.py`,
   `netd/modes.py`, `netd/wifi.py`, `netd/wifi_share.py`, `web/server.py`), and four further files or folders through its
   P-10 row (`bin/`, `systemd/`, `os/`): `systemd/` units (P-04, P-10), `bin/nmb-helper` (P-06, P-10),
   `bin/nmb-factory-reset` (P-03, P-10) and `os/nmb-guard.nft` (N-02, P-10). M3 gives every `systemd/` unit to P-10 (P-04
   takes the sandbox lines), `bin/nmb-helper` to P-06, `bin/nmb-factory-reset` to P-03 and `os/nmb-guard.nft` to N-02.
   `parts/python-files.md` `parts/python-files.md` cuts 61 files over
   several packages by line range. Read literally, work plan line 217 ("every Python file ... in exactly one package") fails
   for each of them. The register reads it as **every line of every Python file in exactly one package**, and gives each
   file one package in M3 (the package with most of its lines); the other packages take their ranges as written in
   `parts/python-files.md`. The workplan part (1) proposed the same alternative; its other option (one owner per file,
   the rest as "uses") is what M3 records. F-01 confirms the reading.
2. **Moves out of the package the plan names** (python-files C, each with its reason there): `netd/upnp_wan.py` N-11 to M-03;
   `netd/geoip.py@ad3758c:154-214` to N-07; `netd/interfaces.py@ad3758c:666-813` to M-12 and `847-1070` to M-01;
   `report/llm_bundle.py@ad3758c:412-620` to M-12; `netd/services.py` pppd parts to N-05 and miniupnpd/prefix delegation to N-11;
   `common/heal.py@ad3758c:1004-1310` to P-02; `setup.sh@ad3758c:393-454` to P-03; translation tables in code to C-06;
   the `bin/` sub-CLIs to their subsystem packages; the `os/` files to N-01, P-11, P-06, P-03, M-13 and S-13; themes to P-09.
3. **`netd/wifi_share.py`** (testing 4): DHCP guard (`132-251`) N-05, the rest N-08; M3 gives the file to N-08.
4. **Box identity** (readme): REQ-RE-0169 is owned by C-01 (made once at first start, stored with the settings); S-12 is the
   alternative.
5. **Self-heal path** (roadmap-1-507 note): the roadmap names `nmb/report/self_heal.py`, the work plan `common/heal.py`;
   the file of the commit is `nmb/common/heal.py` (P-01, updates part P-02).

## E. Porting-plan sections (part author's map, for information)

The work plan's rule names roadmap sections, not porting-plan sections; the porting-plan part maps them anyway:
1, 1.1-1.3, 2, 4, 5, 6, 7G, 7H, 8 (C parts), 9, 10 to X2-01; 3 to P-10; 7.1 X2-01 (N-15 Also in); 7.2 S-11 (C form X2-01);
7.3 R-03; 7.4 M-14; 7.5 C-09; 7.6 M-12; 7.7 X2-01 (N-12 owns REQ-PP-0079); 7.8 M-04; 7.9 P-01; 7.10 at line 505 N-16; 7.10 at line 698 S-12 (the plan has two
sections numbered 7.10: cite lines, not the number); 7A N-06; 7B M-03; 7F F-04; 7E P-13; 7D P-05; 7C P-11; 8 (Go part) F-00.

## F. Mechanical changes made by the merge

- Id collisions of proposals renumbered: changelog's "NEW: P-14 SSH access" is **P-16** (2 rows and section 3 of that part);
  changelog's optional "N-19 Box resolver and DNS fallback" is **N-20**; roadmap-1-507's and readme's optional "M-15 Shared
  HTTP transfer core" is **M-17**; python-files' "new P-14 Guided live scenarios" is **P-17**. P-14, N-19 and M-15 keep the
  meaning of the part whose rows or file table use them (roadmap-508-1034, python-files, python-files).
- Spelling of the open series made one: `U-02..U-0n` and `S-02..S-0n (area: x)` became `U-02…U-0n` and `S-02…S-0n (API: x)`
  (readme, roadmap parts).
- `S-12/S-13` in Also-in cells became `S-12, S-13` (roadmap-508-1034, 2 rows); the Also-in cell "see REQ-WP-0134" of
  REQ-WP-0072 moved to its Notes.
- No requirement text, owner or source cell was changed.

## G. Round 1 of the checks (2026-11-11)

The checkers' findings of round 1 changed these owners and rows; no id was renumbered, new ids took the next free number.

- **Step-2 rows to X2-01.** REQ-RM-1080 (C hardening flags), -1095 (router hardening limits), -1201 (JSON commands on small
  flash), -1319 (built-in look only on small flash) and REQ-PP-0041 (RAM ring buffer), -0067 (Tier 1 `tsnet`), -0077 (Tier 1
  radix trie) are owned by X2-01; the step-1 package moved to *Also in*. Mixed rows were split into a Go or generic row that
  stays and a C or router row owned by X2-01: REQ-RM-1451 (from -1074), -1452 (from -1077), -1453 (from -1254), -1454 (from
  -1376), -1455 (from -1409). REQ-RM-1349 (procd/ujail profile) stays with P-04: OpenWrt x86-64 and `armsr` are step-1 QEMU
  targets (`docs/NATIVE_WORK_PLAN.md@ad3758c:266`, `:285-287`).
- **Other moves:** REQ-RM-1208 P-04 to P-02; -1016 N-17 to F-00 (every package Also in); -1027 I-01 to I-09; -1182 M-01 to
  N-03; -1009 F-02 to F-00; REQ-WP-0032 P-13 to F-00; REQ-WP-0154 N-17 to F-04; REQ-TS-0028 P-01 to P-06; REQ-TS-0248 N-01
  to M-12 (with REQ-CL-0299); REQ-RE-0098 P-02 to M-04; REQ-RE-0191 P-13 to P-11 (with REQ-CL-0104); REQ-RE-0288 P-04 to P-06
  (the privilege rule has one owner, P-06: REQ-RE-0339, REQ-TS-0014, -0018, REQ-CL-0365); REQ-CL-0092 the exports API to
  S-01 (with REQ-TS-0125, REQ-RE-0303); REQ-CL-0341 P-02 to P-11 (with REQ-TS-0168, REQ-CT-0018).
  REQ-RM-1237 (Phase 2 changelog) stays with N-01, with the reason in its Notes.
- **Splits:** REQ-RE-0287 into -0287 (web server user, P-04) and -0339 (privilege rule, P-06); REQ-CL-0235 into -0235
  (address flushes, N-07) and -0452 (cable preferred as measuring link, N-01).
- **Withdrawn:** REQ-PP-0080 (Ookla protocol emulation), superseded by REQ-WP-0179 and -0180; it keeps M-10 as package,
  and the gate expects no test for it.
- **New rows:** REQ-RM-0213 (setup words from one source, C-11), REQ-RE-0340 (SMTP dispatch, R-05), and the new part
  `parts/contributing.md` (REQ-CT-0001 to -0026), which no part covered before.
- `BY_PACKAGE.md` was regenerated from the parts.

## H. Round 2 of the checks (2026-11-11)

The checkers' findings of round 2 changed these rows; no id was renumbered, new ids took the next free number.

- **Split into a step-1 row and a router or C row owned by X2-01** (step-1 package and T-01 under *Also in*): REQ-RM-0214
  (from -0038, router UI budget and flash fit), -0215 (from -0049), -0216 (from -0050), -0217 (from -0052), -0218 (from
  -0106, C11 bot memory), -1456 (from -1413), -1457 (from -1414), -1458 (from -1416), -1459 (from -1418), -1460 (from -1422).
- **Split per protocol owner:** REQ-RM-1066 keeps DHCP and PPPoE (N-05); DNS went to -1461 (N-12) and -1462 (M-12), UPnP
  to -1463 (N-11). REQ-WP-0148 keeps `cpio`, `xorriso`/`bsdtar` and `ostree` (now F-00); the router-firmware readers went
  to REQ-WP-0197 (I-10).
- **Moves:** REQ-RM-0068 M-13 to X2-01; REQ-RM-1073 S-01 to F-00 (a rule for every change, like -1060 and -1061);
  REQ-PP-0184 P-13 to X2-01 (a router-tier design target like -0175 to -0183); REQ-WP-0016, -0141, -0142, -0149, -0150,
  -0151 I-01 to F-00 (shared image-method rules; each I package proves them for its own image); REQ-RE-0209 C-08 to P-13
  (one owner with REQ-TS-0139).
- **Restated:** REQ-RM-0028 (native Ookla engine only under owner decision 3), -0145 (status file only), -1301 (the
  condition of the fourth candidate); REQ-PP-0029 (scope of REQ-RM-0006, with the sandboxed third-party programs and vendor
  tools as exceptions), -0099 (the source's "should"), -0185 (source narrowed to line 584); REQ-AR-0021 ("only" removed);
  REQ-CL-0232, -0249, -0319, -0343, -0386; REQ-WP-0028, -0057 and the notes of REQ-RM-0011, REQ-DV-0009, REQ-RE-0176 (the
  development PC's figures and Go version no longer copied, Quality Bar 8).
- **New rows:** REQ-RM-0219 (C-tier binary size), REQ-RM-1464 (pass criteria of the firewall namespace test), REQ-WP-0196
  (each role a separate agent), REQ-CL-0453 (iPerf3 port from a server's port range).
- **Process rules:** rules no test of the owner's code can prove carry `process rule:` in Notes and are listed in
  `BY_PACKAGE.md` *Flags* (register README, *Owner and Also in*).
- **Mappings:** section A maps the open page labels SLA, Exports, Results, Support and Outages to page packages; section E
  and the porting-plan part map 7.1 and 7.7 to X2-01; M1 *Rows by owner* (stale since round 1 for several sections) and M2
  lines 42 and 327 were recomputed from the parts, and `check.sh` now recomputes both.
- `BY_PACKAGE.md` was regenerated from the parts.

## M1. Every roadmap section in one package

Every `##` and `###` heading of `docs/DEPYTHONIZATION_ROADMAP.md@ad3758c` outside code blocks (69), with the lines up to the
next such heading. *Package* is the one package the section belongs to: the package that owns most of the section's rows,
unless *Basis* says otherwise. *Rows* counts the owners of the REQ-RM rows that cite a line of the section (a row that cites
lines of two sections counts in both; round 2: recomputed from the parts, and `check.sh` now recomputes it).
`check.sh` fails when a heading is missing here, listed twice, or has no valid package.

| Line | Lines | Section | Package | Basis | Rows by owner |
| ---: | :--- | :--- | :--- | :--- | :--- |
| 2 | 2-5 | Migration Roadmap for Embedded Routers, OpenWrt & Resource-Constrained Hardware | F-02 | context: document subtitle, held by the register | none |
| 6 | 6-7 | 1. Executive Summary & Problem Formulation | P-13 | context: chapter heading; 1.1 and 1.2 below are mapped | none |
| 8 | 8-18 | 1.1 The Embedded Reality | P-13 | majority | P-13 3, P-10 2, M-01 1 |
| 19 | 19-25 | 1.2 Strategic Goal: The De-Pythonization Objective | F-00 | tie: language and toolchain goal (REQ-RM-0010 to -0014 are F-00) | F-00 1, P-10 1, P-13 1, X2-01 1 |
| 26 | 26-52 | 2. Technology Evaluation: Go (Golang) vs. C (C11) vs. C++ (C++17) | F-00 | context: language evaluation, decided as the toolchain of F-00 | X2-01 1 |
| 53 | 53-68 | 2.1 Detailed Comparison Matrix | F-00 | tie (F-00 2, P-13 2, X2-01 2): comparison of the languages, toolchain of F-00 | F-00 2, P-13 2, X2-01 2, C-02 1, M-01 1, N-01 1, S-01 1, U-01 1 |
| 69 | 69-70 | 3. Recommended Architectural Decision | F-00 | context: chapter heading of the language decision | none |
| 71 | 71-80 | 🎯 Two-Tier Tiered Strategy: | F-00 | tie: one row each; the tier split is the toolchain decision (C tier rows in X2-01) | F-00 1, P-11 1, U-01 1, X2-01 1 |
| 81 | 81-115 | 4. Component-by-Component Migration Blueprint | P-10 | majority | P-10 1 |
| 116 | 116-127 | 4.1 Latency, Jitter & Loss Probe Engine (`nmb-probe`) | M-01 | majority | M-01 3, X2-01 1 |
| 128 | 128-137 | 4.2 Dynamic Network & Route Daemon (`nmb-netd`) | N-01 | majority | N-01 2, N-02 1, X2-01 1 |
| 138 | 138-145 | 4.3 High-Throughput Speedtest Engine (`nmb-speed`) | M-04 | majority | M-04 3, M-06 2, M-08 1, M-10 1 |
| 146 | 146-156 | 4.4 Web Server, REST API & Embedded Dashboard (`nmb-web`) | S-01 | plan | S-01 3, U-01 1, X2-01 1 |
| 157 | 157-175 | 4.4A Optimised Web Interface for the Native Build (Owner, 2026-11-10; Planned) | U-01 | majority | U-01 15, S-01 4, X2-01 4, S-02 2, C-06 1 |
| 176 | 176-182 | 4.5 In-Memory Database & Persistence (`nmb-db`) | C-02 | majority | C-02 2 |
| 183 | 183-195 | 4.6 Native RTR-NetzTest (Open-RMBT) Engine (`rmbt-client`) | M-08 | majority | M-08 4, X2-01 1 |
| 196 | 196-207 | 4.7 Censorship & Neutrality Engine: In-Process OONI Probe (`ooniprobe`) | M-13 | majority | M-13 3, X2-01 1 |
| 208 | 208-216 | 4.8 High-Precision Local/WAN Bandwidth Engine: iPerf3 Protocol Client | M-09 | majority | M-09 2, X2-01 1 |
| 217 | 217-227 | 4.9 Embedded DNS Ad-Blocker & Privacy Shield (`nmb-adblock` / Pi-hole Alternative) | N-12 | majority | N-12 5 |
| 228 | 228-243 | 4.10 Dual-Hub VPN (WireGuard + OpenVPN EC) & Tailscale Mesh (`tsnet`) | N-13 | plan names N-13, N-14 and N-15 for 4.10; majority N-13 | N-13 2, N-14 1, N-15 1, N-18 1, X2-01 1 |
| 244 | 244-252 | 4.11 Hardware Driver Finder & Firmware Provisioning Engine (`nmb-driver`) | N-17 | majority | N-17 3, X2-01 1 |
| 253 | 253-261 | 4.12 Autonomous Self-Healing & Safety Rollback Watchdog (`nmb-heal`) | P-01 | plan (tie P-01 2, P-02 2) | P-01 2, P-02 2, X2-01 1 |
| 262 | 262-274 | 4.13 Real-Time Task Manager & Hardware Monitor (`/proc` Telemetry Engine) | C-09 | majority | C-09 4 |
| 275 | 275-282 | 4.14 Dedicated Guest Wi-Fi & Station Multi-SSID Wireless Controller | N-07 | majority | N-07 3 |
| 283 | 283-290 | 4.15 Invasive Diagnostic Logger & Crash Trap (`nmb-invasive-log`) | R-05 | majority; plan names C-03 and R-05 for 4.15 | R-05 4 |
| 291 | 291-303 | 4.16 Multilingual Telegram Chatbot & Interactive Menus in Native Code | S-11 | majority | S-11 4, X2-01 1 |
| 304 | 304-312 | 4.17 Regulatory-Grade Broadband SLA Dossier & Pure SVG Vector Engine | R-03 | majority; plan names R-01 and R-03 for 4.17 | R-03 4 |
| 313 | 313-322 | 4.18 Micro-Outage Detection & Loss Pattern Correlation Engine | M-02 | majority | M-02 3 |
| 323 | 323-334 | 4.19 Wi-Fi RF & PHY Telemetry via `nl80211` Netlink | M-14 | majority | M-14 2 |
| 335 | 335-348 | 4.20 DNS Hijacking & Comparative Multi-Resolver Matrix Engine | M-12 | majority | M-12 4 |
| 349 | 349-360 | 4.21 RFC 5389 STUN Client & Carrier-Grade NAT (CGNAT) Engine | M-12 | majority | M-12 4 |
| 361 | 361-368 | 4.22 CPU Interrupts (IRQ/SoftIRQ), RPS/XPS & NIC Offload Telemetry | C-09 | majority | C-09 3 |
| 369 | 369-376 | 4.23 Storage I/O & Flash Filesystem Health Monitor | C-08 | majority | C-08 3 |
| 377 | 377-389 | 4.24 Sanitized Configuration & PII/Secret Masking Engine | C-04 | majority; plan names C-04 and R-04 for 4.24 | C-04 6 |
| 390 | 390-395 | 4.25 Central Collector Client & GitHub Updater | S-13 | majority; plan names M-04, S-12 and S-13 for 4.25 | S-13 6, S-12 3, M-04 1, M-05 1, P-02 1, X2-01 1 |
| 396 | 396-409 | 4.26 Features Added in November 2026 (the Python Reference the Native Build Must Match) | N-10 | majority; the section collects several features (C-06, C-08, M-07, N-08 also cite it) | N-10 9, M-07 5, N-08 5, C-08 2, C-06 1, M-04 1, M-09 1, N-05 1, P-13 1, X2-01 1 |
| 410 | 410-435 | 4.27 USB Tethering & Mobile Uplink (Phones, 5G/4G Hotspots, Cellular Modems) | N-10 | majority | N-10 12, X2-01 3, N-05 2, N-08 1, P-11 1 |
| 436 | 436-445 | 4.28 Automatic Mode (`mode = auto`, the Owner's Rule) | N-04 | majority | N-04 7 |
| 446 | 446-457 | 4.29 The Measurement Setup on Every Result (Owner, 2026-11-10) | C-11 | majority | C-11 10, C-02 1, R-01 1, R-02 1, R-03 1, R-04 1, R-05 1, S-10 1, S-11 1, S-12 1, U-02 1 |
| 458 | 458-493 | 5. Architectural Proposal: The Unified Single-Binary Monolith (`nmbd`) | P-10 | plan (section 5) | C-02 1, P-10 1, S-01 1 |
| 494 | 494-507 | Multi-Call CLI Compatibility (Busybox-style): | P-10 | majority | P-10 2 |
| 508 | 508-535 | 6. Four-Phase Hybrid Migration Roadmap | P-13 | tie (P-13 2, P-11 2): generic first, proven on the PC and the Pi 4 | P-11 2, P-13 2, P-01 1 |
| 536 | 536-541 | Where the Native Files Live (Owner, 2026-11-10) | F-00 | majority | F-00 6 |
| 542 | 542-557 | Generic First: Platform- and Architecture-Independent Code | F-00 | majority | F-00 6, P-13 3, F-04 2 |
| 558 | 558-571 | Order of Work (Owner, 2026-11-10) | I-09 | part: step-1 image analysis first (tie F-03 3, I-09 3; I-10 2) | F-03 3, I-09 3, I-10 2, F-00 1, F-06 1, P-08 1, P-13 1, P-14 1, T-01 1 |
| 572 | 572-605 | Quality Bar: At Least as Good as the Python Version, and Better | P-13 | plan: P-13 names QB 1, 4, 5, 7 (majority F-00 14, P-13 10) | F-00 14, P-13 10, C-10 8, P-03 6, F-03 4, P-02 4, X2-01 4, C-01 3, C-06 3, F-06 3, S-01 3, R-03 2, R-04 2, S-11 2, C-03 1, C-04 1, C-05 1, F-05 1, M-12 1, N-02 1, N-05 1, N-07 1, N-11 1, N-12 1, P-04 1, P-06 1, P-11 1, S-10 1, S-13 1, U-01 1 |
| 606 | 606-620 | Register of Python Faults the Native Build Must Not Repeat (owner, 2026-11-10) | F-03 | plan: F-03 is the roadmap's register (majority N-02 7) | N-02 7, M-04 6, S-11 5, F-03 4, R-04 4, N-12 3, N-14 3, C-11 2, M-02 2, P-06 2, R-02 2, C-01 1, C-04 1, C-07 1, F-06 1, M-01 1, M-10 1, M-12 1, P-02 1, P-03 1, R-01 1, S-02 1, U-02 1 |
| 621 | 621-758 | Phase 0: Freeze Behaviour Before Rewriting | F-05 | plan: F-05 is the Phase 0 reference (majority F-00 18) | F-00 18, M-03 11, S-10 11, P-13 7, C-01 6, M-05 5, N-03 3, S-01 3, F-05 2, M-04 2, N-01 2, N-05 2, P-02 2, P-03 2, P-11 2, X2-01 2, C-02 1, C-08 1, F-04 1, F-06 1, M-06 1, M-09 1, M-10 1, N-02 1, N-14 1, N-17 1, P-04 1, P-06 1, P-10 1 |
| 759 | 759-763 | Phase 1: High-ROI Micro-Engine Replacements | M-01 | tie: first step of the phase (probe) | C-09 1, M-01 1, M-04 1 |
| 764 | 764-768 | Phase 2: Dynamic Routing, DNS Ad-Blocker & Wi-Fi IPC | N-01 | tie: first step of the phase (routing daemon) | N-01 1, N-07 1, N-12 1 |
| 769 | 769-773 | Phase 3: WireGuard VPN, In-Process OONI & Unified Monolith (`nmbd`) | N-13 | tie: first step of the phase (WireGuard) | M-13 1, N-13 1, P-10 1 |
| 774 | 774-780 | Phase 4: Self-Healing Watchdog, Driver Autodiscovery & OpenWrt Packaging | P-01 | tie: first step of the phase (watchdog) | N-17 1, P-01 1, P-11 1 |
| 781 | 781-840 | 6A. NMB's Own Linux Distribution: Python-Free (Planned) | P-11 | plan: P-11 names 6A (majority I-09 15, P-11 12) | I-09 15, P-11 12, P-02 3, C-08 2, C-02 1, I-01 1, I-06 1, I-07 1, I-08 1, N-07 1, P-03 1, P-13 1 |
| 841 | 841-844 | 6B. Product Decisions for the Native Build (Owner, 2026-11-10; Planned) | F-00 | tie: scope rule of the 6B decisions (REQ-RM-1314) | F-00 1, P-08 1, U-01 1 |
| 845 | 845-851 | 6B.1 Look: AMOLED black by default, other themes only as add-ons | P-09 | majority | P-09 3, U-01 1, X2-01 1 |
| 852 | 852-857 | 6B.2 CLI: no colour by default, colour as an add-on | S-10 | majority | S-10 3, P-09 1 |
| 858 | 858-865 | 6B.3 Firmware and software from a file, and back to the maker's firmware | P-02 | plan (tie P-02 8, P-14 8) | P-02 8, P-14 8 |
| 866 | 866-872 | 6B.3A Licence information and disclaimer for every component (owner, 2026-11-10) | P-08 | majority | P-08 6 |
| 873 | 873-883 | 6B.5 Sandboxing as a built-in option (owner, 2026-11-10) | P-04 | majority | P-04 11, F-06 1, N-02 1, N-03 1, P-02 1 |
| 884 | 884-899 | 6B.6 Captive portals on a Wi-Fi uplink (owner, 2026-11-10) | N-09 | majority | N-09 18, X2-01 1 |
| 900 | 900-916 | 6B.7 Full user manual and command-line reference (owner, 2026-11-10) | P-12 | majority | P-12 18 |
| 917 | 917-930 | 6B.8 One behaviour, two languages: the Go and the C code bases learn from each other (owner, 2026-11-10) | F-07 | majority | F-07 11, F-00 1, F-05 1, T-01 1, X2-01 1 |
| 931 | 931-942 | 6B.4 Downloaded images and firmware to work with | I-09 | part (tie) | F-00 1, I-09 1, I-10 1 |
| 943 | 943-975 | 7. Component Language & Dependency Elimination Matrix | P-10 | part (step-1 rows spread, M-12 2, C-09 2; the 5 X2-01 rows are the C halves split from step-1 rows in round 2) | X2-01 5, C-09 2, M-12 2, C-02 1, C-04 1, C-08 1, M-01 1, M-02 1, M-04 1, M-08 1, M-09 1, M-13 1, M-14 1, N-01 1, N-07 1, N-12 1, N-13 1, N-15 1, N-17 1, P-01 1, P-10 1, R-03 1, R-05 1, S-01 1, S-11 1 |
| 976 | 976-999 | 8. Cross-Compilation Matrix (Go) | F-00 | majority | F-00 2 |
| 1000 | 1000-1008 | UPX Binary Compression (Optional for SPI Flash): | T-01 | majority | T-01 1 |
| 1009 | 1009-1019 | 9. Conclusion & Performance Projections | P-13 | majority | P-13 2, M-04 1, P-10 1, P-11 1 |
| 1020 | 1020-1031 | 10. Router Hardware Porting Specifications | I-10 | context: step-2 router hardware list | none |
| 1032 | 1032-1034 | Required for every port: development access and safe fixes | P-13 | part (tie, one row each) | C-03 1, P-02 1, P-05 1, P-13 1 |

Sections whose rows are owned by more than one package: 41 of 69. They keep one package above; the
other packages own single rows of the section (see *Rows by owner*), which the register allows because every row has one owner.

## M2. Every TESTING.md row in one package

Every data row of every table of `docs/TESTING.md@ad3758c` and every item of *Known limits* (212 rows; header and
separator rows left out). Only line numbers are given: some rows name the test site (Quality Bar 8). *Package* is the owner
of most of the row's REQ-TS rows; on a tie, the owner of the row's first REQ-TS id. `check.sh` fails when a row is missing,
listed twice, or has no valid package, and (round 2) when the owners column or the package differs from what the parts give.

| Line | Package | REQ-TS rows | Owners of those rows |
| ---: | :--- | :--- | :--- |
| 17 | P-13 | TS-0007 | P-13 1 |
| 18 | F-06 | TS-0008 | F-06 1 |
| 19 | P-11 | TS-0009 | P-11 1 |
| 20 | F-06 | TS-0010 | F-06 1 |
| 32 | P-11 | TS-0012 | P-11 1 |
| 33 | P-11 | TS-0013, TS-0014 | P-06 1, P-11 1 |
| 34 | N-01 | TS-0015 | N-01 1 |
| 35 | P-11 | TS-0016 | P-11 1 |
| 36 | P-11 | TS-0017, TS-0018 | P-06 1, P-11 1 |
| 37 | P-11 | TS-0019 | P-11 1 |
| 38 | P-11 | TS-0020 | P-11 1 |
| 39 | P-04 | TS-0021, TS-0022, TS-0023 | M-08 1, P-04 1, P-10 1 |
| 40 | P-11 | TS-0024 | P-11 1 |
| 41 | P-11 | TS-0025, TS-0026 | P-04 1, P-11 1 |
| 42 | P-11 | TS-0027, TS-0028 | P-06 1, P-11 1 |
| 43 | P-11 | TS-0029 | P-11 1 |
| 49 | M-10 | TS-0030, TS-0031 | M-05 1, M-10 1 |
| 50 | M-08 | TS-0032 | M-08 1 |
| 51 | M-06 | TS-0033 | M-06 1 |
| 52 | N-05 | TS-0034, TS-0035, TS-0036, TS-0037 | N-05 2, N-08 1, P-10 1 |
| 53 | M-04 | TS-0038 | M-04 1 |
| 54 | M-07 | TS-0039 | M-07 1 |
| 55 | M-09 | TS-0040 | M-09 1 |
| 56 | M-01 | TS-0041 | M-01 1 |
| 57 | M-04 | TS-0042, TS-0043 | M-04 1, M-05 1 |
| 58 | M-11 | TS-0044 | M-11 1 |
| 59 | M-03 | TS-0045 | M-03 1 |
| 60 | M-12 | TS-0046 | M-12 1 |
| 66 | N-03 | TS-0047 | N-03 1 |
| 67 | N-03 | TS-0048 | N-03 1 |
| 68 | N-02 | TS-0049 | N-02 1 |
| 69 | N-05 | TS-0050 | N-05 1 |
| 70 | N-03 | TS-0051 | N-03 1 |
| 71 | N-10 | TS-0052 | N-10 1 |
| 79 | N-05 | TS-0054 | N-05 1 |
| 80 | N-05 | TS-0055 | N-05 1 |
| 81 | N-04 | TS-0056 | N-04 1 |
| 82 | N-02 | TS-0057, TS-0058 | N-02 1, N-03 1 |
| 83 | M-04 | TS-0059 | M-04 1 |
| 84 | N-05 | TS-0060 | N-05 1 |
| 85 | N-03 | TS-0061 | N-03 1 |
| 86 | N-02 | TS-0062 | N-02 1 |
| 87 | N-05 | TS-0063 | N-05 1 |
| 88 | N-11 | TS-0064 | N-11 1 |
| 89 | N-11 | TS-0065 | N-11 1 |
| 90 | N-11 | TS-0066 | N-11 1 |
| 91 | N-16 | TS-0067 | N-16 1 |
| 92 | N-07 | TS-0068 | N-07 1 |
| 93 | M-04 | TS-0069 | M-04 1 |
| 94 | N-05 | TS-0070 | N-05 1 |
| 100 | M-05 | TS-0071 | M-05 1 |
| 101 | M-05 | TS-0072 | M-05 1 |
| 102 | M-04 | TS-0073 | M-04 1 |
| 103 | M-01 | TS-0074, TS-0075 | M-01 1, M-02 1 |
| 104 | N-03 | TS-0076 | N-03 1 |
| 105 | N-07 | TS-0077 | N-07 1 |
| 106 | N-07 | TS-0078 | N-07 1 |
| 107 | N-04 | TS-0079 | N-04 1 |
| 108 | N-03 | TS-0080 | N-03 1 |
| 109 | N-11 | TS-0081 | N-11 1 |
| 110 | S-11 | TS-0082, TS-0083, TS-0084, TS-0085, TS-0086, TS-0087 | C-02 1, C-08 1, N-07 1, N-16 1, P-02 1, S-11 1 |
| 111 | C-03 | TS-0088 | C-03 1 |
| 112 | C-09 | TS-0089 | C-09 1 |
| 113 | C-08 | TS-0090 | C-08 1 |
| 114 | C-08 | TS-0091 | C-08 1 |
| 115 | P-10 | TS-0092 | P-10 1 |
| 116 | P-11 | TS-0093 | P-11 1 |
| 117 | N-04 | TS-0094, TS-0095, TS-0096, TS-0097 | M-01 1, N-01 1, N-04 1, N-08 1 |
| 118 | P-02 | TS-0098, TS-0099, TS-0100 | P-02 2, M-05 1 |
| 119 | N-05 | TS-0101 | N-05 1 |
| 120 | C-02 | TS-0102 | C-02 1 |
| 126 | N-07 | TS-0103 | N-07 1 |
| 127 | N-07 | TS-0104 | N-07 1 |
| 128 | N-07 | TS-0105 | N-07 1 |
| 129 | N-07 | TS-0106 | N-07 1 |
| 130 | N-07 | TS-0107 | N-07 1 |
| 131 | N-07 | TS-0108 | N-07 1 |
| 132 | N-07 | TS-0109 | N-07 1 |
| 138 | N-08 | TS-0110 | N-08 1 |
| 139 | N-08 | TS-0111 | N-08 1 |
| 140 | N-08 | TS-0112 | N-08 1 |
| 141 | P-06 | TS-0113 | P-06 1 |
| 142 | N-08 | TS-0114 | N-08 1 |
| 148 | N-13 | TS-0115 | N-13 1 |
| 149 | N-14 | TS-0116 | N-14 1 |
| 150 | N-11 | TS-0117 | N-11 1 |
| 151 | N-16 | TS-0118 | N-16 1 |
| 152 | N-15 | TS-0119 | N-15 1 |
| 158 | N-12 | TS-0120 | N-12 1 |
| 159 | P-02 | TS-0121, TS-0122, TS-0123 | F-06 1, P-02 1, S-10 1 |
| 160 | U-02…U-0n | TS-0124 | U-02…U-0n 1 |
| 161 | S-01 | TS-0125 | S-01 1 |
| 162 | M-05 | TS-0126 | M-05 1 |
| 163 | P-03 | TS-0127 | P-03 1 |
| 164 | S-10 | TS-0128 | S-10 1 |
| 165 | N-02 | TS-0129, TS-0130, TS-0131 | C-04 1, N-02 1, P-04 1 |
| 166 | N-01 | TS-0132 | N-01 1 |
| 167 | S-11 | TS-0133, TS-0134, TS-0135 | S-11 3 |
| 168 | R-05 | TS-0136, TS-0137, TS-0138 | C-08 1, M-13 1, R-05 1 |
| 169 | P-13 | TS-0139 | P-13 1 |
| 170 | P-03 | TS-0140, TS-0141, TS-0142 | P-03 3 |
| 171 | C-02 | TS-0143 | C-02 1 |
| 172 | M-10 | TS-0144, TS-0145, TS-0146 | M-05 1, M-10 1, S-02…S-0n 1 |
| 173 | M-04 | TS-0147 | M-04 1 |
| 174 | P-02 | TS-0148 | P-02 1 |
| 175 | P-11 | TS-0149 | P-11 1 |
| 176 | S-12 | TS-0150, TS-0151 | S-12 1, S-13 1 |
| 177 | C-08 | TS-0152 | C-08 1 |
| 178 | C-01 | TS-0153, TS-0154 | C-01 2 |
| 179 | N-10 | TS-0155, TS-0156, TS-0157 | M-04 1, M-08 1, N-10 1 |
| 180 | N-10 | TS-0158, TS-0159, TS-0160 | C-07 1, N-01 1, N-10 1 |
| 181 | N-05 | TS-0161, TS-0162, TS-0163, TS-0164, TS-0165 | N-05 2, N-03 1, N-07 1, N-10 1 |
| 182 | N-07 | TS-0166 | N-07 1 |
| 183 | S-10 | TS-0167 | S-10 1 |
| 184 | P-11 | TS-0168 | P-11 1 |
| 185 | C-04 | TS-0169 | C-04 1 |
| 218 | P-07 | TS-0180 | P-07 1 |
| 219 | P-07 | TS-0181 | P-07 1 |
| 220 | P-07 | TS-0182 | P-07 1 |
| 221 | P-07 | TS-0183 | P-07 1 |
| 222 | P-07 | TS-0184 | P-07 1 |
| 223 | P-07 | TS-0185 | P-07 1 |
| 224 | P-07 | TS-0186 | P-07 1 |
| 225 | P-07 | TS-0187 | P-07 1 |
| 226 | P-07 | TS-0188 | P-07 1 |
| 227 | P-07 | TS-0189 | P-07 1 |
| 228 | P-07 | TS-0190 | P-07 1 |
| 229 | P-07 | TS-0191 | P-07 1 |
| 284 | P-13 | TS-0216 | P-13 1 |
| 285 | P-13 | TS-0217 | P-13 1 |
| 286 | P-13 | TS-0218 | P-13 1 |
| 287 | P-13 | TS-0219 | P-13 1 |
| 288 | P-13 | TS-0220 | P-13 1 |
| 289 | P-13 | TS-0221 | P-13 1 |
| 290 | P-13 | TS-0222 | P-13 1 |
| 291 | P-13 | TS-0223 | P-13 1 |
| 292 | P-13 | TS-0224 | P-13 1 |
| 293 | P-13 | TS-0225 | P-13 1 |
| 294 | P-13 | TS-0226 | P-13 1 |
| 295 | P-13 | TS-0227 | P-13 1 |
| 296 | P-13 | TS-0228 | P-13 1 |
| 297 | P-13 | TS-0229 | P-13 1 |
| 298 | P-13 | TS-0230 | P-13 1 |
| 299 | P-13 | TS-0231 | P-13 1 |
| 300 | P-13 | TS-0232 | P-13 1 |
| 301 | P-13 | TS-0233 | P-13 1 |
| 302 | P-13 | TS-0234 | P-13 1 |
| 303 | P-13 | TS-0235 | P-13 1 |
| 304 | P-13 | TS-0236 | P-13 1 |
| 305 | P-13 | TS-0237 | P-13 1 |
| 306 | P-13 | TS-0238 | P-13 1 |
| 307 | P-13 | TS-0239 | P-13 1 |
| 315 | N-08 | TS-0240 | N-08 1 |
| 316 | N-08 | TS-0241 | N-08 1 |
| 317 | N-08 | TS-0242 | N-08 1 |
| 318 | N-07 | TS-0109, TS-0243 | N-07 2 |
| 319 | N-07 | TS-0244 | N-07 1 |
| 320 | P-13 | TS-0245 | P-13 1 |
| 321 | X2-01 | TS-0246 | X2-01 1 |
| 327 | S-11 | TS-0247, TS-0248, TS-0249, TS-0250, TS-0251 | M-08 1, M-12 1, N-01 1, N-07 1, S-11 1 |
| 328 | N-04 | TS-0252, TS-0253 | N-04 1, N-06 1 |
| 329 | P-07 | TS-0254, TS-0255 | N-16 1, P-07 1 |
| 330 | N-11 | TS-0256 | N-11 1 |
| 331 | N-05 | TS-0257, TS-0258, TS-0259 | M-04 1, N-05 1, N-11 1 |
| 332 | P-03 | TS-0260 | P-03 1 |
| 333 | N-03 | TS-0261, TS-0262 | N-03 1, N-15 1 |
| 334 | N-08 | TS-0263 | N-08 1 |
| 335 | N-03 | TS-0164, TS-0264 | N-03 2 |
| 336 | N-01 | TS-0160, TS-0265 | N-01 2 |
| 337 | N-01 | TS-0266 | N-01 1 |
| 338 | N-10 | TS-0267 | N-10 1 |
| 339 | C-04 | TS-0268 | C-04 1 |
| 340 | M-09 | TS-0269 | M-09 1 |
| 341 | M-05 | TS-0031, TS-0270 | M-05 2 |
| 342 | S-13 | TS-0151, TS-0271 | S-13 2 |
| 348 | C-08 | TS-0152, TS-0272 | C-08 2 |
| 349 | P-11 | TS-0273 | P-11 1 |
| 350 | C-01 | TS-0153, TS-0274 | C-01 2 |
| 351 | M-07 | TS-0275 | M-07 1 |
| 352 | P-15 | TS-0276 | P-15 1 |
| 353 | P-15 | TS-0277 | P-15 1 |
| 354 | C-02 | TS-0278, TS-0279 | C-02 1, P-03 1 |
| 355 | N-06 | TS-0280 | N-06 1 |
| 356 | P-15 | TS-0281 | P-15 1 |
| 357 | S-11 | TS-0282 | S-11 1 |
| 358 | S-11 | TS-0283 | S-11 1 |
| 359 | C-01 | TS-0284 | C-01 1 |
| 360 | R-05 | TS-0136, TS-0285 | R-05 2 |
| 361 | P-02 | TS-0148, TS-0286 | P-02 2 |
| 362 | U-01 | TS-0287 | U-01 1 |
| 363 | P-11 | TS-0013, TS-0025, TS-0118, TS-0119, TS-0137, TS-0288, TS-0289, TS-0290, TS-0291, TS-0292 | P-11 3, M-13 2, N-15 2, N-16 2, N-02 1 |
| 371 | P-11 | TS-0029, TS-0294 | P-11 2 |
| 372 | P-11 | TS-0029, TS-0295 | P-11 2 |
| 373 | P-11 | TS-0029, TS-0296 | P-11 2 |
| 374 | P-11 | TS-0029, TS-0297 | P-11 2 |
| 375 | P-11 | TS-0029, TS-0298 | P-11 2 |
| 376 | P-11 | TS-0029, TS-0299 | P-11 2 |
| 377 | X2-01 | TS-0300 | X2-01 1 |
| 378 | P-11 | TS-0301 | P-11 1 |
| 386 | M-05 | TS-0302 | M-05 1 |
| 387 | C-03 | TS-0303 | C-03 1 |
| 388 | M-05 | TS-0304 | M-05 1 |
| 389 | N-01 | TS-0305 | N-01 1 |
| 390 | N-05 | TS-0306 | N-05 1 |
| 394 | N-07 | TS-0307 | N-07 1 |
| 395 | C-11 | TS-0308 | C-11 1 |
| 396 | N-05 | TS-0309 | N-05 1 |
| 397 | M-05 | TS-0031, TS-0310 | M-05 2 |
| 398 | C-02 | TS-0311 | C-02 1 |
| 399 | F-06 | TS-0312 | F-06 1 |
| 400 | S-11 | TS-0283, TS-0313 | S-11 2 |
| 401 | N-07 | TS-0314 | N-07 1 |

Rows whose requirement rows have more than one owner: 28 of 212: lines 33, 36, 39, 41, 42, 49, 52, 57, 82, 103, 110, 117, 118, 159, 165, 168, 172, 176, 179, 180, 181, 327, 328, 329, 331, 333, 354, 363.
Each of them is one TESTING row that checks several things; the package above holds the row (the test record entry), the
others prove their own REQ-TS row. Splitting the row in the native `TESTING.md` is the other way out; P-13 decides.

## M3. Every Python file in one package

Every file of `parts/python-files.md` (351 files of `git ls-tree ad3758c`, without `docs/` and the top-level `*.md`).
*Package* is the package that ports most of the file's lines; *Also cut to* lists the other packages that `python-files.md`
gives line ranges of the same file (rows by line count). `check.sh` fails when a file is missing here or listed twice.

| File | Package | Also cut to |
| :--- | :--- | :--- |
| `.claude/settings.json` | F-00 | - |
| `.github/pull_request_template.md` | F-00 | - |
| `.gitignore` | F-00 | - |
| `bin/nmb-adblock` | N-12 | - |
| `bin/nmb-cli` | P-10 | - |
| `bin/nmb-collector` | P-10 | - |
| `bin/nmb-driver` | N-17 | - |
| `bin/nmb-factory-reset` | P-03 | - |
| `bin/nmb-heal` | P-01 | - |
| `bin/nmb-helper` | P-06 | - |
| `bin/nmb-invasive-log` | R-05 | - |
| `bin/nmb-iperf` | M-09 | - |
| `bin/nmb-netd` | P-10 | - |
| `bin/nmb-ooni` | M-13 | - |
| `bin/nmb-probe` | P-10 | - |
| `bin/nmb-report` | P-10 | - |
| `bin/nmb-selftest` | P-10 | - |
| `bin/nmb-speed` | P-10 | - |
| `bin/nmb-telegram` | P-10 | - |
| `bin/nmb-vpn` | N-13 | - |
| `bin/nmb-web` | P-10 | - |
| `bin/nmb-wifi` | N-08 | - |
| `components/open-rmbt/Makefile` | M-08 | - |
| `components/open-rmbt/README.md` | M-08 | - |
| `components/open-rmbt/src/connection.c` | M-08 | - |
| `components/open-rmbt/src/connection.h` | M-08 | - |
| `components/open-rmbt/src/control.c` | M-08 | - |
| `components/open-rmbt/src/control.h` | M-08 | - |
| `components/open-rmbt/src/gui.c` | M-08 | - |
| `components/open-rmbt/src/gui.h` | M-08 | - |
| `components/open-rmbt/src/main.c` | M-08 | - |
| `components/open-rmbt/src/pretest.c` | M-08 | - |
| `components/open-rmbt/src/pretest.h` | M-08 | - |
| `components/open-rmbt/src/pthread_barrier_compat.h` | M-08 | - |
| `components/open-rmbt/src/tests.c` | M-08 | - |
| `components/open-rmbt/src/tests.h` | M-08 | - |
| `components/open-rmbt/src/uuid_store.c` | M-08 | - |
| `components/open-rmbt/src/uuid_store.h` | M-08 | - |
| `config/nmb.conf.example` | C-01 | - |
| `config/usb-import.example.conf` | P-03 | - |
| `nmb/__init__.py` | P-10 | - |
| `nmb/bot/__init__.py` | S-11 | - |
| `nmb/bot/__main__.py` | S-11 | - |
| `nmb/bot/alerts.py` | S-11 | - |
| `nmb/bot/client.py` | S-11 | - |
| `nmb/bot/commands.py` | S-11 | - |
| `nmb/bot/daemon.py` | S-11 | - |
| `nmb/bot/i18n.py` | C-06 | S-11 (488) |
| `nmb/bot/locales/es.json` | C-06 | - |
| `nmb/bot/locales/it.json` | C-06 | - |
| `nmb/bot/locales/pt.json` | C-06 | - |
| `nmb/bot/locales/pt_br.json` | C-06 | - |
| `nmb/bot/locales/ro.json` | C-06 | - |
| `nmb/bot/locales/zh.json` | C-06 | - |
| `nmb/bot/messages.py` | C-06 | - |
| `nmb/collector/__init__.py` | S-12 | - |
| `nmb/collector/client.py` | S-12 | - |
| `nmb/collector/server.py` | S-13 | - |
| `nmb/common/__init__.py` | C-05 | - |
| `nmb/common/backup.py` | P-03 | - |
| `nmb/common/donate.py` | P-08 | - |
| `nmb/common/heal.py` | P-01 | P-02 (307) |
| `nmb/common/memory.py` | C-08 | - |
| `nmb/common/net_setup.py` | C-11 | - |
| `nmb/common/netaddr.py` | C-05 | - |
| `nmb/common/own_action.py` | C-07 | - |
| `nmb/common/purge.py` | C-08 | - |
| `nmb/common/redact.py` | C-04 | - |
| `nmb/common/secret_keep.py` | C-04 | - |
| `nmb/common/setup_texts.py` | C-11 | - |
| `nmb/common/storage_guard.py` | C-08 | - |
| `nmb/common/tasks.py` | C-09 | - |
| `nmb/common/updater.py` | P-02 | - |
| `nmb/common/usb_import.py` | P-03 | - |
| `nmb/common/user_agent.py` | C-05 | - |
| `nmb/common/version.py` | C-05 | - |
| `nmb/config.py` | C-01 | - |
| `nmb/db/__init__.py` | C-02 | - |
| `nmb/db/corrections.py` | C-02 | - |
| `nmb/db/database.py` | C-02 | - |
| `nmb/db/schema.sql` | C-02 | - |
| `nmb/db/setup_backfill.py` | C-11 | - |
| `nmb/devlog.py` | C-03 | - |
| `nmb/netd/__init__.py` | N-03 | - |
| `nmb/netd/__main__.py` | N-03 | - |
| `nmb/netd/adblock.py` | N-12 | - |
| `nmb/netd/auto_mode.py` | N-04 | - |
| `nmb/netd/daemon.py` | N-03 | - |
| `nmb/netd/detector.py` | N-01 | - |
| `nmb/netd/dns.py` | M-12 | - |
| `nmb/netd/drivers.py` | N-17 | - |
| `nmb/netd/duckdns.py` | N-16 | - |
| `nmb/netd/explorer.py` | N-17 | - |
| `nmb/netd/firewall.py` | N-02 | - |
| `nmb/netd/geoip.py` | M-12 | N-07 (61) |
| `nmb/netd/guest_wifi.py` | N-08 | - |
| `nmb/netd/interfaces.py` | N-01 | M-01 (224), M-12 (148) |
| `nmb/netd/ipv6_lan.py` | N-11 | - |
| `nmb/netd/kernel.py` | N-01 | - |
| `nmb/netd/lan_topology.py` | N-06 | - |
| `nmb/netd/modes.py` | N-03 | N-11 (111), N-10 (60), N-05 (42) |
| `nmb/netd/openvpn.py` | N-14 | - |
| `nmb/netd/os_detector.py` | N-17 | - |
| `nmb/netd/port_forward.py` | N-11 | - |
| `nmb/netd/proxy.py` | N-11 | - |
| `nmb/netd/regdomains.py` | N-07 | - |
| `nmb/netd/services.py` | N-03 | N-05 (103), N-11 (71) |
| `nmb/netd/tailscale.py` | N-15 | - |
| `nmb/netd/upnp_wan.py` | M-03 | - |
| `nmb/netd/usb_uplink.py` | N-10 | - |
| `nmb/netd/wifi.py` | N-07 | M-14 (57) |
| `nmb/netd/wifi_share.py` | N-08 | N-05 (120) |
| `nmb/netd/wifi_uplink.py` | N-08 | - |
| `nmb/netd/wireguard.py` | N-13 | - |
| `nmb/probe/__init__.py` | M-01 | - |
| `nmb/probe/__main__.py` | M-01 | - |
| `nmb/probe/cross_traffic.py` | M-03 | - |
| `nmb/probe/engine.py` | M-01 | M-02 (87) |
| `nmb/probe/icmp.py` | M-01 | - |
| `nmb/probe/ooni.py` | M-13 | - |
| `nmb/probe/traffic_gate.py` | M-03 | - |
| `nmb/report/__init__.py` | R-01 | - |
| `nmb/report/__main__.py` | R-03 | - |
| `nmb/report/cli.py` | R-03 | - |
| `nmb/report/crash_trap.py` | C-03 | - |
| `nmb/report/csv_export.py` | R-02 | - |
| `nmb/report/dossier.py` | R-03 | C-06 (658) |
| `nmb/report/invasive_logger.py` | R-05 | - |
| `nmb/report/llm_bundle.py` | R-04 | M-12 (209) |
| `nmb/report/locales/es.json` | C-06 | - |
| `nmb/report/locales/it.json` | C-06 | - |
| `nmb/report/locales/pt.json` | C-06 | - |
| `nmb/report/locales/pt_br.json` | C-06 | - |
| `nmb/report/locales/ro.json` | C-06 | - |
| `nmb/report/locales/zh.json` | C-06 | - |
| `nmb/report/metrics.py` | R-01 | M-02 (132) |
| `nmb/report/smtp.py` | R-05 | - |
| `nmb/report/status_codes.py` | R-01 | - |
| `nmb/report/svg_charts.py` | R-03 | C-06 (129) |
| `nmb/selftest/__init__.py` | P-07 | - |
| `nmb/selftest/live.py` | P-07 | - |
| `nmb/selftest/scenario.py` | P-07 | - |
| `nmb/selftest/scenario_auto.py` | P-07 | - |
| `nmb/selftest/scenario_pppoe.py` | P-07 | - |
| `nmb/selftest/scenario_router.py` | P-07 | - |
| `nmb/sniffer/__init__.py` | M-11 | - |
| `nmb/sniffer/engine.py` | M-11 | - |
| `nmb/speed/__init__.py` | M-04 | - |
| `nmb/speed/__main__.py` | M-04 | - |
| `nmb/speed/anti_ban.py` | M-05 | - |
| `nmb/speed/cancel.py` | M-04 | - |
| `nmb/speed/cloudflare.py` | M-06 | - |
| `nmb/speed/data_budget.py` | M-05 | - |
| `nmb/speed/data_cap.py` | M-05 | - |
| `nmb/speed/fastcom.py` | M-07 | - |
| `nmb/speed/iperf.py` | M-09 | - |
| `nmb/speed/iperf_directory.py` | M-09 | - |
| `nmb/speed/ookla.py` | M-10 | - |
| `nmb/speed/port_checker.py` | M-11 | - |
| `nmb/speed/rmbt/__init__.py` | M-08 | - |
| `nmb/speed/rmbt/__main__.py` | M-08 | - |
| `nmb/speed/rmbt/connection.py` | M-08 | - |
| `nmb/speed/rmbt/control.py` | M-08 | - |
| `nmb/speed/rmbt/pretest.py` | M-08 | - |
| `nmb/speed/rmbt/tests.py` | M-08 | - |
| `nmb/speed/rmbt/uuid_store.py` | M-08 | - |
| `nmb/speed/rtr.py` | M-08 | - |
| `nmb/speed/runner.py` | M-04 | - |
| `nmb/speed/scheduler.py` | M-04 | - |
| `nmb/speed/traceroute.py` | M-11 | - |
| `nmb/tui/__init__.py` | S-10 | - |
| `nmb/tui/app.py` | S-10 | - |
| `nmb/tui/catalog.py` | S-10 | - |
| `nmb/tui/cli.py` | S-10 | - |
| `nmb/tui/client.py` | S-10 | - |
| `nmb/tui/logo.py` | S-10 | - |
| `nmb/tui/logs.py` | S-10 | - |
| `nmb/tui/render.py` | S-10 | - |
| `nmb/web/__init__.py` | S-01 | - |
| `nmb/web/__main__.py` | S-01 | - |
| `nmb/web/assets/dashboard.css` | U-01 | P-09 (1236) |
| `nmb/web/assets/dashboard.html` | U-16 | U-09 (827), U-17 (707), U-14 (512), U-07 (411), U-08 (207), U-05 (187), U-18 (186), U-12 (177), U-21 (159), U-01 (128), U-06 (119), U-04 (109), P-09 (89), U-10 (81), U-02 (76), U-15 (66), U-20 (61), U-13 (60), U-03 (41), U-11 (33), U-19 (29) |
| `nmb/web/assets/dashboard.js` | U-09 | U-16 (1572), U-02 (1236), U-17 (1162), U-07 (1046), U-14 (784), U-08 (508), U-05 (497), U-04 (424), U-18 (361), U-12 (349), U-01 (340), U-06 (276), U-21 (267), U-20 (145), U-10 (108), U-13 (80), P-09 (50), U-19 (40), U-11 (11) |
| `nmb/web/assets/logout.html` | U-01 | - |
| `nmb/web/i18n.py` | C-06 | - |
| `nmb/web/locales/es.json` | C-06 | - |
| `nmb/web/locales/it.json` | C-06 | - |
| `nmb/web/locales/pt.json` | C-06 | - |
| `nmb/web/locales/pt_br.json` | C-06 | - |
| `nmb/web/locales/ro.json` | C-06 | - |
| `nmb/web/locales/ui_core.json` | C-06 | - |
| `nmb/web/locales/zh.json` | C-06 | - |
| `nmb/web/results_page.py` | S-09 | - |
| `nmb/web/server.py` | S-08 | S-03 (730), S-01 (679), S-06 (501), S-02 (442), S-04 (425), S-05 (161), S-09 (143), S-07 (137) |
| `nmb/web/static/themes/asus_bottom_bg.png` | P-09 | - |
| `nmb/web/static/themes/asus_middown_bg.png` | P-09 | - |
| `nmb/web/static/themes/asus_title.png` | P-09 | - |
| `nmb/web/static/themes/asus_title_bg.png` | P-09 | - |
| `nmb/web/static/themes/hago.jpg` | P-09 | - |
| `nmb/web/static/themes/ndefault_bg_center.png` | P-09 | - |
| `nmb/web/static/themes/ndefault_bg_line.png` | P-09 | - |
| `nmb/web/static/themes/ndefault_box_header.png` | P-09 | - |
| `nmb/web/static/themes/ndefault_header_bg.jpg` | P-09 | - |
| `nmb/web/static/themes/ndefault_info_bar.png` | P-09 | - |
| `nmb/web/static/themes/ndefault_main_top.jpg` | P-09 | - |
| `nmb/web/static/themes/ndefault_menu_bar.png` | P-09 | - |
| `nmb/web/static/themes/ndefault_top_strip.png` | P-09 | - |
| `nmb/web/static/themes/tplink_blue.jpg` | P-09 | - |
| `nmb/web/static/themes/tplink_top1_1.jpg` | P-09 | - |
| `nmb/web/static/themes/tplink_top1_2.jpg` | P-09 | - |
| `nmb/web/static/themes/tplink_top_bg.jpg` | P-09 | - |
| `nmb/web/static/themes/tuf_camo_bg.png` | P-09 | - |
| `nmb/web/static/themes/tuf_logo.png` | P-09 | - |
| `nmb/web/static/themes/tuf_title_bg.png` | P-09 | - |
| `nmb/web/templates.py` | U-01 | - |
| `os/90-nmb-usb-import.rules` | P-03 | - |
| `os/90-nmb.pkla` | P-06 | - |
| `os/90-nmb.rules` | P-06 | - |
| `os/99-nmb-performance.conf` | N-01 | - |
| `os/net-detect.sh` | S-10 | - |
| `os/nic-tune.sh` | N-01 | - |
| `os/nmb-collector.service` | S-13 | P-04 (6) |
| `os/nmb-guard.nft` | N-02 | - |
| `os/nmb-wait-online.conf` | P-10 | - |
| `os/ooni-resolv.conf` | M-13 | - |
| `os/openrc/nmb-db` | P-10 | - |
| `os/openrc/nmb-netd` | P-10 | - |
| `os/openrc/nmb-probe` | P-10 | - |
| `os/openrc/nmb-telegram` | P-10 | - |
| `os/openrc/nmb-web` | P-10 | - |
| `os/storage-setup.sh` | P-11 | - |
| `setup.sh` | P-11 | P-03 (62) |
| `systemd/nmb-adblock-update.service` | P-10 | P-04 (2) |
| `systemd/nmb-adblock-update.timer` | P-10 | - |
| `systemd/nmb-adblock.service` | P-10 | P-04 (8) |
| `systemd/nmb-db.service` | P-10 | P-04 (2) |
| `systemd/nmb-guard.service` | P-10 | - |
| `systemd/nmb-heal.service` | P-10 | P-04 (3) |
| `systemd/nmb-heal.timer` | P-10 | - |
| `systemd/nmb-invasive-log.service` | P-10 | P-04 (2) |
| `systemd/nmb-invasive-log.timer` | P-10 | - |
| `systemd/nmb-iperf3.service` | P-10 | P-04 (2) |
| `systemd/nmb-maintain.service` | P-10 | P-04 (2) |
| `systemd/nmb-maintain.timer` | P-10 | - |
| `systemd/nmb-netd.service` | P-10 | P-04 (13) |
| `systemd/nmb-ooni.service` | P-10 | P-04 (4) |
| `systemd/nmb-ooni.timer` | P-10 | - |
| `systemd/nmb-probe.service` | P-10 | P-04 (13) |
| `systemd/nmb-report.service` | P-10 | P-04 (7) |
| `systemd/nmb-report.timer` | P-10 | - |
| `systemd/nmb-snapshot.service` | P-10 | P-04 (2) |
| `systemd/nmb-snapshot.timer` | P-10 | - |
| `systemd/nmb-speed.service` | P-10 | P-04 (7) |
| `systemd/nmb-telegram.service` | P-10 | P-04 (9) |
| `systemd/nmb-tune.service` | P-10 | - |
| `systemd/nmb-usb-import@.service` | P-10 | - |
| `systemd/nmb-web.service` | P-10 | P-04 (12) |
| `systemd/nmb.target` | P-10 | - |
| `tests/_box_isolation.py` | F-06 | - |
| `tests/test_adblock.py` | N-12 | - |
| `tests/test_adblock_hardening.py` | N-12 | - |
| `tests/test_auto_mode.py` | N-04 | - |
| `tests/test_auto_pppoe_once.py` | N-04 | - |
| `tests/test_backup_send_encrypted.py` | P-03 | - |
| `tests/test_bot_authorization.py` | S-11 | - |
| `tests/test_bot_client_backoff.py` | S-11 | - |
| `tests/test_bot_dispatch.py` | S-11 | - |
| `tests/test_bot_locales.py` | C-06 | - |
| `tests/test_client_address_round2.py` | S-13 | C-05 (51), R-01 (30), M-09 (21), S-11 (12) |
| `tests/test_cloudflare_engine.py` | M-06 | - |
| `tests/test_collector.py` | S-13 | - |
| `tests/test_config_merge.py` | C-01 | - |
| `tests/test_cross_traffic.py` | M-03 | M-04 (56) |
| `tests/test_csv_exports.py` | R-02 | M-04 (21), C-05 (11) |
| `tests/test_data_cap.py` | M-05 | S-11 (27) |
| `tests/test_dns_fallback.py` | M-12 | - |
| `tests/test_donate.py` | P-08 | - |
| `tests/test_dossier.py` | R-03 | - |
| `tests/test_driver_cli.py` | N-17 | - |
| `tests/test_duckdns.py` | N-16 | S-11 (77) |
| `tests/test_engine_errors.py` | M-04 | N-16 (43), M-08 (16) |
| `tests/test_failover_and_cleanup.py` | N-03 | C-08 (24) |
| `tests/test_fastcom.py` | M-07 | M-09 (25), S-02 (23) |
| `tests/test_feature_sweep_fixes.py` | N-11 | P-03 (33), N-08 (24), N-14 (16), C-05 (13), N-07 (7) |
| `tests/test_firewall_every_mode.py` | N-02 | - |
| `tests/test_first_run_fixes.py` | C-01 | N-02 (37), P-06 (31), P-03 (20), M-05 (15), M-09 (15), P-04 (12), N-07 (11), P-01 (11), R-01 (11), M-03 (10), S-01 (9), P-11 (8), N-01 (6) |
| `tests/test_github_updater.py` | P-02 | - |
| `tests/test_headless_features.py` | P-10 | C-03 (69), C-05 (63), N-01 (57), M-05 (43), N-08 (40), M-08 (23), S-01 (23), R-05 (15) |
| `tests/test_heal.py` | P-01 | - |
| `tests/test_helper_allowlists.py` | P-06 | - |
| `tests/test_i18n_locales.py` | C-06 | - |
| `tests/test_iperf3_rest.py` | M-04 | - |
| `tests/test_ipv6_lan.py` | N-11 | - |
| `tests/test_lan_topology.py` | N-06 | - |
| `tests/test_llm_bundle.py` | R-04 | - |
| `tests/test_llm_bundle_honesty.py` | R-04 | - |
| `tests/test_llm_bundle_masking.py` | R-04 | - |
| `tests/test_log_budget.py` | C-03 | C-04 (22), C-09 (9) |
| `tests/test_measure_lock.py` | M-04 | - |
| `tests/test_measurement_setup.py` | C-11 | M-04 (28), N-03 (24), M-02 (13), S-13 (10) |
| `tests/test_mixed_engines.py` | M-04 | - |
| `tests/test_no_made_up_values.py` | C-02 | R-02 (34), M-10 (33), C-05 (32), M-08 (30), M-09 (22), M-06 (21), R-01 (21), S-11 (13) |
| `tests/test_ookla_honesty.py` | M-04 | M-10 (126), M-05 (54), M-09 (41), R-01 (22) |
| `tests/test_ooni_cli.py` | M-13 | - |
| `tests/test_ooni_toggle.py` | M-13 | - |
| `tests/test_openvpn_names.py` | N-14 | - |
| `tests/test_own_outages.py` | M-02 | C-07 (49) |
| `tests/test_port_forward.py` | N-11 | - |
| `tests/test_pppoe_retry.py` | N-05 | N-10 (77), N-02 (54), N-01 (31), N-03 (28) |
| `tests/test_probe_line.py` | M-01 | M-02 (50), N-01 (35), S-11 (16), R-01 (9) |
| `tests/test_probe_outage.py` | M-02 | - |
| `tests/test_public_results_and_ookla_warning.py` | S-09 | M-05 (39) |
| `tests/test_purge_and_verdict.py` | C-08 | C-06 (33) |
| `tests/test_redact.py` | C-04 | - |
| `tests/test_regdomains.py` | N-07 | - |
| `tests/test_review_fixes.py` | R-01 | R-02 (83), R-03 (38), S-11 (32), S-12 (30), C-02 (25), M-02 (22), M-04 (19), R-05 (11) |
| `tests/test_rtr_native.py` | M-08 | - |
| `tests/test_rtr_output_tail.py` | M-08 | - |
| `tests/test_secret_keep.py` | C-04 | - |
| `tests/test_selftest.py` | P-07 | - |
| `tests/test_selftest_auto.py` | P-07 | - |
| `tests/test_selftest_router.py` | P-07 | - |
| `tests/test_selftest_scenario.py` | P-07 | - |
| `tests/test_setup_bot.py` | S-11 | P-07 (71), S-13 (41) |
| `tests/test_setup_reports.py` | R-02 | R-03 (91), R-04 (52), R-05 (27) |
| `tests/test_setup_web.py` | S-02 | U-02 (100), S-10 (48) |
| `tests/test_shipped_files.py` | P-11 | - |
| `tests/test_speed_schedule_default.py` | M-04 | - |
| `tests/test_storage_guard.py` | C-08 | - |
| `tests/test_tailscale_and_telegram_ui.py` | N-15 | C-06 (52), C-01 (27), U-09 (13), U-18 (11), U-01 (10) |
| `tests/test_tasks.py` | C-09 | - |
| `tests/test_telegram_bot.py` | S-11 | - |
| `tests/test_telegram_i18n.py` | C-06 | - |
| `tests/test_theme.py` | P-09 | - |
| `tests/test_traceroute.py` | M-11 | - |
| `tests/test_tui.py` | S-10 | - |
| `tests/test_ui_navigation_targets.py` | U-01 | - |
| `tests/test_ui_text_coverage.py` | C-06 | - |
| `tests/test_usb_import.py` | P-03 | - |
| `tests/test_usb_router.py` | N-10 | N-01 (44), N-03 (18), N-07 (17), N-08 (10) |
| `tests/test_usb_tether.py` | N-10 | - |
| `tests/test_usb_uplink.py` | N-10 | - |
| `tests/test_wait_online.py` | P-10 | - |
| `tests/test_web_ui_integrity.py` | U-01 | - |
| `tests/test_wifi_band_caps.py` | N-07 | - |
| `tests/test_wifi_flags.py` | N-07 | - |
| `tests/test_wifi_psk.py` | N-07 | - |
| `tests/test_wifi_reload.py` | N-03 | - |
| `tests/test_wifi_share.py` | N-08 | N-05 (67), C-01 (30) |
| `tests/test_wifi_uplink.py` | N-08 | - |
| `tests/test_write_private.py` | N-03 | - |

Files cut over more than one package: 61 of 351. Read literally, work plan line 217 ("every Python file
... in exactly one package") is broken by each of them; see section D.1 for the reading the register uses.
