# Requirements register: questions for the owner

The questions the parts raised, de-duplicated (44 asked, 39 numbered here; Q-06 and Q-07 were answered by other rows in round 1
and moved to the end, so 37 are open). Each names the part(s) that asked it, the rows it
affects, the package that needs the answer, and the proposal of the part where it made one. Until a question is answered
the rows stay as written. Questions the owner already settled in `docs/NATIVE_WORK_PLAN.md@ad3758c` section 7 are not asked
again; where a question falls under such a decision, it says so. Answers go to F-01, which writes them into the roadmap.

Part names: RM1 `roadmap-1-507.md`, RM2 `roadmap-508-1034.md`, PP `porting-plan.md`, WP `workplan-architecture-devsetup.md`,
TS `testing.md`, RE `readme.md`, CL `changelog.md`, PY `python-files.md`.

## Scope and order of the work

| No. | Question | Asked by | Rows | Needed by | Proposal in the part |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Q-01 | Must the native components run **alongside the Python version on one box**, component by component (Phase 1.1: native probe writing into the Python SQLite database), or is replacing it with the whole `nmbd` enough (QB 5, decision 5)? | RM2 1 | REQ-RM-1002, -1260 | P-13, M-01 | none |
| Q-02 | **F-07 in the order:** work plan 3.0 puts F-00 to F-07 before anything else, section 4 step 1 lists only F-00 to F-06. Is `spec/` finished before Core starts, or does it grow alongside? | WP 1 | REQ-WP-0004 ff. | F-00, F-07 | done first, as rule 11 of 2.2 implies |
| Q-03 | **ARMv6 (Pi 1, Pi Zero):** roadmap line 73 and section 8 (`GOARM=6`) name them; the step-1 matrix has arm v7 only. Add `linux/arm` v6, or leave Pi Zero and Pi 1 out of step 1? | RM1 7, RM2 6 | REQ-RM-0013, REQ-RM-1440 | F-00, P-13 | none |
| Q-04 | **Projections as limits?** Are the RAM and CPU figures of roadmap sections 7 and 9 (for example 10-18 MB RSS, idle CPU below 0.5 %) acceptance limits for P-13, or targets to measure and report? | RM2 5 | REQ-RM-1413 to -1438, -1442, -1443 | P-13 | measure and record (current rows) |
| Q-05 | **Supported distributions:** must native packages still install on every distribution the Python version supported (Debian, Ubuntu, Alpine, Fedora/RHEL family, Rocky), or only on the own distribution and the bases chosen after the image analysis? | CL 1 | CHANGELOG lines 95, 96, 131 rows | P-11 | none |
| Q-08 | **OpenWrt package format:** Phase 4.3 says `.ipk`/`opkg`; OpenWrt 25.12 (the step-1 images) uses `apk`. Assumed: the format follows the target release (APK for 25.12, `.ipk` for the 22.03 base of the TL-WR1043ND v1). Confirm. | RM2 4 | REQ-RM-1271 | P-11 | as assumed |
| Q-09 | **Documentation languages:** does the English-and-Hungarian rule of the Python repository hold for the native repository (manual and command reference of P-12), or does it follow the 11 UI languages, or English only? | WP 5 | DEV_SETUP line 200 rows | P-12 | none |
| Q-10 | **Test scripts in the repository** (REQ-TS-0312 is marked pending until this is answered): the Python known limit (VM and bot check scripts kept outside the repository) became a requirement that the native scripts live in the native repository. Confirm. | TS 6 | TESTING line 399 row | F-06 | as written |

## Security and privacy

