# Requirements register, part: CHANGELOG.md

Source: `CHANGELOG.md` of the Python repository at commit `ad3758c` (423 lines; this part covers lines 1-423).
Read in pages: 1-75, 76-150, 151-220, 221-290, 291-360, 361-423 (all lines).
Package ids from `docs/NATIVE_WORK_PLAN.md@ad3758c:98-226` (section 3 and section 4).
Id range of this part: REQ-CL-0001 to REQ-CL-0452 (0452 added in round 1). Every Source cell is written out in full as `CHANGELOG.md@ad3758c:<line>`.
API and page packages are written as `S-02…S-0n (API: <area>)` and `U-02…U-0n (page: <page>)`, using the areas of work plan line 189, until F-02 fixes their numbers (same convention as the TESTING.md part).

Rules used: an Added or Changed entry that describes behaviour of the box is a requirement; a Fixed or Security entry is a requirement stating the correct behaviour, marked `fault` in Notes so F-03 can cross-reference it (the native build must not repeat the fault). Entries that only add or change a plan, a development note or a document with no box behaviour are listed as context in the Coverage table, with the part of the register that holds their requirements. One changelog bullet that states several checkable things is split into several rows with the same source line. "Python proof" in Notes records how the Python version proved it, as a hint for the native test.

## Requirements

