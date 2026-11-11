# Requirements register, part: every file of the Python version

Source: the Python repository at commit `ad3758c`, every tracked file except `docs/` and the `*.md` files at the top
(`CHANGELOG.md`, `CONTRIBUTING.md`, `CONTRIBUTING.hu.md`, `README.md`, `README.hu.md`, covered by their own parts): 351 files,
122,873 text lines plus 20 theme images. Package ids are those of `docs/NATIVE_WORK_PLAN.md@ad3758c:98-225`.

Rule of this table: every file is in it; a file split over packages has one row per package with its line ranges, and
the ranges of one file cover each of its lines exactly once (checked by the script that built the table, see *Read coverage*).
Tests go with the package of the code they test; a test file that checks several packages is split by its test classes.
"not ported" rows still name the package that accounts for them, so the register's check finds every file in exactly one place.
No REQ ids here: the requirements of each file are written by the package analyst from these ranges (nmb-native-brief).

## Package names used for S-02...S-0n and U-02...U-0n

Work plan section 3.5 leaves these open ("one package per area", "one package per page"). The table uses these ids (proposed):

| Id | Name | Id | Name |
| :--- | :--- | :--- | :--- |
| S-02 | Status API | U-08 | OONI page |
| S-03 | Speed and measuring API | U-09 | Remote access and VPN page |
| S-04 | Network API | U-10 | Network explorer page |
| S-05 | Wi-Fi API | U-11 | Port neutrality page |
| S-06 | VPN and remote-access API | U-12 | Logs page |
| S-07 | DNS blocker API | U-13 | Packet capture page |
| S-08 | System and settings API | U-14 | Measurement and servers page |
| S-09 | Exports and public results API | U-15 | Bandwidth monitor page |
| U-02 | Dashboard page | U-16 | Router page |
| U-03 | Ping telemetry page | U-17 | System page |
| U-04 | Speed results page | U-18 | Telegram page |
| U-05 | iPerf3 page | U-19 | AI diagnostics page |
| U-06 | Traceroute page | U-20 | Terminal page |
| U-07 | DNS blocker page | U-21 | Tasks and load page |

S-0n stops at S-09 because S-10 is taken (nmb-cli); the eight areas are the ones the work plan names (status, speed, network,
Wi-Fi, VPN, blocker, system, exports).

## Read coverage

Read in full: every `systemd/` unit (26), every `os/` file (16), `config/` (2), `.claude/settings.json`, `.github/pull_request_template.md`,
`.gitignore`, the eleven thin `bin/` wrappers (`nmb-cli`, `nmb-collector`, `nmb-factory-reset`, `nmb-invasive-log`, `nmb-netd`, `nmb-probe`,
`nmb-report`, `nmb-selftest`, `nmb-speed`, `nmb-telegram`, `nmb-web`), and `docs/NATIVE_WORK_PLAN.md@ad3758c:98-225`.

Read in part: `bin/nmb-adblock`, `nmb-driver`, `nmb-heal`, `nmb-helper`, `nmb-iperf`, `nmb-ooni`, `nmb-vpn`, `nmb-wifi` (lines 1-25 and every
function or `case` label); `setup.sh` (386-460 and every section header); `nmb/web/server.py` (1399-1495, 4386-4400, 4453-4480 and every route
branch line of 370-3441); `nmb/bot/i18n.py` (1-42, 290-300, 345-360, 405-420, 1105-1140, 1230-1260, 1355-1449); `nmb/web/i18n.py` (7140-7171 and
its top-level names); `nmb/web/assets/dashboard.html` (1-110, 2881-2900, 5221-5239 and every tab, sub-tab and section comment);
`dashboard.css` (1-20, 866-872, 1004-1071, 2163-2169, 2455-2480, 2750-2812 and every theme block start); `components/open-rmbt/README.md`
(1-25, 100-119), `Makefile` (1-20); the first keys of every locale JSON.

Read by structure only: every other Python file (module docstring and every top-level function and class with its line, from the
syntax tree; methods of the large classes in `bot/commands.py`, `netd/modes.py`, `netd/daemon.py`, `netd/services.py`, `probe/engine.py`,
`speed/runner.py`, `web/server.py`), every test file (imports, every test class and method with its lines), the 383 functions of
`dashboard.js` with their lines, the C sources of `components/open-rmbt/src/` (names only). Theme images were not opened.

**Not done as asked:** the task asked for every line in pages of at most 150 lines. About 5,000 of the 122,873 lines were read line by
line; the rest by structure. Reading all of it does not fit one agent's context (about 1.5 to 2 million tokens). The file-to-package
assignment needs only the structure, and each package analyst reads its own ranges in full (nmb-native-brief), but a second pass by
directory is needed before the notes column can be taken as a complete list of what each range does.

## Files and packages