| No. | Question | Asked by | Rows | Needed by | Proposal in the part |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Q-11 | **Third-party code in-process or sandboxed:** section 7 and Phase 3.2 put OONI's probe-engine and the iPerf3 client inside `nmbd`; 6B.5 runs third-party code in a stricter sandbox that never sees `nmb.conf`, keys or the database. Is a vetted library in-process allowed? (The OONI code also sits in an `internal/` folder and is GPL-3.0, which affects the licence, decision 9.) | RM2 2, RM1 note | REQ-RM-1267, -1350, -1417, -1418 | P-04, M-13, M-09 | none |
| Q-12 | **Browser terminal:** `/api/terminal/exec` runs shell commands as root from the browser (off unless `terminal_enabled`). Port it at all? Same question for the **experimental stop on network traffic** (changelog line 373). | PY 3, CL 3 | CHANGELOG lines 373, 374 rows; page U-20 | S-08, U-20, M-03 | none |
| Q-13 | **Web server capabilities:** the Python web server holds `CAP_NET_ADMIN` and `CAP_NET_RAW`. Keep that, or give the native web server no network-admin rights and pass such tasks to the network daemon or the privileged helper (QB 7)? | WP 2 | ARCHITECTURE line 45 rows | S-01, P-04 | none |
| Q-14 | **`/kill` from Telegram** signals any PID above 1, with `sudo` if needed (it can stop sshd or Tailscale). Limit the native `/kill` to NMB's own processes? | WP 3 | REQ-AR rows of line 135; REQ-RE-0280 | S-11 | none |
| Q-15 | **Secrets in Telegram messages:** a WireGuard QR code holds the client's private key and the bot would send it through Telegram; `/wifi connect <ssid> [pw]` puts the Wi-Fi password in a message Telegram keeps. Keep both, or have the bot point to the Web UI for these? | RM1 4, WP 4 | roadmap line 300 rows; ARCHITECTURE line 136 rows | S-11, N-13 | none |
| Q-16 | **MAC addresses in shared diagnostics:** the redactor keeps MAC addresses in bundles that go to outside LLMs or ISP support. Intended under QB 8, or masked (for example vendor prefix only)? | RM1 3 | roadmap line 386 rows | C-04, R-04 | none |
| Q-17 | **Audit log "the session cannot erase":** a root shell can erase any local file. Send the log off the device as it is written (collector or syslog), make the maintenance shell a restricted user with logged `sudo`, or accept a local append-only log with the limit stated? | PP 6 | REQ-PP-0135 | P-05 | one of the three |
| Q-18 | **Login method:** roadmap 4.4 (line 151) names HTTP Basic Auth; work plan S-01 has login, sessions and CSRF, and the roadmap wins where they differ. Falls under decision 1 (API left to the developer) if the owner agrees. | RM1 1 | roadmap line 151 rows | S-01 | session login (Basic Auth line out of date) |
| Q-19 | **Factory login warning:** the Python installer warns on every run while the Web UI factory login is set. The owner keeps that login on purpose. Should the native installer keep the warning, drop it, or show it once? | CL 2 | CHANGELOG line 211 row | P-11 | kept as written until answered |
| Q-20 | **Donation addresses** appear in the README, the Web UI, the bot and the CLI. Do they go into the native code as product data, or are they supplied at build or release time from outside the code (QB 8)? | RE 4 | REQ-RE-0310, -0311 | P-08 | none |
| Q-21 | **Theme images:** may the vendor-look images (`asus_title.png`, `tuf_logo.png`, `tplink_*.jpg`) ship as add-ons of the open native repository, or be redrawn without logos or left out? And is `themes/hago.jpg` a personal photo or place (then it stays out, QB 8)? | PY 1, 2 | python-files rows of `nmb/web/static/themes/` | P-09 | none |
| Q-22 | **Origin and licence of `components/open-rmbt/`:** the repository has no licence file for the C client. Upstream RTR code (which project and version) or written for NMB? The licence of the native build follows from its components (decision 9). | PY 4 | python-files rows of `components/open-rmbt/` | M-08, P-08 | none |

## Measuring

| No. | Question | Asked by | Rows | Needed by | Proposal in the part |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Q-23 | **NAT type classifier:** a plain RFC 5389 Binding Request cannot tell full cone from restricted cone (needs RFC 5780 or several servers). May the classifier report "unknown" when no such server answers? | RM1 5 | roadmap line 359 rows | M-12 | report "unknown" |
| Q-24 | **Flow Cache bypass:** the plan turns off Broadcom acceleration with a shell call and does not restore it. Fault? | PP 1 | REQ-PP-0051 | X2-01 | call the interface directly (no shell) and always restore the previous state, also after a crash |
| Q-25 | **CAKE shaping on the measured line:** the speed tests would measure the shaper. Off (or lifted) during measurements, or on with the shaper recorded in the measurement setup (4.29)? | PP 2 | REQ-PP-0053, -0186 | X2-01, C-11 | none |
| Q-26 | **DynDNS from the local WAN address** is wrong behind NAT or CGNAT. Fault? | PP 5 | REQ-PP-0088 | N-16 | use the netlink address only when public, otherwise the public-address check (M-12) |
| Q-27 | **Fast.com over plain HTTP:** a fixed rule for the native engine, or may it use TLS when it measures as well? | CL 4 | CHANGELOG line 82 row | M-07 | the M-07 analyst compares both |
| Q-28 | **Manual and bot Ookla runs against the hourly limit:** counting them means a manual run inside the hour is refused or deferred. Intended? (Decision 3 settles how Ookla runs, not this.) | TS 3 | TESTING line 386 row | M-05, M-10 | none |
| Q-29 | **Default data cap on a metered link:** set a default daily cap on USB and mobile uplinks, or only warn (the box ships with 0)? | TS 2 | TESTING line 388 row | M-05, N-10 | none |
| Q-30 | **RTR over 5G** measured far above the other engines and was never settled. Fault to find in the M-08 brief, or accepted engine behaviour? | TS 5 | TESTING line 179 rows | M-08 | none |
| Q-31 | **Bloom filter false positives** (0.1 %) block domains that are on no list. Accepted on small routers, or every hit confirmed by an exact check before blocking? | PP 4, RM1 note | REQ-PP-0078 | N-12, X2-01 | confirm exactly before blocking |

