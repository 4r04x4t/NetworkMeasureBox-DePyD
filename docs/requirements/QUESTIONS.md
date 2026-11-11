# Requirements register: questions, triaged

The parts asked 44 questions; de-duplicated they are Q-01 to Q-39 (ids are permanent and were kept); Q-40 was added in
round 4 from a checker's finding, so there are 40 below. In round 3 (2026-11-11) each was sorted into one of three sections:

1. **Answered by existing decisions**: the owner's decisions in `docs/NATIVE_WORK_PLAN.md@ad3758c` section 7, the owner's
   decisions of 2026-11-11 (given to F-02 by the orchestrator, to be written into the roadmap by F-01), or the roadmap's own text.
2. **Developer's choices**: the owner left the matter to the developer (work plan 7 point 1), or the choice follows from a
   Quality Bar point (QB) or another roadmap rule; the rule is named.
3. **For the owner**: real owner decisions (product scope, security trade-offs, money, legal), each as one question with
   options and a recommendation.

Each entry names the rows it concerns and the packages that need it. Rows settled by sections 1 and 2 were corrected in their
parts (round 3 list: `PACKAGE_CHANGES.md` section I); rows of section 3 stay as written and point here. Round 4 (the
checkers' findings, `PACKAGE_CHANGES.md` section K) moved Q-01, Q-16 and Q-17 to the owner, Q-09 and Q-14 to the
developer's choices, corrected Q-08, Q-12, Q-25 and Q-36, and added Q-40; the rows of Q-01, Q-16 and Q-17 went back to
their source wording. On 2026-11-11 the owner answered the six questions of section 3 (Q-01, Q-05, Q-12, Q-15, Q-16, Q-17): they are in a
fourth section, *Answered by the owner*, with the rows they settled (`PACKAGE_CHANGES.md` section L), and no question is
open. Answers go to F-01, which writes them into the roadmap.

Part names: RM1 `roadmap-1-507.md`, RM2 `roadmap-508-1034.md`, PP `porting-plan.md`, WP `workplan-architecture-devsetup.md`,
TS `testing.md`, RE `readme.md`, CL `changelog.md`, PY `python-files.md`. `contributing.md` and `other-docs.md` asked no
question of their own (their rows follow Q-09 and the owner's decisions of 2026-11-11).

## Answered by existing decisions

**Q-06 Ubuntu 22.04 / polkit 0.105: service user may manage any unit?** (REQ-RE-0288; P-06) Answered in round 1: QB 6 and 7,
REQ-TS-0018, REQ-CL-0365: the native privilege rule is an allow-list of NMB's own units on every base.

**Q-07 Alpine without the public-address guard?** (REQ-RE-0290; N-02) Answered in round 1: QB 7 (nothing listens on the WAN
side unless the owner turned it on, `docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:587`), REQ-TS-0129, REQ-CL-0205: the guard runs
on every base, Alpine/OpenRC included. Alpine is a supported distribution (Q-05, answered by the owner on 2026-11-11); the base of the own image is I-09's.

**Q-11 Third-party engines in-process or sandboxed?** (RM2 2; REQ-RM-1267, -1350, -1417, -1418, -1459; P-04, M-13, M-09)
Answer: sandboxed, in a separate process. Roadmap 6B.5 (owner, 2026-11-10, newer than Phase 3.2 and section 7) lists
`ooniprobe` and iPerf3 among third-party code that runs in the stricter sandbox, and "third-party engines and upload parsers
cannot be taken out of their sandbox". OONI's engine is shipped as upstream's own program with its licence and run under that
profile; as a separate program its GPL-3.0 does not decide the licence of `nmbd` (work plan 7 point 9). An iPerf3 client
that NMB writes itself is NMB code under the service sandbox; `libiperf` (C tier) is third-party and runs apart.
Ordinary libraries (netlink, DNS parsing, SQLite) are dependencies under the licence and vulnerability checks of QB 7 and 6B.3A.

**Q-18 Login method: HTTP Basic Auth or session login?** (RM1 1; REQ-RM-0033; S-01) Answer: session login with CSRF protection.
Work plan 7 point 1 leaves the API to the developer, and QB 7 asks for CSRF protection, secure cookies and a rate-limited
login, which Basic Auth does not give. The Basic Auth line of roadmap 4.4 is out of date; REQ-RM-0033 restated.

**Q-19 Installer warning about the factory login?** (CL 2; REQ-CL-0247; P-11) Answer: no warning is needed. The owner decided
on 2026-11-10 that every native image gets a unique password per device (QB 7), so the native build has no factory login to
warn about. (The owner's own Python box keeps its login by choice; that is not the native build.) REQ-CL-0247 restated.

**Q-28 Manual and bot Ookla runs against the hourly limit?** (TS 3; REQ-TS-0302; M-05, M-10) Answer (owner, 2026-11-11): tests a
person starts obey the refusal pause but not the hourly limit; scheduled runs keep it. REQ-TS-0302 restated.

**Q-29 Default data cap on a metered link?** (TS 2; REQ-TS-0304; M-05, N-10) Answer (owner, 2026-11-11): no default cap;
scheduled tests on a metered link wait until the user has set one, and the box says so. REQ-TS-0304 restated.

**Q-33 Wi-Fi client without NetworkManager?** (TS 1; REQ-TS-0114; N-07, N-08) Answer: the native client works without it. The
own image holds "no NetworkManager or netplan" (roadmap 6A) and the Wi-Fi manager is a direct client of `wpa_supplicant`'s
control socket (section 7, REQ-RM-1436). Where another network manager holds the radio, or `wpa_supplicant` is missing, the
feature is reported as unavailable with the reason. REQ-TS-0114 restated.

**Q-35 No PPPoE on the TL-WR740N?** (PP 3; REQ-PP-0061; X2-01) Answer: roadmap 6B.8 point 6: what a target cannot hold is
listed in its parity table and shown in the UI as not available, never faked. The target is not scheduled yet (*Order of
Work* step 3: the other routers come in an order the owner sets then); the image recipe is checked again at that time.

**Q-36 `ip`, `nft` and `iw` in the image?** (RM2 3; REQ-RM-1263, -1277, -1414, -1425; P-11) Answer: both texts hold. 6A lists
the tools as image contents (for people and recovery); section 7 and Phase 2.1 say that `nmbd` itself never calls
`/sbin/ip` and `ethtool` (section 7 also `/sbin/route` and `iw`). Neither names `nft`: that `nmbd` never calls `nft` either and
drives nftables through its netlink interface is a reading, a developer's choice under QB 7 (no shell) and REQ-RM-0006 (no
routine spawning of tools). Round 4: corrected; the round-3 text cited `nft` as stated in section 7.

**Q-37 Languages of the bot and the dossier?** (RM1 2; REQ-RM-0105, -0107; S-11, R-03, C-06) Answer: all 11. QB 9: "all 11 (and
later more) are always available ... every report and chart label exists in every language"; roadmap line 406 says the
same. The 5-language lists of lines 301 and 308 are out of date; both rows restated.

**Q-39 Collector server "not ported" against work plan S-13?** (PP 7; REQ-PP-0193; S-13) Answer: the porting plan concerns the
router builds; work plan 3.5 keeps S-13, so the collector server is ported natively in step 1 for the PC and server build.

**Q-40 Zero forks on every tier against the sandboxed third-party programs?** (added in round 4 from a checker's finding;
REQ-PP-0029, -0184, REQ-RM-0006, -1350, REQ-WP-0179; P-13, X2-01, M-10, P-04) Porting plan line 686 sets 0 subprocess forks
per minute on Tier 1, 2 and 3; lines 127-128 give the reason (spawning `ip`, `ps`, `wpa_cli` on small CPUs). The owner's
decision 3 (work plan 7: Ookla's own program in the strict sandbox after the user accepts its terms) and roadmap 6B.5 (owner,
2026-11-10: Ookla's program, `ooniprobe`, iPerf3 and other foreign engines run in the stricter sandbox, REQ-RM-1350; Q-11)
need these programs started as processes. Answer: the owner's decisions are newer than line 686 and name the programs, so
the zero-fork target cannot cover their starts; it holds for `nmbd`'s own work, which is what lines 127-128 and REQ-RM-0006
are about. The rows keep their source wording and point here; the P-13 and X2-01 briefs measure the forks of `nmbd`'s own
work against the target and do not count the starts of these sandboxed programs against it.

Merged earlier and not asked again:

- The Web UI factory login admin/admin (REQ-RE-0231): the owner keeps it on his own Python box on purpose; the native build
  has a unique password per device (QB 7). Not a question.
- Ookla protocol emulation (REQ-PP-0080) and Ookla "eliminated" in roadmap section 7: superseded by work plan 7 point 3.
- The disclaimer wording (REQ-RM-1344, work plan 7 point 7): decided by the owner on 2026-11-11: the open-source "as is"
  warranty and liability exclusion, plus "measurement results are indicative, not a certified measurement; for an official
  complaint use the certified tool of your country", with the tool named only from the country data file (C-10); the
  neutral international default names none (QB 9).

## Developer's choices

**Q-02 F-07 before Core, or alongside?** (WP 1; REQ-WP-0056; F-00, F-07) Choice: F-07 belongs to the foundation, which work
plan 3.0 puts "in order, before anything else"; section 4 point 1 is read as F-00 to F-07. Its first version (the `spec/`
layout, the formats and the vector runner) is accepted before Core starts, then `spec/` grows with each package. Rule:
work plan 2.2 rule 11 ("behaviour is changed in `spec/` first") cannot hold for Core without it.

**Q-03 ARMv6 (Pi 1, Pi Zero) in the build matrix?** (RM1 7, RM2 6; REQ-RM-0013, -1440; F-00, P-13) Choice: `linux/arm` GOARM=6
joins the cross-compile-only set (built on every change, not run); Pi Zero and Pi 1 are not a tested or promised target of
step 1. Rule: *Generic First* point 6 (cross-compile what would break later "while it is cheap to fix"), roadmap section 8
names GOARM=6, and the work plan's step-1 targets stay as they are.

**Q-04 RAM and CPU projections as limits?** (RM2 5; REQ-RM-1413 to -1437, -1442, -1443, -1456 to -1460; P-13) Choice: measured
on the Pi 4 and recorded against each projected figure, not pass/fail on the figure. The pass condition is QB 4 ("the
improvement is measured, not claimed": lower than the Python version on the same box and day) and *Generic First* point 7.

**Q-08 OpenWrt package format?** (RM2 4; REQ-RM-1271; P-11) Choice: the format follows the target release: `apk` for OpenWrt
25.12 (the step-1 images), `.ipk` with `opkg` for 22.03 and older, which applies to the TL-WR1043ND v1 if the base chosen for
it is one of those (*Order of Work* step 2 leaves that base open: "the plan must also say which OpenWrt base the box image
uses on it"). Rule: the target's own installer and upgrade method (6B.3, *Generic First* point 3). REQ-RM-1271 restated.
Round 4: no longer presumes the router's base.

**Q-09 Documentation languages of the native repository?** (WP 5; REQ-DV-0026, REQ-CT-0016, -0022; P-12) The user
manual is in all 11 UI languages from one source: that is the owner's own text (roadmap 6B.7: "available in all 11 UI
languages, from one source so no language falls behind"). Choices: the command reference is in all 11 languages too (6B.7
gives the languages for the manual only; the choice follows QB 9: no language is a default, and every text a user reads
exists in every language); developer documents (briefs, test record, contribution rules, development setup) are written in
English as their one source (they are for developers and agents, and one source keeps them in step). The English-and-Hungarian
pair of the Python repository is not carried over. Round 4: moved here from the answered section, since only the manual's
languages are an existing decision.

**Q-10 Test scripts in the native repository?** (TS 6; REQ-TS-0312; F-06) Choice: yes, with reserved values only; private
values (the box's password, the owner's own data) stay outside. Rule: QB 6 (a software-caused known limit is not carried
over), QB 8, work plan 2.3A (what holds nothing private belongs in the native repository). The row is no longer pending.

**Q-13 Network-admin rights of the web server?** (WP 2; REQ-AR-0012; S-01, P-04) Choice: the native web server gets only the
right to bind its port (`CAP_NET_BIND_SERVICE`, or none with a socket passed by the service manager); network changes and
other privileged tasks go to the network daemon or the privileged helper. Rule: QB 7 "only the capabilities they need" and
6B.5. REQ-AR-0012 restated.

**Q-14 `/kill` from Telegram limited to NMB's processes?** (WP 3; REQ-AR-0068, REQ-RE-0280; S-11, P-06) Choice: yes, the
scope stays Python's effective one. In Python the bot runs as user `nmb` and its `sudo -n kill` fallback is not in the sudoers
line (`setup.sh@ad3758c:281`), so it signals only what `nmb` may signal ("a process the bot may signal", README.md:422; "the
task manager cannot act as root", README.md:452). Native: the bot signals NMB's own processes only; where one runs under
another service user than the bot, the privileged helper sends the signal after checking the process against its allow-list
of NMB's own units (never `sudo` or a shell). Every signal needs the confirmation tap of the owner's decision of 2026-11-11;
KILL, after which the process cannot clean up, counts as destructive and also needs the typed word. PID 1 and the bot stay
protected. Rule: QB 6 (the reference is the intended behaviour, here README:422 and :452), QB 7 ("only the capabilities they
need") and the helper allow-list of the fault register (roadmap line 610). Signalling root or foreign processes (sshd,
Tailscale) from the bot would be a wider scope and a security trade-off for the owner; it is not proposed. Round 4: moved here
from the answered section, where the round-3 text had widened the scope.

**Q-20 Donation addresses in the code?** (RE 4; REQ-RE-0310, -0311; P-08) Choice: not in the code base. The addresses come from a
release data file outside it, read by the release step, and are shown in the Web UI, the bot and the CLI as before; a build
without the file shows no donation section. Rule: QB 8 (nothing that belongs to a person in source, defaults, tests or
documentation examples; the build scan fails on the owner's values).

**Q-21 Vendor-look theme images and `hago.jpg`?** (PY 1, 2; the 20 theme-image rows of `parts/python-files.md`; P-09) Choice: no
theme image goes into the native repository. A native theme carries colours, fonts from the image and spacing only (6B.1);
the vendor-look themes are redrawn as colour schemes under neutral names, without makers' logos or names. Whether
`hago.jpg` is personal no longer matters: it stays out either way (QB 8). Any image theme the owner wants for himself is an
add-on file he adds to his own box (6B.1: add-ons are added by hand, never shipped).

**Q-22 Origin and licence of `components/open-rmbt/`?** (PY 4; its rows in `parts/python-files.md`; M-08, P-08) Choice: the native
build copies nothing from it; M-08 writes the Go client from the RMBT protocol documentation and upstream open-rmbt sources
(with their licences in the 6B.3A list) and uses the C client only as a behaviour reference. Rule: work plan 7 point 9 (no
component of unclear licence) and 6B.3A. Found in the history: the folder came in one commit
(`30f7cf8`, 2026-11-04) with no copyright header and no licence file; its origin matters only for the Python repository.

**Q-23 NAT type classifier without RFC 5780?** (RM1 5; REQ-RM-0123; M-12) Choice: RFC 5780 tests where a server supports them;
otherwise the type the probes cannot tell apart is reported as "unknown", never guessed. Rule: the fault register's
*Measurement honesty* (no made-up data). REQ-RM-0123 restated.

**Q-24 Flow Cache bypass by shell, not restored?** (PP 1; REQ-PP-0051; X2-01) Choice: the interface is called directly (ioctl
or `/proc`), never through a shell, and the previous state is always restored, also after a crash or restart. Rule:
REQ-PP-0029 and QB 7 (no shell), and 4.29 (the box's own changes are marked as its own action). REQ-PP-0051 restated.

**Q-25 CAKE shaping on the measured line?** (PP 2; REQ-PP-0053, -0186; X2-01, C-11) Choice: when a shaper is on the
measured line, the shaper and its rate are stored in the result's measurement setup and the result is marked as measured
through it. Rule: roadmap 4.29 (every result states its setup) and *Measurement honesty*. Whether the box lifts the shaper
during its own line speed tests is not decided by these rules: lifting it changes the user's network while the test runs
(REQ-PP-0186: "without slowing gaming"), so it becomes a setting or a question to the owner when X2-01 starts (step 2; no
step-1 row depends on it). Round 4: the lifting half taken out of the choice.

**Q-26 DynDNS from the local WAN address behind NAT?** (PP 5; REQ-PP-0088; N-16) Choice: the netlink address is used only when it
is public; behind NAT or CGNAT the box's public-address check (M-12) gives the address. Rule: the fault register ("the public
IP missing or replaced by the LAN address"). REQ-PP-0088 restated.

**Q-27 Fast.com over plain HTTP or TLS?** (CL 4; REQ-CL-0075; M-07) Choice: the M-07 analyst measures both on the same box and line;
TLS is used where it reaches the line rate on the Pi 4, plain HTTP where the comparison shows TLS limits the result; the
figures go into `docs/TESTING.md`. Rule: QB 4 names "TLS-capable engines at line rate on a Pi 4" as an improvement to prove.
REQ-CL-0075 restated.

**Q-30 RTR far above the other engines over 5G?** (TS 5; REQ-TS-0157; M-08) Choice: treated as unexplained, not as engine
behaviour: the M-08 analyst finds the cause (same link and time, server, stream count, bytes counted on the interface) before
the native engine is accepted, and until then that Python figure is no reference value. If the cause turns out to be a
choice (for example multi-stream by design), the analyst asks then. Rule: QB 6 (the reference is the intended behaviour).

**Q-31 Bloom filter false positives on small routers?** (PP 4; REQ-PP-0078; N-12, X2-01) Choice: a filter hit is confirmed by an
exact check (a compact exact list, on flash where RAM is short) before a domain is blocked. Rule: QB 6 (blocking a domain on
no list is a wrong result). REQ-PP-0078 restated.

**Q-32 Mode B when the box fails: fail-open bypass?** (RE 1; REQ-RE-0028, -0315; N-03) Choice: the two READMEs describe different
faults and both hold: when `nmbd` stops or fails, the kernel bridge keeps passing traffic (required, tested by stopping the
daemon); when the box itself stops (power, kernel), the connection stops. Bypass hardware is on no step-1 target, so it is
not supported and the manual says so. Rule: *Generic First* point 3 (nothing assumed that the hardware lacks) and honest
documentation. REQ-RE-0028 restated.

**Q-34 IPv6 for the LAN on a DHCP or USB WAN, and dual-band with a second radio?** (TS 4; REQ-TS-0244, -0256; N-11, N-07) Choice:
both in step 1. IPv6 on these WANs is tested between virtual ports in the F-06 QEMU guests; dual-band is used when a second
radio is detected and is tested under `mac80211_hwsim`; a real-hardware row needs a second USB Wi-Fi adapter and says "not
tested" until one is used. Rule: *Order of Work* step 1 ("every component, not a subset") and QB 4 (the open items the
Python version never finished). REQ-TS-0244 restated.

**Q-38 Compatibility names `pihole` and `pivpn`?** (RM1 6, PY 5; REQ-RM-0212; P-10, N-12, N-13) Choice: installed only where no
command of that name exists, never over one, and removed only if NMB made them; on NMB's own image they are there. Rule:
*Generic First* point 3 (detect on the running system) and QB 7 (like the teardown that flushed other software's rules,
NMB never breaks other software).


## Answered by the owner

The six questions left for the owner in round 4 were answered by him on 2026-11-11. The answers are given as the owner gave
them; the rows each answer settled were rewritten or added (`PACKAGE_CHANGES.md` section L).

**Q-01 How does the owner's Python box move over to the native build?** (RM2 1; REQ-RM-1002, REQ-RM-1260; P-13, M-01, F-05)
Answer (owner, 2026-11-11): component by component, as roadmap line 512 and Quality Bar 5 say: each native component
replaces its Python counterpart when its evidence is written down, through a bridge that native and Python parts share
during the migration (the Python database over SQLite WAL as in Phase 1.1, or an API). REQ-RM-1260 (native probe writing
into the Python database) is not withdrawn. The bridge contract is F-04's; each mixed state must be tested; the bridge must
not carry Python faults into native code (Quality Bar 6). F-01 changes nothing in roadmap line 512. (Option B; the
recommendation had been A.)
Rows: REQ-RM-1002 (Notes), REQ-RM-1260 (kept as the probe's bridge, owner M-01, F-04 and C-02 under *Also in*); new
REQ-RM-1466 (the bridge and its contract, F-04), REQ-RM-1467 (each mixed state tested, P-13), REQ-RM-1468 (no Python fault
through the bridge, C-02).

**Q-05 Which distributions must the native packages support?** (CL 1; REQ-CL-0090, -0091; P-11)
Answer (owner, 2026-11-11): A: the own image plus packages for every distribution the Python version supports (Debian,
Ubuntu, Alpine, the Fedora/RHEL family, Rocky Linux), each tested in QEMU.
Rows: REQ-CL-0090 restated; REQ-CL-0091 and the context of CHANGELOG line 131 point to the answer.

**Q-12 Should the native Web UI keep the browser terminal?** (PY 3, CL 3; REQ-CL-0416; S-08, U-20)
Answer (owner, 2026-11-11): B: kept, hardened: off by default, only from the LAN or Tailscale, only while the maintenance
mode is on, the password asked again each time, every command in the audit log (P-05), switching itself off after a set
time; it runs as the service user, root only through the privileged helper.
Rows: REQ-CL-0416 restated; new REQ-CL-0454 (`S-02…S-0n (API: system)`, the S-08 range in `parts/python-files.md`; the
tools page, P-05, P-06 and P-04 under *Also in*). `parts/python-files.md` question 3 said the Python terminal runs commands
as root; a correction there says it runs as the service user `nmb`.

**Q-15 May the Telegram bot carry secrets: a VPN profile's private key, a Wi-Fi password?** (RM1 4, WP 4; REQ-RM-0104,
REQ-AR-0054, -0069; S-11, N-13, N-14, N-08)
Answer (owner, 2026-11-11): C: OFF by default; the user switches "secrets over Telegram" on per item; then: the
confirmation tap, a warning, the bot deletes its QR message after a few minutes and the user's password message at once,
and the VPN peer can be withdrawn with one tap. (The recommendation had been A.)
Rows: REQ-AR-0054, -0069 restated, REQ-RM-0104 (Notes); new REQ-AR-0079 (VPN profiles) and REQ-AR-0080 (Wi-Fi passwords),
both S-11. Every other row on the bot sending a QR code, a `.conf` or taking a Wi-Fi password qualified ("only when secrets
over Telegram is switched on for this item") with a Notes pointer: REQ-RM-0104, REQ-AR-0055, -0056, REQ-RE-0121, -0147,
-0259, -0260, -0261, -0270, REQ-TS-0134, -0231, -0282.

**Q-16 MAC addresses in shared diagnostic bundles?** (RM1 3; REQ-RM-0135; C-04, R-04)
Answer (owner, 2026-11-11): C: kept in the SLA dossier and the ISP-support bundle; replaced by the vendor prefix and a
pseudonym stable within one bundle in bundles meant for outside LLMs.
Rows: REQ-RM-0135 restated (private IP subnets kept in every bundle; MAC addresses as above; R-03 added under *Also in*);
new REQ-RM-0220 (C-04).

**Q-17 An audit log "the session itself cannot erase"?** (PP 6; REQ-PP-0135; P-05)
Answer (owner, 2026-11-11): B: the maintenance mode does not start without a log target off the device (collector, syslog
or the hub of the remote paths); each entry is sent as it is written; the hash-chained local log stays as a second record.
Rows: REQ-PP-0135 restated (S-13 added under *Also in*); new REQ-PP-0199 (P-05).

## For the owner

None open. The six questions of round 4 (Q-01, Q-05, Q-12, Q-15, Q-16, Q-17) were answered on 2026-11-11 (section above).
