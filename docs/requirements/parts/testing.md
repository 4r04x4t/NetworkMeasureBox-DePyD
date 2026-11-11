# Requirements register, part: TESTING.md

Source: `docs/TESTING.md` of the Python repository at commit `ad3758c` (401 lines; this part covers lines 1-401).
Read in pages: 1-150, 151-280, 281-401 (all lines). Compared with `docs/TESTING.hu.md@ad3758c` (399 lines, read 1-21, every table row's first cell, and 276-399 in full).
Package ids from `docs/NATIVE_WORK_PLAN.md@ad3758c:98-226` (section 3 and section 4).
Id range of this part: REQ-TS-0001 to REQ-TS-0314. Every Source cell is written out in full as `docs/TESTING.md@ad3758c:<line>`.
API and page packages are written as `S-02…S-0n (API: <area>)` and `U-02…U-0n (page: <page>)`, using the areas of work plan line 189, until F-02 fixes their numbers.

Hungarian comparison: every heading, table row and bullet of `TESTING.md` has a counterpart in `TESTING.hu.md` (same row counts per section; Measurement rows Cloudflare and Fast.com are in the other order). No row exists only in Hungarian. One paragraph exists only in English: `docs/TESTING.md@ad3758c:7` (environment of the 2026-11-10 morning work, REQ-TS-0002 and REQ-TS-0003).

## Requirements

| Id | Requirement | Source | Package | Also in | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| REQ-TS-0001 | The native build keeps a testing record (what was tested, where, how, with what result), updated whenever a test is run again or a new one is added; the README's tested-status table summarises it. | docs/TESTING.md@ad3758c:3; docs/TESTING.md@ad3758c:276 | P-13 | P-12 |  |
| REQ-TS-0002 | Every batch of work records the environment it was made and tested in (box OS, development PC OS, what ran where, what was only read). | docs/TESTING.md@ad3758c:7; docs/TESTING.md@ad3758c:9 | P-13 | F-00 | Line 7 exists in English only; TESTING.hu.md has no counterpart (line 9 has one). |
| REQ-TS-0003 | The unit tests pass in several full runs in a row, the first one with an empty home folder (test isolation: no dependence on the developer's home, earlier runs or the box). | docs/TESTING.md@ad3758c:7 | F-06 | F-00 | Python figure: 558 tests, three runs. |
| REQ-TS-0004 | The testing record states the line type and the hardware every speed, CPU and timing figure comes from. | docs/TESTING.md@ad3758c:13 | P-13 | R-01 | (source names a specific site value; not copied) Source names the ISP's country, the plan and the home router model. |
| REQ-TS-0005 | Lines much slower (DSL, mobile) and much faster (2.5 or 10 Gbps) than the test line are recorded as not tested until they are. | docs/TESTING.md@ad3758c:13 | P-13 |  |  |
| REQ-TS-0006 | The record states that above about 1 Gbps the Pi 4's single gigabit port and its CPU limit the measurement before the line does. | docs/TESTING.md@ad3758c:13 | P-13 | R-01, M-04 |  |
| REQ-TS-0007 | Live tests run on a Raspberry Pi 4 with Debian 13 (arm64) in daily use: in standalone mode wired behind a router, and also directly on the ISP's ONT over PPPoE in router mode with a USB mobile hotspot keeping remote access up. | docs/TESTING.md@ad3758c:17 | P-13 | P-11, N-05, N-10 | (source names a specific site value; not copied) Source names the router model, the ISP and the hotspot model. |
| REQ-TS-0008 | A VM test bed: QEMU/KVM on x86-64, Debian 13 with the full kernel, two NICs (user-mode NAT as WAN, a socket link as LAN), Wi-Fi with mac80211_hwsim (three simulated radios), clients in network namespaces. | docs/TESTING.md@ad3758c:18; docs/TESTING.md@ad3758c:251 | F-06 | P-13 |  |
| REQ-TS-0009 | Clean VMs of each supported distribution, each installed from scratch with the installer. | docs/TESTING.md@ad3758c:19; docs/TESTING.md@ad3758c:250 | P-11 | F-06 |  |
| REQ-TS-0010 | The unit tests run on the developer PC and, as root, inside the VM test bed. | docs/TESTING.md@ad3758c:20 | F-06 | F-00 | Python count 303 is historical. |
| REQ-TS-0011 | Results from virtual NICs and simulated radios are marked as simulated; they do not cover timing, drivers and radio behaviour of real hardware. | docs/TESTING.md@ad3758c:22; docs/TESTING.md@ad3758c:245 | P-13 | F-06 |  |
| REQ-TS-0012 | Clean install, reinstall and reboot pass on Debian 13. | docs/TESTING.md@ad3758c:32 | P-11 | P-10 |  |
| REQ-TS-0013 | Clean install and reinstall pass on Debian 12 with a real speed test; reboot still to prove. | docs/TESTING.md@ad3758c:33; docs/TESTING.md@ad3758c:363 | P-11 |  |  |
| REQ-TS-0014 | The privilege rule lets the service account restart the network daemon but not other units (for example ssh). | docs/TESTING.md@ad3758c:33 | P-06 | P-04, P-11 | Python used polkit; native mechanism is P-06's choice. |
| REQ-TS-0015 | Clean install on Ubuntu 24.04 LTS; the wired interface is detected automatically (predictable names such as ens3, not only eth0). | docs/TESTING.md@ad3758c:34 | N-01 | P-11, N-17 | Reinstall and reboot not tested. |
| REQ-TS-0016 | Ubuntu 26.04 LTS: clean install, reinstall over a running box and reboot pass; measurements are kept and keep growing, a changed admin password is kept, the config is 0640 owned by the service user, a planted stale file in the install folder is removed, no crash traces in the logs; an RTR test runs. | docs/TESTING.md@ad3758c:35 | P-11 | C-01, C-02, M-08 |  |
| REQ-TS-0017 | Ubuntu 22.04 LTS (old polkit): clean install, reinstall and reboot pass; scheduled Ookla tests run. | docs/TESTING.md@ad3758c:36 | P-11 | M-10 |  |
| REQ-TS-0018 | The service account must not be able to manage any system unit: the Python rule for the old polkit (.pkla) granted that and is not carried over; the native privilege rule stays an allow-list on every base. | docs/TESTING.md@ad3758c:36 | P-06 | P-04 | Python fault (QB 6, QB 7). |
| REQ-TS-0019 | Rocky Linux 10.2: the extra package repository (EPEL) is enabled from the OS id; reinstall keeps measurements and the changed admin password, config 0640 service user, stale file removed, no crash traces; RTR runs. | docs/TESTING.md@ad3758c:37 | P-11 |  | Test note: the cloud image needs a disk of at least 10 GB. |
| REQ-TS-0020 | Alpine 3.22 with OpenRC: clean install, reinstall over a running box and reboot pass; measurements and changed admin password kept, config 0640 service user, stale file removed, no crash traces in the log folder; RTR runs. | docs/TESTING.md@ad3758c:38 | P-11 | P-10 |  |
| REQ-TS-0021 | On Alpine/OpenRC the web server may bind port 80 without root (capability granted by the service definition). | docs/TESTING.md@ad3758c:39 | P-04 | P-10, S-01 |  |
| REQ-TS-0022 | The RTR (RMBT) engine builds and runs on Alpine (musl). | docs/TESTING.md@ad3758c:39 | M-08 | P-11 |  |
| REQ-TS-0023 | Service restart works under OpenRC. | docs/TESTING.md@ad3758c:39 | P-10 | P-11 |  |
| REQ-TS-0024 | Clean install, reinstall and reboot pass on the Raspberry Pi 4 with Debian 13 (the box). | docs/TESTING.md@ad3758c:40 | P-11 | P-13 |  |
| REQ-TS-0025 | Fedora 45 (dnf branch): all services, the privilege rule and all optional packages install; RTR and Cloudflare measure; the Ookla client runs; reboot still to prove. | docs/TESTING.md@ad3758c:41; docs/TESTING.md@ad3758c:363 | P-11 | M-06, M-08, M-10 | Ookla was refused by its rate limit on a shared IP. |
| REQ-TS-0026 | With SELinux enforcing there are no SELinux denials. | docs/TESTING.md@ad3758c:41 | P-04 | P-11 |  |
| REQ-TS-0027 | AlmaLinux 10.2: install, reinstall and reboot pass; base detection does not misread ID_LIKE (rebuilds listing "fedora" must still get EPEL); nothing relies on sudo searching /usr/local/bin. | docs/TESTING.md@ad3758c:42 | P-11 | P-06 | Python faults found and fixed; test note: the image needs a SCSI disk in QEMU. |
| REQ-TS-0028 | The update check counts the OS package manager's updates (dnf as well). | docs/TESTING.md@ad3758c:42 | P-01 | P-02 | Python reference unclear between heal and updater; assumed the heal 'updates' view. |
| REQ-TS-0029 | Raspberry Pi OS, Arch and 32-bit systems are listed as not tested until they are. | docs/TESTING.md@ad3758c:43; docs/TESTING.md@ad3758c:371-376 | P-11 |  |  |
| REQ-TS-0030 | A scheduled Ookla test at an hourly schedule stores a row with download, upload, ping, bufferbloat grade and the chosen server. | docs/TESTING.md@ad3758c:49 | M-10 | M-04, R-01 | (source names a specific site value; not copied) Source names the router model, a measured result and the server number. |
| REQ-TS-0031 | On an Ookla refusal ("Limit reached") the box pauses Ookla (30 min, up to 4 h) and stores no other engine's result in its place. | docs/TESTING.md@ad3758c:49; docs/TESTING.md@ad3758c:341; docs/TESTING.md@ad3758c:397 | M-05 | M-10, M-04 |  |
| REQ-TS-0032 | RTR-NetTest measures close to the nominal plan of a gigabit line (Python: about 91 % down, about 107 % of the up plan). | docs/TESTING.md@ad3758c:50 | M-08 | R-01 | (source names a specific site value; not copied) |
| REQ-TS-0033 | The Cloudflare engine works on the box. | docs/TESTING.md@ad3758c:51 | M-06 |  |  |
| REQ-TS-0034 | Wi-Fi shared on a LAN port: a client on the port gets an address by DHCP, pings and downloads over HTTPS through the Wi-Fi uplink, with the wired uplink untouched. | docs/TESTING.md@ad3758c:52 | N-08 | N-05 |  |
| REQ-TS-0035 | With a second DHCP server on the client side the box stops its own DHCP, address, NAT and routing rule; when that server goes away, sharing comes back within 15 s. | docs/TESTING.md@ad3758c:52 | N-05 | N-08 | Guard reference is netd/wifi_share.py; see Proposed package changes. |
| REQ-TS-0036 | The DHCP guard receives the other server's replies (packet socket, not a normal socket that drops them) and waits long enough for slow servers (Python: three requests over 8 s; dnsmasq offers after up to 3 s). | docs/TESTING.md@ad3758c:52 | N-05 | F-06 | Python faults found and fixed. |
| REQ-TS-0037 | Helper daemons started by the box run in the foreground under supervision (Python: --log-facility=- kept dnsmasq in the foreground unexpectedly). | docs/TESTING.md@ad3758c:52 | P-10 | N-03 | Interpretation: lesson for native process control. |
| REQ-TS-0038 | With measurement_link = wifi the test traffic goes over the Wi-Fi interface only, even with the cable in (byte counters prove it). | docs/TESTING.md@ad3758c:53 | M-04 | N-01, C-11 |  |
| REQ-TS-0039 | The Fast.com engine measures close to the line's plan with several servers and a bufferbloat grade. | docs/TESTING.md@ad3758c:54 | M-07 | R-01 | (source names a specific site value; not copied) |
| REQ-TS-0040 | The iPerf3 client works on the box and in the VM test bed. | docs/TESTING.md@ad3758c:55 | M-09 |  |  |
| REQ-TS-0041 | The ping probe and outage detection run continuously from install on. | docs/TESTING.md@ad3758c:56 | M-01 | M-02 |  |
| REQ-TS-0042 | The schedule runs tests at the set times. | docs/TESTING.md@ad3758c:57 | M-04 |  |  |
| REQ-TS-0043 | The daily data limit is enforced. | docs/TESTING.md@ad3758c:57 | M-05 | M-04 |  |
| REQ-TS-0044 | The port check (for example 22, 5060, 6881) works. | docs/TESTING.md@ad3758c:58 | M-11 |  |  |
| REQ-TS-0045 | Traffic stop (experimental): a test aborts when other devices' traffic is present (Python: aborts at 23-40 Mbps); off by default. | docs/TESTING.md@ad3758c:59 | M-03 | M-04 |  |
| REQ-TS-0046 | DNS interception and the NAT type are detected and reported (record: no interception, endpoint-independent NAT). | docs/TESTING.md@ad3758c:60 | M-12 |  |  |
| REQ-TS-0047 | Router mode (mode A): DHCP, NAT and internet for LAN clients (client in a namespace gets an address, pings the internet, loads HTTPS). | docs/TESTING.md@ad3758c:66 | N-03 | N-02, N-05 |  |
| REQ-TS-0048 | Bridge mode (mode B): the box gets its own address on the bridge. | docs/TESTING.md@ad3758c:67 | N-03 |  |  |
| REQ-TS-0049 | DMZ isolation: the DMZ reaches the internet, not the LAN and not the Web UI. | docs/TESTING.md@ad3758c:68 | N-02 | N-03 |  |
| REQ-TS-0050 | PPPoE dial-in and redial against a PPPoE server; the box's own traffic uses the PPP link; a LAN client's full round trip over PPPoE is still to prove. | docs/TESTING.md@ad3758c:69 | N-05 | N-03 |  |
| REQ-TS-0051 | After a box restart the LAN DHCP server is authoritative: a client renewing an old lease gets it at once. | docs/TESTING.md@ad3758c:70 | N-03 | N-05 |  |
| REQ-TS-0052 | A real USB mobile hotspot works on the Pi 4 as measuring link, as router-mode WAN, and as the box's remote-access path. | docs/TESTING.md@ad3758c:71 | N-10 | M-04, N-03 | (source names a specific site value; not copied) Source names the hotspot model. |
| REQ-TS-0053 | The self-test has a guided PPPoE scenario (`nmb-selftest scenario pppoe`) that guides the live PPPoE and router-mode checks, usable when the box is reached only over a remote path. | docs/TESTING.md@ad3758c:75 | P-07 | N-05 | Python result counts and the steps not done at night are history. |
| REQ-TS-0054 | PPPoE discovery, login and kernel-mode PPPoE with the access concentrator's service name; a refused login shows the server's own reason; the saved user name is used exactly as entered. | docs/TESTING.md@ad3758c:79 | N-05 | S-02…S-0n (API: network) | (source names a specific site value; not copied) Concentrator and service names not copied. |
| REQ-TS-0055 | Internet, HTTPS and MTU 1492 over PPPoE; connectivity checks tolerate the ISP dropping the first packets after link-up (first answer about 6 s after the link came up). | docs/TESTING.md@ad3758c:80 | N-05 | M-01 |  |
| REQ-TS-0056 | Automatic mode picks router mode with PPPoE about 15 s after mode = auto with the cable on the ONT (PPPoE login saved). | docs/TESTING.md@ad3758c:81 | N-04 | N-05 |  |
| REQ-TS-0057 | NAT and MSS clamping on the PPP WAN. | docs/TESTING.md@ad3758c:82 | N-02 | N-05 |  |
| REQ-TS-0058 | In router mode on PPPoE the box's own internet and DNS use the PPP link; access point and DHCP serve the LAN; a simulated LAN client gets internet. | docs/TESTING.md@ad3758c:82 | N-03 | N-05, N-07 |  |
| REQ-TS-0059 | Speed tests over PPPoE bind to the PPP interface for every engine (Cloudflare, Fast.com, RTR, iPerf3) and reach close to the plan. | docs/TESTING.md@ad3758c:83 | M-04 | M-06, M-07, M-08, M-09 | (source names a specific site value; not copied) |
| REQ-TS-0060 | Redial: after the PPPoE session hangs up the box is back in about 12 s; when the PPPoE process stops, the network daemon starts it again (Python: 41 s). | docs/TESTING.md@ad3758c:84 | N-05 | P-01 |  |
| REQ-TS-0061 | A real Wi-Fi client gets internet through router mode on PPPoE (2.4 GHz, 802.11n). | docs/TESTING.md@ad3758c:85 | N-03 | N-07 |  |
| REQ-TS-0062 | The WAN is closed to the internet: SSH, DNS (TCP and UDP), Web UI and iperf3 closed on IPv4 and IPv6; ping answered. | docs/TESTING.md@ad3758c:86 | N-02 | M-09, N-12, S-01 |  |
| REQ-TS-0063 | The box's own IPv6 on PPPoE: an address and default route from the ISP's router advertisement; HTTPS over IPv6 works. | docs/TESTING.md@ad3758c:87 | N-05 | N-01 |  |
| REQ-TS-0064 | IPv6 for the LAN: a delegated /64 is advertised to the LAN; a LAN client gets a SLAAC address and reaches the internet over IPv6; a LAN device's open port is not reachable from the internet. | docs/TESTING.md@ad3758c:88 | N-11 | N-02, N-05 |  |
| REQ-TS-0065 | Port forwarding: a LAN device is reached from another network through the line's address; hairpin from a second LAN device works; the self-test port-forward group passes. | docs/TESTING.md@ad3758c:89 | N-11 | N-02, P-07 |  |
| REQ-TS-0066 | UPnP / NAT-PMP / PCP: mappings from a LAN device work from another network; a mapping for port 22 and a mapping for another device are refused; mappings are removed when UPnP is turned off. | docs/TESTING.md@ad3758c:90 | N-11 | N-02 |  |
| REQ-TS-0067 | DynDNS updates at once when the box's way to the internet changes (not only at the next scheduled update) and sends only an IPv4 address as the IPv4 (A) record, never an IPv6 address. | docs/TESTING.md@ad3758c:91 | N-16 | N-01 | Python faults found and fixed. |
| REQ-TS-0068 | Access point with default settings (auto): 802.11ac, 80 MHz, channel 36 at the first try on 5 GHz. | docs/TESTING.md@ad3758c:92 | N-07 |  |  |
| REQ-TS-0069 | In router mode speed tests bind to the router's WAN; never to an interface with only a link-local address, and never to a USB hotspot (mobile data) when the WAN is down or during a network-daemon restart: the guard skips the test ("no test over the USB hotspot"). | docs/TESTING.md@ad3758c:93 | M-04 | N-03, N-10, C-11 | Python faults found and fixed. |
| REQ-TS-0070 | When the network daemon stops or restarts with PPPoE up, it hangs up first (PADT sent within a second) and its stop timeout is long enough for that (Python: 45 s); the next login is accepted at once. | docs/TESTING.md@ad3758c:94 | N-05 | P-10 | Python fault: a stale session blocked logins for about 6 minutes. |
| REQ-TS-0071 | The data volume each link moves under the schedule is counted and can be shown (record: hundreds of GB in hours with every engine every few minutes). | docs/TESTING.md@ad3758c:100 | M-05 | C-09, R-01 | Source marks this row as a figure, not a test. |
| REQ-TS-0072 | Data caps: usage counted per link since the period start; a cap equal to the used amount counts as reached; a scheduled cycle is then skipped without running an engine; the self-test data-cap group passes and restores the caps. | docs/TESTING.md@ad3758c:101 | M-05 | M-04, P-07 |  |
| REQ-TS-0073 | A schedule set to a subset of engines (for example without iPerf3 and Ookla) runs only those in its next cycles. | docs/TESTING.md@ad3758c:102 | M-04 |  |  |
| REQ-TS-0074 | The ping probe follows the active WAN (moves to the PPP interface within 30 s) and switches off targets that do not apply (Local Gateway with no gateway on PPPoE, an ISP first hop that does not answer pings). | docs/TESTING.md@ad3758c:103 | M-01 | N-01, C-11 | Python fault: probe bound to the wired port without address gave 100 % loss for hours. |
| REQ-TS-0075 | Loss caused by the probe's own wrong binding must not be recorded as line outages, and no outage is left open by it. | docs/TESTING.md@ad3758c:103 | M-02 | M-01, R-01 | Python fault: 8 outages left open. |
| REQ-TS-0076 | In router mode the WAN port is not managed by another network manager (no repeated activation attempts on it). | docs/TESTING.md@ad3758c:104 | N-03 | N-01 |  |
| REQ-TS-0077 | A client on the 5 GHz access point with default settings joins, gets IPv4 and an IPv6 address from the delegated prefix, DNS and HTTPS work over both, no ping lost; the record gives link rate, Wi-Fi-only and internet throughput and latency. | docs/TESTING.md@ad3758c:105 | N-07 | N-11, P-13 | See the Known limit on throughput (line 395). |
| REQ-TS-0078 | Changing the access point's band reloads only the access point: the PPPoE session stays (same process, no redial), no ping lost, the AP is up again in about 3 s. | docs/TESTING.md@ad3758c:106 | N-07 | N-03, N-05 | Python fault: a band change rebuilt the whole router (LAN offline about 40 s). |
| REQ-TS-0079 | A network-daemon restart in automatic mode dials once and builds router mode once. | docs/TESTING.md@ad3758c:107 | N-04 | N-05 | Python fault: two dials 20 s apart, most pings lost. |
| REQ-TS-0080 | Changing the mode setting to one that resolves to the running mode (router to automatic and back) keeps the running mode: same PPPoE session, no ping lost. | docs/TESTING.md@ad3758c:108 | N-03 | N-04 |  |
| REQ-TS-0081 | A real phone on the box's Wi-Fi gets IPv6 and IPv4 from the line and passes public IPv6 test sites. | docs/TESTING.md@ad3758c:109 | N-11 | N-07 |  |
| REQ-TS-0082 | Bot commands that change the box work for real through the bot's own handler: mode change (no redial), VPN stop/start, heal and heal fix, ad blocker off/on, backup create and list, reboot now (box back in its mode with one dial, IPv4 and IPv6, Wi-Fi up). | docs/TESTING.md@ad3758c:110 | S-11 | N-03, N-13, N-14, P-01, N-12, P-03, P-10 |  |
| REQ-TS-0083 | The update check never offers older or equal code as an update; when the latest runs it says so with the commit. | docs/TESTING.md@ad3758c:110 | P-02 | S-11, P-01 | Python fault: offered "2.5.0 -> 2.5.0". |
| REQ-TS-0084 | Connecting the Wi-Fi client is refused while the radio is the access point. | docs/TESTING.md@ad3758c:110 | N-07 | S-11, N-08 |  |
| REQ-TS-0085 | Turning DynDNS off stops every updater, also the one inside the web server process (within 35 s); turning it on brings it back; a manual update runs. | docs/TESTING.md@ad3758c:110 | N-16 | S-11 | Python fault: the web server's updater kept running. |
| REQ-TS-0086 | Purging measurements works on a nearly full disk by deleting in slices (no "database or disk is full"). | docs/TESTING.md@ad3758c:110 | C-08 | C-02, S-11 |  |
| REQ-TS-0087 | A reboot keeps the database up to the last minute (snapshot on shutdown). | docs/TESTING.md@ad3758c:110 | C-02 | P-10 |  |
| REQ-TS-0088 | Log rotation fits the log file system (for example a 48 MB RAM disk): logs never fill it; a full log disk must not stop services or the installer. | docs/TESTING.md@ad3758c:111 | C-03 | C-08, P-11 | Python fault: developer log 10 MB x 5 filled /var/log. |
| REQ-TS-0089 | Killing a task: only the bot's own processes are stopped (with TERM), root processes and PID 1 are refused, a missing PID is reported, signal 19 (stop) is refused, and the signal sent is the one named. | docs/TESTING.md@ad3758c:112 | C-09 | S-11, S-10, P-06 | Python fault: sent TERM while naming 19. |
| REQ-TS-0090 | Purging logs (with confirm) empties the NMB logs and their rotated copies, keeps other files, rotates and vacuums the system journal and clears the kernel buffer through the privileged helper, and also clears the logs kept on the SD card. | docs/TESTING.md@ad3758c:113 | C-08 | P-06, S-11, F-06 | Python test ran in an isolated mount namespace with a recorder for journalctl and dmesg. |
| REQ-TS-0091 | The storage guard also watches the log RAM disk: it deletes the oldest rotated log first and stops at its target free space; the live developer log is untouched. | docs/TESTING.md@ad3758c:114 | C-08 | C-03 | Python fault: only the data disk was checked. |
| REQ-TS-0092 | After a reboot in router mode the bot starts together with the other services and does not wait for a network-online target that fails on the ONT; its first poll may fail, then it works. | docs/TESTING.md@ad3758c:115 | P-10 | S-11 |  |
| REQ-TS-0093 | The installer on a nearly full log disk removes old rotated logs first, finishes, and resizes the old default developer log; it warns when the factory Web UI login is in use. | docs/TESTING.md@ad3758c:116 | P-11 | C-03, S-01 | Owner keeps the factory Web UI login on purpose; the warning is product behaviour only. |
| REQ-TS-0094 | Automatic mode, cable moved from the ONT to a router: standalone within about 2 minutes ("a router gives addresses"), DHCP server, PPPoE and access point stopped, no stale boot route left; DynDNS updated at once. | docs/TESTING.md@ad3758c:117 | N-04 | N-03, N-05, N-16 |  |
| REQ-TS-0095 | The probe's target set follows the mode: Local Gateway back on in standalone, ISP first hop off where it does not answer. | docs/TESTING.md@ad3758c:117 | M-01 | N-04 |  |
| REQ-TS-0096 | The wired route metric is set as intended in the active connection profile (Python: 50) and a stale DNS fallback marker is cleared on a mode change. | docs/TESTING.md@ad3758c:117 | N-01 | N-04 | Python faults found and fixed. |
| REQ-TS-0097 | Connecting to a saved Wi-Fi profile works in standalone. | docs/TESTING.md@ad3758c:117 | N-08 | S-11 |  |
| REQ-TS-0098 | The update's test gate runs unprivileged on a copy with throwaway paths and never writes the live settings. | docs/TESTING.md@ad3758c:118 | P-02 | C-01, F-06 | Python fault: tests run as root overwrote nmb.conf (PPPoE login, bot chat). |
| REQ-TS-0099 | An update coordinates with a running scheduled measurement (the measurement lock) instead of failing on it. | docs/TESTING.md@ad3758c:118 | P-02 | M-04 |  |
| REQ-TS-0100 | A measurement cycle does not fail while counting data. | docs/TESTING.md@ad3758c:118 | M-05 | M-04 | Python fault found by the update's test gate. |
| REQ-TS-0101 | A short WAN cable loss (3 s) keeps the PPPoE session: no redial, no outage, no ping lost. | docs/TESTING.md@ad3758c:119 | N-05 | M-02 |  |
| REQ-TS-0102 | A migrated Python database is taken with the owner's manual corrections as they are (deleted false outage rows do not come back; availability is computed from the corrected rows). | docs/TESTING.md@ad3758c:120 | C-02 | M-02, R-01 | Interpretation: the row records a one-off owner-approved data correction after a backup. |
| REQ-TS-0103 | Access point WPA2 on 2.4 GHz: a client connects, gets DHCP, reaches the internet, HTTPS and the Web UI. | docs/TESTING.md@ad3758c:126 | N-07 | N-03 |  |
| REQ-TS-0104 | Access point WPA2 on 5 GHz (channel 36) passes the same checks; the regulatory database it needs is installed. | docs/TESTING.md@ad3758c:127 | N-07 | P-11 |  |
| REQ-TS-0105 | Without the regulatory database a 5 GHz setting falls back to 2.4 GHz instead of leaving the box without Wi-Fi. | docs/TESTING.md@ad3758c:128 | N-07 |  |  |
| REQ-TS-0106 | Access point WPA3 (SAE) on 2.4 GHz passes the same checks as WPA2. | docs/TESTING.md@ad3758c:129 | N-07 |  |  |
| REQ-TS-0107 | A wrong Wi-Fi password gives no address and no traffic. | docs/TESTING.md@ad3758c:130 | N-07 | N-02 |  |
| REQ-TS-0108 | When the AP software rejects Wi-Fi 7 / 6 GHz / 160-320 MHz settings, the box steps down to Wi-Fi 6. | docs/TESTING.md@ad3758c:131 | N-07 |  |  |
| REQ-TS-0109 | The access point on real Wi-Fi hardware is still to prove for every case above. | docs/TESTING.md@ad3758c:132; docs/TESTING.md@ad3758c:318 | N-07 | P-13 |  |
| REQ-TS-0110 | The box as Wi-Fi client: scan, connect, DHCP and ping through the Web UI API. | docs/TESTING.md@ad3758c:138 | N-08 | S-02…S-0n (API: Wi-Fi), N-07 |  |
| REQ-TS-0111 | A wrong password shows the Wi-Fi backend's real error message. | docs/TESTING.md@ad3758c:139 | N-08 | S-02…S-0n (API: Wi-Fi) |  |
| REQ-TS-0112 | Reconnect, disconnect and saved profiles: a stale profile does not block reconnecting; "disconnected" is never reported as connected. | docs/TESTING.md@ad3758c:140 | N-08 |  | Python faults found and fixed. |
| REQ-TS-0113 | The unprivileged service account may control the Wi-Fi client through the privilege rule; an upgrade replaces an old installed rule. | docs/TESTING.md@ad3758c:141 | P-06 | P-11, N-08 |  |
| REQ-TS-0114 | Without the Wi-Fi client backend (Python: NetworkManager, not installed by the installer) the feature is reported as unavailable instead of failing silently. | docs/TESTING.md@ad3758c:142 | N-08 | P-11 | Python: fails by design; see Questions. |
| REQ-TS-0115 | WireGuard server: a peer added in the Web UI, its config used by a client: tunnel, internet, HTTPS and Web UI through the tunnel; peer delete and stop work. | docs/TESTING.md@ad3758c:148 | N-13 | S-02…S-0n (API: VPN), U-02…U-0n (page: VPN) |  |
| REQ-TS-0116 | OpenVPN server: the server certificate has TLS server usage; old certificates without it are reissued from the same CA; a client profile works end to end. | docs/TESTING.md@ad3758c:149 | N-14 |  | Python fault: every client was rejected. |
| REQ-TS-0117 | HTTPS/CONNECT proxy: no long delay on networks with broken IPv6 (Python: 0.4 s), stop frees the port, connections to the box's own loopback services are refused. | docs/TESTING.md@ad3758c:150 | N-11 | N-02 | Python faults found and fixed (20 s delay; port not freed; loopback reachable). |
| REQ-TS-0118 | DuckDNS: a wrong token is rejected and the error shown; a real account is still to prove. | docs/TESTING.md@ad3758c:151; docs/TESTING.md@ad3758c:363 | N-16 |  |  |
| REQ-TS-0119 | Tailscale is still to prove (needs an account). | docs/TESTING.md@ad3758c:152; docs/TESTING.md@ad3758c:363 | N-15 |  |  |
| REQ-TS-0120 | DNS ad blocker under load: 3000 queries from 50 threads all answered; rate, p99 latency and memory recorded and not worse than Python (6900/s, p99 72 ms, 24 MB). | docs/TESTING.md@ad3758c:158 | N-12 | P-13 |  |
| REQ-TS-0121 | Self-update with rollback: a good update is applied; an update whose tests fail is rolled back. | docs/TESTING.md@ad3758c:159 | P-02 | F-06 |  |
| REQ-TS-0122 | Tests do not depend on the machine they run on (Python: three machine-dependent tests rolled back every update on a real box). | docs/TESTING.md@ad3758c:159 | F-06 | P-02 |  |
| REQ-TS-0123 | JSON output of commands is pure JSON. | docs/TESTING.md@ad3758c:159 | S-10 | P-02 | Python fault: --json output was not pure JSON. |
| REQ-TS-0124 | Settings saved in the Web UI are kept and shown after a reload. | docs/TESTING.md@ad3758c:160 | U-02…U-0n (page: settings) | S-02…S-0n (API: system), C-01 |  |
| REQ-TS-0125 | Results without password (public_results): off gives 401 on /results; on, the results, the report and the CSV files open from the LAN without login while /, /api/config, /api/status and every POST still need it; turning it off is saved. | docs/TESTING.md@ad3758c:161 | S-01 | R-02, R-03, C-01 |  |
| REQ-TS-0126 | A schedule that runs Ookla more often than hourly reports ookla_rate_limit_risk; hourly or without Ookla gives none. | docs/TESTING.md@ad3758c:162 | M-05 | M-04, S-02…S-0n (API: speed) |  |
| REQ-TS-0127 | Factory reset needs an explicit confirmation (--yes). | docs/TESTING.md@ad3758c:163 | P-03 | S-10 |  |
| REQ-TS-0128 | The nmb-cli text UI works on the box. | docs/TESTING.md@ad3758c:164 | S-10 |  |  |
| REQ-TS-0129 | The firewall guard is in place in every mode. | docs/TESTING.md@ad3758c:165 | N-02 |  |  |
| REQ-TS-0130 | The web server runs as the unprivileged service user. | docs/TESTING.md@ad3758c:165 | P-04 | S-01 |  |
| REQ-TS-0131 | The secrets copy on the boot partition works. | docs/TESTING.md@ad3758c:165 | C-04 | P-03 |  |
| REQ-TS-0132 | Kernel network tuning is applied by the network daemon (root). | docs/TESTING.md@ad3758c:166 | N-01 | P-06 |  |
| REQ-TS-0133 | Every bot command, menu button and inline button works in all 11 languages through the real command handler with a stand-in client (0 failures), and commands with good and bad input answer in every language with no Hungarian outside Hungarian. | docs/TESTING.md@ad3758c:167 | S-11 | C-06, P-07 |  |
| REQ-TS-0134 | Reversible bot commands work for real and are undone (kill on an own sleep, backup create, send and delete, VPN add, QR, config and delete with a test client, heal fix); a mixed speed test started from the bot arrives; a real alert marked as a test reaches the chat. | docs/TESTING.md@ad3758c:167 | S-11 | C-09, P-03, N-13, M-04 |  |
| REQ-TS-0135 | Bot commands call the current functions with the current arguments (Python: nine commands broke while mocks kept the tests green). | docs/TESTING.md@ad3758c:167 | S-11 | F-06 |  |
| REQ-TS-0136 | SMTP (e-mail) is still to prove. | docs/TESTING.md@ad3758c:168; docs/TESTING.md@ad3758c:360 | R-05 |  |  |
| REQ-TS-0137 | OONI is still to prove. | docs/TESTING.md@ad3758c:168; docs/TESTING.md@ad3758c:363 | M-13 |  |  |
| REQ-TS-0138 | Drive swap (moving the data to another disk) is still to prove. | docs/TESTING.md@ad3758c:168 | C-08 | C-02, P-03 | Python reference not named; meaning of 'drive swap' assumed. |
| REQ-TS-0139 | On 512 MB RAM (463 MB usable) all services with a test every minute run without swap, OOM or restarts; the web server and probe memory stays flat. | docs/TESTING.md@ad3758c:169 | P-13 | P-10, C-09 | Python: 25 min, 26 of 26 tests, 182-222 MB available. |
| REQ-TS-0140 | Backup and restore: create, download, list, restore from the list and from an uploaded file; config and all measurements come back. | docs/TESTING.md@ad3758c:170 | P-03 | S-02…S-0n (API: system), U-02…U-0n (page: system) |  |
| REQ-TS-0141 | A truncated backup upload is refused and changes nothing. | docs/TESTING.md@ad3758c:170 | P-03 |  |  |
| REQ-TS-0142 | A restore must not lose the restored database to an old journal (WAL), must not be overwritten by the older boot-partition config copy, leaves the config readable by the service user, and an uploaded backup is not world-readable. | docs/TESTING.md@ad3758c:170 | P-03 | C-02, C-04 | Four Python faults found and fixed. |
| REQ-TS-0143 | Power loss during a speed test or right after a snapshot starts: both databases pass an integrity check, config intact, all services up; the interrupted test is simply missing; the snapshot is atomic (temporary file, fsync, rename). | docs/TESTING.md@ad3758c:171 | C-02 | C-01, P-10 |  |
| REQ-TS-0144 | A configured Ookla server that does not exist gives a clear error, not a silent failure. | docs/TESTING.md@ad3758c:172 | M-10 |  | (source names a specific site value; not copied) Source names a wrong built-in server id. |
| REQ-TS-0145 | A paused Ookla is never replaced by another engine's result; the pause survives a restart; paused runs add no rows. | docs/TESTING.md@ad3758c:172 | M-05 | M-10, C-02 |  |
| REQ-TS-0146 | The status API does not fail when the last test failed. | docs/TESTING.md@ad3758c:172 | S-02…S-0n (API: status) | M-04 |  |
| REQ-TS-0147 | Per-engine intervals: an engine with interval N minutes is skipped within N and runs after; 0 runs it every time; Ookla cannot be set below 60 min. | docs/TESTING.md@ad3758c:173 | M-04 | M-05 |  |
| REQ-TS-0148 | Update from the code host without Git: safe unpacking (path traversal refused), check against the host, a status file, root and commit checks; a run against the real repository on a box is still to prove. | docs/TESTING.md@ad3758c:174; docs/TESTING.md@ad3758c:361 | P-02 |  |  |
| REQ-TS-0149 | Reinstall over a running box keeps settings, the admin password and measurements, removes stale files from the install folder, and leaves the config 0640 owned by the service user. | docs/TESTING.md@ad3758c:175; docs/TESTING.md@ad3758c:250 | P-11 | C-01, C-02 |  |
| REQ-TS-0150 | Collector client: an upload repeated adds nothing; diagnostics and events go only with consent; remote commands are refused without consent or from another site; unknown commands are refused. | docs/TESTING.md@ad3758c:176 | S-12 | S-13 |  |
| REQ-TS-0151 | Collector server: search by case number, earlier case number and customer number; CSV; a remote speed test runs and is reported; unknown commands refused; secrets stored as hashes in files with mode 0600; a run between two machines is still to prove. | docs/TESTING.md@ad3758c:176; docs/TESTING.md@ad3758c:342 | S-13 | S-12 |  |
| REQ-TS-0152 | Storage guard: the free-space target follows the disk size, oldest logs go first, the newest backup and the last 7 days of measurements are kept; the CLI storage actions refuse a wrong number; the storage status API and the Web UI disk-space card work; a really full disk is still to prove. | docs/TESTING.md@ad3758c:177; docs/TESTING.md@ad3758c:348 | C-08 | S-10, S-02…S-0n (API: system), U-02…U-0n (page: system) |  |
| REQ-TS-0153 | Saving settings from several programs at once: one save at a time, each with its own temporary file, so the file always parses. | docs/TESTING.md@ad3758c:178; docs/TESTING.md@ad3758c:350 | C-01 |  | Python fault: a duplicated section and a cut line broke every program. |
| REQ-TS-0154 | A save writes only what that program changed (three-way merge): a secret saved by one program survives a save of another setting from an older copy. | docs/TESTING.md@ad3758c:178 | C-01 | S-11, S-10 | Python fault: the bot token vanished. |
| REQ-TS-0155 | A USB tethering device (RNDIS and similar) is recognised as tethering, not as wired Ethernet; the box keeps measuring on the cable; enabling the USB uplink takes an address by DHCP, keeps the default route on the cable and adds the source rule to its own routing table (Python: 152). | docs/TESTING.md@ad3758c:179 | N-10 | N-01, C-11 | (source names a specific site value; not copied) Source names the device model, operator and USB ids. |
| REQ-TS-0156 | A test asked for a given interface runs on that interface for every engine (bytes counted there); results are stored with the access kind (for example fwa) and the public address actually used. | docs/TESTING.md@ad3758c:179 | M-04 | M-06, M-07, M-08, M-10, C-11, R-01 | Python fault: the interface option was ignored. |
| REQ-TS-0157 | The RTR engine binds to the asked interface and its row has the client's public address. | docs/TESTING.md@ad3758c:179 | M-08 | M-04 | Source: RTR was much faster than the other engines over 5G, 'to check'; see Questions. |
| REQ-TS-0158 | USB as the only internet: when cable and Wi-Fi are gone, a default route through the hotspot is added within about 12 s with a lower priority (Python: metric 900); when cable or Wi-Fi come back it is removed within 10 s and the fallback-end event is written. | docs/TESTING.md@ad3758c:180 | N-10 | N-01, C-07 |  |
| REQ-TS-0159 | Events use only severities the schema accepts (a test inserts into the real schema). | docs/TESTING.md@ad3758c:180 | C-07 | C-02 | Python fault: fallback events were not stored. |
| REQ-TS-0160 | A static default route the box itself set on a link that went down is removed (routes with link down ignored), so it cannot beat Wi-Fi and USB. | docs/TESTING.md@ad3758c:180; docs/TESTING.md@ad3758c:336 | N-01 | N-10 | Python fault: the box lost its internet and remote access. |
| REQ-TS-0161 | Router mode with USB as WAN: LAN ports and the radio bridged, NAT to the USB link, DHCP to cable and Wi-Fi clients, which get confirmed internet (forwarded packets counted). | docs/TESTING.md@ad3758c:181 | N-10 | N-03, N-02, N-07 |  |
| REQ-TS-0162 | When the radio becomes the access point it drops its client lease and default route. | docs/TESTING.md@ad3758c:181 | N-07 | N-03 | Python fault: the box lost its own internet. |
| REQ-TS-0163 | The DHCP guard sends its probe as a raw frame from 0.0.0.0 and listens for all protocols, so it finds a foreign DHCP server on a plain and on a bridged port without false alarms. | docs/TESTING.md@ad3758c:181 | N-05 | N-06, P-07 | Python fault: the guard never saw the home router on the bridged port. |
| REQ-TS-0164 | Router mode drops the wired LAN ports' link for a few seconds after setup (Python: 4 s) so clients renew their lease. | docs/TESTING.md@ad3758c:181; docs/TESTING.md@ad3758c:335 | N-03 |  |  |
| REQ-TS-0165 | When a router appears on the port in router mode, the guard stops serving DHCP (Python: after 50 s) and the wired port becomes the WAN; when the USB device is pulled the mode rebuilds with the wired WAN. | docs/TESTING.md@ad3758c:181 | N-05 | N-03, N-04, N-10 |  |
| REQ-TS-0166 | Access point password: the shared factory password is refused by the access point and by the Web UI save, as are an empty password and one outside 8-63 characters; the config API reports only wifi_psk_set; the installer log has no password line; each box has its own valid password. | docs/TESTING.md@ad3758c:182 | N-07 | S-02…S-0n (API: Wi-Fi), P-11, P-03, C-04 |  |
| REQ-TS-0167 | nmb-cli covers the Web UI: a test fails for a Web UI endpoint without a CLI action, and the endpoints the UI really calls are compared with the catalog. | docs/TESTING.md@ad3758c:183; docs/TESTING.md@ad3758c:200 | S-10 | U-01 |  |
| REQ-TS-0168 | The installer restarts every service: each service's start time is later than the start of the install. | docs/TESTING.md@ad3758c:184; docs/TESTING.md@ad3758c:259 | P-11 | P-10 | Python fault: the bot ran code from two days earlier. |
| REQ-TS-0169 | No secrets in output: the install log, the config API, bot answers and the diagnostic bundle carry only presence flags; tests check secrets only for empty or factory value and never print them. | docs/TESTING.md@ad3758c:185 | C-04 | P-11, S-11, R-04, P-07 |  |
| REQ-TS-0170 | Every change is tested at the level it can break; what ran, where, and what was seen is written down; a change without a recorded test is not finished. | docs/TESTING.md@ad3758c:189 | P-13 | F-00 |  |
| REQ-TS-0171 | The gate runs the unit tests (must end OK), a syntax check of the Web UI script when it changed and of the installer script when it changed. | docs/TESTING.md@ad3758c:193-197 | F-00 | U-01, P-11 | Python commands; native equivalents are F-00's choice. |
| REQ-TS-0172 | New behaviour gets a test; the system (processes, /sys, sockets) is faked; tests pass on a normal PC without root, without network and without the box. | docs/TESTING.md@ad3758c:199 | F-06 | F-00 |  |
| REQ-TS-0173 | Translation check: every key in every language, the same placeholders and HTML as English, no Hungarian in the other languages; the key-difference check prints only zeros. | docs/TESTING.md@ad3758c:201-202 | C-06 | F-00 |  |
| REQ-TS-0174 | A change is installed on a real box the way users get it (release archive and installer), not by copying single files except a quick try that is declared. | docs/TESTING.md@ad3758c:206 | P-13 | P-11 |  |
| REQ-TS-0175 | After installing, all services are active and no NMB unit has failed. | docs/TESTING.md@ad3758c:207 | P-13 | P-10 |  |
| REQ-TS-0176 | The feature is run for real and its result checked where users see it: Web UI, CLI latest, the database read-only, the logs. | docs/TESTING.md@ad3758c:208 | P-13 |  |  |
| REQ-TS-0177 | The test measures what the feature claims; for network features interface byte counters are read before and after to prove which link the traffic took. | docs/TESTING.md@ad3758c:209 | P-13 | F-06 |  |
| REQ-TS-0178 | After the test the box is put back as it was (settings, test interfaces, rules), and the record says what was changed and undone. | docs/TESTING.md@ad3758c:210 | P-13 |  |  |
| REQ-TS-0179 | The live self-test checks the running box in about a minute; --full also drives the bot through every command and button in all 11 languages; single groups can be run; the report goes to a timestamped JSON file under the log folder; the exit code is 1 when a check failed. | docs/TESTING.md@ad3758c:214 | P-07 | S-11 |  |
| REQ-TS-0180 | Self-test group services: NMB units active, no crash traces in the last hour, network-daemon state file present. | docs/TESTING.md@ad3758c:218 | P-07 | P-10 |  |
| REQ-TS-0181 | Self-test group internet: default route, ping, DNS, HTTPS. | docs/TESTING.md@ad3758c:219 | P-07 | N-01 |  |
| REQ-TS-0182 | Self-test group config: Wi-Fi password not weak or factory, bot token set (never shown), config not world-readable, event severities fit the schema. | docs/TESTING.md@ad3758c:220 | P-07 | C-01, C-04, C-07 |  |
| REQ-TS-0183 | Self-test group web: all 11 languages complete and without Hungarian outside Hungarian; status, USB, storage and config endpoints; the config API returns no secrets. | docs/TESTING.md@ad3758c:221 | P-07 | C-06, S-01 |  |
| REQ-TS-0184 | Self-test group cli: the CLI menu, every screen, latest for every engine. | docs/TESTING.md@ad3758c:222 | P-07 | S-10 |  |
| REQ-TS-0185 | Self-test group usb: device kind, address, profile never default, source rule to its own table, USB default route only without other internet, ping and HTTPS over USB. | docs/TESTING.md@ad3758c:223 | P-07 | N-10 |  |
| REQ-TS-0186 | Self-test group router (router mode only): bridge members, DHCP server, NAT on the WAN, the box's own route; a simulated LAN client (namespace on a veth in the LAN bridge) gets a lease and reaches internet, DNS and HTTPS. | docs/TESTING.md@ad3758c:224 | P-07 | N-03 |  |
| REQ-TS-0187 | Self-test group dhcp-probe: the DHCP guard against a throwaway DHCP server is found on a plain port and on a bridge port, with no false alarm without it. | docs/TESTING.md@ad3758c:225 | P-07 | N-05 |  |
| REQ-TS-0188 | Self-test group speed: last 48 h per engine: errors with the most common messages, rate-limit pauses, rows without the client IP, single ping samples, median download against Ookla/Cloudflare on the same link. | docs/TESTING.md@ad3758c:226 | P-07 | M-04, M-05 |  |
| REQ-TS-0189 | Self-test group wifi-ap: radio bands (HT40/VHT per band) against the running AP config, refused AP starts. | docs/TESTING.md@ad3758c:227 | P-07 | N-07 |  |
| REQ-TS-0190 | Self-test group collector: a database copy uploaded to a local collector (or a given URL, token from an environment variable) delivers every speed test once, sending again adds nothing; the box's own upload position is not touched. | docs/TESTING.md@ad3758c:228 | P-07 | S-12, S-13 |  |
| REQ-TS-0191 | Self-test group telegram: runs only with --full or by name. | docs/TESTING.md@ad3758c:229 | P-07 | S-11 |  |
| REQ-TS-0192 | Self-test checks only build throwaway objects (namespaces, veth pairs, a test bridge, a database copy) and remove them; they never change the mode, the settings or the uplinks, and never print secrets or public addresses. | docs/TESTING.md@ad3758c:231 | P-07 | C-04 |  |
| REQ-TS-0193 | Guided USB tethering scenario: runs as a transient service because the link changes; the current instruction is in a run-time status file and the log; steps: speed tests bound to USB (every engine, bytes and data counter on USB), Wi-Fi shared on the LAN port, USB fallback, router mode with USB as WAN (forwarded packets counted in a throwaway firewall table), the DHCP guard when the cable goes back (limit 90 s), restore, optional unplug; --plan prints the steps, --steps runs some. | docs/TESTING.md@ad3758c:233 | P-07 | N-10, N-08, N-05 |  |
| REQ-TS-0194 | A guided scenario saves the settings first and puts them back at the end, after a crash, or by a separate restore service after 45 minutes; a restore command does it by hand. | docs/TESTING.md@ad3758c:233 | P-07 | C-01 |  |
| REQ-TS-0195 | Guided router-mode scenario (Ethernet WAN, a device on the box's Wi-Fi): router mode up, a real device on the Wi-Fi (fresh lease, packets counted out of the WAN), the Web UI from that device, the WAN cable out for a minute and back, restore to standalone; stopping the scenario service restores too. | docs/TESTING.md@ad3758c:235 | P-07 | N-03, N-07 |  |
| REQ-TS-0196 | Guided automatic-mode scenario (needs USB tethering for remote access, or an allow-offline switch): stays standalone a minute with the cable in the router, router mode not before the 90 s grace with the cable out, standalone again within about a minute with the cable back, PPPoE on the ONT with a saved login, restore. | docs/TESTING.md@ad3758c:237 | P-07 | N-04 |  |
| REQ-TS-0197 | Self-test group mode checks the running mode against the setting: the automatic decision and its reason, no access point outside router mode, a router address in standalone, DHCP served in router mode, no router answering on the WAN port while automatic mode keeps router mode. | docs/TESTING.md@ad3758c:239 | P-07 | N-04, N-03 |  |
| REQ-TS-0198 | Every self-test run ends with a 'To look at' list of failures and warnings; --md writes a Markdown table. | docs/TESTING.md@ad3758c:239 | P-07 |  |  |
| REQ-TS-0199 | A LAN client is simulated with a veth pair, one end in a network namespace; DHCP or routing is never tested on the port the tester is connected through. | docs/TESTING.md@ad3758c:243 | F-06 | P-13 |  |
| REQ-TS-0200 | A second DHCP server is simulated inside the namespace. | docs/TESTING.md@ad3758c:244 | F-06 | N-05 |  |
| REQ-TS-0201 | Simulated Wi-Fi uses mac80211_hwsim in a VM. | docs/TESTING.md@ad3758c:245 | F-06 |  |  |
| REQ-TS-0202 | Speed engines are tested one at a time; Ookla at most once an hour from the same public IP. | docs/TESTING.md@ad3758c:246 | P-13 | M-05, M-10 |  |
| REQ-TS-0203 | Per distribution: a clean cloud image installed from scratch, then installed again and rebooted; settings, admin password and measurements survive; the config is 0640 owned by the service user. | docs/TESTING.md@ad3758c:250 | P-11 | F-06 |  |
| REQ-TS-0204 | Network modes and the access point are tested in the VM test bed. | docs/TESTING.md@ad3758c:251 | F-06 | N-03, N-07 |  |
| REQ-TS-0205 | Test doubles match the real code: a renamed argument fails the test, and a fake returns the real function's shape (captured from a real run when in doubt). | docs/TESTING.md@ad3758c:255 | F-06 | F-00 | Python: patch with autospec=True; Go: fakes implement the real interface. |
| REQ-TS-0206 | After an install, and after a settings change a service reads only at start, every service's start time is newer than the change. | docs/TESTING.md@ad3758c:259 | P-11 | P-10, C-01 |  |
| REQ-TS-0207 | Long checks on the box run as a transient unit as the service user, not in the SSH session; the output file is not created beforehand. | docs/TESTING.md@ad3758c:263 | P-13 |  |  |
| REQ-TS-0208 | Every step of a long check has its own time limit and prints one line, so a hang shows where it is. | docs/TESTING.md@ad3758c:264 | P-13 | F-06 |  |
| REQ-TS-0209 | Never kill by a pattern that also matches the tester's own SSH command line. | docs/TESTING.md@ad3758c:265 | P-13 |  |  |
| REQ-TS-0210 | Before cutting a link the box is reached through, a way back is armed that does not depend on the tester, as a sleeping service and not a timer; a second path stays open. | docs/TESTING.md@ad3758c:266 | P-13 | P-07 |  |
| REQ-TS-0211 | Actions that change the line or the box (speed tests from a menu button, reboot, purge, updates) are skipped unless they are the test; a button label is mapped to its command first. | docs/TESTING.md@ad3758c:267 | P-13 | S-11 |  |
| REQ-TS-0212 | Bot tests use chat IDs that are not real chats; the record says which real messages were sent. | docs/TESTING.md@ad3758c:268 | P-13 | S-11 |  |
| REQ-TS-0213 | While testing no token, password or key is read, printed or copied, not even into a log; checks look only at empty or not, factory value or not, length; a lost secret is entered again by the owner. | docs/TESTING.md@ad3758c:272 | P-13 | C-04, P-07 |  |
| REQ-TS-0214 | Every record row gives where (box, VM, unit tests), the result (works, works with a limit, fails, not tested) and how (set-up, what was measured, numbers, what was found and fixed); what was not tried is written as not tried. | docs/TESTING.md@ad3758c:26; docs/TESTING.md@ad3758c:276 | P-13 |  |  |
| REQ-TS-0215 | The release checklist runs before every release on the box (and on one VM for installer changes); each line says what to do and what passes; results go into the record. | docs/TESTING.md@ad3758c:280 | P-13 | P-11 |  |
| REQ-TS-0216 | Checklist 1, unit tests: the unit tests and the script syntax checks all pass. | docs/TESTING.md@ad3758c:284 | P-13 | F-00 |  |
| REQ-TS-0217 | Checklist 2, upgrade install: exit 0; web, network daemon, probe (and bot if used) active and started after the install; no failed NMB unit; the installed source commit is the new one; settings, admin password, bot token and measurements kept; config readable and 0640 service user; no password or token in the install log. | docs/TESTING.md@ad3758c:285 | P-13 | P-11, C-04 |  |
| REQ-TS-0218 | Checklist 3, clean install: one VM per supported family (Debian/Ubuntu, Fedora/RHEL, Alpine) passes as in 2, the dashboard answers and a first RTR test runs. | docs/TESTING.md@ad3758c:286 | P-13 | P-11, M-08 |  |
| REQ-TS-0219 | Checklist 4, Ookla: one scheduled run at most hourly stores a row with download, upload, ping and the chosen server; if refused: one error row, a pause, no other engine's result in its place. | docs/TESTING.md@ad3758c:287 | P-13 | M-10, M-05 |  |
| REQ-TS-0220 | Checklist 5, RTR, Cloudflare, Fast.com: one manual run each stores a row under its own engine name with speeds plausible for the line; the Fast.com row has raw JSON with servers, samples and loaded latency. | docs/TESTING.md@ad3758c:288 | P-13 | M-06, M-07, M-08, S-10 |  |
| REQ-TS-0221 | Checklist 6, iPerf3: one run against the configured server stores a row named after the server (place and provider), not only the address. | docs/TESTING.md@ad3758c:289 | P-13 | M-09 |  |
| REQ-TS-0222 | Checklist 7, mixed run: schedule engine all for one cycle gives one row per engine with its own name; Ookla only if its hour is over. | docs/TESTING.md@ad3758c:290 | P-13 | M-04, M-05 |  |
| REQ-TS-0223 | Checklist 8, measurement link: with wifi set, byte counters show the traffic on the chosen link; auto uses Ethernet when a cable is in; set back to auto. | docs/TESTING.md@ad3758c:291 | P-13 | M-04, N-01 |  |
| REQ-TS-0224 | Checklist 9, Wi-Fi shared on a LAN port: the client gets an address and internet through Wi-Fi; with a second DHCP server the box stops its DHCP, address, NAT and rule within about 2 minutes; everything removed when turned off. | docs/TESTING.md@ad3758c:292 | P-13 | N-08, N-05 |  |
| REQ-TS-0225 | Checklist 10, wireless roles: setting the radio to main and then to guest in Ports & Drivers shows the right banner in the device explorer, and guest status no longer claims the main network. | docs/TESTING.md@ad3758c:293 | P-13 | N-17, N-08, U-02…U-0n (page: ports and drivers) |  |
| REQ-TS-0226 | Checklist 11, storage guard: forced run on the box deletes nothing and shows the target; on a VM with a filled disk rotated logs go first, measurements of the last 7 days are kept, an event is written. | docs/TESTING.md@ad3758c:294 | P-13 | C-08 |  |
| REQ-TS-0227 | Checklist 12, settings: saving from Web UI and CLI at the same time a few times, then one setting in each, keeps the config parsing, no duplicated section, both changes present, secrets unchanged (the bot still starts). | docs/TESTING.md@ad3758c:295 | P-13 | C-01 |  |
| REQ-TS-0228 | Checklist 13, languages: all 11 languages on the dashboard, settings tabs, logout page, the results page, the report with ?lang= and the bot's language command show no raw keys or placeholders, no broken layout, no Hungarian outside Hungarian. | docs/TESTING.md@ad3758c:296 | P-13 | C-06, U-01, S-11, R-03 |  |
| REQ-TS-0229 | Checklist 14, Web UI: every tab opened with the browser console open shows no JavaScript errors and loads its data. | docs/TESTING.md@ad3758c:297 | P-13 | U-01 |  |
| REQ-TS-0230 | Checklist 15, reports and exports: report, speed CSV, detailed CSV and ping CSV download and open; the new engines and iPerf3 names appear. | docs/TESTING.md@ad3758c:298 | P-13 | R-02, R-03 |  |
| REQ-TS-0231 | Checklist 16, bot (if configured): every menu button and the status, ping, report, tailscale, vpn, wifi, wifi scan, tasks, heal, backup, drivers and mode commands answer in the chat's language with real data (no 0.0 B, no 'no controllers', Tailscale state matching the system); no error text; one test arrives; kill on an own sleep, backup create and delete, vpn add and del with a test name, one test alert; everything created is removed; the stand-in client run gives 0 failures in all 11 languages. | docs/TESTING.md@ad3758c:299 | P-13 | S-11 |  |
| REQ-TS-0232 | Checklist 17, reboot: after a reboot and 3 minutes all services are up; settings, the last snapshot and the Ookla pause state survive. | docs/TESTING.md@ad3758c:300 | P-13 | P-10, C-02, M-05 |  |
| REQ-TS-0233 | Checklist 18, Wi-Fi access point: the factory password and a 6-character one are both refused; the config API has wifi_psk_set and no password; the AP does not start without a valid password. | docs/TESTING.md@ad3758c:301 | P-13 | N-07 |  |
| REQ-TS-0234 | Checklist 19, nmb-cli: the menu lists the new actions (storage save with keep days, storage cleanup, diagnostic bundle) and they work; a bad value is refused with a message. | docs/TESTING.md@ad3758c:302 | P-13 | S-10, C-08, R-04 |  |
| REQ-TS-0235 | Checklist 20, disk-space card: values load, save and come back after a reload; 'Check and clean up now' updates the free-space line. | docs/TESTING.md@ad3758c:303 | P-13 | U-02…U-0n (page: system), C-08 |  |
| REQ-TS-0236 | Checklist 21, services restarted: every NMB service's start time is later than the install. | docs/TESTING.md@ad3758c:304 | P-13 | P-11, P-10 |  |
| REQ-TS-0237 | Checklist 22, secrets: the install log and the config API show only flags such as wifi_psk_set, no values. | docs/TESTING.md@ad3758c:305 | P-13 | C-04, P-11 |  |
| REQ-TS-0238 | Checklist 23, logs after the run: no new errors in the NMB units' logs since the start except the ones the test caused on purpose. | docs/TESTING.md@ad3758c:306 | P-13 | C-03 |  |
| REQ-TS-0239 | Checklist 24, put back: test settings, interfaces, rules, test clients and test chat languages undone; the box is as before the checklist. | docs/TESTING.md@ad3758c:307 | P-13 | P-07 |  |
| REQ-TS-0240 | Must prove: Wi-Fi shared on a LAN port with a real cable to a laptop (no router) gives it an address of the sharing subnet and internet; with the cable then in the home router the box stops its DHCP within seconds and takes the router's address. | docs/TESTING.md@ad3758c:315 | N-08 | N-05 |  |
| REQ-TS-0241 | Must prove: Wi-Fi uplink with the box's own profile on a box not yet on Wi-Fi, then a speed test with measurement_link = wifi with the bytes on the Wi-Fi interface. | docs/TESTING.md@ad3758c:316 | N-08 | M-04 |  |
| REQ-TS-0242 | Must prove: guest and main Wi-Fi on two radios (a USB adapter as second radio), both banners in the device explorer. | docs/TESTING.md@ad3758c:317 | N-08 | N-17, N-07 |  |
| REQ-TS-0243 | Real radios as an access point: 2.4 GHz 802.11n and 5 GHz 802.11ac 80 MHz pass with real clients; Wi-Fi 6/7 and 6 GHz stay to prove on other hardware. | docs/TESTING.md@ad3758c:318 | N-07 | P-13 |  |
| REQ-TS-0244 | Both bands at once need a second radio (USB adapter) as a second access point; not built in Python. | docs/TESTING.md@ad3758c:319 | N-07 | N-17 | See Questions (scope). |
| REQ-TS-0245 | Must prove on a Raspberry Pi 5. | docs/TESTING.md@ad3758c:320 | P-13 | P-11 |  |
| REQ-TS-0246 | Nothing runs on routers yet; router targets follow the porting plan. | docs/TESTING.md@ad3758c:321 | NEW: X2-01 Step-2 router tiers | P-11 |  |
| REQ-TS-0247 | During an outage the bot backs off with growing pauses, not several requests a second. | docs/TESTING.md@ad3758c:327 | S-11 |  | DynDNS off/on part of this row is done (line 110). |
| REQ-TS-0248 | When only DNS is blocked the DNS fallback starts while IP works, and the normal DNS comes back within 15 minutes after the block ends. | docs/TESTING.md@ad3758c:327 | N-01 | M-12 | Python reference for the fallback not named in the plan; assumed N-01. |
| REQ-TS-0249 | Router mode on the Pi starts the access point without 'did not start' errors. | docs/TESTING.md@ad3758c:327 | N-07 | N-03 |  |
| REQ-TS-0250 | The RTR client UUID stays the same across RTR tests. | docs/TESTING.md@ad3758c:327 | M-08 |  |  |
| REQ-TS-0251 | A cable pull leaves no repeated 'dead default route' log lines. | docs/TESTING.md@ad3758c:327 | N-01 | C-03 |  |
| REQ-TS-0252 | Automatic mode: with the cable in the router it stays standalone; cable out for 90 s gives router mode with the access point while remote access stays up; on the ONT with a saved PPPoE login it gives router mode with PPPoE; back in the router it is standalone within about 30 s; a router reboot under 90 s keeps standalone. | docs/TESTING.md@ad3758c:328 | N-04 | N-03, N-05, N-15 |  |
| REQ-TS-0253 | On a box with several wired ports, automatic mode bridges the other ports to the WAN port and a PC behind them gets the router's address. | docs/TESTING.md@ad3758c:328 | N-06 | N-04 |  |
| REQ-TS-0254 | Self-test group ipv6: the LAN router's advertisement (lifetime, prefixes), the box's public IPv6 address and route, HTTPS over IPv6. | docs/TESTING.md@ad3758c:329 | P-07 | N-01 | (source names a specific site value; not copied) Source records a router model and its advertisement. |
| REQ-TS-0255 | When the box has a public IPv6 address, DynDNS also sets the AAAA record. | docs/TESTING.md@ad3758c:329 | N-16 |  |  |
| REQ-TS-0256 | IPv6 for the LAN in router mode on a DHCP WAN or a USB WAN (Python: built for a PPPoE WAN only; the LAN gets IPv4 only otherwise). | docs/TESTING.md@ad3758c:330 | N-11 | N-10 | See Questions (scope). |
| REQ-TS-0257 | After a refused PPPoE login the box waits 10 minutes before dialling again (log: 'dialing again in 600 s'). | docs/TESTING.md@ad3758c:331 | N-05 |  |  |
| REQ-TS-0258 | A long line outage (not only a restart) causes no test over the USB hotspot. | docs/TESTING.md@ad3758c:331 | M-04 | N-10, M-02 |  |
| REQ-TS-0259 | Must prove: a forwarded port with a real phone or PC, and UPnP from a real application. | docs/TESTING.md@ad3758c:331 | N-11 |  |  |
| REQ-TS-0260 | Settings from a USB stick: an import file on a FAT32 stick with a new bot token changes the token, a result file lists the changed key names, a config backup is kept, plugging in again changes nothing. | docs/TESTING.md@ad3758c:332 | P-03 | C-01, C-04 |  |
| REQ-TS-0261 | Router mode with the Ethernet WAN: a device on the box's Wi-Fi gets internet, Web UI and SSH; a WAN cable loss and return recover. | docs/TESTING.md@ad3758c:333 | N-03 | P-07, N-07 |  |
| REQ-TS-0262 | The Tailscale guard stops and restores Tailscale when remote access drops for 2 minutes. | docs/TESTING.md@ad3758c:333 | N-15 | P-01 |  |
| REQ-TS-0263 | Wi-Fi sharing on the LAN port with USB as upstream: when the box leaves its Wi-Fi the laptop keeps browsing over USB (client packets out of the USB link counted). | docs/TESTING.md@ad3758c:334 | N-08 | N-10, P-07 |  |
| REQ-TS-0264 | After router mode drops the LAN ports' link, a laptop with an old lease gets a router-mode address and browses within a minute without replugging. | docs/TESTING.md@ad3758c:335 | N-03 |  |  |
| REQ-TS-0265 | Dead default routes are removed: with the cable moved from the router to a laptop no default route stays on the wired port after one daemon loop; Wi-Fi or USB carries the box's traffic and remote access stays up (self-test internet: 'no dead default route'). | docs/TESTING.md@ad3758c:336 | N-01 | N-10, P-07 |  |
| REQ-TS-0266 | Wired route priority: the cable moved to a laptop leaves no default route on the wired port (priority behind the router already proven). | docs/TESTING.md@ad3758c:337 | N-01 |  |  |
| REQ-TS-0267 | Unplugging the USB device in router mode and in standalone leaves no USB route or rule; the mode rebuilds with the wired port. | docs/TESTING.md@ad3758c:338 | N-10 | N-03 |  |
| REQ-TS-0268 | Secrets restore: with the bot token cleared by hand in the config, the self-test config group fails naming it, the secrets-restore command brings it back and the bot answers. | docs/TESTING.md@ad3758c:339 | C-04 | P-07, S-11 |  |
| REQ-TS-0269 | An iPerf3 server that fails five tests in a row is left out for 6 hours (log line), then tried again. | docs/TESTING.md@ad3758c:340 | M-09 |  |  |
| REQ-TS-0270 | During an Ookla rate-limit pause the schedule stores skipped rows, not errors. | docs/TESTING.md@ad3758c:341 | M-05 | M-04, P-07 |  |
| REQ-TS-0271 | Must prove: the collector over the internet with HTTPS and a real certificate on another network. | docs/TESTING.md@ad3758c:342 | S-13 | S-12 |  |
| REQ-TS-0272 | Storage guard on a full disk (VM with 2-4 GB): within 10 minutes or on a forced run rotated logs go first, measurements of the last 7 days stay, an event is written. | docs/TESTING.md@ad3758c:348 | C-08 | F-06 |  |
| REQ-TS-0273 | The installer with less than 500 MB free stops with a clear message; on a disk under 8 GB the new config gets 2 x 2 MB logs. | docs/TESTING.md@ad3758c:349 | P-11 | C-03 |  |
| REQ-TS-0274 | Web UI and CLI saving in a loop for a few minutes keep the config readable (across processes, not only threads). | docs/TESTING.md@ad3758c:350 | C-01 |  |  |
| REQ-TS-0275 | Fast.com compared with RTR on a line other than the test line. | docs/TESTING.md@ad3758c:351 | M-07 | M-08 |  |
| REQ-TS-0276 | Long run of 24-72 hours with per-minute tests: memory, database size, log rotation and storage guard stay in bounds. | docs/TESTING.md@ad3758c:352 | NEW: P-15 Endurance and hostile-network campaign | P-13, C-02, C-03, C-08 |  |
| REQ-TS-0277 | Low memory for 24 hours on 512 MB. | docs/TESTING.md@ad3758c:353 | NEW: P-15 Endurance and hostile-network campaign | P-13 |  |
| REQ-TS-0278 | An older Python release's database and config are taken over by the native build. | docs/TESTING.md@ad3758c:354 | C-02 | C-01, P-11 | Python row: install 2.4.0, update with its config. |
| REQ-TS-0279 | A backup made by an older release restores. | docs/TESTING.md@ad3758c:354 | P-03 | C-02 |  |
| REQ-TS-0280 | WAN failover (two WAN NICs, pull one), load balancing, VLAN (a bridge with VLAN filtering) and link aggregation. | docs/TESTING.md@ad3758c:355 | N-06 | N-01, N-03 |  |
| REQ-TS-0281 | Hostile networks: IPv6-only WAN, CGNAT, broken DNS, captive portal, flapping link, clock jump. | docs/TESTING.md@ad3758c:356 | NEW: P-15 Endurance and hostile-network campaign | N-01, N-09, M-12, C-05, M-02 |  |
| REQ-TS-0282 | Must prove: the bot's heal apply with a real update, purge on the real data, and Wi-Fi connect to a new network with a password. | docs/TESTING.md@ad3758c:357 | S-11 | P-02, C-08, N-08 |  |
| REQ-TS-0283 | A backup over the bot's 50 MB file limit: the bot says so instead of failing quietly. | docs/TESTING.md@ad3758c:358; docs/TESTING.md@ad3758c:400 | S-11 | P-03 |  |
| REQ-TS-0284 | Web UI, CLI, bot and network daemon saving in turns for an hour: every value set stays. | docs/TESTING.md@ad3758c:359 | C-01 | S-10, S-11, N-03 |  |
| REQ-TS-0285 | SMTP against a local mock server, then a real mailbox. | docs/TESTING.md@ad3758c:360 | R-05 | F-06 |  |
| REQ-TS-0286 | Must prove: an update from the code host on a box. | docs/TESTING.md@ad3758c:361 | P-02 |  |  |
| REQ-TS-0287 | Web UI: all 11 languages read by native speakers (the six newer ones were machine-assisted), several users at once, wrong logins, Firefox, Chrome and a phone. | docs/TESTING.md@ad3758c:362 | U-01 | C-06, S-01 |  |
| REQ-TS-0288 | Must prove: firewalld ports. | docs/TESTING.md@ad3758c:363 | N-02 | P-11 |  |
| REQ-TS-0289 | Must prove: reboot on Debian 12 and Fedora. | docs/TESTING.md@ad3758c:363 | P-11 |  |  |
| REQ-TS-0290 | Must prove: Tailscale. | docs/TESTING.md@ad3758c:363 | N-15 |  |  |
| REQ-TS-0291 | Must prove: OONI. | docs/TESTING.md@ad3758c:363 | M-13 |  |  |
| REQ-TS-0292 | Must prove: a real DuckDNS domain. | docs/TESTING.md@ad3758c:363 | N-16 |  |  |
| REQ-TS-0293 | The native build supports systemd with nftables and OpenRC (the Python base requirement of Python 3.10 and only the standard library does not apply). | docs/TESTING.md@ad3758c:367 | P-10 | P-11, N-02 |  |
| REQ-TS-0294 | Raspberry Pi OS (arm64) Lite should work (Debian 13 based); to test on a Pi or Debian arm64 under QEMU. | docs/TESTING.md@ad3758c:371 | P-11 | F-06 |  |
| REQ-TS-0295 | Raspberry Pi OS (armhf, 32-bit): unknown; third-party engine builds have run only on arm64 and x86-64. | docs/TESTING.md@ad3758c:372 | P-11 | M-08, M-10 | Outside the step-1 targets amd64 and arm64. |
| REQ-TS-0296 | Armbian and DietPi should work (Debian-based); test on the board or their x86 images. | docs/TESTING.md@ad3758c:373 | P-11 |  |  |
| REQ-TS-0297 | Linux Mint 22 and Pop!_OS should work (Ubuntu 24.04 based); a desktop network manager may claim the ports; test in a VM. | docs/TESTING.md@ad3758c:374 | P-11 | N-03 |  |
| REQ-TS-0298 | A base without a supported service manager (Devuan, no systemd) is refused by the installer with a clear message. | docs/TESTING.md@ad3758c:375 | P-11 | P-10 |  |
| REQ-TS-0299 | Arch, Manjaro, openSUSE, Void and NixOS are not supported (other package managers or init systems). | docs/TESTING.md@ad3758c:376 | P-11 |  |  |
| REQ-TS-0300 | OpenWrt is not supported in step 1; planned in the router port. | docs/TESTING.md@ad3758c:377 | NEW: X2-01 Step-2 router tiers | P-11 |  |
| REQ-TS-0301 | NMB's own distribution is planned (porting plan). | docs/TESTING.md@ad3758c:378 | P-11 |  |  |
| REQ-TS-0302 | To fix, Ookla rate limits at an hourly schedule: Ookla runs as one client identity (not a new client each run), and manual and bot runs count against the hour too. | docs/TESTING.md@ad3758c:386 | M-05 | M-10, S-11, S-10 | Python still saw refusals every few hours; see Questions. |
| REQ-TS-0303 | To fix, SD card writes: the developer log must not grow about 10 MB a day on flash; flash writes are limited (the boot-partition diagnostic snapshot at most once an hour). | docs/TESTING.md@ad3758c:387 | C-03 | C-08, R-04 |  |
| REQ-TS-0304 | To fix, data used by the schedule: a metered link (USB) without a data limit is warned about; tests in router mode never fall back to USB. | docs/TESTING.md@ad3758c:388 | M-05 | N-10, M-04 | See Questions (default USB cap). |
| REQ-TS-0305 | To fix, putting IPv4 DNS first must not reapply the connection profile in a way that resets open connections (Python reset the bot's connection once). | docs/TESTING.md@ad3758c:389 | N-01 | S-11 |  |
| REQ-TS-0306 | To fix, every PPPoE dial logs a harmless permission warning for a per-user options file; the native PPPoE dial logs no such warning. | docs/TESTING.md@ad3758c:390 | N-05 | P-04 |  |
| REQ-TS-0307 | Known limit: with one Wi-Fi radio the access point uses one band at a time (2.4 or 5 GHz; auto picks 5 GHz); both bands at once need a second radio. | docs/TESTING.md@ad3758c:394 | N-07 | P-12 |  |
| REQ-TS-0308 | Known limit: the Pi 4's onboard Wi-Fi as an access point carries about 85-95 Mbit/s on 5 GHz; Wi-Fi results show the radio, not the line, and the docs say to measure the line over the cable. | docs/TESTING.md@ad3758c:395 | C-11 | N-07, R-01, P-12 |  |
| REQ-TS-0309 | Known limit: an ISP may allow one PPPoE session per login, so another router and the box cannot be online on the same login at once, and a session not hung up (no PADT) blocks the next login for minutes; the box always hangs up cleanly. | docs/TESTING.md@ad3758c:396 | N-05 | P-12 | (source names a specific site value; not copied) Source names the ISP. |
| REQ-TS-0310 | Known limit: Ookla blocks a public IP that tests too often and refused attempts prolong it; the box pauses Ookla (30 min, doubling up to 4 h) and does not change the IP to get around it; Ookla runs at most hourly. | docs/TESTING.md@ad3758c:397 | M-05 | M-10, P-12 |  |
| REQ-TS-0311 | Known limit: up to 15 minutes of results can be lost on a power cut, by design (the live database is in RAM). | docs/TESTING.md@ad3758c:398 | C-02 | P-12 |  |
| REQ-TS-0312 | The VM scripts and bot check scripts used for the record are part of the native repository (the Python limit is not carried over). | docs/TESTING.md@ad3758c:399 | F-06 | P-13 | Interpretation of a Python known limit; see Questions. |
| REQ-TS-0313 | Known limit: the bot can send a backup only up to the messenger's 50 MB file limit for bots. | docs/TESTING.md@ad3758c:400 | S-11 | P-12 |  |
| REQ-TS-0314 | Known limit: a box that had the old factory Wi-Fi password and no new one gets no access point until a password is set. | docs/TESTING.md@ad3758c:401 | N-07 | P-03, P-12 |  |

## Coverage

Blank lines are listed with the element before them. Every line 1-401 is in exactly one row.

| Lines | Kind | Ids or context |
| :--- | :--- | :--- |
| 1-2 | heading (line 2 blank) | context: document title |
| 3-4 | paragraph (line 4 blank) | REQ-TS-0001 |
| 5-6 | paragraph (line 6 blank) | context: dates of the last Python sweep and update; history, nothing to build |
| 7-8 | paragraph (line 8 blank) | REQ-TS-0002; REQ-TS-0003 |
| 9-10 | paragraph (line 10 blank) | REQ-TS-0002 |
| 11-12 | heading (line 12 blank) | context: section heading |
| 13-14 | paragraph (line 14 blank) | REQ-TS-0004; REQ-TS-0005; REQ-TS-0006 |
| 15-16 | table row | context: table header |
| 17 | table row | REQ-TS-0007 |
| 18 | table row | REQ-TS-0008 |
| 19 | table row | REQ-TS-0009 |
| 20-21 | table row (line 21 blank) | REQ-TS-0010 |
| 22-23 | paragraph (line 23 blank) | REQ-TS-0011 |
| 24-25 | heading (line 25 blank) | context: section heading |
| 26-27 | paragraph (line 27 blank) | context: legend of result symbols; the recording rule is REQ line 276 |
| 28-29 | heading (line 29 blank) | context: section heading |
| 30-31 | table row | context: table header |
| 32 | table row | REQ-TS-0012 |
| 33 | table row | REQ-TS-0013; REQ-TS-0014 |
| 34 | table row | REQ-TS-0015 |
| 35 | table row | REQ-TS-0016 |
| 36 | table row | REQ-TS-0017; REQ-TS-0018 |
| 37 | table row | REQ-TS-0019 |
| 38 | table row | REQ-TS-0020 |
| 39 | table row | REQ-TS-0021; REQ-TS-0022; REQ-TS-0023 |
| 40 | table row | REQ-TS-0024 |
| 41 | table row | REQ-TS-0025; REQ-TS-0026 |
| 42 | table row | REQ-TS-0027; REQ-TS-0028 |
| 43-44 | table row (line 44 blank) | REQ-TS-0029 |
| 45-46 | heading (line 46 blank) | context: section heading |
| 47-48 | table row | context: table header |
| 49 | table row | REQ-TS-0030; REQ-TS-0031 |
| 50 | table row | REQ-TS-0032 |
| 51 | table row | REQ-TS-0033 |
| 52 | table row | REQ-TS-0034; REQ-TS-0035; REQ-TS-0036; REQ-TS-0037 |
| 53 | table row | REQ-TS-0038 |
| 54 | table row | REQ-TS-0039 |
| 55 | table row | REQ-TS-0040 |
| 56 | table row | REQ-TS-0041 |
| 57 | table row | REQ-TS-0042; REQ-TS-0043 |
| 58 | table row | REQ-TS-0044 |
| 59 | table row | REQ-TS-0045 |
| 60-61 | table row (line 61 blank) | REQ-TS-0046 |
| 62-63 | heading (line 63 blank) | context: section heading |
| 64-65 | table row | context: table header |
| 66 | table row | REQ-TS-0047 |
| 67 | table row | REQ-TS-0048 |
| 68 | table row | REQ-TS-0049 |
| 69 | table row | REQ-TS-0050 |
| 70 | table row | REQ-TS-0051 |
| 71-72 | table row (line 72 blank) | REQ-TS-0052 |
| 73-74 | heading (line 74 blank) | context: section heading; (source names a specific site value; not copied) |
| 75-76 | paragraph (line 76 blank) | REQ-TS-0053 |
| 77-78 | table row | context: table header |
| 79 | table row | REQ-TS-0054 |
| 80 | table row | REQ-TS-0055 |
| 81 | table row | REQ-TS-0056 |
| 82 | table row | REQ-TS-0057; REQ-TS-0058 |
| 83 | table row | REQ-TS-0059 |
| 84 | table row | REQ-TS-0060 |
| 85 | table row | REQ-TS-0061 |
| 86 | table row | REQ-TS-0062 |
| 87 | table row | REQ-TS-0063 |
| 88 | table row | REQ-TS-0064 |
| 89 | table row | REQ-TS-0065 |
| 90 | table row | REQ-TS-0066 |
| 91 | table row | REQ-TS-0067 |
| 92 | table row | REQ-TS-0068 |
| 93 | table row | REQ-TS-0069 |
| 94-95 | table row (line 95 blank) | REQ-TS-0070 |
| 96-97 | heading (line 97 blank) | context: section heading |
| 98-99 | table row | context: table header |
| 100 | table row | REQ-TS-0071 |
| 101 | table row | REQ-TS-0072 |
| 102 | table row | REQ-TS-0073 |
| 103 | table row | REQ-TS-0074; REQ-TS-0075 |
| 104 | table row | REQ-TS-0076 |
| 105 | table row | REQ-TS-0077 |
| 106 | table row | REQ-TS-0078 |
| 107 | table row | REQ-TS-0079 |
| 108 | table row | REQ-TS-0080 |
| 109 | table row | REQ-TS-0081 |
| 110 | table row | REQ-TS-0082; REQ-TS-0083; REQ-TS-0084; REQ-TS-0085; REQ-TS-0086; REQ-TS-0087 |
| 111 | table row | REQ-TS-0088 |
| 112 | table row | REQ-TS-0089 |
| 113 | table row | REQ-TS-0090 |
| 114 | table row | REQ-TS-0091 |
| 115 | table row | REQ-TS-0092 |
| 116 | table row | REQ-TS-0093 |
| 117 | table row | REQ-TS-0094; REQ-TS-0095; REQ-TS-0096; REQ-TS-0097 |
| 118 | table row | REQ-TS-0098; REQ-TS-0099; REQ-TS-0100 |
| 119 | table row | REQ-TS-0101 |
| 120-121 | table row (line 121 blank) | REQ-TS-0102 |
| 122-123 | heading (line 123 blank) | context: section heading |
| 124-125 | table row | context: table header |
| 126 | table row | REQ-TS-0103 |
| 127 | table row | REQ-TS-0104 |
| 128 | table row | REQ-TS-0105 |
| 129 | table row | REQ-TS-0106 |
| 130 | table row | REQ-TS-0107 |
| 131 | table row | REQ-TS-0108 |
| 132-133 | table row (line 133 blank) | REQ-TS-0109 |
| 134-135 | heading (line 135 blank) | context: section heading |
| 136-137 | table row | context: table header |
| 138 | table row | REQ-TS-0110 |
| 139 | table row | REQ-TS-0111 |
| 140 | table row | REQ-TS-0112 |
| 141 | table row | REQ-TS-0113 |
| 142-143 | table row (line 143 blank) | REQ-TS-0114 |
| 144-145 | heading (line 145 blank) | context: section heading |
| 146-147 | table row | context: table header |
| 148 | table row | REQ-TS-0115 |
| 149 | table row | REQ-TS-0116 |
| 150 | table row | REQ-TS-0117 |
| 151 | table row | REQ-TS-0118 |
| 152-153 | table row (line 153 blank) | REQ-TS-0119 |
| 154-155 | heading (line 155 blank) | context: section heading |
| 156-157 | table row | context: table header |
| 158 | table row | REQ-TS-0120 |
| 159 | table row | REQ-TS-0121; REQ-TS-0122; REQ-TS-0123 |
| 160 | table row | REQ-TS-0124 |
| 161 | table row | REQ-TS-0125 |
| 162 | table row | REQ-TS-0126 |
| 163 | table row | REQ-TS-0127 |
| 164 | table row | REQ-TS-0128 |
| 165 | table row | REQ-TS-0129; REQ-TS-0130; REQ-TS-0131 |
| 166 | table row | REQ-TS-0132 |
| 167 | table row | REQ-TS-0133; REQ-TS-0134; REQ-TS-0135 |
| 168 | table row | REQ-TS-0136; REQ-TS-0137; REQ-TS-0138 |
| 169 | table row | REQ-TS-0139 |
| 170 | table row | REQ-TS-0140; REQ-TS-0141; REQ-TS-0142 |
| 171 | table row | REQ-TS-0143 |
| 172 | table row | REQ-TS-0144; REQ-TS-0145; REQ-TS-0146 |
| 173 | table row | REQ-TS-0147 |
| 174 | table row | REQ-TS-0148 |
| 175 | table row | REQ-TS-0149 |
| 176 | table row | REQ-TS-0150; REQ-TS-0151 |
| 177 | table row | REQ-TS-0152 |
| 178 | table row | REQ-TS-0153; REQ-TS-0154 |
| 179 | table row | REQ-TS-0155; REQ-TS-0156; REQ-TS-0157 |
| 180 | table row | REQ-TS-0158; REQ-TS-0159; REQ-TS-0160 |
| 181 | table row | REQ-TS-0161; REQ-TS-0162; REQ-TS-0163; REQ-TS-0164; REQ-TS-0165 |
| 182 | table row | REQ-TS-0166 |
| 183 | table row | REQ-TS-0167 |
| 184 | table row | REQ-TS-0168 |
| 185-186 | table row (line 186 blank) | REQ-TS-0169 |
| 187-188 | heading (line 188 blank) | context: section heading |
| 189-190 | paragraph (line 190 blank) | REQ-TS-0170 |
| 191-192 | heading (line 192 blank) | context: section heading |
| 193-198 | code block (line 198 blank) | REQ-TS-0171 |
| 199 | bullet | REQ-TS-0172 |
| 200 | bullet | REQ-TS-0167 |
| 201-203 | bullet (line 203 blank) | REQ-TS-0173 |
| 204-205 | heading (line 205 blank) | context: section heading |
| 206 | bullet | REQ-TS-0174 |
| 207 | bullet | REQ-TS-0175 |
| 208 | bullet | REQ-TS-0176 |
| 209 | bullet | REQ-TS-0177 |
| 210-211 | bullet (line 211 blank) | REQ-TS-0178 |
| 212-213 | heading (line 213 blank) | context: section heading |
| 214-215 | paragraph (line 215 blank) | REQ-TS-0179 |
| 216-217 | table row | context: table header |
| 218 | table row | REQ-TS-0180 |
| 219 | table row | REQ-TS-0181 |
| 220 | table row | REQ-TS-0182 |
| 221 | table row | REQ-TS-0183 |
| 222 | table row | REQ-TS-0184 |
| 223 | table row | REQ-TS-0185 |
| 224 | table row | REQ-TS-0186 |
| 225 | table row | REQ-TS-0187 |
| 226 | table row | REQ-TS-0188 |
| 227 | table row | REQ-TS-0189 |
| 228 | table row | REQ-TS-0190 |
| 229-230 | table row (line 230 blank) | REQ-TS-0191 |
| 231-232 | paragraph (line 232 blank) | REQ-TS-0192 |
| 233-234 | paragraph (line 234 blank) | REQ-TS-0193; REQ-TS-0194 |
| 235-236 | paragraph (line 236 blank) | REQ-TS-0195 |
| 237-238 | paragraph (line 238 blank) | REQ-TS-0196 |
| 239-240 | paragraph (line 240 blank) | REQ-TS-0197; REQ-TS-0198 |
| 241-242 | heading (line 242 blank) | context: section heading |
| 243 | bullet | REQ-TS-0199 |
| 244 | bullet | REQ-TS-0200 |
| 245 | bullet | REQ-TS-0201; REQ-TS-0011 |
| 246-247 | bullet (line 247 blank) | REQ-TS-0202 |
| 248-249 | heading (line 249 blank) | context: section heading |
| 250 | bullet | REQ-TS-0203; REQ-TS-0009; REQ-TS-0149 |
| 251-252 | bullet (line 252 blank) | REQ-TS-0204; REQ-TS-0008 |
| 253-254 | heading (line 254 blank) | context: section heading |
| 255-256 | paragraph (line 256 blank) | REQ-TS-0205 |
| 257-258 | heading (line 258 blank) | context: section heading |
| 259-260 | paragraph (line 260 blank) | REQ-TS-0206; REQ-TS-0168 |
| 261-262 | heading (line 262 blank) | context: section heading |
| 263 | bullet | REQ-TS-0207 |
| 264 | bullet | REQ-TS-0208 |
| 265 | bullet | REQ-TS-0209 |
| 266 | bullet | REQ-TS-0210 |
| 267 | bullet | REQ-TS-0211 |
| 268-269 | bullet (line 269 blank) | REQ-TS-0212 |
| 270-271 | heading (line 271 blank) | context: section heading |
| 272-273 | paragraph (line 273 blank) | REQ-TS-0213 |
| 274-275 | heading (line 275 blank) | context: section heading |
| 276-277 | paragraph (line 277 blank) | REQ-TS-0214; REQ-TS-0001 |
| 278-279 | heading (line 279 blank) | context: section heading |
| 280-281 | paragraph (line 281 blank) | REQ-TS-0215 |
| 282-283 | table row | context: table header |
| 284 | table row | REQ-TS-0216 |
| 285 | table row | REQ-TS-0217 |
| 286 | table row | REQ-TS-0218 |
| 287 | table row | REQ-TS-0219 |
| 288 | table row | REQ-TS-0220 |
| 289 | table row | REQ-TS-0221 |
| 290 | table row | REQ-TS-0222 |
| 291 | table row | REQ-TS-0223 |
| 292 | table row | REQ-TS-0224 |
| 293 | table row | REQ-TS-0225 |
| 294 | table row | REQ-TS-0226 |
| 295 | table row | REQ-TS-0227 |
| 296 | table row | REQ-TS-0228 |
| 297 | table row | REQ-TS-0229 |
| 298 | table row | REQ-TS-0230 |
| 299 | table row | REQ-TS-0231 |
| 300 | table row | REQ-TS-0232 |
| 301 | table row | REQ-TS-0233 |
| 302 | table row | REQ-TS-0234 |
| 303 | table row | REQ-TS-0235 |
| 304 | table row | REQ-TS-0236 |
| 305 | table row | REQ-TS-0237 |
| 306 | table row | REQ-TS-0238 |
| 307-308 | table row (line 308 blank) | REQ-TS-0239 |
| 309-310 | heading (line 310 blank) | context: section heading |
| 311-312 | heading (line 312 blank) | context: section heading |
| 313-314 | table row | context: table header |
| 315 | table row | REQ-TS-0240 |
| 316 | table row | REQ-TS-0241 |
| 317 | table row | REQ-TS-0242 |
| 318 | table row | REQ-TS-0243 |
| 319 | table row | REQ-TS-0244 |
| 320 | table row | REQ-TS-0245 |
| 321-322 | table row (line 322 blank) | REQ-TS-0246 |
| 323-324 | heading (line 324 blank) | context: section heading |
| 325-326 | table row | context: table header |
| 327 | table row | REQ-TS-0247; REQ-TS-0248; REQ-TS-0249; REQ-TS-0250; REQ-TS-0251 |
| 328 | table row | REQ-TS-0252; REQ-TS-0253 |
| 329 | table row | REQ-TS-0254; REQ-TS-0255 |
| 330 | table row | REQ-TS-0256 |
| 331 | table row | REQ-TS-0257; REQ-TS-0258; REQ-TS-0259 |
| 332 | table row | REQ-TS-0260 |
| 333 | table row | REQ-TS-0261; REQ-TS-0262 |
| 334 | table row | REQ-TS-0263 |
| 335 | table row | REQ-TS-0264 |
| 336 | table row | REQ-TS-0265; REQ-TS-0160 |
| 337 | table row | REQ-TS-0266 |
| 338 | table row | REQ-TS-0267 |
| 339 | table row | REQ-TS-0268 |
| 340 | table row | REQ-TS-0269 |
| 341 | table row | REQ-TS-0270; REQ-TS-0031 |
| 342-343 | table row (line 343 blank) | REQ-TS-0271 |
| 344-345 | heading (line 345 blank) | context: section heading |
| 346-347 | table row | context: table header |
| 348 | table row | REQ-TS-0272 |
| 349 | table row | REQ-TS-0273 |
| 350 | table row | REQ-TS-0274; REQ-TS-0153 |
| 351 | table row | REQ-TS-0275 |
| 352 | table row | REQ-TS-0276 |
| 353 | table row | REQ-TS-0277 |
| 354 | table row | REQ-TS-0278; REQ-TS-0279 |
| 355 | table row | REQ-TS-0280 |
| 356 | table row | REQ-TS-0281 |
| 357 | table row | REQ-TS-0282 |
| 358 | table row | REQ-TS-0283 |
| 359 | table row | REQ-TS-0284 |
| 360 | table row | REQ-TS-0285 |
| 361 | table row | REQ-TS-0286 |
| 362 | table row | REQ-TS-0287 |
| 363-364 | table row (line 364 blank) | REQ-TS-0288; REQ-TS-0289; REQ-TS-0290; REQ-TS-0291; REQ-TS-0292 |
| 365-366 | heading (line 366 blank) | context: section heading |
| 367-368 | paragraph (line 368 blank) | REQ-TS-0293 |
| 369-370 | table row | context: table header |
| 371 | table row | REQ-TS-0294 |
| 372 | table row | REQ-TS-0295 |
| 373 | table row | REQ-TS-0296 |
| 374 | table row | REQ-TS-0297 |
| 375 | table row | REQ-TS-0298 |
| 376 | table row | REQ-TS-0299 |
| 377 | table row | REQ-TS-0300 |
| 378-379 | table row (line 379 blank) | REQ-TS-0301 |
| 380-381 | heading (line 381 blank) | context: section heading |
| 382-383 | paragraph (line 383 blank) | context: introduction of the To fix table; the log review is summarised in the CHANGELOG (covered by the changelog part) |
| 384-385 | table row | context: table header |
| 386 | table row | REQ-TS-0302 |
| 387 | table row | REQ-TS-0303 |
| 388 | table row | REQ-TS-0304 |
| 389 | table row | REQ-TS-0305 |
| 390-391 | table row (line 391 blank) | REQ-TS-0306 |
| 392-393 | heading (line 393 blank) | context: section heading |
| 394 | bullet | REQ-TS-0307 |
| 395 | bullet | REQ-TS-0308 |
| 396 | bullet | REQ-TS-0309 |
| 397 | bullet | REQ-TS-0310; REQ-TS-0031 |
| 398 | bullet | REQ-TS-0311 |
| 399 | bullet | REQ-TS-0312 |
| 400 | bullet | REQ-TS-0283; REQ-TS-0313 |
| 401 | bullet | REQ-TS-0314 |

## Proposed package changes

1. **NEW: P-15 Endurance and hostile-network campaign** (REQ-TS ids at lines 352, 353, 356). The long run (24-72 h), the 24 h low-memory run and the hostile-network suite (IPv6-only WAN, CGNAT, broken DNS, captive portal, flapping link, clock jump) need days of box or VM time and touch C-02, C-03, C-05, C-08, M-02, M-12, N-01 and N-09; inside P-13 they would make the release proof too big for one sitting (size rule, work plan line 100) and block it on multi-day runs. P-13 keeps the release checklist, the record rules and the Python-against-native comparison; P-15 runs before P-13 and hands it its results.
2. **Split P-07 into P-07a Live self-test groups and P-07b Guided scenarios.** The Python reference `nmb/selftest/` is 2,837 lines (`live.py` 1,605; `scenario.py` 593, `scenario_pppoe.py` 288, `scenario_auto.py` 187, `scenario_router.py` 164), above the size rule. P-07a owns the group rows (lines 214, 218-229, 231, 239 groups, 329 ipv6 group); P-07b owns the scenarios with their restore safety (lines 75, 233, 235, 237). This part writes both as P-07 until the split is accepted.
3. **Use NEW: X2-01 Step-2 router tiers** (proposed in `roadmap-1-507.md`) for the two rows that point to router targets (lines 321, 377), so they leave step 1 with an owner and do not fall into P-11.
4. **One owner for `netd/wifi_share.py`.** This part gives the DHCP guard rows (lines 52, 181, 225) to N-05 and the sharing rows to N-08, as the work plan's table says; the file itself must still get one owner (same point as made in `workplan-architecture-devsetup.md`).
5. **Fix the S-0n and U-0n numbers.** Rows here name areas (status, speed, network, Wi-Fi, VPN, system) and pages (settings, VPN, system, ports and drivers); F-02 must give them ids before the register check runs.

## Questions for the owner

1. **Wi-Fi client without NetworkManager** (line 142, REQ for that row): Python fails there by design. Should the native Wi-Fi client work without NetworkManager (own nl80211/WPA control), or keep the dependency and only report it clearly?
2. **Default data cap on a metered link** (line 388): the source says a USB limit is advisable but the box ships with 0 (no limit). Should the native build set a default daily cap on USB/mobile uplinks, or only warn?
3. **Manual and bot Ookla runs against the hour** (line 386): counting them means a manual Ookla run inside the hour is refused or deferred. Intended?
4. **IPv6 for the LAN on a DHCP or USB WAN** (line 330) and **dual-band with a second radio** (line 319): neither is built in Python. In step 1 of the native build, or recorded as a limit?
5. **RTR over 5G** (line 179): RTR measured far above the other engines over the mobile hotspot, marked 'to check' and never settled. Fault to find in the native M-08 brief, or accepted as the engine's behaviour (multi-stream, other server)?
6. **Test scripts in the repository** (line 399): this part turns the Python known limit (VM and bot check scripts kept outside the repository) into a requirement that the native scripts live in the native repository (F-06). Confirm.