## Network and modes

| No. | Question | Asked by | Rows | Needed by | Proposal in the part |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Q-32 | **Mode B when the box fails:** README.md:92 says the connection stops with the box; README.hu.md:169 claims the software bridge or bypass hardware still passes traffic. Is a fail-open bypass required? | RE 1 | REQ-RE-0028, -0315 | N-03 | none |
| Q-33 | **Wi-Fi client without NetworkManager:** Python fails there by design. Own nl80211/WPA control natively, or keep the dependency and report it clearly? | TS 1 | TESTING line 142 row | N-07 | none |
| Q-34 | **Not built in Python:** IPv6 for the LAN on a DHCP or USB WAN, and dual-band with a second radio. In step 1, or recorded as a limit? | TS 4 | TESTING lines 319, 330 rows | N-11, N-07 | none |
| Q-35 | **No PPPoE on the TL-WR740N:** the 4 MiB recipe removes PPP, so that target has no router mode. Accepted for Tier 3 (the UI then says so plainly)? | PP 3 | REQ-PP-0061 | X2-01 | none |
| Q-36 | **ip, nft and iw in the image:** 6A lists them as image contents; section 7 and Phase 2.1 remove `nmbd`'s calls of them. Kept only for people and recovery, or left out? | RM2 3 | REQ-RM-1263, -1277, -1414, -1425 | P-11 | none |

## Surfaces and names

| No. | Question | Asked by | Rows | Needed by | Proposal in the part |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Q-37 | **Languages of the bot and the dossier:** roadmap lines 301 and 308 list 5 languages, line 406 and the Web UI have 11. Are the 5-language lists out of date (all 11 used)? The Python bot has 5 core languages plus 6 locale files (`nmb/bot/i18n.py@ad3758c:27`). | RM1 2, PP note, RE note | roadmap lines 301, 308 rows | S-11, R-03, C-06 | all 11 (C-06) |
| Q-38 | **Compatibility command names `pihole` and `pivpn`** (`setup.sh@ad3758c:486-489`; roadmap line 503): install them from the native multi-call binary (they can clash with a real Pi-hole or PiVPN and with scripts that expect their own CLI), or only the `nmb-*` names? | RM1 6, PY 5 | roadmap line 503 rows | P-10, N-12, N-13 | none |
| Q-39 | **Collector server "not ported"** (porting plan) against work plan package S-13. Reading: not part of the router builds; S-13 stays in step 1 for the PC/server build. Correct? | PP 7 | REQ-PP-0193 | S-13 | as read |

## Merged and not asked again

- The admin/admin default login of the Web UI (REQ-RE-0231) against "unique password per device" of P-03: recorded in the
  row's Notes by the readme part; the owner keeps the default on purpose, so it is not a question.
- Ookla protocol emulation (REQ-PP-0080) and Ookla "eliminated" in roadmap section 7: superseded by decision 3.
- The disclaimer wording (REQ-RM-1344): already open as point 7 of work plan section 7.
- Q-06 (Ubuntu 22.04, polkit 0.105, service user may manage any unit; REQ-RE-0288): answered in round 1 by the Quality Bar,
  points 6 and 7 (known Python faults are not carried over; no security hole is left open) and by REQ-TS-0018 and
  REQ-CL-0365: the native privilege rule is an allow-list of NMB's own units on every base. REQ-RE-0288 states that; owner P-06.
- Q-07 (Alpine without the public-address guard; REQ-RE-0290): answered in round 1 by QB 7 (nothing listens on the WAN side
  unless the owner turned it on, `docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:587`) and by REQ-TS-0129 and REQ-CL-0205 (the
  firewall in every mode): the native guard runs on every base, Alpine/OpenRC included. Whether Alpine is a supported base at
  all stays with Q-05 and the image comparison (I-09).
