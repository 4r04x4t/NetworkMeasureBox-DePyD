# Requirements register, part: README

Source: `README.md` of the Python repository at commit `ad3758c` (487 lines; this part covers lines 1-487, the whole file, read with `git show ad3758c:README.md`; the working tree was at the same commit).
Read in pages: 1-150, 151-300, 301-450, 451-487 (all lines).
Compared with `README.hu.md@ad3758c` (721 lines, read in pages 1-150, 151-300, 301-450, 451-600, 601-721, all lines): content that exists only in the Hungarian file is listed as its own rows (REQ-RE-0312 to REQ-RE-0338, marked "HU only") or added as a second source to the English row it extends; Hungarian statements that contradict the English file are listed under "Coverage" (README.hu.md table), where the English text wins.
Package ids from `docs/NATIVE_WORK_PLAN.md@ad3758c:98-218` (section 3) and `:281-294` (image packages I-01..I-10).
Id range of this part: REQ-RE-0001 to REQ-RE-0338.

Reading notes for the register owner:
- `NEW: X2-01 Step-2 router tiers` is the holding package proposed in `roadmap-1-507.md` and reused by `porting-plan.md`; it is reused here for the router-target rows of README section 16, so the register keeps one id for it.
- `U-02…U-0n (page: X)` and `S-02…S-0n (API: X)` name the page or API area so the split of those package ranges can pick the row up. The page list in REQ-RE-0327 to -0333 (README.hu.md section 8) is a ready list for that split.
- Quality Bar 8: the README names the test site's ISP, its router model, a tethering device and the repository's account (in clone URLs), and gives donation addresses and an example box ID. None is copied; rows that drop one say so.
- The README describes the Python version, so some rows name Python tools (`dnsmasq`, `hostapd`, NetworkManager, `iw`, `malloc_trim`). The requirement is the behaviour; the Notes say where the native build may do it differently.

## Requirements