| Id | Requirement | Source | Package | Also in | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| REQ-CL-0001 | Every measurement (speed test, outage, ping minute) carries the box's mode (standalone, router, bridge), the link the measurement used (Ethernet, PPPoE, Wi-Fi client, USB tethering, inline bridge, LAN passthrough, link aggregation) and the path to the ISP: directly on it (own PPPoE session, or a public or carrier-grade NAT address on that link) or behind another router (a private address). | CHANGELOG.md@ad3758c:14 | C-11 | M-01, M-02, M-04, N-01, N-03, R-01, R-02, R-03, R-04, R-05, S-02…S-0n (API: status), U-02…U-0n (page: home), S-10, S-11, S-12, S-13 |  |
| REQ-CL-0002 | The position on the path to the ISP is read from the link's IPv4 address, never assumed from the mode; what cannot be told for sure is stored and shown as "unknown". | CHANGELOG.md@ad3758c:14 | C-11 | N-01 |  |
| REQ-CL-0003 | A speed test during which the box switched mode is marked as changed. | CHANGELOG.md@ad3758c:14 | C-11 | M-04 |  |
| REQ-CL-0004 | The network daemon keeps the history of the setup (Python table `network_state`). | CHANGELOG.md@ad3758c:14 | C-11 | N-03, C-02 |  |
| REQ-CL-0005 | The setup words are the same in the dossier, the Web UI and the bot, in all 11 languages, from one text source. | CHANGELOG.md@ad3758c:14 | C-11 | C-06, R-03, U-01, S-11 |  |
| REQ-CL-0006 | Speed test rows from before the setup was recorded are reconstructed once from the event log and the stored interface and address, marked `derived`; nothing measured is replaced; every value written is kept in `speedtest_correction`, so the collector receives it too. A migrated Python database keeps these marks. | CHANGELOG.md@ad3758c:14 | C-02 | C-11, S-12 | Python data migration |
| REQ-CL-0007 | Every speed test row records the address of the server it actually talked to (`server_ip`): Ookla's server address from its result, the Cloudflare edge address the test connected to, the iPerf3 server's, the Fast.com server's, and the RTR server's when RTR gives one. | CHANGELOG.md@ad3758c:19 | M-04 | M-06, M-07, M-08, M-09, M-10 |  |
| REQ-CL-0008 | The speed CSV has `server_ip` and `servers_all` (all servers of a multi-server test; Fast.com measures to up to five at once); the detailed CSV has "Server" and "Server address". | CHANGELOG.md@ad3758c:19 | R-02 | M-07 |  |
| REQ-CL-0009 | The dossier shows the server address next to the server name. | CHANGELOG.md@ad3758c:19 | R-03 |  |  |
| REQ-CL-0010 | The collector receives the server address of every speed test. | CHANGELOG.md@ad3758c:19 | S-12 | S-13 |  |
| REQ-CL-0011 | RTR rows are named after the measurement server RTR assigned, not after the engine. | CHANGELOG.md@ad3758c:19 | M-08 | R-02, R-03 |  |
| REQ-CL-0012 | Every speed test row records when that engine itself started (`ts`) and the start of the run it belonged to (`cycle_ts`). | CHANGELOG.md@ad3758c:22 | M-04 | C-02 | fault (all engines of a mixed run shared the run's start) |
| REQ-CL-0013 | Every speed test row stores how long the engine ran, the NMB version that measured and the box's network mode at the time; older rows stay empty and nothing is filled in afterwards. | CHANGELOG.md@ad3758c:22 | M-04 | C-05, C-11, C-02 |  |
| REQ-CL-0014 | The speed CSV carries row id, time zone, interface, local address, retransmits, TCP window and box ID as columns after the older ones, which keep their order; the ping CSV carries the target's role. | CHANGELOG.md@ad3758c:23 | R-02 |  |  |
| REQ-CL-0015 | Downloads `/export/outages.csv` (with whether the box caused the outage itself) and `/export/events.csv` (the event log; password only) exist. | CHANGELOG.md@ad3758c:23 | R-02 | S-02…S-0n (API: exports), C-07 |  |
| REQ-CL-0016 | The speed page has buttons for every export (11 languages), and `nmb-cli` has an action for each. | CHANGELOG.md@ad3758c:23 | U-02…U-0n (page: speed) | S-10, C-06 |  |
| REQ-CL-0017 | Every export takes a range (`?from=` / `?to=` as a date, ISO time or epoch ms, or `?days=` / `?hours=`), and its file name states the range. | CHANGELOG.md@ad3758c:23 | R-02 | S-02…S-0n (API: exports) |  |
| REQ-CL-0018 | Text from outside the box (server names, error texts, event details) can never act as a spreadsheet formula in an export. | CHANGELOG.md@ad3758c:23 | R-02 | R-03 | fault (CSV formula injection) |
| REQ-CL-0019 | The NMB version comes from the installed version file everywhere. | CHANGELOG.md@ad3758c:23 | C-05 | R-02, R-03, R-04 |  |
| REQ-CL-0020 | Optional SSH with keys only (Web UI System > SSH access, `nmb-cli run ssh.keys_only`), off by default; it can be switched on only when an account already has an SSH key. Texts in all 11 languages. | CHANGELOG.md@ad3758c:27 | NEW: P-16 SSH access | U-02…U-0n (page: system), S-10, C-06, S-02…S-0n (API: system) | Python reference: handlers in nmb/web/server.py |
| REQ-CL-0021 | The box checks that sshd really takes the keys-only setting (ahead of the base image's own SSH settings) and switches back by itself if not; Tailscale SSH keeps working. | CHANGELOG.md@ad3758c:27 | NEW: P-16 SSH access | P-06 |  |
| REQ-CL-0022 | Donation addresses come from one list whose checksums a test checks, and are shown in the READMEs, the Web UI (a Support button with copy buttons, a footer on every page), the logout page, the public results page, the bot (`/donate`, a line under `/start` and `/help`), `nmb-cli donate` (a line under `--help`) and at the end of the installer; texts in all 11 languages. | CHANGELOG.md@ad3758c:29 | P-08 | U-01, U-02…U-0n (page: results and logout), S-11, S-10, P-11, P-12, C-06 |  |
| REQ-CL-0023 | Donation addresses never appear in the SLA dossier or the CSV exports, which are evidence for the ISP. | CHANGELOG.md@ad3758c:29 | P-08 | R-02, R-03 |  |
| REQ-CL-0024 | The Wi-Fi regulatory domain offers every country of the Linux wireless regulatory database (181 in the version the source names) with its DFS rule set (ETSI, FCC, JP). | CHANGELOG.md@ad3758c:31 | N-07 | C-10 | fault (only 8 countries before) |
| REQ-CL-0025 | Country names are shown in the Web UI's language (from the browser), sorted by that name, with a short hint in all 11 languages. | CHANGELOG.md@ad3758c:31 | U-02…U-0n (page: Wi-Fi) | C-06, C-10 |  |
| REQ-CL-0026 | An unknown country code is refused by the Web UI's API. | CHANGELOG.md@ad3758c:31 | S-02…S-0n (API: Wi-Fi) | N-07 |  |
| REQ-CL-0027 | An unknown country code in `nmb.conf` falls back to automatic (hostapd would not start the access point with it). | CHANGELOG.md@ad3758c:31 | C-01 | N-07 |  |
| REQ-CL-0028 | Data caps for automatic speed tests (`[data_cap]`): a cap per link (Ethernet/PPPoE, Wi-Fi, USB, link aggregation) and per connection type (fiber, cable, DSL, mobile/FWA, dial-up, satellite, leased line, other; USB counts as mobile), per day or per month with a reset day. | CHANGELOG.md@ad3758c:35 | M-05 | C-01 |  |
| REQ-CL-0029 | The cap counts the data the tests themselves move on the measured interface; at 80 % an event; at 100 % automatic tests on that link stop until the period ends, while tests started by hand still run and count. | CHANGELOG.md@ad3758c:35 | M-05 | M-04, C-07 |  |
| REQ-CL-0030 | The data cap card sits under the schedule in the Web UI (11 languages); `nmb-cli datacap.save` sets it. | CHANGELOG.md@ad3758c:35 | U-02…U-0n (page: speed) | S-10, C-06 |  |
| REQ-CL-0031 | Self-test `datacap`: a reached cap skips a scheduled cycle without running an engine, and the counters are not touched. | CHANGELOG.md@ad3758c:35 | P-07 | M-05 |  |
| REQ-CL-0032 | The bot's `/status` shows the data the tests used per link in the current cap period, with the cap when one is set. | CHANGELOG.md@ad3758c:36 | S-11 | M-05 |  |
| REQ-CL-0033 | The bot alerts when a data cap stops the automatic tests on a link (which link, how much, when they start again) and when they run again, in all 11 languages. | CHANGELOG.md@ad3758c:37 | S-11 | M-05, C-06 |  |
| REQ-CL-0034 | The engines of the scheduled mixed run can be picked (`[speedtest] mixed_engines`, a tick per engine next to its interval, 11 languages, `nmb-cli`); tests started by hand still run them all. | CHANGELOG.md@ad3758c:38 | M-04 | U-02…U-0n (page: speed), S-10, C-06 |  |
| REQ-CL-0035 | `[probe] isp_first_hop = false` leaves the ISP's first hop out of the ping telemetry; a PPP peer placeholder address (one that never answers ping) is never taken as the first hop. (source names a specific site value; not copied) | CHANGELOG.md@ad3758c:39 | M-01 | C-01 |  |
| REQ-CL-0036 | Port forwarding in router mode, switchable on and off (`[port_forward]`, `pf.save` / `pf.rules`, 11 languages): rules for TCP, UDP or both, single ports or ranges, each rule on or off. | CHANGELOG.md@ad3758c:41 | N-11 | N-02, U-02…U-0n (page: router), S-10, C-06 | Python proof: live on a PPPoE line |
| REQ-CL-0037 | Hairpin NAT: LAN devices reach a forwarded port through the line's own address. | CHANGELOG.md@ad3758c:41 | N-11 | N-02 |  |
| REQ-CL-0038 | Port forwarding rules are checked against the LAN subnet, overlaps and the box's own VPN ports. | CHANGELOG.md@ad3758c:41 | N-11 | S-02…S-0n (API: network) |  |
| REQ-CL-0039 | UPnP IGD / NAT-PMP / PCP in router mode, off by default: devices may open ports only for themselves and only 1024-65535. | CHANGELOG.md@ad3758c:42 | N-11 | N-02, P-11 | Python used miniupnpd-nftables, installed with its own service off |
| REQ-CL-0040 | What devices opened via UPnP shows in the Web UI and is removed when UPnP or router mode stops. | CHANGELOG.md@ad3758c:42 | N-11 | U-02…U-0n (page: router) |  |
| REQ-CL-0041 | IPv6 for LAN and Wi-Fi clients in router mode on a PPPoE line (`[network] lan_ipv6`, on by default): the box asks the ISP for a delegated prefix, puts the /64 on the LAN bridge and advertises it (SLAAC, the box as DNS). | CHANGELOG.md@ad3758c:43 | N-11 | C-01 |  |
| REQ-CL-0042 | On PPPoE the box takes its own IPv6 address and route from the ISP's router advert on the PPP link. | CHANGELOG.md@ad3758c:43 | N-05 | N-11 |  |
| REQ-CL-0043 | LAN devices are not reachable from the internet over IPv6. | CHANGELOG.md@ad3758c:43 | N-02 | N-11 |  |
| REQ-CL-0044 | Scenario `pppoe`: router mode on a real PPPoE line (automatic mode chooses it, NAT, MSS clamp, the box's internet and DNS over the PPP link, speed tests over it, a line drop and a stopped pppd, every self-test group, DynDNS following the new address, a real Wi-Fi client and the Web UI, restore), guarded by the remote-access link. | CHANGELOG.md@ad3758c:44 | P-07 | N-04, N-05, N-02, M-04, N-16 |  |
| REQ-CL-0045 | Self-test `portfwd`: a temporary rule, the DNAT and hairpin path from a LAN device in a namespace, a NAT-PMP mapping and the low-port refusal, then the owner's settings back; the router group checks IPv6 for the LAN; the real-line PPPoE test reports the ISP's refusal text, the time to the first answer, HTTPS and IPv6 over the line. | CHANGELOG.md@ad3758c:45 | P-07 | N-11, N-05 |  |
| REQ-CL-0046 | Scenario `auto` (guided test of automatic mode with real cables, PPPoE on the ONT included), a `mode` check group, a "To look at" summary after every run and Markdown reports; scenarios restore the network mode too. | CHANGELOG.md@ad3758c:46 | P-07 | N-04 |  |
| REQ-CL-0047 | Automatic mode (`mode = auto`, the default for new installs, an "Automatic" choice in the Web UI in 11 languages): standalone while a router gives the box an address over DHCP; with several wired ports the other ports are bridged to the WAN port like a switch. | CHANGELOG.md@ad3758c:47 | N-04 | N-03, N-06, U-02…U-0n (page: router), C-06 |  |
| REQ-CL-0048 | Automatic mode goes to router mode when there has been no router and no internet for 90 s, dialling PPPoE when its login is saved and a PPPoE server answers; it goes back to standalone as soon as a router answers on the WAN port. | CHANGELOG.md@ad3758c:47 | N-04 | N-05 | Python: unit-tested only at this entry |
| REQ-CL-0049 | Existing installs keep their configured mode; only new installs default to automatic. | CHANGELOG.md@ad3758c:47 | C-01 | N-04 |  |
| REQ-CL-0050 | Self-test `ipv6`: reads the LAN router's IPv6 advertisement and says why the box has no IPv6 (for example no prefix advertised), then checks the box's public IPv6 address, route and HTTPS over IPv6. | CHANGELOG.md@ad3758c:48 | P-07 | N-11 |  |
| REQ-CL-0051 | Settings from a USB stick: `nmb.conf` at the top level of a stick is applied once when plugged in: only its keys, secrets never wiped by empty values, a backup first, a result file written back with names only. | CHANGELOG.md@ad3758c:50 | P-03 | C-01, C-04 |  |
| REQ-CL-0052 | A template for the USB settings file is shipped; `[usb_import] enabled = false` turns the import off. | CHANGELOG.md@ad3758c:50 | P-03 | C-01 |  |
| REQ-CL-0053 | Self-test `pppoe`: the box's own pppd command against a throwaway PPPoE server (right password connects and carries traffic, wrong password is refused, the password is on no command line); the real-line variant dials on the WAN port in a separate namespace. | CHANGELOG.md@ad3758c:51 | P-07 | N-05 | Python proof: passed on the Pi 4 (simulated line) |
| REQ-CL-0054 | Scenarios stop and restore when the remote-access link (Tailscale) is offline for 2 minutes; the `internet` group checks Tailscale; the `router` group checks the Web UI and SSH from the LAN. | CHANGELOG.md@ad3758c:52 | P-07 | N-15 |  |
| REQ-CL-0055 | Scenario `router`: guided live test of router mode with the Ethernet port as WAN (a simulated client, a real device on the box's Wi-Fi and its traffic counted, the Web UI from that device, a WAN cable loss and return, restore); stopping a scenario through the service manager restores the settings too. | CHANGELOG.md@ad3758c:53 | P-07 | N-03 |  |
| REQ-CL-0056 | Scenario `usb`: guided live test with USB tethering (speed tests bound to USB, Wi-Fi shared on the LAN port, USB fallback, router mode on USB with real clients, the DHCP guard, restore, unplugging); it waits for each cable move, counts the clients' forwarded packets and always puts the settings back (also after 45 minutes by a separate service). | CHANGELOG.md@ad3758c:54 | P-07 | N-10, N-08, N-05 |  |
| REQ-CL-0057 | Self-test live checks on a running box (services, internet, settings, web and terminal interfaces, USB tethering, router mode with a simulated LAN client, the DHCP guard against a throwaway server, speed engines against each other, the access point, the collector, and with `--full` the bot in all 11 languages) use throwaway namespaces only and never print secrets. | CHANGELOG.md@ad3758c:55 | P-07 | F-06 | See TESTING.md 2a |
| REQ-CL-0058 | USB tethering as the only internet: with no wired or Wi-Fi default route the box adds one through USB (metric 900) and removes it when another is back; both switches are events. | CHANGELOG.md@ad3758c:57 | N-10 | N-01, C-07 |  |
| REQ-CL-0059 | In router mode, USB becomes the WAN when no DHCP server answers on the Ethernet port, which then joins the LAN with the access point; a DHCP server on that port stops the box's DHCP at once. | CHANGELOG.md@ad3758c:57 | N-10 | N-03, N-05 |  |
| REQ-CL-0060 | LAN-port sharing takes USB as its upstream when the Wi-Fi uplink is down. | CHANGELOG.md@ad3758c:57 | N-08 | N-10 |  |
| REQ-CL-0061 | USB tethering (`[usb_uplink]`): a phone, mobile hotspot or modem on USB as a measuring link with its own profile that never takes the default route, a source rule (table 152), the measurement link `usb`, its own daily data limit, results stored as mobile; Web UI card, `nmb-cli` (`usb.uplink`, `usb.off`), 11 languages. | CHANGELOG.md@ad3758c:59 | N-10 | M-04, M-05, C-01, U-02…U-0n (page: network), S-10, C-06 | Python proof: Pi 4 over 5G with a USB hotspot (source names a specific site value; not copied) |
| REQ-CL-0062 | A Disk space card in the settings: automatic cleanup on/off, minimum free space, days of measurements always kept, "Check and clean up now" and the current free space, in all 11 languages. | CHANGELOG.md@ad3758c:66 | U-02…U-0n (page: system) | C-08, C-06 |  |
| REQ-CL-0063 | `nmb-cli` has the automatic disk cleanup settings and "clean up now" (`storage.save`, `storage.cleanup`, `/api/storage/*`), the main Wi-Fi client port role in `ports.roles`, and the diagnostic bundle and deep snapshot downloads. | CHANGELOG.md@ad3758c:67 | S-10 | S-02…S-0n (API: system), C-08, R-04, N-17 |  |
| REQ-CL-0064 | Eleven languages (including Romanian, Italian, Spanish, Portuguese for Portugal and Brazil, Simplified Chinese) in every Web UI text, the logout page, the terminal UI, the results page, the charts and the report; the bot's menus, answers, language menu, admin answers and alerts too. | CHANGELOG.md@ad3758c:71 | C-06 | U-01, U-02…U-0n (page: results and logout), S-10, S-11, R-03 | fault family: missing translations |
| REQ-CL-0065 | The documentation explains why scheduled Ookla tests run at most once an hour: Ookla blocks a public IP address that tests too often. | CHANGELOG.md@ad3758c:72 | P-12 | M-04, M-10 |  |
| REQ-CL-0066 | Storage guard (`[storage]`, on by default): before the disk fills up, the box deletes the oldest logs, backups, reports and journal, and only then the oldest measurements (the last 7 days are kept); the free space it keeps depends on the disk size. | CHANGELOG.md@ad3758c:74 | C-08 | C-01, C-02 |  |
| REQ-CL-0067 | The installer stops with a clear message when less than 500 MB is free and uses smaller logs on disks under 8 GB. | CHANGELOG.md@ad3758c:74 | P-11 | C-08, C-03 |  |
| REQ-CL-0068 | Wi-Fi shared on a LAN port (`[wifi_share]`, off by default): when the box is on Wi-Fi and the port has no internet of its own, devices on that port get addresses from the box (a default private /24) and reach the internet through the box's Wi-Fi. | CHANGELOG.md@ad3758c:79 | N-08 | U-02…U-0n (page: network), S-10, C-01 |  |
| REQ-CL-0069 | The box never hands out addresses where another DHCP server answers: it checks before starting, every minute and right after a cable is plugged in (stopping its own DHCP first). | CHANGELOG.md@ad3758c:79 | N-05 | N-08 |  |
| REQ-CL-0070 | Each wireless interface can be the guest Wi-Fi client or the main Wi-Fi client (measuring), never both; taking the guest client's radio turns the guest client off; the automatic choice of the main client skips the guest client's radio. | CHANGELOG.md@ad3758c:80 | N-08 | N-07, N-17, U-02…U-0n (page: ports and drivers) |  |
| REQ-CL-0071 | Fast.com engine through Fast.com's own API: up to five nearby servers, latency to each, parallel streams over them for download and upload, the "stable" speed Fast.com shows, latency under load and bufferbloat. | CHANGELOG.md@ad3758c:82 | M-07 |  | The source line says both "8 streams" and, later, "16 streams"; the stream count to check in the code is 16 (REQ-CL-0075, README REQ-RE-0079). Count dropped here in round 1. |
| REQ-CL-0072 | Each Fast.com row keeps the ISP, ASN and location Fast.com sees, every server with its latency and share of the transfer, and the speed every half second. | CHANGELOG.md@ad3758c:82 | M-07 | C-02, R-02 |  |
| REQ-CL-0073 | The home page shows a Fast.com table: how Fast.com sees the box (ISP, ASN, place), the stable, mean and peak speed, the latency under load and every server with its latency and share. | CHANGELOG.md@ad3758c:82 | U-02…U-0n (page: home) | M-07 |  |
| REQ-CL-0074 | Fast.com is available in the Web UI, the bot (`/test fastcom`), the speed command (`--engine fastcom`) and `nmb-cli`, and is part of the mixed run (`fastcom_interval_min`, 0 by default). | CHANGELOG.md@ad3758c:82 | M-07 | M-04, U-02…U-0n (page: speed), S-11, S-10 |  |
| REQ-CL-0075 | Fast.com transfers use plain HTTP with 16 streams for 15 s each way. | CHANGELOG.md@ad3758c:82 | M-07 |  | Pi 4: TLS used most of the CPU; the same line also says 8 streams earlier (see REQ-CL-0071): 16 is the figure. |
| REQ-CL-0076 | iPerf3 rows that stored only an address show the server's name in the statistics, the history and the CSV: from the public iPerf3 list, else city, country and ISP of the address from GeoIP. | CHANGELOG.md@ad3758c:83 | M-09 | R-02, U-02…U-0n (page: speed) |  |
| REQ-CL-0077 | Wi-Fi uplink for measuring (`[wifi_uplink]`): the box joins the owner's Wi-Fi with its own profile, separate from the guest Wi-Fi client. | CHANGELOG.md@ad3758c:84 | N-08 | C-01, S-10 |  |
| REQ-CL-0078 | `[network] measurement_link` chooses where speed tests run: `auto` (Ethernet when a cable is connected, else Wi-Fi), `ethernet` or `wifi`; with both links up, tests bound to Wi-Fi really leave over Wi-Fi (source routing). | CHANGELOG.md@ad3758c:84 | M-04 | N-01, N-08, U-02…U-0n (page: speed), S-10 |  |
| REQ-CL-0079 | iPerf3 direction "Download, then upload" (the default in the iPerf3 tab) is shown as one result. | CHANGELOG.md@ad3758c:85 | M-09 | U-02…U-0n (page: speed) |  |
| REQ-CL-0080 | iPerf3 results name the server (place and provider, from the public list or GeoIP) instead of only its address. | CHANGELOG.md@ad3758c:86 | M-09 | M-12 |  |
| REQ-CL-0081 | Several iPerf3 servers for automatic tests (`[iperf] extra_servers`, "Add to the automatic list", `nmb-cli run iperf.auto`) are tested one after the other, each in its own row; servers that do not answer are left out. | CHANGELOG.md@ad3758c:87 | M-09 | M-04, S-10 |  |
| REQ-CL-0082 | "Use this server for automatic tests" saves the iPerf3 server with its streams, duration, protocol and direction. | CHANGELOG.md@ad3758c:88 | M-09 | U-02…U-0n (page: speed) |  |
| REQ-CL-0083 | `nmb-cli latest [engine]` gives the newest result of every engine, or of one. | CHANGELOG.md@ad3758c:89 | S-10 |  |  |
| REQ-CL-0084 | iPerf3 can be the only automatic engine and joins the mixed run when a server is set and answers, with its own interval (`iperf3_interval_min`); without a working server it is left out without an empty row. | CHANGELOG.md@ad3758c:90 | M-04 | M-09 |  |
| REQ-CL-0085 | Collector client, on only when the owner turns it on: sends speed tests and hourly ping summaries; with separate consent also events and hourly full diagnostics, and remote control (speed test, fresh diagnostics, upload; nothing else). | CHANGELOG.md@ad3758c:91 | S-12 | U-02…U-0n (page: system), S-10 |  |
| REQ-CL-0086 | The collector finds boxes by a permanent unique ID, name, current or earlier case number, customer number or ISP. | CHANGELOG.md@ad3758c:91 | S-13 |  |  |
| REQ-CL-0087 | Each engine has its own minimum interval for scheduled tests (`ookla_interval_min` at least 60; `rtr_interval_min`, `cloudflare_interval_min`), set under the schedule. | CHANGELOG.md@ad3758c:92 | M-04 | U-02…U-0n (page: speed), C-01 |  |
| REQ-CL-0088 | A "Refresh from Ookla" button next to the Ookla server choice on the home page and in the server list. | CHANGELOG.md@ad3758c:93 | M-10 | U-02…U-0n (page: home) |  |
| REQ-CL-0089 | Update in a few clicks (Web UI, `nmb-cli update --apply --yes`, the heal command): download the newest version, run its tests, back up settings and measurements, install, and put the previous version back if the install fails; the Web UI shows each step. | CHANGELOG.md@ad3758c:94 | P-02 | U-02…U-0n (page: system), S-10, P-01 | Native: signed packages instead of a Git download (owner decision) |
| REQ-CL-0090 | The install is proven on each supported distribution: clean install, reinstall and reboot (Python: Ubuntu 26.04 LTS, Alpine 3.22, Rocky Linux 10). | CHANGELOG.md@ad3758c:95 | P-11 |  | question for the owner |
| REQ-CL-0091 | Fedora and RHEL family support: the package-manager branch (dnf, EPEL on RHEL and rebuilds), the NMB ports opened in firewalld; the helper's upgrade and update check work with dnf and apk. | CHANGELOG.md@ad3758c:96 | P-11 | P-06 | question for the owner |
| REQ-CL-0092 | Optional results without password (`public_results = true`, off by default): devices on the local network open `/results` and download the report and CSV files without logging in; settings still need the password. | CHANGELOG.md@ad3758c:97 | S-01 | S-02…S-0n (API: exports), U-02…U-0n (page: results and logout), C-01 | Owner moved from S-02…S-0n (API: exports) in round 1: one owner (S-01, which decides the routes that need a login) with REQ-TS-0125 and REQ-RE-0303. |
| REQ-CL-0093 | A warning in the Web UI, the API and the log when the schedule runs Ookla more often than hourly, showing how long Ookla stays paused. | CHANGELOG.md@ad3758c:98 | M-05 | U-02…U-0n (page: speed), S-02…S-0n (API: speed), C-03 |  |
| REQ-CL-0094 | A real DMZ in router mode: its own subnet and DHCP, internet access, and no access to the LAN or the box beyond DHCP and DNS. | CHANGELOG.md@ad3758c:102 | N-02 | N-03, N-05 |  |
| REQ-CL-0095 | A DNS interception test and a NAT mapping check (endpoint-independent or symmetric) are in the diagnostic bundle. | CHANGELOG.md@ad3758c:103 | M-12 | R-04 |  |
| REQ-CL-0096 | Port checks cover ports 22, 5060 and 6881. | CHANGELOG.md@ad3758c:104 | M-11 |  |  |
| REQ-CL-0097 | Telegram alerts for high latency and high jitter exist. | CHANGELOG.md@ad3758c:105 | S-11 | M-01, M-02 | fault (latency alerts were documented but did not exist) |
| REQ-CL-0098 | A daily data limit for scheduled tests, for mobile and other metered connections. | CHANGELOG.md@ad3758c:106 | M-05 | M-04 |  |
| REQ-CL-0099 | The default contracted speed is 0 (not set); the settings form fills in no value; nothing is judged against a plan the owner never entered; a configuration that names its speeds is not affected. | CHANGELOG.md@ad3758c:115 | C-01 | R-01, U-02…U-0n (page: system) | fault (a default plan was judged against); (source names a specific site value; not copied) |
| REQ-CL-0100 | The bot's `/report` caption, progress message and help never call the box's evaluation "official" or "regulatory" in any language: it is NMB's evaluation of its own measurements. | CHANGELOG.md@ad3758c:116 | S-11 | C-06 | fault |
| REQ-CL-0101 | The dossier and the email summary say that only scheduled tests wait for low traffic; tests started by hand (Web UI, Telegram, remote) start at once (11 languages). | CHANGELOG.md@ad3758c:117 | R-03 | R-05, C-06 |  |
| REQ-CL-0102 | Ookla runs only after the owner accepted Ookla's licence and privacy terms (links to licence, privacy policy and terms of use, an accept button, a way to withdraw); until then Ookla is skipped (not stored, nothing stands in) and the other engines measure; existing installs start as not accepted. | CHANGELOG.md@ad3758c:118 | M-10 | U-02…U-0n (page: speed), C-01 | fault (terms were accepted silently for the owner) |
| REQ-CL-0103 | The Wi-Fi access point runs only in router mode; in standalone and bridge mode the radio stays a Wi-Fi client of the home network and is handed back and reconnected if an access point held it. | CHANGELOG.md@ad3758c:119 | N-07 | N-03 | fault (an access point in standalone cut the box off its Wi-Fi) |
| REQ-CL-0104 | No device (router, Raspberry Pi 4, x86-64, own distribution) ships or needs Python. | CHANGELOG.md@ad3758c:120 | P-11 | P-10, P-13 | Same rule as REQ-RE-0191 (one owner, P-11). |
| REQ-CL-0105 | Scheduled iPerf3 tests run at most hourly by default (`iperf3_interval_min = 60`), since public servers take one client at a time. | CHANGELOG.md@ad3758c:121 | M-09 | M-04, C-01 |  |
| REQ-CL-0106 | The home page only shows results: buttons choose the newest result of any engine (mixed) or of one engine; engine, Ookla server, interface and schedule are set under Measurement & Servers, which the home page links to; in `nmb-cli` the schedule is on the same screen. | CHANGELOG.md@ad3758c:122 | U-02…U-0n (page: home) | U-02…U-0n (page: speed), S-10 |  |
| REQ-CL-0107 | Recommended iPerf3 defaults for automatic tests: 4 parallel TCP streams, 10 s, download then upload (`auto_direction = both`), with a short pause and a retry before the upload. | CHANGELOG.md@ad3758c:123 | M-09 | C-01 |  |
| REQ-CL-0108 | The Ookla server choice comes only from Ookla's current list of servers near the box; no built-in server list. | CHANGELOG.md@ad3758c:124 | M-10 | C-10 | fault (built-in list had wrong IDs) (source names a specific site value; not copied) |
| REQ-CL-0109 | The Wi-Fi country defaults to automatic. | CHANGELOG.md@ad3758c:124 | N-07 | C-10, C-01 |  |
| REQ-CL-0110 | No country is built in: ISP names, examples and placeholders name no country's providers. | CHANGELOG.md@ad3758c:124 | C-10 | C-06, U-01 | QB 9 |
| REQ-CL-0111 | Status codes are neutral codes with English descriptions, not tied to one regulator. | CHANGELOG.md@ad3758c:124 | R-01 | R-03, R-02 |  |
| REQ-CL-0112 | The Ookla server choice is labelled as the Ookla Speedtest server everywhere, with a note that RTR-NetTest, Cloudflare and iPerf3 choose their own servers. | CHANGELOG.md@ad3758c:125 | U-02…U-0n (page: speed) | C-06, M-10 |  |
| REQ-CL-0113 | Scheduled Ookla tests run at most once an hour however often the schedule runs; RTR-NetTest, Cloudflare and iPerf3 still run every time; manual Ookla tests are not held back. | CHANGELOG.md@ad3758c:126 | M-04 | M-05, M-10 |  |
| REQ-CL-0114 | The speed test list does not show skipped runs (busy network, manual mode); a line says how many were left out; failed runs show the reason instead of 0.00 Mbps. | CHANGELOG.md@ad3758c:127 | U-02…U-0n (page: speed) | S-02…S-0n (API: speed), R-01 | fault (failed runs showed 0.00 Mbps) |
| REQ-CL-0115 | The web terminal and the task manager never act as root; the web server runs as an unprivileged user. | CHANGELOG.md@ad3758c:128 | P-04 | U-02…U-0n (page: tools), S-01, C-09 | security |
| REQ-CL-0116 | The editable boot partition copy of the settings uses `__KEEP_CURRENT__` for passwords and tokens. | CHANGELOG.md@ad3758c:129 | C-01 | C-04 |  |
| REQ-CL-0117 | The diagnostic bundle's MTR runs few enough cycles to finish quickly without root. | CHANGELOG.md@ad3758c:130 | R-04 | M-11 |  |
| REQ-CL-0118 | The READMEs state what has been tested and describe features as they work (for example a bridge needs an IP address). | CHANGELOG.md@ad3758c:132 | P-12 | P-13 | QB 1 honest results |
| REQ-CL-0119 | The documentation describes the security model and which services run as root. | CHANGELOG.md@ad3758c:133 | P-12 | P-04, P-06 |  |
| REQ-CL-0120 | The Web UI header shows the link actually in use and the path to the ISP from the recorded setup, and the real system time; it never shows a fixed interface name or `--` for fields the server does not send (page and server agree on fields). | CHANGELOG.md@ad3758c:137 | U-01 | C-11, S-02…S-0n (API: status) | fault |
| REQ-CL-0121 | The latency chart's labels are translated in every language. | CHANGELOG.md@ad3758c:137 | U-02…U-0n (page: home) | C-06 | fault (Hungarian labels in every language) |
| REQ-CL-0122 | The bot's `/status` and `/mode` show the mode in words, never the raw mode code. | CHANGELOG.md@ad3758c:138 | S-11 | N-03, C-06 | fault |
| REQ-CL-0123 | The bot's `/ai` uses the diagnostic bundle's cache instead of rebuilding it on every request. | CHANGELOG.md@ad3758c:138 | S-11 | R-04 | fault |
| REQ-CL-0124 | The live bot self-test leaves no test chats' language rows in the database. | CHANGELOG.md@ad3758c:138 | P-07 | S-11 | fault |
| REQ-CL-0125 | The diagnostic bundle states the active mode with the configured one beside it, never the configured mode alone (`auto` read as "LAN probe behind a NAT router" when the box was the router). | CHANGELOG.md@ad3758c:139 | R-04 | N-03, N-04 | fault |
| REQ-CL-0126 | Without ping data the uptime is shown as unknown ("-", "--" or "N/A") in the dossier, the email, the bot (`/status`, `/report`), the API and the report command, never 100 %. | CHANGELOG.md@ad3758c:140 | R-01 | R-03, R-05, S-11, S-02…S-0n (API: status), S-10 | fault |
| REQ-CL-0127 | The SLA compliance is "--" when no test could be judged. | CHANGELOG.md@ad3758c:140 | R-01 | R-03 | fault |
| REQ-CL-0128 | The dossier's measurement success rate is "-" without tests and counts the line's tests only. | CHANGELOG.md@ad3758c:140 | R-03 | R-01 | fault |
| REQ-CL-0129 | `/ping` shows "--" for a target without samples and lists switched-on targets only. | CHANGELOG.md@ad3758c:140 | S-11 | M-01 | fault (showed 100 % loss) |
| REQ-CL-0130 | `/test` says when a guaranteed speed is set but the test did not measure that direction (11 languages). | CHANGELOG.md@ad3758c:140 | S-11 | R-01, C-06 | fault |
| REQ-CL-0131 | The local gateway does not count into the line's loss, latency and uptime timeline; it stays in the per-target table. | CHANGELOG.md@ad3758c:141 | M-02 | R-01, M-01 | fault |
| REQ-CL-0132 | An outage that began before the report window is put on the line when its real start lies outside the box's own windows (not judged by the clamped start). | CHANGELOG.md@ad3758c:142 | R-01 | M-02, C-07 | fault |
| REQ-CL-0133 | LAN tests: the speed CSV leaves the share of the offered speed empty and names the ISP "LAN". | CHANGELOG.md@ad3758c:143 | R-01 | R-02, R-03, M-09 | fault |
| REQ-CL-0134 | The Web UI speed chart leaves LAN tests out and shows an unmeasured direction as a gap, not 0 Mbps. | CHANGELOG.md@ad3758c:143 | U-02…U-0n (page: speed) | R-01 | fault |
| REQ-CL-0135 | The home card of an iPerf3 test against a LAN server shows no percentage; the 24 h SLA averages show "--" without data. | CHANGELOG.md@ad3758c:143 | U-02…U-0n (page: home) | R-01 | fault |
| REQ-CL-0136 | The collector uses the box's own LAN rule and counts the "below 80 %" share only among tests that measured download. | CHANGELOG.md@ad3758c:143 | S-13 | S-12, R-01 | fault |
| REQ-CL-0137 | The LLM bundle: no 0.00 % loss for a target without samples; a "no ping data" verdict; "gateway does not answer" only when internet targets have data; unset contract and guaranteed speeds as "not set"; LAN tests apart in engine statistics; upload-only tests counted as successful. | CHANGELOG.md@ad3758c:144 | R-04 | R-01 | fault |
| REQ-CL-0138 | RTR: a direction without transferred bytes is "not measured", not 0 bit/s (native and fallback client); a run without any direction measured is an error. | CHANGELOG.md@ad3758c:145 | M-08 | R-01 | fault |
| REQ-CL-0139 | The iPerf3 command record names the directions that really ran. | CHANGELOG.md@ad3758c:145 | M-09 |  | fault |
| REQ-CL-0140 | Cloudflare never stores the idle latency as the loaded latency (also not with a rounding difference), and never stores 0 bit/s for an unmeasured direction of a successful test. | CHANGELOG.md@ad3758c:146 | M-06 | C-02 | fault |
| REQ-CL-0141 | A migrated Python database keeps the correction markers (for example `data_correction_2026_11_10c`) and the originals in `speedtest_correction`. | CHANGELOG.md@ad3758c:146 | C-02 | F-05 | Python data migration |
| REQ-CL-0142 | The streamed ping CSV reads in key pages, each finished before it is written (no read transaction open during a slow download); a client that stops reading times out. | CHANGELOG.md@ad3758c:147 | R-02 | S-02…S-0n (API: exports), C-02 | fault |
| REQ-CL-0143 | A logged-in admin gets the 120-day range also on the public ping path. | CHANGELOG.md@ad3758c:147 | R-02 | S-01 | fault |
| REQ-CL-0144 | The email summary shows unset contracted speeds as "not set", not "0 Mbps". | CHANGELOG.md@ad3758c:148 | R-05 | R-01 | fault |
| REQ-CL-0145 | The "Change password / Settings" button of the default-password notice opens the administrator credentials settings with the cursor in the password field. | CHANGELOG.md@ad3758c:150 | U-02…U-0n (page: system) | U-01 | fault |
| REQ-CL-0146 | A test checks that every tab and subtab button points to one that exists. | CHANGELOG.md@ad3758c:150 | U-01 |  | fault |
| REQ-CL-0147 | The SLA dossier shows the installed version only, with no fixed version number and no claim of conformity to a standard that nobody checked. | CHANGELOG.md@ad3758c:151 | R-03 | C-05 | fault |
| REQ-CL-0148 | The dossier appendix (11 languages) says what the box does about background traffic: it waits until the other traffic it can see is low; traffic it cannot see (for example devices behind another router) is not excluded. | CHANGELOG.md@ad3758c:151 | R-03 | M-03, C-06 | fault |
| REQ-CL-0149 | The dossier cites no national act (the report is not tied to one country); the EU regulation stays. | CHANGELOG.md@ad3758c:151 | R-03 | C-10 | fault; QB 9 |
| REQ-CL-0150 | The dossier escapes server names and outage causes exactly once. | CHANGELOG.md@ad3758c:152 | R-03 |  | fault |
| REQ-CL-0151 | The per-minute ping CSV is streamed in pieces, never built in memory (up to 120 days, about a million rows). | CHANGELOG.md@ad3758c:153 | R-02 | C-02 | fault |
| REQ-CL-0152 | A minute without pings has no loss figure in the ping CSV, not 0. | CHANGELOG.md@ad3758c:153 | R-02 | M-01 | fault |
| REQ-CL-0153 | With no pings and no judged test the SLA verdict is "unknown"; compliance counts only tests judged against a required speed; a direction a test did not measure is never a failure. | CHANGELOG.md@ad3758c:155 | R-01 | R-03, S-11, R-04 | fault |
| REQ-CL-0154 | Values that were not measured show "-" or "--", never 0, in the dossier (statistics table, KPIs, per-target table), the Web UI (speed table, averages per direction, SLA card, chart tooltips), the bot (`/status`, `/test` and its per-engine lines), the diagnostic bundle, the speed test event text and the developer snapshot. | CHANGELOG.md@ad3758c:156 | R-01 | R-03, U-02…U-0n (page: speed), U-02…U-0n (page: home), S-11, R-04, C-07, C-03 | fault |
| REQ-CL-0155 | An unmeasured upload is never reported as an "upload deficit". | CHANGELOG.md@ad3758c:156 | S-11 | R-01 | fault |
| REQ-CL-0156 | In the diagnostic bundle the last test is the newest successful one; the recent-runs table has a status column and no speeds for skipped or failed runs; the developer snapshot reads columns that exist. | CHANGELOG.md@ad3758c:156 | R-04 | R-01 | fault |
| REQ-CL-0157 | The dossier shows no invented packet loss minimum or maximum. | CHANGELOG.md@ad3758c:156 | R-03 | R-01 | fault (invented figures) |
| REQ-CL-0158 | Tests against an iPerf3 server on the local network (also older ones, recognised by the server's local address) are marked LAN everywhere and left out of the line's figures (dossier and charts, detailed CSV with no status and ISP "LAN", Web UI averages and SLA card, bot, diagnostic bundle, collector). | CHANGELOG.md@ad3758c:157 | R-01 | R-03, R-02, U-02…U-0n (page: speed), S-11, R-04, S-13, M-09 | fault |
| REQ-CL-0159 | A local iPerf3 server is never named after the box's own town and ISP. | CHANGELOG.md@ad3758c:157 | M-09 | M-12 | fault |
| REQ-CL-0160 | The ISP name per test comes from the settings or from the test's own public address, never from a lookup of today's address. | CHANGELOG.md@ad3758c:158 | M-04 | M-12 | fault |
| REQ-CL-0161 | The ISP ticket email describes the traffic gate and the required speed as they are set, never fixed thresholds. | CHANGELOG.md@ad3758c:159 | R-05 | M-03 | fault |
| REQ-CL-0162 | The bundle never calls a gateway that does not answer pings "loss inside the local network". | CHANGELOG.md@ad3758c:160 | R-04 |  | fault |
| REQ-CL-0163 | A skip because other devices' traffic is not visible is code 10, not 9 ("other traffic"). | CHANGELOG.md@ad3758c:161 | R-01 | M-03 | fault |
| REQ-CL-0164 | No invented values survive a migration: 0 ms Cloudflare jitter from a single sample, the RTR fallback's invented ping, Cloudflare "loaded" latency equal to the idle one (and its grade), LAN iPerf3 rows not marked LAN; originals stay in `speedtest_correction`. | CHANGELOG.md@ad3758c:162 | C-02 | M-06, M-08, M-09 | fault; Python data migration |
| REQ-CL-0165 | A rebuild the box makes in reaction to the line (USB failover in or out, PPPoE appearing or going in auto mode) is recorded but not taken out of the line's outage; an outage counts as the box's own only when it began while the box had its connection down. | CHANGELOG.md@ad3758c:163 | M-02 | C-07, N-10, N-04 | fault |
| REQ-CL-0166 | An engine's interval counts from the start of the run it belonged to. | CHANGELOG.md@ad3758c:164 | M-04 |  | fault (hourly iPerf3 skipped every other hour) |
| REQ-CL-0167 | Aborted runs store their facts too. | CHANGELOG.md@ad3758c:164 | M-04 | C-02 | fault |
| REQ-CL-0168 | RTR fallback rows say where their address came from. | CHANGELOG.md@ad3758c:164 | M-08 |  | fault |
| REQ-CL-0169 | A silent ISP first hop stays unpinged (remembered for 30 days) instead of being switched on again once its lost pings aged out. | CHANGELOG.md@ad3758c:165 | M-01 |  | fault |
| REQ-CL-0170 | The collector's `/api/poll` checks the box-to-token binding, so another token cannot take, fake or refuse a box's remote commands. | CHANGELOG.md@ad3758c:166 | S-13 | S-12 | fault; security |
| REQ-CL-0171 | The collector receives the per-test facts (`cycle_ts`, duration, version, mode, `client_ip_source`), leaves LAN tests out of its averages and shows missing values as "--". | CHANGELOG.md@ad3758c:166 | S-12 | S-13, M-04 | fault |
| REQ-CL-0172 | The per-minute ping export without login (public results) covers at most 31 days, with login 120. | CHANGELOG.md@ad3758c:167 | S-02…S-0n (API: exports) | R-02, S-01 | fault |
| REQ-CL-0173 | iPerf3 client address: against an internet server the row has the line's public address (asked over the same source address after the test behind NAT, or the test's own address when already public); against a local server it has the box's local address and the test is marked LAN and left out of the line's figures. | CHANGELOG.md@ad3758c:168 | M-09 | M-12 | fault |
| REQ-CL-0174 | A failed public-address lookup leaves the address empty, never the LAN address in its place. | CHANGELOG.md@ad3758c:168 | M-09 | M-08 | fault |
| REQ-CL-0175 | Every row says where its client address came from (`client_ip_source`: `engine`, `lookup`, `socket` or `local`; last column of the speed CSV); RTR uses the same lookup. | CHANGELOG.md@ad3758c:168 | M-04 | R-02, M-08, M-09 | fault |
| REQ-CL-0176 | A test meets the required speed when it reaches the guaranteed speed the owner set; without one, 80 % of the offered speed; a direction the test did not measure is not counted as failed. | CHANGELOG.md@ad3758c:169 | R-01 | R-03 | fault |
| REQ-CL-0177 | The bot's result states use the dossier's status numbers (6 = below the guaranteed speed, with "download" or "upload" in the text), never the Wi-Fi codes. | CHANGELOG.md@ad3758c:170 | S-11 | R-01, C-06 | fault |
| REQ-CL-0178 | A box belongs to the token it first uploaded with; another valid token cannot write into it until the first token is revoked. | CHANGELOG.md@ad3758c:171 | S-13 | S-12 | fault; security |
| REQ-CL-0179 | The box sends the corrections of values that were filled in instead of measured (`speedtest_correction`) for rows it sent earlier, and the collector clears them in its copy too. | CHANGELOG.md@ad3758c:171 | S-12 | S-13, C-02 | fault |
| REQ-CL-0180 | The diagnostic bundle concludes only what its data shows. | CHANGELOG.md@ad3758c:172 | R-04 |  | fault; lead-in of sub-bullets 173-178 |
| REQ-CL-0181 | A private hop 2 gives "possibly double NAT" with the reason a traceroute cannot tell; a hop 2 that is a known PPP peer placeholder is called that; CGNAT is "consistent with", not "detected". | CHANGELOG.md@ad3758c:173 | R-04 | M-11, N-05 | fault |
| REQ-CL-0182 | The bundle's LAN, ISP-hop and transit verdicts use only the loss left once the box's own actions and recorded outages are taken out, say how many bursts fell into them, and name possible causes instead of a certainty; a target that answers almost nothing gets no "max streak". | CHANGELOG.md@ad3758c:174 | R-04 | C-07, M-02 | fault |
| REQ-CL-0183 | The path MTU is "unknown" when nothing answers; the advice depends on who clamps the MSS (this box in router mode, or the router in front). | CHANGELOG.md@ad3758c:175 | R-04 | N-02 | fault |
| REQ-CL-0184 | The MTR is "loaded" only when a speed test really ran during it, in any process on the box. | CHANGELOG.md@ad3758c:176 | R-04 | M-04 | fault |
| REQ-CL-0185 | The bundle's events section uses the stored severities, and lists warnings, outages, the box's own actions and mode changes. | CHANGELOG.md@ad3758c:177 | R-04 | C-07 | fault |
| REQ-CL-0186 | The bundle's version is the installed one. | CHANGELOG.md@ad3758c:178 | R-04 | C-05 | fault |
| REQ-CL-0187 | An ISP first hop that answered none of at least 2,000 pings in 24 hours is no longer pinged as a target; a PPP peer placeholder seen while the box dialled PPPoE is remembered for 90 days. | CHANGELOG.md@ad3758c:179 | M-01 | N-05 | fault (made up "100 % loss") |
| REQ-CL-0188 | The network daemon records every window in which it takes its own connection down (mode change, rebuild after a settings change, PPPoE on or off in auto mode, a WAN change in router mode, its own stop, which covers updates and reboots) in a table, with an event at the start and at the end. | CHANGELOG.md@ad3758c:180 | C-07 | N-03, N-04, N-05, P-02, P-10 | fault (box's own outages counted against the ISP) |
| REQ-CL-0189 | Reports take these windows (plus 2 minutes for the line to come back) out of every outage: an outage inside them is the box's own and listed apart ("caused by the box itself; not counted"), a longer one counts only for the rest; the windows are left out of the measured time and pings lost in them are not packet loss. | CHANGELOG.md@ad3758c:180 | M-02 | C-07, R-01, R-03 | fault |
| REQ-CL-0190 | For time before own-action windows were recorded, the mode change events stand in (each marks the end of a switch that began up to 3 minutes earlier). | CHANGELOG.md@ad3758c:180 | C-07 | C-02 | Python data migration |
| REQ-CL-0191 | The outages CSV has `cause_kind`, `counted_sec` and `own_action_sec`; the dossier (11 languages) and the report API show the split. | CHANGELOG.md@ad3758c:180 | R-02 | C-07 | fault |
| REQ-CL-0192 | No outage alarm is sent while the box itself has the line down; it is sent if the line is still down afterwards. | CHANGELOG.md@ad3758c:180 | S-11 | C-07, M-02 | fault |
| REQ-CL-0193 | Values filled in instead of measured are never stored or exported; already stored ones are cleared once with each original kept with its reason (`speedtest_correction`), so it can be checked or undone. | CHANGELOG.md@ad3758c:181 | M-04 | C-02 | fault |
| REQ-CL-0194 | Packet loss from engines that do not measure it (Cloudflare, RTR, Ookla without a loss figure) is empty (unknown), never 0 %. | CHANGELOG.md@ad3758c:182 | M-04 | M-06, M-08, M-10 | fault |
| REQ-CL-0195 | RTR stores no fixed jitter; jitter is the mean change between pings from the server's timings; no estimated latency under load or bufferbloat; no fixed place or colo; no 0 ms when no ping came back; no project page as the result link; the measurement server RTR assigned is stored instead of the control host. | CHANGELOG.md@ad3758c:183 | M-08 |  | fault |
| REQ-CL-0196 | No bufferbloat grade without a latency sample under load; only the directions really sampled count. | CHANGELOG.md@ad3758c:184 | M-04 | M-06, M-07 | fault |
| REQ-CL-0197 | Cloudflare: the server place is "Cloudflare <colo>" (its city and country describe the client; that stays in raw data); one latency sample gives no jitter; no sample under load gives no loaded latency; no generic link as the result link. | CHANGELOG.md@ad3758c:185 | M-06 |  | fault |
| REQ-CL-0198 | iPerf3: never store block size, MSS or "Auto" as TCP window (the largest congestion window when reported); a missing retransmit count is empty; TCP RTT under upload load is loaded latency, not ping; the port is never stored as colo. | CHANGELOG.md@ad3758c:186 | M-09 |  | fault |
| REQ-CL-0199 | Only a public address is stored as the test's client address: never the LAN address, never a cached address of the other IP family the test did not use, never today's address next to old tests in the report. | CHANGELOG.md@ad3758c:187 | M-04 | M-12, R-03 | fault |
| REQ-CL-0200 | The detailed CSV and dossier judge no status against a speed that was never set (no filled-in defaults); "below the guaranteed speed" checks upload too; Wi-Fi tests are recognised from their interface; a missing value is "--". | CHANGELOG.md@ad3758c:188 | R-02 | R-01 | fault |
| REQ-CL-0201 | The detailed CSV has time zone (UTC), status text, engine and interface as last columns; its file name is neutral (`nmb_speedtests_detailed_<date>.csv`) in every language. | CHANGELOG.md@ad3758c:188 | R-02 | C-06 | fault |
| REQ-CL-0202 | The speed CSV writes no made-up ISP or engine name for an unknown value, and exports a measured 0 as 0. | CHANGELOG.md@ad3758c:188 | R-02 |  | fault |
| REQ-CL-0203 | The ISP settings never pre-fill guaranteed speeds that a save would store as if the owner set them. | CHANGELOG.md@ad3758c:189 | U-02…U-0n (page: system) | C-01 | fault |
| REQ-CL-0204 | Without a guaranteed speed the bot calls a result "not judged" (11 languages); it shows only shares of the speeds that are set, "--" for a missing ping or jitter, and never another address in place of the test's own. | CHANGELOG.md@ad3758c:190 | S-11 | R-01, C-06 | fault |
| REQ-CL-0205 | In every mode the box has a host firewall: from the internet only answers to its own traffic, ICMP, DHCP, Tailscale and the VPN servers that are on reach the box; private addresses, the home network's own IPv6 prefix (followed when it changes), Tailscale, WireGuard and OpenVPN reach everything. | CHANGELOG.md@ad3758c:191 | N-02 | N-03, N-04, N-13, N-14, N-15 | fault; security |
| REQ-CL-0206 | Router mode: the WAN guard is up before PPPoE dials. | CHANGELOG.md@ad3758c:191 | N-02 | N-05 | fault; security |
| REQ-CL-0207 | IP forwarding is switched on only after the firewall rules took, and stays off if they fail (shown in the status). | CHANGELOG.md@ad3758c:191 | N-02 | N-03, S-02…S-0n (API: status) | fault; security |
| REQ-CL-0208 | A second WAN gets the same NAT and protection as the main WAN, IPv6 included. | CHANGELOG.md@ad3758c:191 | N-02 |  | fault; security |
| REQ-CL-0209 | Any fallback firewall backend covers IPv6 too, or turns IPv6 forwarding off. | CHANGELOG.md@ad3758c:191 | N-02 |  | fault; security |
| REQ-CL-0210 | Switching modes stops PPPoE before removing the firewall rules. | CHANGELOG.md@ad3758c:191 | N-03 | N-05, N-02 | fault |
| REQ-CL-0211 | Only NMB's own firewall rules are removed (never other software's NAT), and they never pile up. | CHANGELOG.md@ad3758c:191 | N-02 |  | fault |
| REQ-CL-0212 | Proof: from a LAN address SSH and DNS are reachable, from a public one not; the VPN port and ping are. | CHANGELOG.md@ad3758c:191 | N-02 | F-06 | Python proof: real nft in namespaces |
| REQ-CL-0213 | The bot's backup send sends only an encrypted archive (a standard tool can open it) with a backup passphrase of at least 12 characters set on the Web UI's Telegram card (never shown again); without a passphrase it sends nothing and says how to set one; an explicit plain option sends unencrypted with a warning; 11 languages. | CHANGELOG.md@ad3758c:192 | S-11 | P-03, C-04 | fault; security |
| REQ-CL-0214 | The hint under the authorised chat IDs says how groups work. | CHANGELOG.md@ad3758c:192 | U-02…U-0n (page: system) | S-11 | fault |
| REQ-CL-0215 | DNS blocker: each upstream query has a random ID and its own connected socket, and only a reply to that exact question is used (the client's ID put back). | CHANGELOG.md@ad3758c:193 | N-12 |  | fault; security |
| REQ-CL-0216 | DNS blocker refuses names with bad bytes, pointers in the question, over-long names or several questions. | CHANGELOG.md@ad3758c:193 | N-12 | F-06 | fault; security |
| REQ-CL-0217 | DNS blocker cache keys on class, flags and the DNSSEC OK bit and keeps only complete NOERROR/NXDOMAIN answers. | CHANGELOG.md@ad3758c:193 | N-12 |  | fault; security |
| REQ-CL-0218 | One socket error never stops DNS for the whole network; queries are answered by a small pool so one slow upstream answer does not hold up the others. | CHANGELOG.md@ad3758c:193 | N-12 |  | fault |
| REQ-CL-0219 | The bot runs a command only when the sender's own user id is listed (in a private chat, the listed user's own chat), not just the chat. | CHANGELOG.md@ad3758c:194 | S-11 |  | fault; security |
| REQ-CL-0220 | Button presses run only the commands the bot's own buttons carry; edited messages are not run again; commands queued while the box was off (older than 2 minutes) are dropped. | CHANGELOG.md@ad3758c:194 | S-11 |  | fault; security |
| REQ-CL-0221 | A stranger is told "Access Denied" at most every 10 minutes and members of a listed group not at all; a reload uses the new bot token and alert recipients. | CHANGELOG.md@ad3758c:194 | S-11 |  | fault |
| REQ-CL-0222 | The home page names the saved engine for automatic tests and, for the mixed run, the engines that take part. | CHANGELOG.md@ad3758c:195 | U-02…U-0n (page: home) | M-04 | fault |
| REQ-CL-0223 | An unsent ping (no route) counts as lost. | CHANGELOG.md@ad3758c:196 | M-02 | M-01 | fault |
| REQ-CL-0224 | The internet targets decide an outage (a local gateway answering never hides one; recorded as "internet targets unreachable while the local gateway answers"); in reports only an answer from the internet ends an outage. | CHANGELOG.md@ad3758c:196 | M-02 | M-01, R-01 | fault |
| REQ-CL-0225 | The Web UI shows the installed NMB version (the commit), no client when there is none, and the packet counts the sniffer really saw; never invented firmware numbers, clients or rates. | CHANGELOG.md@ad3758c:197 | U-01 | C-05, M-11, U-02…U-0n (page: tools) | fault (made-up data shown as the box's own) |
| REQ-CL-0226 | Every visible Web UI text (page and script) has a translation key and is translated into all 11 languages; the language menu keeps each language's own name; no Hungarian (or other) fallback; tests fail on visible text without a key and on Hungarian left in the page. | CHANGELOG.md@ad3758c:198 | C-06 | U-01 | fault |
| REQ-CL-0227 | UI logic never reads translated text. | CHANGELOG.md@ad3758c:198 | U-01 | C-06 | fault |
| REQ-CL-0228 | The "purge all logs" action works. | CHANGELOG.md@ad3758c:198 | U-02…U-0n (page: system) | C-08 | fault |
| REQ-CL-0229 | Making a report only reads: it never changes the Wi-Fi regulatory domain; the access point sets it when it starts. | CHANGELOG.md@ad3758c:199 | R-04 | N-07 | fault |
| REQ-CL-0230 | Tests never ask a public GeoIP service and never set the regulatory domain. | CHANGELOG.md@ad3758c:199 | F-06 | N-07, M-12 | fault |
| REQ-CL-0231 | A test checks that every file the installer installs or the README names is in the repository. | CHANGELOG.md@ad3758c:200 | P-11 | P-13, F-00 | fault |
| REQ-CL-0232 | At boot in router mode the box counts as online with one active connection (the LAN bridge in router mode, the uplink in standalone mode), so dependent services start without a timeout; a box with no network at all still fails it. (source names a specific site value; not copied) | CHANGELOG.md@ad3758c:201 | P-10 | N-03, N-13, N-14 | fault |
| REQ-CL-0233 | Tests never read the box's real state (network mode, line idle state, a running service): they use fixed inputs. | CHANGELOG.md@ad3758c:202 | F-06 | P-02 | fault |
| REQ-CL-0234 | Unit tests do not depend on the machine they run on: no system tools, no real commands, no internet, no order between tests, no real database, no time-of-boot or wall-clock boundary; they run on fixed inputs and throwaway files. | CHANGELOG.md@ad3758c:203 | F-06 |  | fault |
| REQ-CL-0235 | Address flushes are limited to DHCP addresses. | CHANGELOG.md@ad3758c:203 | N-07 | N-08 | fault; the second rule of the line (cable on another port preferred) is REQ-CL-0452 (split in round 1). |
| REQ-CL-0452 | With the measuring link on automatic, a cable on another port than the configured WAN is still preferred as the measuring interface. | CHANGELOG.md@ad3758c:203 | N-01 | N-08, M-04 | fault; split from REQ-CL-0235 in round 1. Python: `find_preferred_measurement_interface` in `nmb/netd/interfaces.py@ad3758c:465` (N-01 in python-files.md), checked by `tests/test_wifi_uplink.py`. |
| REQ-CL-0236 | The update's test run never runs as root on the live box and never touches real settings, database, locks, DynDNS state, blocker lists, crash reports, bundles, backups or kept logs: everything points at throwaway files; a test that reaches a real file fails and the update is refused. | CHANGELOG.md@ad3758c:204 | P-02 | F-06, C-01 | fault; security; data loss |
| REQ-CL-0237 | Proof: after the test run the settings, secrets copy, boot copy and blocker files are byte for byte unchanged. | CHANGELOG.md@ad3758c:204 | P-02 |  | Python proof: checked as root on the box |
| REQ-CL-0238 | The diagnostic bundle never packs secrets: tokens and passwords become `<redacted, N chars>`. | CHANGELOG.md@ad3758c:205 | R-04 | C-04 | fault; security |
| REQ-CL-0239 | A log file the bundle may not read does not make the whole bundle fail. | CHANGELOG.md@ad3758c:205 | R-04 |  | fault |
| REQ-CL-0240 | The update's test run uses its own measurement lock and DynDNS state, so a scheduled test on the box cannot make it fail. | CHANGELOG.md@ad3758c:206 | P-02 | M-04, N-16 | fault |
| REQ-CL-0241 | The bot's heal apply never claims an update is installed when it only started: it says that it runs, which version, and where to follow it (11 languages). | CHANGELOG.md@ad3758c:207 | S-11 | P-02, C-06 | fault |
| REQ-CL-0242 | Counting a cycle's data usage cannot make the cycle fail (every counter is written inside the guard). | CHANGELOG.md@ad3758c:208 | M-05 | M-04 | fault |
| REQ-CL-0243 | The wired route priority is set by the privileged network daemon at each mode change and every 5 minutes outside router mode, not by the unprivileged web server. | CHANGELOG.md@ad3758c:209 | N-01 | N-03 | fault |
| REQ-CL-0244 | When the system rewrites the resolver file while the DNS fallback is on, the fallback marker is dropped, so a later restore never puts back an old network's file. | CHANGELOG.md@ad3758c:210 | M-12 | N-03 | Python reference nmb/netd/dns.py (box resolver fallback); see proposed split N-20; fault |
| REQ-CL-0245 | The installer makes room on a nearly full log disk (rotated NMB logs, oldest first) before writing anything. | CHANGELOG.md@ad3758c:211 | P-11 | C-03 | fault |
| REQ-CL-0246 | Log sizes fit the log disk (smaller on boards with a small log RAM disk). | CHANGELOG.md@ad3758c:211 | P-11 | C-03 | fault |
| REQ-CL-0247 | The installer never prints or stores a Web UI password the owner set, and warns on every run while the factory login is still set. | CHANGELOG.md@ad3758c:211 | P-11 | P-03 | fault; security |
| REQ-CL-0248 | The developer log default is 5 MB with 2 copies. | CHANGELOG.md@ad3758c:212 | C-03 |  | fault |
| REQ-CL-0249 | The documentation has the full bot command list in every language, and describes the storage guard, the update check and the 5 GHz result as they work. | CHANGELOG.md@ad3758c:213 | P-12 | S-11, C-08, P-02 |  |
| REQ-CL-0250 | The installed version is recorded as a full commit id, so the update check knows it; deploy instructions run the install detached with its log on persistent storage. | CHANGELOG.md@ad3758c:213 | P-02 | P-12 | fault |
| REQ-CL-0251 | The bot does not wait for the network-online target; it retries Telegram by itself. | CHANGELOG.md@ad3758c:214 | S-11 | P-10 | fault |
| REQ-CL-0252 | The storage guard also keeps 20 % of a separate log disk free (rotated logs first, oldest first, then active NMB logs over 1 MB emptied). | CHANGELOG.md@ad3758c:215 | C-08 | C-03 | fault |
| REQ-CL-0253 | Purging logs also clears the warnings kept on persistent storage and counts them in the summary. | CHANGELOG.md@ad3758c:216 | C-08 | S-11, U-02…U-0n (page: system) | fault |
| REQ-CL-0254 | The bot's kill command refuses a signal it does not allow instead of sending another. | CHANGELOG.md@ad3758c:217 | S-11 | C-09 | fault |
| REQ-CL-0255 | Log rotation shrinks to fit 40 % of the file system the log is on. | CHANGELOG.md@ad3758c:218 | C-03 | C-08 | fault |
| REQ-CL-0256 | DynDNS: the updater follows the settings when another process saves them (within 30 s) and shares its state through a runtime file; only one updater runs; the bot only saves the setting and reads the shared state. | CHANGELOG.md@ad3758c:219 | N-16 | S-11 | fault |
| REQ-CL-0257 | A manual DynDNS update uses the configured provider. | CHANGELOG.md@ad3758c:219 | N-16 | S-11 | fault |
| REQ-CL-0258 | Purging measurements works in time slices with a checkpoint after each, so the write-ahead log never needs as much space as the database. | CHANGELOG.md@ad3758c:220 | C-08 | C-02, S-11 | fault |
| REQ-CL-0259 | The update check offers an update only when the published version is newer than the installed one; applying without an update available changes nothing; the bot shows the commits. | CHANGELOG.md@ad3758c:221 | P-02 | S-11 | fault |
| REQ-CL-0260 | The bot's Wi-Fi connect in router mode says the radio is the access point (switch to standalone or turn the access point off) instead of trying; 11 languages. | CHANGELOG.md@ad3758c:222 | S-11 | N-07, C-06 | fault |
| REQ-CL-0261 | Changing the mode setting keeps the running mode when the new setting leads to it (no rebuild, no redial). Python proof: same pppd process, no ping lost. | CHANGELOG.md@ad3758c:223 | N-03 | N-04 | fault |
| REQ-CL-0262 | A resolver file without a nameserver line (the C library then asks 127.0.0.1) is tested like any other, not taken as broken. | CHANGELOG.md@ad3758c:224 | M-12 | N-03 | Python reference nmb/netd/dns.py (box resolver fallback); see proposed split N-20; fault |
| REQ-CL-0263 | Outages a stopped probe left open are ended once after the probe starts, where the data shows; the newest carries on if the line is still down, so a restart never opens a second row for the same outage. | CHANGELOG.md@ad3758c:225 | M-02 | C-02 | fault |
| REQ-CL-0264 | The Web UI header shows the mode the network daemon runs (Router, Bridge or Standalone, "PPPoE" when the router dials, "auto" in automatic mode), in 11 languages. | CHANGELOG.md@ad3758c:226 | U-01 | N-03, S-02…S-0n (API: status), C-06 | fault |
| REQ-CL-0265 | Automatic mode dials PPPoE once at start and settings reload: the decision is applied before the first setup, and router mode is rebuilt only when the dialling really changes. | CHANGELOG.md@ad3758c:227 | N-04 | N-05 | fault |
| REQ-CL-0266 | Changing only the access point's settings in router mode restarts only the access point; WAN, DHCP, IPv6 and firewall stay up. | CHANGELOG.md@ad3758c:228 | N-07 | N-03 | fault |
| REQ-CL-0267 | Availability and outage totals end an outage left open by a stopped probe at the next one or at the first answered ping, and count overlapping outages once; stored rows are not changed. | CHANGELOG.md@ad3758c:229 | R-01 | M-02, S-11 | fault |
| REQ-CL-0268 | The bot's outage alert does not stay "firing" on such a row, so a real outage alerts again. | CHANGELOG.md@ad3758c:229 | S-11 | M-02 | fault |
| REQ-CL-0269 | The "server left out" warning comes once per rest period (a check without the test history does not clear the mark). | CHANGELOG.md@ad3758c:230 | M-09 | C-03 | fault |
| REQ-CL-0270 | The ping probe follows the line (the router's WAN in router mode, never a link without an address) and rebinds within 30 s. | CHANGELOG.md@ad3758c:231 | M-01 | N-01, N-05 | fault |
| REQ-CL-0271 | The Local Gateway target is switched off when the box has no gateway (PPPoE); switched-off targets do not show in the table. | CHANGELOG.md@ad3758c:231 | M-01 |  | fault |
| REQ-CL-0272 | While router mode runs, the WAN/PPPoE port is released from the system's network manager (no DHCP on it) and given back afterwards. | CHANGELOG.md@ad3758c:232 | N-03 | N-05 | fault |
| REQ-CL-0273 | The wired route priority is set in the connection's profile, and an old separate route is removed. | CHANGELOG.md@ad3758c:233 | N-01 |  | fault |
| REQ-CL-0274 | DynDNS asks the IPv4 address over IPv4 only and sends it explicitly (every provider). | CHANGELOG.md@ad3758c:234 | N-16 |  | fault |
| REQ-CL-0275 | The self-test's simulated LAN client keeps its IPv6 link-local address, so the IPv6 check proves something. | CHANGELOG.md@ad3758c:234 | P-07 | N-11 | fault |
| REQ-CL-0276 | A settings save re-applies the kernel tuning only when the tuning file or the set of interfaces changed. | CHANGELOG.md@ad3758c:235 | N-01 | N-03 | fault |
| REQ-CL-0277 | The DHCP/DNS server of router mode writes no pid file that needs extra capabilities. | CHANGELOG.md@ad3758c:235 | N-03 |  | fault |
| REQ-CL-0278 | The PPPoE options file is written new and moved over; the installer never changes the owner of runtime files. | CHANGELOG.md@ad3758c:236 | N-05 | P-11 | fault |
| REQ-CL-0279 | A refused PPPoE login is not resent for ever: give up after 3 failures and dial again after 10, 20, then 30 minutes (30 s for a line problem); the reason shows in the status; the real-line self-test dials once and shows the ISP's message. | CHANGELOG.md@ad3758c:237 | N-05 | S-02…S-0n (API: status) | fault |
| REQ-CL-0280 | The PPP link accepts the ISP's router advert from the moment it is created, also with forwarding on. | CHANGELOG.md@ad3758c:238 | N-05 | N-11 | fault |
| REQ-CL-0281 | A USB hotspot or modem carries a test only when it is the chosen measuring link, the router's WAN, or the only internet in standalone or bridge mode. | CHANGELOG.md@ad3758c:239 | M-04 | N-10 | fault |
| REQ-CL-0282 | In router mode tests measure the router's WAN; a link-local address never counts as a test link. | CHANGELOG.md@ad3758c:240 | M-04 | N-01, N-05 | fault |
| REQ-CL-0283 | An engine that raises an error is stored as a failed test of that engine and the other engines run. | CHANGELOG.md@ad3758c:241 | M-04 |  | fault |
| REQ-CL-0284 | RTR control requests never retry over an unbound connection. | CHANGELOG.md@ad3758c:241 | M-08 |  | fault |
| REQ-CL-0285 | The access point's HT and VHT flags come from what the radio reports for each band, with the right HT40 direction; 160 MHz is never claimed from "neither 160 nor 80+80". | CHANGELOG.md@ad3758c:242 | N-07 |  | fault |
| REQ-CL-0286 | USB tethering takes over when every other route has had no internet for about 90 s, also when the Wi-Fi network has a default route, until another route works again. | CHANGELOG.md@ad3758c:243 | N-10 | N-01 | fault |
| REQ-CL-0287 | The box checks its route every 30 s and updates DynDNS at once when it changes. | CHANGELOG.md@ad3758c:244 | N-16 | N-01 | fault |
| REQ-CL-0288 | UPnP starts once the WAN has an address and restarts when the address changes. | CHANGELOG.md@ad3758c:245 | N-11 |  | fault |
| REQ-CL-0289 | The real-line PPPoE self-test link uses no peer DNS (the box's resolver is untouched). | CHANGELOG.md@ad3758c:246 | P-07 | N-05 | fault |
| REQ-CL-0290 | The self-test's groups read the real WAN name (the PPP link in router mode) from the state; a PPPoE line without an answer to a solicitation is not called "no IPv6"; the Wi-Fi group reports width and standard. | CHANGELOG.md@ad3758c:247 | P-07 |  | fault |
| REQ-CL-0291 | PPPoE uses the kernel's PPPoE support (with the saved service name), never a user-space one that caps the line. | CHANGELOG.md@ad3758c:248 | N-05 | N-02 | fault |
| REQ-CL-0292 | Router mode clamps the TCP MSS to the path MTU. | CHANGELOG.md@ad3758c:248 | N-02 | N-05 | fault |
| REQ-CL-0293 | Automatic mode: a PPPoE server answering on the WAN port with a saved login switches to router mode with PPPoE at once; only internet through the WAN port counts; PPPoE discovery runs at most every 30 s. | CHANGELOG.md@ad3758c:249 | N-04 | N-05 | fault |
| REQ-CL-0294 | Self-test and scenario reports are rotated (the newest 30 of each kind stay). | CHANGELOG.md@ad3758c:250 | P-07 | C-08 | fault |
| REQ-CL-0295 | The snapshot tool never says it mirrored to the boot partition when the hourly limit skipped it. | CHANGELOG.md@ad3758c:250 | R-04 |  | fault |
| REQ-CL-0296 | On an HTTP error or without network the bot waits (the server's retry-after, or a growing pause); a message refused with "retry after" is sent once more, not lost. | CHANGELOG.md@ad3758c:251 | S-11 |  | fault |
| REQ-CL-0297 | The bot's sandbox lets it write exactly what its commands need (its state, the settings, the boot copy). | CHANGELOG.md@ad3758c:252 | P-04 | S-11, N-16, P-01, M-10 | fault |
| REQ-CL-0298 | Dead default routes are removed by their selector and success is checked; the log never says "removed" when it failed. | CHANGELOG.md@ad3758c:253 | N-01 |  | fault |
| REQ-CL-0299 | The DNS fallback acts only when IP works but DNS does not (three tries), keeps a fresh copy of the old settings, and puts the normal DNS back after 15 minutes if it works. | CHANGELOG.md@ad3758c:254 | M-12 | N-03 | Python reference nmb/netd/dns.py (box resolver fallback); see proposed split N-20; fault; same behaviour as REQ-TS-0248 (one owner, M-12) |
| REQ-CL-0300 | IPv4 DNS servers come first on the wired and Wi-Fi profiles, so IPv6-only servers without an IPv6 route never fill the slots. | CHANGELOG.md@ad3758c:255 | N-01 | M-12, N-07, N-08 | Python reference nmb/netd/interfaces.py; fault |
| REQ-CL-0301 | The access point's channel width follows what the radio does on that band; 802.11ac is never asked on 2.4 GHz. | CHANGELOG.md@ad3758c:256 | N-07 |  | fault |
| REQ-CL-0302 | The RTR client keeps its client ID in the box's state directory, not in a home folder. | CHANGELOG.md@ad3758c:257 | M-08 | P-04 | fault |
| REQ-CL-0303 | Root-only files (the secrets copy) stay root-only after an install. | CHANGELOG.md@ad3758c:258 | P-11 | C-04 | fault; security |
| REQ-CL-0304 | A rested iPerf3 server is logged once per rest period; scheduler errors carry a stack trace. | CHANGELOG.md@ad3758c:259 | C-03 | M-09, M-04 | fault |
| REQ-CL-0305 | DynDNS card: the badge says "Unsaved" and the header wraps; after saving, the live status is fetched again once the update has run; the card names the provider in use (11 languages). | CHANGELOG.md@ad3758c:260 | U-02…U-0n (page: network) | N-16, C-06 | fault |
| REQ-CL-0306 | Self-test `dyndns` checks the name in public DNS and warns about an AAAA record on a line without IPv6. | CHANGELOG.md@ad3758c:261 | P-07 | N-16 |  |
| REQ-CL-0307 | DynDNS on with a domain but no token is logged and the status shows "the token is missing". | CHANGELOG.md@ad3758c:262 | N-16 | S-02…S-0n (API: status) | fault |
| REQ-CL-0308 | Self-test `dyndns` checks the token, a recent successful update and that the name points at the current line. | CHANGELOG.md@ad3758c:262 | P-07 | N-16 |  |
| REQ-CL-0309 | A bot speed test runs in its own thread, so the bot answers meanwhile; a second one is refused politely; the result shows only this test's rows (one line per engine); no promised duration. | CHANGELOG.md@ad3758c:263 | S-11 | M-04 | fault |
| REQ-CL-0310 | One speed test at a time on the whole box: a scheduled test skips, a manual, web or Telegram test waits up to 10 minutes. | CHANGELOG.md@ad3758c:264 | M-04 | S-11, U-02…U-0n (page: speed) | fault |
| REQ-CL-0311 | PPPoE credentials are never on a command line; they are in a private options file. | CHANGELOG.md@ad3758c:265 | N-05 | C-04 | fault; security |
| REQ-CL-0312 | Starting the access point removes only a DHCP client's dynamic addresses and routes, never its own address. | CHANGELOG.md@ad3758c:266 | N-07 | N-08 | fault |
| REQ-CL-0313 | Self-test `wifi-ap` fails when an access point has no address or no DHCP. | CHANGELOG.md@ad3758c:266 | P-07 | N-07 | fault |
| REQ-CL-0314 | RTR native client: binds to the asked link, asks the client IP over the same link, and reports the real ping minimum and jitter (an older binary leaves them empty). | CHANGELOG.md@ad3758c:267 | M-08 | M-12 | fault |
| REQ-CL-0315 | Scheduled Ookla runs during a rate-limit pause are stored as skipped, not failed. | CHANGELOG.md@ad3758c:268 | M-05 | M-10 | fault |
| REQ-CL-0316 | An iPerf3 server whose last five tests all failed is left out for 6 hours. | CHANGELOG.md@ad3758c:269 | M-09 |  | fault |
| REQ-CL-0317 | No shared factory Wi-Fi password exists in defaults or example config; the installer first restores the owner's password from the root-only copy and generates a new one only when there is none. | CHANGELOG.md@ad3758c:270 | P-03 | N-07, C-01, P-11 | fault; security |
| REQ-CL-0318 | After setting up, router mode drops the wired LAN ports' link for 4 seconds so devices ask for addresses again. | CHANGELOG.md@ad3758c:271 | N-03 | N-05 | fault |
| REQ-CL-0319 | The DHCP guard stops serving within about a minute after another router appears. | CHANGELOG.md@ad3758c:271 | N-05 | N-10 | Python proof: live on USB tethering |
| REQ-CL-0320 | The scenario packet counter uses a valid firewall rule; the "laptop got an address" check never counts the simulated client's lease. | CHANGELOG.md@ad3758c:272 | P-07 |  | fault |
| REQ-CL-0321 | Default routes through a port with no link or no address are removed, and the kernel ignores routes on a port without link; the self-test fails on such a route. | CHANGELOG.md@ad3758c:273 | N-01 | N-08, N-10 | fault |
| REQ-CL-0322 | Secrets the owner typed are kept in a root-only copy for 30 days; the self-test reports one that went missing and a root command puts it back; the USB scenario restores them itself. | CHANGELOG.md@ad3758c:274 | C-04 | P-07, P-11 |  |
| REQ-CL-0323 | The editable boot copy of the settings is taken over only when it was really edited, and an empty secret there never replaces a set one. | CHANGELOG.md@ad3758c:275 | C-01 | P-02 | fault |
| REQ-CL-0324 | The network daemon logs a loop part that takes longer than 15 s, and a signal dumps its stacks to the journal. | CHANGELOG.md@ad3758c:276 | C-03 | N-03 | fault |
| REQ-CL-0325 | The self-test fails on a weak or factory Wi-Fi password, even with the access point off. | CHANGELOG.md@ad3758c:277 | P-07 | N-07 | fault |
| REQ-CL-0326 | The speed command's interface option is used: a test asked for a link runs on that link. | CHANGELOG.md@ad3758c:278 | M-04 | S-10 | fault |
| REQ-CL-0327 | The collector reads the upload before refusing a wrong token, so the box sees "the collector refused the token". | CHANGELOG.md@ad3758c:279 | S-13 | S-12 | fault |
| REQ-CL-0328 | Every router settings text is translated into all 11 languages. | CHANGELOG.md@ad3758c:280 | C-06 | U-02…U-0n (page: router) | fault |
| REQ-CL-0329 | Router mode moves its WAN to USB tethering when the WAN port keeps only an IPv6 link-local address, and takes the wired port back when a router answers there. | CHANGELOG.md@ad3758c:281 | N-10 | N-03 | fault |
| REQ-CL-0330 | The DHCP guard sends a complete frame from 0.0.0.0 and reads all protocols on the port, so it sees a router on a bridged port and on any port. | CHANGELOG.md@ad3758c:282 | N-05 | F-06 | fault |
| REQ-CL-0331 | In router mode on USB, the access point starts on a radio without client addresses, and every default route that does not go through USB is removed. | CHANGELOG.md@ad3758c:283 | N-10 | N-07 | fault |
| REQ-CL-0332 | In router mode the USB route stays; routes through bridge ports or on a down link never count as another internet; router mode removes default routes through its LAN ports. | CHANGELOG.md@ad3758c:284 | N-10 | N-01 | fault |
| REQ-CL-0333 | LAN ports are released from the system's network manager while router mode runs and given back afterwards; so is the radio when the access point stops. | CHANGELOG.md@ad3758c:285 | N-03 | N-07 | fault |
| REQ-CL-0334 | Generated configs replace an old file (written new and moved over), whoever owned it. | CHANGELOG.md@ad3758c:286 | N-07 | N-03 | fault |
| REQ-CL-0335 | A USB phone, hotspot or modem is never a LAN port, and tearing down router mode does not remove its address. | CHANGELOG.md@ad3758c:287 | N-10 | N-03, N-17 | fault |
| REQ-CL-0336 | USB phones, hotspots and modems are recognised by their driver (rndis_host, cdc_ncm, cdc_ether, ipheth, qmi_wwan, cdc_mbim) and never used as the wired line; USB Ethernet dongles stay wired. | CHANGELOG.md@ad3758c:288 | N-10 | N-17, M-04 | fault |
| REQ-CL-0337 | The access point does not start with the shared factory password, an empty one or one outside 8-63 characters; the Web UI refuses to save such a password; the API shows only whether a password is set; the installer never logs a generated password. (source names a specific site value; not copied) | CHANGELOG.md@ad3758c:289 | N-07 | S-02…S-0n (API: Wi-Fi), U-02…U-0n (page: Wi-Fi), P-11 | fault; security |
| REQ-CL-0338 | A settings save writes only the settings that service changed itself, never its whole in-memory copy. | CHANGELOG.md@ad3758c:290 | C-01 |  | fault |
| REQ-CL-0339 | The installer warns when Telegram is enabled without a bot token. | CHANGELOG.md@ad3758c:291 | P-11 | S-11 | fault |
| REQ-CL-0340 | Every bot command is tested against the functions it calls (heal, Wi-Fi, scan, connect, tasks, kill, Tailscale, backup, drivers), so a changed interface cannot break it silently. | CHANGELOG.md@ad3758c:292 | S-11 | F-06 | fault |
| REQ-CL-0341 | An install or upgrade restarts every service, so none keeps old code. | CHANGELOG.md@ad3758c:293 | P-11 | P-02, P-10, S-11 | fault; owner moved from P-02 in round 1: one owner (P-11) with REQ-TS-0168 and REQ-CT-0018 |
| REQ-CL-0342 | A bot error with no text names the error type. | CHANGELOG.md@ad3758c:294 | S-11 |  | fault |
| REQ-CL-0343 | The bot reads the SLA result from its code, never by matching words (a non-compliant result showed as compliant). | CHANGELOG.md@ad3758c:295 | S-11 | R-01 | fault |
| REQ-CL-0344 | Web UI texts describe the current behaviour in every language. | CHANGELOG.md@ad3758c:296 | C-06 | U-01 | fault |
| REQ-CL-0345 | The Ookla program has its own home in the box's state directory, so it keeps its settings; an accepted licence carries over. | CHANGELOG.md@ad3758c:297 | M-10 | P-04 | fault |
| REQ-CL-0346 | Settings saves run one at a time, each with its own temporary file, so two programs saving at once never leave a broken file. | CHANGELOG.md@ad3758c:298 | C-01 |  | fault |
| REQ-CL-0347 | The Wi-Fi status and the access-point check work for the unprivileged service (no dependence on a search path holding system tools). | CHANGELOG.md@ad3758c:299 | M-14 | N-07, P-10 | fault |
| REQ-CL-0348 | A test bound to Wi-Fi leaves over Wi-Fi also when the Wi-Fi was joined outside NMB and both links are on the same subnet. | CHANGELOG.md@ad3758c:300 | N-08 | N-01, M-04 | fault |
| REQ-CL-0349 | Cloudflare results never link to Cloudflare's front page as the result; a missing bufferbloat grade is "--". | CHANGELOG.md@ad3758c:301 | M-06 | R-01 | fault |
| REQ-CL-0350 | The speed test list labels every row with its own engine; an upload-only run shows no download. | CHANGELOG.md@ad3758c:302 | U-02…U-0n (page: speed) | M-09 | fault |
| REQ-CL-0351 | The local iPerf3 server is started and stopped through service control, so "Stop" works. | CHANGELOG.md@ad3758c:303 | M-09 | P-10, U-02…U-0n (page: speed) | fault |
| REQ-CL-0352 | A busy or unreachable iPerf3 server adds no failed rows to scheduled runs. | CHANGELOG.md@ad3758c:304 | M-09 | M-04 | fault |
| REQ-CL-0353 | The home card shows the last real result, says when a newer attempt failed, when the next Ookla test can run, and when the result is from an earlier chosen server. | CHANGELOG.md@ad3758c:305 | U-02…U-0n (page: home) | M-10 | fault |
| REQ-CL-0354 | The home schedule line shows a short note (full text on hover), wraps, and shows each label once. | CHANGELOG.md@ad3758c:306 | U-02…U-0n (page: home) | C-06 | fault |
| REQ-CL-0355 | Ookla failures are reported as Ookla failures; no other engine runs in its place; while Ookla is paused no rows are added. | CHANGELOG.md@ad3758c:307 | M-10 | U-02…U-0n (page: home) | fault |
| REQ-CL-0356 | The Ookla pause survives restarts (kept on persistent storage). | CHANGELOG.md@ad3758c:308 | M-05 | M-10 | fault |
| REQ-CL-0357 | The status API works when the history holds a failed speed test. | CHANGELOG.md@ad3758c:309 | S-02…S-0n (API: status) | M-04 | fault |
| REQ-CL-0358 | Reinstall or upgrade saves the measurements still in RAM first and replaces the program files, so removed files do not stay behind. | CHANGELOG.md@ad3758c:310 | P-11 | C-02 | fault |
| REQ-CL-0359 | A backup restore goes through SQLite's backup API, never replaying an old journal onto the new file. | CHANGELOG.md@ad3758c:311 | P-03 | C-02 | fault |
| REQ-CL-0360 | A restored config is not overwritten by the older boot copy and is readable by the services that need it; an uploaded backup file is not readable by every user. | CHANGELOG.md@ad3758c:312 | P-03 | C-01 | fault; security |
| REQ-CL-0361 | The helper's upgrade really upgrades packages, with the right package manager for the system. | CHANGELOG.md@ad3758c:313 | P-06 | P-02 | fault |
| REQ-CL-0362 | The privileged helper is called by its absolute path. | CHANGELOG.md@ad3758c:314 | P-06 | P-02 | fault |
| REQ-CL-0363 | A program given a config file path saves to that file. | CHANGELOG.md@ad3758c:315 | C-01 | S-01 | fault |
| REQ-CL-0364 | The installer stops with a clear message on a system without the expected service manager. | CHANGELOG.md@ad3758c:316 | P-11 |  | fault |
| REQ-CL-0365 | The privilege rule that lets the service user manage NMB's own services works on every supported system version, and on each it stays an allow-list of NMB's own units (never any unit). | CHANGELOG.md@ad3758c:317 | P-06 | P-11 | fault; restated in round 1 to match REQ-TS-0018 (the Python `.pkla` rule for polkit 0.105 allowed any unit, which is not carried over); one owner with REQ-RE-0339. |
| REQ-CL-0366 | OpenVPN server certificates carry TLS server usage; an old one is reissued from the same CA so issued profiles keep working. | CHANGELOG.md@ad3758c:318 | N-14 |  | fault |
| REQ-CL-0367 | The HTTPS proxy tries IPv6 and IPv4 addresses side by side; stopping it frees its port; it never forwards to the box's own loopback or link-local addresses. | CHANGELOG.md@ad3758c:319 | N-11 |  | fault |
| REQ-CL-0368 | Guest Wi-Fi: the service may control the network manager, reconnecting works with a stale profile, errors show the network manager's message, a disconnected client is not reported as connected. | CHANGELOG.md@ad3758c:320 | N-08 | P-06 | fault |
| REQ-CL-0369 | 5 GHz starts where the regulatory database exists (shipped with the box); the access point falls back to 2.4 GHz rather than none. | CHANGELOG.md@ad3758c:321 | N-07 | P-11 | fault |
| REQ-CL-0370 | Router mode's DHCP server answers a client's renewal at once after the box restarts (authoritative). | CHANGELOG.md@ad3758c:322 | N-03 | N-05 | fault |
| REQ-CL-0371 | Tests are independent of the machine's interfaces, config and network, so an update is not rolled back for that. | CHANGELOG.md@ad3758c:323 | F-06 | P-02 | fault |
| REQ-CL-0372 | Machine-readable output modes of the heal command print only JSON. | CHANGELOG.md@ad3758c:323 | P-01 | S-10 | fault |
| REQ-CL-0373 | The software reports its real version. | CHANGELOG.md@ad3758c:324 | C-05 |  | fault |
| REQ-CL-0374 | A fallback test is stored and logged under its own engine, never as an Ookla result under the chosen server's name. | CHANGELOG.md@ad3758c:325 | M-10 | M-04 | fault |
| REQ-CL-0375 | An Ookla rate limit pauses Ookla for the whole box (30 min, doubling up to 4 h) instead of quarantining servers one by one. | CHANGELOG.md@ad3758c:326 | M-05 | M-10 | fault |
| REQ-CL-0376 | The chosen Ookla server is never silently swapped; a result from a different server is flagged with `requested_server_id`. | CHANGELOG.md@ad3758c:326 | M-10 |  | fault |
| REQ-CL-0377 | Kernel network tuning runs in the privileged network daemon; the tuning button asks it to re-apply. | CHANGELOG.md@ad3758c:328 | N-01 | P-04, U-02…U-0n (page: system) | fault |
| REQ-CL-0378 | The installer's summary prints no raw terminal escape codes. | CHANGELOG.md@ad3758c:329 | P-11 |  | fault |
| REQ-CL-0379 | Wireless router mode: Wi-Fi and LAN share a bridge; a Wi-Fi failure does not stop wired DHCP; the access point steps down to an older standard if a setting is rejected. | CHANGELOG.md@ad3758c:331 | N-03 | N-07 | fault |
| REQ-CL-0380 | Restoring a lost main config from the boot copy never writes `__KEEP_CURRENT__` marks as values or brings back a shared factory Wi-Fi password; passwords and tokens must then be entered again. | CHANGELOG.md@ad3758c:332 | C-01 | C-04, N-07 | fault; security |
| REQ-CL-0381 | Running the installer again does not stop when a service is already restarting. | CHANGELOG.md@ad3758c:333 | P-11 |  | fault |
| REQ-CL-0382 | On every supported service manager the Web UI listens on port 80, the RTR client is built, and the summary shows the right status command. | CHANGELOG.md@ad3758c:334 | P-11 | S-01, M-08 | fault |
| REQ-CL-0383 | A new config uses the real network interface names, not a fixed name. | CHANGELOG.md@ad3758c:335 | N-17 | P-11, C-01 | fault |
| REQ-CL-0384 | PPPoE redials after the line drops. | CHANGELOG.md@ad3758c:336 | N-05 |  | fault |
| REQ-CL-0385 | The DHCP/DNS server of router mode does not conflict with the DNS blocker. | CHANGELOG.md@ad3758c:337 | N-12 | N-03 | fault |
| REQ-CL-0386 | The box connects to WPA3 access points (cipher set). | CHANGELOG.md@ad3758c:338 | N-07 |  | fault |
| REQ-CL-0387 | Network settings saved in the Web UI take effect without a restart. | CHANGELOG.md@ad3758c:339 | N-03 | S-02…S-0n (API: network) | fault |
| REQ-CL-0388 | The diagnostic snapshot copies nothing big into the system journal and is written to the boot drive at most once an hour. | CHANGELOG.md@ad3758c:340 | R-04 | C-03 | fault |
| REQ-CL-0389 | The privileged helper does only what NMB needs: runs from a fixed place with no code loaded from the caller's folder, repairs only NMB's own links and folders, restores settings only from the boot copy through the normal merge, installs engine programs only to their fixed places and only files root prepared itself, and installs only the packages NMB lists. | CHANGELOG.md@ad3758c:344 | P-06 | P-13 | fault; security |
| REQ-CL-0390 | No privileged write rule for the settings file exists; tests check that each allow-list matches the code and that the helper refuses everything else. | CHANGELOG.md@ad3758c:344 | P-06 | C-01 | fault; security |
| REQ-CL-0391 | Settings that leave the box are masked by one shared list of secrets and personal details (auth keys, PPPoE user name, user names, chat ids, domains, every line of a multi-line secret), used by the diagnostic text, the boot copy, crash reports and the bundle. | CHANGELOG.md@ad3758c:345 | C-04 | R-04, C-01, C-03 | fault; security |
| REQ-CL-0392 | Reading the masked boot copy back keeps all current values; a failed save never falls back to writing the full unmasked settings onto the boot partition. | CHANGELOG.md@ad3758c:345 | C-01 | C-04 | fault; security |
| REQ-CL-0393 | OpenVPN hands out a client profile only for a recorded client name; names of the server's own files cannot be used for clients. | CHANGELOG.md@ad3758c:346 | N-14 | P-06 | fault; security |
| REQ-CL-0394 | The OpenVPN server accepts only clients on a root-kept allow-list; deleting a client removes it from the list and ends its session; existing installs get the list on the next start without breaking an old config. | CHANGELOG.md@ad3758c:346 | N-14 | P-06 | fault; security |
| REQ-CL-0395 | Router mode with a public address: the WAN lets in only answers to the box's own traffic, ICMP, DHCP renewals, private sources and the VPN ports that are on; WAN-to-LAN only answers and forwarded ports; on IPv4 and IPv6. Python proof: checked from another network on the real line. | CHANGELOG.md@ad3758c:347 | N-02 | N-05, N-13, N-14, N-15 | fault; security |
| REQ-CL-0396 | No code path runs files from a world-writable directory (self-repair and speed test code ran files from `/tmp`). | CHANGELOG.md@ad3758c:348 | P-06 | M-04, P-01 | fault; security |
| REQ-CL-0397 | The service user cannot run network tools with root rights. | CHANGELOG.md@ad3758c:349 | P-06 |  | fault; security |
| REQ-CL-0398 | The privileged helper refuses VPN configs with command lines and package names that look like options. | CHANGELOG.md@ad3758c:350 | P-06 | N-13, N-14 | fault; security |
| REQ-CL-0399 | The language setting cannot be changed without logging in. | CHANGELOG.md@ad3758c:351 | S-01 | C-06 | fault; security |
| REQ-CL-0400 | The DNS blocker's control socket is limited to the box's own services. | CHANGELOG.md@ad3758c:352 | N-12 | P-04 | fault; security |
| REQ-CL-0401 | API values with control characters are refused, so no lines can be added to service config files. | CHANGELOG.md@ad3758c:353 | S-01 | C-01 | fault; security |
| REQ-CL-0402 | Update references that a tool could read as options are refused. | CHANGELOG.md@ad3758c:354 | P-02 |  | fault; security |
| REQ-CL-0403 | The Web UI, DNS and iPerf3 refuse public addresses on IPv4 and IPv6 through a firewall rule. | CHANGELOG.md@ad3758c:355 | N-02 | S-01, N-12, M-09 | fault; security |
| REQ-CL-0404 | Passwords and tokens are never copied to the boot partition, and a chosen admin password is never stored there. | CHANGELOG.md@ad3758c:356 | C-01 | C-04 | fault; security |
| REQ-CL-0405 | The web server runs as an unprivileged user, not root. | CHANGELOG.md@ad3758c:357 | P-04 | S-01, P-10 | security |
| REQ-CL-0406 | A downloaded engine program (Ookla) is installed only if its checksum matches. | CHANGELOG.md@ad3758c:358 | M-10 | P-02 | security |
| REQ-CL-0407 | No unused service that comes with an engine package (Tor with OONI) runs. | CHANGELOG.md@ad3758c:359 | M-13 | P-11 | security |
| REQ-CL-0408 | Network modes: standalone probe, gateway with NAT, DHCP and Wi-Fi access point, and transparent bridge, with multi-WAN failover and LAN port roles. | CHANGELOG.md@ad3758c:367 | N-03 | N-01, N-06 |  |
| REQ-CL-0409 | LAN passthrough and link aggregation for a probe behind a router. | CHANGELOG.md@ad3758c:368 | N-06 |  |  |
| REQ-CL-0410 | Continuous ICMP latency, jitter, loss and outage measurement. | CHANGELOG.md@ad3758c:369 | M-01 | M-02 |  |
| REQ-CL-0411 | Speed tests with Ookla, Cloudflare, RTR (native client) and iPerf3, plus port throttling checks. | CHANGELOG.md@ad3758c:370 | M-04 | M-06, M-08, M-09, M-10, M-11 |  |
| REQ-CL-0412 | Scheduled tests down to second-level intervals. | CHANGELOG.md@ad3758c:371 | M-04 | M-05 | server quarantine was replaced by the box-wide Ookla pause (REQ at line 326) |
| REQ-CL-0413 | A traffic gate postpones tests when the box sees other traffic. | CHANGELOG.md@ad3758c:372 | M-03 | M-04 |  |
| REQ-CL-0414 | Experimental, off-by-default stop on network traffic: tests only a provably idle network and stops running tests when other traffic appears, reading client-side ports or the router's UPnP counters. | CHANGELOG.md@ad3758c:373 | M-03 | M-04 |  |
| REQ-CL-0415 | Web dashboard with selectable themes (now 11 languages; AMOLED-only base with add-on themes per owner decision). | CHANGELOG.md@ad3758c:374 | U-01 | P-09, C-06 |  |
| REQ-CL-0416 | The dashboard has a device explorer, packet sniffer, traceroute, task manager and web terminal. | CHANGELOG.md@ad3758c:374 | U-02…U-0n (page: tools) | N-17, M-11, C-09 |  |
| REQ-CL-0417 | `nmb-cli` terminal interface with the Web UI's menus, mouse and hotkeys, JSON commands for scripts and LLMs, log listing and a logo (plain CLI per owner decision). | CHANGELOG.md@ad3758c:375 | S-10 |  |  |
| REQ-CL-0418 | Multilingual Telegram bot with menus and alerts. | CHANGELOG.md@ad3758c:376 | S-11 | C-06 |  |
| REQ-CL-0419 | HTML dossier, CSV export and LLM diagnostic bundle. | CHANGELOG.md@ad3758c:377 | R-03 | R-02, R-04 |  |
| REQ-CL-0420 | WireGuard, OpenVPN and Tailscale management, DynDNS (DuckDNS, dynv6) and an HTTP proxy. | CHANGELOG.md@ad3758c:378 | N-13 | N-14, N-15, N-16, N-11 |  |
| REQ-CL-0421 | Built-in DNS blocker with Pi-hole style lists and CLI compatibility. | CHANGELOG.md@ad3758c:379 | N-12 |  |  |
| REQ-CL-0422 | Optional weekly OONI Probe runs, off by default. | CHANGELOG.md@ad3758c:380 | M-13 |  |  |
| REQ-CL-0423 | Hardware driver detection and installation. | CHANGELOG.md@ad3758c:381 | N-17 | P-06 |  |
| REQ-CL-0424 | Backup, restore, factory reset and data purge. | CHANGELOG.md@ad3758c:382 | P-03 | C-08 |  |
| REQ-CL-0425 | Self-healing audit and repair with update rollback. | CHANGELOG.md@ad3758c:383 | P-01 | P-02 |  |
| REQ-CL-0426 | Low-memory profiles. | CHANGELOG.md@ad3758c:384 | C-08 | P-11 | Python: tested only on the Pi 4 |
| REQ-CL-0427 | New installs run speed tests every hour; upgrades keep the configured schedule. | CHANGELOG.md@ad3758c:389 | M-04 | C-01 |  |
| REQ-CL-0428 | The schedule setting is the only thing that starts speed tests (no separate timer). | CHANGELOG.md@ad3758c:390 | M-04 | P-10 |  |
| REQ-CL-0429 | The default log level is INFO. | CHANGELOG.md@ad3758c:392 | C-03 |  |  |
| REQ-CL-0430 | Cloudflare tests use fixed-time transfers. | CHANGELOG.md@ad3758c:393 | M-06 |  |  |
| REQ-CL-0431 | TLS streams that drop and reconnect during a Cloudflare test are logged as information, and their number is stored with the result. | CHANGELOG.md@ad3758c:394 | M-06 | C-03 |  |
| REQ-CL-0432 | An old `nmb.conf` with the removed settings (SQM selector, `sample_window_sec`, `cooldown_sec`) loads; the settings are not brought back. | CHANGELOG.md@ad3758c:399 | C-01 | F-04 |  |
| REQ-CL-0433 | A single invalid setting or a `%` sign in a value never stops the services. | CHANGELOG.md@ad3758c:403 | C-01 |  | fault |
| REQ-CL-0434 | A wrong web port or address never locks users out of the dashboard. | CHANGELOG.md@ad3758c:404 | S-01 | C-01 | fault |
| REQ-CL-0435 | Bridge mode requests an IPv4 address, measures on the bridge and reads forwarded traffic from the WAN port. | CHANGELOG.md@ad3758c:405 | N-03 | M-03 | fault |
| REQ-CL-0436 | Factory reset leaves a config the services can read, and updates the saved login notes. | CHANGELOG.md@ad3758c:406 | P-03 | C-01 | fault |
| REQ-CL-0437 | Installation leaves no failed units, and the main service target runs afterwards. | CHANGELOG.md@ad3758c:407 | P-11 | P-10 | fault |
| REQ-CL-0438 | Self-heal from Telegram and alerts calls the repair function correctly. | CHANGELOG.md@ad3758c:408 | P-01 | S-11 | fault |
| REQ-CL-0439 | Update verification has enough time, runs from the update's own code and does not touch the live database. | CHANGELOG.md@ad3758c:409 | P-02 | F-06 | fault |
| REQ-CL-0440 | Skipped speed tests never appear as 0 Mbps or replace real results. | CHANGELOG.md@ad3758c:410 | R-01 | U-02…U-0n (page: speed), M-04 | fault |
| REQ-CL-0441 | The local iPerf3 server check does not connect to its port (no failed test logged every few seconds). | CHANGELOG.md@ad3758c:411 | M-09 | C-03 | fault |
| REQ-CL-0442 | Impossible Cloudflare results are not stored. | CHANGELOG.md@ad3758c:412 | M-06 | R-01 | fault |
| REQ-CL-0443 | Invalid request bodies give a clean API error, not a failure. | CHANGELOG.md@ad3758c:412 | S-01 | S-02…S-0n (API: status) | fault |
| REQ-CL-0444 | Commands in the READMEs are correct (checked by the build). | CHANGELOG.md@ad3758c:412 | P-12 |  | fault |
| REQ-CL-0445 | Mobile layout works, the theme choice persists, and no Hungarian text appears in other languages. | CHANGELOG.md@ad3758c:413 | U-01 | C-06 | fault |
| REQ-CL-0446 | Factory reset and log purge require explicit confirmation in the CLI and the web API. | CHANGELOG.md@ad3758c:417 | P-03 | S-10, S-02…S-0n (API: system) | security |
| REQ-CL-0447 | Each box gets its own Wi-Fi password, also after a factory reset. | CHANGELOG.md@ad3758c:418 | P-03 | N-07 | security |
| REQ-CL-0448 | Narrow privilege rights, brute-force and CSRF protection. | CHANGELOG.md@ad3758c:419 | S-01 | P-06 | security |
| REQ-CL-0449 | External values are escaped in the dashboard. | CHANGELOG.md@ad3758c:419 | U-01 | S-01 | security |
| REQ-CL-0450 | No personal data in the repository. | CHANGELOG.md@ad3758c:420 | P-13 | F-00 | security; QB 8 |
| REQ-CL-0451 | Subscriber data is masked in the LLM bundle. | CHANGELOG.md@ad3758c:420 | C-04 | R-04 | security |

## Coverage

Every line 1-423 in order; no gaps. Blank lines are listed as such.

| Lines | Kind | Ids or context |
| :--- | :--- | :--- |
| 1 | heading | context: document title |
| 2 | blank | context: blank line |
| 3 | paragraph | context: statement of purpose of the file; the native changelog rule is F-00's (work plan 3.0) |
| 4 | blank | context: blank line |
| 5 | paragraph | context: format note (Keep a Changelog); the native changelog follows the same format, already required by F-00 |
| 6 | paragraph | context: versioning note (Semantic Versioning); covered by the work plan's F-00 and P-02 rows |
| 7 | blank | context: blank line |
| 8 | heading | context: version heading |
| 9 | blank | context: blank line |
| 10 | heading | context: section heading |
| 11 | blank | context: blank line |
| 12 | bullet | context: plan text (work plan 2.3A, development on a rented server); not box behaviour; registered from the work plan part |
| 13 | bullet | context: plan text (porting plan 7D, remote access without Tailscale, package N-18); registered from the porting plan part |
| 14 | bullet | REQ-CL-0001, REQ-CL-0002, REQ-CL-0003, REQ-CL-0004, REQ-CL-0005, REQ-CL-0006 |
| 15 | bullet | context: plan text (roadmap 6B.8, Go and C code bases learn from each other; package F-07); registered from the roadmap part |
| 16 | bullet | context: plan text (roadmap 4.29 and the fault register; package C-11); the behaviour is REQ at line 14, the plan in the roadmap part |
| 17 | bullet | context: plan text (the native work plan itself); registered from the work plan part |
| 18 | bullet | context: plan text (roadmap 6B.7, manual and command reference; package P-12); registered from the roadmap part |
| 19 | bullet | REQ-CL-0007, REQ-CL-0008, REQ-CL-0009, REQ-CL-0010, REQ-CL-0011 |
| 20 | bullet | context: plan text (roadmap 6B.6, automatic captive portal re-login; package N-09); registered from the roadmap part |
| 21 | bullet | context: plan text (roadmap 6B.6, captive portals on a Wi-Fi uplink; package N-09); registered from the roadmap part |
| 22 | bullet | REQ-CL-0012, REQ-CL-0013 |
| 23 | bullet | REQ-CL-0014, REQ-CL-0015, REQ-CL-0016, REQ-CL-0017, REQ-CL-0018, REQ-CL-0019 |
| 24 | bullet | context: plan text (porting plan 7H, parts from stock firmware); registered from the porting plan part |
| 25 | bullet | context: plan text (roadmap register of Python faults); registered by F-03 and the roadmap part |
| 26 | bullet | context: plan text (roadmap 6B.5, sandboxing; package P-04); registered from the roadmap part |
| 27 | bullet | REQ-CL-0020, REQ-CL-0021 |
| 28 | bullet | context: documentation (comparison with a national regulator's measuring box, linked from the READMEs); no box behaviour (source names a specific site value; not copied) |
| 29 | bullet | REQ-CL-0022, REQ-CL-0023 |
| 30 | bullet | context: development note (environment of a batch); recorded by the testing part, no box behaviour |
| 31 | bullet | REQ-CL-0024, REQ-CL-0025, REQ-CL-0026, REQ-CL-0027 |
| 32 | bullet | context: development setup document (WSL clock); not box behaviour |
| 33 | bullet | context: development setup document (native tool chain on the development PC); registered from the dev-setup part |
| 34 | bullet | context: development note (environment of a batch) (source names a specific site value; not copied) |
| 35 | bullet | REQ-CL-0028, REQ-CL-0029, REQ-CL-0030, REQ-CL-0031 |
| 36 | bullet | REQ-CL-0032 |
| 37 | bullet | REQ-CL-0033 |
| 38 | bullet | REQ-CL-0034 |
| 39 | bullet | REQ-CL-0035 |
| 40 | bullet | context: development note (environment of a batch) (source names a specific site value; not copied) |
| 41 | bullet | REQ-CL-0036, REQ-CL-0037, REQ-CL-0038 |
| 42 | bullet | REQ-CL-0039, REQ-CL-0040 |
| 43 | bullet | REQ-CL-0041, REQ-CL-0042, REQ-CL-0043 |
| 44 | bullet | REQ-CL-0044 |
| 45 | bullet | REQ-CL-0045 |
| 46 | bullet | REQ-CL-0046 |
| 47 | bullet | REQ-CL-0047, REQ-CL-0048, REQ-CL-0049 |
| 48 | bullet | REQ-CL-0050 |
| 49 | bullet | context: development setup document (Windows with WSL2); not box behaviour |
| 50 | bullet | REQ-CL-0051, REQ-CL-0052 |
| 51 | bullet | REQ-CL-0053 |
| 52 | bullet | REQ-CL-0054 |
| 53 | bullet | REQ-CL-0055 |
| 54 | bullet | REQ-CL-0056 |
| 55 | bullet | REQ-CL-0057 |
| 56 | bullet | context: plan text (roadmap 4.27, USB tethering in the native build; porting plan 7G); registered from the roadmap and porting plan parts |
| 57 | bullet | REQ-CL-0058, REQ-CL-0059, REQ-CL-0060 |
| 58 | bullet | context: plan text (roadmap: native collector server, same protocol and data layout; USB features to match); registered from the roadmap part |
| 59 | bullet | REQ-CL-0061 |
| 60 | bullet | context: plan text (porting plan 1.3, MikroTik and MediaTek targets); router ports are planned only; registered from the porting plan part |
| 61 | bullet | context: plan text (roadmap 6A, base candidates and criteria); registered from the roadmap part |
| 62 | bullet | context: plan text (porting plan, base distribution candidates); registered from the porting plan part |
| 63 | bullet | context: plan text (porting plan, own distribution and stock kernels); registered from the porting plan part |
| 64 | bullet | context: plan text (Generic First and the build matrix); registered from the roadmap and porting plan parts |
| 65 | bullet | context: testing record change (docs/TESTING.md); registered from the testing part |
| 66 | bullet | REQ-CL-0062 |
| 67 | bullet | REQ-CL-0063 |
| 68 | bullet | context: plan text (quality bar of the roadmap and porting plan); registered from the roadmap part |
| 69 | bullet | context: testing record change (release checklist); registered from the testing part |
| 70 | bullet | context: plan text (features of November 2026 and rules from fixes listed in the plans); registered from the roadmap part |
| 71 | bullet | REQ-CL-0064 |
| 72 | bullet | REQ-CL-0065 |
| 73 | bullet | context: plan text (porting plan: maintenance mode, logging, crash evidence, rollback; planned only); registered from the porting plan part |
| 74 | bullet | REQ-CL-0066, REQ-CL-0067 |
| 75 | bullet | context: contributor documents (CONTRIBUTING, pull request template); no box behaviour |
| 76 | bullet | context: testing record change (how to test, what is left to test); registered from the testing part |
| 77 | bullet | context: plan text (own distribution, planned only); registered from the roadmap part |
| 78 | bullet | context: plan text (a planned router target; router ports are planned only); registered from the porting plan part |
| 79 | bullet | REQ-CL-0068, REQ-CL-0069 |
| 80 | bullet | REQ-CL-0070 |
| 81 | bullet | context: documentation of the development line (source names a specific site value; not copied) |
| 82 | bullet | REQ-CL-0071, REQ-CL-0072, REQ-CL-0073, REQ-CL-0074, REQ-CL-0075 |
| 83 | bullet | REQ-CL-0076 |
| 84 | bullet | REQ-CL-0077, REQ-CL-0078 |
| 85 | bullet | REQ-CL-0079 |
| 86 | bullet | REQ-CL-0080 |
| 87 | bullet | REQ-CL-0081 |
| 88 | bullet | REQ-CL-0082 |
| 89 | bullet | REQ-CL-0083 |
| 90 | bullet | REQ-CL-0084 |
| 91 | bullet | REQ-CL-0085, REQ-CL-0086 |
| 92 | bullet | REQ-CL-0087 |
| 93 | bullet | REQ-CL-0088 |
| 94 | bullet | REQ-CL-0089 |
| 95 | bullet | REQ-CL-0090 |
| 96 | bullet | REQ-CL-0091 |
| 97 | bullet | REQ-CL-0092 |
| 98 | bullet | REQ-CL-0093 |
| 99 | bullet | context: testing record added; registered from the testing part |
| 100 | bullet | context: this changelog added; the native changelog is F-00's |
| 101 | bullet | context: plan for the native changelog; done in F-00 |
| 102 | bullet | REQ-CL-0094 |
| 103 | bullet | REQ-CL-0095 |
| 104 | bullet | REQ-CL-0096 |
| 105 | bullet | REQ-CL-0097 |
| 106 | bullet | REQ-CL-0098 |
| 107 | bullet | context: plan text (clean code rules); registered from the roadmap part |
| 108 | bullet | context: architecture document change (`nmb-cli` and the changelog rule); registered from the architecture part |
| 109 | blank | context: blank line |
| 110 | heading | context: section heading |
| 111 | blank | context: blank line |
| 112 | bullet | context: plan text (native build stays open, own repository, licence); registered from the work plan part |
| 113 | bullet | context: plan text (owner's decisions, work plan section 7); registered from the work plan part |
| 114 | bullet | context: plan text (native folder location); registered from the work plan part |
| 115 | bullet | REQ-CL-0099 |
| 116 | bullet | REQ-CL-0100 |
| 117 | bullet | REQ-CL-0101 |
| 118 | bullet | REQ-CL-0102 |
| 119 | bullet | REQ-CL-0103 |
| 120 | bullet | REQ-CL-0104 |
| 121 | bullet | REQ-CL-0105 |
| 122 | bullet | REQ-CL-0106 |
| 123 | bullet | REQ-CL-0107 |
| 124 | bullet | REQ-CL-0108, REQ-CL-0109, REQ-CL-0110, REQ-CL-0111 |
| 125 | bullet | REQ-CL-0112 |
| 126 | bullet | REQ-CL-0113 |
| 127 | bullet | REQ-CL-0114 |
| 128 | bullet | REQ-CL-0115 |
| 129 | bullet | REQ-CL-0116 |
| 130 | bullet | REQ-CL-0117 |
| 131 | bullet | context: README record of OSes tested in a simulated environment; see question on REQ for line 95 |
| 132 | bullet | REQ-CL-0118 |
| 133 | bullet | REQ-CL-0119 |
| 134 | blank | context: blank line |
| 135 | heading | context: section heading |
| 136 | blank | context: blank line |
| 137 | bullet | REQ-CL-0120, REQ-CL-0121 |
| 138 | bullet | REQ-CL-0122, REQ-CL-0123, REQ-CL-0124 |
| 139 | bullet | REQ-CL-0125 |
| 140 | bullet | REQ-CL-0126, REQ-CL-0127, REQ-CL-0128, REQ-CL-0129, REQ-CL-0130 |
| 141 | bullet | REQ-CL-0131 |
| 142 | bullet | REQ-CL-0132 |
| 143 | bullet | REQ-CL-0133, REQ-CL-0134, REQ-CL-0135, REQ-CL-0136 |
| 144 | bullet | REQ-CL-0137 |
| 145 | bullet | REQ-CL-0138, REQ-CL-0139 |
| 146 | bullet | REQ-CL-0140, REQ-CL-0141 |
| 147 | bullet | REQ-CL-0142, REQ-CL-0143 |
| 148 | bullet | REQ-CL-0144 |
| 149 | blank | context: blank line inside the Fixed list |
| 150 | bullet | REQ-CL-0145, REQ-CL-0146 |
| 151 | bullet | REQ-CL-0147, REQ-CL-0148, REQ-CL-0149 |
| 152 | bullet | REQ-CL-0150 |
| 153 | bullet | REQ-CL-0151, REQ-CL-0152 |
| 154 | bullet | context: lead-in of a nested list; the requirements are in its sub-bullets 155-162 |
| 155 | bullet (nested) | REQ-CL-0153 |
| 156 | bullet (nested) | REQ-CL-0154, REQ-CL-0155, REQ-CL-0156, REQ-CL-0157 |
| 157 | bullet (nested) | REQ-CL-0158, REQ-CL-0159 |
| 158 | bullet (nested) | REQ-CL-0160 |
| 159 | bullet (nested) | REQ-CL-0161 |
| 160 | bullet (nested) | REQ-CL-0162 |
| 161 | bullet (nested) | REQ-CL-0163 |
| 162 | bullet (nested) | REQ-CL-0164 |
| 163 | bullet | REQ-CL-0165 |
| 164 | bullet | REQ-CL-0166, REQ-CL-0167, REQ-CL-0168 |
| 165 | bullet | REQ-CL-0169 |
| 166 | bullet | REQ-CL-0170, REQ-CL-0171 |
| 167 | bullet | REQ-CL-0172 |
| 168 | bullet | REQ-CL-0173, REQ-CL-0174, REQ-CL-0175 |
| 169 | bullet | REQ-CL-0176 |
| 170 | bullet | REQ-CL-0177 |
| 171 | bullet | REQ-CL-0178, REQ-CL-0179 |
| 172 | bullet | REQ-CL-0180 |
| 173 | bullet (nested) | REQ-CL-0181 |
| 174 | bullet (nested) | REQ-CL-0182 |
| 175 | bullet (nested) | REQ-CL-0183 |
| 176 | bullet (nested) | REQ-CL-0184 |
| 177 | bullet (nested) | REQ-CL-0185 |
| 178 | bullet (nested) | REQ-CL-0186 |
| 179 | bullet | REQ-CL-0187 |
| 180 | bullet | REQ-CL-0188, REQ-CL-0189, REQ-CL-0190, REQ-CL-0191, REQ-CL-0192 |
| 181 | bullet | REQ-CL-0193 |
| 182 | bullet (nested) | REQ-CL-0194 |
| 183 | bullet (nested) | REQ-CL-0195 |
| 184 | bullet (nested) | REQ-CL-0196 |
| 185 | bullet (nested) | REQ-CL-0197 |
| 186 | bullet (nested) | REQ-CL-0198 |
| 187 | bullet (nested) | REQ-CL-0199 |
| 188 | bullet (nested) | REQ-CL-0200, REQ-CL-0201, REQ-CL-0202 |
| 189 | bullet (nested) | REQ-CL-0203 |
| 190 | bullet (nested) | REQ-CL-0204 |
| 191 | bullet | REQ-CL-0205, REQ-CL-0206, REQ-CL-0207, REQ-CL-0208, REQ-CL-0209, REQ-CL-0210, REQ-CL-0211, REQ-CL-0212 |
| 192 | bullet | REQ-CL-0213, REQ-CL-0214 |
| 193 | bullet | REQ-CL-0215, REQ-CL-0216, REQ-CL-0217, REQ-CL-0218 |
| 194 | bullet | REQ-CL-0219, REQ-CL-0220, REQ-CL-0221 |
| 195 | bullet | REQ-CL-0222 |
| 196 | bullet | REQ-CL-0223, REQ-CL-0224 |
| 197 | bullet | REQ-CL-0225 |
| 198 | bullet | REQ-CL-0226, REQ-CL-0227, REQ-CL-0228 |
| 199 | bullet | REQ-CL-0229, REQ-CL-0230 |
| 200 | bullet | REQ-CL-0231 |
| 201 | bullet | REQ-CL-0232 |
| 202 | bullet | REQ-CL-0233 |
| 203 | bullet | REQ-CL-0234, REQ-CL-0235, REQ-CL-0452 |
| 204 | bullet | REQ-CL-0236, REQ-CL-0237 |
| 205 | bullet | REQ-CL-0238, REQ-CL-0239 |
| 206 | bullet | REQ-CL-0240 |
| 207 | bullet | REQ-CL-0241 |
| 208 | bullet | REQ-CL-0242 |
| 209 | bullet | REQ-CL-0243 |
| 210 | bullet | REQ-CL-0244 |
| 211 | bullet | REQ-CL-0245, REQ-CL-0246, REQ-CL-0247 |
| 212 | bullet | REQ-CL-0248 |
| 213 | bullet | REQ-CL-0249, REQ-CL-0250 |
| 214 | bullet | REQ-CL-0251 |
| 215 | bullet | REQ-CL-0252 |
| 216 | bullet | REQ-CL-0253 |
| 217 | bullet | REQ-CL-0254 |
| 218 | bullet | REQ-CL-0255 |
| 219 | bullet | REQ-CL-0256, REQ-CL-0257 |
| 220 | bullet | REQ-CL-0258 |
| 221 | bullet | REQ-CL-0259 |
| 222 | bullet | REQ-CL-0260 |
| 223 | bullet | REQ-CL-0261 |
| 224 | bullet | REQ-CL-0262 |
| 225 | bullet | REQ-CL-0263 |
| 226 | bullet | REQ-CL-0264 |
| 227 | bullet | REQ-CL-0265 |
| 228 | bullet | REQ-CL-0266 |
| 229 | bullet | REQ-CL-0267, REQ-CL-0268 |
| 230 | bullet | REQ-CL-0269 |
| 231 | bullet | REQ-CL-0270, REQ-CL-0271 |
| 232 | bullet | REQ-CL-0272 |
| 233 | bullet | REQ-CL-0273 |
| 234 | bullet | REQ-CL-0274, REQ-CL-0275 |
| 235 | bullet | REQ-CL-0276, REQ-CL-0277 |
| 236 | bullet | REQ-CL-0278 |
| 237 | bullet | REQ-CL-0279 |
| 238 | bullet | REQ-CL-0280 |
| 239 | bullet | REQ-CL-0281 |
| 240 | bullet | REQ-CL-0282 |
| 241 | bullet | REQ-CL-0283, REQ-CL-0284 |
| 242 | bullet | REQ-CL-0285 |
| 243 | bullet | REQ-CL-0286 |
| 244 | bullet | REQ-CL-0287 |
| 245 | bullet | REQ-CL-0288 |
| 246 | bullet | REQ-CL-0289 |
| 247 | bullet | REQ-CL-0290 |
| 248 | bullet | REQ-CL-0291, REQ-CL-0292 |
| 249 | bullet | REQ-CL-0293 |
| 250 | bullet | REQ-CL-0294, REQ-CL-0295 |
| 251 | bullet | REQ-CL-0296 |
| 252 | bullet | REQ-CL-0297 |
| 253 | bullet | REQ-CL-0298 |
| 254 | bullet | REQ-CL-0299 |
| 255 | bullet | REQ-CL-0300 |
| 256 | bullet | REQ-CL-0301 |
| 257 | bullet | REQ-CL-0302 |
| 258 | bullet | REQ-CL-0303 |
| 259 | bullet | REQ-CL-0304 |
| 260 | bullet | REQ-CL-0305 |
| 261 | bullet | REQ-CL-0306 |
| 262 | bullet | REQ-CL-0307, REQ-CL-0308 |
| 263 | bullet | REQ-CL-0309 |
| 264 | bullet | REQ-CL-0310 |
| 265 | bullet | REQ-CL-0311 |
| 266 | bullet | REQ-CL-0312, REQ-CL-0313 |
| 267 | bullet | REQ-CL-0314 |
| 268 | bullet | REQ-CL-0315 |
| 269 | bullet | REQ-CL-0316 |
| 270 | bullet | REQ-CL-0317 |
| 271 | bullet | REQ-CL-0318, REQ-CL-0319 |
| 272 | bullet | REQ-CL-0320 |
| 273 | bullet | REQ-CL-0321 |
| 274 | bullet | REQ-CL-0322 |
| 275 | bullet | REQ-CL-0323 |
| 276 | bullet | REQ-CL-0324 |
| 277 | bullet | REQ-CL-0325 |
| 278 | bullet | REQ-CL-0326 |
| 279 | bullet | REQ-CL-0327 |
| 280 | bullet | REQ-CL-0328 |
| 281 | bullet | REQ-CL-0329 |
| 282 | bullet | REQ-CL-0330 |
| 283 | bullet | REQ-CL-0331 |
| 284 | bullet | REQ-CL-0332 |
| 285 | bullet | REQ-CL-0333 |
| 286 | bullet | REQ-CL-0334 |
| 287 | bullet | REQ-CL-0335 |
| 288 | bullet | REQ-CL-0336 |
| 289 | bullet | REQ-CL-0337 |
| 290 | bullet | REQ-CL-0338 |
| 291 | bullet | REQ-CL-0339 |
| 292 | bullet | REQ-CL-0340 |
| 293 | bullet | REQ-CL-0341 |
| 294 | bullet | REQ-CL-0342 |
| 295 | bullet | REQ-CL-0343 |
| 296 | bullet | REQ-CL-0344 |
| 297 | bullet | REQ-CL-0345 |
| 298 | bullet | REQ-CL-0346 |
| 299 | bullet | REQ-CL-0347 |
| 300 | bullet | REQ-CL-0348 |
| 301 | bullet | REQ-CL-0349 |
| 302 | bullet | REQ-CL-0350 |
| 303 | bullet | REQ-CL-0351 |
| 304 | bullet | REQ-CL-0352 |
| 305 | bullet | REQ-CL-0353 |
| 306 | bullet | REQ-CL-0354 |
| 307 | bullet | REQ-CL-0355 |
| 308 | bullet | REQ-CL-0356 |
| 309 | bullet | REQ-CL-0357 |
| 310 | bullet | REQ-CL-0358 |
| 311 | bullet | REQ-CL-0359 |
| 312 | bullet | REQ-CL-0360 |
| 313 | bullet | REQ-CL-0361 |
| 314 | bullet | REQ-CL-0362 |
| 315 | bullet | REQ-CL-0363 |
| 316 | bullet | REQ-CL-0364 |
| 317 | bullet | REQ-CL-0365 |
| 318 | bullet | REQ-CL-0366 |
| 319 | bullet | REQ-CL-0367 |
| 320 | bullet | REQ-CL-0368 |
| 321 | bullet | REQ-CL-0369 |
| 322 | bullet | REQ-CL-0370 |
| 323 | bullet | REQ-CL-0371, REQ-CL-0372 |
| 324 | bullet | REQ-CL-0373 |
| 325 | bullet | REQ-CL-0374 |
| 326 | bullet | REQ-CL-0375, REQ-CL-0376 |
| 327 | blank | context: blank line inside the Fixed list |
| 328 | bullet | REQ-CL-0377 |
| 329 | bullet | REQ-CL-0378 |
| 330 | blank | context: blank line inside the Fixed list |
| 331 | bullet | REQ-CL-0379 |
| 332 | bullet | REQ-CL-0380 |
| 333 | bullet | REQ-CL-0381 |
| 334 | bullet | REQ-CL-0382 |
| 335 | bullet | REQ-CL-0383 |
| 336 | bullet | REQ-CL-0384 |
| 337 | bullet | REQ-CL-0385 |
| 338 | bullet | REQ-CL-0386 |
| 339 | bullet | REQ-CL-0387 |
| 340 | bullet | REQ-CL-0388 |
| 341 | blank | context: blank line |
| 342 | heading | context: section heading |
| 343 | blank | context: blank line |
| 344 | bullet | REQ-CL-0389, REQ-CL-0390 |
| 345 | bullet | REQ-CL-0391, REQ-CL-0392 |
| 346 | bullet | REQ-CL-0393, REQ-CL-0394 |
| 347 | bullet | REQ-CL-0395 |
| 348 | bullet | REQ-CL-0396 |
| 349 | bullet | REQ-CL-0397 |
| 350 | bullet | REQ-CL-0398 |
| 351 | bullet | REQ-CL-0399 |
| 352 | bullet | REQ-CL-0400 |
| 353 | bullet | REQ-CL-0401 |
| 354 | bullet | REQ-CL-0402 |
| 355 | bullet | REQ-CL-0403 |
| 356 | bullet | REQ-CL-0404 |
| 357 | bullet | REQ-CL-0405 |
| 358 | bullet | REQ-CL-0406 |
| 359 | bullet | REQ-CL-0407 |
| 360 | blank | context: blank line |
| 361 | heading | context: version heading |
| 362 | blank | context: blank line |
| 363 | paragraph | context: paragraph: scope of the first changelog entry (history) |
| 364 | blank | context: blank line |
| 365 | heading | context: section heading |
| 366 | blank | context: blank line |
| 367 | bullet | REQ-CL-0408 |
| 368 | bullet | REQ-CL-0409 |
| 369 | bullet | REQ-CL-0410 |
| 370 | bullet | REQ-CL-0411 |
| 371 | bullet | REQ-CL-0412 |
| 372 | bullet | REQ-CL-0413 |
| 373 | bullet | REQ-CL-0414 |
| 374 | bullet | REQ-CL-0415, REQ-CL-0416 |
| 375 | bullet | REQ-CL-0417 |
| 376 | bullet | REQ-CL-0418 |
| 377 | bullet | REQ-CL-0419 |
| 378 | bullet | REQ-CL-0420 |
| 379 | bullet | REQ-CL-0421 |
| 380 | bullet | REQ-CL-0422 |
| 381 | bullet | REQ-CL-0423 |
| 382 | bullet | REQ-CL-0424 |
| 383 | bullet | REQ-CL-0425 |
| 384 | bullet | REQ-CL-0426 |
| 385 | bullet | context: plan statement (native rewrite and router ports planned; no router port exists) |
| 386 | blank | context: blank line |
| 387 | heading | context: section heading |
| 388 | blank | context: blank line |
| 389 | bullet | REQ-CL-0427 |
| 390 | bullet | REQ-CL-0428 |
| 391 | bullet | context: implementation history (dashboard split into files); the native page budgets are U-01's |
| 392 | bullet | REQ-CL-0429 |
| 393 | bullet | REQ-CL-0430 |
| 394 | bullet | REQ-CL-0431 |
| 395 | blank | context: blank line |
| 396 | heading | context: section heading |
| 397 | blank | context: blank line |
| 398 | bullet | context: removal of a timer; requirement is REQ at line 390 |
| 399 | bullet | REQ-CL-0432 |
| 400 | blank | context: blank line |
| 401 | heading | context: section heading |
| 402 | blank | context: blank line |
| 403 | bullet | REQ-CL-0433 |
| 404 | bullet | REQ-CL-0434 |
| 405 | bullet | REQ-CL-0435 |
| 406 | bullet | REQ-CL-0436 |
| 407 | bullet | REQ-CL-0437 |
| 408 | bullet | REQ-CL-0438 |
| 409 | bullet | REQ-CL-0439 |
| 410 | bullet | REQ-CL-0440 |
| 411 | bullet | REQ-CL-0441 |
| 412 | bullet | REQ-CL-0442, REQ-CL-0443, REQ-CL-0444 |
| 413 | bullet | REQ-CL-0445 |
| 414 | blank | context: blank line |
| 415 | heading | context: section heading |
| 416 | blank | context: blank line |
| 417 | bullet | REQ-CL-0446 |
| 418 | bullet | REQ-CL-0447 |
| 419 | bullet | REQ-CL-0448, REQ-CL-0449 |
| 420 | bullet | REQ-CL-0450, REQ-CL-0451 |
| 421 | blank | context: blank line |
| 422-423 | paragraph (link reference) | context: link reference (compare link; names an account) (source names a specific site value; not copied) |

## Proposed package changes

1. **NEW: P-16 SSH access (keys only).** Source line 27. No package of work plan section 3 owns the box's SSH server setting: the keys-only switch, the check that sshd really takes it ahead of the base image's own settings, and the automatic switch back. The Python code sits in `nmb/web/server.py` handlers and `bin/nmb-helper`, so the API part stays with `S-02…S-0n (API: system)` and the privileged write with P-06; P-16 owns the behaviour and its live test. Another option is to merge it into P-05 (maintenance mode, which also gives keys-only access), if the owner prefers fewer packages.
2. **Optional split: N-20 Box resolver and DNS fallback**, out of M-12. Source lines 210, 224, 254. These rows are assigned to M-12 because their Python reference is `nmb/netd/dns.py`, which the work plan gives to M-12, but they are the box's own resolver management (fallback resolvers, restore, marker), not a measurement. If M-12's analyst finds the package too big (DNS audit, STUN, NAT, public address and the fallback), split `nmb/netd/dns.py` into M-12 (audit, lookups) and N-20 (fallback, restore), and move REQ rows of lines 210, 224 and 254 there.
3. **Fix the numbers of the API and page packages.** This part uses `S-02…S-0n (API: <area>)` and `U-02…U-0n (page: <page>)`, like the TESTING.md part. Proposal, so every row gets one id: S-02 status, S-03 speed, S-04 network, S-05 Wi-Fi, S-06 VPN, S-07 blocker, S-08 system, S-09 exports (this fills S-02 to S-09 exactly, before S-10); U-02 home, U-03 speed (Measurement & Servers), U-04 network and router (including ports and drivers, DynDNS), U-05 Wi-Fi, U-06 VPN, U-07 blocker, U-08 system and settings, U-09 tools (device explorer, sniffer, traceroute, task manager, web terminal), U-10 results, logout and public pages. The integrator may renumber; the area words in this part make the mapping mechanical.

## Questions for the owner

1. **Supported distributions (lines 95, 96, 131).** The Python version was installed and tested on Debian, Ubuntu, Alpine, Fedora/RHEL family and Rocky Linux. Must the native packages still install on all of these (REQ rows of lines 95 and 96, package P-11), or only on the own distribution and the bases chosen after the image analysis? Until answered, the rows stay with P-11 marked "question for the owner".
2. **Factory login warning (line 211).** The Python installer warns on every run while the Web UI factory login is still set. The owner keeps the default login on purpose. Should the native installer keep this warning (it is only a warning), drop it, or show it once? The row is kept as written until answered.
3. **Experimental stop on network traffic (line 373)** and **web terminal (line 374).** Both exist in the Python version. Are they carried into the native build (rows assigned to M-03 and to the tools page), or left out on purpose? The web terminal gives a shell from the browser, so the security review of the native build will ask about it in any case.
4. **Fast.com over plain HTTP (line 82).** Python uses plain HTTP because TLS used most of a Pi 4's CPU. Is this a fixed rule for the native engine, or may the native engine use TLS when it measures as well? The row records the Python behaviour; M-07's analyst should compare both.
