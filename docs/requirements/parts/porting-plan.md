# Requirements register, part: porting plan

Source: `docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md` of the Python repository at commit `ad3758c` (700 lines; this part covers lines 1-700, the whole file; the working copy was checked identical to the commit with `git show ad3758c:...`).
Read in pages: 1-150, 151-300, 301-450, 451-600, 600-700 (all lines). Package ids from `docs/NATIVE_WORK_PLAN.md@ad3758c:98-219` (section 3), read with sections 1-2 (lines 1-97) for the step-1 scope and sections 4-7 (lines 220-343) for the image packages I-01..I-10 and the owner's decisions.
Id range of this part: REQ-PP-0001 to REQ-PP-0199 (0194-0198 added in round 4, the router or C halves split from step-1 rows; 0199 from the owner's answer to Q-17 of 2026-11-11). "Source" cells give the file path in full.

Reading notes for the register owner:
- The whole plan is router work (step 2 and later), except where it states rules for every device (sections 7D, 7E, 7F and parts of 7, 7A, 7B). Router-only rows go to `NEW: X2-01 Step-2 router tiers`, the same holding package that the roadmap part (`roadmap-1-507.md`) proposes; see "Proposed package changes".
- The plan names router models; these are the product's port targets, not the owner's site data, and are kept. Example device lists that may include a device of the owner's site are not copied.
- Several rows are superseded or contradicted by later decisions (Ookla, bot languages, Python reference); this is said in Notes, never by changing the requirement text.

## Requirements

| Id | Requirement | Source | Package | Also in | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| REQ-PP-0001 | All binary sizes, memory figures, CPU loads and throughput values of the porting plan are design targets, not measurements; none is reported as reached until it was measured on the target (no router build exists to measure). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:4; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:673 | NEW: X2-01 Step-2 router tiers | P-13 | Honest results rule; work plan 2.2 point 8. |
| REQ-PP-0002 | Order of work: step 1 (the generic x86-64 and ARM images and the complete native program) comes first; NMB is today developed and run only on the Raspberry Pi 4, and no router software or firmware exists yet. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:4; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:10 | P-13 | NEW: X2-01 Step-2 router tiers, P-11 | Work plan 1: nothing router-specific is written before step 1 is done. |
| REQ-PP-0003 | The first router target is the TP-Link TL-WR1043ND v1; its firmware is unpacked and analysed while it is being programmed. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:10 | NEW: X2-01 Step-2 router tiers | I-10 | The roadmap part `roadmap-508-1034.md` proposes `NEW: T-01` for this target; see Proposed package changes. |
| REQ-PP-0004 | The other targets of the plan follow only after the first router, in an order the owner sets at that time. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:10 | NEW: X2-01 Step-2 router tiers | | Not a question now: the owner sets it later by the source's own words. |
| REQ-PP-0005 | Every target uses the optimised Web UI of roadmap 4.4A with its router size budget (the whole UI at most 150 KB compressed). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:10 | U-01 | NEW: X2-01 Step-2 router tiers | Same budget as REQ-RM-0038. |
| REQ-PP-0006 | 100 % de-pythonization: zero Python runtime, zero external Python libraries, zero interpreted dependencies; no device runs Python. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:12; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:110; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:494; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:638 | P-13 | P-10, P-11 | The Asuswrt hooks of 4.2 are `/bin/sh` scripts of the host firmware (lines 215-227); read as integration glue, not as an interpreted dependency of NMB's logic. |
| REQ-PP-0007 | NMB is re-engineered as a native single-binary appliance `nmbd`, in Go for Tier 1. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:14 | P-10 | F-07 | Round 4: split; the C11/musl half (Tiers 2 and 3) is REQ-PP-0197 (X2-01). |
| REQ-PP-0197 | The native `nmbd` is written in C11 with musl libc for Tiers 2 and 3. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:14 | NEW: X2-01 Step-2 router tiers | P-10, F-07 | Split from REQ-PP-0007 in round 4; the C code base starts in step 2 from the shared spec (F-07, roadmap 6B.8). |
| REQ-PP-0008 | Tier 1 (ARM64 flagship router, ASUS RT-AXE7800 class: 512 MB RAM, 256 MB NAND): Go 1.23+ or C11 monolith on Asuswrt, Asuswrt-Merlin or OpenWrt (`bcm4908`); storage in RAM tmpfs with NAND flash persistence; 2.5 Gbps line-rate saturation. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:18-21; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:33; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:56-57 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0009 | Tier 2 (MIPS little-endian SOHO router, ASUS RT-N12+ class: 32 MB RAM, 8 MB SPI flash): pure C11 musl monolith on OpenWrt 23.05 (`ramips/mt7620`, profile `asus_rt-n12p`); storage in RAM tmpfs plus the SPI flash overlay; 100 Mbps line rate under 14 % CPU. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:23-26; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:34; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:56-57; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:266 | NEW: X2-01 Step-2 router tiers | I-10 | The downloaded set (work plan 5.4) is OpenWrt 25.12.6; the OpenWrt release per target is fixed when its port starts. |
| REQ-PP-0010 | Tier 3 (MIPS big-endian legacy router, TP-Link TL-WR740N v4 class: 32 MB RAM, 4 MiB flash): micro-C11 (musl, `-Os`, UPX LZMA) on a minimal stripped OpenWrt `ath79/generic` (or custom Buildroot); 100 % RAM-only circular ring buffer storage; 100 Mbps line rate under 22 % CPU. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:28-31; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:35; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:56-57; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:334 | NEW: X2-01 Step-2 router tiers | I-10 | |
| REQ-PP-0011 | Tier 2 and Tier 3 CPUs have no FPU: their builds use soft-float (`-msoft-float`); Tier 1 may use the hardware FPU and NEON. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:48 | NEW: X2-01 Step-2 router tiers | F-00 | The step-1 Go cross builds for mips/mipsle need `GOMIPS=softfloat` to run on these CPUs (work plan 2.4 builds them). |
| REQ-PP-0012 | The NMB installation fits the flash available to it: Tier 1 more than 100 MB writable (`/jffs`), Tier 2 about 2.1 MB on the SquashFS overlay, Tier 3 under 650 KB (compressed). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:52 | NEW: X2-01 Step-2 router tiers | | Tier 3 figures disagree inside the plan (576 KB overlay at line 342, 398 KB left at line 349, about 750 KB at line 375); to be measured on the built image. |
| REQ-PP-0013 | Executable size targets: Tier 1 about 12 MB (Go) or about 1.5 MB (C11); Tier 2 about 850 KB (C11), about 320 KB with UPX; Tier 3 under 380 KB (C11), about 178 KB with UPX; measured and recorded. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:58 | NEW: X2-01 Step-2 router tiers | P-13 | Design targets (REQ-PP-0001). |
| REQ-PP-0014 | Active memory envelope targets: Tier 1 18-24 MB RSS, Tier 2 2.5-3.8 MB RSS, Tier 3 1.2-1.8 MB RSS; measured and recorded. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:59 | NEW: X2-01 Step-2 router tiers | P-13 | |
| REQ-PP-0015 | Primary data storage per tier: Tier 1 SQLite WAL in RAM with flash sync; Tier 2 in-memory SQLite or a flat ring buffer; Tier 3 zero flash writes, RAM-only ring buffer. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:60 | NEW: X2-01 Step-2 router tiers | C-02 | On-disk formats come from F-04 (same on every architecture). |
| REQ-PP-0016 | The firmware basis of each additional target is confirmed before its port: TL-WR1043ND v1 OpenWrt (target to confirm), Archer C6 v2 OpenWrt `ath79`, RT-AC57U v1 OpenWrt `ramips/mt7621` (to confirm; WikiDevi lists Padavan as work in progress). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:77 | NEW: X2-01 Step-2 router tiers | I-10 | |
| REQ-PP-0017 | Planned implementation of the additional targets: TL-WR1043ND v1 C11/musl under the same constraints as the 32 MiB Tier 2 target; Archer C6 v2 C11/musl with the same 8 MiB flash limit; RT-AC57U v1 C11/musl or Go. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:78 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0018 | Gigabit ports are not presented as a promise of gigabit test throughput; the CPU limit (for example gigabit tests on an 880 MHz MIPS) is measured on real hardware. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:78; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:80 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0019 | Hardware revisions are told apart: the Archer C6 target covers v2.x only (v3.0/v3.2 are MediaTek hardware; several C6 revisions exist), the RT-AC57U target covers v1 only (v2 uses Qualcomm chips). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:80 | NEW: X2-01 Step-2 router tiers | N-17, P-02 | Firmware upload must check model and revision (roadmap 6B.3). |
| REQ-PP-0020 | Before work on any model starts, every hardware figure is checked against WikiDevi and the OpenWrt Table of Hardware; the device-family figures are from memory and unverified until then. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:85; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:94-97; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:103-108 | NEW: X2-01 Step-2 router tiers | I-10 | Covers all hardware fact rows of sections 1.1, 1.2, 4.1, 5.1, 6.1. |
| REQ-PP-0021 | For MikroTik devices both ways onto the device (RouterOS v7 container, OpenWrt instead of RouterOS) are compared before one is chosen. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:87 | NEW: X2-01 Step-2 router tiers |  |  |
| REQ-PP-0022 | RouterOS container way: `nmbd` runs as a small OCI image next to RouterOS (container package on ARM, ARM64 and x86 builds, not on MIPS); to verify: enabling the container in device mode (physical confirmation), external storage for the image, the container sees only a virtual interface so its measurements go through RouterOS routing and queues, and memory limits per model. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:89; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:94-97 | NEW: X2-01 Step-2 router tiers | P-11, C-11 | Measurements through RouterOS routing and queues belong in the measurement setup (roadmap 4.29). |
| REQ-PP-0023 | OpenWrt way on MikroTik: models supported by OpenWrt are installed by netboot and NMB runs as on any OpenWrt target with full access to ports and radios; the comparison records that the owner loses RouterOS on that device. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:90 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0024 | MediaTek SoCs are ported per chip, so one port covers the many devices that share it. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:99 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0025 | Where hardware flow offload (MediaTek PPE, WED for Wi-Fi) forwards traffic past the CPU, NMB reads the counters where offloaded traffic is still counted, or turns offload off during a test (as for the Broadcom Flow Cache). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:99 | NEW: X2-01 Step-2 router tiers | M-03, M-04 | |
| REQ-PP-0026 | Planned implementation per MediaTek SoC: MT7620N/A and MT7628 C11/musl with Tier 2 limits; MT7621A/AT C11/musl or Go (`GOARCH=mipsle`, soft float) with gigabit tests to be measured; MT7622, MT7981B, MT7986A and MT7988A Go static binary. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:103-108 | NEW: X2-01 Step-2 router tiers | | Example device column not copied (source names example products, one of which may be a device of the owner's site; not copied). |
| REQ-PP-0027 | Generic First: every device family uses the same code as on the PC and the Pi 4; only the backends (service control, firewall, Wi-Fi, flash layout, offload) differ. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:110 | F-04 | NEW: X2-01 Step-2 router tiers | |
| REQ-PP-0028 | Runtime pauses (interpreter lock, garbage collection) must not put latency spikes into sub-millisecond jitter, bufferbloat tracking or microsecond ICMP timing. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:125-126 | M-01 | M-04, R-01 | Go has a garbage collector too: proof is the timestamp spread under load (same as REQ-RM-0005). |
| REQ-PP-0029 | `nmbd` does not spawn `ip`, `ps`, `wpa_cli` or similar tools for its network, process and Wi-Fi work: on 400/575 MHz MIPS CPUs each spawn drives the CPU to 100 % and causes false packet drops and interface congestion. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:127-128 | P-13 | N-01, N-07, C-09, M-14 | Round 4: restated to what lines 127-128 give (the problem of spawning these tools, the reason for REQ-RM-0006 and REQ-PP-0184); the round-2 exception list is removed, it is in neither source line. **Conflict:** line 686 (REQ-PP-0184) sets 0 forks per minute on every tier, while REQ-RM-1350 runs Ookla's program, ooniprobe, iPerf3 and other foreign engines as sandboxed programs and REQ-WP-0179 starts Ookla's own program: Q-40 (answered from owner decision 3 and 6B.5). Each package proves its own part; P-13 checks the whole binary. |
| REQ-PP-0030 | The whole suite is one multicall binary; its Tier 1 variant is `nmbd` (Go). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:134-137 | P-10 | | Step 1 builds `nmbd` only. Round 4: split; the C11 variants are REQ-PP-0198 (X2-01). |
| REQ-PP-0198 | The C11 variants of the multicall binary: `nmbd` for Tier 1 (line 135: "Go / C11"), `nmbd-tiny` (Tier 2) and `nmbd-micro` (Tier 3). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:134-137 | NEW: X2-01 Step-2 router tiers | P-10 | Split from REQ-PP-0030 in round 4. |
| REQ-PP-0031 | Event loop: Linux `epoll`, `timerfd`, `signalfd` and non-blocking sockets. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:143 | NEW: X2-01 Step-2 router tiers | P-10 | For Go the runtime's poller does this; the C tiers implement it. |
| REQ-PP-0032 | Probe engine: direct raw ICMP/UDP sockets with Welford running statistics. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:144 | M-01 | P-04 | Raw sockets need `CAP_NET_RAW` (P-04); the roadmap also allows unprivileged ICMP datagram sockets (REQ-RM-0017). |
| REQ-PP-0033 | Speed engine: zero-copy HTTP/1.1, HTTP/2 and RMBT TCP multi-stream transfers. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:145; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:192-201 | M-04 | M-06, M-07, M-08, M-10 | |
| REQ-PP-0034 | Dynamic routing through a direct netlink socket (`RTM_NEWLINK`, `RTM_NEWROUTE`). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:146; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:192-201 | N-01 | N-03 | |
| REQ-PP-0035 | Wi-Fi telemetry through a direct nl80211 netlink socket (RSSI, MCS, PHY rate, retries). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:147; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:192-201 | M-14 | N-07 | |
| REQ-PP-0036 | Telemetry and tasks read directly from `/proc` (`stat`, `meminfo`, `status`). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:148 | C-09 | | |
| REQ-PP-0037 | DNS forwarder and ad-blocker with an in-memory compact radix tree or Bloom filter. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:149 | N-12 | NEW: X2-01 Step-2 router tiers | Per tier: REQ-PP-0077, -0078. |
| REQ-PP-0038 | STUN and NAT: direct RFC 5389 UDP client for CGNAT and mapping detection. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:150 | M-12 | | Detail: REQ-PP-0075, -0076. |
| REQ-PP-0039 | Web server: embedded HTTP engine serving pre-gzipped SVG charts and UI. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:151; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:192-201 | S-01 | U-01, R-03 | |
| REQ-PP-0040 | Telegram bot: direct TLS long-polling worker with a multi-language message matrix (source: 5 languages). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:152; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:192-201 | S-11 | C-06 | The Python bot has 5 core languages (`nmb/bot/i18n.py@ad3758c:27`) plus 6 locale files (`nmb/bot/locales/`); the native bot follows C-06 (11 languages, roadmap 4.26). |
| REQ-PP-0041 | Storage engine: RAM-only binary circular ring buffer with optional SQLite sync. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:153 | NEW: X2-01 Step-2 router tiers | C-02 | On the PC/Pi class C-02 keeps the RAM live copy with snapshots (roadmap 4.5); the ring buffer is the small-router form. Owner moved from C-02 to X2-01 in round 1 (router form, step 2: docs/NATIVE_WORK_PLAN.md@ad3758c:21-22); C-02 is Also in for its step-1 choice. |
| REQ-PP-0042 | Tier 1 runs on stock Asuswrt (Linux 4.19+, USB Entware or `/jffs` user scripts), Asuswrt-Merlin (GNUton fork for the BCM6756 platform) and OpenWrt (`bcm4908`/`broadcom`). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:177-180 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0043 | Tier 1 CPU use plan: core 0 kernel and IRQs, core 1 speed workers, core 2 Web UI, bot and probe, core 3 CAKE SQM and routing. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:182-191; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:202-206 | NEW: X2-01 Step-2 router tiers | C-09 | Design; the effect on measured jitter is to be measured. |
| REQ-PP-0044 | On Asuswrt-Merlin `nmbd` is integrated without touching read-only flash partitions, using the persistent `/jffs` partition. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:209 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0045 | Asuswrt-Merlin file layout: binary `/jffs/nmb/nmbd` (or `/opt/bin/nmbd` via Entware), configuration `/jffs/configs/nmb/nmb.conf`. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:210-211 | NEW: X2-01 Step-2 router tiers | C-01 | |
| REQ-PP-0046 | Asuswrt-Merlin database: the active database lives in tmpfs RAM (`/tmp/nmb/live.db`); the persistent copy (`/jffs/nmb/persist/nmb.db`) is written every 30 minutes to prevent NAND wear. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:212-214; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:192-201 | NEW: X2-01 Step-2 router tiers | C-02 | |
| REQ-PP-0047 | Lifecycle hook `init-start` mounts the tmpfs directories `/tmp/nmb/run` and `/tmp/nmb/live`. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:215-216 | NEW: X2-01 Step-2 router tiers | P-10 | |
| REQ-PP-0048 | Lifecycle hook `nat-start` sets the forwarding rules (iptables/nftables) for WireGuard, OpenVPN and the guest Wi-Fi probe isolation. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:217 | NEW: X2-01 Step-2 router tiers | N-02, N-08, N-13, N-14 | |
| REQ-PP-0049 | Lifecycle hook `services-start` starts the daemon (`nmbd --daemon --config <path>`) with a watchdog run every 2 minutes from the router's cron (`nmbd --watchdog`), output to a log in `/tmp/nmb/`. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:218-226 | NEW: X2-01 Step-2 router tiers | P-01, P-10 | The example appends to `/tmp/nmb/nmbd.log` without a bound; in RAM tmpfs that conflicts with "logs never fill the disk" (REQ-PP-0153): the log must be bounded. |
| REQ-PP-0050 | Lifecycle hook `services-stop` shuts down gracefully and flushes the database to `/jffs/nmb/persist/`. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:227 | NEW: X2-01 Step-2 router tiers | C-02 | |
| REQ-PP-0051 | Broadcom Flow Cache awareness: because acceleration can stop `/proc/net/dev` counters and let tests bypass software shaping, `nmbd` queries the Flow Cache status (`ioctl` or `/proc/fcache/`) before a line-rate speed test or CAKE calibration and temporarily bypasses acceleration for the measuring interface through the interface directly (never a shell), restoring the previous state after the test, also after a crash or restart. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:231-239 | NEW: X2-01 Step-2 router tiers | M-03, M-04 | The example uses `system("fc config --flush")`, a shell call against REQ-PP-0029 and Quality Bar 7; restoring acceleration after the test is not stated. Q-24 settled (developer's choice: REQ-PP-0029 and Quality Bar 7, no shell; the box's own changes are undone and marked as its own action, 4.29). Restated in round 3. |
| REQ-PP-0052 | On Tier 1 the raw measuring sockets are bound to the 2.5 Gbps interface with `SO_BINDTODEVICE`; the claim of no Flow Cache drops at full line rate is proven by measurement. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:240 | NEW: X2-01 Step-2 router tiers | M-04 | Source says "guaranteeing"; not measured. |
| REQ-PP-0053 | On the 2.5 Gbps WAN `nmbd` configures CAKE (`sch_cake`) multi-queue qdiscs dynamically (example: bandwidth 2400mbit, diffserv4, nat, dual-srchost, ack-filter, memlimit 64M). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:243-246 | NEW: X2-01 Step-2 router tiers | N-03, C-11 | Shaping the measured line changes results. Q-25 (round 4): the shaper and its rate are stored in the result's measurement setup and the result is marked as measured through it (developer's choice: 4.29 and *Measurement honesty*). Whether the shaper is lifted during the box's own tests changes the user's network (REQ-PP-0186) and is asked when X2-01 starts (or offered as a setting). |
| REQ-PP-0054 | On Tier 1 `nmbd` assigns RPS/XPS CPU affinity across all 4 cores (example: `rps_cpus` mask `f`). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:247-249 | NEW: X2-01 Step-2 router tiers | C-09 | |
| REQ-PP-0055 | Tier 2 `nmbd-tiny`: binary 850 KB (320 KB with UPX), 2.8 MB RSS, no dynamic heap fragmentation; parts: raw ICMP ping loop (`epoll`, no allocation), 100 Mbps HTTP/RMBT client, flat ring buffer in `/tmp`, embedded web server (Mongoose or uhttpd CGI), OpenWrt `ubus` and `uci` IPC. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:268-286 | NEW: X2-01 Step-2 router tiers | | The diagram's free RAM (~10 MB) differs from the table (12-14 MB, line 50). |
| REQ-PP-0056 | `nmbd-tiny` is installed at `/overlay/usr/bin/nmbd-tiny` and leaves more than 1.4 MB of the overlay for configuration and TLS certificates. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:298-301 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0057 | Tier 2 SQLite (amalgamation compiled in) uses a minimal configuration so it cannot cause OOM: page size 1024 bytes, cache 50 pages (about 50 KB), temp store `MEMORY`, journal `OFF` or minimal WAL, allocations from a fixed pool. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:304-311 | NEW: X2-01 Step-2 router tiers | C-02, F-04 | `mbed_alloc` is not a known SQLite interface (fixed pool would be SQLite's own heap configuration); SQLite's default cache is about 2 MB, not 8 MB. Schema stays that of F-04. |
| REQ-PP-0058 | Alternative Tier 2 store: a pre-allocated binary ring buffer `/tmp/nmb/ring.bin` in tmpfs (1,048,576 bytes, 10,000 fixed-size 48-byte sample records), O(1) wrap-around, no allocation, never grows, cannot cause OOM. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:312-318 | NEW: X2-01 Step-2 router tiers | F-04 | 10,000 x 48 bytes = 480,000 bytes, not 1,048,576: figures disagree; the format is fixed in F-04/F-07. |
| REQ-PP-0059 | `nmbd-tiny` saturates 100 Mbps Fast Ethernet with a non-blocking socket loop and a 64 KB static receive buffer at under 14 % CPU on the 580 MHz MIPS. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:323-324 | NEW: X2-01 Step-2 router tiers | | Design target. |
| REQ-PP-0060 | Tier 3 `nmbd-micro` budget: stripped binary 380 KB, 178 KB with UPX LZMA; active RSS 1.4 MB with a strict 1.5 MB limit; RAM tmpfs `/tmp` at most 4 MB. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:346-359 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0061 | Tier 3 uses a custom minimal OpenWrt ImageBuilder image for the TL-WR740N v4 profile without non-essential packages (LuCI, PPP/PPPoE, USB, `opkg`, full `wpad`; `firewall4`/`nftables` replaced by iptables; `wpad-basic-wolfssl`, `libubox`, `libuci`, `uclient-fetch` kept), so kernel plus root drop to about 3.1 MB with about 750 KB writable. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:334; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:364-376 | NEW: X2-01 Step-2 router tiers | I-10 | Without PPP the box has no PPPoE (router mode) on Tier 3; the source says "if connected behind an upstream router". Q-35 answered by 6B.8 point 6: listed in the target's parity table and shown in the UI as not available, never faked; the recipe is checked again when the owner schedules this target (*Order of Work* step 3). |
| REQ-PP-0062 | The Tier 3 firewall backend uses iptables (no `nftables`/`firewall4` in the image). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:368-370 | NEW: X2-01 Step-2 router tiers | N-02 | N-02's interface must allow an iptables backend. |
| REQ-PP-0063 | `nmbd-micro` is compiled with the OpenWrt MIPS musl toolchain for size: `-Os`, `-flto`, `-mno-branch-likely`, `-mips32r2 -msoft-float`, no unwind tables, `-ffunction-sections -fdata-sections -Wl,--gc-sections`, stripped; then `upx --best --lzma` (380 KB stripped, 178 KB compressed). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:377-397 | NEW: X2-01 Step-2 router tiers | F-00 | |
| REQ-PP-0064 | `nmbd-micro` never writes logs or database records to the SPI NOR flash (limited erase cycles). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:398-399 | NEW: X2-01 Step-2 router tiers | C-03, C-02 | |
| REQ-PP-0065 | On Tier 3 all runtime metrics are kept in the circular ring buffer `/tmp/nmb/ring.bin` (tmpfs); a power cut loses the history but never wears the flash. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:400-401 | NEW: X2-01 Step-2 router tiers | F-04 | |
| REQ-PP-0066 | For long-term history `nmbd-micro` sends daily/hourly rollups to an external NMB server, a syslog server or a cloud endpoint in a single compressed HTTP POST. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:402 | NEW: X2-01 Step-2 router tiers | S-12 | Must follow the collector consents (REQ-PP-0190); a "cloud endpoint" is configuration, never built in (Quality Bar 8). |
| REQ-PP-0067 | Tier 1: Tailscale is embedded with `tsnet` in Go as an in-process userspace node inside `nmbd`, listening on the tailnet address and giving HTTPS access to the dashboard without the external `tailscaled` or a TUN interface. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:410-413 | NEW: X2-01 Step-2 router tiers | N-15 | Stated for the Tier 1 router; the N-15 brief decides `tsnet` or `tailscaled` for the PC and Pi 4 (binary size, REQ-RM-0003). Example address is from the reserved shared range. Owner moved from N-15 to X2-01 in round 1 (router form, step 2: docs/NATIVE_WORK_PLAN.md@ad3758c:21-22); N-15 is Also in for its step-1 choice. |
| REQ-PP-0068 | Tiers 2 and 3 do not run `tailscaled` (about 18 MB RSS); they use the kernel WireGuard (`kmod-wireguard`) configured over netlink (under 400 KB RAM, no CPU when idle). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:414-417 | NEW: X2-01 Step-2 router tiers | N-13, N-18 | |
| REQ-PP-0069 | C Telegram bot: polls `getUpdates` from an asynchronous thread or an `epoll` timer; TLS by mbedTLS or WolfSSL; message strings embedded read-only in the binary; inline keyboards written by a zero-allocation streaming JSON formatter; under 150 KB RAM. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:420-435 | NEW: X2-01 Step-2 router tiers | S-11, C-06, C-04 | The example hard-codes a version string (must come from C-05) and 5 languages (C-06 has 11); the token sits in the request path and must be masked in logs (C-04); the formatter must escape strings. |
| REQ-PP-0070 | The SLA dossier's HTML/CSS template is pre-minified and stored as a constant in the binary. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:438-439 | R-03 | | |
| REQ-PP-0071 | A zero-dependency SVG chart generator draws the charts (`viewBox 0 0 800 360`) by linear projection of time and value onto the plot area and writes `path` and touch `circle` elements straight to the output; no external chart libraries. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:440-445 | R-03 | U-01 | The formula divides by zero for a single point or a flat series (t_max = t_min or v_max = v_min); R-03 must define that case (no NaN in the SVG); candidate for F-03. |
| REQ-PP-0072 | Chart generation time target 0.8 ms (Python: 220 ms), measured. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:446-447 | R-03 | P-13 | Design target. |
| REQ-PP-0073 | Wi-Fi telemetry opens a generic netlink socket, resolves `nl80211` through `GENL_ID_CTRL`, sends `NL80211_CMD_GET_STATION` and `NL80211_CMD_GET_INTERFACE`, and reads signal (dBm), TX bitrate, MCS and channel width from the attributes; no subprocess (Python spawns `iw ... station dump`). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:450-455 | M-14 | N-07 | |
| REQ-PP-0074 | The task manager walks `/proc` and reads `/proc/[pid]/stat`, `status` and `cmdline` into a pre-allocated array sorted in memory; under 0.1 % CPU at a 1 s refresh (Python spawns `ps` every 2 s). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:458-462 | C-09 | | |
| REQ-PP-0075 | STUN client: a UDP socket sends a 20-byte RFC 5389 Binding Request to a STUN server and parses `XOR-MAPPED-ADDRESS` to learn the public address and NAT mapping. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:464-466 | M-12 | C-10 | Source names one public STUN server; the server list is configuration (Quality Bar 9, location-aware), not built in. |
| REQ-PP-0076 | NAT verdict from comparing the mapped address with the local WAN interface address (`eth0`/`pppoe-wan`): WAN address in `100.64.0.0/10` means ISP CGNAT; WAN address in RFC 1918 space means double NAT; WAN address equal to the STUN address means a direct public address. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:467-471 | M-12 | C-11, R-01 | Incomplete as written: RFC 1918 also includes 172.16.0.0/12; the cases "public WAN address differing from the mapped one", IPv6 and "STUN failed" are not given; M-12 takes the full rule from the Python reference and roadmap 4.21. |
| REQ-PP-0077 | Tier 1 ad-blocker: full radix trie of 300,000+ domains in RAM (about 14 MB), sub-microsecond exact and wildcard matching, integration with `dnsmasq` ipset / nftables. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:474-477 | NEW: X2-01 Step-2 router tiers | N-12 | The N-12 brief decides the structure for the PC/Pi class. Owner moved from N-12 to X2-01 in round 1 (router form, step 2: docs/NATIVE_WORK_PLAN.md@ad3758c:21-22); N-12 is Also in for its step-1 choice. |
| REQ-PP-0078 | Tiers 2/3 ad-blocker: two-stage filter, a counting Bloom filter (512 KB, MurmurHash3/FNV-1a, 0.1 % false-positive rate, 50,000 top threat domains) plus a small static hash table for exact matches of the local allow list and custom LAN domains (`custom.list`); under 640 KB RAM; a Bloom filter hit is confirmed by an exact check (a compact exact list, on flash where RAM is short) before a domain is blocked. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:478-483 | NEW: X2-01 Step-2 router tiers | N-12 | A "counting" filter is not a bit vector; a false positive would block a domain not on any list. Q-31 settled (developer's choice: a wrong block is a wrong result, Quality Bar 6). Restated in round 3. |
| REQ-PP-0079 | Blocked domains are answered with `0.0.0.0` or `::`, in under 0.05 ms per query. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:484-485 | N-12 | NEW: X2-01 Step-2 router tiers | The answer form must equal the Python blocker's (F-05). |
| REQ-PP-0080 | *Withdrawn.* Ookla protocol emulation: connect to the nearest Speedtest.net server over TCP port 8080 or port 5060 and stream parallel HTTP GET/POST transfers (`/download`, `/upload`), replacing the closed Ookla binary. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:488-489; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:494 | M-10 | M-05 | withdrawn: superseded by REQ-WP-0179, -0180 (owner decision 3, docs/NATIVE_WORK_PLAN.md@ad3758c:313-316): Ookla's own program in the sandbox first, after the user accepts Ookla's terms; a native engine only if it gives the same results and Ookla's terms allow it. M-10 has no test for this row. Port 5060 is not HTTPS. Kept for traceability (round 1). |
| REQ-PP-0081 | Cloudflare engine: fetches 25, 50 and 100 MB chunks from `speed.cloudflare.com/__down?bytes=...` over pre-warmed keep-alive HTTP/1.1 sockets. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:490 | M-06 | M-05 | Chunk sizes count against data caps (M-05). |
| REQ-PP-0082 | RTR-NetTest / Open-RMBT engine: the official JSON token handshake and multi-thread TCP socket sink. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:491 | M-08 | | |
| REQ-PP-0083 | Fast.com engine uses the same method as the Python reference (`nmb/speed/fastcom.py`): Fast.com's API for up to five Open Connect servers, `range/0-0` latency, 16 streams over the servers for 15 s each way over plain HTTP, stable speed, loaded latency and per-server detail; plain HTTP because TLS decryption on an 880 MHz MIPS or on a Cortex-A72 without AES instructions (the Raspberry Pi 4's CPU) caps the result below a gigabit line. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:492 | M-07 |  | The reason holds for the step-1 Pi 4 target too (round 1). |
| REQ-PP-0084 | Zero-copy receive sinks: data is received with `splice()` or `recv()` into a static 64 KB buffer without allocation, and Mbps is computed in real time. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:493 | M-04 | M-06, M-07, M-08, M-09, M-10, NEW: X2-01 Step-2 router tiers | |
| REQ-PP-0085 | The self-healing watchdog inspects link carrier state, DNS resolution latency, default route persistence and CPU softirq peaks. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:497-498 | P-01 | N-01, C-09 | |
| REQ-PP-0086 | Multi-tier recovery sequence: 1) soft restart of network daemons; 2) interface bounce with hardware MAC re-registration and DHCP release/renew; 3) driver offload toggle (disable EEE, enable RPS steering); 4) safe fallback to the baseline network configuration with the rollback log kept in flash/RAM. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:499-504 | P-01 | C-07, N-03, N-05 | The example `ubus call service restart` is OpenWrt's; service control is a platform interface (F-04). Recovery steps are the box's own action (C-07). |
| REQ-PP-0087 | Dynamic DNS client for DuckDNS and dynv6: an embedded HTTP client sends TLS 1.3 update GET requests to their APIs. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:506-507 | N-16 | C-04 | dynv6 is new against the Python reference (`netd/duckdns.py`); the token is in the URL and must be masked in logs. |
| REQ-PP-0088 | The DynDNS client reads the local WAN IPv4 and IPv6 addresses through netlink and uses them only when they are public; behind NAT or CGNAT (private or shared address) it uses the box's public-address check (M-12). It detects a change at once and submits the update in under 80 ms; under 60 KB RAM. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:508-510 | N-16 | N-01 | Behind NAT or CGNAT the local WAN address is not the public one; "signed updates" is not how DuckDNS works (token). Q-26 settled (developer's choice: fault register, "the public IP missing or replaced by the LAN address"). Restated in round 3. |
| REQ-PP-0089 | The Python LAN topology reference (`nmb/netd/lan_topology.py`) defines the behaviour every port reproduces; it applies when the probe is connected behind a router, with or without aggregation. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:515-520 | N-06 | NEW: X2-01 Step-2 router tiers, F-05 | Python faults are not reproduced (F-03). |
| REQ-PP-0090 | Passthrough: a layer-2 bridge between the uplink and the selected ports; devices stay on the router's subnet and the probe's address moves to the bridge. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:521 | N-06 | C-11 | |
| REQ-PP-0091 | Link aggregation: user-selected ports form one bond, modes `802.3ad`, `balance-rr`, `balance-xor`, `balance-tlb`, `balance-alb`, `active-backup`. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:522 | N-06 | | |
| REQ-PP-0092 | Cross-traffic counters come from the uplink device, never the bridge (a Linux bridge does not count forwarded frames, so speed tests would start during client traffic). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:523-524 | M-03 | N-06 | |
| REQ-PP-0093 | Passthrough and link aggregation are verified per router target on real hardware: RT-AXE7800 (LAG across 2.5GbE and 1GbE, Flow Cache), RT-N12+ (switch trunking, CPU link cap), TL-WR740N (whether the 4 MiB image can hold the bonding module), TL-WR1043ND v1 and Archer C6 v2 (trunking in the switch driver, CPU-to-switch link as ceiling). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:525-534 | NEW: X2-01 Step-2 router tiers | N-06 | |
| REQ-PP-0094 | Each target records whether aggregation can raise the probe's own throughput: where all ports sit behind one CPU link software bonding cannot, and may only help traffic forwarded inside the switch. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:535-536 | NEW: X2-01 Step-2 router tiers | N-06 | |
| REQ-PP-0095 | Each router target proves that its counters include hardware-offloaded traffic (Broadcom Runner, MediaTek PPE), or reads the switch or offload counters instead. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:543 | NEW: X2-01 Step-2 router tiers | M-03 | |
| REQ-PP-0096 | The test's own traffic is separated by subtracting it using its sockets or by reading client-side ports; the WAN counter alone is never assumed to be client traffic. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:544 | M-03 | M-04 | |
| REQ-PP-0097 | Wi-Fi radios count as client-side ports, because their traffic may not cross the wired LAN ports. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:545 | M-03 | M-14, N-07 | |
| REQ-PP-0098 | If traffic cannot be read, the network is treated as busy, as the Python version does. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:546 | M-03 | M-04 | |
| REQ-PP-0099 | Stopping a test on client traffic should stay available on targets with a reduced feature set, because it needs only counter reads and a cancel signal. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:547-548 | NEW: X2-01 Step-2 router tiers | M-03, M-04 | Round 2: the source's modality ("should") kept; the X2-01 brief records each target where it is not available, with the reason. |
| REQ-PP-0100 | Each port passes the Phase 0 checks of the roadmap as far as its hardware allows. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:549-550 | NEW: X2-01 Step-2 router tiers | F-05 | |
| REQ-PP-0101 | `nmb-cli` depends on neither Python nor full `curses`; a target that cannot hold the menu keeps the JSON commands. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:553 | S-10 | NEW: X2-01 Step-2 router tiers | |
| REQ-PP-0102 | Ports follow the clean code rules: hardware differences stay behind one interface per feature, in separate files per target, never as `#ifdef` or device checks spread through shared code. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:554 | F-04 | F-00, NEW: X2-01 Step-2 router tiers | F-00's gate can check build tags / device checks outside backend files. |
| REQ-PP-0103 | Each port records its support in `NetworkMeasureBox-DePyD/CHANGELOG.md` under the roadmap's rules, and only after it has run on that hardware. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:555-556 | NEW: X2-01 Step-2 router tiers | F-00 | |
| REQ-PP-0104 | Routers with a USB port get USB tethering (phone, 5G/4G hotspot or modem as measuring link, only internet, or backup WAN) with the Pi 4 rules: never the wired line, never a LAN port, no default route unless nothing else has one, its own data limit, results stored as mobile, the DHCP guard in router mode. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:561 | NEW: X2-01 Step-2 router tiers | N-10, N-05, M-05, C-11 | The Pi 4 rules themselves are owned by N-10 (roadmap 4.27). |
| REQ-PP-0105 | Targets without USB (TL-WR740N, RT-N12+) do not offer USB tethering; it is planned on the RT-AXE7800, RT-AC57U v1, RB5009 and the Filogic boards (kernel packages, flash limits and the MikroTik container limit per roadmap 4.27). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:561-562 | NEW: X2-01 Step-2 router tiers | N-10 | |
| REQ-PP-0106 | When something a target needs does not work under OpenWrt or the other candidate systems (closed Wi-Fi driver, switch/NPU, flow offload, calibration data, LED/button wiring, DSL/modem firmware, NVRAM setting), the maker's stock firmware for that model is searched for the missing part before the feature is dropped. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:565-566 | NEW: X2-01 Step-2 router tiers | I-10 | |
| REQ-PP-0107 | Order of preference for a missing part: the open driver in mainline Linux or OpenWrt; another open system that supports it (for example DD-WRT, Asuswrt-Merlin); the maker's own GPL source release, rebuilt for NMB's kernel; only then a binary part from the stock firmware. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:567 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0108 | Firmware files and data loaded by a driver (Wi-Fi firmware, board and calibration data, NVRAM defaults, device tree values such as GPIO numbers and MAC offsets) may be used with an open driver. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:569 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0109 | Vendor kernel modules are used only when NMB runs on the maker's kernel (an add-on on stock-based firmware); they are never force-loaded into another kernel (`vermagic`). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:570 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0110 | Vendor user-space tools (for example `wl`, `nvram`, `et`, switch tools) run with their own libraries in the strict sandbox of roadmap 6B.5, without network access of their own and without access to NMB's settings and keys. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:571 | NEW: X2-01 Step-2 router tiers | P-04 | |
| REQ-PP-0111 | Closed parts of the maker's firmware are never copied into the NMB code base or images unless their licence allows redistribution; the box takes them on the device from the user's own copy of the stock firmware (uploaded, or downloaded from the maker's official address and verified as in roadmap 6B.3) or from the device's own flash; GPL parts come with their source. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:572 | NEW: X2-01 Step-2 router tiers | P-02, P-11 | |
| REQ-PP-0112 | Every borrowed part is listed on the licence page (roadmap 6B.3A) with its origin, version and checksum. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:572 | NEW: X2-01 Step-2 router tiers | P-08 | |
| REQ-PP-0113 | The device's calibration, MAC address and factory partitions (`art`, `factory`, `nvram`, EEPROM) are read and backed up, never rewritten; the bootloader stays untouched. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:573 | NEW: X2-01 Step-2 router tiers | P-03, P-02 | |
| REQ-PP-0114 | Stock parts: their version is recorded and checked against known vulnerabilities, they never face the WAN, a part with a known exploitable hole is not used, and Quality Bar point 7 applies to them. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:574 | NEW: X2-01 Step-2 router tiers | N-02 | |
| REQ-PP-0115 | Each target has, before use, a table of missing parts: what is missing, where it was found (file path, version, SHA-256, source firmware file), how it is used (open driver plus firmware file, rebuilt from GPL source, vendor kernel only, sandboxed tool, or not usable with the reason), its licence and the real-device test; a part that cannot be used makes the feature unavailable on that target, said plainly in the UI, never faked. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:575 | NEW: X2-01 Step-2 router tiers | I-10, U-01 | |
| REQ-PP-0116 | The stock firmware files downloaded on 2026-11-10 are unpacked and inventoried read-only on the development PC (never flashed for this): kernel version and configuration, modules, `/lib/firmware`, NVRAM defaults, device tree, user-space tools and their libraries; findings go into the per-target table; device tests follow 7D (recovery path first). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:576-577 | I-10 | NEW: X2-01 Step-2 router tiers | Work plan 5: nothing from a maker's closed firmware goes into the repository. |
| REQ-PP-0117 | The native code is first written platform- and architecture-independent and proven on a PC (x86-64) and the Raspberry Pi 4 (ARM64), where development continues. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:580-581 | P-13 | F-00 | |
| REQ-PP-0118 | Build matrix from the start: amd64 and arm64 built and tested; arm, mips and mipsle cross-compiled. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:580 | F-00 | | Work plan 2.4 also runs arm v7 tests under QEMU. |
| REQ-PP-0119 | Router targets are added afterwards as backends of the same code (service control, firewall, Wi-Fi, flash layout), not as separate programs. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:580 | F-04 | NEW: X2-01 Step-2 router tiers | The C tiers are a second code base from the same spec (roadmap 6B.8, F-07); "same code" is read as same design and spec there. |
| REQ-PP-0120 | Every port runs at least as flawlessly as the Python version on the Raspberry Pi 4. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:584-585 | P-13 | NEW: X2-01 Step-2 router tiers | |
| REQ-PP-0121 | Bugs found on the way are fixed, not carried over. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:584 | F-03 | | |
| REQ-PP-0122 | No port shows a behaviour known to be worse in the Python version (list of known Python faults, Quality Bar point 6). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:584 | F-03 | P-13 | Each package proves its faults (work plan 2.4). |
| REQ-PP-0123 | Every change is searched for security holes, with none left open (Quality Bar point 7). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:584 | P-13 | F-00 | Every package gate has the written security check (work plan 2.4). |
| REQ-PP-0124 | Personal data and site-specific values stay out of the code and the image (Quality Bar point 8). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:584 | P-13 | F-00 | F-00's gate runs the scan per package. |
| REQ-PP-0125 | A port shows only what applies in the box's own country, with nothing tied to Hungary built in (Quality Bar point 9). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:584 | C-10 | C-06 | |
| REQ-PP-0126 | Every port keeps all its files inside `NetworkMeasureBox-DePyD/`. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:584 | F-00 | | |
| REQ-PP-0127 | Every port is measurably better where it can be (CPU, memory, start-up, logging, safe updates). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:584 | P-13 | | Evidence rule: roadmap Quality Bar. |
| REQ-PP-0128 | Development access, detailed logging and safe fixes (section 7D) apply to every port and to the native `nmbd` on all devices (routers, Raspberry Pi 4, x86 boxes, NMB's own distribution); a port that does not meet them is not finished. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:588-589 | P-13 | NEW: X2-01 Step-2 router tiers | |
| REQ-PP-0129 | Every device stays open to its developers and to developer LLM agents after release. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:590 | P-05 | | Through the maintenance mode below (owner's consent). |
| REQ-PP-0130 | A fix is possible while the device runs, with as little risk of bricking as the hardware allows. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:590-591 | P-02 | | |
| REQ-PP-0131 | Maintenance mode is off by default; the owner turns it on in the Web UI, in `nmb-cli`, or with a button/serial console where there is no network. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:594 | P-05 | U-02, S-10 | |
| REQ-PP-0132 | Maintenance mode gives a root shell over SSH with keys only (no passwords), on the LAN or through the box's VPN/Tailscale; never open on the WAN. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:594 | P-05 | N-02 | |
| REQ-PP-0133 | Developer LLM agents use the same channel and the same rights as a human developer, with the full logs, the configuration and the database (read-only by default); there is no separate, weaker path that behaves differently. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:595 | P-05 | S-10 | |
| REQ-PP-0134 | `nmb-cli` has JSON output for every Web UI function. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:595 | S-10 | S-02 | |
| REQ-PP-0135 | Every session is recorded: who connected (key name, person or agent), when, and every command and file change, in an audit log the session itself cannot erase: each entry is sent to the log target off the device (REQ-PP-0199) as it is written, and the hash-chained local log stays as a second record. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:596 | P-05 | C-03, S-13 | Q-17 answered by the owner (2026-11-11): B. Round 4 had undone the round-3 restatement. |
| REQ-PP-0199 | Maintenance mode does not start without a log target off the device: the collector, syslog or the hub of the remote paths. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:596 | P-05 | S-13, N-18, C-03 | New: the owner's answer to Q-17 (2026-11-11). |
| REQ-PP-0136 | The owner sees the audit log in the Web UI and can end a session or turn maintenance mode off at any time. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:596 | P-05 | U-02 | |
| REQ-PP-0137 | Maintenance mode turns itself off after a set time (for example 24 h) unless renewed. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:597-598 | P-05 | | |
| REQ-PP-0138 | Remote access without Tailscale on the device: a box is reached, managed and programmed from anywhere through the paths below, in their order of preference. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:601-602 | N-18 | | Round 4: split; that the C builds carry no Tailscale is REQ-PP-0194 (X2-01). Work plan N-18 (`docs/NATIVE_WORK_PLAN.md@ad3758c:170`) builds these paths in step 1. |
| REQ-PP-0194 | The C builds do not carry Tailscale (it exists only in Go and is too large for the small routers' 32 MB RAM); on them the remote paths of 7D are the way in. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:601-602 | NEW: X2-01 Step-2 router tiers | N-18, N-15 | Split from REQ-PP-0138 in round 4. |
| REQ-PP-0139 | Path 1: kernel WireGuard dials out as a client to a hub the owner chooses (the Pi 4 box, a larger router with a WireGuard server, or a rented server), so it works behind CGNAT and firewalls with no open WAN port; SSH (`dropbear`, keys only), the Web UI and `nmb-cli` run through the tunnel. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:603 | N-18 | N-13 | |
| REQ-PP-0140 | `nmbd` sets the hub tunnel up itself (keys made on the device, the hub's public key and address entered once or scanned as a QR code) and keeps it alive. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:603 | N-18 | N-13, U-02 | |
| REQ-PP-0141 | Path 2: a box that runs Tailscale (Pi 4 or a Go-capable router) announces the hub's tunnel network or the local network into the owner's tailnet as a subnet router, so the small routers are reached through it with no Go code on them; on the same LAN no tunnel is needed. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:604 | N-18 | N-15 | |
| REQ-PP-0142 | Path 3: when WireGuard cannot be used (no kernel module, UDP blocked), a reverse SSH tunnel to a jump host the owner names: `dropbear`'s client keeps an outgoing connection with a remote forward, keys only. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:605 | N-18 | | |
| REQ-PP-0143 | Path 4: the collector's fixed commands (speed test, diagnostics, upload) and the Telegram bot work over outgoing connections only, stay available on every target, and never give a shell. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:606 | S-12 | S-11 | |
| REQ-PP-0144 | Each router target has a written recovery path (vendor rescue mode, OpenWrt failsafe, TFTP/U-Boot recovery, or serial console) before its port starts, with the steps tested once on the real device. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:607-608; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:627 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0145 | Every remote path is off until the owner switches it on, uses keys only (never a password), never listens on the WAN, and is under the maintenance mode's time limit and audit log. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:609 | N-18 | P-05, N-02 | |
| REQ-PP-0146 | The hub's address and keys are configuration, never built in (Quality Bar point 8). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:609 | N-18 | C-01 | |
| REQ-PP-0147 | A lost tunnel never stops measuring. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:609 | N-18 | M-01, M-04 | |
| REQ-PP-0148 | Programming over the remote paths is as on the Pi: a cross-built `nmbd` is uploaded through the tunnel or the Web UI, started beside the running one, and rolled back by itself when it does not report healthy. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:609 | P-02 | N-18 | |
| REQ-PP-0149 | On routers logs go to the collector or syslog instead of the flash. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:609 | NEW: X2-01 Step-2 router tiers | C-03, S-12 | Router only (step 2): owner moved from C-03 in round 4 (register README, *Owner and Also in*). |
| REQ-PP-0150 | Each router target's table states which remote paths it has (8 MiB flash: WireGuard and `dropbear` fit; 4 MiB TL-WR740N: maybe only one, path 2 over the LAN or path 3 planned there, confirmed when its image is built), and each path is tested on the real device: connect from outside through CGNAT, upload a build, forced rollback, tunnel loss during a test. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:611-612 | NEW: X2-01 Step-2 router tiers | N-18 | |
| REQ-PP-0151 | Every component writes structured logs: one JSON object per line with time, level, component, event and values. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:615 | C-03 | | |
| REQ-PP-0152 | The debug level is switchable at run time per component, without a restart, and returns to normal by itself after a set time. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:616 | C-03 | S-10, U-02 | |
| REQ-PP-0153 | Logs are flash-safe: the storage guard rules of the Pi version apply to them, so logs never fill the disk. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:617 | C-03 | C-08 | Round 4: split; the RAM ring buffer and the rare flash writes on routers with small NOR/NAND flash are REQ-PP-0195 (X2-01). |
| REQ-PP-0195 | On routers with small NOR/NAND flash the logs live in a RAM ring buffer (as for `nmbd-micro`) and are written to flash only in small, rare blocks, or sent out (syslog, the collector, a developer's machine). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:617 | NEW: X2-01 Step-2 router tiers | C-03, C-08 | Split from REQ-PP-0153 in round 4. |
| REQ-PP-0154 | Crash evidence survives a reboot: the last log lines, the stack/core summary and the reason of the last restart are kept (reserved RAM area or a small dedicated file) and shown after boot. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:618 | C-03 | U-02 | |
| REQ-PP-0155 | One command (`nmb-cli bundle`, or the Web UI) packs logs, configuration with secrets masked, versions, interfaces, routes and recent events for the developer. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:619-620 | R-04 | S-10, C-04, U-02 | |
| REQ-PP-0156 | An update or fix never changes the only working copy: it goes beside the running version, in A/B system slots where flash allows (Pi, x86, own distribution). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:623 | P-02 | P-11 | Round 4: split; routers (A/B slots from 16 MB, otherwise a second copy of the binary and its configuration) are REQ-PP-0196 (X2-01). |
| REQ-PP-0196 | On routers an update or fix also goes beside the running version: A/B system slots on routers with 16 MB or more, otherwise a second copy of the NMB binary and its configuration, with the router firmware itself untouched. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:623 | NEW: X2-01 Step-2 router tiers | P-02, P-11 | Split from REQ-PP-0156 in round 4. |
| REQ-PP-0157 | After an update or hot fix the device must confirm itself healthy (services up, network reachable, Web UI answering) within a set time; otherwise it goes back to the previous version by itself, with no person present. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:624 | P-02 | P-01 | |
| REQ-PP-0158 | Fixes replace only the NMB binary or its configuration where possible, not the kernel or the bootloader; NMB never writes the bootloader. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:625 | P-02 | | |
| REQ-PP-0159 | A configuration is checked before it is applied and the previous configuration is kept, so a wrong setting cannot cut the device off. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:626 | C-01 | P-02 | |
| REQ-PP-0160 | Network changes that could lock out the developer (firewall, routes, addresses) are applied with a timed confirm and undone if not confirmed (commit-confirmed). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:626 | N-03 | N-01, N-02, S-02, U-02 | |
| REQ-PP-0161 | The Raspberry Pi 4, x86 boxes and NMB's own distribution also have a written recovery path, tested once on the real device. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:627 | P-11 | P-02 | Read from "a recovery path for each target" with 7D's scope (line 588). |
| REQ-PP-0162 | Each target table states the brick risk: how updates are applied, how rollback works and what recovery needs (network only, a button, or a serial cable); on 4 MB flash targets (TL-WR740N) only the NMB binary is swapped and the firmware is never rewritten from the field. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:628-629 | NEW: X2-01 Step-2 router tiers | P-02 | |
| REQ-PP-0163 | Before a port counts as done, on the real device: maintenance mode is turned on and off. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:632 | P-13 | P-05, NEW: X2-01 Step-2 router tiers | |
| REQ-PP-0164 | Release test on the real device: connect with a key as a person and as an agent; both appear in the audit log. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:632 | P-13 | P-05 | |
| REQ-PP-0165 | Release test on the real device: switch a component to debug and back. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:632 | P-13 | C-03 | |
| REQ-PP-0166 | Release test on the real device: force a crash and read the evidence after the reboot. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:632 | P-13 | C-03 | |
| REQ-PP-0167 | Release test on the real device: install a deliberately broken update and see the automatic rollback. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:632 | P-13 | P-02 | |
| REQ-PP-0168 | Release test on the real device: apply a network change without confirming it and see it undone. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:632 | P-13 | N-03 | |
| REQ-PP-0169 | Release test on the real device: go through the written recovery path once. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:632 | P-13 | P-11, NEW: X2-01 Step-2 router tiers | |
| REQ-PP-0170 | Each result of these release tests goes into `docs/TESTING.md`. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:632-633 | P-13 | | Native files live in `NetworkMeasureBox-DePyD/` (REQ-PP-0126): read as the native repository's `docs/TESTING.md`. |
| REQ-PP-0171 | Every router target documents its recovery path and its stock-firmware method before the feature of going back to the maker's firmware (roadmap 6B) is offered on it. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:636-637 | NEW: X2-01 Step-2 router tiers | P-02 | The other 6B product decisions named here are registered from the roadmap part. |
| REQ-PP-0172 | Go builds are static: `CGO_ENABLED=0`, `-ldflags="-s -w"`, `GOOS=linux` with the target `GOARCH` (arm64 for Tier 1), from `./cmd/nmbd`, one named output per target. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:643-650 | F-00 | P-11 | `export CC=...` in the example has no effect with cgo off. |
| REQ-PP-0173 | Tier 2 build: OpenWrt `ramips/mt7620` SDK toolchain (`mipsel` musl gcc), `-Os -flto -mips32r2 -msoft-float -s`, linked with `ubox`, `uci`, `mbedtls`, `mbedcrypto`; UPX `--best --lzma` copy. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:651-658 | NEW: X2-01 Step-2 router tiers | I-10 | Toolchain version (gcc 12.3.0) is that of OpenWrt 23.05; it follows the downloaded SDK set (roadmap 6B.4). |
| REQ-PP-0174 | Tier 3 build: OpenWrt `ath79/generic` SDK toolchain (`mips` musl gcc), `-Os -flto -mips32r2 -msoft-float -ffunction-sections -fdata-sections -Wl,--gc-sections -s`, same libraries; UPX `--best --lzma` copy. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:659-668 | NEW: X2-01 Step-2 router tiers | I-10 | |
| REQ-PP-0175 | Design target, to measure: idle RAM (RSS) Tier 1 18 MB, Tier 2 2.8 MB, Tier 3 1.4 MB (Python on the Pi 4: 180-220 MB). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:677 | NEW: X2-01 Step-2 router tiers | P-13 | |
| REQ-PP-0176 | Design target, to measure: RAM during an active speed test Tier 1 24 MB, Tier 2 3.6 MB, Tier 3 1.8 MB (Python: 240 MB). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:678 | NEW: X2-01 Step-2 router tiers | P-13 | |
| REQ-PP-0177 | Design target, to measure: flash space Tier 1 12 MB (NAND), Tier 2 320 KB (UPX), Tier 3 178 KB (UPX) (Python: about 75 MB). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:679 | NEW: X2-01 Step-2 router tiers | P-13 | |
| REQ-PP-0178 | Design target, to measure: idle CPU load Tier 1 under 0.2 %, Tier 2 under 1.0 %, Tier 3 under 2.5 % (Python: 4-12 %). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:680 | NEW: X2-01 Step-2 router tiers | P-13 | |
| REQ-PP-0179 | Design target, to measure: CPU load during a 100 Mbps test Tier 1 under 1.5 %, Tier 2 14 %, Tier 3 22 % (Python: 45 %). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:681 | NEW: X2-01 Step-2 router tiers | P-13 | |
| REQ-PP-0180 | Design target, to measure: CPU load during a 1 Gbps test on Tier 1 4.2 % (Tiers 2/3 have 100 Mbps ports; Python: 100 %, interpreter-lock limited). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:682 | NEW: X2-01 Step-2 router tiers | P-13 | |
| REQ-PP-0181 | Design target, to measure: CPU load during a 2.5 Gbps test on Tier 1 11.5 % (Python cannot reach 2.5 Gbps). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:683 | NEW: X2-01 Step-2 router tiers | P-13 | |
| REQ-PP-0182 | Design target, to measure: `/proc` task scan time Tier 1 0.8 ms, Tier 2 1.4 ms, Tier 3 2.1 ms (Python: 22 ms). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:684 | NEW: X2-01 Step-2 router tiers | C-09, P-13 | |
| REQ-PP-0183 | Design target, to measure: ICMP probe jitter added by the engine Tier 1 under 0.04 ms, Tier 2 under 0.08 ms, Tier 3 under 0.12 ms (Python: 2.5 ms from garbage collection). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:685 | NEW: X2-01 Step-2 router tiers | M-01, P-13 | |
| REQ-PP-0184 | Design target, to measure: zero subprocess forks per minute on Tier 1, Tier 2 and Tier 3 (Python on the Pi 4: 120 per minute). | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:686-687 | NEW: X2-01 Step-2 router tiers | P-13 | Round 2: a router-tier design target like REQ-PP-0175 to -0183 (line 673: the router columns are design targets); owner moved from P-13 to X2-01. Step-1 coverage comes through REQ-PP-0029 and REQ-PP-0185. Round 4: the added rule on counting the sandboxed programs apart is removed (not in lines 686-687). **Conflict** with REQ-RM-1350 and REQ-WP-0179 (separate sandboxed programs): Q-40. |
| REQ-PP-0185 | On the Raspberry Pi 4 the native build is measured for the parameters of the performance table of section 9 (idle and test RAM, flash, idle and test CPU, /proc scan time, probe jitter, forks per minute) and compared with the Python version, and is better where it can be. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:584 | P-13 |  | Derived (round 1) from line 584 (QB 5: at least as flawless as the Python version on the Pi 4, measurably better where it can). The parameter list is that of the table at lines 675-686, whose figures are router design targets only; the Tier 1 column is not a Pi 4 target. Round 2: the source narrowed to line 584, the table rows no longer name this id. |
| REQ-PP-0186 | Goal: the RT-AXE7800 becomes a multi-gigabit probe capable of 2.5 Gbps line-rate tests, CAKE shaping and native Tailscale management, without slowing gaming or Wi-Fi 6E performance. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:693 | NEW: X2-01 Step-2 router tiers | | "Without slowing" must be measured (forwarding throughput and latency with NMB idle and active). CAKE during the tests: Q-25 (REQ-PP-0053); lifting the shaper is asked when X2-01 starts. |
| REQ-PP-0187 | Goal: the RT-N12+ becomes an SLA probe with web reporting and Telegram alerts. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:694 | NEW: X2-01 Step-2 router tiers | | |
| REQ-PP-0188 | Goal: the TL-WR740N v4 runs a reduced diagnostic suite, aiming at 178 KB of flash and 1.4 MB of RAM. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:695 | NEW: X2-01 Step-2 router tiers | | The reduced suite is not listed; it must at least keep stopping tests on traffic (REQ-PP-0099) and the collector subset (REQ-PP-0192). |
| REQ-PP-0189 | Goal: the TL-WR1043ND v1 and the Archer C6 v2 add gigabit-port targets, with throughput limits measured on hardware. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:696-697 | NEW: X2-01 Step-2 router tiers | | The RT-AC57U v1 and the families of 1.3 have no goal line. |
| REQ-PP-0190 | The router build sends to the same `nmb-collector` as the Python box, with the same three consents (`enabled`, `share_diagnostics`, `allow_remote_control`) and the same fixed command list. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:699 | NEW: X2-01 Step-2 router tiers | S-12 | Router build only (step 2): owner moved from S-12 in round 4; the step-1 client keeps the same protocol under roadmap 4.25 (S-12). |
| REQ-PP-0191 | The collector client only polls outward, so it needs no open port on the router's WAN. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:699 | S-12 | N-02 | |
| REQ-PP-0192 | On 4 MB flash targets the diagnostics report is left out of the collector client; only speed tests and ping summaries are sent. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:699 | NEW: X2-01 Step-2 router tiers | S-12 | |
| REQ-PP-0193 | The collector server is not ported to the router builds; it runs on the ISP's or the specialist's server. | docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:699-700 | NEW: X2-01 Step-2 router tiers | S-13 | Q-39 answered: the porting plan concerns the router builds; work plan S-13 ports the collector server natively in step 1 for the PC and server build. A rule of the router builds only: owner moved from S-13 in round 4. |

## Coverage

Each range includes the blank line(s) that end it. Every line 1-700 is in exactly one row.

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1 | heading | context: document title |
| 2-3 | heading | context: subtitle naming the planned targets (their requirements are in sections 1.1-1.3) |
| 4-5 | paragraph | REQ-PP-0001, REQ-PP-0002 |
| 6-7 | separator | context: horizontal rule |
| 8-9 | heading | context: section 1 title |
| 10-11 | paragraph | REQ-PP-0002, REQ-PP-0003, REQ-PP-0004, REQ-PP-0005 |
| 12-13 | paragraph | REQ-PP-0006 |
| 14-15 | paragraph | REQ-PP-0007, REQ-PP-0197 |
| 16-17 | code block | context: diagram opening lines |
| 18-22 | code block | REQ-PP-0008 |
| 23-27 | code block | REQ-PP-0009 |
| 28-32 | code block | REQ-PP-0010 |
| 33-35 | code block | REQ-PP-0008, REQ-PP-0009, REQ-PP-0010 |
| 36-37 | code block | context: closing fence |
| 38-39 | heading | context: 1.1 title |
| 40-41 | table row | context: table header and alignment row |
| 42-43 | table row | context: reference links and FCC ids (hardware identity, checked under REQ-PP-0020) |
| 44-47 | table row | context: SoC, switch chip, CPU architecture, cores and clock (hardware facts, REQ-PP-0020) |
| 48 | table row | REQ-PP-0011 |
| 49-51 | table row | context: RAM, free RAM and flash chips (hardware facts, REQ-PP-0020; inputs to REQ-PP-0012 to -0014) |
| 52 | table row | REQ-PP-0012 |
| 53-55 | table row | context: wired ports, USB, Wi-Fi (hardware facts, REQ-PP-0020; USB presence used by REQ-PP-0105) |
| 56-57 | table row | REQ-PP-0008, REQ-PP-0009, REQ-PP-0010 |
| 58 | table row | REQ-PP-0013 |
| 59 | table row | REQ-PP-0014 |
| 60-61 | table row | REQ-PP-0015 |
| 62-63 | separator | context: horizontal rule |
| 64-65 | heading | context: 1.2 title |
| 66-67 | table row | context: table header and alignment row |
| 68-76 | table row | context: hardware facts of the additional targets (REQ-PP-0020) |
| 77 | table row | REQ-PP-0016 |
| 78-79 | table row | REQ-PP-0017, REQ-PP-0018 |
| 80-82 | paragraph | REQ-PP-0018, REQ-PP-0019 |
| 83-84 | heading | context: 1.3 title |
| 85-86 | paragraph | REQ-PP-0020 |
| 87-88 | paragraph | REQ-PP-0021 |
| 89 | bullet | REQ-PP-0022 |
| 90-91 | bullet | REQ-PP-0023 |
| 92-93 | table row | context: table header and alignment row |
| 94-98 | table row | REQ-PP-0020, REQ-PP-0022 (example MikroTik models, figures to verify) |
| 99-100 | paragraph | REQ-PP-0024, REQ-PP-0025 |
| 101-102 | table row | context: table header and alignment row |
| 103-109 | table row | REQ-PP-0026, REQ-PP-0020 |
| 110 | paragraph | REQ-PP-0027, REQ-PP-0006 |
| 111-112 | separator | context: horizontal rule |
| 113-114 | heading | context: section 2 title |
| 115 | heading | context: 2.1 title |
| 116-119 | bullet | context: rationale (Python does not fit router flash); the rule it supports is REQ-PP-0006 |
| 120-124 | bullet | context: rationale (Python exhausts router RAM, OOM example); supports REQ-PP-0006 |
| 125-126 | bullet | REQ-PP-0028 |
| 127-129 | bullet | REQ-PP-0029 |
| 130-131 | separator | context: horizontal rule |
| 132-133 | heading | context: section 3 title |
| 134-138 | paragraph | REQ-PP-0030, REQ-PP-0198 |
| 139-142 | code block | context: diagram frame and title |
| 143 | code block | REQ-PP-0031 |
| 144 | code block | REQ-PP-0032 |
| 145 | code block | REQ-PP-0033 |
| 146 | code block | REQ-PP-0034 |
| 147 | code block | REQ-PP-0035 |
| 148 | code block | REQ-PP-0036 |
| 149 | code block | REQ-PP-0037 |
| 150 | code block | REQ-PP-0038 |
| 151 | code block | REQ-PP-0039 |
| 152 | code block | REQ-PP-0040 |
| 153 | code block | REQ-PP-0041 |
| 154-156 | code block | context: diagram frame and closing fence |
| 157-158 | separator | context: horizontal rule |
| 159-160 | heading | context: section 4 title |
| 161 | heading | context: 4.1 title |
| 162-176 | bullet | context: Tier 1 hardware facts (REQ-PP-0020) |
| 177-181 | bullet | REQ-PP-0042 |
| 182-191 | code block | REQ-PP-0043 |
| 192-201 | code block | REQ-PP-0033, REQ-PP-0034, REQ-PP-0035, REQ-PP-0039, REQ-PP-0040, REQ-PP-0046, REQ-PP-0053, REQ-PP-0067 (component list of the Tier 1 build) |
| 202-207 | code block | REQ-PP-0043 |
| 208 | heading | context: 4.2 title |
| 209 | paragraph | REQ-PP-0044 |
| 210-211 | bullet | REQ-PP-0045 |
| 212-214 | bullet | REQ-PP-0046 |
| 215-216 | bullet | REQ-PP-0047 |
| 217 | bullet | REQ-PP-0048 |
| 218-226 | bullet | REQ-PP-0049 |
| 227-228 | bullet | REQ-PP-0050 |
| 229 | heading | context: 4.3 title |
| 230 | paragraph | context: states that the Broadcom SoC has hardware acceleration (fact behind REQ-PP-0051) |
| 231-239 | bullet | REQ-PP-0051 |
| 240-241 | bullet | REQ-PP-0052 |
| 242 | heading | context: 4.4 title |
| 243-246 | bullet | REQ-PP-0053 |
| 247-250 | bullet | REQ-PP-0054 |
| 251-252 | separator | context: horizontal rule |
| 253-254 | heading | context: section 5 title |
| 255 | heading | context: 5.1 title |
| 256-265 | bullet | context: Tier 2 hardware facts (REQ-PP-0020) |
| 266-267 | bullet | REQ-PP-0009 |
| 268-286 | code block | REQ-PP-0055 |
| 287 | heading | context: 5.2 title |
| 288 | paragraph | context: introduction to the flash layout |
| 289-295 | code block | context: flash layout figures (hardware facts, REQ-PP-0020) |
| 296-297 | bullet | context: OpenWrt base packages and the overlay space left (input to REQ-PP-0056) |
| 298-302 | bullet | REQ-PP-0056 |
| 303 | heading | context: 5.3 title |
| 304 | paragraph | REQ-PP-0057 |
| 305-311 | bullet | REQ-PP-0057 |
| 312-318 | bullet | REQ-PP-0058 |
| 319 | heading | context: 5.4 title |
| 320-322 | bullet | context: Fast Ethernet line rate and the CPU use of Python/curl/iperf scripts (comparison only) |
| 323-324 | bullet | REQ-PP-0059 |
| 325-326 | separator | context: horizontal rule |
| 327-328 | heading | context: section 6 title |
| 329 | heading | context: 6.1 title |
| 330-333 | bullet | context: Tier 3 hardware facts (REQ-PP-0020) |
| 334-335 | bullet | REQ-PP-0010, REQ-PP-0061 |
| 336-345 | code block | context: 4 MB flash partition figures (hardware facts, REQ-PP-0020; overlay figure noted at REQ-PP-0012) |
| 346-351 | code block | REQ-PP-0060 |
| 352-359 | code block | REQ-PP-0060 |
| 360 | heading | context: 6.2 title |
| 361-363 | paragraph | context: why standard OpenWrt images do not fit 4 MB flash |
| 364-376 | bullet | REQ-PP-0061, REQ-PP-0062 |
| 377-397 | bullet | REQ-PP-0063 |
| 398-399 | bullet | REQ-PP-0064 |
| 400-401 | bullet | REQ-PP-0065 |
| 402-403 | bullet | REQ-PP-0066 |
| 404-405 | separator | context: horizontal rule |
| 406-407 | heading | context: section 7 title |
| 408 | heading | context: 7.1 title |
| 409 | bullet | context: current Python architecture (Tailscale CLI wrapper) |
| 410-413 | bullet | REQ-PP-0067 |
| 414-417 | bullet | REQ-PP-0068 |
| 418 | heading | context: 7.2 title |
| 419 | bullet | context: current Python architecture (urllib long polling) |
| 420-435 | bullet | REQ-PP-0069 |
| 436 | heading | context: 7.3 title |
| 437 | bullet | context: current Python architecture (HTML dossier with SVG) |
| 438-439 | bullet | REQ-PP-0070 |
| 440-445 | bullet | REQ-PP-0071 |
| 446-447 | bullet | REQ-PP-0072 |
| 448 | heading | context: 7.4 title |
| 449 | bullet | context: current Python architecture (`iw` subprocess and regex) |
| 450-455 | bullet | REQ-PP-0073 |
| 456 | heading | context: 7.5 title |
| 457 | bullet | context: current Python architecture (`ps` every 2 s) |
| 458-462 | bullet | REQ-PP-0074 |
| 463 | heading | context: 7.6 title |
| 464-466 | bullet | REQ-PP-0075 |
| 467-471 | bullet | REQ-PP-0076 |
| 472 | heading | context: 7.7 title |
| 473 | bullet | context: current Python architecture (DNS sniffer, SQLite query log, hosts lists) |
| 474-477 | bullet | REQ-PP-0077 |
| 478-483 | bullet | REQ-PP-0078 |
| 484-485 | bullet | REQ-PP-0079 |
| 486 | heading | context: 7.8 title |
| 487 | bullet | context: current Python architecture (engine subprocesses) |
| 488-489 | bullet | REQ-PP-0080 |
| 490 | bullet | REQ-PP-0081 |
| 491 | bullet | REQ-PP-0082 |
| 492 | bullet | REQ-PP-0083 |
| 493 | bullet | REQ-PP-0084 |
| 494-495 | bullet | REQ-PP-0080, REQ-PP-0006 |
| 496 | heading | context: 7.9 title |
| 497-498 | bullet | REQ-PP-0085 |
| 499-504 | bullet | REQ-PP-0086 |
| 505 | heading | context: 7.10 title (DynDNS; a second 7.10 follows at line 698) |
| 506-507 | bullet | REQ-PP-0087 |
| 508-510 | bullet | REQ-PP-0088 |
| 511-512 | separator | context: horizontal rule |
| 513-514 | heading | context: 7A title |
| 515-516 | paragraph | REQ-PP-0089 |
| 517-520 | code block | REQ-PP-0089 (topology diagram: router, uplink or bond, NMB, passthrough ports) |
| 521 | bullet | REQ-PP-0090 |
| 522 | bullet | REQ-PP-0091 |
| 523-524 | bullet | REQ-PP-0092 |
| 525-534 | table row | REQ-PP-0093 (line 525 introduction, 527-528 header and alignment row, 529-533 one row per target) |
| 535-536 | paragraph | REQ-PP-0094 |
| 537-538 | separator | context: horizontal rule |
| 539-540 | heading | context: 7B title |
| 541-542 | paragraph | context: introduction (router ports would see client traffic; not yet implemented); its rules follow in 543-547 |
| 543 | bullet | REQ-PP-0095 |
| 544 | bullet | REQ-PP-0096 |
| 545 | bullet | REQ-PP-0097 |
| 546 | bullet | REQ-PP-0098 |
| 547-548 | bullet | REQ-PP-0099 |
| 549-550 | paragraph | REQ-PP-0100 |
| 551-552 | paragraph | context: introduction to the list that follows |
| 553 | bullet | REQ-PP-0101 |
| 554 | bullet | REQ-PP-0102 |
| 555-556 | bullet | REQ-PP-0103 |
| 557-558 | separator | context: horizontal rule |
| 559-560 | heading | context: 7G title |
| 561-562 | paragraph | REQ-PP-0104, REQ-PP-0105 |
| 563-564 | heading | context: 7H title |
| 565-566 | paragraph | REQ-PP-0106 |
| 567 | bullet | REQ-PP-0107 |
| 568 | bullet | context: lead-in of point 2; its rules are in lines 569-571 |
| 569 | bullet | REQ-PP-0108 |
| 570 | bullet | REQ-PP-0109 |
| 571 | bullet | REQ-PP-0110 |
| 572 | bullet | REQ-PP-0111, REQ-PP-0112 |
| 573 | bullet | REQ-PP-0113 |
| 574 | bullet | REQ-PP-0114 |
| 575 | bullet | REQ-PP-0115 |
| 576-577 | bullet | REQ-PP-0116 |
| 578-579 | heading | context: 7F title |
| 580-581 | paragraph | REQ-PP-0117, REQ-PP-0118, REQ-PP-0119 |
| 582-583 | heading | context: 7E title |
| 584-585 | paragraph | REQ-PP-0120, REQ-PP-0121, REQ-PP-0122, REQ-PP-0123, REQ-PP-0124, REQ-PP-0125, REQ-PP-0126, REQ-PP-0127, REQ-PP-0185 |
| 586-587 | heading | context: 7D title |
| 588-589 | paragraph | REQ-PP-0128 |
| 590-591 | paragraph | REQ-PP-0129, REQ-PP-0130 |
| 592-593 | heading | context: "Access" title |
| 594 | bullet | REQ-PP-0131, REQ-PP-0132 |
| 595 | bullet | REQ-PP-0133, REQ-PP-0134 |
| 596 | bullet | REQ-PP-0135, REQ-PP-0136, REQ-PP-0199 |
| 597-598 | bullet | REQ-PP-0137 |
| 599-600 | heading | context: "Remote access without Tailscale" title |
| 601-602 | paragraph | REQ-PP-0138, REQ-PP-0194 |
| 603 | bullet | REQ-PP-0139, REQ-PP-0140 |
| 604 | bullet | REQ-PP-0141 |
| 605 | bullet | REQ-PP-0142 |
| 606 | bullet | REQ-PP-0143 |
| 607-608 | bullet | REQ-PP-0144 |
| 609-610 | paragraph | REQ-PP-0145, REQ-PP-0146, REQ-PP-0147, REQ-PP-0148, REQ-PP-0149 |
| 611-612 | paragraph | REQ-PP-0150 |
| 613-614 | heading | context: "Detailed logging" title |
| 615 | bullet | REQ-PP-0151 |
| 616 | bullet | REQ-PP-0152 |
| 617 | bullet | REQ-PP-0153, REQ-PP-0195 |
| 618 | bullet | REQ-PP-0154 |
| 619-620 | bullet | REQ-PP-0155 |
| 621-622 | heading | context: "Fixing without bricking" title |
| 623 | bullet | REQ-PP-0156, REQ-PP-0196 |
| 624 | bullet | REQ-PP-0157 |
| 625 | bullet | REQ-PP-0158 |
| 626 | bullet | REQ-PP-0159, REQ-PP-0160 |
| 627 | bullet | REQ-PP-0144, REQ-PP-0161 |
| 628-629 | bullet | REQ-PP-0162 |
| 630-631 | heading | context: "Testing these requirements" title |
| 632-633 | paragraph | REQ-PP-0163, REQ-PP-0164, REQ-PP-0165, REQ-PP-0166, REQ-PP-0167, REQ-PP-0168, REQ-PP-0169, REQ-PP-0170 |
| 634-635 | heading | context: 7C title |
| 636-637 | paragraph | REQ-PP-0171; context for the rest: pointer to the product decisions of roadmap 6B and the firmware list of 6B.4 (registered from the roadmap part) |
| 638 | paragraph | REQ-PP-0006; context for the rest: pointer to the moved distribution text in roadmap 6A |
| 639-640 | heading | context: section 8 title |
| 641-642 | heading | context: 8.1 title |
| 643-650 | code block | REQ-PP-0172 |
| 651-658 | code block | REQ-PP-0173 |
| 659-668 | code block | REQ-PP-0174 |
| 669-670 | separator | context: horizontal rule |
| 671-672 | heading | context: section 9 title |
| 673-674 | paragraph | REQ-PP-0001 |
| 675-676 | table row | context: table header (Python column and the three router tiers) and alignment row; REQ-PP-0185 takes its parameter list from this table |
| 677 | table row | REQ-PP-0175 |
| 678 | table row | REQ-PP-0176 |
| 679 | table row | REQ-PP-0177 |
| 680 | table row | REQ-PP-0178 |
| 681 | table row | REQ-PP-0179 |
| 682 | table row | REQ-PP-0180 |
| 683 | table row | REQ-PP-0181 |
| 684 | table row | REQ-PP-0182 |
| 685 | table row | REQ-PP-0183 |
| 686-687 | table row | REQ-PP-0184 |
| 688-689 | separator | context: horizontal rule |
| 690-691 | heading | context: section 10 title |
| 692 | paragraph | context: lead-in ("if implemented and the targets are met") |
| 693 | bullet | REQ-PP-0186 |
| 694 | bullet | REQ-PP-0187 |
| 695 | bullet | REQ-PP-0188 |
| 696-697 | bullet | REQ-PP-0189 |
| 698 | heading | context: second section numbered 7.10 (collector client), placed after section 10 |
| 699-700 | paragraph | REQ-PP-0190, REQ-PP-0191, REQ-PP-0192, REQ-PP-0193 |

## Proposed package changes

1. **NEW: X2-01 Step-2 router tiers** (same id and name as proposed in `roadmap-1-507.md`). Work plan section 1 says the C11 tier and everything router-specific start after step 1, and section 3 has no package for them; most of this plan is router work, so without X2-01 the register's check fails. Owner of 93 rows here (after rounds 1 and 2): REQ-PP-0001, -0003, -0004, -0008 to -0026, -0031, -0041 to -0069, -0077, -0078, -0093 to -0095, -0099, -0100, -0103 to -0115, -0144, -0150, -0162, -0171, -0173 to -0184, -0186 to -0189, -0192. Proposal: keep X2-01 as one holding package in step 1 (its rows go into `spec/` through F-07 so step 2 starts from them) and split it when step 2 is planned:
   - one package per router target, starting with the TL-WR1043ND v1 (REQ-PP-0003; this is the `NEW: T-01` of `roadmap-508-1034.md`, which the register owner should make a child of X2-01 rather than a second holding package), then T-02... in the owner's order (REQ-PP-0004);
   - X2-01a Tier 1 Asuswrt/Merlin integration (4.2-4.4: REQ-PP-0042 to -0054);
   - X2-01b C tier storage and budgets (5.3, 6.2, section 9: REQ-PP-0055 to -0066, -0175 to -0184);
   - X2-01c Missing parts from the maker's firmware (7H: REQ-PP-0106 to -0115), next to I-10 (REQ-PP-0116);
   - X2-01d Remote access, recovery and brick risk per target (7D: REQ-PP-0144, -0150, -0162, -0171);
   - X2-01e MikroTik and MediaTek families (1.3: REQ-PP-0021 to -0026).
2. **Duplicate section number.** The plan has two sections "7.10" (DynDNS at line 505, collector client at line 698). The work plan's N-16 row cites "porting plan 7.10" for DynDNS; the register should cite lines, not the number, and S-12 also owns the second 7.10 (REQ-PP-0191; the router-only REQ-PP-0190, -0192 and -0193 are X2-01's, round 4).
3. **Porting plan sections to their main package** (for the register's "every section in exactly one package" check): 1, 1.1, 1.2, 1.3, 2 (rationale), 4, 5, 6, 7G, 7H, 8 (C parts), 9, 10 → X2-01; 3 → P-10; 7.1 → X2-01 (N-15 under Also in; both 7.1 rows, REQ-PP-0067 and -0068, are owned by X2-01 since round 1); 7.2 → S-11 (C form X2-01); 7.3 → R-03; 7.4 → M-14; 7.5 → C-09; 7.6 → M-12; 7.7 → X2-01 (its main rows REQ-PP-0077 and -0078 are owned by X2-01 since round 1; N-12 owns REQ-PP-0079); 7.8 → M-04 (engines M-06 to M-10); 7.9 → P-01; 7.10 (line 505) → N-16; 7.10 (line 698) → S-12; 7A → N-06; 7B → M-03; 7F → F-04; 7E → P-13; 7D → P-05 (with N-18, C-03, P-02, R-04); 7C → P-11; 8 (Go part) → F-00.
4. **Round 1 moves.** REQ-PP-0041 (RAM ring buffer), -0067 (Tier 1 `tsnet`) and -0077 (Tier 1 radix trie) are owned by
   X2-01, with C-02, N-15 and N-12 under *Also in* for their step-1 choice; X2-01 owned 92 rows here after round 1, 93 after round 2 (REQ-PP-0184, a router-tier design target), 101 after round 4. REQ-PP-0080 (Ookla
   protocol emulation) is withdrawn, superseded by REQ-WP-0179 and -0180.
   **Round 4:** the router-only rows REQ-PP-0149, -0190 and -0193 moved to X2-01 (step-1 owner under *Also in*); REQ-PP-0007,
   -0030, -0138, -0153 and -0156 were split into a step-1 row and a router or C row owned by X2-01 (REQ-PP-0197, -0198, -0194,
   -0195, -0196).
5. No other split or merge proposed from this source. P-05 owns 8 rows from 7D (REQ-PP-0199 added on 2026-11-11) and N-18 owns 8; both stay within one sitting by the size rule as far as this source shows.

## Questions for the owner

*Triaged in round 3 (2026-11-11), answers and choices in `QUESTIONS.md`:* 1 → Q-24, 2 → Q-25, 4 → Q-31, 5 → Q-26, 6 → Q-17 (developer's choices); 3 → Q-35, 7 → Q-39 (answered). Round 4: Q-17 moved to the owner, Q-25 keeps only the stored shaper (lifting it is asked when X2-01 starts), and the fork-target conflict of REQ-PP-0029 and -0184 is Q-40 (answered). Answered by the owner on 2026-11-11: Q-17 (REQ-PP-0135 restated, REQ-PP-0199 added). The rows were updated where an answer or a choice settled them; the text below is the question as the part asked it.

1. **Flow Cache bypass (REQ-PP-0051), fault or intended?** The plan's example turns off Broadcom acceleration with `system("fc config --flush")`, a shell call, which REQ-PP-0029 and Quality Bar 7 rule out, and it does not say that acceleration is turned back on after the test. Proposed: call the interface directly (ioctl or `/proc`), never through a shell, and always restore the previous state after the test, also after a crash.
2. **CAKE shaping on the measured line (REQ-PP-0053, -0186), fault or intended?** If `nmbd` shapes the WAN with CAKE, the speed tests measure the shaper, not the ISP's line. Is shaping meant to stay off (or be lifted) during measurements, or to stay on with the shaper recorded in the measurement setup (roadmap 4.29) of every result?
3. **No PPPoE on the TL-WR740N (REQ-PP-0061), intended?** The 4 MiB image recipe removes PPP, so that target cannot dial PPPoE and has no router mode; it would only work behind another router. Is that accepted for Tier 3 (then the UI says so plainly, as 7H point 6 requires)?
4. **Bloom filter false positives (REQ-PP-0078), fault or intended?** A 0.1 % false-positive rate blocks some domains that are on no list. Is that accepted on the small routers, or must a hit be confirmed by an exact check before blocking (more memory)?
5. **DynDNS from the local WAN address (REQ-PP-0088), fault or intended?** Behind NAT or CGNAT the router's WAN address is not the public one, so "no polling of outside servers" would publish a wrong address. Proposed: use the netlink address only when it is public, and otherwise the box's public-address check (M-12).
6. **Audit log "the session cannot erase" (REQ-PP-0135).** A root shell can erase any local file. Should the log be sent off the device as it is written (collector or syslog), should the maintenance shell be a restricted user with logged `sudo`, or is a local append-only log accepted with that limit stated?
7. **Collector server (REQ-PP-0193).** The plan says the collector server "is not ported"; work plan S-13 ports it natively in step 1. Proposed reading: it is not part of the router builds; S-13 stays in step 1 for the PC/server build. Correct?