| Id | Requirement | Source | Package | Also in | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| REQ-RE-0001 | The README states the project's licence with a badge, a link to `LICENSE` and a licence section; the Python version is under the MIT License, free for personal, community and commercial use. | README.md@ad3758c:5; README.md@ad3758c:485-487 | P-08 | P-12 | The native build's licence follows from its components (owner decision, native folder). This row asks for a licence statement of the same kind, not for MIT. |
| REQ-RE-0002 | Router ports are documented as planned only: NMB is developed and run on the Raspberry Pi 4, no router software or firmware exists, and router figures are design targets, not measurements. | README.md@ad3758c:9; README.md@ad3758c:265 | P-12 | NEW: X2-01 Step-2 router tiers, P-13 | Honest results rule. Stays true until a router build has passed live. |
| REQ-RE-0003 | The English README links the full Hungarian guide (`README.hu.md`). | README.md@ad3758c:12 | P-12 | | |
| REQ-RE-0004 | The README points to a support section: NMB is free and open source and developed in free time; a donation helps; the section tells the reader to double-check a pasted address and to send only the asset of that network. | README.md@ad3758c:16; README.md@ad3758c:474 | P-08 | P-12 | |
| REQ-RE-0005 | NMB runs 24/7 and records connection quality, to give evidence when a problem is disputed with the ISP. | README.md@ad3758c:22 | M-01 | P-01, R-03 | |
| REQ-RE-0006 | The box can run behind the router (`standalone`), as a router (Mode A), or as a transparent bridge between modem and router (Mode B). | README.md@ad3758c:26 | N-03 | N-04 | |
| REQ-RE-0007 | The box keeps latency, jitter, packet loss, outage and speed data for reports. | README.md@ad3758c:26 | C-02 | M-01, M-02, M-04, R-01 | |
| REQ-RE-0008 | The user is told that the box sees only traffic passing through it, so where it is connected affects what it can measure. | README.md@ad3758c:26 | M-03 | C-11, P-12, U-02…U-0n (page: Home) | |
| REQ-RE-0009 | The raw CSV telemetry export follows RFC 4180. | README.md@ad3758c:52 | R-02 | F-04 | The detailed spreadsheet table uses `;` and a UTF-8 BOM (REQ-RE-0295), a deliberate variant. |
| REQ-RE-0010 | The system journal is kept volatile (in RAM), to protect the flash. | README.md@ad3758c:61 | P-11 | C-03 | Crash evidence must still survive a reboot (REQ-RE-0186). |
| REQ-RE-0011 | Automatic mode (`mode = auto`) is the default for new installs. | README.md@ad3758c:82 | N-04 | C-01 | |
| REQ-RE-0012 | In automatic mode, when a router gives the box an address over DHCP the box runs `standalone`; on a box with several wired ports the other ports are bridged to it (LAN passthrough, almost like a switch). | README.md@ad3758c:82 | N-04 | N-06 | |
| REQ-RE-0013 | With no router and no internet for 90 s (cable on an ONT or modem, on a PC, or no cable) the box runs router mode until a router answers again: it dials PPPoE when a PPPoE login is saved and a PPPoE server answers, otherwise it offers the Wi-Fi access point and the LAN ports for setup. | README.md@ad3758c:82 | N-04 | N-03, N-05, N-07 | |
| REQ-RE-0014 | A router appearing on the WAN port switches automatic mode to `standalone` at once. | README.md@ad3758c:82 | N-04 | N-01 | |
| REQ-RE-0015 | A fixed mode, a manual override or the owner's own passthrough or aggregation settings always win over the automatic choice. | README.md@ad3758c:82 | N-04 | N-03, N-06 | |
| REQ-RE-0016 | Automatic mode is proven by unit tests and live before it is called working (the source says "Unit-tested; not yet tested live"). | README.md@ad3758c:82 | N-04 | P-13 | Stale in the source: Tested Status (line 314) says automatic mode chose router mode on a real PPPoE line on 2026-11-09. |
| REQ-RE-0017 | Mode A router WAN: DHCP, static IP or PPPoE; PPPoE in the kernel, with 802.1Q VLAN tagging and MSS clamping. | README.md@ad3758c:84 | N-05 | N-03 | Python: tested on a real 1000/300 Mbps PPPoE line on 2026-11-09. |
| REQ-RE-0018 | PPPoE redials after a dropped line and waits longer after a refused login. | README.md@ad3758c:84 | N-05 | | |
| REQ-RE-0019 | In Mode A the WAN is closed to the internet except replies, ICMP and the VPN ports that are on; a home router in front keeps full access for the owner. | README.md@ad3758c:85 | N-02 | N-03, N-13, N-14 | |
| REQ-RE-0020 | Port forwarding, switchable in the Web UI (Router > Port forwarding): rules for TCP, UDP or both, single ports or ranges, with hairpin NAT. | README.md@ad3758c:86 | N-11 | N-02, U-02…U-0n (page: Router), S-02…S-0n (API: network) | |
| REQ-RE-0021 | UPnP / NAT-PMP / PCP for LAN devices: off by default; a device may map only its own address, and only ports 1024-65535. | README.md@ad3758c:86 | N-11 | N-02 | |
| REQ-RE-0022 | IPv6 for the LAN on a PPPoE line: the ISP's delegated prefix (DHCPv6-PD) is advertised to the clients (SLAAC); LAN devices are not reachable from the internet. | README.md@ad3758c:87 | N-11 | N-02, N-05 | |
| REQ-RE-0023 | Mode A provides the local gateway services: DHCP server, DNS caching and NAT masquerading. | README.md@ad3758c:88 | N-03 | N-02 | Python uses `dnsmasq` and iptables/nftables; the native build may provide them itself (contract in F-04). |
| REQ-RE-0024 | The Wi-Fi access point runs on the LAN bridge in router mode only; in standalone and bridge mode the radio stays a Wi-Fi client of the home network. | README.md@ad3758c:89; README.hu.md@ad3758c:672 | N-07 | N-03 | |
| REQ-RE-0025 | The access point adapts to the radio and steps down when the access point daemon rejects a setting. | README.md@ad3758c:89 | N-07 | | See REQ-RE-0203 (Wi-Fi 7). |
| REQ-RE-0026 | Mode B bridges an existing modem/ONT and router at Layer 2; the box requests its own address over DHCP so it can measure, so the upstream must offer one. | README.md@ad3758c:91 | N-03 | N-05 | |
| REQ-RE-0027 | Mode B measures on the bridge (`br0`). | README.md@ad3758c:92 | N-03 | M-01 | |
| REQ-RE-0028 | Documented limits of Mode B: if the box stops, the connection stops too; PPPoE or single-address connections are not handled yet. | README.md@ad3758c:92 | N-03 | P-12 | README.hu.md:169 claims a hardware bypass; see Questions. |
| REQ-RE-0029 | WAN failover: with WAN2 configured, Mode A moves the default route to WAN2 after failed connectivity checks. | README.md@ad3758c:93 | N-03 | N-01 | |
| REQ-RE-0030 | The Web UI has a port role matrix that maps physical network cards (PCIe, USB, onboard) to roles: WAN 1 (primary uplink), WAN 2, LAN (`br-lan`), DMZ, guest Wi-Fi client probe. | README.md@ad3758c:96-97; README.hu.md@ad3758c:177-184 | N-01 | U-02…U-0n (page: Ports & Drivers), S-02…S-0n (API: network), C-01 | The "not assigned" role is REQ-RE-0316 (HU only). |
| REQ-RE-0031 | WAN 2: failover backup, or load balancing through a multipath default route, with its own DHCP or static settings. | README.md@ad3758c:98 | N-03 | N-01 | Python: not tested on real lines (see REQ-RE-0206). |
| REQ-RE-0032 | LAN role: any number of physical Ethernet ports are grouped automatically into the bridge `br-lan` (`192.168.100.1/24`), with the DHCP server and NAT bound to it. | README.md@ad3758c:99; README.hu.md@ad3758c:181 | N-03 | N-01 | DHCP and NAT binding from the Hungarian file. |
| REQ-RE-0033 | DMZ in Mode A: the DMZ port gets its own subnet (`192.168.200.0/24`) with DHCP; its devices reach the internet but not the LAN, and the box only for DHCP and DNS. | README.md@ad3758c:100 | N-02 | N-03 | |
| REQ-RE-0034 | Guest Wi-Fi client probe role: a dedicated wireless card for client probe tests. | README.md@ad3758c:101 | N-08 | N-01 | Details in REQ-RE-0314 (HU only). |
| REQ-RE-0035 | Storage guard (`[storage]`, on by default): the network daemon checks the disk every 10 minutes and keeps free a size-dependent share: 15 % up to 2 GB, 10 % up to 16 GB, 5 % above; at least 200 MB, at most 2 GB; `min_free_mb` overrides it. | README.md@ad3758c:102 | C-08 | C-01 | |
| REQ-RE-0036 | When less is free, the guard deletes oldest first in this order: rotated logs, diagnostic bundles and crash reports; update backups except the newest; old reports; the system journal down to 32 MB; large active NMB logs; and only then measurements, a day at a time, keeping the last `keep_days` (7) days. | README.md@ad3758c:102 | C-08 | C-02, C-03 | |
| REQ-RE-0037 | Measurements deleted by the storage guard are recorded as an event. | README.md@ad3758c:102 | C-08 | C-07 | |
| REQ-RE-0038 | The installer stops when less than 500 MB is free, and uses smaller logs on disks under 8 GB. | README.md@ad3758c:102 | P-11 | C-08 | |
| REQ-RE-0039 | `/var/log` is a RAM disk of its own (48 MB, 16 MB with 512 MB RAM): the guard keeps 20 % of it free, the developer log rotates within 40 % of it (3 x 5 MB by default), and the installer makes room there before it writes anything. | README.md@ad3758c:102 | C-08 | C-03, P-11 | |
| REQ-RE-0040 | Measuring over Wi-Fi: the box can join the owner's Wi-Fi (`[wifi_uplink]`: `ssid`, `password`, `interface`, `hidden`, `enabled`) with its own connection profile (`nmb-uplink`), separate from the guest client. | README.md@ad3758c:103 | N-08 | C-01, N-07 | Python needs NetworkManager for it; the native build may manage the client itself. |
| REQ-RE-0041 | `[network] measurement_link` chooses where speed tests run: `auto` (Ethernet when a cable is connected, otherwise Wi-Fi), `ethernet` or `wifi` (and `usb`, REQ-RE-0048). | README.md@ad3758c:103 | M-04 | N-08, N-10, C-01, C-11 | |
| REQ-RE-0042 | With a cable and Wi-Fi both up, the default route stays on Ethernet, and a source rule (table 150) makes tests bound to Wi-Fi really leave over Wi-Fi. | README.md@ad3758c:103 | N-01 | N-08, M-04 | |
| REQ-RE-0043 | The Wi-Fi uplink is set in the Web UI under the schedule, or with `nmb-cli run wifi.uplink`. | README.md@ad3758c:103 | N-08 | U-02…U-0n (page: Schedule), S-10, S-02…S-0n (API: Wi-Fi) | |
| REQ-RE-0044 | Each wireless interface can be given its role in Ports & Drivers: guest Wi-Fi client or main Wi-Fi client (measuring); one interface cannot be both, and with `interface = auto` the main client skips the radio of an enabled guest client. | README.md@ad3758c:103 | N-08 | U-02…U-0n (page: Ports & Drivers) | |
| REQ-RE-0045 | Settings from a USB stick: an `nmb.conf` with only the keys to change at the top level of a stick is applied once when plugged in; the old settings are backed up; a set secret is never wiped by an empty value; `nmb-import-result.txt` (names only) is written back to the stick. | README.md@ad3758c:104 | P-03 | C-01, C-04 | |
| REQ-RE-0046 | USB import has a template (`config/usb-import.example.conf`) and is switched off with `[usb_import] enabled = false`; it is tested with a real stick before it is called working. | README.md@ad3758c:104 | P-03 | P-13 | Python: not tested with a real stick yet. |
| REQ-RE-0047 | USB tethering (`[usb_uplink]`, off by default): a phone, mobile hotspot or cellular modem on USB is recognised by its driver (`rndis_host`, `cdc_ncm`, `cdc_ether`, `ipheth`, `qmi_wwan`, `cdc_mbim`) and never taken for the wired line. | README.md@ad3758c:105 | N-10 | N-01, N-17 | Source names a specific device model as example (source names a specific site value; not copied). |
| REQ-RE-0048 | Turned on, USB tethering gets its own profile (`nmb-usb-uplink`) with an address but never the default route, so the box's traffic and the outage probe stay on the cable; `measurement_link = usb` (or a manual test on that interface) measures over it through a source rule (table 152). | README.md@ad3758c:105 | N-10 | N-01, M-04, M-01 | |
| REQ-RE-0049 | When USB is the only internet (no wired or Wi-Fi default route), the box uses it automatically: a default route with metric 900, removed when the cable or Wi-Fi is back; each switch is an event. | README.md@ad3758c:105 | N-10 | N-01, C-07 | |
| REQ-RE-0050 | In router mode, when no DHCP server answers on the Ethernet port, USB becomes the WAN and the single Ethernet port plus the Wi-Fi access point share the mobile internet; a DHCP server appearing on the port stops the box's DHCP at once. | README.md@ad3758c:105 | N-10 | N-03, N-05, N-07 | |
| REQ-RE-0051 | The LAN-port sharing (`[wifi_share]`) also takes USB as its upstream. | README.md@ad3758c:105 | N-08 | N-10 | |
| REQ-RE-0052 | Results measured over USB are stored as mobile (`fwa`). | README.md@ad3758c:105 | N-10 | C-11, C-02 | |
| REQ-RE-0053 | Tests over USB have their own daily data limit (`daily_limit_mb`, default 1024 MB, 0 = none). | README.md@ad3758c:105 | M-05 | N-10, C-01 | |
| REQ-RE-0054 | USB tethering has a Web UI card and `nmb-cli run usb.uplink ...`, in all 11 languages. | README.md@ad3758c:105 | N-10 | U-02…U-0n (page: Schedule), S-10, C-06 | |
| REQ-RE-0055 | USB tethering is proven live: detection and speed tests over USB, the box's own fallback to USB, router mode on USB sharing the internet to a wired client and a client on the box's Wi-Fi; Wi-Fi sharing on the LAN port through USB is not confirmed in Python and is proven before it is called working. | README.md@ad3758c:105 | N-10 | P-13 | Python status: fallback worked once; router mode on USB worked live on 2026-11-08 after two fixes. Test device not copied (source names a specific site value; not copied). |
| REQ-RE-0056 | Data caps for automatic tests (`[data_cap]`, Web UI card under the schedule): a cap per link (Ethernet/PPPoE, Wi-Fi, USB, link aggregation) and per connection type (fiber, cable, DSL, mobile, satellite ...), per day or per month with a reset day, counted from the data the tests use. | README.md@ad3758c:106 | M-05 | C-01, C-11, U-02…U-0n (page: Schedule) | Rationale given: an every-engine schedule on a gigabit line can move hundreds of GB a day. |
| REQ-RE-0057 | When a cap is reached, automatic tests on that link stop until the period ends; tests started by hand still run. | README.md@ad3758c:106 | M-05 | M-04 | |
| REQ-RE-0058 | The scheduled mixed run can leave engines out (`[speedtest] mixed_engines`, a tick per engine). | README.md@ad3758c:106 | M-04 | C-01, U-02…U-0n (page: Schedule) | |
| REQ-RE-0059 | Wi-Fi shared on a LAN port (`[wifi_share]`, off by default): when the box is on Wi-Fi and a wired port has no internet of its own, devices on that port get an address from the box (`192.168.77.1/24` by default, DHCP .100-.200, the Wi-Fi's DNS servers) and reach the internet through the box's Wi-Fi (NAT, routing table 151); Web UI card under the schedule and `nmb-cli run wifi.share enabled=true lan_interface=eth0`. | README.md@ad3758c:107 | N-08 | N-02, U-02…U-0n (page: Schedule), S-10 | |
| REQ-RE-0060 | DHCP guard: the box never hands out addresses on a network that has its own DHCP; it sends a DHCP request on the port before it starts, every minute while it runs, and right after a cable is plugged in (with its own DHCP stopped first), and stops if any other server answers or the port gets a default route. | README.md@ad3758c:107 | N-05 | N-08, N-10 | Work plan puts the guard in N-05. |
| REQ-RE-0061 | The DHCP guard reads the replies from a packet socket, so the box's firewall cannot hide them. | README.md@ad3758c:107 | N-05 | N-02 | |
| REQ-RE-0062 | LAN passthrough (standalone probe behind a router; Web UI Ports & Drivers): selected ports are bridged to the uplink at Layer 2, so devices behind the box stay on the router's subnet. | README.md@ad3758c:108-109 | N-06 | U-02…U-0n (page: Ports & Drivers) | |
| REQ-RE-0063 | Link aggregation: two or more wired ports or cards are combined into `bond0` (802.3ad LACP, balance-rr, balance-xor, balance-tlb, balance-alb, active-backup), with or without passthrough. | README.md@ad3758c:110 | N-06 | | |
| REQ-RE-0064 | The user is told that 802.3ad needs LACP on the router or switch, balance-rr and balance-xor need a static LAG there, and except with balance-rr one TCP connection stays on one link, so the combined speed appears only with multi-stream tests. | README.md@ad3758c:111 | N-06 | P-12, U-02…U-0n (page: Ports & Drivers) | |
| REQ-RE-0065 | Speed tests can be pinned to a network interface from the CLI (`-i` / `--interface`) or a Web UI drop-down, to compare WAN1 and WAN2 throughput. | README.md@ad3758c:112 | M-04 | S-10, U-02…U-0n (page: Speed), N-01 | |
| REQ-RE-0066 | The driver finder detects physical controllers (Realtek RTL8111/8168/8125/8126, Intel e1000e/igb/igc/ixgbe, Broadcom, MediaTek MT7921, Qualcomm Atheros, ASIX USB, Aquantia 10G). | README.md@ad3758c:115 | N-17 | | Longer list in REQ-RE-0317 (HU only). |
| REQ-RE-0067 | The driver finder inspects the kernel ring buffer to detect missing non-free firmware. | README.md@ad3758c:116 | N-17 | C-03 | |
| REQ-RE-0068 | One-click driver and firmware package installer through the privileged helper (`install-drivers`), with real-time terminal log in the Web UI; tested with `apt` only; the `opkg`, `apk`, `pacman` and `dnf` paths exist but are untested. | README.md@ad3758c:117 | N-17 | P-06, U-02…U-0n (page: Ports & Drivers) | The source says "tested on Raspberry Pi OS (apt)", while line 321 says Raspberry Pi OS was not tested. README.hu.md:87 lists all package managers as supported; the English wording (untested) wins. |
| REQ-RE-0069 | ICMP probe: 500 ms continuous sampling to the gateway and to public targets (by default `1.1.1.1`, `8.8.8.8`, `9.9.9.9`, `185.222.222.222`). | README.md@ad3758c:120 | M-01 | C-01 | README.hu.md:199-203 lists only three public targets; README.hu.md:97 has all four. |
| REQ-RE-0070 | The probe computes RFC 3550 jitter and p95/p99 latency, with average, minimum and maximum RTT. | README.md@ad3758c:120; README.hu.md@ad3758c:204-205 | M-01 | R-01 | Average, minimum and maximum from the Hungarian file. |
| REQ-RE-0071 | Every outage is logged with its exact start, end and duration (to the second). | README.md@ad3758c:120; README.hu.md@ad3758c:206 | M-02 | C-02 | |
| REQ-RE-0072 | Ookla engine: runs the official Ookla binary with automatic server detection or a chosen server; the server list always comes from Ookla (the servers nearest to the box, refreshed every few minutes or with one click); no list is built in. | README.md@ad3758c:121 | M-10 | U-02…U-0n (page: Speed) | Owner decision: Ookla "whichever works". |
| REQ-RE-0073 | Scheduled Ookla tests run at most once an hour (`ookla_interval_min`, at least 60), however often the schedule runs; a manual Ookla test runs when asked. | README.md@ad3758c:122-123 | M-04 | M-10, C-01 | |
| REQ-RE-0074 | When Ookla refuses ("Limit reached"), the box pauses Ookla (30 min, doubling up to 4 h, kept across restarts) instead of retrying, and stores no other engine's result in its place. | README.md@ad3758c:122-123 | M-05 | M-10, C-02 | |
| REQ-RE-0075 | The Web UI shows the hourly Ookla limit and any pause. | README.md@ad3758c:122 | U-02…U-0n (page: Speed) | M-05 | |
| REQ-RE-0076 | NMB does not change the public IP address by itself to get around an Ookla block; manual tests are not held back; the server list refresh also counts as a request to Ookla. | README.md@ad3758c:123 | M-05 | M-10, P-12 | Source tells the test line's history with dates, a router model and figures (source names a specific site value; not copied). Finding kept: the block is on the public IP and covers every device behind it. |
| REQ-RE-0077 | Cloudflare engine: timed 4-stream download and upload against the Cloudflare edge (about 1.2 GB per run at gigabit speed). | README.md@ad3758c:124 | M-06 | M-05 | README.hu.md:215 (100 MB and 25 MB payloads) is stale; see Coverage. |
| REQ-RE-0078 | Fast.com engine built in (no Node, Go or browser), using Fast.com's own API as the known fast-cli clients do: Fast.com names up to five Open Connect servers near the box; the box measures latency to each (the `range/0-0` request Fast.com times, plus the TCP handshake). | README.md@ad3758c:125 | M-07 | | API and token detail in REQ-RE-0319 (HU only). |
| REQ-RE-0079 | Fast.com runs over plain HTTP with 16 parallel streams spread over the servers, 15 s each way, leaving out the first 3 s of TCP ramp-up. | README.md@ad3758c:125 | M-07 | | Reason given: on a Pi 4 TLS took 77-88 % CPU at 910 Mbps against 50-56 % for HTTP; in isolated runs every variant reached the line speed (Python figures). |
| REQ-RE-0080 | The Fast.com result is the "stable" speed, as Fast.com shows it: the mean from the moment the speed settles. | README.md@ad3758c:125 | M-07 | R-01 | |
| REQ-RE-0081 | Each Fast.com row keeps in `raw_json`: the ISP, ASN, city and country Fast.com sees; every server with city, country, latency samples and its share of download and upload; the speed every half second; peak and plain mean; latency under load both ways; the bufferbloat grade. | README.md@ad3758c:125; README.hu.md@ad3758c:221 | M-07 | F-04, C-02 | |
| REQ-RE-0082 | Fast.com is part of every mixed run, like Cloudflare (`fastcom_interval_min = 0`); about 2.2 GB per run at gigabit speed. | README.md@ad3758c:125 | M-04 | M-07, M-05 | |
| REQ-RE-0083 | RTR-Netztest (RMBT) engine, the open-source benchmark of the Austrian regulator. | README.md@ad3758c:126 | M-08 | | |
| REQ-RE-0084 | iPerf3 client and managed server: LAN and public WAN throughput over TCP and UDP with parallel streams. | README.md@ad3758c:127 | M-09 | | |
| REQ-RE-0085 | iPerf3 results, statistics and the CSV name each server (place and provider from the public iPerf3 list; for a server known only by address, city, country and ISP from GeoIP), also for older rows that stored only the address. | README.md@ad3758c:127 | M-09 | R-01, R-02, M-12, C-02 | |
| REQ-RE-0086 | Port throttling check: compares TCP throughput on common ports (80, 443, 8080, 8443, 53, 22, 5060, 6881 by default) against a port 443 baseline; the result says it can show a difference but cannot prove the ISP caused it. | README.md@ad3758c:128 | M-11 | R-01, U-02…U-0n (page: Port check) | README.hu.md:233 claims it checks whether the ISP throttles; the English wording wins. |
| REQ-RE-0087 | Traffic gate: reads the interface counters and postpones a test when the box sees traffic above the threshold (2 % of capacity by default). | README.md@ad3758c:129 | M-03 | M-04 | |
| REQ-RE-0088 | In plain `standalone` mode the gate sees only the box's own traffic, so other devices do not count; other traffic reaches its counters only when it passes through the box. | README.md@ad3758c:129 | M-03 | P-12 | |
| REQ-RE-0089 | Stop on network traffic (experimental, off by default), switched on with `experimental_stop = true` under `[traffic_gate]` or a Web UI switch. | README.md@ad3758c:130 | M-03 | C-01, U-02…U-0n (page: Schedule) | |
| REQ-RE-0090 | With stop on traffic on, scheduled tests need a limit of at most 500 kbps, retry every 5 minutes after a deferral, and stop without saving partial results when other traffic appears; ping keeps running and manual tests are not stopped. | README.md@ad3758c:130 | M-03 | M-04, M-01 | A voice call fits under the limit; video and torrents do not. |
| REQ-RE-0091 | Stop on traffic reads other devices' traffic from client-side ports (bridge, gateway or the box's own access point) or from the router's UPnP counters; when neither is available, no scheduled speed tests run. | README.md@ad3758c:130 | M-03 | N-11, N-03 | |
| REQ-RE-0092 | Stop on traffic is tried on a real bridge before it is called working. | README.md@ad3758c:130 | M-03 | P-13 | Python: not yet tried on a real bridge. The wiring advice (bridge between ISP endpoint and router, a second USB 3 adapter on a Pi 4) is context. |
| REQ-RE-0093 | OONI probe (optional, off by default): a weekly censorship and net-neutrality run, enabled in the Web UI or with `enabled = true` under `[ooni]`; it does not use the traffic gate. | README.md@ad3758c:131 | M-13 | C-01, M-03 | |
| REQ-RE-0094 | New installs test every hour (`schedule_cron = 0 * * * *`), changeable in the Web UI. | README.md@ad3758c:133 | M-04 | C-01, U-02…U-0n (page: Schedule) | |
| REQ-RE-0095 | The schedule sets how often the box checks; each engine has its own minimum interval under `[speedtest]` or in the Web UI (`ookla_interval_min` at least 60; `rtr_interval_min` and `cloudflare_interval_min`, 0 = every run; `fastcom_interval_min`, 0 by default), so RTR can run every minute while Ookla runs hourly. | README.md@ad3758c:133 | M-04 | C-01 | |
| REQ-RE-0096 | iPerf3 runs on the same schedule with its own `iperf3_interval_min`, against one or more servers (`target_host` and `extra_servers`; default 4 TCP streams, 10 s, download then upload), alone or in the mixed run, but only when an iPerf3 server is set under `[iperf]` and answers; otherwise it is left out without an empty row. | README.md@ad3758c:133 | M-09 | M-04, C-01 | |
| REQ-RE-0097 | The speed test list shows only real measurements; skipped and failed runs are counted in one line. | README.md@ad3758c:133 | U-02…U-0n (page: Speed) | S-02…S-0n (API: speed), M-04 | |
| REQ-RE-0098 | Scheduled tests run as long as the scheduling service runs, and upgrades keep the configured schedule. | README.md@ad3758c:133 | P-02 | M-04, C-01 | In Python `nmb-web` starts scheduled tests; in the native build the scheduler is M-04 inside `nmbd`. |
| REQ-RE-0099 | Memory-saving profiles exist for constrained Linux systems; the documentation says where they have run (Python: only on the Raspberry Pi 4, not on routers). | README.md@ad3758c:136 | C-08 | P-12 | |
| REQ-RE-0100 | `low` profile: hot database in RAM limited to 48 MB, SQLite cache 2 MB (mmap 2 MB), TCP buffers limited to 8 MB. | README.md@ad3758c:137; README.hu.md@ad3758c:270-271 | C-08 | C-02, P-11 | mmap from the Hungarian file. |
| REQ-RE-0101 | Unused heap memory is released back to the kernel proactively (periodically and after measurements). | README.md@ad3758c:138; README.hu.md@ad3758c:272 | C-08 | | Python does it with `malloc_trim(0)` and garbage collection; the native mechanism is the runtime's own. |
| REQ-RE-0102 | `standard` (1-2 GB) and `performance` (4 GB+) profiles scale the buffers for multi-gigabit workloads. | README.md@ad3758c:139 | C-08 | P-11 | Figures in REQ-RE-0323 (HU only). |
| REQ-RE-0103 | The live database runs in tmpfs (`/var/lib/nmb/live/`) with SQLite write-ahead logging, keeping most writes away from flash. | README.md@ad3758c:142 | C-02 | | |
| REQ-RE-0104 | A consistent snapshot is copied with the online backup method to persistent storage (`/var/lib/nmb/persist/nmb.db`) every 15 minutes and at system shutdown. | README.md@ad3758c:143; README.hu.md@ad3758c:294 | C-02 | P-10 | Hungarian: on clean shutdown (`SIGTERM`), as one sequential write. |
| REQ-RE-0105 | Filesystem settings: ext4 with `noatime,commit=600`; `/tmp`, `/var/tmp` and `/var/log` in RAM, so logs are lost on restart (stated to the user). | README.md@ad3758c:144 | P-11 | C-03, P-12 | |
| REQ-RE-0106 | The Web UI listens on standard HTTP port 80, with an automatic unprivileged fallback to port 8080 for development. | README.md@ad3758c:147 | S-01 | P-10 | |
| REQ-RE-0107 | The Web UI is a dark-themed router dashboard. | README.md@ad3758c:147 | U-01 | | Native decision: AMOLED-only black look (U-01). |
| REQ-RE-0108 | Sub-navigation bars are horizontal swipe ribbons (no wrapping, container-relative scrolling), so buttons are not squeezed or wrapped letter by letter on 360-420 px phone screens. | README.md@ad3758c:148 | U-01 | | |
| REQ-RE-0109 | The Web UI has live inline SVG graphs and historical speed records. | README.md@ad3758c:149 | U-02…U-0n (pages: Ping, Speed) | R-03 | |
| REQ-RE-0110 | Network device explorer with ARP sweep and MAC vendor lookup. | README.md@ad3758c:149 | N-17 | U-02…U-0n (page: Devices) | Details in REQ-RE-0313 (HU only). |
| REQ-RE-0111 | Real-time system log inspector in the Web UI. | README.md@ad3758c:149 | S-02…S-0n (API: system) | C-03, U-02…U-0n (page: Developer console) | |
| REQ-RE-0112 | One-click generation of the diagnostic bundle for LLMs. | README.md@ad3758c:149 | R-04 | U-02…U-0n (page: AI diagnostics) | |
| REQ-RE-0113 | The Web UI, the terminal UI (`nmb-cli`), the results page, the charts and the dossier share the same text keys in 11 languages; some technical terms stay in English, and `nmb-cli` action names are English. | README.md@ad3758c:150 | C-06 | U-01, S-10, R-03 | |
| REQ-RE-0114 | The Telegram bot speaks all 11 languages (menus, answers, admin answers, alerts). | README.md@ad3758c:150 | S-11 | C-06 | Supersedes line 218 (five languages). |
| REQ-RE-0115 | The 11 languages are English, Hungarian, German, Ukrainian, French, Romanian, Italian, Spanish, Portuguese (Portugal), Portuguese (Brazil) and Chinese (Simplified). | README.md@ad3758c:151-161 | C-06 | | |
| REQ-RE-0116 | A missing text is shown in English, never in another language; the newer languages are one file per language in the Web UI and report locale folders. | README.md@ad3758c:162 | C-06 | U-01, R-03 | Owner rule: no Hungarian in foreign locales. |
| REQ-RE-0117 | The Web UI is protected by HTTP authentication and has a Logout. | README.md@ad3758c:163 | S-01 | U-01 | Python uses HTTP Basic; work plan S-01 names login, sessions and CSRF for the native build. Logout detail in REQ-RE-0333 (HU only). |
| REQ-RE-0118 | WireGuard VPN server (UDP 51820). | README.md@ad3758c:167 | N-13 | N-02 | |
| REQ-RE-0119 | OpenVPN server (UDP 1194, elliptic curve `prime256v1`). | README.md@ad3758c:167 | N-14 | N-02 | |
| REQ-RE-0120 | Dual-stack tunnels: IPv4 `10.8.0.0/24` and `10.9.0.0/24`, IPv6 ULA `fd42:42:42::/64` and `fd42:42:43::/64`. | README.md@ad3758c:168 | N-13 | N-14 | Which range belongs to which server is not said; taken from the code in the briefs. |
| REQ-RE-0121 | Client pairing: entering a client name generates the configuration profile and a camera-scannable QR code, in the Web UI and in Telegram. | README.md@ad3758c:169 | N-13 | N-14, S-11, U-02…U-0n (page: VPN) | |
| REQ-RE-0122 | Peer management: connected peer list with endpoint IPs, latest handshake, RX/TX transfer and revocation. | README.md@ad3758c:170 | N-13 | N-14, U-02…U-0n (page: VPN) | |
| REQ-RE-0123 | Tailscale: encrypted access to the Web UI from anywhere without public port forwarding, dynamic DNS or opening the firewall, also through carrier-grade NAT and router cascades. | README.md@ad3758c:172-173 | N-15 | N-18 | Remote access without Tailscale is N-18. |
| REQ-RE-0124 | The Web UI shows the MagicDNS name and the Tailscale IPv4/IPv6 status, with one-click connect and disconnect. | README.md@ad3758c:174 | N-15 | U-02…U-0n (page: Remote) | |
| REQ-RE-0125 | Built-in DuckDNS dynamic DNS updater. | README.md@ad3758c:176 | N-16 | | dynv6 also: REQ-RE-0195, -0266. |
| REQ-RE-0126 | Optional built-in HTTP CONNECT forward proxy with required authentication. | README.md@ad3758c:177 | N-11 | N-02 | |
| REQ-RE-0127 | DNS sinkhole: blocks tracking, telemetry, malware and advertisement domains before they reach LAN devices. | README.md@ad3758c:181 | N-12 | | |
| REQ-RE-0128 | Preconfigured blocklist subscriptions (for example StevenBlack unified hosts, AdAway) with automatic background updates. | README.md@ad3758c:182 | N-12 | | |
| REQ-RE-0129 | DNS statistics: total queries today, blocked queries, blocked share, active blocklist domain count. | README.md@ad3758c:184 | N-12 | U-02…U-0n (page: Blocker), S-11 | |
| REQ-RE-0130 | Custom domain blacklist, whitelist and wildcard regex rules (API `/api/adblock/regex/add`). | README.md@ad3758c:185 | N-12 | S-02…S-0n (API: blocker) | |
| REQ-RE-0131 | DNS query log in the Web UI. | README.md@ad3758c:186 | N-12 | U-02…U-0n (page: Blocker) | |
| REQ-RE-0132 | Wi-Fi telemetry: signal strength (RSSI) in dBm where the driver reports it. | README.md@ad3758c:190 | M-14 | | Python reads it through `iw`; native through nl80211. Noise floor: REQ-RE-0335 (HU only). |
| REQ-RE-0133 | Wi-Fi telemetry: RX/TX PHY bitrates, MCS index, channel width (20/40/80/160 MHz), and hardware retry/fail counters (`tx_retries`, `tx_failed`). | README.md@ad3758c:191 | M-14 | | |
| REQ-RE-0134 | DNS audit queries 4 resolver tiers concurrently: local gateway/DHCP, Cloudflare (`1.1.1.1`), Google (`8.8.8.8`), Quad9 (`9.9.9.9`). | README.md@ad3758c:193 | M-12 | | |
| REQ-RE-0135 | DNS audit checks DNSSEC validation (the `ad` bit) and compares cached and uncached response times. | README.md@ad3758c:194 | M-12 | | |
| REQ-RE-0136 | DNS interception test: a query to an address where no DNS server exists; any answer means port 53 is intercepted. | README.md@ad3758c:195 | M-12 | | |
| REQ-RE-0137 | STUN (RFC 5389) classifies the upstream: direct public IP (local WAN IP equals the reflexive address), carrier-grade NAT (WAN IP in `100.64.0.0/10`, RFC 6598), double NAT (WAN IP is an RFC 1918 private address). | README.md@ad3758c:196-200 | M-12 | C-11 | The source gives `192.168.x.x` or `10.x.x.x` as examples; RFC 1918 also has `172.16.0.0/12`. The M-12 brief checks the code; a missing range would be a Python fault (QB 6). |
| REQ-RE-0138 | NAT mapping is classified as endpoint-independent (same public port towards two STUN servers) or symmetric; full cone and restricted types are not distinguished. | README.md@ad3758c:201 | M-12 | | |
| REQ-RE-0139 | Per-core interrupt and SoftIRQ distribution is tracked, to catch CPU0 saturation. | README.md@ad3758c:203 | C-09 | | |
| REQ-RE-0140 | RPS/XPS bitmasks and NIC hardware offloads (checksumming, TSO, GRO) are inspected. | README.md@ad3758c:204 | C-09 | N-17 | Python uses `ethtool`. |
| REQ-RE-0141 | Lost packets are classified, separating micro-bursts from macro-outages (at least 5 s). | README.md@ad3758c:206 | M-02 | | |
| REQ-RE-0142 | Loss at the gateway and on public targets is compared to suggest whether a problem is local or further upstream; this is shown as an indication, not proof. | README.md@ad3758c:207 | M-02 | R-01, R-03 | README.hu.md:417 claims an automatic determination; the English wording wins. |
| REQ-RE-0143 | `%iowait` is tracked to detect failing or saturated flash storage. | README.md@ad3758c:209 | C-08 | C-09 | |
| REQ-RE-0144 | Free filesystem capacity is tracked. | README.md@ad3758c:210; README.hu.md@ad3758c:420 | C-08 | | Hungarian names `statvfs()`. |
| REQ-RE-0145 | Kernel log traps catch hardware I/O errors and ext4/squashfs corruption. | README.md@ad3758c:211 | C-08 | C-03 | |
| REQ-RE-0146 | A built-in redactor masks bot tokens, chat IDs, Wi-Fi keys, ISP account numbers and circuit IDs before dossiers or diagnostic snapshots are made (and before anything goes to an outside LLM or the ISP). | README.md@ad3758c:213; README.hu.md@ad3758c:422 | C-04 | R-03, R-04, S-12 | |
| REQ-RE-0147 | Telegram bot for management without a browser, with inline keyboards for status, ping matrix, on-demand speed tests, SLA dossier and VPN QR codes. | README.md@ad3758c:216-217; README.md@ad3758c:384; README.hu.md@ad3758c:429 | S-11 | | SLA dossier button from the Hungarian file. |
| REQ-RE-0148 | The bot sends the complete `.html` dossier into the chat. | README.md@ad3758c:219 | S-11 | R-03 | |
| REQ-RE-0149 | The bot alerts on outages, and when packet loss, latency or jitter stays above its threshold for 5 minutes. | README.md@ad3758c:220 | S-11 | M-02, C-01 | The diagram (line 67) and README.hu.md:109 say "loss > 5 %"; the S-11 brief takes the thresholds from the code. |
| REQ-RE-0150 | The dossier has a neutral title: "Broadband Quality & ISP SLA Performance Dossier" (and its Hungarian equivalent). | README.md@ad3758c:224 | R-03 | C-06 | |
| REQ-RE-0151 | The dossier makes no official or regulatory claim: the SLA result is NMB's own evaluation against fixed thresholds, not a decision by an authority or an ISP. | README.md@ad3758c:225 | R-01 | R-03, P-08 | README.hu.md:583 calls the dossier "official" (hivatalos); that is a fault of the Hungarian text. |
| REQ-RE-0152 | Guaranteed-speed cards with `0 Mb/s` or an unset value are suppressed, and disclaimers guide the SLA check. | README.md@ad3758c:226; README.hu.md@ad3758c:441 | R-03 | R-01 | |
| REQ-RE-0153 | Every measurement sample and telemetry table shows the detected client IPv4/IPv6 addresses with the ISP identity. | README.md@ad3758c:227 | R-03 | M-12 | |
| REQ-RE-0154 | Charts are SVG without external charting dependencies: responsive `viewBox="0 0 800 380"`, cubic Bezier curves, dual y-axes (Mbps and ms), SLA threshold lines and touch dots. | README.md@ad3758c:229 | R-03 | U-02…U-0n (pages: Ping, Speed) | |
| REQ-RE-0155 | The dossier is exportable in 11 languages via `/dossier?lang=...` or Telegram `/report`. | README.md@ad3758c:230 | R-03 | S-02…S-0n (API: exports), S-11, C-06 | |
| REQ-RE-0156 | Task manager reads `/proc` directly instead of spawning `ps`. | README.md@ad3758c:233 | C-09 | | |
| REQ-RE-0157 | Task manager shows PID, user, priority, nice, VIRT, RES, CPU %, MEM %, runtime and command arguments. | README.md@ad3758c:234 | C-09 | U-02…U-0n (page: Tasks) | |
| REQ-RE-0158 | Task manager filters system services, user apps and NMB daemons. | README.md@ad3758c:235 | C-09 | U-02…U-0n (page: Tasks) | |
| REQ-RE-0159 | SIGTERM or SIGKILL from the Web UI, limited to processes the service user may signal. | README.md@ad3758c:236 | C-09 | S-02…S-0n (API: system), P-04 | |
| REQ-RE-0160 | Self-healing: background health audits of binaries, database, permissions and connectivity, with repairs. | README.md@ad3758c:239 | P-01 | | |
| REQ-RE-0161 | An update is started from the Web UI ("Apply Stable Update"), `nmb-cli update --apply --yes` or `nmb-heal update --apply`. | README.md@ad3758c:240 | P-02 | S-10, U-02…U-0n (page: System), P-01 | |
| REQ-RE-0162 | Update: download the newest version, run its tests (nothing changes if they fail), back up settings and measurements, install it keeping settings and measurements and replacing the program files, and put the previous version back if the install or the web server check fails. | README.md@ad3758c:240 | P-02 | P-03, P-07 | Native: signed updates and file upload (owner decision, roadmap 6B.3). |
| REQ-RE-0163 | The Web UI shows each update step; the update log is `/var/log/nmb/update.log`. | README.md@ad3758c:240 | P-02 | U-02…U-0n (page: System), C-03 | |
| REQ-RE-0164 | An update is offered only when the source is ahead of the installed version: a box running newer code is never "updated" back to older code. | README.md@ad3758c:240 | P-02 | P-01, S-11 | |
| REQ-RE-0165 | Central collector client: everything is off until the box owner turns it on (Web UI → System → Central measurement collector, or `[collector]`), with three separate consents (`enabled`, `share_diagnostics`, `allow_remote_control`). | README.md@ad3758c:243; README.md@ad3758c:460 | S-12 | C-01, U-02…U-0n (page: System) | |
| REQ-RE-0166 | With `enabled` the collector gets: speed tests (failed ones too), an hourly ping summary per target, the box's unique ID, its name, the case number, the ISP, the customer number and the subscribed speeds. | README.md@ad3758c:247 | S-12 | S-13, C-11 | |
| REQ-RE-0167 | With `share_diagnostics` it also gets the box's events and, hourly, the full network diagnostics (interfaces, routes, MTU, DNS, NAT, Wi-Fi, traceroute, LAN devices) and the settings with every secret masked. | README.md@ad3758c:248 | S-12 | C-04, S-13 | |
| REQ-RE-0168 | With `allow_remote_control` the collector may only start a speed test (engine and Ookla server), ask for fresh diagnostics or an upload; no shell commands, no setting changes; the box asks the collector for work every minute, so nothing connects into the owner's network and no port is opened. | README.md@ad3758c:249; README.md@ad3758c:460 | S-12 | S-13, M-04 | |
| REQ-RE-0169 | Every box has a permanent random ID (prefix `nmb-` and three groups of four characters), shown in the Web UI, because the box name is `nmb-01` on every box unless changed. | README.md@ad3758c:251 | C-01 | S-12, R-02, U-02…U-0n (page: System) | Example ID not copied. No package names the ID; C-01 (first-start defaults) proposed as owner. |
| REQ-RE-0170 | Collector web pages are password-protected; a box is found by its ID, name, current or earlier case number, customer number or ISP; its page shows averages against the subscribed speeds, results per engine, ping per target, events, diagnostics history, remote commands and a CSV download. | README.md@ad3758c:251 | S-13 | | |
| REQ-RE-0171 | Collector setup: its own system user, an admin password (`set-password`), a token per client (`token-add <name>`), a data directory owned by the collector user, and a service unit. | README.md@ad3758c:254-261 | S-13 | P-11 | Clone URL names the repository account (source names a specific site value; not copied). |
| REQ-RE-0172 | The collector listens on port 8600; it is put behind an HTTPS reverse proxy or started with `--cert` and `--key`, since tokens and the admin password travel with every request. | README.md@ad3758c:262 | S-13 | | |
| REQ-RE-0173 | Router plan target: commercial routers with as little as 4 MB flash and 32 MB RAM. | README.md@ad3758c:267 | NEW: X2-01 Step-2 router tiers | | Design target. |
| REQ-RE-0174 | Generic first: the native code is platform- and architecture-independent and is built, tested and proven on a PC (x86-64) and the Raspberry Pi 4 (ARM64) before any router work; router targets are added later as backends of the same code. | README.md@ad3758c:268 | P-13 | F-00, NEW: X2-01 Step-2 router tiers | |
| REQ-RE-0175 | Quality bar: the native build runs at least as flawlessly as the Python version, fixes the bugs found on the way instead of carrying them over, and is measurably better (CPU, memory, start-up, logging, safe updates); a component replaces its Python counterpart only with that evidence written down. | README.md@ad3758c:269 | P-13 | F-03, F-05 | |
| REQ-RE-0176 | Single binary `nmbd`, written in Go 1.23+ (Tier 1) and C11 with musl (Tiers 2 and 3), with no Python, pip or runtime packages. | README.md@ad3758c:270 | P-10 | F-00, NEW: X2-01 Step-2 router tiers | Work plan F-00: the PC has Go 1.22; the toolchain version is fixed there. |
| REQ-RE-0177 | Router target ASUS RT-AXE7800 (BCM6756, 4 x Cortex-A7 1.7 GHz, 512 MB RAM, 256 MB NAND, 2.5GbE): Asuswrt/Merlin or OpenWrt `bcm4908`; Go or C11; 2.5 Gbps line-rate testing, multi-threaded CAKE, Flow Cache bypass, in-process `tsnet` Tailscale. | README.md@ad3758c:272-275 | NEW: X2-01 Step-2 router tiers | I-10 | Product port target, not site data. Design targets. |
| REQ-RE-0178 | Router target ASUS RT-N12+ (MT7620N MIPS 580 MHz, 32 MB RAM, 8 MB flash, 100 Mbps): OpenWrt `ramips/mt7620`; C11 musl under 850 KB, about 2.8 MB RSS; 100 Mbps testing under 14 % CPU. | README.md@ad3758c:276-279 | NEW: X2-01 Step-2 router tiers | I-10 | Design targets. |
| REQ-RE-0179 | Router target TP-Link TL-WR740N v4 (AR9331 MIPS 400 MHz, 32 MB RAM, 4 MB flash, 100 Mbps): stripped OpenWrt `ath79/generic`; micro-C11 `-Os`, UPX (178 KB), 1.4 MB RSS, RAM-only ring buffer with no flash writes. | README.md@ad3758c:280-283 | NEW: X2-01 Step-2 router tiers | I-10 | Design targets. |
| REQ-RE-0180 | Router target TP-Link TL-WR1043ND v1 (AR9132 400 MHz, 32 MiB RAM, 8 MiB flash, RTL8366RB, 1 + 4 x 1GbE). | README.md@ad3758c:284-285 | NEW: X2-01 Step-2 router tiers | I-10 | First router target (porting plan, work plan I-10). |
| REQ-RE-0181 | Router target TP-Link Archer C6 v2 (QCA9563 775 MHz, 128 MiB RAM, 8 MiB flash, QCA8337N, 1 + 4 x 1GbE, AC1200). | README.md@ad3758c:286-287 | NEW: X2-01 Step-2 router tiers | I-10 | |
| REQ-RE-0182 | Router target ASUS RT-AC57U v1 only (MT7621AT, 128 MiB RAM, 16 MiB flash, MT7603EN + MT7612EN, 1 x USB 2.0); likely base OpenWrt `ramips/mt7621`, to be confirmed; v2 is different hardware. | README.md@ad3758c:288-290 | NEW: X2-01 Step-2 router tiers | I-10 | |
| REQ-RE-0183 | Planned device families: MikroTik (native `nmbd` as a RouterOS v7 container, or OpenWrt instead of RouterOS) and MediaTek-SoC routers (MT7620/MT7628, MT7621, MT7622, MT7981B, MT7986A, MT7988A), one port per chip; figures to verify. | README.md@ad3758c:291 | NEW: X2-01 Step-2 router tiers | | |
| REQ-RE-0184 | Every device keeps an owner-enabled maintenance mode with full root access for developers and developer LLM agents: SSH keys only, every session in an audit log, switched off after a set time. | README.md@ad3758c:292 | P-05 | P-06, N-02 | |
| REQ-RE-0185 | Every device has detailed structured logs that do not wear out flash. | README.md@ad3758c:292 | C-03 | | |
| REQ-RE-0186 | Every device keeps crash evidence that survives a reboot. | README.md@ad3758c:292 | C-03 | P-01 | |
| REQ-RE-0187 | Fixes are applied next to the running version with automatic rollback, so a field fix is unlikely to brick the device. | README.md@ad3758c:292 | P-02 | P-01 | |
| REQ-RE-0188 | NMB's own distribution (planned): a minimal image for the Raspberry Pi 4 and later x86-64, built on `nmbd`, with a read-only system, a separate data partition and A/B updates with automatic rollback. | README.md@ad3758c:293 | P-11 | P-02 | |
| REQ-RE-0189 | The base of the distribution (Debian, OpenWrt, Alpine or another network-oriented one) is chosen by a written comparison. | README.md@ad3758c:293 | I-09 | P-11 | The owner decides (work plan I-09). |
| REQ-RE-0190 | The distribution uses the stock kernel of its base unchanged (the Raspberry Pi kernel on the Pi 4, the base's LTS kernel on x86-64). | README.md@ad3758c:293 | P-11 | I-05, I-06 | |
| REQ-RE-0191 | No device runs Python. | README.md@ad3758c:293 | P-13 | P-10 | |
| REQ-RE-0192 | Router DNS sinkhole: a 512 KB counting Bloom filter (MurmurHash3, 0.1 % false positives) and a trie, under 640 KB RAM on 32 MB routers. | README.md@ad3758c:295 | NEW: X2-01 Step-2 router tiers | N-12 | Design target. |
| REQ-RE-0193 | Speed engines are in-process TCP/TLS clients towards Ookla, Cloudflare, Fast.com and RTR, without external binaries (zero-copy buffers). | README.md@ad3758c:296 | M-04 | M-06, M-07, M-08, M-10, NEW: X2-01 Step-2 router tiers | In-process Ookla conflicts with the official-binary engine (REQ-RE-0072); owner decision: Ookla "whichever works". `roadmap-1-507.md` proposes `NEW: M-17` shared HTTP transfer core, which would own this row. |
| REQ-RE-0194 | Self-healing is a multi-tier watchdog: service restart, link bounce, driver offload toggling, configuration rollback. | README.md@ad3758c:297 | P-01 | C-01, N-17 | |
| REQ-RE-0195 | Dynamic DNS client for `duckdns` and `dynv6`, TLS 1.3, aware of address changes from netlink, under 60 KB RAM. | README.md@ad3758c:298 | N-16 | NEW: X2-01 Step-2 router tiers | RAM figure is a router design target. |
| REQ-RE-0196 | The README has a tested-status section that says what has actually been run, and links the full testing record. | README.md@ad3758c:307 | P-12 | P-13 | |
| REQ-RE-0197 | Every speed, CPU and timing figure names the line and hardware it came from; untested line classes (much slower, much faster) are stated, and so is the Pi 4's limit above about 1 Gbps (one gigabit port and its CPU). | README.md@ad3758c:309 | P-13 | P-12 | The test line's ISP, country and router model are not copied (source names a specific site value; not copied). |
| REQ-RE-0198 | The native build is in daily use on a Raspberry Pi 4 with Debian 13 in `standalone` mode (the Python status) before it replaces the Python version. | README.md@ad3758c:313 | P-13 | | |
| REQ-RE-0199 | Router mode on a real PPPoE line on the Pi 4 is proven: automatic mode chooses it, kernel PPPoE at about 900/320 Mbit/s, redial, a real phone on the box's Wi-Fi, the WAN closed when seen from another network, the box's IPv6 and IPv6 for a LAN client, port forwarding and UPnP reached from another network. | README.md@ad3758c:314 | P-13 | N-04, N-05, N-07, N-02, N-11 | Python figures from the test line. |
| REQ-RE-0200 | Clean install on x86-64 is tested on the supported systems (Python: Debian 13 and 12, Ubuntu 26.04, 24.04 and 22.04 LTS, Fedora 45, AlmaLinux 10, Rocky Linux 10, Alpine 3.22 and 3.20): install, services, Web UI, a real speed test, reinstall and reboot. | README.md@ad3758c:315; README.md@ad3758c:339 | P-11 | P-13, F-06 | |
| REQ-RE-0201 | `mode_a` router and `mode_b` bridge are tested with DHCP, NAT, DMZ isolation and client traffic (Python: in virtual machines only, not on Raspberry Pi hardware). | README.md@ad3758c:316 | N-03 | F-06, P-13 | Owner decision: everything is tested on the Pi 4 too. |
| REQ-RE-0202 | Wi-Fi access point is proven: WPA2 on 2.4 and 5 GHz and WPA3 on 2.4 GHz with simulated radios and a real client; on the Pi 4 radio router mode on 2.4 GHz 802.11n 20 MHz and on 5 GHz 802.11ac 80 MHz with real clients over IPv4 and IPv6; one radio serves one band at a time. | README.md@ad3758c:317 | N-07 | F-06, P-13 | Python: about 85-95 Mbit/s on 5 GHz, the limit of the Pi 4 radio as an access point. |
| REQ-RE-0203 | Wi-Fi 7, 6 GHz and 160/320 MHz are not claimed until verified; when the access point daemon rejects Wi-Fi 7 settings the box steps down to an older standard. | README.md@ad3758c:318 | N-07 | | |
| REQ-RE-0204 | PPPoE is proven on a real line and against a simulated ISP: dial-in, the ISP's refusal shown, redial, router mode on it. | README.md@ad3758c:319 | N-05 | F-06 | ISP name not copied (source names a specific site value; not copied). |
| REQ-RE-0205 | VLAN tagging is tested on a real line before it is called working (Python: code exists, not tested). | README.md@ad3758c:320 | N-05 | P-13 | |
| REQ-RE-0206 | Multi-WAN load balancing is tested on real lines before it is called working (Python: code exists, not tested). | README.md@ad3758c:320 | N-03 | P-13 | |
| REQ-RE-0207 | Link aggregation is tested with a real switch or router before it is called working (Python: code exists, not tested). | README.md@ad3758c:320 | N-06 | P-13 | |
| REQ-RE-0208 | The supported install bases are stated; others (such as Raspberry Pi OS, Arch) and 32-bit systems are listed as not tested (Python setup supports Debian-based systems, Fedora, RHEL and its rebuilds, Alpine). | README.md@ad3758c:321 | P-11 | P-12 | Work plan build matrix decides 32-bit ARM for the native build. |
| REQ-RE-0209 | On a 512 MB RAM device: 25 minutes of tests every minute with no out-of-memory kill and no swap used (Python: in a 512 MB VM). | README.md@ad3758c:322 | C-08 | P-13 | README.hu.md:98 and :268 say "not tested on 512 MB": stale. |
| REQ-RE-0210 | Reinstall and upgrade keep settings and measurements, clean up removed files, and bring the services back after a reboot. | README.md@ad3758c:323 | P-11 | P-02 | |
| REQ-RE-0211 | Bot commands that change the box are run live on the Pi 4 in router mode on PPPoE through the bot's command handler; destructive ones (`/purge`) only on copies of the data or in an isolated mount namespace, so nothing real is deleted; `/heal apply` only when nothing newer exists. | README.md@ad3758c:324 | S-11 | P-13, P-07 | Python list: `/mode`, `/vpn start`/`stop`, `/heal`, `/heal fix`, `/heal updates`, `/wifi connect`, `/dyndns on`/`off`/`update`, `/adblock on`/`off`, `/backup create`, `/kill`, `/reboot now`. |
| REQ-RE-0212 | The update is run end to end against the real update source (Python: unit tests only for download, safe unpacking, test gate and status). | README.md@ad3758c:325 | P-02 | P-13 | |
| REQ-RE-0213 | Collector and remote control are tested: upload, search by case number, CSV, diagnostics only with consent, remote speed test, refusal without consent, cross-site requests refused; also between two machines (Python: one machine only). | README.md@ad3758c:326 | S-13 | S-12, P-13 | |
| REQ-RE-0214 | WireGuard and OpenVPN servers are tested with a real client: tunnel, internet and Web UI through the tunnel. | README.md@ad3758c:327 | N-13 | N-14, F-06 | |
| REQ-RE-0215 | The HTTPS (CONNECT) proxy is tested with a real client. | README.md@ad3758c:327 | N-11 | F-06 | |
| REQ-RE-0216 | The guest Wi-Fi client is tested against a simulated network (Python: needs NetworkManager). | README.md@ad3758c:328 | N-08 | F-06 | |
| REQ-RE-0217 | Self-update with rollback is tested: a good update is applied, a broken one rolled back. | README.md@ad3758c:329 | P-02 | F-06 | |
| REQ-RE-0218 | DuckDNS is tested beyond a wrong token (Python: only a rejected wrong token). | README.md@ad3758c:330 | N-16 | P-13 | |
| REQ-RE-0219 | Telegram is tested with a real account before it is called working. | README.md@ad3758c:331 | S-11 | P-13 | |
| REQ-RE-0220 | SMTP mail is tested with a real account before it is called working. | README.md@ad3758c:331 | R-05 | P-13 | |
| REQ-RE-0221 | Tailscale is tested with a real account before it is called working. | README.md@ad3758c:331 | N-15 | P-13 | |
| REQ-RE-0222 | OONI is tested with a real run before it is called working. | README.md@ad3758c:331 | M-13 | P-13 | |
| REQ-RE-0223 | NMB states that it is a personal project, not a certified measuring instrument: its results can support a complaint but do not replace a regulator's or ISP's own measurement. | README.md@ad3758c:333 | P-08 | R-03, P-12 | |
| REQ-RE-0224 | Installing on a clean supported system is one documented command run as root (Python: clone the repository and run `sudo ./setup.sh`). | README.md@ad3758c:339-345 | P-11 | P-12 | Clone URL not copied (source names a specific site value; not copied). |
| REQ-RE-0225 | The installer installs the network tools the build needs (Python: `iproute2`, `nftables`, `iptables`, `ethtool`, `pciutils`, `usbutils`, `dnsmasq`, `hostapd`, `ppp`, `speedtest`). | README.md@ad3758c:348 | P-11 | | The native build may need fewer of them. |
| REQ-RE-0226 | The installer sets up the tmpfs mounts and the flash longevity settings for microSD cards. | README.md@ad3758c:349 | P-11 | C-02 | |
| REQ-RE-0227 | The installer applies kernel settings: socket buffers for 10G bursts, BBR congestion control, unprivileged ICMP. | README.md@ad3758c:350 | P-11 | P-06, M-01 | |
| REQ-RE-0228 | The application goes to `/opt/nmb` and its commands to `/usr/local/bin/`. | README.md@ad3758c:351 | P-11 | P-10 | |
| REQ-RE-0229 | The installer enables and starts the service target (`nmb.target`). | README.md@ad3758c:352 | P-10 | P-11 | |
| REQ-RE-0230 | After install the dashboard opens at `http://<appliance-ip>/` without a port. | README.md@ad3758c:354-357 | S-01 | | |
| REQ-RE-0231 | Default login is `admin` / `admin` unless installed with `RANDOM_ADMIN_PASSWORD=1` (then saved to `/var/lib/nmb/credentials.txt`); it is changed in the Settings tab. | README.md@ad3758c:358 | P-03 | S-01, P-11 | The owner keeps `admin`/`admin` on the Python box on purpose. Work plan P-03 names a unique password per device for the native build; the register owner reconciles this row with the roadmap row. |
| REQ-RE-0232 | The installer warns on every run while the factory login is set; a password the owner set is never printed and never written to `credentials.txt`. | README.md@ad3758c:358 | P-11 | P-03, C-04 | |
| REQ-RE-0233 | `nmb-netd --status` and `nmb-netd --switch auto\|standalone\|mode_a\|mode_b`. | README.md@ad3758c:366 | P-10 | N-03 | Native: multi-call name of `nmbd`. |
| REQ-RE-0234 | `nmb-probe --verbose`. | README.md@ad3758c:367 | P-10 | M-01 | |
| REQ-RE-0235 | `nmb-speed --trigger manual`, `--engine ookla\|cloudflare\|fastcom\|rtr\|iperf3\|all`, `--interface <name>`. | README.md@ad3758c:368 | P-10 | M-04 | |
| REQ-RE-0236 | `nmb-ooni status` and `nmb-ooni run websites`. | README.md@ad3758c:369 | P-10 | M-13 | |
| REQ-RE-0237 | `nmb-report --window 24h --html <file>`, `--json-summary`, `--send-email`. | README.md@ad3758c:370 | P-10 | R-01, R-03, R-05 | |
| REQ-RE-0238 | `nmb-adblock status`, `update`, `-b <domain>`. | README.md@ad3758c:371 | P-10 | N-12 | |
| REQ-RE-0239 | `nmb-heal audit`, `status`, `fix`. | README.md@ad3758c:372 | P-10 | P-01 | |
| REQ-RE-0240 | `nmb-driver scan`, `missing`, and `install` (as root). | README.md@ad3758c:373 | P-10 | N-17, P-06 | |
| REQ-RE-0241 | `nmb-telegram` runs the bot. | README.md@ad3758c:374 | P-10 | S-11 | |
| REQ-RE-0242 | `nmb-cli`: the Web UI menus in the terminal (over SSH, mouse, hotkeys) and JSON commands, for example `menu`, `show tab-home`, `run speed.run engine=...`, `logs dev -n 50 --level warning`, `latest [engine]`, `update [--apply --yes]`, `run collector.send`. | README.md@ad3758c:375 | S-10 | | Owner decision for the native build: plain CLI (6B.2). |
| REQ-RE-0243 | `nmb-collector set-password`, `token-add <name>`, `serve --port 8600`; it runs on the collector operator's server, not on the box. | README.md@ad3758c:376 | S-13 | P-10 | |
| REQ-RE-0244 | `nmb-web` serves the Web UI and API (port 80, development 8080). | README.md@ad3758c:377 | P-10 | S-01 | |
| REQ-RE-0245 | Privileged helper: `nmb-helper install-drivers all` and `nmb-helper tune` (as root). | README.md@ad3758c:378 | P-06 | N-17, P-11 | |
| REQ-RE-0246 | Bot access: `[telegram] allowed_chat_ids` lists Telegram ids; a private chat works when it is a listed user's own chat; in a group, the group's id and the user id of every member who may give commands must be listed; other members are ignored. | README.md@ad3758c:386 | S-11 | C-01 | |
| REQ-RE-0247 | The bot does not run edited messages, nor commands older than 2 minutes (queued while the box was off). | README.md@ad3758c:386 | S-11 | | |
| REQ-RE-0248 | `/menu` opens the main touch menu (inline keyboards). | README.md@ad3758c:391 | S-11 | | |
| REQ-RE-0249 | `/status`: uptime, WAN mode, public IP, ISP name, CPU/RAM load, data the tests used in the data cap period. | README.md@ad3758c:392 | S-11 | M-05, C-09, M-12 | |
| REQ-RE-0250 | `/ping`: latency, jitter and packet loss to the reference targets. | README.md@ad3758c:393 | S-11 | M-01 | |
| REQ-RE-0251 | `/test [engine]` runs a speed test (`ookla`, `cloudflare`, `fastcom`, `rtr`, `iperf`) with bufferbloat rating. | README.md@ad3758c:394 | S-11 | M-04 | |
| REQ-RE-0252 | `/report [window]` makes and sends the SLA dossier (`.html`) into the chat. | README.md@ad3758c:395 | S-11 | R-03 | |
| REQ-RE-0253 | `/sla` explains what the SLA result means. | README.md@ad3758c:396 | S-11 | R-01 | |
| REQ-RE-0254 | `/ai` sends the masked Markdown diagnostic snapshot for outside LLM analysis. | README.md@ad3758c:397 | S-11 | R-04, C-04 | |
| REQ-RE-0255 | `/mode [auto\|standalone\|mode_a\|mode_b]` shows the running mode and the override, or sets the override. | README.md@ad3758c:398 | S-11 | N-03 | |
| REQ-RE-0256 | Setting a mode or override that leads to the mode already running rebuilds nothing (no PPPoE redial). | README.md@ad3758c:398 | N-03 | S-11, N-05 | |
| REQ-RE-0257 | `/vpn`: active WireGuard and OpenVPN tunnels and connected client counts. | README.md@ad3758c:399 | S-11 | N-13, N-14 | |
| REQ-RE-0258 | `/vpn clients`: connected peers with handshake times, endpoints and RX/TX traffic. | README.md@ad3758c:400 | S-11 | N-13, N-14 | |
| REQ-RE-0259 | `/vpn qr <name>` sends a camera-scannable QR code PNG. | README.md@ad3758c:401 | S-11 | N-13 | |
| REQ-RE-0260 | `/vpn conf <name>` sends the `.conf` profile for WireGuard or OpenVPN. | README.md@ad3758c:402 | S-11 | N-13, N-14 | |
| REQ-RE-0261 | `/vpn add <name> [wg\|ovpn]` creates a client profile and returns the QR code. | README.md@ad3758c:403 | S-11 | N-13, N-14 | |
| REQ-RE-0262 | `/vpn del <name>` deletes a client profile. | README.md@ad3758c:404 | S-11 | N-13, N-14 | |
| REQ-RE-0263 | `/vpn start\|stop [wg\|ovpn]` starts or stops the WireGuard or OpenVPN server. | README.md@ad3758c:405 | S-11 | N-13, N-14 | |
| REQ-RE-0264 | `/adblock`: total queries today, blocked queries, blocklist domain count. | README.md@ad3758c:406 | S-11 | N-12 | |
| REQ-RE-0265 | `/adblock on\|off` enables or disables the DNS blocker. | README.md@ad3758c:407 | S-11 | N-12 | |
| REQ-RE-0266 | `/dyndns`: status of the DynDNS (DuckDNS or dynv6) updater: last update, address, next run. | README.md@ad3758c:408 | S-11 | N-16 | Python runs the updater inside the web server. |
| REQ-RE-0267 | `/dyndns on\|off\|update` turns the updater on or off (it follows within 30 s) or updates the name now with the configured provider. | README.md@ad3758c:409 | S-11 | N-16 | |
| REQ-RE-0268 | `/dyndns interval <min>\|domain <name>\|token <token>` changes the interval, the name or the token. | README.md@ad3758c:410 | S-11 | N-16, C-04 | |
| REQ-RE-0269 | `/wifi [scan\|saved]`: Wi-Fi status, nearby networks, saved profiles. | README.md@ad3758c:411 | S-11 | N-07 | |
| REQ-RE-0270 | `/wifi connect <ssid> [password]` joins a network as a client; in router mode, where the radio is the access point, the bot says so instead. | README.md@ad3758c:412 | S-11 | N-07, N-08 | |
| REQ-RE-0271 | `/tailscale` shows the box's Tailscale status. | README.md@ad3758c:413 | S-11 | N-15 | |
| REQ-RE-0272 | `/heal`: health audit of services, database, permissions, missing programs. | README.md@ad3758c:414 | S-11 | P-01 | |
| REQ-RE-0273 | `/heal fix [tier]` repairs what the audit found (service restarts first; higher tiers go further). | README.md@ad3758c:415 | S-11 | P-01 | |
| REQ-RE-0274 | `/heal updates` checks for an update and offers one only when the source is newer than the installed version. | README.md@ad3758c:416 | S-11 | P-02 | |
| REQ-RE-0275 | `/heal apply` installs that update with a backup first and rollback on failure; it does nothing when nothing newer exists. | README.md@ad3758c:417 | S-11 | P-02 | |
| REQ-RE-0276 | `/backup [list]` lists the stored backups. | README.md@ad3758c:418 | S-11 | P-03 | |
| REQ-RE-0277 | `/backup create` makes a full backup (settings, database, VPN keys). | README.md@ad3758c:419 | S-11 | P-03 | |
| REQ-RE-0278 | `/backup send <file> [plain]` sends a backup encrypted with the backup passphrase set on the Web UI's Telegram card (a `.gpg` file that `gpg -d` or Kleopatra opens); without a passphrase it sends nothing; `plain` sends it unencrypted; Telegram's 50 MB bot limit applies. | README.md@ad3758c:420 | S-11 | P-03, C-04 | Security review in S-11 (QB 7): `plain` sends secrets in clear. |
| REQ-RE-0279 | `/tasks`: load, memory and the busiest processes. | README.md@ad3758c:421 | S-11 | C-09 | |
| REQ-RE-0280 | `/kill <pid> [15\|9\|1\|2]` sends TERM, KILL, HUP or INT to a process the bot may signal; PID 1 and the bot itself are protected; other signals are refused. | README.md@ad3758c:422 | S-11 | C-09, P-04 | |
| REQ-RE-0281 | `/drivers`: network adapters and their drivers. | README.md@ad3758c:423 | S-11 | N-17 | |
| REQ-RE-0282 | `/purge [category] confirm` shows or deletes stored data: `measurements`, `events`, `reports`, `backups`, `logs` (NMB logs, logs kept on the SD card, system journal, kernel buffer) or `all`; settings, passwords and VPN keys are kept. | README.md@ad3758c:424 | S-11 | C-08 | |
| REQ-RE-0283 | `/reboot now` restarts the box (about 1-2 minutes; in router mode the home network is offline meanwhile). | README.md@ad3758c:425 | S-11 | P-06 | |
| REQ-RE-0284 | `/lang [code]` switches the bot language: `hu`, `en`, `de`, `fr`, `uk`, `ro`, `it`, `es`, `pt`, `pt-br`, `zh`. | README.md@ad3758c:426 | S-11 | C-06 | |
| REQ-RE-0285 | The whole suite is grouped under one service target (`nmb.target`), each component with its own unit; status and restart of the whole suite go through the target. | README.md@ad3758c:432-439; README.hu.md@ad3758c:554 | P-10 | P-04 | Native: one `nmbd` binary; units per role are P-10's design. |
| REQ-RE-0286 | Each service's log can be followed live with the system journal (for example the probe, and the scheduled speed tests under the service that runs the schedule). | README.md@ad3758c:440-446 | C-03 | P-10 | |
| REQ-RE-0287 | The web server runs as the unprivileged `nmb` user with only the network capabilities it needs; polkit lets it manage its own services and reboot; the privileged helper does every other privileged task, so the CLI terminal and the task manager cannot act as root. | README.md@ad3758c:452 | P-04 | P-06 | |
| REQ-RE-0288 | Known Python limit: on polkit 0.105 (Ubuntu 22.04) the rule cannot name units, so there the service user may manage any systemd service. | README.md@ad3758c:452 | P-04 | F-03 | See Questions. |
| REQ-RE-0289 | The public-address guard blocks ports 80, 53 and 5201 from public IPv4 and IPv6 addresses; local networks, VPN and Tailscale still work, and the box's own outgoing tests are not affected. | README.md@ad3758c:453 | N-02 | | |
| REQ-RE-0290 | Known Python limit: the guard needs systemd and `nftables`, so it is not installed on Alpine. | README.md@ad3758c:453 | N-02 | P-11, F-03 | See Questions. |
| REQ-RE-0291 | The boot partition copy (`/boot/firmware/nmb.conf`) shows `__KEEP_CURRENT__` instead of passwords and tokens: leaving the mark keeps the value, typing a new value changes it; `nmb_credentials.txt` stays there only while the password is the default or was generated at install. | README.md@ad3758c:454 | C-01 | C-04, P-03, F-04 | |
| REQ-RE-0292 | Factory reset needs `--yes` in the CLI and `RESET` in the API. | README.md@ad3758c:455 | P-03 | S-10, S-02…S-0n (API: system) | |
| REQ-RE-0293 | Data and log purges need `PURGE`. | README.md@ad3758c:455 | C-08 | S-02…S-0n (API: system), S-11 | |
| REQ-RE-0294 | The Ookla archive is installed only if its SHA-256 matches. | README.md@ad3758c:456 | M-10 | P-11 | |
| REQ-RE-0295 | CSV exports (Web UI buttons, `nmb-cli` or the URLs): `/export/csv` (ping per minute and target, with the target's role), `/export/speedtests.csv` (every speed test field), `/export/speedtests_detailed.csv` (short spreadsheet table, `;` and UTF-8 BOM), `/export/outages.csv` (with `cause_kind` `isp` or `own` and the counted seconds) and `/export/events.csv` (event log, password only). | README.md@ad3758c:457 | R-02 | S-02…S-0n (API: exports), U-02…U-0n (page: Exports), S-10, S-01 | |
| REQ-RE-0296 | `speedtests.csv` carries, after `error`, the columns row id, time zone, run start, duration, interface and local address, retransmits, TCP window, box ID, NMB version and network mode. | README.md@ad3758c:457 | R-02 | F-04 | |
| REQ-RE-0297 | Each export takes `?from=YYYY-MM-DD&to=YYYY-MM-DD`, an ISO time or epoch ms, or `?days=N` / `?hours=N`; times are UTC. | README.md@ad3758c:457 | R-02 | S-02…S-0n (API: exports) | |
| REQ-RE-0298 | A value that was not measured is empty (or `--`), never 0. | README.md@ad3758c:457 | R-02 | R-01, R-03, U-02…U-0n, S-11 | |
| REQ-RE-0299 | Every speed test, outage and ping minute records how the box was connected then: its mode (standalone, router, bridge), the link the test used (Ethernet, PPPoE, Wi-Fi, USB tethering, bridge, LAN passthrough, link aggregation), and whether it was directly on the ISP (own PPPoE session, or a public or carrier-grade NAT address on that link) or behind another router (a private address). | README.md@ad3758c:458 | C-11 | M-01, M-02, M-04, N-01, N-03, R-01, R-02, R-03, R-04, R-05, S-02…S-0n, U-02…U-0n, S-10, S-11, S-12, S-13 | Roadmap 4.29; "Also in" as the work plan lists. |
| REQ-RE-0300 | The position (direct or behind a router) comes from the address, not from the mode; what cannot be told for sure is shown as unknown. | README.md@ad3758c:458 | C-11 | M-12 | |
| REQ-RE-0301 | The dossier lists the setups of the period; the CSVs carry them as columns (`box_mode`, `uplink_kind`, `isp_path`, `setup_source`); the Web UI, the bot and the collector show them. | README.md@ad3758c:458 | C-11 | R-02, R-03, U-02…U-0n, S-02…S-0n, S-10, S-11, S-12, S-13 | |
| REQ-RE-0302 | Rows from before the setup was recorded are reconstructed once from the event log and the stored interface and address, and are marked `derived`. | README.md@ad3758c:458 | C-11 | C-02 | Native: the migration from a Python database (C-02) keeps the `derived` marks. |
| REQ-RE-0303 | Results without password (optional, off by default): with `public_results = true` in `[web]`, devices on the local network or VPN can open `/results` and download the report and CSV files without login; only these read-only pages open; settings, tests and the diagnostic bundle still need the password, and requests from public addresses always do. | README.md@ad3758c:459 | S-01 | C-01, U-02…U-0n (page: Results), R-02, R-03 | `/export/events.csv` stays password-only (REQ-RE-0295). |
| REQ-RE-0304 | The collector stores tokens and its admin password only as hashes. | README.md@ad3758c:460 | S-13 | | |
| REQ-RE-0305 | The collector refuses admin forms from other sites. | README.md@ad3758c:460 | S-13 | | |
| REQ-RE-0306 | The collector token is masked in the boot partition copy. | README.md@ad3758c:460 | C-01 | C-04, S-12 | |
| REQ-RE-0307 | When the OONI package brings Tor, its service is disabled, because NMB does not use it. | README.md@ad3758c:461 | M-13 | P-11 | |
| REQ-RE-0308 | The Web UI uses HTTP by default, and the user is told to use it only on a trusted network or through a VPN. | README.md@ad3758c:462 | S-01 | P-12 | |
| REQ-RE-0309 | Changes are listed in `CHANGELOG.md`, linked from the README. | README.md@ad3758c:468 | F-00 | P-12 | Owner rule: update the changelog with each user-visible change. |
| REQ-RE-0310 | The support section lists donation addresses for Solana, Bitcoin (Native SegWit), Bitcoin (Taproot) and EVM networks (Ethereum, BNB, Polygon, ERC-20). | README.md@ad3758c:476-481 | P-08 | | Addresses not copied (source names a specific site value; not copied). See Questions. |
| REQ-RE-0311 | The same donation addresses are shown in the Web UI (Support), by the bot (`/donate`) and by `nmb-cli donate`. | README.md@ad3758c:483 | P-08 | U-02…U-0n (page: Support), S-11, S-10 | `/donate` is missing from the bot command table (lines 389-426). |
| REQ-RE-0312 | HU only: the dossier is made with one click as a standalone HTML report with embedded SVG charts and an SLA compliance table against the subscriber contract, ready to hand to the ISP or, in a dispute, to the communications authority. | README.hu.md@ad3758c:60 | R-03 | R-01 | |
| REQ-RE-0313 | HU only: the device explorer scans the local subnet with ARP continuously (also the guest network) and lists the active clients with IP, MAC, host name, first and last seen; the vendor comes from the MAC OUI. | README.hu.md@ad3758c:100; README.hu.md@ad3758c:349-352 | N-17 | U-02…U-0n (page: Devices), N-08 | |
| REQ-RE-0314 | HU only: the guest Wi-Fi probe joins the router's guest SSID as a separate client with a second Wi-Fi adapter, checks its availability, signal strength, address assignment, gateway, bandwidth and stability, and checks that it is isolated (does not see the internal network), independent of the main measurements. | README.hu.md@ad3758c:101; README.hu.md@ad3758c:354-358 | N-08 | M-14, N-07 | |
| REQ-RE-0315 | HU only: mode comparison: Mode A has a gateway address (`192.168.100.1` or custom), built-in PPPoE and VLAN, DHCP and DNS service; Mode B passes PPPoE and VLAN through untouched and serves no DHCP or DNS (the existing router hands out addresses). | README.hu.md@ad3758c:160-168 | N-03 | N-05 | README.hu.md:165 says Mode B has no address, against :168, :74 and README.md:91 (the box requests one); the address rows win. |
| REQ-RE-0316 | HU only: the port role matrix has a "not assigned / inactive" role for unused ports. | README.hu.md@ad3758c:184 | N-01 | U-02…U-0n (page: Ports & Drivers) | |
| REQ-RE-0317 | HU only: the driver finder knows more hardware (Realtek RTL8152/8153 USB, RTL8125 2.5G, RTL8126 5G; Intel igc 2.5G, ixgbe 10G, AX200/AX210; MediaTek MT7921/MT7922; Aquantia/Marvell 10G; ASIX AX88179), alerts at once when a network card failed to start for missing firmware, and installs the firmware packages (Debian names: `firmware-realtek`, `firmware-linux-nonfree`, `firmware-iwlwifi`, `firmware-atheros`, plus `pciutils`, `usbutils`, `ethtool`) and then triggers udev to load the modules. | README.hu.md@ad3758c:188-190 | N-17 | P-06, S-11 | README.hu.md:187 links a file path under a personal home folder (source names a specific site value; not copied). |
| REQ-RE-0318 | HU only: the Ookla engine measures download, upload and latency under load (bufferbloat). | README.hu.md@ad3758c:211 | M-10 | R-01 | |
| REQ-RE-0319 | HU only: Fast.com uses the API `api.fast.com/netflix/speedtest/v2` and reads the token from the Fast.com page script, falling back to the known token when that fails. | README.hu.md@ad3758c:218 | M-07 | | |
| REQ-RE-0320 | HU only: the built-in iPerf3 server can be started and stopped from the Web UI, to measure the LAN's maximum throughput (cable and the box's access point). | README.hu.md@ad3758c:229-230; README.hu.md@ad3758c:310 | M-09 | U-02…U-0n (page: iPerf3) | |
| REQ-RE-0321 | HU only: the memory profile is chosen with `[system] memory_profile = auto\|low\|standard\|performance` or in the Web UI. | README.hu.md@ad3758c:259-266 | C-08 | C-01, U-02…U-0n (page: Settings) | |
| REQ-RE-0322 | HU only: the `low` profile is for less than 768 MB RAM; when the hot database reaches its 48 MB limit, the oldest ping records are thinned. | README.hu.md@ad3758c:268-269 | C-02 | C-08 | |
| REQ-RE-0323 | HU only: `standard` profile: 96 MB live database, 8 MB cache, 16 MB mmap; `performance` profile: 192 MB live database, 32 MB cache, 64 MB mmap, 64 MB TCP buffers. | README.hu.md@ad3758c:274-277 | C-08 | C-02, P-11 | README.hu.md:273 (Python garbage collection every 3 minutes) is Python-only. |
| REQ-RE-0324 | HU only: the Web UI System tab shows free RAM and the active memory profile, with a one-click memory trim. | README.hu.md@ad3758c:279 | U-02…U-0n (page: System) | C-08, S-02…S-0n (API: system) | |
| REQ-RE-0325 | HU only: at boot the live database is restored from the persistent copy. | README.hu.md@ad3758c:287-291 | C-02 | P-10 | |
| REQ-RE-0326 | HU only: NMB does not switch off swap. | README.hu.md@ad3758c:296 | P-11 | | |
| REQ-RE-0327 | HU only: Dashboard page: summary of the latest measurements, SLA status, one-click test start, interface selector and a live run terminal. | README.hu.md@ad3758c:305 | U-02…U-0n (page: Dashboard) | S-02…S-0n (API: status) | |
| REQ-RE-0328 | HU only: Ping page (SVG latency graph, packet loss, p95/p99, CSV export); Speed results page (download, upload, bufferbloat score, engine); Outages page (every outage with start, end, duration and lost packets). | README.hu.md@ad3758c:306; README.hu.md@ad3758c:309; README.hu.md@ad3758c:311 | U-02…U-0n (pages: Ping, Speed, Outages) | M-02, R-02 | |
| REQ-RE-0329 | HU only: SLA and contract page: real measurements compared with the guaranteed and nominal values of the subscriber contract. | README.hu.md@ad3758c:312 | U-02…U-0n (page: SLA) | R-01 | |
| REQ-RE-0330 | HU only: traceroute with per-hop latency (MTR-like) and its own page, next to the port throttling and RTR pages. | README.hu.md@ad3758c:313-315 | M-11 | U-02…U-0n (pages: Traceroute, Port check, RTR), M-08 | |
| REQ-RE-0331 | HU only: Developer console page with the system log and the internal module log as a live stream; AI diagnostics page with a one-click copyable masked Markdown bundle. | README.hu.md@ad3758c:317-318 | U-02…U-0n (pages: Developer console, AI diagnostics) | C-03, R-04, S-01 | |
| REQ-RE-0332 | HU only: Settings pages: WAN (DHCP, PPPoE with VLAN, static IP, WAN MAC cloning), LAN (gateway address, DHCP range, lease time), Wi-Fi (SSID, password, 2.4/5/6 GHz, channel width, WPA2/WPA3), Ports & Drivers, operating mode (A, B, automatic), memory profiles and RAM maintenance, admin credentials. | README.hu.md@ad3758c:319-326 | U-02…U-0n (page: Settings) | N-03, N-05, N-07, N-01, S-02…S-0n (API: network) | WAN MAC cloning appears nowhere in the English file. |
| REQ-RE-0333 | HU only: Logout clears the browser's cached credentials. | README.hu.md@ad3758c:327 | S-01 | U-01 | |
| REQ-RE-0334 | HU only: the language is switched with one click in the top right corner of the Web UI. | README.hu.md@ad3758c:330 | U-01 | C-06 | |
| REQ-RE-0335 | HU only: Wi-Fi telemetry includes the noise floor. | README.hu.md@ad3758c:400 | M-14 | | Only where the driver reports it (README.md:190 wording for RSSI). |
| REQ-RE-0336 | HU only: chart generation takes less than 1.5 ms. | README.hu.md@ad3758c:445 | R-03 | P-13 | Python figure; measured again on the native build. |
| REQ-RE-0337 | HU only: hardware requirements are stated: tested on a Raspberry Pi 4 (Debian 13); hardware with less memory or 32-bit is not tested; recommended for full router mode: Raspberry Pi 4/5 or an x86-64 mini PC with at least 2 Ethernet ports (1 WAN, 1 LAN) and a Wi-Fi adapter. | README.hu.md@ad3758c:620-622 | P-12 | P-11 | |
| REQ-RE-0338 | HU only: after install, the owner sets the contract values (nominal and guaranteed download and upload) and the network card roles, and the box then starts measuring. | README.hu.md@ad3758c:651 | R-01 | C-01, U-02…U-0n (page: Settings), N-01 | |

## Coverage

### README.md (lines 1-487)

| Lines | Kind (heading, paragraph, bullet, table row, code block) | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-2 | heading | context: document title |
| 3-4 | paragraph | context: tagline summarising the product; its parts are rows REQ-RE-0005, -0006 |
| 5 | badge | REQ-RE-0001 |
| 6 | badge | context: development platform (status; see REQ-RE-0198) |
| 7 | badge | context: Python version badge; Python only, replaced by REQ-RE-0176, -0191 |
| 8 | badge | context: link badge to section 8 (REQ-RE-0118 to -0126) |
| 9 | badge | REQ-RE-0002 |
| 10-11 | badge | context: stale language badge (5 of the 11 languages); the languages are REQ-RE-0115 |
| 12-13 | paragraph (quote) | REQ-RE-0003 |
| 14-15 | paragraph (quote) | context: link to a separate comparison document with a national regulator's box; that document is not a source of this part |
| 16-17 | paragraph (quote) | REQ-RE-0004 |
| 18-19 | rule | context: separator |
| 20-21 | heading | context: section heading |
| 22-23 | paragraph | REQ-RE-0005 |
| 24-25 | paragraph | context: problem statement (motivation) |
| 26-27 | paragraph | REQ-RE-0006, REQ-RE-0007, REQ-RE-0008 |
| 28-29 | rule | context: separator |
| 30-31 | heading | context: section heading |
| 32-75 | code block | REQ-RE-0009, REQ-RE-0010; the rest of the diagram restates rows of the sections below: REQ-RE-0017, -0026, -0029, -0069 to -0072, -0077, -0078, -0083, -0086, -0087, -0100, -0101, -0103 to -0106, -0113, -0117, -0149, -0154, -0162, -0285. Its engine list stops after RTR (iPerf3 missing): context |
| 76-80 | rule, heading | context: separator and section heading |
| 81 | heading | context: section heading |
| 82 | bullet | REQ-RE-0011 to REQ-RE-0016 |
| 83 | bullet | context: sub-heading "Mode A" |
| 84 | bullet | REQ-RE-0017, REQ-RE-0018 |
| 85 | bullet | REQ-RE-0019 |
| 86 | bullet | REQ-RE-0020, REQ-RE-0021 |
| 87 | bullet | REQ-RE-0022 |
| 88 | bullet | REQ-RE-0023 |
| 89 | bullet | REQ-RE-0024, REQ-RE-0025 |
| 90 | bullet | context: sub-heading "Mode B" |
| 91 | bullet | REQ-RE-0026 |
| 92 | bullet | REQ-RE-0027, REQ-RE-0028 |
| 93 | bullet | REQ-RE-0029 |
| 94-95 | heading | context: section heading |
| 96-97 | bullet | REQ-RE-0030 |
| 98 | bullet | REQ-RE-0031 |
| 99 | bullet | REQ-RE-0032 |
| 100 | bullet | REQ-RE-0033 |
| 101 | bullet | REQ-RE-0034 |
| 102 | bullet | REQ-RE-0035 to REQ-RE-0039 |
| 103 | bullet | REQ-RE-0040 to REQ-RE-0044 |
| 104 | bullet | REQ-RE-0045, REQ-RE-0046 |
| 105 | bullet | REQ-RE-0047 to REQ-RE-0055 |
| 106 | bullet | REQ-RE-0056 to REQ-RE-0058 |
| 107 | bullet | REQ-RE-0059 to REQ-RE-0061 |
| 108-109 | bullet | REQ-RE-0062 |
| 110 | bullet | REQ-RE-0063 |
| 111 | bullet | REQ-RE-0064 |
| 112 | bullet | REQ-RE-0065 |
| 113-114 | heading | context: section heading |
| 115 | bullet | REQ-RE-0066 |
| 116 | bullet | REQ-RE-0067 |
| 117 | bullet | REQ-RE-0068 |
| 118-119 | heading | context: section heading |
| 120 | numbered item | REQ-RE-0069, REQ-RE-0070, REQ-RE-0071 |
| 121 | numbered item | REQ-RE-0072 |
| 122 | bullet | REQ-RE-0073, REQ-RE-0074, REQ-RE-0075 |
| 123 | bullet | REQ-RE-0073, REQ-RE-0074, REQ-RE-0076 (the test line's history is context, not copied) |
| 124 | numbered item | REQ-RE-0077 |
| 125 | numbered item | REQ-RE-0078 to REQ-RE-0082 |
| 126 | numbered item | REQ-RE-0083 |
| 127 | numbered item | REQ-RE-0084, REQ-RE-0085 |
| 128 | numbered item | REQ-RE-0086 |
| 129 | numbered item | REQ-RE-0087, REQ-RE-0088 |
| 130 | numbered item | REQ-RE-0089 to REQ-RE-0092 |
| 131-132 | numbered item | REQ-RE-0093 |
| 133 | paragraph | REQ-RE-0094 to REQ-RE-0098 |
| 134-135 | heading | context: section heading |
| 136 | paragraph | REQ-RE-0099 |
| 137 | bullet | REQ-RE-0100 |
| 138 | bullet | REQ-RE-0101 |
| 139 | bullet | REQ-RE-0102 |
| 140-141 | heading | context: section heading |
| 142 | bullet | REQ-RE-0103 |
| 143 | bullet | REQ-RE-0104 |
| 144 | bullet | REQ-RE-0105 |
| 145-146 | heading | context: section heading |
| 147 | bullet | REQ-RE-0106, REQ-RE-0107 |
| 148 | bullet | REQ-RE-0108 |
| 149 | bullet | REQ-RE-0109 to REQ-RE-0112 |
| 150 | bullet | REQ-RE-0113, REQ-RE-0114 |
| 151-161 | bullet (list) | REQ-RE-0115 |
| 162 | paragraph | REQ-RE-0116 |
| 163 | bullet | REQ-RE-0117 |
| 164-165 | heading | context: section heading |
| 166 | bullet | context: sub-heading "client generator" |
| 167 | bullet | REQ-RE-0118, REQ-RE-0119 |
| 168 | bullet | REQ-RE-0120 |
| 169 | bullet | REQ-RE-0121 |
| 170 | bullet | REQ-RE-0122 |
| 171 | bullet | context: sub-heading "Tailscale" |
| 172-173 | bullet | REQ-RE-0123 |
| 174 | bullet | REQ-RE-0124 |
| 175 | bullet | context: sub-heading "DynDNS and proxy" |
| 176 | bullet | REQ-RE-0125 |
| 177 | bullet | REQ-RE-0126 |
| 178-179 | heading | context: section heading |
| 180 | bullet | context: sub-heading "sinkhole" |
| 181 | bullet | REQ-RE-0127 |
| 182 | bullet | REQ-RE-0128 |
| 183 | bullet | context: sub-heading "custom control" |
| 184 | bullet | REQ-RE-0129 |
| 185 | bullet | REQ-RE-0130 |
| 186 | bullet | REQ-RE-0131 |
| 187-188 | heading | context: section heading |
| 189 | bullet | context: sub-heading naming the Python tool (`iw`); the values are REQ-RE-0132, -0133 |
| 190 | bullet | REQ-RE-0132 |
| 191 | bullet | REQ-RE-0133 |
| 192 | bullet | context: sub-heading "DNS audit" |
| 193 | bullet | REQ-RE-0134 |
| 194 | bullet | REQ-RE-0135 |
| 195 | bullet | REQ-RE-0136 |
| 196-200 | bullet | REQ-RE-0137 |
| 201 | bullet | REQ-RE-0138 |
| 202 | bullet | context: sub-heading "IRQ and offload" |
| 203 | bullet | REQ-RE-0139 |
| 204 | bullet | REQ-RE-0140 |
| 205 | bullet | context: sub-heading "loss patterns" |
| 206 | bullet | REQ-RE-0141 |
| 207 | bullet | REQ-RE-0142 |
| 208 | bullet | context: sub-heading "storage health" |
| 209 | bullet | REQ-RE-0143 |
| 210 | bullet | REQ-RE-0144 |
| 211 | bullet | REQ-RE-0145 |
| 212 | bullet | context: sub-heading "masking" |
| 213 | bullet | REQ-RE-0146 |
| 214-215 | heading | context: section heading |
| 216-217 | bullet | REQ-RE-0147 |
| 218 | bullet | context: stale (five bot languages); lines 150 and 426 give 11: REQ-RE-0114, -0284 |
| 219 | bullet | REQ-RE-0148 |
| 220 | bullet | REQ-RE-0149 |
| 221-222 | heading | context: section heading |
| 223 | bullet | context: sub-heading "neutral branding" |
| 224 | bullet | REQ-RE-0150 |
| 225 | bullet | REQ-RE-0151 |
| 226 | bullet | REQ-RE-0152 |
| 227 | bullet | REQ-RE-0153 |
| 228 | bullet | context: sub-heading "SVG engine" |
| 229 | bullet | REQ-RE-0154 |
| 230 | bullet | REQ-RE-0155 |
| 231-232 | heading | context: section heading |
| 233 | bullet | REQ-RE-0156 |
| 234 | bullet | REQ-RE-0157 |
| 235 | bullet | REQ-RE-0158 |
| 236 | bullet | REQ-RE-0159 |
| 237-238 | heading | context: section heading |
| 239 | bullet | REQ-RE-0160 |
| 240 | bullet | REQ-RE-0161 to REQ-RE-0164; the Git-checkout update path ("git pull", roll back to the previous commit) is context: a Python developer install with no native counterpart |
| 241-242 | heading | context: section heading |
| 243 | paragraph | REQ-RE-0165 |
| 244-246 | table header | context: table header |
| 247 | table row | REQ-RE-0166 |
| 248 | table row | REQ-RE-0167 |
| 249 | table row | REQ-RE-0168 |
| 250-251 | paragraph | REQ-RE-0169, REQ-RE-0170 |
| 252-253 | paragraph | context: Python-only prerequisite (standard library, Python 3.10); the setup itself is REQ-RE-0171 |
| 254-261 | code block | REQ-RE-0171 |
| 262 | paragraph | REQ-RE-0172 |
| 263-264 | heading | context: section heading |
| 265 | paragraph (quote) | REQ-RE-0002 |
| 266-267 | paragraph | REQ-RE-0173 |
| 268 | bullet | REQ-RE-0174 |
| 269 | bullet | REQ-RE-0175 |
| 270 | bullet | REQ-RE-0176 |
| 271 | bullet | context: sub-heading "target hardware" |
| 272-275 | bullet | REQ-RE-0177 |
| 276-279 | bullet | REQ-RE-0178 |
| 280-283 | bullet | REQ-RE-0179 |
| 284-285 | bullet | REQ-RE-0180 |
| 286-287 | bullet | REQ-RE-0181 |
| 288-290 | bullet | REQ-RE-0182 |
| 291 | bullet | REQ-RE-0183 |
| 292 | bullet | REQ-RE-0184 to REQ-RE-0187 |
| 293 | bullet | REQ-RE-0188 to REQ-RE-0191 |
| 294 | bullet | context: sub-heading "native subsystems" |
| 295 | bullet | REQ-RE-0192 |
| 296 | bullet | REQ-RE-0193 |
| 297 | bullet | REQ-RE-0194 |
| 298 | bullet | REQ-RE-0195 |
| 299-301 | bullet | context: links to the porting plan and the roadmap, which have their own register parts |
| 302-305 | rule, heading | context: separator and section heading |
| 306-307 | paragraph | REQ-RE-0196 |
| 308-309 | paragraph | REQ-RE-0197 |
| 310-312 | table header | context: table header |
| 313 | table row | REQ-RE-0198 |
| 314 | table row | REQ-RE-0199 |
| 315 | table row | REQ-RE-0200 |
| 316 | table row | REQ-RE-0201 |
| 317 | table row | REQ-RE-0202 |
| 318 | table row | REQ-RE-0203 |
| 319 | table row | REQ-RE-0204 |
| 320 | table row | REQ-RE-0205, REQ-RE-0206, REQ-RE-0207 |
| 321 | table row | REQ-RE-0208 |
| 322 | table row | REQ-RE-0209 |
| 323 | table row | REQ-RE-0210 |
| 324 | table row | REQ-RE-0211 |
| 325 | table row | REQ-RE-0212 |
| 326 | table row | REQ-RE-0213 |
| 327 | table row | REQ-RE-0214, REQ-RE-0215 |
| 328 | table row | REQ-RE-0216 |
| 329 | table row | REQ-RE-0217 |
| 330 | table row | REQ-RE-0218 |
| 331 | table row | REQ-RE-0219 to REQ-RE-0222 |
| 332-333 | paragraph | REQ-RE-0223 |
| 334-337 | rule, heading | context: separator and section heading |
| 338-339 | paragraph | REQ-RE-0224, REQ-RE-0200 |
| 340-346 | code block | REQ-RE-0224 |
| 347 | paragraph | context: lead-in to the installer list |
| 348 | numbered item | REQ-RE-0225 |
| 349 | numbered item | REQ-RE-0226 |
| 350 | numbered item | REQ-RE-0227 |
| 351 | numbered item | REQ-RE-0228 |
| 352 | numbered item | REQ-RE-0229 |
| 353-357 | paragraph, code block | REQ-RE-0230 |
| 358 | bullet | REQ-RE-0231, REQ-RE-0232 |
| 359-362 | rule, heading | context: separator and section heading |
| 363-365 | table header | context: table header |
| 366 | table row | REQ-RE-0233 |
| 367 | table row | REQ-RE-0234 |
| 368 | table row | REQ-RE-0235 |
| 369 | table row | REQ-RE-0236 |
| 370 | table row | REQ-RE-0237 |
| 371 | table row | REQ-RE-0238 |
| 372 | table row | REQ-RE-0239 |
| 373 | table row | REQ-RE-0240 |
| 374 | table row | REQ-RE-0241 |
| 375 | table row | REQ-RE-0242 |
| 376 | table row | REQ-RE-0243 |
| 377 | table row | REQ-RE-0244 |
| 378 | table row | REQ-RE-0245 |
| 379-382 | rule, heading | context: separator and section heading |
| 383-384 | paragraph | REQ-RE-0147 |
| 385-386 | paragraph | REQ-RE-0246, REQ-RE-0247 |
| 387-390 | table header | context: blank lines and table header |
| 391 | table row | REQ-RE-0248 |
| 392 | table row | REQ-RE-0249 |
| 393 | table row | REQ-RE-0250 |
| 394 | table row | REQ-RE-0251 |
| 395 | table row | REQ-RE-0252 |
| 396 | table row | REQ-RE-0253 |
| 397 | table row | REQ-RE-0254 |
| 398 | table row | REQ-RE-0255, REQ-RE-0256 |
| 399 | table row | REQ-RE-0257 |
| 400 | table row | REQ-RE-0258 |
| 401 | table row | REQ-RE-0259 |
| 402 | table row | REQ-RE-0260 |
| 403 | table row | REQ-RE-0261 |
| 404 | table row | REQ-RE-0262 |
| 405 | table row | REQ-RE-0263 |
| 406 | table row | REQ-RE-0264 |
| 407 | table row | REQ-RE-0265 |
| 408 | table row | REQ-RE-0266 |
| 409 | table row | REQ-RE-0267 |
| 410 | table row | REQ-RE-0268 |
| 411 | table row | REQ-RE-0269 |
| 412 | table row | REQ-RE-0270 |
| 413 | table row | REQ-RE-0271 |
| 414 | table row | REQ-RE-0272 |
| 415 | table row | REQ-RE-0273 |
| 416 | table row | REQ-RE-0274 |
| 417 | table row | REQ-RE-0275 |
| 418 | table row | REQ-RE-0276 |
| 419 | table row | REQ-RE-0277 |
| 420 | table row | REQ-RE-0278 |
| 421 | table row | REQ-RE-0279 |
| 422 | table row | REQ-RE-0280 |
| 423 | table row | REQ-RE-0281 |
| 424 | table row | REQ-RE-0282 |
| 425 | table row | REQ-RE-0283 |
| 426 | table row | REQ-RE-0284 |
| 427-430 | rule, heading | context: separator and section heading |
| 431-432 | paragraph | REQ-RE-0285 |
| 433-446 | code block | REQ-RE-0285, REQ-RE-0286 |
| 447-451 | rule, heading | context: separator and section heading |
| 452 | bullet | REQ-RE-0287, REQ-RE-0288 |
| 453 | bullet | REQ-RE-0289, REQ-RE-0290 |
| 454 | bullet | REQ-RE-0291 |
| 455 | bullet | REQ-RE-0292, REQ-RE-0293 |
| 456 | bullet | REQ-RE-0294 |
| 457 | bullet | REQ-RE-0295 to REQ-RE-0298 |
| 458 | bullet | REQ-RE-0299 to REQ-RE-0302 |
| 459 | bullet | REQ-RE-0303 |
| 460 | bullet | REQ-RE-0165, REQ-RE-0168, REQ-RE-0304, REQ-RE-0305, REQ-RE-0306 |
| 461 | bullet | REQ-RE-0307 |
| 462 | bullet | REQ-RE-0308 |
| 463-466 | rule, heading | context: separator and section heading |
| 467-468 | paragraph | REQ-RE-0309 |
| 469-472 | rule, heading | context: separator and section heading |
| 473-474 | paragraph | REQ-RE-0004 |
| 475-477 | table header | context: table header |
| 478-481 | table row | REQ-RE-0310 (addresses not copied) |
| 482-483 | paragraph | REQ-RE-0311 |
| 484-485 | heading | context: section heading |
| 486-487 | paragraph | REQ-RE-0001 |

### README.hu.md (lines 1-721): content only in Hungarian, and contradictions

The other lines of `README.hu.md` translate `README.md` (same content in a different order: the Hungarian file has a table of contents, the licence before Tested Status, and the CLI and services in one section) and are covered by the English rows above. Lines that add content or contradict the English file:

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 18-41 | numbered list | context: table of contents (Hungarian only) |
| 49-53 | heading, paragraph | context: problem statement, longer than README.md:24 |
| 59 | numbered item | context: same as README.md:129-130 (REQ-RE-0087, -0089) |
| 60 | numbered item | REQ-RE-0312 |
| 66 | bullet | context: Python-only (no pip dependencies); replaced by REQ-RE-0176 |
| 87 | bullet | context: contradiction; claims every package manager is supported, README.md:117 says untested (REQ-RE-0068, English wins) |
| 94 | numbered item | context: claim that RMBT is a BEREC standard; no requirement beyond REQ-RE-0083 |
| 98 | bullet | context: stale ("not tested on 512 MB"); README.md:322 and README.hu.md:677 say tested in a VM (REQ-RE-0209) |
| 100 | bullet | REQ-RE-0313 |
| 101 | bullet | REQ-RE-0314 |
| 109 | bullet | context: "loss > 5 %" alert, see Notes of REQ-RE-0149 |
| 160-169 | table | REQ-RE-0315; line 169 (software bridge, hardware bypass on failure) contradicts README.md:92: see Questions |
| 175 | paragraph | context: lead-in |
| 181 | bullet | REQ-RE-0032 (DHCP and NAT bound to the bridge) |
| 184 | bullet | REQ-RE-0316 |
| 187-190 | paragraph, bullets | REQ-RE-0317 |
| 196 | paragraph | context: "5 independent layers", lead-in |
| 199-206 | bullets | REQ-RE-0069 (target list conflict noted), REQ-RE-0070, REQ-RE-0071 |
| 211 | bullet | REQ-RE-0318 |
| 214-215 | bullets | context: stale; "100 MB and 25 MB payloads" contradicts the timed 4-stream engine of README.md:124 (REQ-RE-0077); the M-06 brief takes the behaviour from the code |
| 218 | bullet | REQ-RE-0319 |
| 225-226 | bullets | context: claims (BEREC-accepted, free of ISP influence); no requirement beyond REQ-RE-0083 |
| 229-230 | bullets | REQ-RE-0320 |
| 233 | bullet | context: overclaim ("checks whether the ISP throttles"); README.md:128 wins (REQ-RE-0086) |
| 259-279 | heading, code block, list, paragraph | REQ-RE-0321 to REQ-RE-0324, REQ-RE-0100, REQ-RE-0101; line 273 (Python garbage collection) is Python-only context |
| 285-296 | paragraph, diagram, list | REQ-RE-0325, REQ-RE-0326, REQ-RE-0104 |
| 302 | paragraph | context: "14 dedicated categories", lead-in to REQ-RE-0327 to -0333 |
| 305-327 | numbered list | REQ-RE-0327 to REQ-RE-0333 (pages 4 and 10 are REQ-RE-0320 and -0313) |
| 330 | paragraph | REQ-RE-0334 |
| 364 | paragraph | context: why remote access matters |
| 400 | bullet | REQ-RE-0335 |
| 417 | bullet | context: overclaim ("determines automatically"); README.md:207 wins (REQ-RE-0142) |
| 420 | bullet | REQ-RE-0144 |
| 422 | bullet | REQ-RE-0146 |
| 429 | bullet | REQ-RE-0147 |
| 445 | bullet | REQ-RE-0336 |
| 535 | paragraph | context: modular daemons, lead-in |
| 554 | paragraph | REQ-RE-0285 |
| 583 | table row | context: fault in the Hungarian text ("official" SLA dossier), against README.md:225 and README.hu.md:440 (REQ-RE-0151) |
| 620-622 | heading, bullets | REQ-RE-0337 |
| 649-651 | bullets | REQ-RE-0231 (repeated login), REQ-RE-0338 |
| 656 | paragraph | REQ-RE-0001 (wording: free to use, modify and embed) |

## Proposed package changes

- No new package beyond `NEW: X2-01 Step-2 router tiers`, already proposed in `roadmap-1-507.md` and used in `porting-plan.md`. Here it owns REQ-RE-0173, -0177 to -0183 and -0192, and is "Also in" for REQ-RE-0002, -0174, -0176, -0193 and -0195.
- If `NEW: M-17 Shared HTTP transfer core` from `roadmap-1-507.md` is accepted, it should own REQ-RE-0193 (in-process engine clients) in place of M-04.
- Split of `U-02…U-0n`: README.hu.md section 8 (lines 305-327) gives the Python page list. With the pages named in this part, a split could be: Dashboard/Home, Ping, Speed (with Schedule cards: Wi-Fi uplink, USB, data caps, Wi-Fi share, traffic gate), iPerf3, Outages, SLA, Traceroute/Port check/RTR, Devices, Developer console and AI diagnostics, Router (port forwarding), Ports & Drivers, VPN and Remote, Blocker, Tasks, Settings, System (update, collector, memory), Results (public), Support. The rows here name their page so the split can take them over.
- Box identity (REQ-RE-0169) has no package of its own; C-01 is proposed as owner, since the ID is made once at first start and stored with the settings. If the register owner prefers, S-12 can own it.

## Questions for the owner

1. Mode B when the box fails: README.md:92 says "if the box stops, the connection stops too", but README.hu.md:169 says that on a software fault the software bridge still passes traffic and bypass hardware lets it through. Is a fail-open bypass (hardware or software) required for the native Mode B, or is the Hungarian row stale? (REQ-RE-0028, -0315)
2. polkit 0.105 (Ubuntu 22.04): the Python rule cannot name units, so there the service user may manage every systemd service (README.md:452). Is that accepted for the native build on Ubuntu 22.04, or must it be fixed (unit control moved to the privileged helper, or 22.04 dropped)? It is a candidate for F-03. (REQ-RE-0288)
3. Public-address guard on Alpine: the Python guard needs systemd and nftables and is not installed on Alpine (README.md:453), so ports 80, 53 and 5201 stay open to public addresses there. Must the native firewall (N-02, "for every mode") also cover Alpine/OpenRC, or is Alpine without the guard accepted? (REQ-RE-0290)
4. Donation addresses: README.md:476-483 shows the owner's crypto addresses in the README, the Web UI, the bot and the CLI. Quality Bar 8 keeps personal data out of the native code. Do the addresses go into the native build as product data (as now), or are they given at build or release time from outside the code? (REQ-RE-0310, -0311)