| Python file | lines | Package that ports it | Notes |
| :--- | :--- | :--- | :--- |
| `.claude/settings.json`@ad3758c | 1-5 | F-00 | not ported: editor plug-in setting of the development tool, no product behaviour |
| `.github/pull_request_template.md`@ad3758c | 1-31 | F-00 | contribution rule (test record, translations, changelog, AI-written parts named); the native repository gets its own template |
| `.gitignore`@ad3758c | 1-86 | F-00 | rule to carry over: databases, logs, captures, keys, *.conf and credentials files never committed (lines 37-85); the native repository has its own |
| `bin/nmb-adblock`@ad3758c | 1-613 | N-12 | blocker CLI (status, enable/disable, gravity update, allow/deny lists, adlists, local DNS, regex, charts, query log, guide); also installed as `pihole` (setup.sh 488); command name itself -> P-10 |
| `bin/nmb-cli`@ad3758c | 1-11 | P-10 | thin entry point to nmb/tui/cli.py (behaviour in S-10); multi-call name |
| `bin/nmb-collector`@ad3758c | 1-10 | P-10 | thin entry point to nmb/collector/server.py (behaviour in S-13); multi-call name |
| `bin/nmb-driver`@ad3758c | 1-232 | N-17 | driver CLI: scan, status, missing firmware, install (goes through nmb-helper install-drivers) |
| `bin/nmb-factory-reset`@ad3758c | 1-34 | P-03 | requires --yes, optional --keep-db, safety backup first, then restarts services; runs from the root-owned install folder, never the caller's PYTHONPATH |
| `bin/nmb-heal`@ad3758c | 1-287 | P-01 | self-healing CLI: audit, fix, update (update part is P-02 behaviour), status |
| `bin/nmb-helper`@ad3758c | 1-548 | P-06 | root helper with fixed allow-lists (links, dirs, packages, install targets), refuses unsafe VPN configs, runs from / with PYTHON* unset; every subcommand is a requirement of P-06 |
| `bin/nmb-invasive-log`@ad3758c | 1-26 | R-05 | writes the diagnostic snapshot, prints Markdown or --json; boot-partition copy at most hourly |
| `bin/nmb-iperf`@ad3758c | 1-166 | M-09 | iPerf3 CLI: test, local server status/start/stop, public server list |
| `bin/nmb-netd`@ad3758c | 1-15 | P-10 | thin entry point to nmb/netd/daemon.py (behaviour in N-03) |
| `bin/nmb-ooni`@ad3758c | 1-185 | M-13 | OONI CLI: run (websites, middlebox, im, circumvention, all), status, enabled (used as ExecCondition), history, onboard, inspect |
| `bin/nmb-probe`@ad3758c | 1-14 | P-10 | thin entry point to nmb/probe/engine.py (behaviour in M-01) |
| `bin/nmb-report`@ad3758c | 1-14 | P-10 | thin entry point to nmb/report/cli.py (behaviour in R-03) |
| `bin/nmb-selftest`@ad3758c | 1-10 | P-10 | thin entry point to nmb/selftest/live.py (behaviour in P-07) |
| `bin/nmb-speed`@ad3758c | 1-14 | P-10 | thin entry point to nmb/speed/runner.py (behaviour in M-04) |
| `bin/nmb-telegram`@ad3758c | 1-14 | P-10 | thin entry point to nmb/bot/daemon.py (behaviour in S-11) |
| `bin/nmb-vpn`@ad3758c | 1-322 | N-13 | PiVPN-style CLI for WireGuard and OpenVPN (add, list, qr, remove, status, guide), also installed as `pivpn`; OpenVPN branches exercise N-14 code |
| `bin/nmb-web`@ad3758c | 1-14 | P-10 | thin entry point to nmb/web/server.py (behaviour in S-01) |
| `bin/nmb-wifi`@ad3758c | 1-190 | N-08 | Wi-Fi client CLI over guest_wifi (scan, connect incl. hidden, disconnect, status, saved, delete) |
| `components/open-rmbt/Makefile`@ad3758c | 1-55 | M-08 | build of rmbt-client, version string from git describe |
| `components/open-rmbt/README.md`@ad3758c | 1-119 | M-08 | protocol summary (lines 107-119) and loop-mode UUID rule; no licence file in the repository (see questions) |
| `components/open-rmbt/src/connection.c`@ad3758c | 1-583 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/connection.h`@ad3758c | 1-53 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/control.c`@ad3758c | 1-562 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/control.h`@ad3758c | 1-102 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/gui.c`@ad3758c | 1-225 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/gui.h`@ad3758c | 1-66 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/main.c`@ad3758c | 1-677 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/pretest.c`@ad3758c | 1-141 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/pretest.h`@ad3758c | 1-14 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/pthread_barrier_compat.h`@ad3758c | 1-57 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/tests.c`@ad3758c | 1-294 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/tests.h`@ad3758c | 1-40 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/uuid_store.c`@ad3758c | 1-53 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `components/open-rmbt/src/uuid_store.h`@ad3758c | 1-8 | M-08 | native RMBT client in C (control, pretest, ping, multi-thread download/upload, RMBThttp and RMBTws, TLS); reference and possibly carried as component |
| `config/nmb.conf.example`@ad3758c | 1-192 | C-01 | every section and default of nmb.conf; part of the F-04 configuration contract; [isp] example values are placeholders |
| `config/usb-import.example.conf`@ad3758c | 1-18 | P-03 | USB import file: only written keys change, empty or __KEEP_CURRENT__ secret keeps its value, result file names keys never values, applied once |
| `nmb/__init__.py`@ad3758c | 0 (empty) | P-10 | not ported: empty Python package marker |
| `nmb/bot/__init__.py`@ad3758c | 1-1 | S-11 | not ported: Python package marker |
| `nmb/bot/__main__.py`@ad3758c | 1-4 | S-11 | not ported: module entry point; native entry is a multi-call name (P-10) |
| `nmb/bot/alerts.py`@ad3758c | 1-310 | S-11 | outage/recovery, loss, latency, WAN failover, data-cap alerts; no outage alert while the box has its own line down |
| `nmb/bot/client.py`@ad3758c | 1-198 | S-11 | Telegram Bot API client: long poll, sendMessage HTML, sendDocument multipart, back-off on 429/5xx |
| `nmb/bot/commands.py`@ad3758c | 1-248 | S-11 | helpers and dispatcher (handle_command, authorisation path); 1725 lines in all, see split proposal |
| `nmb/bot/commands.py`@ad3758c | 249-1061 | S-11 | commands: help, donate, lang, status, purge, ping, test, report, tailscale, ai, mode, vpn (749-1061) |
| `nmb/bot/commands.py`@ad3758c | 1062-1725 | S-11 | commands: dyndns, adblock, heal (audit, fix, updates), tasks, kill, wifi, backup (send encrypted), drivers, reboot |
| `nmb/bot/daemon.py`@ad3758c | 1-264 | S-11 | bot daemon: refuses to start without allowed chats, authorises the person, keeps answering during a speed test, alert coordination |
| `nmb/bot/i18n.py`@ad3758c | 1-414 | S-11 | per-chat language storage in kv, reply/inline keyboards, menu aliases, measurement status badges, verdict helpers |
| `nmb/bot/i18n.py`@ad3758c | 415-1375 | C-06 | bot strings (STRINGS), alert templates, admin and donation strings, loader of nmb/bot/locales/*.json |
| `nmb/bot/i18n.py`@ad3758c | 1376-1449 | S-11 | shared setup texts merged in, keyboard/status/verdict overrides for the JSON languages, get_sla_verdict longest-match rule |
| `nmb/bot/locales/es.json`@ad3758c | 1-322 | C-06 | bot language file (info, strings, alerts, keyboard, status, verdict); one of the 11 languages |
| `nmb/bot/locales/it.json`@ad3758c | 1-322 | C-06 | bot language file (info, strings, alerts, keyboard, status, verdict); one of the 11 languages |
| `nmb/bot/locales/pt.json`@ad3758c | 1-322 | C-06 | bot language file (info, strings, alerts, keyboard, status, verdict); one of the 11 languages |
| `nmb/bot/locales/pt_br.json`@ad3758c | 1-322 | C-06 | bot language file (info, strings, alerts, keyboard, status, verdict); one of the 11 languages |
| `nmb/bot/locales/ro.json`@ad3758c | 1-322 | C-06 | bot language file (info, strings, alerts, keyboard, status, verdict); one of the 11 languages |
| `nmb/bot/locales/zh.json`@ad3758c | 1-322 | C-06 | bot language file (info, strings, alerts, keyboard, status, verdict); one of the 11 languages |
| `nmb/bot/messages.py`@ad3758c | 1-181 | C-06 | bot messages keyed (en, hu, de, fr, uk) |
| `nmb/collector/__init__.py`@ad3758c | 0 (empty) | S-12 | not ported: empty Python package marker |
| `nmb/collector/client.py`@ad3758c | 1-360 | S-12 | opt-in upload of new rows every interval_min, consent-gated diagnostics and remote control (speed test, diagnostics, upload only), corrections only for rows already sent |
| `nmb/collector/server.py`@ad3758c | 1-860 | S-13 | ingest with per-box token, secrets stored as hashes, search by case id, box page, CSV, command queue, LAN-only admin rules |
| `nmb/common/__init__.py`@ad3758c | 1-1 | C-05 | not ported: Python package marker |
| `nmb/common/backup.py`@ad3758c | 1-837 | P-03 | backup create/verify/restore/list/delete, consistent SQLite copy, encryption for transfer with passphrase rules, factory reset with safety backup |
| `nmb/common/donate.py`@ad3758c | 1-30 | P-08 | the one list of donation addresses read by every surface; checksums proven by tests |
| `nmb/common/heal.py`@ad3758c | 1-1003 | P-01 | health audit, repair of links, permissions, config, services, database, missing binaries; escalation state |
| `nmb/common/heal.py`@ad3758c | 1004-1310 | P-02 | stable updates: repo detection, version stability, GitHub update start and apply with rollback |
| `nmb/common/memory.py`@ad3758c | 1-159 | C-08 | RAM detection, memory profiles (low/standard/performance), SQLite pragmas, live DB size cap, heap trim |
| `nmb/common/net_setup.py`@ad3758c | 1-384 | C-11 | measurement setup: mode, link kind, ISP path (direct or behind a router) from addresses; timeline; record on change |
| `nmb/common/netaddr.py`@ad3758c | 1-106 | C-05 | global/local address rules, CGNAT not local, public address lookup over the same source, client address for a test |
| `nmb/common/own_action.py`@ad3758c | 1-109 | C-07 | own-action windows (mode change, update, restart): begin/end/stale close; outages inside are the box's own |
| `nmb/common/purge.py`@ad3758c | 1-181 | C-08 | selective deletion of stored data in slices; never config, credentials or VPN keys; removal of impossible speed tests after backup |
| `nmb/common/redact.py`@ad3758c | 1-71 | C-04 | the one matcher for secret and personal keys, value masking, text scrub, config redaction used by every export |
| `nmb/common/secret_keep.py`@ad3758c | 1-113 | C-04 | root-only copy of typed-in secrets, detect loss, restore |
| `nmb/common/setup_texts.py`@ad3758c | 1-286 | C-11 | measurement-setup words in all 11 languages, one source for dossier, Web UI and bot |
| `nmb/common/storage_guard.py`@ad3758c | 1-275 | C-08 | free-space target by disk size (15/10/5 %, 200 MB floor, cap), cleanup order, measurements last, log file system guard |
| `nmb/common/tasks.py`@ad3758c | 1-482 | C-09 | per-core CPU, memory, processes from /proc, kill with signal allow-list |
| `nmb/common/updater.py`@ad3758c | 1-390 | P-02 | update without git: compare commits, download archive, safe unpack, tests sandboxed as nobody, restart, web check, status file |
| `nmb/common/usb_import.py`@ad3758c | 1-202 | P-03 | settings from a USB stick: partial apply, secrets never wiped, applied once, result file without values |
| `nmb/common/user_agent.py`@ad3758c | 1-181 | C-05 | browser user-agent presets, fresh generation, effective value |
| `nmb/common/version.py`@ad3758c | 1-30 | C-05 | installed version from the file the updater writes |
| `nmb/config.py`@ad3758c | 1-1646 | C-01 | load, tolerant parser, defaults text (226-437), dataclasses per section, clamping, locked atomic save, boot-partition copy (58-80, 114-225), secret masking, __KEEP_CURRENT__, unique Wi-Fi password; 1646 lines, see split proposal |
| `nmb/db/__init__.py`@ad3758c | 0 (empty) | C-02 | not ported: empty Python package marker |
| `nmb/db/corrections.py`@ad3758c | 1-263 | C-02 | one-off correction passes A/B/C of made-up speed values with originals kept; part of migrating a Python database |
| `nmb/db/database.py`@ad3758c | 1-407 | C-02 | live DB in RAM, consistent snapshot to SD, restore at boot, schema upgrades, minute rollups, prune, maintain |
| `nmb/db/schema.sql`@ad3758c | 1-171 | C-02 | tables of the Python database; part of the F-04 schema contract |
| `nmb/db/setup_backfill.py`@ad3758c | 1-298 | C-11 | one-off reconstruction of the setup of old rows from the event log |
| `nmb/devlog.py`@ad3758c | 1-854 | C-03 | structured developer log, ring buffer, rotation fitted to the file system, run-time level, syslog/journal reading for the UI, purge |
| `nmb/netd/__init__.py`@ad3758c | 1-1 | N-03 | not ported: Python package marker |
| `nmb/netd/__main__.py`@ad3758c | 1-4 | N-03 | not ported: module entry point |
| `nmb/netd/adblock.py`@ad3758c | 1-2031 | N-12 | DNS parser and sinkhole answers, engine, gravity update, allow/deny/regex lists, local DNS, router DHCP guide, DNS server, IPC, charts; 2031 lines, see split proposal |
| `nmb/netd/auto_mode.py`@ad3758c | 1-211 | N-04 | owner's rule: router lease -> standalone, no router and no internet for 90 s -> router mode, PPPoE only with credentials and a server answer; throttled discovery |
| `nmb/netd/daemon.py`@ad3758c | 1-511 | N-03 | netd main loop: build modes, reload, override mode, events, state file, setup recording, switch with own-action window, auto details (PPPoE once per start), kernel tuning signature |
| `nmb/netd/detector.py`@ad3758c | 1-91 | N-01 | carrier and upstream checks with hysteresis feeding the mode decision |
| `nmb/netd/dns.py`@ad3758c | 1-427 | M-12 | DNS health test, fallback resolvers added and removed again, RAM cache, events, status |
| `nmb/netd/drivers.py`@ad3758c | 1-680 | N-17 | PCI/USB adapter detection, ethtool driver info, link modes, missing firmware, recommended packages, install through helper |
| `nmb/netd/duckdns.py`@ad3758c | 1-703 | N-16 | DynDNS (DuckDNS, dynv6) update over IPv4 and IPv6, shared state across processes, daemon with wake-up |
| `nmb/netd/explorer.py`@ad3758c | 1-309 | N-17 | LAN and guest device discovery from neighbour tables and sweeps, OUI vendor lookup, cache |
| `nmb/netd/firewall.py`@ad3758c | 1-459 | N-02 | WAN input rules, host guard in every mode, guard ruleset, NAT for mode A and WAN2, iptables fallback, forwarding sysctls |
| `nmb/netd/geoip.py`@ad3758c | 1-153 | M-12 | public GeoIP context of the box's address |
| `nmb/netd/geoip.py`@ad3758c | 154-214 | N-07 | Wi-Fi regulatory domain read and set |
| `nmb/netd/guest_wifi.py`@ad3758c | 1-476 | N-08 | Wi-Fi client (station) next to the wired uplink: scan, status, connect incl. hidden, saved profiles |
| `nmb/netd/interfaces.py`@ad3758c | 1-665 | N-01 | links, carrier, MAC (factory reset of MAC), addresses, bridges, DHCP lease request, wired/USB classification, measurement interface choice, wired default-route priority, stats |
| `nmb/netd/interfaces.py`@ad3758c | 666-813 | M-12 | public and dual-stack addresses resolved in the background |
| `nmb/netd/interfaces.py`@ad3758c | 814-846 | N-01 | default gateway |
| `nmb/netd/interfaces.py`@ad3758c | 847-1070 | M-01 | ISP first hop, PPP peers, gateway ping targets switched on and off to fit the line |
| `nmb/netd/interfaces.py`@ad3758c | 1071-1181 | N-01 | dead default routes reaped, IPv4 DNS preferred, accept_ra with forwarding, router solicitation |
| `nmb/netd/ipv6_lan.py`@ad3758c | 1-62 | N-11 | delegated prefix from the ISP via dhcpcd, router adverts on the LAN (source names a specific site value; not copied) |
| `nmb/netd/kernel.py`@ad3758c | 1-364 | N-01 | kernel network tuning (sysctls, RPS/XPS masks, offloads) and its status |
| `nmb/netd/lan_topology.py`@ad3758c | 1-253 | N-06 | LAN passthrough bridge and bond (LAG) in standalone mode: validation, setup, rollback, status |
| `nmb/netd/modes.py`@ad3758c | 1-142 | N-03 | NetworkMode base, ModeA init, VLAN, WAN address, WAN2 routing, WAN failover |
| `nmb/netd/modes.py`@ad3758c | 143-184 | N-05 | pppd watch: refused login waits longer, capped pause, line problem back-off |
| `nmb/netd/modes.py`@ad3758c | 185-227 | N-11 | IPv6 on the PPPoE WAN (RA watch) |
| `nmb/netd/modes.py`@ad3758c | 228-287 | N-10 | USB as router WAN when nothing is upstream on the port |
| `nmb/netd/modes.py`@ad3758c | 288-382 | N-03 | ModeA setup, firewall application |
| `nmb/netd/modes.py`@ad3758c | 383-422 | N-11 | IPv6 LAN and UPnP checks, WAN open ports |
| `nmb/netd/modes.py`@ad3758c | 423-614 | N-03 | LAN clients renewed, LAN and NAT setup, Wi-Fi reapply (hostapd only), teardown order, status |
| `nmb/netd/modes.py`@ad3758c | 615-642 | N-11 | IPv6 LAN and port-forward status, box open ports |
| `nmb/netd/modes.py`@ad3758c | 643-812 | N-03 | ModeB transparent bridge and ModeStandalone |
| `nmb/netd/openvpn.py`@ad3758c | 1-567 | N-14 | OpenVPN server: EC PKI, server cert with TLS server usage, client names never reach PKI files, CCD allow-list, unified .ovpn, status |
| `nmb/netd/os_detector.py`@ad3758c | 1-387 | N-17 | distribution, architecture, package manager, dynamic interfaces, update check and system upgrade through helper |
| `nmb/netd/port_forward.py`@ad3758c | 1-281 | N-11 | owner's port-forward rules (router mode only), checks, nft rules, miniupnpd config for UPnP/NAT-PMP/PCP, mappings |
| `nmb/netd/proxy.py`@ad3758c | 1-502 | N-11 | HTTP CONNECT proxy: box-local targets refused, IPv6-first interleave, optional auth and TLS |
| `nmb/netd/regdomains.py`@ad3758c | 1-51 | N-07 | regulatory domain list from wireless-regdb with DFS sets, country normalisation |
| `nmb/netd/services.py`@ad3758c | 1-289 | N-03 | service supervisor: private file writes (new inode), ManagedProcess, dnsmasq, NetworkManager release/return, hostapd |
| `nmb/netd/services.py`@ad3758c | 290-312 | N-05 | pppd start, exit code, stop |
| `nmb/netd/services.py`@ad3758c | 313-383 | N-11 | miniupnpd and dhcpcd prefix delegation |
| `nmb/netd/services.py`@ad3758c | 384-391 | N-03 | stop all |
| `nmb/netd/services.py`@ad3758c | 392-471 | N-05 | kernel PPPoE plugin choice, quoting, pppd command (gives up after a few failures) |
| `nmb/netd/tailscale.py`@ad3758c | 1-259 | N-15 | Tailscale installed/active, status parse, up with flags, down, logout |
| `nmb/netd/upnp_wan.py`@ad3758c | 1-119 | M-03 | gateway WAN byte counters over UPnP for other devices' traffic; foreign host and DTD rejected (plan lists it under N-11, see proposals) |
| `nmb/netd/usb_uplink.py`@ad3758c | 1-374 | N-10 | USB tethering link: own profile without default route, daily usage, source routing, rescue of a dead default route, fallback (source names a device model; stated generically) |
| `nmb/netd/wifi.py`@ad3758c | 1-344 | N-07 | rfkill unblock, iw lookup, band HT40/flags, radio capability detection |
| `nmb/netd/wifi.py`@ad3758c | 345-401 | M-14 | Wi-Fi radio status telemetry |
| `nmb/netd/wifi.py`@ad3758c | 402-693 | N-07 | VHT/6 GHz centre frequencies, weak password rule, hostapd config with a step-down to a plain 2.4 GHz setup |
| `nmb/netd/wifi_share.py`@ad3758c | 1-131 | N-08 | sharing the Wi-Fi or USB internet on a LAN port: settings, carrier, routes, DNS |
| `nmb/netd/wifi_share.py`@ad3758c | 132-251 | N-05 | DHCP guard: DISCOVER build, offer parsing, probe for other DHCP servers |
| `nmb/netd/wifi_share.py`@ad3758c | 252-487 | N-08 | sharing state, dnsmasq, NAT for the shared subnet only, sync, leases, status |
| `nmb/netd/wifi_uplink.py`@ad3758c | 1-211 | N-08 | the box as Wi-Fi client on the owner's network as measuring link, own profile, source routing |
| `nmb/netd/wireguard.py`@ad3758c | 1-725 | N-13 | WireGuard server: keys, PSK, peers, endpoints, config, live stats, client config, QR (PNG, SVG, data URI) |
| `nmb/probe/__init__.py`@ad3758c | 1-1 | M-01 | not ported: Python package marker |
| `nmb/probe/__main__.py`@ad3758c | 1-4 | M-01 | not ported: module entry point |
| `nmb/probe/cross_traffic.py`@ad3758c | 1-178 | M-03 | other devices' traffic apart from the box's own, access-point clients, network confirmed idle, watch during a test |
| `nmb/probe/engine.py`@ad3758c | 1-300 | M-01 | probe engine: binding to the line, targets, send, replies, timeouts |
| `nmb/probe/engine.py`@ad3758c | 301-387 | M-02 | open outages settled after restart, outage state and evaluation |
| `nmb/probe/engine.py`@ad3758c | 388-480 | M-01 | sample flush, signals, run, main |
| `nmb/probe/icmp.py`@ad3758c | 1-190 | M-01 | ICMP echo build/parse, checksum, DGRAM or RAW socket, kernel timestamps, socket tuning |
| `nmb/probe/ooni.py`@ad3758c | 1-698 | M-13 | OONI probe-cli wrapper: install check, onboarding, history, measurements, anomaly breakdown, URL validation, sync/async runs, events |
| `nmb/probe/traffic_gate.py`@ad3758c | 1-300 | M-03 | interface counters, capacity, load percent, effective threshold (percent capped on fast lines), idle check |
| `nmb/report/__init__.py`@ad3758c | 1-1 | R-01 | not ported: Python package marker |
| `nmb/report/__main__.py`@ad3758c | 1-4 | R-03 | not ported: module entry point |
| `nmb/report/cli.py`@ad3758c | 1-106 | R-03 | nmb-report: window parsing, dossier, CSV export, ISP e-mail |
| `nmb/report/crash_trap.py`@ad3758c | 1-435 | C-03 | crash evidence: exception and thread hooks, sanitised locals, reports with pruning, boot-drive copy, list/read/clear, diagnostic tarball |
| `nmb/report/csv_export.py`@ad3758c | 1-429 | R-02 | RFC 4180 CSVs: ping summary (streamed), speed tests, detailed, outages, events; ranges; formula-safe cells; setup columns appended |
| `nmb/report/dossier.py`@ad3758c | 1-30 | R-03 | imports |
| `nmb/report/dossier.py`@ad3758c | 31-688 | C-06 | DOSSIER_I18N: dossier texts per language |
| `nmb/report/dossier.py`@ad3758c | 689-1674 | R-03 | verdict notes, setup section, HTML dossier generation, save |
| `nmb/report/invasive_logger.py`@ad3758c | 1-876 | R-05 | diagnostic snapshot: hardware, interfaces, systemd, DB health, dev logs, journal, threads, DNS, disks, dmesg; persisted and boot copy |
| `nmb/report/llm_bundle.py`@ad3758c | 1-411 | R-04 | CPU/thermal, PHY/ethtool, Wi-Fi PHY telemetry for the bundle |
| `nmb/report/llm_bundle.py`@ad3758c | 412-620 | M-12 | NAT and STUN telemetry, DNS interception detection, DNS matrix |
| `nmb/report/llm_bundle.py`@ad3758c | 621-2078 | R-04 | explained windows, loss pattern, config sanitising, statistics, setup lines, PMTU, ARP flux, MTR, bundle generation |
| `nmb/report/locales/es.json`@ad3758c | 1-108 | C-06 | dossier language file |
| `nmb/report/locales/it.json`@ad3758c | 1-108 | C-06 | dossier language file |
| `nmb/report/locales/pt.json`@ad3758c | 1-108 | C-06 | dossier language file |
| `nmb/report/locales/pt_br.json`@ad3758c | 1-108 | C-06 | dossier language file |
| `nmb/report/locales/ro.json`@ad3758c | 1-108 | C-06 | dossier language file |
| `nmb/report/locales/zh.json`@ad3758c | 1-108 | C-06 | dossier language file |
| `nmb/report/metrics.py`@ad3758c | 1-132 | M-02 | effective outage ends, span merge/subtract, own-action windows |
| `nmb/report/metrics.py`@ad3758c | 133-696 | R-01 | latest measured test, display rules, setup per row, report data aggregation |
| `nmb/report/smtp.py`@ad3758c | 1-217 | R-05 | ISP ticket e-mail: plain summary, setup words, HTML dossier and CSV attachments |
| `nmb/report/status_codes.py`@ad3758c | 1-75 | R-01 | one status code per speed row (1, 6, 7, 9, 10, 12, 14 only), same in CSV and dossier |
| `nmb/report/svg_charts.py`@ad3758c | 1-12 | R-03 | imports |
| `nmb/report/svg_charts.py`@ad3758c | 13-141 | C-06 | CHART_I18N chart texts |
| `nmb/report/svg_charts.py`@ad3758c | 142-549 | R-03 | latency, loss, jitter, compliance, speed charts as inline SVG |
| `nmb/selftest/__init__.py`@ad3758c | 0 (empty) | P-07 | not ported: empty Python package marker |
| `nmb/selftest/live.py`@ad3758c | 1-1605 | P-07 | nmb-selftest groups: services, internet, config, web, cli, telegram, usb, router, port forward, data cap, DHCP probe, PPPoE in a namespace, DynDNS, IPv6, mode, speed, Wi-Fi AP, collector; report pruning |
| `nmb/selftest/scenario.py`@ad3758c | 1-593 | P-07 | guided USB tethering scenario run as a service, snapshot and restore |
| `nmb/selftest/scenario_auto.py`@ad3758c | 1-187 | P-07 | guided automatic-mode scenario |
| `nmb/selftest/scenario_pppoe.py`@ad3758c | 1-288 | P-07 | guided PPPoE line scenario (up, speed, redial, groups, DynDNS, Wi-Fi client, Web UI, restore) |
| `nmb/selftest/scenario_router.py`@ad3758c | 1-164 | P-07 | guided router-mode scenario with WAN loss |
| `nmb/sniffer/__init__.py`@ad3758c | 1-4 | M-11 | not ported: Python package marker |
| `nmb/sniffer/engine.py`@ad3758c | 1-402 | M-11 | AF_PACKET capture: protocols, top talkers, DNS queries, TCP anomalies, PCAP stream |
| `nmb/speed/__init__.py`@ad3758c | 1-1 | M-04 | not ported: Python package marker |
| `nmb/speed/__main__.py`@ad3758c | 1-4 | M-04 | not ported: module entry point |
| `nmb/speed/anti_ban.py`@ad3758c | 1-295 | M-05 | Ookla ban/429 detection, server quarantine, client cooldown (doubles, capped, survives restart), rotation, schedule warning |
| `nmb/speed/cancel.py`@ad3758c | 1-74 | M-04 | cancellation scoped to the measuring thread, kills the running process |
| `nmb/speed/cloudflare.py`@ad3758c | 1-500 | M-06 | Cloudflare engine: metadata, idle and loaded latency, bound connections, throughput, rate limit, plausibility, edge address recorded |
| `nmb/speed/data_budget.py`@ad3758c | 1-47 | M-05 | daily data limit for scheduled tests (0 = off, counts only today) |
| `nmb/speed/data_cap.py`@ad3758c | 1-270 | M-05 | caps per link and per connection type, periods with reset day, usage, skip rows at most hourly |
| `nmb/speed/fastcom.py`@ad3758c | 1-473 | M-07 | Fast.com engine: token, targets (HTTP only, OCA names), unloaded/loaded latency, stable speed ignoring ramp-up |
| `nmb/speed/iperf.py`@ad3758c | 1-547 | M-09 | iPerf3 client (TCP/UDP, both directions in one row) and local server control |
| `nmb/speed/iperf_directory.py`@ad3758c | 1-390 | M-09 | public iPerf3 server lists with cache, probe, server description and row labels |
| `nmb/speed/ookla.py`@ad3758c | 1-558 | M-10 | Ookla CLI: server list, binary with pinned checksum, writable HOME, official run, speedtest-cli, bufferbloat grade; no disguised fallback |
| `nmb/speed/port_checker.py`@ad3758c | 1-214 | M-11 | port neutrality TCP checks against the 443 baseline, results stored |
| `nmb/speed/rmbt/__init__.py`@ad3758c | 1-1 | M-08 | not ported: Python package marker |
| `nmb/speed/rmbt/__main__.py`@ad3758c | 1-242 | M-08 | Python RMBT client phases and entry |
| `nmb/speed/rmbt/connection.py`@ad3758c | 1-264 | M-08 | RMBT connection bound to the chosen link |
| `nmb/speed/rmbt/control.py`@ad3758c | 1-160 | M-08 | RMBT control server requests (settings, testRequest, result), bound handlers, no unbound retry |
| `nmb/speed/rmbt/pretest.py`@ad3758c | 1-112 | M-08 | RMBT pretest, thread count |
| `nmb/speed/rmbt/tests.py`@ad3758c | 1-221 | M-08 | RMBT ping, download, upload (C loop when available) |
| `nmb/speed/rmbt/uuid_store.py`@ad3758c | 1-37 | M-08 | RMBT client UUID persistence |
| `nmb/speed/rtr.py`@ad3758c | 1-410 | M-08 | RTR engine: native client lookup, output tail on failure, env, client IPs, bound run, result |
| `nmb/speed/runner.py`@ad3758c | 1-1343 | M-04 | measurement cycle: one test at a time box-wide (lock), engines due by interval, skip and aborted rows, setup capture, iPerf3 server selection and rest, USB guard, CPU sampling; 1343 lines |
| `nmb/speed/scheduler.py`@ad3758c | 1-517 | M-04 | cron and interval parser, next run, retry before the next test, in-process scheduler (runs inside nmb-web) |
| `nmb/speed/traceroute.py`@ad3758c | 1-678 | M-11 | traceroute/mtr/tracepath and ping-TTL fallback, hop classes and enrichment, streaming, stop |
| `nmb/tui/__init__.py`@ad3758c | 1-1 | S-10 | not ported: Python package marker |
| `nmb/tui/app.py`@ad3758c | 1-503 | S-10 | interactive curses menu; its colour scheme is an add-on matter (P-09) per the plain-CLI decision |
| `nmb/tui/catalog.py`@ad3758c | 1-422 | S-10 | menu tree read from the Web UI page, actions, value coercion, request bodies |
| `nmb/tui/cli.py`@ad3758c | 1-215 | S-10 | JSON commands for scripts, menu JSON, setup text, latest results, update |
| `nmb/tui/client.py`@ad3758c | 1-67 | S-10 | client of the box's own REST API |
| `nmb/tui/logo.py`@ad3758c | 1-23 | S-10 | product logo text |
| `nmb/tui/logs.py`@ad3758c | 1-68 | S-10 | logs as text lines |
| `nmb/tui/render.py`@ad3758c | 1-65 | S-10 | API answers to plain text lines, secrets hidden |
| `nmb/web/__init__.py`@ad3758c | 1-1 | S-01 | not ported: Python package marker |
| `nmb/web/__main__.py`@ad3758c | 1-4 | S-01 | not ported: module entry point |
| `nmb/web/assets/dashboard.css`@ad3758c | 1-868 | U-01 | base look, layout, components |
| `nmb/web/assets/dashboard.css`@ad3758c | 869-1005, 1067-2165 | P-09 | themes default, PrTD, ProudRide, AsusWRT, AsusCamouflage, TP-Gold, MountainHaGo, nDefault: add-on themes, not in the AMOLED-only core |
| `nmb/web/assets/dashboard.css`@ad3758c | 1006-1066 | U-01 | AMOLED pitch-black theme: the native look (owner decision) |
| `nmb/web/assets/dashboard.css`@ad3758c | 2166-2812 | U-01 | responsive/mobile rules, focus, donation and Ookla-terms styles (a few vendor-theme lines inside, to drop or move to P-09) |
| `nmb/web/assets/dashboard.html`@ad3758c | 1-109, 5221-5239 | U-01 | head, translation and donation injection, header (language and theme menus), tab navigation, toast, script slot, donation footer and dialog |
| `nmb/web/assets/dashboard.html`@ad3758c | 110-185 | U-02 (Dashboard page) | Dashboard (home) page |
| `nmb/web/assets/dashboard.html`@ad3758c | 186-226 | U-03 (Ping telemetry page) | Ping telemetry page |
| `nmb/web/assets/dashboard.html`@ad3758c | 227-335 | U-04 (Speed results page) | Speed results page |
| `nmb/web/assets/dashboard.html`@ad3758c | 336-522 | U-05 (iPerf3 page) | iPerf3 page and public server dialog |
| `nmb/web/assets/dashboard.html`@ad3758c | 523-641 | U-06 (Traceroute page) | Traceroute page |
| `nmb/web/assets/dashboard.html`@ad3758c | 642-1052 | U-07 (DNS blocker page) | DNS blocker page and sub-tabs |
| `nmb/web/assets/dashboard.html`@ad3758c | 1053-1259 | U-08 (OONI page) | OONI page and details dialog |
| `nmb/web/assets/dashboard.html`@ad3758c | 1260-2086 | U-09 (Remote access and VPN page) | Remote access and VPN page (hub, guide, DynDNS, WireGuard, OpenVPN, proxy, Tailscale, QR dialog) |
| `nmb/web/assets/dashboard.html`@ad3758c | 2087-2167 | U-10 (Network explorer page) | Network explorer page |
| `nmb/web/assets/dashboard.html`@ad3758c | 2168-2200 | U-11 (Port neutrality page) | Port neutrality page |
| `nmb/web/assets/dashboard.html`@ad3758c | 2201-2377 | U-12 (Logs page) | Developer and system logs page |
| `nmb/web/assets/dashboard.html`@ad3758c | 2378-2437 | U-13 (Packet capture page) | Packet capture page |
| `nmb/web/assets/dashboard.html`@ad3758c | 2438-2813 | U-14 (Measurement and servers page) | Measurement and servers page (schedule, Ookla terms, links, data caps, storage) |
| `nmb/web/assets/dashboard.html`@ad3758c | 2814-2879 | U-15 (Bandwidth monitor page) | Bandwidth monitor page |
| `nmb/web/assets/dashboard.html`@ad3758c | 2880-3567, 3704-3780, 4224-4423, 4802-4810 | U-16 (Router page) | Router page: WAN (PPPoE, static, MAC, DNS), LAN, port forwarding, Wi-Fi AP, guest Wi-Fi client, operation mode, ISP profile, ports and drivers, save bar |
| `nmb/web/assets/dashboard.html`@ad3758c | 3568-3703 | U-14 (Measurement and servers page) | measurement schedule and user-agent sub-tab of the router page; moved to the measurement page in the native UI |
| `nmb/web/assets/dashboard.html`@ad3758c | 3781-3862, 3952-4223, 4424-4776 | U-17 (System page) | System page: kernel and hardware, memory profile, SSH, credentials and public results, collector, updates, drive config, factory reset, diagnostics status, backup and purge, self-healing |
| `nmb/web/assets/dashboard.html`@ad3758c | 3863-3951 | P-09 | theme chooser card |
| `nmb/web/assets/dashboard.html`@ad3758c | 4777-4801 | U-18 (Telegram page) | Telegram sub-tab of the router page |
| `nmb/web/assets/dashboard.html`@ad3758c | 4811-4971 | U-18 (Telegram page) | Telegram page |
| `nmb/web/assets/dashboard.html`@ad3758c | 4972-5000 | U-19 (AI diagnostics page) | AI diagnostics page |
| `nmb/web/assets/dashboard.html`@ad3758c | 5001-5061 | U-20 (Terminal page) | CLI terminal page |
| `nmb/web/assets/dashboard.html`@ad3758c | 5062-5220 | U-21 (Tasks and load page) | Tasks and load page |
| `nmb/web/assets/dashboard.js`@ad3758c | 1-40 | U-01 | state and translation loading |
| `nmb/web/assets/dashboard.js`@ad3758c | 41-78 | U-04 (Speed results page) | speed error and Ookla warning texts |
| `nmb/web/assets/dashboard.js`@ad3758c | 79-116 | U-16 (Router page) | Wi-Fi country list |
| `nmb/web/assets/dashboard.js`@ad3758c | 117-176 | U-01 | donation list and dialog |
| `nmb/web/assets/dashboard.js`@ad3758c | 177-211 | U-14 (Measurement and servers page) | Ookla terms |
| `nmb/web/assets/dashboard.js`@ad3758c | 212-236 | U-17 (System page) | SSH keys only |
| `nmb/web/assets/dashboard.js`@ad3758c | 237-300 | U-01 | applyTranslations |
| `nmb/web/assets/dashboard.js`@ad3758c | 301-350 | P-09 | theme switching |
| `nmb/web/assets/dashboard.js`@ad3758c | 351-403 | U-01 | language switch, escaping, URL sanitising, toast |
| `nmb/web/assets/dashboard.js`@ad3758c | 404-449 | U-17 (System page) | memory profile |
| `nmb/web/assets/dashboard.js`@ad3758c | 450-530 | U-01 | tab switching |
| `nmb/web/assets/dashboard.js`@ad3758c | 531-623 | U-16 (Router page) | LAN topology |
| `nmb/web/assets/dashboard.js`@ad3758c | 624-668 | U-17 (System page) | data purge |
| `nmb/web/assets/dashboard.js`@ad3758c | 669-678 | U-01 | accessibility |
| `nmb/web/assets/dashboard.js`@ad3758c | 679-1015 | U-16 (Router page) | interface matrix, WAN2, drivers |
| `nmb/web/assets/dashboard.js`@ad3758c | 1016-2061 | U-07 (DNS blocker page) | DNS blocker |
| `nmb/web/assets/dashboard.js`@ad3758c | 2062-3131 | U-02 (Dashboard page) | engine cards, setup line, refreshStatus (live status of all pages) |
| `nmb/web/assets/dashboard.js`@ad3758c | 3132-3517 | U-04 (Speed results page) | speed chart and tooltip |
| `nmb/web/assets/dashboard.js`@ad3758c | 3518-3607 | U-02 (Dashboard page) | gauge, speedometer, start/stop speed test |
| `nmb/web/assets/dashboard.js`@ad3758c | 3608-3618 | U-11 (Port neutrality page) | port check |
| `nmb/web/assets/dashboard.js`@ad3758c | 3619-3733 | U-14 (Measurement and servers page) | server catalog and default server |
| `nmb/web/assets/dashboard.js`@ad3758c | 3734-4082 | U-12 (Logs page) | log filters, syslog, developer log, log level, purge |
| `nmb/web/assets/dashboard.js`@ad3758c | 4083-4162 | U-13 (Packet capture page) | sniffer |
| `nmb/web/assets/dashboard.js`@ad3758c | 4163-4611 | U-16 (Router page) | MAC, DNS, PPPoE, router sub-tabs, WAN protocol, Wi-Fi AP form |
| `nmb/web/assets/dashboard.js`@ad3758c | 4612-4655 | U-02 (Dashboard page) | home server choice, copy script |
| `nmb/web/assets/dashboard.js`@ad3758c | 4656-4989 | U-14 (Measurement and servers page) | schedule form, traffic gate, user agent, save schedule |
| `nmb/web/assets/dashboard.js`@ad3758c | 4990-5021 | U-02 (Dashboard page) | quick schedule on home |
| `nmb/web/assets/dashboard.js`@ad3758c | 5022-5242 | U-16 (Router page) | load config, save router settings |
| `nmb/web/assets/dashboard.js`@ad3758c | 5243-5837 | U-17 (System page) | reload, reboot, backups, factory reset, drive config, healing, updates |
| `nmb/web/assets/dashboard.js`@ad3758c | 5838-5982 | U-20 (Terminal page) | terminal |
| `nmb/web/assets/dashboard.js`@ad3758c | 5983-6156 | U-17 (System page) | diagnostic snapshot, anti-ban and wired status, crashes, Ookla quarantine |
| `nmb/web/assets/dashboard.js`@ad3758c | 6157-6653 | U-05 (iPerf3 page) | iPerf3 |
| `nmb/web/assets/dashboard.js`@ad3758c | 6654-6929 | U-06 (Traceroute page) | traceroute |
| `nmb/web/assets/dashboard.js`@ad3758c | 6930-7437 | U-08 (OONI page) | OONI |
| `nmb/web/assets/dashboard.js`@ad3758c | 7438-7545 | U-10 (Network explorer page) | network devices |
| `nmb/web/assets/dashboard.js`@ad3758c | 7546-7905 | U-16 (Router page) | guest Wi-Fi client |
| `nmb/web/assets/dashboard.js`@ad3758c | 7906-7945 | U-19 (AI diagnostics page) | LLM bundle |
| `nmb/web/assets/dashboard.js`@ad3758c | 7946-7986 | U-17 (System page) | admin credentials, public results |
| `nmb/web/assets/dashboard.js`@ad3758c | 7987-8231 | U-14 (Measurement and servers page) | Wi-Fi uplink, Wi-Fi share, USB uplink, storage |
| `nmb/web/assets/dashboard.js`@ad3758c | 8232-8299 | U-17 (System page) | collector |
| `nmb/web/assets/dashboard.js`@ad3758c | 8300-8331 | U-01 | logout |
| `nmb/web/assets/dashboard.js`@ad3758c | 8332-8499 | U-17 (System page) | system info, OS updates, kernel tuning |
| `nmb/web/assets/dashboard.js`@ad3758c | 8500-10072 | U-09 (Remote access and VPN page) | remote access hub, port guide, DynDNS, WireGuard, OpenVPN, proxy, Tailscale |
| `nmb/web/assets/dashboard.js`@ad3758c | 10073-10433 | U-18 (Telegram page) | Telegram |
| `nmb/web/assets/dashboard.js`@ad3758c | 10434-10700 | U-21 (Tasks and load page) | tasks |
| `nmb/web/assets/dashboard.js`@ad3758c | 10701-10774 | U-16 (Router page) | port forwarding |
| `nmb/web/assets/dashboard.js`@ad3758c | 10775-10829 | U-14 (Measurement and servers page) | data caps |
| `nmb/web/assets/logout.html`@ad3758c | 1-250 | U-01 | logout page |
| `nmb/web/i18n.py`@ad3758c | 1-7171 | C-06 | Web UI translations (de, en, fr, hu, uk dictionaries, extras, corrections, storage and USB texts), 11-language list, fallback to English; data of 7171 lines |
| `nmb/web/locales/es.json`@ad3758c | 1-2447 | C-06 | Web UI language file |
| `nmb/web/locales/it.json`@ad3758c | 1-2447 | C-06 | Web UI language file |
| `nmb/web/locales/pt.json`@ad3758c | 1-2447 | C-06 | Web UI language file |
| `nmb/web/locales/pt_br.json`@ad3758c | 1-2447 | C-06 | Web UI language file |
| `nmb/web/locales/ro.json`@ad3758c | 1-2447 | C-06 | Web UI language file |
| `nmb/web/locales/ui_core.json`@ad3758c | 1-4912 | C-06 | Web UI core keys (en, hu, de, fr, uk) |
| `nmb/web/locales/zh.json`@ad3758c | 1-2447 | C-06 | Web UI language file |
| `nmb/web/results_page.py`@ad3758c | 1-165 | S-09 (Exports and public results API) | password-free read-only results for LAN clients only when public_results is on; escaped; donation footer |
| `nmb/web/server.py`@ad3758c | 1-147, 234-442, 1399-1494, 4386-4612 | S-01 | transport: login rate limit, auth (Basic), JSON/HTML senders, static files, index, logout, CSRF Origin check, body size limits, backup upload body, control-character refusal, server and daemon (also hosts the scheduler) |
| `nmb/web/server.py`@ad3758c | 148-233, 443-447, 3442-3774, 4368-4385 | S-02 (Status API) | status: /api/status payload, latest results and current setup, page language, version, hardware model |
| `nmb/web/server.py`@ad3758c | 448-450, 459-490, 618-630, 780-886, 908-911, 1332-1337, 1495-1552, 1571-1792, 2057-2187, 2208-2293, 2433-2437, 2622-2635, 3420-3426, 3775-3786, 4288-4317 | S-03 (Speed and measuring API) | speed and measuring: speed test run/stop, Ookla terms, servers, engine and schedule, user agent, port check, sniffer, iPerf3, traceroute, OONI, scheduler, anti-ban, data cap |
| `nmb/web/server.py`@ad3758c | 451-458, 602-617, 631-667, 724-750, 887-898, 922-926, 1074-1153, 1200-1304, 1311-1331, 1342-1346, 1360-1380, 1553-1570, 1810-1814, 1947-2056, 2418-2432, 2481-2565, 2697-2728, 3083-3419, 3427-3433, 3787-4208, 4342-4367 | S-08 (System and settings API) | system and settings: /api/config read/write (3787-4208), SSH keys only, drive config, i18n, theme, info, tasks/kill, memory, heal, updates, storage, collector, Telegram, backup/restore/factory reset, dev logs, syslog, crashes, invasive log, purge, reboot, credentials, terminal exec, daemon reload |
| `nmb/web/server.py`@ad3758c | 491-596, 751-769, 1381-1398 | S-09 (Exports and public results API) | exports: public results page, dossier, CSV routes, LLM bundle and its download, diagnostic tarball |
| `nmb/web/server.py`@ad3758c | 597-601, 668-680, 770-779, 899-907, 1190-1199, 1305-1310, 1338-1341, 1347-1359, 1793-1809, 1874-1946, 2188-2207, 2294-2319, 2414-2417, 2438-2480, 2636-2696, 3434-3441, 4209-4287, 4318-4341 | S-04 (Network API) | network: DNS, interfaces/ports and roles, drivers, devices, USB uplink, port forward, MAC, kernel tuning, LAN topology, preferred interface, mode, wired priority |
| `nmb/web/server.py`@ad3758c | 681-723, 2320-2413 | S-07 (DNS blocker API) | DNS blocker: status, charts, query log, lists, adlists, local DNS, regex, gravity |
| `nmb/web/server.py`@ad3758c | 912-921, 1154-1189, 1815-1873, 2566-2621 | S-05 (Wi-Fi API) | Wi-Fi: capabilities, status, guest client (scan, saved, connect, disconnect), Wi-Fi share, Wi-Fi uplink |
| `nmb/web/server.py`@ad3758c | 927-1073, 2729-3082 | S-06 (VPN and remote-access API) | VPN and remote access: DynDNS, WireGuard (peers, config, QR), OpenVPN, proxy, Tailscale |
| `nmb/web/static/themes/asus_bottom_bg.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/asus_middown_bg.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/asus_title.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/asus_title_bg.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/hago.jpg`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/ndefault_bg_center.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/ndefault_bg_line.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/ndefault_box_header.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/ndefault_header_bg.jpg`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/ndefault_info_bar.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/ndefault_main_top.jpg`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/ndefault_menu_bar.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/ndefault_top_strip.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/tplink_blue.jpg`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/tplink_top1_1.jpg`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/tplink_top1_2.jpg`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/tplink_top_bg.jpg`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/tuf_camo_bg.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/tuf_logo.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/static/themes/tuf_title_bg.png`@ad3758c | binary | P-09 | theme image (binary) |
| `nmb/web/templates.py`@ad3758c | 1-55 | U-01 | page assembly: CSS/JS/translations/theme/donation injected into the HTML |
| `os/90-nmb-usb-import.rules`@ad3758c | 1-2 | P-03 | udev: a USB file system starts the import unit |
| `os/90-nmb.pkla`@ad3758c | 1-15 | P-06 | polkit (old format) grants for the nmb user; cannot limit by unit name |
| `os/90-nmb.rules`@ad3758c | 1-20 | P-06 | polkit grants: nmb may manage only its own and named units, NetworkManager actions, reboot |
| `os/99-nmb-performance.conf`@ad3758c | 1-92 | N-01 | kernel sysctls (buffers, BBR+fq, ping group, forwarding, bridge netfilter off, rp_filter loose, VM); installed by setup.sh |
| `os/net-detect.sh`@ad3758c | 1-52 | S-10 | not ported: stand-alone link print-out, installed but called by no code; nmb-cli status covers it |
| `os/nic-tune.sh`@ad3758c | 1-88 | N-01 | NIC ring buffers, offloads, coalescing, EEE off, txqueuelen by RAM, RPS/XPS, CPU governor (run by nmb-tune.service) |
| `os/nmb-collector.service`@ad3758c | 1-9, 11-12, 18-20 | S-13 | unit of the central collector on its own server |
| `os/nmb-collector.service`@ad3758c | 10, 13-17 | P-04 | collector sandbox: own user, NoNewPrivileges, ProtectSystem strict, one writable path |
| `os/nmb-guard.nft`@ad3758c | 1-12 | N-02 | public sources dropped for ports 53, 80, 5201 on IPv4 and IPv6; private, link-local, CGNAT/Tailscale allowed |
| `os/nmb-wait-online.conf`@ad3758c | 1-12 | P-10 | network-online means one active connection with an address, not NetworkManager's full start-up (source names a specific site value; not copied) |
| `os/ooni-resolv.conf`@ad3758c | 1-4 | M-13 | public resolvers for the OONI run only, so a LAN DNS filter is not reported as ISP blocking |
| `os/openrc/nmb-db`@ad3758c | 1-19 | P-10 | OpenRC service for an OpenRC base (kept only if P-11 picks one) |
| `os/openrc/nmb-netd`@ad3758c | 1-19 | P-10 | OpenRC service for an OpenRC base (kept only if P-11 picks one) |
| `os/openrc/nmb-probe`@ad3758c | 1-20 | P-10 | OpenRC service for an OpenRC base (kept only if P-11 picks one) |
| `os/openrc/nmb-telegram`@ad3758c | 1-20 | P-10 | OpenRC service for an OpenRC base (kept only if P-11 picks one) |
| `os/openrc/nmb-web`@ad3758c | 1-22 | P-10 | OpenRC service for an OpenRC base (kept only if P-11 picks one) |
| `os/storage-setup.sh`@ad3758c | 1-109 | P-11 | OS baseline: nmb user, folders and modes, tmpfs for /tmp, /var/log, live DB, root ext4 noatime/commit, volatile journal sized by RAM, swap off, writeback sysctls, chatty timers off |
| `setup.sh`@ad3758c | 1-392 | P-11 | installer: Ookla with pinned checksum, OONI binary, room on /var/log, OS and init detection, packages, RMBT build, user, sudoers, storage, sysctls, files to /opt/nmb root-owned, installed commit, config merge, log size fix |
| `setup.sh`@ad3758c | 393-454 | P-03 | unique Wi-Fi password per box, restore from secrets copy, admin password kept and not printed, first-boot credentials file, boot-partition copy only for the default or a generated password |
| `setup.sh`@ad3758c | 455-664 | P-11 | low-memory profile, group, symlinks (pivpn, pihole), DB init, Ookla home, /run/nmb and root-only files, units, udev, polkit, OpenRC, summary, donation list |
| `systemd/nmb-adblock-update.service`@ad3758c | 1-6, 9-12 | P-10 | weekly blocklist update run |
| `systemd/nmb-adblock-update.service`@ad3758c | 7-8 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-adblock-update.timer`@ad3758c | 1-11 | P-10 | Sunday 03:00 with random delay, persistent |
| `systemd/nmb-adblock.service`@ad3758c | 1-6, 9-16, 23-26 | P-10 | blocker daemon, restart always |
| `systemd/nmb-adblock.service`@ad3758c | 7-8, 17-22 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-db.service`@ad3758c | 1-13, 16-22 | P-10 | restore live DB at boot, snapshot at shutdown while file systems are mounted |
| `systemd/nmb-db.service`@ad3758c | 14-15 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-guard.service`@ad3758c | 1-13 | P-10 | loads os/nmb-guard.nft before the network and the web server |
| `systemd/nmb-heal.service`@ad3758c | 1-7, 11-16 | P-10 | self-healing run |
| `systemd/nmb-heal.service`@ad3758c | 8-10 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-heal.timer`@ad3758c | 1-12 | P-10 | 3 min after boot, then every 15 min |
| `systemd/nmb-invasive-log.service`@ad3758c | 1-8, 11-17 | P-10 | diagnostic snapshot, idle I/O, output not copied to the journal |
| `systemd/nmb-invasive-log.service`@ad3758c | 9-10 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-invasive-log.timer`@ad3758c | 1-11 | P-10 | every 3 min |
| `systemd/nmb-iperf3.service`@ad3758c | 1-6, 9-18 | P-10 | local iPerf3 server on 5201 |
| `systemd/nmb-iperf3.service`@ad3758c | 7-8 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-maintain.service`@ad3758c | 1-7, 10-12 | P-10 | DB rollup, prune, optimise |
| `systemd/nmb-maintain.service`@ad3758c | 8-9 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-maintain.timer`@ad3758c | 1-10 | P-10 | every 5 min |
| `systemd/nmb-netd.service`@ad3758c | 1-9, 12-21, 33-35 | P-10 | netd; 45 s stop time so pppd can send PADT; reload by HUP |
| `systemd/nmb-netd.service`@ad3758c | 10-11, 22-32 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-ooni.service`@ad3758c | 1-6, 9-14, 17 | P-10 | weekly OONI websites run only when enabled; public resolvers bound in |
| `systemd/nmb-ooni.service`@ad3758c | 7-8, 15-16 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-ooni.timer`@ad3758c | 1-11 | P-10 | Sunday 03:30, persistent |
| `systemd/nmb-probe.service`@ad3758c | 1-8, 11-15, 27-29 | P-10 | probe daemon |
| `systemd/nmb-probe.service`@ad3758c | 9-10, 16-26 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-report.service`@ad3758c | 1-7, 10-14 | P-10 | daily dossier |
| `systemd/nmb-report.service`@ad3758c | 8-9, 15-19 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-report.timer`@ad3758c | 1-11 | P-10 | 00:05 UTC daily |
| `systemd/nmb-snapshot.service`@ad3758c | 1-7, 10-13 | P-10 | live DB snapshot to SD |
| `systemd/nmb-snapshot.service`@ad3758c | 8-9 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-snapshot.timer`@ad3758c | 1-10 | P-10 | every 15 min |
| `systemd/nmb-speed.service`@ad3758c | 1-7, 10-14 | P-10 | scheduled speed test run |
| `systemd/nmb-speed.service`@ad3758c | 8-9, 15-19 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-telegram.service`@ad3758c | 1-10, 13-18, 26-28 | P-10 | bot; not after network-online (router mode boot delay) |
| `systemd/nmb-telegram.service`@ad3758c | 11-12, 19-25 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb-tune.service`@ad3758c | 1-12 | P-10 | NIC tuning before the network |
| `systemd/nmb-usb-import@.service`@ad3758c | 1-9 | P-10 | USB settings import per partition |
| `systemd/nmb-web.service`@ad3758c | 1-7, 15-19, 25-27 | P-10 | web server on port 80 |
| `systemd/nmb-web.service`@ad3758c | 8-14, 20-24 | P-04 | user, capabilities and sandbox directives of this unit |
| `systemd/nmb.target`@ad3758c | 1-7 | P-10 | the whole stack |
| `tests/_box_isolation.py`@ad3758c | 1-14 | F-06 | not ported: Python test helper; its rule (tests never touch the box's shared files) is a test-bed requirement |
| `tests/test_adblock.py`@ad3758c | 1-291 | N-12 | tests of the code of this package |
| `tests/test_adblock_hardening.py`@ad3758c | 1-200 | N-12 | tests of the code of this package |
| `tests/test_auto_mode.py`@ad3758c | 1-159 | N-04 | tests of the code of this package |
| `tests/test_auto_pppoe_once.py`@ad3758c | 1-54 | N-04 | tests of the code of this package |
| `tests/test_backup_send_encrypted.py`@ad3758c | 1-117 | P-03 | bot /backup send encrypted by default (command side in S-11) |
| `tests/test_bot_authorization.py`@ad3758c | 1-156 | S-11 | tests of the code of this package |
| `tests/test_bot_client_backoff.py`@ad3758c | 1-52 | S-11 | tests of the code of this package |
| `tests/test_bot_dispatch.py`@ad3758c | 1-54 | S-11 | tests of the code of this package |
| `tests/test_bot_locales.py`@ad3758c | 1-61 | C-06 | tests of the code of this package |
| `tests/test_client_address_round2.py`@ad3758c | 1-51 | C-05 | address rule |
| `tests/test_client_address_round2.py`@ad3758c | 52-72 | M-09 | iPerf3 row |
| `tests/test_client_address_round2.py`@ad3758c | 73-102 | R-01 | LAN test not in the line's figures, CSV source |
| `tests/test_client_address_round2.py`@ad3758c | 103-114 | S-11 | bot status codes match the dossier |
| `tests/test_client_address_round2.py`@ad3758c | 115-172 | S-13 | collector token and corrections |
| `tests/test_cloudflare_engine.py`@ad3758c | 1-46 | M-06 | tests of the code of this package |
| `tests/test_collector.py`@ad3758c | 1-195 | S-13 | client and server together; server is the code mostly checked |
| `tests/test_config_merge.py`@ad3758c | 1-83 | C-01 | tests of the code of this package |
| `tests/test_cross_traffic.py`@ad3758c | 1-106 | M-03 | threshold, monitor, strictness |
| `tests/test_cross_traffic.py`@ad3758c | 107-162 | M-04 | cancel, retry, abort path |
| `tests/test_cross_traffic.py`@ad3758c | 163-181 | M-03 | UPnP WAN counters |
| `tests/test_csv_exports.py`@ad3758c | 1-96 | R-02 | ranges, columns, formula safety, events and ping CSV |
| `tests/test_csv_exports.py`@ad3758c | 97-117 | M-04 | runner row starts |
| `tests/test_csv_exports.py`@ad3758c | 118-128 | C-05 | installed version |
| `tests/test_data_cap.py`@ad3758c | 1-179 | M-05 | periods, links, counting, settings, runner stop |
| `tests/test_data_cap.py`@ad3758c | 180-206 | S-11 | bot alert once and resolved |
| `tests/test_dns_fallback.py`@ad3758c | 1-88 | M-12 | tests of the code of this package |
| `tests/test_donate.py`@ad3758c | 1-122 | P-08 | tests of the code of this package |
| `tests/test_dossier.py`@ad3758c | 1-225 | R-03 | tests of the code of this package |
| `tests/test_driver_cli.py`@ad3758c | 1-58 | N-17 | tests of the code of this package |
| `tests/test_duckdns.py`@ad3758c | 1-150 | N-16 | DynDNS client |
| `tests/test_duckdns.py`@ad3758c | 151-227 | S-11 | bot /dyndns |
| `tests/test_duckdns.py`@ad3758c | 228-301 | N-16 | state across processes |
| `tests/test_engine_errors.py`@ad3758c | 1-34 | M-04 | engine error is a failed row |
| `tests/test_engine_errors.py`@ad3758c | 35-50 | M-08 | RMBT control bound |
| `tests/test_engine_errors.py`@ad3758c | 51-91 | M-04 | router WAN down |
| `tests/test_engine_errors.py`@ad3758c | 92-134 | N-16 | DynDNS IPv4 detection |
| `tests/test_engine_errors.py`@ad3758c | 135-180 | M-04 | USB guard, iPerf3 rest warning |
| `tests/test_failover_and_cleanup.py`@ad3758c | 1-44 | N-03 | WAN failover |
| `tests/test_failover_and_cleanup.py`@ad3758c | 45-68 | C-08 | impossible test cleanup |
| `tests/test_fastcom.py`@ad3758c | 1-66 | M-07 | Fast.com |
| `tests/test_fastcom.py`@ad3758c | 67-91 | M-09 | iPerf3 row labels |
| `tests/test_fastcom.py`@ad3758c | 92-114 | S-02 (Status API) | home page detail |
| `tests/test_feature_sweep_fixes.py`@ad3758c | 1-58 | N-11 | proxy targets |
| `tests/test_feature_sweep_fixes.py`@ad3758c | 59-82 | N-08 | guest Wi-Fi state |
| `tests/test_feature_sweep_fixes.py`@ad3758c | 83-98 | N-14 | OpenVPN cert |
| `tests/test_feature_sweep_fixes.py`@ad3758c | 99-105 | N-07 | hostapd fallback |
| `tests/test_feature_sweep_fixes.py`@ad3758c | 106-138 | P-03 | restore with stale WAL, group-readable config |
| `tests/test_feature_sweep_fixes.py`@ad3758c | 139-151 | C-05 | version |
| `tests/test_firewall_every_mode.py`@ad3758c | 1-194 | N-02 | tests of the code of this package |
| `tests/test_first_run_fixes.py`@ad3758c | 1-35 | C-01 | config resilience |
| `tests/test_first_run_fixes.py`@ad3758c | 36-41 | N-01 | IPv6 route check import |
| `tests/test_first_run_fixes.py`@ad3758c | 42-61 | P-03 | factory reset safety, unique Wi-Fi password |
| `tests/test_first_run_fixes.py`@ad3758c | 62-72 | P-01 | self-heal callers, update verification |
| `tests/test_first_run_fixes.py`@ad3758c | 73-82 | M-03 | blind box never idle |
| `tests/test_first_run_fixes.py`@ad3758c | 83-93 | R-01 | skipped runs shown |
| `tests/test_first_run_fixes.py`@ad3758c | 94-108 | M-09 | iPerf3 port check |
| `tests/test_first_run_fixes.py`@ad3758c | 109-139 | P-06 | security regressions (paths, sudoers, helper) |
| `tests/test_first_run_fixes.py`@ad3758c | 140-156 | C-01 | boot copy secrets |
| `tests/test_first_run_fixes.py`@ad3758c | 157-171 | M-05 | data budget |
| `tests/test_first_run_fixes.py`@ad3758c | 172-182 | N-07 | hostapd step-down, reload via systemd |
| `tests/test_first_run_fixes.py`@ad3758c | 183-194 | P-04 | service hardening, public guard |
| `tests/test_first_run_fixes.py`@ad3758c | 195-204 | C-01 | lost config |
| `tests/test_first_run_fixes.py`@ad3758c | 205-241 | N-02 | DMZ rules (N-02); its other three tests belong to M-12 (DNS interception), C-06 (alert texts in all languages) and M-11 (more ports checked): the analyst of N-02 hands them over |
| `tests/test_first_run_fixes.py`@ad3758c | 242-249 | P-11 | Ookla checksum shared |
| `tests/test_first_run_fixes.py`@ad3758c | 250-258 | S-01 | POST body must be an object |
| `tests/test_github_updater.py`@ad3758c | 1-176 | P-02 | tests of the code of this package |
| `tests/test_headless_features.py`@ad3758c | 1-63 | C-05 | user agent (setUp shared) |
| `tests/test_headless_features.py`@ad3758c | 64-94 | N-01 | wired prioritisation |
| `tests/test_headless_features.py`@ad3758c | 95-137 | M-05 | anti-ban and quarantine |
| `tests/test_headless_features.py`@ad3758c | 138-152 | R-05 | invasive logger |
| `tests/test_headless_features.py`@ad3758c | 153-221 | C-03 | crash trap and bundle |
| `tests/test_headless_features.py`@ad3758c | 222-247 | N-01 | wired default route |
| `tests/test_headless_features.py`@ad3758c | 248-270 | S-01 | web default port, unprivileged fallback |
| `tests/test_headless_features.py`@ad3758c | 271-310 | N-08 | guest Wi-Fi saved and hidden |
| `tests/test_headless_features.py`@ad3758c | 311-333 | M-08 | RTR dual-stack binding |
| `tests/test_headless_features.py`@ad3758c | 334-427 | P-10 | CLI tools and components present |
| `tests/test_heal.py`@ad3758c | 1-210 | P-01 | tests of the code of this package |
| `tests/test_helper_allowlists.py`@ad3758c | 1-118 | P-06 | tests of the code of this package |
| `tests/test_i18n_locales.py`@ad3758c | 1-68 | C-06 | tests of the code of this package |
| `tests/test_iperf3_rest.py`@ad3758c | 1-52 | M-04 | runner rests a public iPerf3 server that always fails |
| `tests/test_ipv6_lan.py`@ad3758c | 1-104 | N-11 | tests of the code of this package |
| `tests/test_lan_topology.py`@ad3758c | 1-115 | N-06 | tests of the code of this package |
| `tests/test_llm_bundle.py`@ad3758c | 1-186 | R-04 | tests of the code of this package |
| `tests/test_llm_bundle_honesty.py`@ad3758c | 1-122 | R-04 | tests of the code of this package |
| `tests/test_llm_bundle_masking.py`@ad3758c | 1-35 | R-04 | tests of the code of this package |
| `tests/test_log_budget.py`@ad3758c | 1-57 | C-03 | rotation fit, log purge |
| `tests/test_log_budget.py`@ad3758c | 58-66 | C-09 | kill signal allow-list |
| `tests/test_log_budget.py`@ad3758c | 67-88 | C-04 | bundle config redacted |
| `tests/test_measure_lock.py`@ad3758c | 1-42 | M-04 | tests of the code of this package |
| `tests/test_measurement_setup.py`@ad3758c | 1-208 | C-11 | classify, finish, history, backfill |
| `tests/test_measurement_setup.py`@ad3758c | 209-236 | M-04 | runner stores setup |
| `tests/test_measurement_setup.py`@ad3758c | 237-254 | C-11 | netd link |
| `tests/test_measurement_setup.py`@ad3758c | 255-267 | M-02 | outage setup |
| `tests/test_measurement_setup.py`@ad3758c | 268-291 | N-03 | netd records once |
| `tests/test_measurement_setup.py`@ad3758c | 292-301 | S-13 | collector codes |
| `tests/test_measurement_setup.py`@ad3758c | 302-338 | C-11 | texts |
| `tests/test_mixed_engines.py`@ad3758c | 1-76 | M-04 | tests of the code of this package |
| `tests/test_no_made_up_values.py`@ad3758c | 1-32 | C-05 | public addresses only |
| `tests/test_no_made_up_values.py`@ad3758c | 33-65 | M-10 | bufferbloat, Ookla |
| `tests/test_no_made_up_values.py`@ad3758c | 66-86 | M-06 | Cloudflare |
| `tests/test_no_made_up_values.py`@ad3758c | 87-116 | M-08 | RTR |
| `tests/test_no_made_up_values.py`@ad3758c | 117-138 | M-09 | iPerf3 row |
| `tests/test_no_made_up_values.py`@ad3758c | 139-159 | R-01 | status codes |
| `tests/test_no_made_up_values.py`@ad3758c | 160-193 | R-02 | detailed CSV |
| `tests/test_no_made_up_values.py`@ad3758c | 194-228 | C-02 | correction pass |
| `tests/test_no_made_up_values.py`@ad3758c | 229-241 | S-11 | bot status |
| `tests/test_ookla_honesty.py`@ad3758c | 1-54 | M-05 | client cooldown |
| `tests/test_ookla_honesty.py`@ad3758c | 55-113 | M-10 | no disguised fallback, server mismatch |
| `tests/test_ookla_honesty.py`@ad3758c | 114-135 | R-01 | error rows in report |
| `tests/test_ookla_honesty.py`@ad3758c | 136-181 | M-10 | Ookla terms |
| `tests/test_ookla_honesty.py`@ad3758c | 182-345 | M-04 | hourly Ookla, engine intervals, fallback mix, busy iPerf3 |
| `tests/test_ookla_honesty.py`@ad3758c | 346-386 | M-09 | iPerf3 both directions, server name |
| `tests/test_ookla_honesty.py`@ad3758c | 387-407 | M-10 | writable HOME |
| `tests/test_ooni_cli.py`@ad3758c | 1-27 | M-13 | tests of the code of this package |
| `tests/test_ooni_toggle.py`@ad3758c | 1-23 | M-13 | tests of the code of this package |
| `tests/test_openvpn_names.py`@ad3758c | 1-109 | N-14 | tests of the code of this package |
| `tests/test_own_outages.py`@ad3758c | 1-115 | M-02 | own outages apart, CSV and dossier |
| `tests/test_own_outages.py`@ad3758c | 116-164 | C-07 | window recording, netd mode switch, alert |
| `tests/test_port_forward.py`@ad3758c | 1-202 | N-11 | tests of the code of this package |
| `tests/test_pppoe_retry.py`@ad3758c | 1-96 | N-05 | pppd command and watch |
| `tests/test_pppoe_retry.py`@ad3758c | 97-173 | N-10 | USB rescue route |
| `tests/test_pppoe_retry.py`@ad3758c | 174-227 | N-02 | WAN firewall |
| `tests/test_pppoe_retry.py`@ad3758c | 228-255 | N-03 | options file replaced, installer chown |
| `tests/test_pppoe_retry.py`@ad3758c | 256-286 | N-01 | router measurement link |
| `tests/test_pppoe_retry.py`@ad3758c | 287-335 | N-05 | shutdown order, port given back |
| `tests/test_probe_line.py`@ad3758c | 1-93 | M-01 | probe device, gateway targets |
| `tests/test_probe_line.py`@ad3758c | 94-102 | R-01 | switched-off targets |
| `tests/test_probe_line.py`@ad3758c | 103-137 | N-01 | wired metric |
| `tests/test_probe_line.py`@ad3758c | 138-187 | M-02 | outage reading and settling |
| `tests/test_probe_line.py`@ad3758c | 188-203 | S-11 | stale outage alert |
| `tests/test_probe_outage.py`@ad3758c | 1-71 | M-02 | tests of the code of this package |
| `tests/test_public_results_and_ookla_warning.py`@ad3758c | 1-56 | S-09 (Exports and public results API) | public results |
| `tests/test_public_results_and_ookla_warning.py`@ad3758c | 57-95 | M-05 | Ookla schedule warning |
| `tests/test_purge_and_verdict.py`@ad3758c | 1-63 | C-08 | purge |
| `tests/test_purge_and_verdict.py`@ad3758c | 64-96 | C-06 | verdict wording and bot language purity |
| `tests/test_redact.py`@ad3758c | 1-77 | C-04 | tests of the code of this package |
| `tests/test_regdomains.py`@ad3758c | 1-76 | N-07 | tests of the code of this package |
| `tests/test_review_fixes.py`@ad3758c | 1-48 | R-01 | verdict |
| `tests/test_review_fixes.py`@ad3758c | 49-77 | R-02 | LAN scope in CSV |
| `tests/test_review_fixes.py`@ad3758c | 78-101 | R-03 | dossier tables and charts |
| `tests/test_review_fixes.py`@ad3758c | 102-113 | R-01 | ISP name |
| `tests/test_review_fixes.py`@ad3758c | 114-124 | R-05 | e-mail texts |
| `tests/test_review_fixes.py`@ad3758c | 125-146 | M-02 | own window vs outage |
| `tests/test_review_fixes.py`@ad3758c | 147-165 | M-04 | runner |
| `tests/test_review_fixes.py`@ad3758c | 166-190 | C-02 | correction pass B |
| `tests/test_review_fixes.py`@ad3758c | 191-220 | S-12 | collector poll and summary |
| `tests/test_review_fixes.py`@ad3758c | 221-252 | S-11 | bot test command |
| `tests/test_review_fixes.py`@ad3758c | 253-306 | R-02 | server facts in CSV |
| `tests/test_review_fixes.py`@ad3758c | 307-320 | R-03 | escaping |
| `tests/test_review_fixes.py`@ad3758c | 321-399 | R-01 | fourth review (mixed: uptime, gateway, bot caption, RTR, correction C, CSV) |
| `tests/test_rtr_native.py`@ad3758c | 1-44 | M-08 | tests of the code of this package |
| `tests/test_rtr_output_tail.py`@ad3758c | 1-19 | M-08 | tests of the code of this package |
| `tests/test_secret_keep.py`@ad3758c | 1-46 | C-04 | tests of the code of this package |
| `tests/test_selftest.py`@ad3758c | 1-119 | P-07 | tests of the code of this package |
| `tests/test_selftest_auto.py`@ad3758c | 1-91 | P-07 | tests of the code of this package |
| `tests/test_selftest_router.py`@ad3758c | 1-78 | P-07 | tests of the code of this package |
| `tests/test_selftest_scenario.py`@ad3758c | 1-104 | P-07 | tests of the code of this package |
| `tests/test_setup_bot.py`@ad3758c | 1-223 | S-11 | bot says the setup, status, mode, outage alert |
| `tests/test_setup_bot.py`@ad3758c | 224-294 | P-07 | selftest cleanup and setup checks |
| `tests/test_setup_bot.py`@ad3758c | 295-335 | S-13 | collector setup column |
| `tests/test_setup_reports.py`@ad3758c | 1-166 | R-02 | CSVs |
| `tests/test_setup_reports.py`@ad3758c | 167-257 | R-03 | dossier |
| `tests/test_setup_reports.py`@ad3758c | 258-309 | R-04 | bundle |
| `tests/test_setup_reports.py`@ad3758c | 310-336 | R-05 | e-mail |
| `tests/test_setup_web.py`@ad3758c | 1-194 | S-02 (Status API) | latest result setup, current setup, language, status endpoint |
| `tests/test_setup_web.py`@ad3758c | 195-294 | U-02 (Dashboard page) | dashboard labels and wiring |
| `tests/test_setup_web.py`@ad3758c | 295-342 | S-10 | CLI latest |
| `tests/test_shipped_files.py`@ad3758c | 1-47 | P-11 | every file setup.sh installs exists in git |
| `tests/test_speed_schedule_default.py`@ad3758c | 1-23 | M-04 | tests of the code of this package |
| `tests/test_storage_guard.py`@ad3758c | 1-133 | C-08 | tests of the code of this package |
| `tests/test_tailscale_and_telegram_ui.py`@ad3758c | 1-52 | C-06 | i18n keys |
| `tests/test_tailscale_and_telegram_ui.py`@ad3758c | 53-65 | U-09 (Remote access and VPN page) | dashboard elements |
| `tests/test_tailscale_and_telegram_ui.py`@ad3758c | 66-130 | N-15 | Tailscale |
| `tests/test_tailscale_and_telegram_ui.py`@ad3758c | 131-157 | C-01 | Telegram config persistence |
| `tests/test_tailscale_and_telegram_ui.py`@ad3758c | 158-167 | U-01 | mobile CSS |
| `tests/test_tailscale_and_telegram_ui.py`@ad3758c | 168-178 | U-18 (Telegram page) | token input protection |
| `tests/test_tasks.py`@ad3758c | 1-85 | C-09 | tests of the code of this package |
| `tests/test_telegram_bot.py`@ad3758c | 1-305 | S-11 | tests of the code of this package |
| `tests/test_telegram_i18n.py`@ad3758c | 1-247 | C-06 | tests of the code of this package |
| `tests/test_theme.py`@ad3758c | 1-53 | P-09 | theme choice and persistence |
| `tests/test_traceroute.py`@ad3758c | 1-87 | M-11 | tests of the code of this package |
| `tests/test_tui.py`@ad3758c | 1-117 | S-10 | tests of the code of this package |
| `tests/test_ui_navigation_targets.py`@ad3758c | 1-34 | U-01 | tests of the code of this package |
| `tests/test_ui_text_coverage.py`@ad3758c | 1-96 | C-06 | tests of the code of this package |
| `tests/test_usb_import.py`@ad3758c | 1-68 | P-03 | tests of the code of this package |
| `tests/test_usb_router.py`@ad3758c | 1-100 | N-10 | USB as router WAN |
| `tests/test_usb_router.py`@ad3758c | 101-118 | N-03 | renew LAN clients |
| `tests/test_usb_router.py`@ad3758c | 119-168 | N-10 | fallback and router WAN route |
| `tests/test_usb_router.py`@ad3758c | 169-178 | N-08 | share from USB |
| `tests/test_usb_router.py`@ad3758c | 179-197 | N-10 | fallback event schema |
| `tests/test_usb_router.py`@ad3758c | 198-220 | N-01 | dead routes |
| `tests/test_usb_router.py`@ad3758c | 221-237 | N-07 | no AP outside router mode |
| `tests/test_usb_router.py`@ad3758c | 238-258 | N-01 | IPv4 DNS first |
| `tests/test_usb_tether.py`@ad3758c | 1-68 | N-10 | tests of the code of this package |
| `tests/test_usb_uplink.py`@ad3758c | 1-82 | N-10 | tests of the code of this package |
| `tests/test_wait_online.py`@ad3758c | 1-30 | P-10 | wait-online drop-in |
| `tests/test_web_ui_integrity.py`@ad3758c | 1-67 | U-01 | tests of the code of this package |
| `tests/test_wifi_band_caps.py`@ad3758c | 1-69 | N-07 | tests of the code of this package |
| `tests/test_wifi_flags.py`@ad3758c | 1-122 | N-07 | tests of the code of this package |
| `tests/test_wifi_psk.py`@ad3758c | 1-93 | N-07 | tests of the code of this package |
| `tests/test_wifi_reload.py`@ad3758c | 1-138 | N-03 | change to AP settings restarts hostapd only (ModeA.reapply_wifi) |
| `tests/test_wifi_share.py`@ad3758c | 1-48 | N-05 | DHCP packets |
| `tests/test_wifi_share.py`@ad3758c | 49-123 | N-08 | sharing sync |
| `tests/test_wifi_share.py`@ad3758c | 124-142 | N-05 | DISCOVER frame |
| `tests/test_wifi_share.py`@ad3758c | 143-172 | C-01 | parallel config saves |
| `tests/test_wifi_uplink.py`@ad3758c | 1-115 | N-08 | tests of the code of this package |
| `tests/test_write_private.py`@ad3758c | 1-25 | N-03 | tests of the code of this package |
## Proposed package changes

Ids are permanent: the new ids below are the next free numbers of their series and are never reused.

### A. Names for the open series (no new scope)

1. **S-02 to S-09** as named in the table at the top: the eight API areas of the work plan, cut along the route branches of
   `nmb/web/server.py@ad3758c:370-3441` and its helpers. `S-08` holds `/api/config` read and write (`nmb/web/server.py@ad3758c:3787-4208`).
2. **U-02 to U-21**, one per tab of `nmb/web/assets/dashboard.html@ad3758c:87-108`. Two Python sub-tabs change page:
   the measurement schedule and user-agent sub-tab of the router page (`dashboard.html@ad3758c:3568-3703`) goes to U-14 with the
   rest of the schedule, and the Telegram sub-tab (`dashboard.html@ad3758c:4777-4801`) goes to U-18.

### B. Files the work plan lists twice, cut by line range

Each of these is listed under two packages in work plan section 3; the table gives every line one owner:

| File | Plan lists it under | Cut used in the table |
| :--- | :--- | :--- |
| `nmb/probe/engine.py` | M-01, M-02 | 301-387 (outage settling and evaluation) M-02, rest M-01 |
| `nmb/report/metrics.py` | M-02, R-01 | 1-132 (outage spans, own-action windows) M-02, rest R-01 |
| `nmb/netd/modes.py` | N-03, N-05 | PPPoE watch 143-184 N-05; IPv6/UPnP/port-forward parts 185-227, 383-422, 615-642 N-11; USB WAN 228-287 N-10; rest N-03 |
| `nmb/netd/wifi_share.py` | N-05, N-08 | DHCP guard 132-251 N-05, rest N-08 |
| `nmb/netd/wifi.py` | M-14, N-07 | status 345-401 M-14, rest N-07 |
| `systemd/*` | P-04, P-10 | user, capability and sandbox lines P-04, the rest of each unit P-10 |

### C. Moves out of the package the plan names (with reason)

1. `nmb/netd/upnp_wan.py` from N-11 to **M-03**: it only reads the gateway's WAN byte counters for the traffic gate
   (`nmb/netd/upnp_wan.py@ad3758c:1`), and its test is in `tests/test_cross_traffic.py@ad3758c:163-181`.
2. `nmb/netd/geoip.py@ad3758c:154-214` (regulatory domain read and set) to **N-07**; the rest stays M-12.
3. `nmb/netd/interfaces.py@ad3758c:666-813` (public and dual-stack addresses) to **M-12** and `847-1070` (first hop, PPP peers,
   gateway ping targets) to **M-01**; the rest stays N-01.
4. `nmb/report/llm_bundle.py@ad3758c:412-620` (NAT and STUN, DNS interception, DNS matrix) to **M-12**: that is where the plan
   puts STUN, NAT and the DNS audit.
5. `nmb/netd/services.py@ad3758c:290-312, 392-471` (pppd) to **N-05**, `313-383` (miniupnpd, dhcpcd prefix delegation) to **N-11**.
6. `nmb/common/heal.py@ad3758c:1004-1310` (stable updates with rollback) to **P-02**; the rest stays P-01.
7. `setup.sh@ad3758c:393-454` (unique Wi-Fi password, admin password not printed, first-boot credentials file and its boot copy)
   to **P-03** ("unique password per device"); the rest of `setup.sh` stays P-11.
8. Translation tables inside code to **C-06**: `nmb/report/dossier.py@ad3758c:31-688`, `nmb/report/svg_charts.py@ad3758c:13-141`,
   `nmb/bot/i18n.py@ad3758c:415-1375`. The bot's language logic (per-chat language, keyboards, verdict matching) stays S-11.
9. The sub-CLIs in `bin/` go to the package of their subsystem (`nmb-adblock` N-12, `nmb-driver` N-17, `nmb-heal` P-01, `nmb-helper` P-06,
   `nmb-iperf` M-09, `nmb-ooni` M-13, `nmb-vpn` N-13, `nmb-wifi` N-08, `nmb-factory-reset` P-03, `nmb-invasive-log` R-05); P-10 keeps
   the thin wrappers and the multi-call names.
10. `os/99-nmb-performance.conf` and `os/nic-tune.sh` to **N-01** with `netd/kernel.py` (same tuning, at install and at run time);
    `os/storage-setup.sh` to **P-11**; `os/90-nmb.rules` and `os/90-nmb.pkla` to **P-06**; `os/90-nmb-usb-import.rules` to **P-03**;
    `os/ooni-resolv.conf` to **M-13**; `os/nmb-collector.service` to **S-13** (its sandbox lines P-04).
11. Themes to **P-09** (add-on themes, AMOLED-only core): the eight non-AMOLED theme blocks of `dashboard.css`, the theme chooser
    (`dashboard.html@ad3758c:3863-3951`, `dashboard.js@ad3758c:301-350`) and all 20 images in `nmb/web/static/themes/`.
    The AMOLED block (`dashboard.css@ad3758c:1006-1066`) is the U-01 look.

### D. Splits for size (the work plan's guide is about 1,500 lines of Python reference per package)

| Package | Reference after the cuts above | Proposed split |
| :--- | :--- | :--- |
| S-11 Telegram bot | about 3,000 lines (commands 1,725) | S-11 keeps daemon, client, authorisation, alerts, language logic and `commands.py@ad3758c:1-248`; **new S-14** "Bot commands: measuring, reports, mode, VPN" = `commands.py@ad3758c:249-1061`; **new S-15** "Bot commands: DynDNS, blocker, healing, tasks, Wi-Fi, backup, drivers, reboot" = `commands.py@ad3758c:1062-1725` |
| N-12 DNS blocker | 2,644 (adblock.py 2,031 + CLI 613) | N-12 keeps the DNS server, parser, engine and CLI; **new N-19** "Blocker lists: gravity update, allow/deny/regex lists, adlists, local DNS records" = `netd/adblock.py@ad3758c:1003-1515` |
| M-04 Runner and schedule | about 1,930 | M-04 keeps runner and cancel; **new M-15** "Scheduler: cron and interval parser, next run, retry timing" = `speed/scheduler.py` (also used inside the web server, `nmb/web/server.py@ad3758c:4478-4481`) |
| M-08 RTR | about 1,450 Python + about 2,900 C | M-08 keeps the engine rules (`speed/rtr.py`); **new M-16** "RMBT protocol client" = `speed/rmbt/` and `components/open-rmbt/` |
| P-07 Self-test | 2,837 | P-07 keeps `selftest/live.py`; **new P-17** "Guided live scenarios" = `selftest/scenario*.py` |
| R-04 Diagnostic bundle | about 1,870 | move `llm_bundle.py@ad3758c:39-168` (CPU, thermal) to C-09, `169-305` (PHY, ethtool) to N-17, `306-411` (Wi-Fi PHY) to M-14; no new id |
| C-01 Configuration | 1,646 | move the boot-partition settings copy (`config.py@ad3758c:58-80, 114-225`: find, target, merge, install, sync) to P-03 next to the USB import; secret mark, masking and the unique Wi-Fi password (81-113) stay C-01; no new id |
| C-06 Translations | about 27,000 lines of strings | C-06 keeps the key system, loader, English fallback and the build checks (complete, same placeholders and HTML, no Hungarian abroad); **new C-12** "Translation data conversion" = moving every string of `web/i18n.py`, `web/locales/`, `bot/messages.py`, the bot strings, `bot/locales/`, `report/locales/` and the dossier and chart tables into the F-04 key files |
| U-09 Remote access page | HTML 827 + JS 1,573 | U-09 keeps hub, guide, WireGuard, OpenVPN (`dashboard.html@ad3758c:1260-1529, 1670-1841`, `dashboard.js@ad3758c:8500-8787, 9107-9673`); **new U-22** "DynDNS, proxy and Tailscale page" (`dashboard.html@ad3758c:1530-1669, 1842-2086`, `dashboard.js@ad3758c:8788-9106, 9674-10072`) |
| U-16 Router page | HTML about 1,000 + JS about 1,500 | **new U-23** "Wi-Fi page: access point and Wi-Fi client" = `dashboard.html@ad3758c:3123-3511`, `dashboard.js@ad3758c:79-116, 4318-4611, 7546-7905` |
| U-17 System page | HTML about 800 + JS about 1,300 | **new U-24** "Backup, restore and stored data page" = `dashboard.html@ad3758c:4505-4651`, `dashboard.js@ad3758c:624-668, 5265-5546` |

Note for U-02: `refreshStatus` (`dashboard.js@ad3758c:2481-3131`) fills every page from one status poll; each page package
takes its own part of it.

### E. Not ported (accounted to a package only for the check)

`.claude/settings.json` (tool setting), every `__init__.py` and `__main__.py` (Python package markers and module entry points; the
native entry is a multi-call name in P-10), `tests/_box_isolation.py` (Python test helper; its rule goes to F-06), `os/net-detect.sh`
(installed by `setup.sh@ad3758c:336` but called by no code). The `os/openrc/` services are ported only if P-11 picks an OpenRC base.

### F. Packages with no Python file

C-10, N-09, N-18, P-05 and P-12 are new in the native build; F-01 to F-05 and F-07 produce documents and contracts; P-13 ports
`docs/TESTING.md`, which another part of the register covers. The register's check must accept these as "no Python reference".

### G. Site and personal values seen (Quality Bar 8)

`os/nmb-wait-online.conf@ad3758c:7` and `nmb/netd/ipv6_lan.py@ad3758c:3` name the owner's ISP or its ONT; the notes above state
them generically. No other personal value was copied. `nmb/netd/usb_uplink.py@ad3758c:3` names a commercial hotspot model; stated
as "a mobile hotspot".

## Questions for the owner

1. **Vendor-look add-on themes.** `nmb/web/static/themes/` holds images named after router vendors and a vendor gaming logo
   (`asus_title.png`, `tuf_logo.png`, `tplink_*.jpg`), used by the AsusWRT, AsusCamouflage and TP-Gold themes. May these ship as
   add-ons of the open native repository, or must they be redrawn without vendor logos (or kept out)?
2. **`nmb/web/static/themes/hago.jpg`** (MountainHaGo theme): is this a personal photo or a place of yours? If it is, it stays out
   of the native repository (Quality Bar 8).
3. **Browser terminal.** `/api/terminal/exec` (`nmb/web/server.py@ad3758c:3322-3407`, page U-20) runs shell commands as root from
   the browser; it is off unless `terminal_enabled` is set (`config/nmb.conf.example@ad3758c:143-144`). Should the native build have
   it at all? Package S-08/U-20 need the answer before their brief.
4. **Origin of `components/open-rmbt/`.** The repository has no licence file for the C client. Is it upstream RTR code (and which
   project and version), or was it written for NMB? The licence of the native repository follows from its components.
5. **`pivpn` and `pihole` command names** (`setup.sh@ad3758c:486-489`): keep these compatibility names in the native multi-call
   binary, or only the `nmb-*` names?
