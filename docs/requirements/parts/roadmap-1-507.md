# Requirements register, part: roadmap lines 1-507

Source: `docs/DEPYTHONIZATION_ROADMAP.md` of the Python repository at commit `ad3758c` (1,034 lines; this part covers lines 1-507).
Read in pages: 1-150, 151-300, 301-450, 451-507 (all lines). Package ids from `docs/NATIVE_WORK_PLAN.md@ad3758c:98-219` (section 3), read with sections 1-2 (lines 1-97) for the step-1 scope.
Id range of this part: REQ-RM-0001 to REQ-RM-0219 (0213 added in round 1, 0214-0219 in round 2). "RM" below in Source means `docs/DEPYTHONIZATION_ROADMAP.md@ad3758c`, written out in full in every cell.
Proposed package `NEW: X2-01` is explained in section 3 of this file.

## Requirements

| Id | Requirement | Source | Package | Also in | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| REQ-RM-0001 | The native build provides the function of every Python daemon (`nmb-probe`, `nmb-netd`, `nmb-speed`, `nmb-web`, `nmb-invasive-log`, `nmb-adblock`, `nmb-heal`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:9 | P-10 | M-01, N-03, M-04, S-01, R-05, N-12, P-01 | Source list omits `nmb-telegram`, `nmb-cli`/TUI and `nmb-collector`; P-10 takes the full unit list from the Python `systemd/` folder. |
| REQ-RM-0002 | The resident memory (RSS) of all native components together is measured on every target and recorded; it must fit the target's RAM without the OOM killer (estimate for Go, all-in-one: 10-18 MB). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:13; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:58 | P-13 | P-10 | Figures in the source are estimates; to measure (work plan 2.4 records memory). |
| REQ-RM-0003 | The binary size is measured and recorded per target; the Go build fits 16 MB+ flash (estimate 8-12 MB, 3.2 MB with UPX). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:14; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:57; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:59 | P-13 | P-11, F-00 | UPX is optional; it unpacks into RAM and raises RSS, so its use is a measured decision. |
| REQ-RM-0004 | Start-up is fast enough that the service manager (systemd, procd) never times out ("instantaneous expected"); start-up time is measured. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:15 | P-10 | P-13 | |
| REQ-RM-0005 | Probe timing has sub-millisecond precision; runtime pauses (garbage collection) must not distort latency, jitter or bufferbloat figures. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:16 | M-01 | M-04, R-01 | Proof: timestamp spread measured under load. |
| REQ-RM-0006 | No routine spawning of external tools (`ip`, `ps`, `nmcli`, `sysctl`) for scans and polling. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:17 | P-13 | N-01, N-07, C-09, M-14 | Each package proves its own part; P-13 checks the whole binary (scan for exec calls). |
| REQ-RM-0007 | One zero-dependency single static binary `nmbd`; no external runtime packages. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:20 | P-10 | F-00, P-11 | |
| REQ-RM-0008 | `nmbd` runs on Linux devices with as little as 16 MB SPI flash and 32-64 MB RAM. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:20 | NEW: X2-01 Step-2 router tiers | P-13 | 32 MB RAM / 16 MB flash is the C tier (line 75); the Go tier is 64 MB+ (REQ-RM-0013). |
| REQ-RM-0009 | All core and extended components (Open-RMBT, OONI Probe, iPerf3, DNS ad-blocker, WireGuard/OpenVPN gateways, hardware driver installer, self-healing watchdog, real-time task manager) are fully integrated into the native architecture. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:22 | P-13 | M-08, M-13, M-09, N-12, N-13, N-14, N-17, P-01, C-09 | Work plan 1: done only when every component of section 5 passes. |
| REQ-RM-0010 | Implementation in compiled systems languages: Go (step 1) and C11 (router tiers); no Python at run time. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:22 | F-00 | NEW: X2-01 Step-2 router tiers | |
| REQ-RM-0011 | Go toolchain 1.23 or newer. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:55 | F-00 | | Work plan F-00: the PC's installed Go version is older (source names a specific site value; not copied); F-00 fixes the version. |
| REQ-RM-0012 | The Go build cross-compiles for `linux/mips` (`GOOS=linux GOARCH=mips`) without a C SDK. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:60 | F-00 | C-02 | Work plan 2.4 matrix: mips and mipsle build only. |
| REQ-RM-0013 | Go is the primary universal target for 64 MB+ RAM devices: Raspberry Pi (Zero, 2, 3, 4, 5), modern GL.iNet routers, x86 appliances, Teltonika routers, all Linux distributions. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:72-73 | P-11 | F-00, NEW: X2-01 Step-2 router tiers | Source names router models (product targets). Pi Zero and Pi 1 are ARMv6: Q-03 settled, cross-compiled only in step 1 (REQ-RM-1440), not tested or promised. Routers are step 2. |
| REQ-RM-0014 | Static build with `CGO_ENABLED=0`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:74 | F-00 | P-10 | |
| REQ-RM-0015 | The whole frontend (HTML, JS, CSS, SVG icons) is embedded in the executable (`//go:embed`); no web files are installed on disk. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:74; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:62; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:150 | U-01 | S-01, P-10 | |
| REQ-RM-0016 | C11/musl tier for ultra-constrained OpenWrt (classic MIPS routers, 32 MB RAM, 16 MB flash): sub-2 MB RAM, SQLite amalgamation, raw epoll loop, direct rtnetlink calls. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:75-77; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:59 | NEW: X2-01 Step-2 router tiers | F-07 | Step 2; F-07 spec and vectors are its starting point. |
| REQ-RM-0219 | The binary size of the C11/musl tier is measured and recorded per target; it fits 8 MB / 16 MB SPI flash (estimate: under 1.5 MB, 450 KB-1.2 MB, 280 KB with UPX). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:42; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:57; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:59 | NEW: X2-01 Step-2 router tiers | P-13, T-01 | Added in round 2 (C-tier counterpart of REQ-RM-0003). Figures in the source are estimates; to measure. UPX is optional (it unpacks into RAM). |
| REQ-RM-0017 | ICMP via `golang.org/x/net/icmp`, with unprivileged ICMP datagram sockets (`SOCK_DGRAM`) or raw ICMP. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:119; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:65 | M-01 | P-04 | Unprivileged ICMP needs `net.ipv4.ping_group_range`; sandbox profile must allow the chosen way. |
| REQ-RM-0018 | Running mean, variance and jitter in constant memory (Welford's algorithm). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:120-121 | M-01 | R-01 | The jitter definition must equal the Python one (F-05 golden files). |
| REQ-RM-0019 | Aggregated probe metrics are flushed to the database every 10 s. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:122 | M-01 | C-02 | Check against the Python minute summaries (M-01 brief); flush interval is not the summary period. |
| REQ-RM-0020 | C tier probe: single-threaded `epoll` + `timerfd` loop, non-blocking echo requests to all targets at once, zero heap allocation in the steady state. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:123-126 | NEW: X2-01 Step-2 router tiers | M-01 | |
| REQ-RM-0021 | netd opens a netlink socket and subscribes to link and route events (`RTM_NEWLINK`, `RTM_DELLINK`, `RTM_NEWROUTE`) instead of polling. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:130-132; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:64 | N-01 | N-03 | |
| REQ-RM-0022 | Strict wired Ethernet priority: on link up/down the default route metric is changed at once by a netlink message, without external binaries. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:129; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:133 | N-01 | N-03, N-08, N-10 | Must respect `measurement_link` (REQ-RM-0156) and that USB is never the wired line (REQ-RM-0166). |
| REQ-RM-0023 | Firewall rules are set without spawning `nft` or other external tools. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:129 | N-02 | | Read from the bottleneck line (Python spawns `nftables`); N-02 brief confirms the method (nftables netlink). |
| REQ-RM-0024 | C tier netd: direct `socket(AF_NETLINK, SOCK_RAW, NETLINK_ROUTE)`; on OpenWrt linked against `libnl-tiny`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:134-136 | NEW: X2-01 Step-2 router tiers | N-01 | |
| REQ-RM-0025 | The speed engines sustain multi-gigabit throughput (1-10 Gbps) without saturating the CPU; CPU use during tests is measured. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:139 | M-04 | M-06, M-07, M-08, M-09, M-10, P-13 | Step-1 test hardware is 1 Gbps; above that only on suitable hardware (record as "not tested" otherwise). |
| REQ-RM-0026 | Multi-stream parallel HTTP/2 and HTTP/1.1 download worker pool (4-16 streams). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:140-141 | M-06 | M-07, M-10 | Shared by HTTP engines; see Proposed package changes (optional M-17). |
| REQ-RM-0027 | Every engine binds its sockets explicitly to the measuring interface/source address (`SO_BINDTODEVICE` or `IP_BIND_ADDRESS_NO_PORT` via `net.Dialer.Control`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:142 | M-04 | M-06, M-07, M-08, M-09, M-10, N-08, N-10, C-11 | `SO_BINDTODEVICE` needs `CAP_NET_RAW` (P-04). |
| REQ-RM-0028 | A native Ookla HTTP protocol engine replaces the sandboxed Ookla program only when it gives the same results on the same line and Ookla's terms allow it; otherwise the native engine is recorded as not done and the sandboxed program stays. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:143; docs/NATIVE_WORK_PLAN.md@ad3758c:313-316 | M-10 | M-05 | The roadmap line asks for the native protocol without a condition; the owner decision (work plan section 7 point 3) makes it conditional, so M-10's gate needs either the equal-results proof and the terms finding, or the not-done record. Same rule as REQ-WP-0180; REQ-RM-1415 treats the Ookla part of line 949 the same way. Round 2: restated. |
| REQ-RM-0029 | Native Cloudflare speed protocol (`/cdn-cgi/trace` plus chunked download and upload). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:143 | M-06 | | |
| REQ-RM-0030 | Native RTR-Netztest RMBT protocol in the speed engine. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:143 | M-08 | M-04 | Detail in 4.6 (REQ-RM-0060 to -0063). |
| REQ-RM-0031 | A concurrent background ICMP ping runs during every speed test and a bufferbloat grade is calculated. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:144 | M-04 | M-01, R-01, U-02 | Grade scale must match the Python reference (F-05). |
| REQ-RM-0032 | Web server: standard `net/http` with HTTP/2 and keep-alive. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:148-149; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:62 | S-01 | | |
| REQ-RM-0033 | Password comparison is constant-time (`subtle.ConstantTimeCompare`); the login is a session login with CSRF protection (work plan S-01), not the HTTP Basic Auth the source names. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:151 | S-01 | | Q-18 answered: work plan 7 point 1 (API left to the developer) and Quality Bar 7 (CSRF protection, secure cookies, rate-limited login); the Basic Auth line is out of date. Restated in round 3. |
| REQ-RM-0034 | The web server listens on port 80 by default and falls back to 8080 when the bind fails with `EACCES` (unprivileged). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:152; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:471 | S-01 | P-04, P-10 | |
| REQ-RM-0035 | C tier web: embedded `mongoose.c` or `libmicrohttpd`; frontend gzipped into a C byte array. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:153-155 | NEW: X2-01 Step-2 router tiers | U-01 | |
| REQ-RM-0036 | The native Web UI has every function of the Python UI (nothing dropped to make it small); the Python UI is the reference for what pages do, not how they are built. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:159 | U-01 | U-02, P-13 | Each page package U-02…U-0n proves its page against the Python page. |
| REQ-RM-0037 | A small shell (header, menu, login state, AMOLED black styles) loads first; each page is its own module, loaded only when opened, then kept in the browser cache. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:165 | U-01 | U-02 | |
| REQ-RM-0038 | Compressed size budgets are checked at every build and the build fails above them: shell with first page at most 60 KB on the PC/Pi class; whole UI at most 150 KB. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:166 | U-01 | F-00 | Round 2: the router budget and the flash fit moved to REQ-RM-0214 (split rule of the register README). |
| REQ-RM-0214 | Router builds: the shell with the first page is at most 40 KB compressed, and the whole UI fits the 8 MiB flash of the TL-WR1043ND v1 next to `nmbd`; the build fails above these budgets. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:166 | NEW: X2-01 Step-2 router tiers | U-01, F-00, T-01 | Split from REQ-RM-0038 in round 2. The TL-WR1043ND v1 is the plans' public first router target, not site data (naming rule in the register README). |
| REQ-RM-0039 | Each build prints the compressed size per module; growth is a reviewed change. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:166 | U-01 | F-00 | |
| REQ-RM-0040 | Assets are stored pre-compressed in the binary (gzip, and Brotli where the client accepts it), with content-hash file names, long cache lifetime, strong `ETag` and `304 Not Modified`; HTTP keep-alive (HTTP/2 where the platform has it). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:167 | S-01 | U-01 | |
| REQ-RM-0041 | Nothing is loaded from the internet (no CDN, no web fonts); the UI works without internet on the box's own LAN. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:167 | U-01 | S-01 | |
| REQ-RM-0042 | Live data over one stream (SSE or WebSocket) carrying what the open page shows; the server sends only changes, at the rate the page needs, and nothing while no page is open. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:168 | S-01 | U-01 | Python page keeps 11 polling timers (line 161). |
| REQ-RM-0043 | The browser stops all updates while the tab is hidden (Page Visibility API) and resumes with one catch-up request. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:168 | U-01 | | |
| REQ-RM-0044 | Long lists (logs, processes, DNS queries, results) are paged and filtered on the server, never sent whole. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:168 | S-02 | U-02 | Applies to every API area package S-02…S-0n. |
| REQ-RM-0045 | No front-end framework; standard JavaScript modules; one style sheet on CSS variables (add-on themes replace only variables); icons in one SVG sprite; the DOM is updated in place, never rebuilt every second. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:169 | U-01 | P-09, U-02 | |
| REQ-RM-0046 | No inline scripts or styles; a strict Content Security Policy is served. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:169 | U-01 | S-01 | |
| REQ-RM-0047 | Every visible text goes through a translation key; only the current language is sent, per page module; the language menu shows each language in its own name; a build check fails on any visible text without a key and any key missing from a language; logic never depends on visible text (states, data attributes). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:170 | C-06 | U-01, U-02 | Python fault (line 161): logic read visible text; F-03. |
| REQ-RM-0048 | The API answers from data the daemons already hold; expensive values (hardware detection, Wi-Fi capabilities, reports) are computed once and cached, not on every page load. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:171 | S-02 | N-07, N-17, R-03 | Python fault: `/api/config` runs radio detection on each call; F-03 entry. |
| REQ-RM-0049 | An open UI must not visibly change measurement results: box CPU use with the dashboard open is measured and kept low (target, to measure: under 2 % on the Pi 4). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:171 | U-01 | P-13 | Target to measure. Round 2: the router figure moved to REQ-RM-0215. |
| REQ-RM-0215 | Router builds: box CPU use with the dashboard open is measured and kept low (target, to measure: under 10 % on the TL-WR1043ND v1). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:171 | NEW: X2-01 Step-2 router tiers | U-01, P-13, T-01 | Split from REQ-RM-0049 in round 2. |
| REQ-RM-0050 | Speed of use (targets to confirm by measurement): first useful screen under 1 s on the Pi 4 over Wi-Fi from a mid-range phone; page switch under 300 ms after first load; no layout jumps while data arrives. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:172 | U-01 | P-13 | Round 2: the router figure moved to REQ-RM-0216. |
| REQ-RM-0216 | Router builds (target to confirm by measurement): first useful screen under 2 s on the TL-WR1043ND v1 over Wi-Fi from a mid-range phone. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:172 | NEW: X2-01 Step-2 router tiers | U-01, P-13, T-01 | Split from REQ-RM-0050 in round 2. |
| REQ-RM-0051 | The UI works on a phone screen, with the keyboard alone, and with enough contrast. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:172 | U-01 | U-02 | |
| REQ-RM-0052 | A headless browser test opens every page in every language on the PC build and records size, load time, requests and console errors; the numbers go into `docs/TESTING.md` next to the Python UI's. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:173 | U-01 | P-13 | Round 2: the router build's place in the release checklist moved to REQ-RM-0217. |
| REQ-RM-0217 | The release checklist includes the router build (the headless browser figures of REQ-RM-0052 for the router build). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:173 | NEW: X2-01 Step-2 router tiers | U-01, P-13 | Split from REQ-RM-0052 in round 2; there is no router build in step 1. |
| REQ-RM-0053 | Every page and sub-page has its own readable path (e.g. `/speed`, `/speed/schedule`, `/router/wifi`) set with the History API; no `#` fragments and no single `/` for everything. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:174 | U-01 | U-02 | Owner, 2026-11-10. |
| REQ-RM-0054 | Reload, bookmark, shared link, new tab, and back/forward all land on that page, in its state where it makes sense (selected log source, filter). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:174 | U-01 | U-02 | |
| REQ-RM-0055 | The server answers every page path with the shell; an unknown path gets a proper "not found" page, not the dashboard. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:174 | S-01 | U-01 | |
| REQ-RM-0056 | The page title follows the page (`<Page> – NetworkMeasureBox`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:174 | U-01 | C-06 | |
| REQ-RM-0057 | API paths stay under `/api/` and never collide with page paths. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:174 | S-01 | F-04 | |
| REQ-RM-0058 | Pure Go SQLite driver (`modernc.org/sqlite` or `github.com/glebarez/go-sqlite`), zero cgo, static for mips, mipsel, arm, arm64 and x86_64. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:178-180; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:63 | C-02 | F-00 | |
| REQ-RM-0059 | Live database in RAM at `/run/nmb/live.db` (WAL), synchronised to flash at `/var/lib/nmb/persist/nmb.db` every 15 minutes. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:181; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:489-490 | C-02 | F-04 | Paths and interval checked against the Python `nmb/db/` (F-04 contract). |
| REQ-RM-0060 | RMBT implemented natively in Go (`net.Dial`, `crypto/tls`); no external `rmbt-client` binary. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:186-187; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:192 | M-08 | P-10 | |
| REQ-RM-0061 | RMBT control negotiation over TLS (REST/JSON) with the RTR control server. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:188-189 | M-08 | | |
| REQ-RM-0062 | RMBT raw TCP chunk streams with nanosecond timestamps, several parallel streams over the dedicated test ports (e.g. 8081/8082). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:190-191 | M-08 | | |
| REQ-RM-0063 | The native RMBT keeps full fidelity with the official RTR test (results comparable to the reference client). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:192 | M-08 | F-05 | Proof: side-by-side runs against the C client built from `components/open-rmbt`. |
| REQ-RM-0064 | C tier: `components/open-rmbt/` statically linked as `libopenrmbt.a`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:193-194 | NEW: X2-01 Step-2 router tiers | M-08 | |
| REQ-RM-0065 | OONI runs in-process as a Go module (upstream probe-engine / `libminiooni`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:199-201 | M-13 | P-08 | Upstream merged into `probe-cli`; `libminiooni` lives under `internal/` and may not be importable from another module: analyst verifies (external docs skill). Licence (GPL-3.0) affects the project licence. |
| REQ-RM-0066 | OONI with zero external process spawning and no external package (no `ooniprobe` Debian package or PPA). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:202-204 | M-13 | | |
| REQ-RM-0067 | OONI results (DNS tampering, TCP blocking, HTTP diff, Web Connectivity) are streamed in memory into the database. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:205 | M-13 | C-02 | |
| REQ-RM-0068 | Lightweight OONI configuration for routers: the heavy WebSockets and WebRTC tests are stripped so a run stays within an 8 MB RAM envelope. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:206 | NEW: X2-01 Step-2 router tiers | M-13, T-01 | RSS during a run is measured. Round 2: owner moved from M-13 to X2-01 (router-only row). |
| REQ-RM-0069 | Native iPerf3 client protocol: TCP/UDP stream framing with JSON control negotiation on port 5201. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:211-212 | M-09 | | Source says "RFC-compliant"; iPerf3 has no RFC, the protocol is defined by the iperf3 sources (two-source rule). |
| REQ-RM-0070 | Zero-fork iPerf3 throughput, jitter and packet-loss measurement for LAN benchmarks and remote WAN tests. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:213 | M-09 | | |
| REQ-RM-0071 | C tier: stripped `libiperf.a` (<300 KB) linked into `nmbd`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:214-215 | NEW: X2-01 Step-2 router tiers | M-09 | |
| REQ-RM-0072 | The native DNS blocker keeps the Python function: DNS on UDP and TCP port 53, blocklists, forwarding to the upstream DNS. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:218 | N-12 | | |
| REQ-RM-0073 | The DNS blocker is built on `github.com/miekg/dns`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:220-221 | N-12 | P-08 | |
| REQ-RM-0074 | Blocked domains in a compact in-memory structure (radix tree, or Bloom filter + Aho-Corasick); target ~3.5 MB RAM for 150,000 domains and under 0.2 ms per lookup. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:222-224 | N-12 | P-13 | A Bloom filter has false positives: a hit must be confirmed exactly so no allowed domain is blocked. Targets to measure. |
| REQ-RM-0075 | Built-in DoH and DoT upstream resolvers with connection pooling. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:225 | N-12 | | |
| REQ-RM-0076 | In-memory DNS query-log ring buffer exposed over the web API and the UNIX socket `/run/nmb/adblock.sock`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:226 | N-12 | S-02 | Socket permissions are part of the security check (QB 7). |
| REQ-RM-0077 | Embedded userspace Tailscale node (`tsnet`) inside `nmbd`: obtains a 100.x.y.z mesh address, needs no `tailscaled`, gives remote dashboard access without WAN port forwarding or a public IP. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:231-235 | N-15 | S-01, P-08 | |
| REQ-RM-0078 | Kernel WireGuard managed in-process via `wgctrl` and netlink `RTM_NEWLINK` (no `wg-quick`/`wg`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:236 | N-13 | N-01 | |
| REQ-RM-0079 | Curve25519 keypairs generated in-process (`crypto/ecdh`); SVG and PNG QR codes rendered in memory; no `qrencode` or shell scripts. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:237 | N-13 | S-11, U-02 | `skip2/go-qrcode` writes PNG only; SVG needs an own renderer from its bitmap. |
| REQ-RM-0080 | Where Tailscale cannot run: remote access through a WireGuard tunnel dialled out to a hub, a Go box as Tailscale subnet router, or a reverse SSH tunnel (porting plan 7D). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:238 | N-18 | N-13, P-05 | |
| REQ-RM-0081 | OpenVPN (EC) gateway managed through a management-socket client to the local OpenVPN daemon (`/run/openvpn/mgmt.sock`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:239 | N-14 | | Socket permissions in the security check. |
| REQ-RM-0082 | C tier VPN: netlink configuration with `libnl-tiny`; OpenWrt `kmod-wireguard` control within <400 KB RAM. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:240-242 | NEW: X2-01 Step-2 router tiers | N-13 | |
| REQ-RM-0083 | Driver finder walks `/sys/bus/pci/devices` and `/sys/bus/usb/devices` reading `modalias`, `vendor` and `device`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:245; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:247-248 | N-17 | | |
| REQ-RM-0084 | A compressed lookup table of known Wi-Fi (Realtek, MediaTek, Atheros, Broadcom) and Ethernet chipsets is compiled into the binary. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:249 | N-17 | | |
| REQ-RM-0085 | On Debian/Ubuntu the driver finder installs firmware through `apt-get` or precompiled `.deb` packages. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:245; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:250 | N-17 | P-06 | Privileged: only through the helper's allow-list (QB 7). |
| REQ-RM-0086 | On OpenWrt the driver finder produces exact `opkg install kmod-<chipset>` commands or fetches modules from the official feed. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:251 | NEW: X2-01 Step-2 router tiers | N-17 | |
| REQ-RM-0087 | Self-healing keeps the Python function: periodic audits, binary verification, permission fixes, connectivity tests, rollback backups. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:254 | P-01 | P-02, P-03 | Source path `nmb/report/self_heal.py` differs from the plan's `common/heal.py`; analyst checks. "Updating via Git or tarballs" is replaced by P-02 (signed upload/channel). |
| REQ-RM-0088 | Independent watchdog supervisor inside `nmbd` with hardware watchdog support (`/dev/watchdog`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:256-257 | P-01 | P-10 | A hardware watchdog reboots the device on a hang; tested under the owner's conditions (work plan 2.2 point 10). |
| REQ-RM-0089 | Update payloads are verified cryptographically (SHA-256 and Ed25519 signature). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:258 | P-02 | | Owner decision: signing yes. |
| REQ-RM-0090 | Atomic update: new binary written to `nmbd.new`, verified, renamed over `nmbd` atomically, then a graceful restart (`execve`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:259 | P-02 | P-01, P-10 | |
| REQ-RM-0091 | OpenWrt A/B sysupgrade compatibility: state kept across firmware flashing (`sysupgrade -k`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:260 | NEW: X2-01 Step-2 router tiers | P-02 | Meaning of `-k` to verify against OpenWrt documentation. |
| REQ-RM-0092 | The task manager keeps the Python fields and refresh: CPU from `/proc/stat`, memory from `/proc/meminfo`, process list with pid, ppid, user, pri, ni, vsz, rss, stat, %cpu, %mem, time, comm, args. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:263 | C-09 | U-02, S-02 | Python refresh every 2 s; source line 272 speaks of a 1-second auto-refresh: take the period from the Python reference. |
| REQ-RM-0093 | No `ps` subprocess: `/proc/stat` and `/proc/meminfo` read with pre-allocated buffers; `/proc/[pid]` iterated, reading `stat` (ticks, state, priority, nice, RSS), `cmdline` (NUL-separated) and `status` (UID to user). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:265-271 | C-09 | | |
| REQ-RM-0094 | Task list generation target under 1.2 ms, near-zero CPU at a 1-second refresh. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:272 | C-09 | P-13 | Target to measure. |
| REQ-RM-0095 | Signals are sent to processes natively (`kill` syscall), no subprocess. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:273 | C-09 | P-06, S-02 | Root boundary: authenticated, allow-listed signals and targets (QB 7). |
| REQ-RM-0096 | Wi-Fi control talks directly to the `wpa_supplicant` control socket (`/var/run/wpa_supplicant/<iface>`) and the `hostapd` control socket (`/var/run/hostapd/<iface>`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:278-279 | N-07 | N-08 | Python uses NetworkManager profiles (line 400); the native build is NetworkManager-free (line 412); F-04 platform interface decides the owner of `wpa_supplicant`. |
| REQ-RM-0097 | Raw control messages (`SCAN`, `SCAN_RESULTS`, `ADD_NETWORK`, `SET_NETWORK`, `SELECT_NETWORK`, `ENABLE_NETWORK`) over the UNIX socket, without `wpa_cli`/`hostapd_cli`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:280 | N-07 | | |
| REQ-RM-0098 | Association and disconnection events (`CTRL-EVENT-CONNECTED`, `CTRL-EVENT-DISCONNECTED`) are received asynchronously from the socket. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:281 | N-07 | C-07 | |
| REQ-RM-0099 | The diagnostic logger keeps the Python function: watches packet loss, ping latency and memory anomalies and writes structured events. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:284 | R-05 | C-03 | |
| REQ-RM-0100 | In-memory circular ring buffer holding the last 5,000 diagnostic events. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:286-287 | R-05 | | |
| REQ-RM-0101 | Interface drops/errors from netlink and `/proc/net/dev` delta counters. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:288 | R-05 | N-01 | |
| REQ-RM-0102 | No flash wear: diagnostic events kept in memory, written to the database only when an anomaly is detected or on user request. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:289 | R-05 | C-02 | |
| REQ-RM-0103 | Telegram bot: asynchronous worker long-polling `getUpdates` over HTTPS. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:294-295 | S-11 | | |
| REQ-RM-0104 | Bot inline-keyboard menus for: status and instant ping matrix; on-demand speed test (Ookla, Cloudflare, RMBT, iPerf3); SLA dossier (.html) and LLM diagnostic snapshot (.md) export; Tailscale, WireGuard and OpenVPN status and one-click QR delivery. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:296-300 | S-11 | M-04, R-03, R-04, N-13, N-14, N-15 | QR delivery puts a VPN private key into Telegram: owner question Q-15. |
| REQ-RM-0105 | Bot language switcher over every UI language (11 now) with zero mixed-language leaks. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:301 | S-11 | C-06 | Source lists 5 languages; Q-37 answered by Quality Bar 9 (all languages always available) and line 406: the 5-language list is out of date. Restated in round 3. |
| REQ-RM-0106 | Bot memory target: under 1.5 MB RAM in Go. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:302 | S-11 | P-13 | Target to measure. Round 2: the C11 figure moved to REQ-RM-0218. |
| REQ-RM-0218 | Bot memory target in the C11 code base: under 200 KB RAM. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:302 | NEW: X2-01 Step-2 router tiers | S-11, P-13, T-01 | Target to measure. Split from REQ-RM-0106 in round 2. |
| REQ-RM-0107 | Dossier templates (HTML skeleton, CSS, translation dictionaries of every UI language) compiled into the binary. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:307-308 | R-03 | C-06 | Source lists 5 languages; Q-37 answered by Quality Bar 9 (every report and chart label in every language) and line 406. Restated in round 3. |
| REQ-RM-0108 | In-process SVG chart renderer: responsive `<svg viewBox="0 0 800 380">`, cubic Bézier curves, dual y-axes (Mbps and ms), SLA threshold baselines, touch dot overlays. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:309 | R-03 | | |
| REQ-RM-0109 | Neutral dossier title ("Broadband Quality & ISP SLA Performance Dossier"), regulator-neutral, following BEREC guidelines and international telecom standards. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:310 | R-03 | C-06 | |
| REQ-RM-0110 | Dossier generation without subprocesses; target under 1.5 ms. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:311 | R-03 | | Target to measure. |
| REQ-RM-0111 | Real-time streaming classification of packet loss: micro-spikes / bufferbloat (1-3 consecutive drops during high-throughput saturation) and macro-outages (at least 5 consecutive seconds of loss). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:316-319 | M-02 | R-01 | Python does this post hoc (line 314); compare thresholds with the Python engine (F-05). |
| REQ-RM-0112 | Root-cause correlation: local CPE failure (default gateway unreachable) vs. ISP upstream failure (gateway answers, public DNS/peering targets fail). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:320 | M-02 | R-01 | |
| REQ-RM-0113 | Outage events recorded with start, end, duration in seconds and root-cause class. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:321 | M-02 | C-07, C-02, R-01 | |
| REQ-RM-0114 | Wi-Fi telemetry from nl80211: signal (dBm) and noise floor, TX/RX bitrate, MCS index and guard interval, channel width, `tx_retries`, `tx_failed`, `rx_drop_misc`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:326-332 | M-14 | U-02 | |
| REQ-RM-0115 | Wi-Fi telemetry without `iw` subprocesses or regex parsing; refresh target under 0.5 ms. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:324; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:333 | M-14 | | |
| REQ-RM-0116 | Concurrent multi-resolver DNS probe: local gateway / DHCP-assigned DNS, Cloudflare (1.1.1.1, 1.0.0.1), Google (8.8.8.8, 8.8.4.4), Quad9 (9.9.9.9). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:338-343 | M-12 | C-10 | New function (Python: single `getaddrinfo`, line 336). The public resolver list belongs in the country/data files if location-aware (QB 9, C-10, REQ-RM-1112). |
| REQ-RM-0117 | Per resolver: response latency, cached vs. uncached, in ms. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:344-345 | M-12 | | |
| REQ-RM-0118 | Per resolver: DNSSEC validation status from the AD bit. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:346 | M-12 | | |
| REQ-RM-0119 | Transparent hijack check: a canary domain with known response signatures detects silent interception of UDP port 53. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:347 | M-12 | | Canary domain is chosen in the brief. |
| REQ-RM-0120 | RFC 5389 STUN: 20-byte Binding Request over UDP to a public STUN server (source: `stun.l.google.com:19302` or `stun.cloudflare.com:3478`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:352-353 | M-12 | | Server list belongs in the country/data files if location-aware (QB 9, C-10). |
| REQ-RM-0121 | Public reflexive endpoint read from the `XOR-MAPPED-ADDRESS` attribute. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:354 | M-12 | | |
| REQ-RM-0122 | Topology: ISP CGNAT when the WAN IPv4 is in 100.64.0.0/10; double NAT when it is RFC 1918 private; direct public when it equals the STUN reflexive address. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:355-358 | M-12 | C-11 | Source lists only 192.168/16 and 10/8; RFC 1918 also has 172.16/12 (all three apply). VPN addresses are excluded (REQ-RM-0195). |
| REQ-RM-0123 | NAT traversal classifier: full cone, restricted cone, port-restricted or symmetric; a type the probes cannot tell apart (no RFC 5780 server answered) is reported as unknown, never guessed. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:359 | M-12 | | Not possible with RFC 5389 Binding alone. Q-23 settled (developer's choice): RFC 5780 where a server supports it, otherwise "unknown" (no made-up results, fault register *Measurement honesty*). Restated in round 3. |
| REQ-RM-0124 | Per-core interrupt distribution from `/proc/interrupts` and `/proc/softirqs`, detecting CPU0 single-core saturation. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:364-365 | C-09 | U-02 | Python: missing or ad hoc (line 362). |
| REQ-RM-0125 | Packet steering check: RPS bitmasks (`/sys/class/net/<iface>/queues/rx-*/rps_cpus`) and irqbalance status. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:366 | C-09 | | |
| REQ-RM-0126 | NIC offload features (rx/tx checksumming, TSO, GRO) read through `SIOCETHTOOL`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:367 | C-09 | | ethtool netlink is an alternative; brief decides. |
| REQ-RM-0127 | Disk I/O wait (%iowait) from `/proc/stat` deltas, to detect failing microSD/flash. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:372-373 | C-08 | | |
| REQ-RM-0128 | Filesystem free capacity and wear tracking (POSIX `statvfs()`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:374 | C-08 | | `statvfs` gives capacity only; wear needs another source (e.g. eMMC life-time in sysfs): analyst. |
| REQ-RM-0129 | Kernel I/O error trap: `/dev/kmsg` read non-blocking for storage errors (`mmc0: timeout`, `I/O error`, `EXT4-fs error`, squashfs corruption). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:375 | C-08 | C-03, R-05 | |
| REQ-RM-0130 | A streaming secret redactor is applied when dossiers are generated and diagnostic snapshots are shared with outside LLMs or ISP support. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:380-381 | C-04 | R-03, R-04, R-05 | |
| REQ-RM-0131 | Passwords, Wi-Fi PSKs and auth tokens become `[REDACTED_SECRET]`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:382 | C-04 | | |
| REQ-RM-0132 | Telegram bot tokens and chat IDs become `[REDACTED_CHAT_ID]`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:383 | C-04 | S-11 | A token under a chat-id marker looks odd; take exact markers from the Python `common/redact.py` (F-05). |
| REQ-RM-0133 | ISP account IDs and circuit IDs become `[REDACTED_CIRCUIT_ID]`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:384 | C-04 | | |
| REQ-RM-0134 | Dynamic DNS domain names become `[REDACTED_DDNS]`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:385 | C-04 | N-16 | |
| REQ-RM-0135 | Private IP subnets are kept for topology diagnostics; MAC addresses are replaced by their vendor prefix and a pseudonym that stays the same within one bundle, unless the user asks for full MAC addresses for that one export. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:386 | C-04 | R-04 | Q-16 settled (developer's choice): keeps the topology purpose of line 386 without sending device identifiers to outside LLMs or support by default (Quality Bar 8, data minimisation). Restated in round 3. |
| REQ-RM-0136 | Collector client keeps the protocol: `POST /api/ingest` (JSON: `protocol`, `box_uid`, new speed tests, hourly ping summaries; events and diagnostics report only with consent) and `POST /api/poll` (command results in, queued commands out), so one collector serves Python and native boxes. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:391 | S-12 | S-13, F-04 | |
| REQ-RM-0137 | Collector client uses a bearer token and outgoing HTTPS only. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:391 | S-12 | | |
| REQ-RM-0138 | Remote commands stay a fixed list (`speedtest`, `diagnostics`, `upload`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:391 | S-12 | S-13 | |
| REQ-RM-0139 | The collector server is not part of the router build. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:391 | S-13 | P-11, NEW: X2-01 Step-2 router tiers | |
| REQ-RM-0140 | Native `nmb-collector`: one static binary in Go, same code base as `nmbd`, for Linux x86-64 and ARM64 servers. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:392 | S-13 | F-00, P-11 | |
| REQ-RM-0141 | Collector server keeps the protocol (`POST /api/ingest`, `GET /api/commands`, `POST /api/command-result`), the data directory layout and token hashes: a Python collector's data opens in the native one, and Python and native boxes report to either. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:392 | S-13 | S-12, F-04, F-05 | The client paths of line 391 (`/api/poll`) and the server paths here differ; analyst checks both Python files. |
| REQ-RM-0142 | Collector keeps the password-protected pages: search by box ID, name, current or earlier case number, customer number and provider; averages against subscribed speeds; CSV. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:392 | S-13 | R-02 | |
| REQ-RM-0143 | Collector keeps the three consents and the fixed command list. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:392 | S-13 | S-12 | |
| REQ-RM-0144 | Collector serves HTTPS itself (`--cert`, `--key`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:392 | S-13 | | |
| REQ-RM-0145 | Of the Python updater the native build keeps the update status file format (`update_status.json`), so the Web UI and `nmb-cli update` work unchanged. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:393 | P-02 | S-10, U-02, F-04 | Status-file rule only (round 2: narrowed to what the line says). That the source is not downloaded comes from this line for routers (REQ-RM-0146) and, for the Pi and x86-64 boxes, from docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:794 (A/B slots replace setup.sh upgrades, REQ-RM-1281, P-02). |
| REQ-RM-0146 | Routers update by firmware or package (`opkg`), not by downloading the source. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:393 | NEW: X2-01 Step-2 router tiers | P-02 | |
| REQ-RM-0147 | The scheduler keeps `ookla_interval_min` (never below 60), `rtr_interval_min` and `cloudflare_interval_min`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:394 | M-04 | C-01 | |
| REQ-RM-0148 | The Ookla cooldown is persisted (`ookla_cooldown.json`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:394 | M-05 | M-10, F-04 | |
| REQ-RM-0149 | The native build keeps the settings, API endpoints and behaviour of every feature in 4.26. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:397 | P-13 | C-01, S-02 | The source's run on the Pi 4 test box and the record in `docs/TESTING.md` describe the Python reference; the native proof duty comes from QB 1 and QB 5 (REQ-RM-1037, REQ-RM-1048). Restated in round 1. |
| REQ-RM-0150 | Fast.com: token from fast.com's page script with the known token as fallback; server list from `api.fast.com/netflix/speedtest/v2`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:399 | M-07 | | |
| REQ-RM-0151 | Fast.com latency from `range/0-0` on an open connection plus the TCP handshake. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:399 | M-07 | | |
| REQ-RM-0152 | Fast.com: 16 parallel streams over up to five Open Connect servers, 15 s each way, first 3 s left out; "stable" speed = mean from the first sample at 90 % of the median of the second half. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:399 | M-07 | R-01 | |
| REQ-RM-0153 | Fast.com transfers over plain HTTP (`https=false`); server names from the HTTPS list, matched by IP. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:399 | M-07 | | Reason in source: TLS CPU cost at near-gigabit speed on a Pi 4 (source names a specific site value; not copied). |
| REQ-RM-0154 | The native Fast.com engine reaches the line speed on the same box with the same method and keeps the full per-server detail in `raw_json`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:399 | M-07 | C-02, P-13 | Proof on the Pi 4 against the Python result of the same day (P-13). |
| REQ-RM-0155 | Fast.com is part of the mixed run, with `fastcom_interval_min`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:399 | M-04 | M-07, C-01 | |
| REQ-RM-0156 | `[network] measurement_link` = `auto`, `ethernet` or `wifi` (and `usb`, REQ-RM-0169). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:400 | N-08 | M-04, C-01, C-11 | |
| REQ-RM-0157 | Measuring over Wi-Fi: `[wifi_uplink]` with its own connection profile (`nmb-uplink`) and a source rule `from <wifi ip> lookup 150`, so a Wi-Fi-bound test leaves over Wi-Fi even when Ethernet holds the default route; also when the Wi-Fi was joined outside NMB. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:400 | N-08 | N-01, N-07 | The Python profile is a NetworkManager profile; native is NetworkManager-free (line 412). |
| REQ-RM-0158 | Wi-Fi shared on a LAN port: own DHCP server (Python: dnsmasq with `--port=0`), NAT in its own nftables table, the shared subnet in routing table 151. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:401 | N-08 | N-02, N-05 | |
| REQ-RM-0159 | DHCP guard: never serve DHCP where another server answers: a DHCPDISCOVER sent three times, replies read from a packet socket for 8 s, before starting, every minute, and right after a cable is plugged in (own DHCP stopped first). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:401; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:426 | N-05 | N-08, N-10, N-04 | |
| REQ-RM-0160 | Each wireless interface is the guest client or the main client, never both; with `interface = auto` the main client skips the radio of an enabled guest client. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:402 | N-08 | N-07 | |
| REQ-RM-0161 | Storage guard free-space target from disk size: 15 % up to 2 GB, 10 % up to 16 GB, 5 % above; limited to 200 MB - 2 GB. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:403 | C-08 | | |
| REQ-RM-0162 | Storage guard deletes in fixed order: rotated logs, bundles, crash reports; backups but the newest; reports; journal; large active logs; only then measurements, one day at a time, keeping `keep_days`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:403 | C-08 | C-02, P-03 | |
| REQ-RM-0163 | On routers the storage guard applies to the overlay or data partition. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:403 | NEW: X2-01 Step-2 router tiers | C-08 | |
| REQ-RM-0164 | iPerf3 server names from the public iPerf3 list, else city, country and ISP of the address from GeoIP, also for older rows that stored only the address; a failed lookup is not repeated for an hour. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:404 | M-09 | M-12, R-02, U-02 | |
| REQ-RM-0165 | A USB phone, hotspot or modem is recognised by its driver (`rndis_host`, `ipheth`, `cdc_eem`, `huawei_cdc_ncm`, `qmi_wwan`, `cdc_mbim`) or `DEVTYPE=wwan` in `uevent`; for generic `cdc_ether`, `cdc_ncm`, `cdc_subset` unless `idVendor` is a known Ethernet dongle maker (Realtek `0bda`, ASIX `0b95`, Microchip `0424`, D-Link `2001`, Linksys `13b1`, ICS `0fe6`, Lenovo `17ef`); USB Ethernet dongles stay wired. Detection from sysfs, no tools. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:414-415 | N-10 | N-17 | The driver list of line 405 is shorter; line 415 is the full rule. The vendor IDs are public chipset-maker IDs (copied in round 1). |
| REQ-RM-0166 | A tethering interface is never the wired line (never chosen by wired priority, never measured as Ethernet) and never a LAN port (never bridged, never gets the box's DHCP server). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:420 | N-10 | N-01, N-03, N-06, N-05 | |
| REQ-RM-0167 | The interface name never decides (`usb0`, `enx<MAC>`, `eth1`; `eth1` is the default LAN name and often the hotspot). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:416 | N-10 | N-06 | |
| REQ-RM-0168 | The USB link gets its own address and no default route by default: the lease's router goes into routing table 152 with source rule `from <usb address> lookup 152`, not the main table; the box's traffic and the outage probe stay on the wired or Wi-Fi line. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:421 | N-10 | N-01, M-01 | |
| REQ-RM-0169 | Measuring over USB with `measurement_link = usb` or a test aimed at that interface; results stored with `connection_type = fwa` and the mobile public IP; per-engine binding to the USB address. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:422 | N-10 | M-04, C-02, C-11 | |
| REQ-RM-0170 | USB has its own daily data limit (`[usb_uplink] daily_limit_mb`, default 1024 MB, 0 = none), counted on the interface's own byte counters, for every trigger (scheduled and manual). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:423 | N-10 | M-05, C-01 | |
| REQ-RM-0171 | Only internet: with no default route through Ethernet or Wi-Fi, a default route through USB with metric 900, removed as soon as another returns; each switch is an event (`uplink_usb_fallback`, `uplink_usb_fallback_end`, severity `warn`); remote access (Tailscale) keeps working over it. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:424 | N-10 | C-07, N-15, R-01 | |
| REQ-RM-0172 | Router mode on USB (Mode A): when no DHCP server answers on the WAN port (a PC is fine, a router is not), USB becomes the WAN and the Ethernet port joins the LAN bridge with the Wi-Fi AP; checked every minute; a DHCP server appearing on the port stops the box's DHCP at once and makes the port the WAN again; PPPoE or a static WAN always keep the port as WAN. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:425 | N-10 | N-03, N-05, N-07 | Python: sharing USB internet to clients not yet live-tested (line 412). |
| REQ-RM-0173 | Sharing on the LAN port (`wifi_share`) takes USB as its upstream when the Wi-Fi uplink is down, with the same DHCP guard. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:426 | N-08 | N-10, N-05 | |
| REQ-RM-0174 | The USB rules apply on every device with a USB port, routers included (a hotspot as backup WAN), without NetworkManager and without Python. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:405; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:412 | N-10 | NEW: X2-01 Step-2 router tiers | Source names the test device and its carrier (source names a specific site value; not copied). |
| REQ-RM-0175 | Web UI texts are keyed; the native build reads the same keys and JSON files in `nmb/web/locales/` and `nmb/report/locales/` (11 languages) and shows a missing text in English. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:406 | C-06 | U-01, R-03, S-11 | |
| REQ-RM-0176 | Sharing the USB internet to clients has not passed a live test in Python; the native build must prove it live before it is called working. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:412 | N-10 | P-13 | Honest results rule. |
| REQ-RM-0177 | USB hot-plug comes from netlink (`RTM_NEWLINK`/`RTM_DELLINK`), seen at once, not on the next poll. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:417 | N-10 | N-01 | |
| REQ-RM-0178 | `nmbd` runs its own small built-in DHCPv4 client on the USB interface (as elsewhere). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:421 | N-05 | N-10 | |
| REQ-RM-0179 | Raspberry Pi 4, x86-64 and NMB's own distribution: the image includes `rndis_host`, `cdc_ether`, `cdc_ncm`, `ipheth`, `qmi_wwan`, `cdc_mbim` (criterion 1 of 6A); `usbmuxd` for iPhone pairing is the only userspace helper. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:429 | P-11 | N-10 | |
| REQ-RM-0180 | OpenWrt: USB kmods (`kmod-usb-net-rndis`, `-cdc-ether`, `-cdc-ncm`, `-ipheth` + `usbmuxd`; `-qmi-wwan`/`-cdc-mbim` for modems); `nmbd` configures the link itself instead of `netifd`/`uqmi`/`umbim` where possible, or through `ubus` when `netifd` owns the interface. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:430 | NEW: X2-01 Step-2 router tiers | N-10 | |
| REQ-RM-0181 | Flash-limited targets (4 MiB, 8 MiB): only `rndis_host` and `cdc_ether`/`cdc_ncm`; modem drivers and `usbmuxd` optional packages; targets without USB do not have the feature. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:431 | NEW: X2-01 Step-2 router tiers | N-10 | |
| REQ-RM-0182 | MikroTik RouterOS container: no USB access; USB tethering is RouterOS's own and measured as any other WAN. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:432 | NEW: X2-01 Step-2 router tiers | N-10 | |
| REQ-RM-0183 | USB tests: the same as the Python reference (`tests/test_usb_tether.py`, `tests/test_usb_uplink.py`, `tests/test_usb_router.py`) plus a real device per platform: plug in, measure over USB with the byte counters, pull the cable and the Wi-Fi (fallback and back), router mode with a PC on the Ethernet port and a Wi-Fi client, and a home router on the port (the box must stop its DHCP). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:434 | N-10 | P-13, F-06 | Real-device steps on the Pi 4 under work plan 2.2 point 10. |
| REQ-RM-0184 | Automatic mode behaves as the Python `nmb/netd/auto_mode.py` on every device (Python unit-tested, not live-tested: the native build proves it live). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:438 | N-04 | P-13 | |
| REQ-RM-0185 | A router giving a DHCP lease on the WAN port means standalone; with several wired ports the others are bridged to the WAN port (LAN passthrough, transparent bridge). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:440 | N-04 | N-06, N-03 | |
| REQ-RM-0186 | No router lease, no DHCP answer and no internet for 90 seconds (ONT/modem, PC, or no cable) means router mode, until a router answers again. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:441 | N-04 | N-03 | |
| REQ-RM-0187 | In router mode: PPPoE is dialled when credentials are saved and a PPPoE server answers discovery; USB tethering is used when present; the Wi-Fi AP and the LAN ports always serve their own network so the device can be reached. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:441 | N-04 | N-05, N-10, N-07 | |
| REQ-RM-0188 | A DHCP server appearing on the WAN port in router mode switches to standalone at once; the device never hands out addresses into another network. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:442 | N-04 | N-05 | |
| REQ-RM-0189 | Probes: a real DHCP DISCOVER from 0.0.0.0 on the WAN port (read before the bridge sees the frames), a PPPoE discovery, and a TCP connection to the internet; an answer stays valid until the next probe (30 s), so the debounce sees a stable answer. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:443 | N-04 | N-05 | |
| REQ-RM-0190 | An explicit setting always wins over automatic mode: a fixed mode, a manual override, or an owner-configured LAN passthrough or link aggregation. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:444 | N-04 | N-03, N-06, C-01 | |
| REQ-RM-0191 | Every speed test, outage and ping minute states how the box was connected, with the same codes as the Python reference, so data moves between the two. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:448 | C-11 | M-01, M-02, M-04, N-01, N-03, R-01, R-02, R-03, R-04, R-05, S-02, U-02, S-10, S-11, S-12, S-13 | Owner, 2026-11-10. |
| REQ-RM-0213 | The words that describe the measurement setup come from one source, in all 11 languages, and every surface that shows the setup uses them. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:448 | C-11 | C-06, R-03, U-02, S-11 | Python reference `nmb/common/setup_texts.py`; same rule as REQ-CL-0005 and REQ-WP-0072. Added in round 1. |
| REQ-RM-0192 | Three facts with fixed codes: mode (`standalone`, `mode_a`, `mode_b`; for `auto` the resolved mode), link (`ethernet`, `pppoe`, `wifi`, `usb`, `bridge`, `passthrough`, `lag`), path to the ISP (`direct_pppoe`, `direct_public`, `direct_cgnat`, `behind_router`). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:450 | C-11 | F-04, F-07 | |
| REQ-RM-0193 | Provenance of the facts: `recorded` (measured), `changed` (mode switched during the test), `derived` (reconstructed later). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:450 | C-11 | F-04 | |
| REQ-RM-0194 | The position comes from the address, never from the mode; only IPv4 decides; a global IPv6 address proves nothing. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:451 | C-11 | M-12 | Source cites a double-NAT case seen on a real box (source names a specific site value; not copied). |
| REQ-RM-0195 | VPN interfaces (Tailscale, WireGuard) are never the measuring link; Tailscale's 100.64/10 addresses are not the ISP's CGNAT. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:451 | C-11 | M-12, M-04, N-13, N-15 | |
| REQ-RM-0196 | Without an address to bind to, the engines leave through the default route, and that link is described, not the one asked for. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:452 | C-11 | M-04 | |
| REQ-RM-0197 | Unknown stays unknown: the field is empty and every surface says "unknown"; nothing is guessed, nothing filled with today's setup, nothing measured replaced. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:453 | C-11 | R-01, R-02, R-03, R-04, R-05, S-02, U-02, S-10, S-11, S-12 | |
| REQ-RM-0198 | netd records a setup history row whenever mode, link, its kind or its position changes (and at stop); a ping minute or outage is read against the setup of its time; an outage takes the last record before it. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:454 | C-11 | N-03, M-01, M-02 | |
| REQ-RM-0199 | The speed, detailed, ping and outage CSVs carry the setup codes, appended after the old columns. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455 | R-02 | C-11 | |
| REQ-RM-0200 | The SLA dossier has a setup section with the periods, a setup column in the test and outage tables, and a warning for tests behind another router. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455 | R-03 | C-11, C-06 | |
| REQ-RM-0201 | The diagnostic bundle states the setup. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455 | R-04 | C-11 | |
| REQ-RM-0202 | The e-mail states the setup. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455 | R-05 | C-11 | |
| REQ-RM-0203 | The Web UI shows the setup (latest result, history, header). | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455 | U-02 | S-02, C-11 | |
| REQ-RM-0204 | The bot shows the setup in `/test` and `/status`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455 | S-11 | C-11 | |
| REQ-RM-0205 | `nmb-cli` shows the setup. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455 | S-10 | C-11 | |
| REQ-RM-0206 | The collector receives the setup codes only; the box's own address stays on the box. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455 | S-12 | S-13, C-11 | |
| REQ-RM-0207 | A report window with more than one setup says so. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455 | R-01 | R-03, R-04, R-05, U-02 | |
| REQ-RM-0208 | The native build reads the Python database including the setup columns and the `network_state` history. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:456 | C-02 | C-11, F-05 | |
| REQ-RM-0209 | The one-time reconstruction (backfill) of older rows is not repeated. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:456 | C-11 | C-02 | |
| REQ-RM-0210 | All components are unified into one multi-call static binary; no external helper binaries are spawned. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:460; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:85-114 | P-10 | P-13 | |
| REQ-RM-0211 | Multi-call names (symlinks to the binary) keep existing systemd/OpenRC units and user workflows working. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:494-495 | P-10 | P-11, P-04 | |
| REQ-RM-0212 | Multi-call map: `nmb-probe`, `nmb-speed`, `nmb-netd`, `nmb-web`, `nmb-adblock`, `nmb-heal`, `nmb-driver`, `pihole` (to adblock) and `rmbt-client` (to rmbt) each run `nmbd --mode <name>`. | docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:496-504 | P-10 | N-12, M-08, N-17 | Map lacks `nmb-telegram`, `nmb-invasive-log`, `nmb-cli`; collector is its own binary (REQ-RM-0140). `pihole` alias: Q-38 settled, the compatibility names `pihole` and `pivpn` are installed only where no command of that name exists, never over one, and removed only if NMB made them (*Generic First* 3, Quality Bar 7). |

## Coverage

Each range includes the blank line(s) that end it. Every line 1-507 is in exactly one row.

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1 | heading | context: document title |
| 2 | heading | context: subtitle |
| 3-5 | separator | context: horizontal rule |
| 6-7 | heading | context: section 1 title |
| 8 | heading | context: 1.1 title |
| 9-10 | paragraph | REQ-RM-0001 |
| 11-12 | table row | context: table header and alignment row |
| 13 | table row | REQ-RM-0002 |
| 14 | table row | REQ-RM-0003 |
| 15 | table row | REQ-RM-0004 |
| 16 | table row | REQ-RM-0005 |
| 17-18 | table row | REQ-RM-0006 |
| 19 | heading | context: 1.2 title |
| 20-21 | paragraph | REQ-RM-0007, REQ-RM-0008 |
| 22 | paragraph | REQ-RM-0009, REQ-RM-0010 |
| 23-25 | separator | context: horizontal rule |
| 26-27 | heading | context: section 2 title |
| 28-29 | paragraph | context: introduces the language evaluation |
| 30-52 | code block | context: Mermaid evaluation diagram; estimates restated as targets in REQ-RM-0002, -0003, -0013, -0016, REQ-RM-0219 (line 42, C-tier binary size) |
| 53-54 | heading | context: 2.1 title |
| 55 | table row | REQ-RM-0011 |
| 56 | table row | context: alignment row |
| 57 | table row | REQ-RM-0003, REQ-RM-0219 |
| 58 | table row | REQ-RM-0002 |
| 59 | table row | REQ-RM-0003, REQ-RM-0016, REQ-RM-0219 |
| 60 | table row | REQ-RM-0012 |
| 61 | table row | context: concurrency model per language, a design choice with nothing to check |
| 62 | table row | REQ-RM-0032, REQ-RM-0015 |
| 63 | table row | REQ-RM-0058 |
| 64 | table row | REQ-RM-0021 |
| 65 | table row | REQ-RM-0017 |
| 66-68 | separator | context: horizontal rule |
| 69-70 | heading | context: section 3 title |
| 71 | heading | context: tier strategy title |
| 72-74 | bullet | REQ-RM-0013, REQ-RM-0014, REQ-RM-0015 |
| 75-77 | bullet | REQ-RM-0016 |
| 78-80 | separator | context: horizontal rule |
| 81-82 | heading | context: section 4 title |
| 83-84 | paragraph | context: introduces the per-component blueprints |
| 85-115 | code block | REQ-RM-0210 (component map into the unified binary; each component is covered by its own section) |
| 116 | heading | context: 4.1 title |
| 117 | bullet | context: Python bottleneck figures (CPU, RAM), motivation only; the probe cadence it names (5 targets every 500 ms) is taken from the Python reference (M-01 brief), not from this line |
| 118-119 | bullet | REQ-RM-0017 |
| 120-121 | bullet | REQ-RM-0018 |
| 122 | bullet | REQ-RM-0019 |
| 123-127 | bullet | REQ-RM-0020 |
| 128 | heading | context: 4.2 title |
| 129 | bullet | REQ-RM-0022, REQ-RM-0023 |
| 130-132 | bullet | REQ-RM-0021 |
| 133 | bullet | REQ-RM-0022 |
| 134-137 | bullet | REQ-RM-0024 |
| 138 | heading | context: 4.3 title |
| 139 | bullet | REQ-RM-0025 |
| 140-141 | bullet | REQ-RM-0026 |
| 142 | bullet | REQ-RM-0027 |
| 143 | bullet | REQ-RM-0028, REQ-RM-0029, REQ-RM-0030 |
| 144-145 | bullet | REQ-RM-0031 |
| 146 | heading | context: 4.4 title |
| 147 | bullet | context: Python bottleneck, motivation only |
| 148-149 | bullet | REQ-RM-0032 |
| 150 | bullet | REQ-RM-0015 |
| 151 | bullet | REQ-RM-0033 |
| 152 | bullet | REQ-RM-0034 |
| 153-156 | bullet | REQ-RM-0035 |
| 157-158 | heading | context: 4.4A title (owner, 2026-11-10, planned) |
| 159-160 | paragraph | REQ-RM-0036 |
| 161-162 | paragraph | context: measured baseline of the Python UI (sizes, 11 timers, missing translations); the requirements below answer it (REQ-RM-0042, -0047) |
| 163-164 | paragraph | context: "Requirements" label |
| 165 | bullet | REQ-RM-0037 |
| 166 | bullet | REQ-RM-0038, REQ-RM-0039, REQ-RM-0214 |
| 167 | bullet | REQ-RM-0040, REQ-RM-0041 |
| 168 | bullet | REQ-RM-0042, REQ-RM-0043, REQ-RM-0044 |
| 169 | bullet | REQ-RM-0045, REQ-RM-0046 |
| 170 | bullet | REQ-RM-0047 |
| 171 | bullet | REQ-RM-0048, REQ-RM-0049, REQ-RM-0215 |
| 172 | bullet | REQ-RM-0050, REQ-RM-0051, REQ-RM-0216 |
| 173 | bullet | REQ-RM-0052, REQ-RM-0217 |
| 174-175 | bullet | REQ-RM-0053, REQ-RM-0054, REQ-RM-0055, REQ-RM-0056, REQ-RM-0057 |
| 176 | heading | context: 4.5 title |
| 177 | bullet | context: Python bottleneck, motivation only |
| 178-180 | bullet | REQ-RM-0058 |
| 181-182 | bullet | REQ-RM-0059 |
| 183 | heading | context: 4.6 title |
| 184 | bullet | context: current Python architecture (subprocess), replaced by REQ-RM-0060 |
| 185 | bullet | context: target language choice; Go for step 1 per REQ-RM-0010 |
| 186-187 | bullet | REQ-RM-0060 |
| 188-189 | bullet | REQ-RM-0061 |
| 190-191 | bullet | REQ-RM-0062 |
| 192 | bullet | REQ-RM-0060, REQ-RM-0063 |
| 193-195 | bullet | REQ-RM-0064 |
| 196 | heading | context: 4.7 title |
| 197 | bullet | context: current Python architecture, replaced by REQ-RM-0065/-0066 |
| 198 | bullet | context: target language choice |
| 199-201 | bullet | REQ-RM-0065 |
| 202-204 | bullet | REQ-RM-0066 |
| 205 | bullet | REQ-RM-0067 |
| 206-207 | bullet | REQ-RM-0068 |
| 208 | heading | context: 4.8 title |
| 209 | bullet | context: current Python architecture (spawns iperf3) |
| 210 | bullet | context: target language choice |
| 211-212 | bullet | REQ-RM-0069 |
| 213 | bullet | REQ-RM-0070 |
| 214-216 | bullet | REQ-RM-0071 |
| 217 | heading | context: 4.9 title |
| 218 | bullet | REQ-RM-0072 |
| 219 | bullet | context: target language choice |
| 220-221 | bullet | REQ-RM-0073 |
| 222-224 | bullet | REQ-RM-0074 |
| 225 | bullet | REQ-RM-0075 |
| 226-227 | bullet | REQ-RM-0076 |
| 228 | heading | context: 4.10 title |
| 229 | bullet | context: current Python architecture (wg-quick, iptables, systemctl, qrencode, tailscale CLI), replaced by REQ-RM-0077 to -0081 |
| 230 | bullet | context: target language choice |
| 231-235 | bullet | REQ-RM-0077 |
| 236 | bullet | REQ-RM-0078 |
| 237 | bullet | REQ-RM-0079 |
| 238 | bullet | REQ-RM-0080 |
| 239 | bullet | REQ-RM-0081 |
| 240-243 | bullet | REQ-RM-0082 |
| 244 | heading | context: 4.11 title |
| 245 | bullet | REQ-RM-0083, REQ-RM-0085 |
| 246 | bullet | context: target language choice |
| 247-248 | bullet | REQ-RM-0083 |
| 249 | bullet | REQ-RM-0084 |
| 250 | bullet | REQ-RM-0085 |
| 251-252 | bullet | REQ-RM-0086 |
| 253 | heading | context: 4.12 title |
| 254 | bullet | REQ-RM-0087 |
| 255 | bullet | context: target language choice |
| 256-257 | bullet | REQ-RM-0088 |
| 258 | bullet | REQ-RM-0089 |
| 259 | bullet | REQ-RM-0090 |
| 260-261 | bullet | REQ-RM-0091 |
| 262 | heading | context: 4.13 title |
| 263 | bullet | REQ-RM-0092 |
| 264 | bullet | context: target language choice |
| 265-271 | bullet | REQ-RM-0093 |
| 272 | bullet | REQ-RM-0094 |
| 273-274 | bullet | REQ-RM-0095 |
| 275 | heading | context: 4.14 title |
| 276 | bullet | context: current Python architecture (edits wpa_supplicant.conf, spawns wpa_cli/hostapd_cli) |
| 277 | bullet | context: target choice |
| 278-279 | bullet | REQ-RM-0096 |
| 280 | bullet | REQ-RM-0097 |
| 281-282 | bullet | REQ-RM-0098 |
| 283 | heading | context: 4.15 title |
| 284 | bullet | REQ-RM-0099 |
| 285 | bullet | context: target choice |
| 286-287 | bullet | REQ-RM-0100 |
| 288 | bullet | REQ-RM-0101 |
| 289-290 | bullet | REQ-RM-0102 |
| 291 | heading | context: 4.16 title |
| 292 | bullet | context: current Python architecture |
| 293 | bullet | context: target language choice |
| 294-295 | bullet | REQ-RM-0103 |
| 296-300 | bullet | REQ-RM-0104 |
| 301 | bullet | REQ-RM-0105 |
| 302-303 | bullet | REQ-RM-0106, REQ-RM-0218 |
| 304 | heading | context: 4.17 title |
| 305 | bullet | context: current Python architecture |
| 306 | bullet | context: target choice |
| 307-308 | bullet | REQ-RM-0107 |
| 309 | bullet | REQ-RM-0108 |
| 310 | bullet | REQ-RM-0109 |
| 311-312 | bullet | REQ-RM-0110 |
| 313 | heading | context: 4.18 title |
| 314-315 | bullet | context: current architecture (post hoc) and target choice |
| 316-319 | bullet | REQ-RM-0111 |
| 320 | bullet | REQ-RM-0112 |
| 321-322 | bullet | REQ-RM-0113 |
| 323 | heading | context: 4.19 title |
| 324 | bullet | REQ-RM-0115 |
| 325 | bullet | context: target choice |
| 326-332 | bullet | REQ-RM-0114 |
| 333-334 | bullet | REQ-RM-0115 |
| 335 | heading | context: 4.20 title |
| 336-337 | bullet | context: current architecture (single getaddrinfo) and target choice |
| 338-343 | bullet | REQ-RM-0116 |
| 344-345 | bullet | REQ-RM-0117 |
| 346 | bullet | REQ-RM-0118 |
| 347-348 | bullet | REQ-RM-0119 |
| 349 | heading | context: 4.21 title |
| 350-351 | bullet | context: current architecture (heuristics) and target choice |
| 352-353 | bullet | REQ-RM-0120 |
| 354 | bullet | REQ-RM-0121 |
| 355-358 | bullet | REQ-RM-0122 |
| 359-360 | bullet | REQ-RM-0123 |
| 361 | heading | context: 4.22 title |
| 362-363 | bullet | context: current state (missing) and target choice |
| 364-365 | bullet | REQ-RM-0124 |
| 366 | bullet | REQ-RM-0125 |
| 367-368 | bullet | REQ-RM-0126 |
| 369 | heading | context: 4.23 title |
| 370-371 | bullet | context: current architecture (shutil.disk_usage) and target choice |
| 372-373 | bullet | REQ-RM-0127 |
| 374 | bullet | REQ-RM-0128 |
| 375-376 | bullet | REQ-RM-0129 |
| 377 | heading | context: 4.24 title |
| 378-379 | bullet | context: current architecture (regex filters) and target choice |
| 380-381 | bullet | REQ-RM-0130 |
| 382 | bullet | REQ-RM-0131 |
| 383 | bullet | REQ-RM-0132 |
| 384 | bullet | REQ-RM-0133 |
| 385 | bullet | REQ-RM-0134 |
| 386 | bullet | REQ-RM-0135 |
| 387-389 | separator | context: horizontal rule |
| 390 | heading | context: 4.25 title |
| 391 | bullet | REQ-RM-0136, REQ-RM-0137, REQ-RM-0138, REQ-RM-0139 |
| 392 | bullet | REQ-RM-0140, REQ-RM-0141, REQ-RM-0142, REQ-RM-0143, REQ-RM-0144 |
| 393 | bullet | REQ-RM-0145, REQ-RM-0146 |
| 394-395 | bullet | REQ-RM-0147, REQ-RM-0148 |
| 396 | heading | context: 4.26 title |
| 397-398 | paragraph | REQ-RM-0149 |
| 399 | bullet | REQ-RM-0150, REQ-RM-0151, REQ-RM-0152, REQ-RM-0153, REQ-RM-0154, REQ-RM-0155 |
| 400 | bullet | REQ-RM-0156, REQ-RM-0157 |
| 401 | bullet | REQ-RM-0158, REQ-RM-0159 |
| 402 | bullet | REQ-RM-0160 |
| 403 | bullet | REQ-RM-0161, REQ-RM-0162, REQ-RM-0163 |
| 404 | bullet | REQ-RM-0164 |
| 405 | bullet | REQ-RM-0165, REQ-RM-0166, REQ-RM-0167, REQ-RM-0168, REQ-RM-0169, REQ-RM-0170, REQ-RM-0171, REQ-RM-0172, REQ-RM-0173, REQ-RM-0174 (test device not copied) |
| 406 | bullet | REQ-RM-0175 |
| 407-409 | separator | context: horizontal rule |
| 410-411 | heading | context: 4.27 title |
| 412-413 | paragraph | REQ-RM-0174, REQ-RM-0176 (test device not copied) |
| 414 | paragraph | REQ-RM-0165 (label: detection from sysfs, no tools) |
| 415 | bullet | REQ-RM-0165 |
| 416 | bullet | REQ-RM-0167 |
| 417-418 | bullet | REQ-RM-0177 |
| 419 | paragraph | context: "Rules that must hold everywhere" label |
| 420 | bullet | REQ-RM-0166 |
| 421 | bullet | REQ-RM-0168, REQ-RM-0178 |
| 422 | bullet | REQ-RM-0169 |
| 423 | bullet | REQ-RM-0170 |
| 424 | bullet | REQ-RM-0171 |
| 425 | bullet | REQ-RM-0172 |
| 426-427 | bullet | REQ-RM-0173, REQ-RM-0159 |
| 428 | paragraph | context: "Per platform" label |
| 429 | bullet | REQ-RM-0179 |
| 430 | bullet | REQ-RM-0180 |
| 431 | bullet | REQ-RM-0181 |
| 432-433 | bullet | REQ-RM-0182 |
| 434-435 | paragraph | REQ-RM-0183 |
| 436-437 | heading | context: 4.28 title |
| 438-439 | paragraph | REQ-RM-0184 |
| 440 | bullet | REQ-RM-0185 |
| 441 | bullet | REQ-RM-0186, REQ-RM-0187 |
| 442 | bullet | REQ-RM-0188 |
| 443 | bullet | REQ-RM-0189 |
| 444-445 | bullet | REQ-RM-0190 |
| 446-447 | heading | context: 4.29 title |
| 448-449 | paragraph | REQ-RM-0191, REQ-RM-0213 |
| 450 | bullet | REQ-RM-0192, REQ-RM-0193 |
| 451 | bullet | REQ-RM-0194, REQ-RM-0195 (site event not copied) |
| 452 | bullet | REQ-RM-0196 |
| 453 | bullet | REQ-RM-0197 |
| 454 | bullet | REQ-RM-0198 |
| 455 | bullet | REQ-RM-0199, REQ-RM-0200, REQ-RM-0201, REQ-RM-0202, REQ-RM-0203, REQ-RM-0204, REQ-RM-0205, REQ-RM-0206, REQ-RM-0207 |
| 456-457 | bullet | REQ-RM-0208, REQ-RM-0209 |
| 458-459 | heading | context: section 5 title |
| 460-461 | paragraph | REQ-RM-0210 |
| 462-493 | code block | context: architecture diagram of section 5; its contents are stated in REQ-RM-0034 (port 80/8080), REQ-RM-0059 (RAM SQLite) and REQ-RM-0210; "in-memory channels" is a design choice; the "QoS" in the netd box (line 471) is a diagram label whose function the roadmap specifies nowhere else, so it holds no requirement (the Python bandwidth monitor page, tab id `tab-qos`, reaches U-15 through `parts/python-files.md`) |
| 494 | heading | REQ-RM-0211 |
| 495 | paragraph | REQ-RM-0211 |
| 496-504 | bullet | REQ-RM-0212 |
| 505-507 | separator | context: horizontal rule |

## Proposed package changes

1. **NEW: X2-01 Step-2 router tiers (C11/musl and OpenWrt/RouterOS integration).** Work plan section 1 says the C11 tier and everything router-specific start with step 2, and section 3 has no package for them, yet the roadmap states requirements for them. Without an owner, the register's check would fail. Rows: REQ-RM-0008, -0016, -0020, -0024, -0035, -0064, -0068, -0071, -0082, -0086, -0091, -0146, -0163, -0180, -0181, -0182, -0214 to -0219 (owner); also in REQ-RM-0010, -0013, -0139, -0174. Round 2: the router and C11 halves of REQ-RM-0038, -0049, -0050, -0052 and -0106 were split off as REQ-RM-0214 to -0218, REQ-RM-0068 (router-only OONI configuration) moved here from M-13, and REQ-RM-0219 (C-tier binary size) was added. Proposal: keep X2-01 as one holding package in step 1 (its requirements go into `spec/` through F-07 so step 2 starts from them). When step 2 is planned, split it per component in the same way as the step-1 packages (probe, netd, web, RMBT, iPerf3, VPN, driver finder, updater, storage guard, USB tethering on OpenWrt, RouterOS container).
2. **Optional split of M-12** into M-12a "DNS resolver matrix and hijack check" (4.20: REQ-RM-0116 to -0119) and M-12b "STUN, NAT type and WAN topology, public address" (4.21: REQ-RM-0120 to -0123). Together with `netd/dns.py` and `netd/geoip.py` these are two protocols, above the one-protocol size rule of section 3.
3. **Optional split of N-10** (17 requirements from 4.26 and 4.27) into N-10a "USB detection, uplink, table 152, data limit, only-internet fallback" (REQ-RM-0165 to -0171, -0176, -0177) and N-10b "Router mode on USB and LAN-port sharing upstream" (REQ-RM-0172, -0173), if the analyst finds `netd/usb_uplink.py` plus tests above the size rule.
4. **Optional NEW: M-17 Shared HTTP transfer core** (multi-stream worker pool and per-socket interface binding, REQ-RM-0026, -0027), used by M-06, M-07, M-10 and M-04, so the binding rule is written and tested once. If not taken, the rows stay with M-06 and M-04 as assigned.

## Questions for the owner

*Triaged in round 3 (2026-11-11), answers and choices in `QUESTIONS.md`:* 1 → Q-18, 2 → Q-37 (answered); 3 → Q-16, 5 → Q-23, 6 → Q-38, 7 → Q-03 (developer's choices); 4 → Q-15 (for the owner). The rows were updated where an answer or a choice settled them; the text below is the question as the part asked it.

1. **Login method (4.4, line 151):** the roadmap names HTTP Basic Auth; the work plan's S-01 has login, sessions and CSRF. The work plan says the roadmap wins where they differ. Should the native build use the session login (Basic Auth line out of date), or Basic Auth as written?
2. **Languages of the bot and the dossier (lines 301, 308):** they list 5 languages (HU, EN, DE, FR, UK), but line 406 says the report locales have 11 and the Web UI has 11. Are the 5-language lists out of date (so the bot and dossier use all 11), or intended?
3. **MAC addresses in shared diagnostics (line 386):** the redactor keeps MAC addresses in bundles that go to outside LLMs or ISP support. Is that intended under Quality Bar 8, or should MACs be masked (for example, keeping the vendor prefix only)?
4. **QR delivery over Telegram (line 300):** a WireGuard QR code holds the client's private key, and the menu would send it through Telegram's servers. Is that intended, or should the bot only point to the Web UI for the QR?
5. **NAT type classifier (line 359):** a plain RFC 5389 Binding Request against the listed servers cannot tell full cone from restricted cone. That needs RFC 5780 support or several servers. May the classifier report "unknown" when no such server answers, or must one be provided?
6. **`pihole` alias (line 503):** should `nmbd` install a `/usr/local/bin/pihole` name? It could clash with a real Pi-hole installation or with scripts that expect Pi-hole's own CLI.
7. **Raspberry Pi Zero and Pi 1 (line 73):** these are ARMv6, and the step-1 build matrix has arm v7 only. Should `linux/arm` v6 be added to the matrix, or are Pi Zero and Pi 1 out of step 1?
