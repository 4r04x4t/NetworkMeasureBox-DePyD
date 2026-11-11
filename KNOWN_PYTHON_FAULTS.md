# Known faults register (Python reference, Go and C code bases)

This is the shared fault register of the native build (roadmap `docs/DEPYTHONIZATION_ROADMAP.md`, *Register of
Python Faults*, line 606, and section 6B.8 point 3, line 923, of the Python repository at commit `ad3758c`). Every
fault found in the Python version, fixed or not, is a requirement for the native build: the native component is
designed so that the fault cannot happen, and it has a test that fails if it does (Quality Bar points 2, 6 and 7).
A native component is not accepted while one of its entries has no test. Every fault found later in the Go or the C
code base is added here too, with a status per code base and a conformance vector in `spec/vectors/`.

## Sources

Read in full for this register (package F-03; the reading proof is `docs/briefs/F-03.md`):

- the Python repository at `ad3758c`: `docs/DEPYTHONIZATION_ROADMAP.md` (Quality Bar, Register, Phase 0, 4.29,
  6B.8), `docs/TESTING.md` (results, release checklist, *To fix*, *Known limits*), `CHANGELOG.md` (every *Fixed* and
  *Security* entry), `docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md` section 7D, the git history, and the Python
  code where a fault's state was unclear (source analysis);
- the private log reviews of 2026-11-09 and 2026-11-10, the private pre-rewrite audit of 2026-11-10 and the private
  data-quality list of 2026-11-10. They stay outside every repository; they are cited only as "private audit
  2026-11-10, <section> item <n>" (or "private log review", "private data-quality list"), never quoted in detail;
- the private package list of 2026-11-11, written by this package: for each open security category, the packages
  that must also respect it, next to the private audit items. It stays outside every repository too.

A source is written as `<file>@ad3758c:<line>` (Python repository) or as one of the private names above.

## Rules of this register

- **Open security holes by category only.** A security hole that is still open in the Python version is recorded by
  its category, never with the place, the input or how to use it (roadmap 608). A category holds several private
  items, and its correct behaviour is written only from the generic rules of Quality Bar 7 (roadmap 587-591) and the
  rules from fixes, never clause by clause against one item. The specific requirements stay in the private audit:
  the analyst of the proving package reads the items named under *Source* there, until the hole is fixed or
  published. A category entry names only its proving package; the packages that must also respect it are kept with
  the item numbers in the private package list of 2026-11-11, and their analysts read the same items. When the
  hole is fixed in the Python version or the details are no longer private, the entry may be completed. Fixed holes say what was wrong, as the public CHANGELOG does. Open faults
  that are not security holes never name a place or a mechanism of an open hole either, and an open fault that an
  outside party could trigger is described by its effect only: its trigger and its code lines stay in the private
  notes.
- **No personal or site data** (Quality Bar point 8): the owner's ISP, devices, addresses, places, account names and
  tokens are written generically.
- **Ids are permanent.** `KF-0001` to `KF-0047` are the roadmap register (lines 610-619), in its order; new entries
  start at `KF-0048`. An id is never reused or renumbered; a fault found to be no fault is marked so, not deleted.
- **One package proves each fault**; others that must also respect it are listed as "also" (packages: work plan
  `docs/NATIVE_WORK_PLAN.md` section 3; `S-02` and `U-02` stand for the API-area and Web UI page packages
  `S-02…S-0n` and `U-02…U-0n` that the integrator splits).
- **Correct behaviour** is taken from the sources; where a source does not state it, it is derived from the
  roadmap's rules and marked "derived" with the rule and its line. Cases where it is unclear whether something is a
  fault or intended are listed under *Questions for the owner* before the entries, not guessed; an answer is
  written into the entries it decides as "Owner, <date>: ..." in *Correct behaviour*, and the question moves to
  *Answered by the owner*.
- **Python fix commit**: one or more commit ids, or `not recorded in the source` (also for a fault that is not
  fixed). **Status, Python**: `fixed`, `being fixed` or `open`, as the newest source says. The Python version is
  closed (owner, 2026-11-10): its open faults are not fixed there, they are requirements here.
- **Field values come first, notes after a fixed separator.** In *Python fix commit* and the three *Status* fields the
  value from the list above stands first; anything else follows after `; note: ` (for example
  `open; note: partly fixed`), so a machine check reads the value up to the first `;`.
- **Status, Go**: `open` until the entry's native test passes in the gate; then `fixed (test passes)`.
  **Status, C**: `not started` until step 2.
- **Groups by id range** (the order of the entries; no group headings stand between them): KF-0001 to KF-0047 the
  roadmap register (lines 610-619); KF-0048 to KF-0099 the Quality Bar examples, Phase 0, the rules from fixes and
  the *Changed* section; KF-0100 to KF-0218 the *Fixed* and *Security* entries of the CHANGELOG and the testing
  record; KF-0219 to KF-0266 faults still open in the Python version that are not security holes; KF-0267 to
  KF-0282 security holes still open in the Python version, by category only; KF-0283 to KF-0290 added after the
  first review of this register (2026-11-11); KF-0291 and KF-0292 added after the second review (2026-11-11);
  KF-0293 and up added after the third review and the owner's answers (2026-11-11).
- **Layout.** *Entry format* and *Questions for the owner* (with *Answered by the owner*) stand before *Entries*;
  `## Entries` is the last section of the file and holds nothing but entries, so an entry runs from its `### KF-`
  line to the next one or to the end of the file.
- Machine-checkable later by F-07: every entry has the twelve fields of *Entry format*, in that order.

## Entry format

Every entry has this form (F-00 brief section 6), the twelve fields in this order and nothing else:

```
### KF-<nnnn>: <short title>

- Area: <area>
- What went wrong: <the fault>
- Correct behaviour: <the behaviour the native build must have, with its source>
- Source: <source>; <source>; ...
- Python fix commit: <commit id>[, <commit id> ...] | not recorded in the source
- Package that proves it: <package id> [(also <package id>, ...)]
- Spec vector: <vector id> | none yet (F-07)
- Native test, Go: <test name> | none yet
- Native test, C: <test name> | none yet (step 2)
- Status, Go: open | fixed (test passes) | not applicable (<reason>)
- Status, C: not started | open | fixed (test passes) | not applicable (<reason>)
- Status, Python: fixed | being fixed | open
```

- Allowed values: *Status, Go* is `open`, `fixed (test passes)` or `not applicable (<reason>)`; *Status, C* is
  `not started`, `open`, `fixed (test passes)` or `not applicable (<reason>)`; *Status, Python* is `fixed`,
  `being fixed` or `open`; *Python fix commit* is one or more commit ids or `not recorded in the source`. In these
  four fields a note may follow the value after `; note: `.
- Citation form: a line of the Python repository is written `<file>@ad3758c:<line>` or
  `<file>@ad3758c:<first>-<last>`, with the path from the repository root (for example
  `docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:610`); inside *Correct behaviour* a short form names the document and the
  line ("roadmap 587", "CHANGELOG 191"). The private notes are cited by name, date and item number only (see
  *Sources*).
- An open security category names only its proving package (see *Rules*).

## Questions for the owner

These are not entries yet: it is unclear whether the Python behaviour is a fault or intended. The answers go into
this register (as an entry, or as a note that it is intended) and into the roadmap.

None is open: questions 1 to 7 are answered (below). A new question gets the next number, 8.

### Answered by the owner

Answered on 2026-11-11. Each answer is written into the entries it decides, as "Owner, 2026-11-11: ..." in
*Correct behaviour*.

- **Question 1, manual and bot Ookla runs and the hourly limit:** a test started by a person (Web UI, CLI, bot)
  is not held back by the hourly schedule limit, but during a refusal pause it is refused with the time left.
  Written into KF-0094 and KF-0292.
- **Question 2, a VPN server running while its setting is off:** show and ask. The status shows the mismatch, the
  firewall keeps the running server's port open (a running VPN is never cut off), and the user chooses to stop it
  or to turn the setting on. Written into KF-0021.
- **Question 3, a default data cap for metered links:** no default cap number; on a metered link scheduled speed
  tests wait until the user sets a cap or says unlimited; ping keeps running. Written into KF-0294 (new), KF-0136
  and KF-0230.
- **Question 4, raw ping samples for an authority:** minute summaries only; no raw ping-sample export. Written into
  KF-0289.
- **Question 5, which actions need a confirmation:** every Telegram bot command that changes the box needs a
  confirmation tap; destructive ones (factory reset, purge) also need the typed word, as in the Web UI; read-only
  commands run at once. Written into KF-0075 and KF-0277.
- **Question 6, what a missing translation shows:** the build refuses every built-in text that is missing in any of
  the 11 languages. A key that is still missing at run time (for example from an add-on) shows the English text,
  visibly marked as untranslated; it is logged and listed in the self-test result and in the diagnostic bundle. It
  is never a silent fallback to any language. This agrees with roadmap 4.26 (line 406, English) and Quality Bar 9
  (line 602, no silent default). Written into KF-0048 and KF-0049, and into the other entries on missing
  translations: KF-0057, KF-0058, KF-0239 and KF-0287.
- **Question 7, a test started by a person on a metered link without a cap decision:** it waits like a scheduled
  test until the user sets a cap or says unlimited, and the user is told why it waits; ping keeps running. A cap
  that is reached (`CHANGELOG.md@ad3758c:35`: tests started by hand still run and count) is a different case and
  is not changed by this answer. Written into KF-0294, KF-0136 and KF-0230.
- **Trust of the uplink network** (not a numbered question; raised by the third review of KF-0015, KF-0084 and
  KF-0130): the role of a network the box is plugged into is never inferred from address classes. Nothing is open
  from the uplink side (Web UI, SSH, DNS) until the owner marks that network as their own; the first-start setup
  offers the network it was done from and the box remembers it; an unknown or new network gets nothing until it is
  confirmed; the LAN ports, the VPNs and Tailscale are not affected. Written into KF-0015, KF-0084 and KF-0130. It
  replaces the address-class wording of roadmap 681 ("local service ports must refuse public addresses on IPv4 and
  IPv6"); that roadmap line needs a correction, which is outside F-03.

## Entries

### KF-0001: Update test run as root overwrote the real settings

- Area: Privileged helper and updates
- What went wrong: the update's test run as root overwrote the real settings
- Correct behaviour: Tests of an update run unprivileged in a sandbox on a copy of the code, with the settings, database, locks, state files, lists, crash reports, bundles, backups and logs all pointed at throwaway files; a test that still reaches a real file fails and the update is refused; afterwards the real settings, secrets copy and boot copy are byte for byte unchanged. A related update path is covered under the category KF-0270.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:610; CHANGELOG.md@ad3758c:204; docs/TESTING.md@ad3758c:118; private log review 2026-11-10 (night), update section
- Python fix commit: bfbba44
- Package that proves it: P-02 (also F-00: the gate refuses to run tests as root; P-04)
- Spec vector: none yet (F-07)
- Native test, Go: gate self-test case R12 of the F-00 brief section 5 (planned: tests refuse to run as root); P-02: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0002: Root helper took paths, packages and files from the caller

- Area: Privileged helper and updates
- What went wrong: root helper actions took paths, packages and files from the caller without an allow-list
- Correct behaviour: The privileged helper does only what NMB needs: it repairs only NMB's own links and folders, restores settings only from the boot-partition copy and through the normal merge, installs engine programs only to fixed places and only files root prepared itself, installs only listed driver, firmware and repair packages, and refuses everything else; a test checks that each allow-list matches the code.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:610; CHANGELOG.md@ad3758c:344; private audit 2026-11-10, privileged-operations item 1
- Python fix commit: e6512aa
- Package that proves it: P-06
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0003: Root code could import a module from the caller's folder

- Area: Privileged helper and updates
- What went wrong: a stray `PYTHONPATH` or working directory could make root code import the wrong module
- Correct behaviour: Privileged code runs from `/` with a fixed code path; it never loads code from the working directory, the environment or any folder the service user can write (CHANGELOG 344; rules from fixes, roadmap 679). For a static native binary: no plug-in or script loading from such places at all (derived: roadmap 679).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:610; CHANGELOG.md@ad3758c:344; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:679; private audit 2026-11-10, privileged-operations item 2 and reports item 2
- Python fix commit: e6512aa
- Package that proves it: P-06 (also P-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0004: Plaintext settings fallback on the boot partition

- Area: Privileged helper and updates
- What went wrong: a plaintext fallback wrote the settings to the boot partition
- Correct behaviour: A failed settings save never falls back to writing the full, unmasked settings anywhere, least of all onto the boot partition; secrets stay off the boot partition (rules from fixes).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:610; CHANGELOG.md@ad3758c:345; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:682; private audit 2026-11-10, reports item 3
- Python fix commit: ffceb45
- Package that proves it: C-01 (also C-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0005: Diagnostic bundle, crash reports and boot copy carried secrets

- Area: Secrets and personal data
- What went wrong: the diagnostic bundle, crash reports and the boot-partition copy carried tokens, passwords, user names and multi-line values
- Correct behaviour: Everything that leaves the box or lands on the boot partition is masked by one shared list of secrets and personal details; a value is replaced by one fixed marker for its kind (roadmap 4.24, 382-385): passwords, Wi-Fi keys and auth tokens by `[REDACTED_SECRET]`, Telegram bot tokens and chat ids by `[REDACTED_CHAT_ID]`, ISP account and circuit ids by `[REDACTED_CIRCUIT_ID]`, Dynamic DNS domain names by `[REDACTED_DDNS]`. The marker never carries the length or any other property of the value: the Python form `<redacted, N chars>` is not carried over. Every line of a multi-line secret is masked; private subnets and MAC addresses stay for topology diagnosis (roadmap 386); reading the boot copy back still keeps all current values; crash reports and the diagnostic bundle use the same list (roadmap 4.24).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:611; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:377-386; CHANGELOG.md@ad3758c:205; CHANGELOG.md@ad3758c:345; private audit 2026-11-10, bot item H5, web item 7, reports items 4, 11, 12
- Python fix commit: bfbba44, 5beb63d, ffceb45
- Package that proves it: C-04 (also R-04, C-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0006: Backup sent to Telegram in plain text

- Area: Secrets and personal data
- What went wrong: `/backup send` put the plain archive with every secret into Telegram
- Correct behaviour: Encrypted by default: the backup is sent only encrypted (AES-256 with a backup passphrase of at least 12 characters that is never shown again); without a passphrase nothing is sent and the bot says how to set one; plain only when the user asks for it explicitly, with a warning.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:611; CHANGELOG.md@ad3758c:192; private audit 2026-11-10, bot item H1
- Python fix commit: 3890378
- Package that proves it: S-11 (also P-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0007: Bot authorised the chat, not the person

- Area: Telegram bot
- What went wrong: it checked the chat, not the person (in a group every member was an admin)
- Correct behaviour: The sender's own user id must be listed too, and a private chat must be the listed user's own chat; members of a listed group are not admins; a stranger is told "Access Denied" at most every 10 minutes, group members not at all.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:612; CHANGELOG.md@ad3758c:194; private audit 2026-11-10, bot items H2 and L2
- Python fix commit: 301e1bd
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0008: Button data could carry any command

- Area: Telegram bot
- What went wrong: button data could carry any command
- Correct behaviour: A button press runs only the commands that the bot's own buttons carry.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:612; CHANGELOG.md@ad3758c:194; private audit 2026-11-10, bot item H2
- Python fix commit: 301e1bd
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0009: Edited messages ran again

- Area: Telegram bot
- What went wrong: edited messages ran again
- Correct behaviour: An edited message is never run again.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:612; CHANGELOG.md@ad3758c:194; private audit 2026-11-10, bot item M2
- Python fix commit: 301e1bd
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0010: Commands queued while the box was off ran hours later

- Area: Telegram bot
- What went wrong: commands queued while the box was off ran hours later
- Correct behaviour: Commands older than 2 minutes when they reach the box are dropped.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:612; CHANGELOG.md@ad3758c:194; private audit 2026-11-10, bot item M2
- Python fix commit: 301e1bd
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0011: Reload kept the old bot token

- Area: Telegram bot
- What went wrong: a reload kept the old token
- Correct behaviour: A reload (SIGHUP) uses the new bot token and the new alert recipients.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:612; CHANGELOG.md@ad3758c:194; private audit 2026-11-10, bot item M4
- Python fix commit: 301e1bd
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0012: OpenVPN profile name could reach the server's keys

- Area: VPN
- What went wrong: an OpenVPN profile name could leave its directory and read the server's keys
- Correct behaviour: A client profile is handed out only for a recorded client name; names of the server's own files can never be used for new clients; no path built from a name leaves its directory (roadmap 587).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:613; CHANGELOG.md@ad3758c:346; private audit 2026-11-10, web item 1 and netd item H6
- Python fix commit: 4c1eea2
- Package that proves it: N-14
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0013: Deleted OpenVPN client could still connect

- Area: VPN
- What went wrong: a deleted client could still connect (no revocation)
- Correct behaviour: The server accepts only clients on an allow-list kept by root; deleting a client removes it from the list and ends its session.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:613; CHANGELOG.md@ad3758c:346; private audit 2026-11-10, web item 2
- Python fix commit: 4c1eea2
- Package that proves it: N-14
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0014: Config line without final newline glued two directives

- Area: VPN
- What went wrong: a config line without a final newline glued the next directive onto it and stopped OpenVPN
- Correct behaviour: Lines appended to an existing configuration never join its last line: a missing final line break is added first, and a generated file is checked before the service is restarted with it (derived: porting plan 7D 626, a configuration check before applying).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:613; CHANGELOG.md@ad3758c:346; private audit 2026-11-10, fix progress (OpenVPN incident)
- Python fix commit: f5b0177
- Package that proves it: N-14 (also C-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0015: No firewall for the box in standalone and transparent mode

- Area: Firewall
- What went wrong: standalone and transparent mode had no firewall for the box itself
- Correct behaviour: A host guard in every mode: from the uplink side the box is reached only by answers to its own traffic, ICMP, DHCP, Tailscale and the VPN servers that are on or running (KF-0021). Owner, 2026-11-11: the role of a network the box is plugged into is never inferred from address classes; nothing is open from the uplink side (Web UI, SSH, DNS) until the owner marks that network as their own; the first-start setup offers the network it was done from and the box remembers it; an unknown or new network gets nothing until it is confirmed; the LAN ports, the VPNs and Tailscale are not affected. Access follows the interface a packet arrives on and the role of its network: the LAN ports, Tailscale, WireGuard, OpenVPN and an uplink network the owner marked as their own (its IPv6 prefix followed when it changes, CHANGELOG 191) reach the box's services; on the uplink no source is trusted for its address alone, and reverse-path (anti-spoofing) filtering is on for IPv4 and IPv6 (derived: roadmap 587, nothing listens on the WAN side unless the owner turned it on). Automatic mode choosing standalone on a public ISP DHCP line is covered.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:614; CHANGELOG.md@ad3758c:191; private audit 2026-11-10, netd item H1
- Python fix commit: d3703b1
- Package that proves it: N-02 (also U-02: the first-start setup offers the network)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0016: PPPoE came up before the router's rules

- Area: Firewall
- What went wrong: PPPoE came up before the router's rules
- Correct behaviour: The WAN guard is up before PPPoE dials; the address of the new link is never reachable without rules.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:614; CHANGELOG.md@ad3758c:191; private audit 2026-11-10, netd item H3
- Python fix commit: d3703b1
- Package that proves it: N-02 (also N-05)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0017: Forwarding on before the rules, and kept on when they failed

- Area: Firewall
- What went wrong: forwarding was switched on before the rules and stayed on when they failed
- Correct behaviour: IP forwarding is switched on only after the rules took, stays off if they fail (fail closed), and the status shows it.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:614; CHANGELOG.md@ad3758c:191; private audit 2026-11-10, netd item H3
- Python fix commit: d3703b1
- Package that proves it: N-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0018: WAN2 and IPv6 unfiltered in the fallback

- Area: Firewall
- What went wrong: WAN2 and IPv6 in the fallback were unfiltered
- Correct behaviour: WAN2 gets the same NAT and protection as the main WAN; a fallback firewall covers IPv6 too, or IPv6 forwarding is turned off when IPv6 cannot be filtered.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:614; CHANGELOG.md@ad3758c:191; private audit 2026-11-10, netd items H2 and H3
- Python fix commit: d3703b1
- Package that proves it: N-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0019: Teardown flushed other software's NAT rules

- Area: Firewall
- What went wrong: teardown flushed other software's NAT rules
- Correct behaviour: Only NMB's own rules are removed; other software's tables and rules (containers, Tailscale, VPNs) are never touched.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:614; CHANGELOG.md@ad3758c:191; private audit 2026-11-10, netd item M9
- Python fix commit: d3703b1
- Package that proves it: N-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0020: Fallback firewall rules piled up

- Area: Firewall
- What went wrong: fallback rules piled up
- Correct behaviour: Applying the rules again gives the same rule set (idempotent); rules never pile up.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:614; CHANGELOG.md@ad3758c:191; private audit 2026-11-10, netd item M9
- Python fix commit: d3703b1
- Package that proves it: N-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0021: Firewall following only the settings would cut off a running VPN server

- Area: Firewall
- What went wrong: a firewall that followed only the settings would have cut off a VPN server that was running
- Correct behaviour: The guard keeps the port of every VPN server that runs, even when the settings say it is off. Owner, 2026-11-11: show and ask: the status shows the mismatch (running while its setting is off), the firewall keeps the running server's port open (a running VPN is never cut off), and the user chooses to stop it or to turn the setting on.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:614; CHANGELOG.md@ad3758c:191; private audit 2026-11-10, fix progress (firewall); nmb/netd/openvpn.py@ad3758c:372-385 (source analysis: the status reports only whether the server runs)
- Python fix commit: f43db0c
- Package that proves it: N-02 (also N-13, N-14, U-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the port part; the status does not show the mismatch or offer the choice (owner, 2026-11-11)

### KF-0022: DNS blocker cache poisoning

- Area: DNS ad-blocker
- What went wrong: any datagram reaching the upstream port was taken as the answer and cached for everyone (cache poisoning)
- Correct behaviour: Each upstream query has a random ID and its own connected socket; only a reply to that exact question is used, and the client's ID is put back.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:615; CHANGELOG.md@ad3758c:193; private audit 2026-11-10, netd item H4
- Python fix commit: e27a151
- Package that proves it: N-12
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0023: Malformed DNS names read as other names

- Area: DNS ad-blocker
- What went wrong: malformed names were read as other names
- Correct behaviour: Names with bad bytes, pointers in the question, over-long names or several questions are refused; the cache keys on class, flags and the DNSSEC OK bit and keeps only complete NOERROR/NXDOMAIN answers. The parser has a fuzz test (roadmap 588).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:615; CHANGELOG.md@ad3758c:193; private audit 2026-11-10, netd item M2
- Python fix commit: e27a151
- Package that proves it: N-12
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0024: One socket error stopped DNS for the whole network

- Area: DNS ad-blocker
- What went wrong: one socket error stopped DNS for the whole network
- Correct behaviour: A socket error (an oversized datagram, a moment without file descriptors) never stops the service; queries are answered by a bounded pool, so one slow upstream answer does not hold up the others.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:615; CHANGELOG.md@ad3758c:193; private audit 2026-11-10, netd items H5 and M3
- Python fix commit: e27a151
- Package that proves it: N-12
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0025: Unsent pings not counted as lost

- Area: Measurement honesty
- What went wrong: pings that could not even be sent (no route, link down) were not counted as lost, so such outages were never recorded
- Correct behaviour: A ping that cannot be sent counts as lost, exactly like a timeout, and counts towards an outage.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:616; CHANGELOG.md@ad3758c:196; private audit 2026-11-10, engines item 1
- Python fix commit: 6673b88
- Package that proves it: M-01 (also M-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0026: Local router's answer ended an internet outage

- Area: Measurement honesty
- What went wrong: an answer from the local router ended an internet outage
- Correct behaviour: The internet targets decide an outage; the local gateway answering never hides or ends one (such outages are recorded as "internet targets unreachable while the local gateway answers"); in reports only an answer from the internet ends an outage.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:616; CHANGELOG.md@ad3758c:196; private audit 2026-11-10, engines item 2
- Python fix commit: 6673b88
- Package that proves it: M-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0027: Web UI showed made-up data

- Area: Measurement honesty
- What went wrong: the Web UI showed made-up data (fake clients, another device's firmware number, invented rates)
- Correct behaviour: The box shows its own installed version, no client when none was seen, and only counts the sniffer really saw; nothing is invented to fill a screen.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:616; CHANGELOG.md@ad3758c:197; private audit 2026-11-10, web item 11
- Python fix commit: 930be52
- Package that proves it: S-02 (also U-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0028: Ookla's terms accepted silently for the user

- Area: Measurement honesty
- What went wrong: Ookla's terms were accepted silently for the user
- Correct behaviour: Ookla runs only after the owner has accepted Ookla's licence and privacy terms (links to them, an accept button, a way to withdraw); until then Ookla is skipped, nothing is stored for it and nothing stands in for it; existing installations start as not accepted.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:616; CHANGELOG.md@ad3758c:118; docs/NATIVE_WORK_PLAN.md@ad3758c:313
- Python fix commit: 7e10bf4
- Package that proves it: M-10
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0029: Automatic-engine badge did not follow the setting

- Area: Measurement honesty
- What went wrong: the automatic-engine badge did not follow the setting
- Correct behaviour: The home page follows the saved engine and, for the mixed run, lists the engines that take part.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:616; CHANGELOG.md@ad3758c:195
- Python fix commit: 0485124
- Package that proves it: U-02 (also S-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0030: Packet loss 0 % from engines that do not measure it

- Area: Measurement data
- What went wrong: packet loss written as 0 % by engines that do not measure it
- Correct behaviour: A value an engine does not measure is stored empty (unknown), never 0; this applies to loss from Cloudflare and RTR and to Ookla results without a loss figure. Each engine writes its own loss value, so the engines prove it: M-08 (RTR) proves it, M-06 (Cloudflare) and M-10 (Ookla) do the same for theirs, and M-04 (the runner that stores the row) keeps an empty value empty.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:182; private data-quality list 2026-11-10, item D1
- Python fix commit: c2d0ed6
- Package that proves it: M-08 (also M-06, M-10, M-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0031: Box's own actions counted as ISP outages

- Area: Measurement data
- What went wrong: outages caused by the box's own mode changes, updates and settings counted against the ISP
- Correct behaviour: Every window in which the box takes its own connection down (mode change, rebuild after a settings change, PPPoE switched in auto mode, WAN change, its own stop including updates and reboots) is recorded; reports take these windows plus 2 minutes out of every outage, list such outages apart as the box's own, leave the windows out of the measured time and do not count pings lost in them; no outage alarm is sent while the box itself has the line down. A rebuild the box makes in reaction to the line (failover, PPPoE appearing) is not taken out, and an outage counts as the box's own only when it began while the box had its connection down.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:180; CHANGELOG.md@ad3758c:163; private data-quality list 2026-11-10, item D4
- Python fix commit: f101c63, da88eb1
- Package that proves it: C-07 (also M-02, R-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0032: ISP name and line type from the settings

- Area: Measurement data
- What went wrong: ISP name and line type taken from the settings instead of the measurement
- Correct behaviour: The ISP name per test comes from the settings or from the test's own public address, never from a lookup of today's address (CHANGELOG 158); the line type follows the link the test really used (a test over USB tethering is mobile, CHANGELOG 59). Derived (roadmap 4.29, 450: "stated, not inferred by the reader"): a value taken from the settings is marked as the owner's declaration, a measured one as measured.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:158; CHANGELOG.md@ad3758c:59; private data-quality list 2026-11-10, item D5
- Python fix commit: c2d0ed6, da88eb1
- Package that proves it: M-04 (also R-01, M-12)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0033: Invented contract and guaranteed speeds used for verdicts

- Area: Measurement data
- What went wrong: invented contract and guaranteed speeds used for verdicts when none were set
- Correct behaviour: No status is judged against a speed that was never set: the default contracted speed is 0 = not set, forms do not pre-fill speeds that a save would store, the bot says "not judged" and shows only the shares of speeds that are set, and e-mails show "not set".
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:115; CHANGELOG.md@ad3758c:188; CHANGELOG.md@ad3758c:189; CHANGELOG.md@ad3758c:190; CHANGELOG.md@ad3758c:148; private data-quality list 2026-11-10, item D6
- Python fix commit: c2d0ed6, 428b608
- Package that proves it: R-01 (also S-11, U-02, R-05)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0034: Missing values exported as 0 ms / 0 %

- Area: Measurement data
- What went wrong: missing values exported as `0 ms` / `0 %`
- Correct behaviour: A missing value is empty or "--" in every export and surface; a measured 0 is exported as 0.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:188; CHANGELOG.md@ad3758c:156; private data-quality list 2026-11-10, item D6
- Python fix commit: c2d0ed6, da88eb1
- Package that proves it: R-02 (also R-03, U-02, S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0035: Times without a time zone

- Area: Measurement data
- What went wrong: times without a time zone
- Correct behaviour: Every exported time states its zone (the detailed CSV has a time zone column, UTC; the speed CSV a time zone column).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:188; CHANGELOG.md@ad3758c:23; private data-quality list 2026-11-10, item D6
- Python fix commit: c2d0ed6, 1012b59
- Package that proves it: R-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0036: Public IP missing or replaced by the LAN address

- Area: Measurement data
- What went wrong: the public IP missing or replaced by the LAN address
- Correct behaviour: Only a public address is stored as the test's address; against an internet server the row has the line's public address asked over the same source address after the test, against a LAN server the box's local address (and the test is marked LAN); a failed lookup leaves it empty, never another address in its place; each row says where the address came from (`client_ip_source`: engine, lookup, socket, local).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:168; CHANGELOG.md@ad3758c:187; private data-quality list 2026-11-10, missing items
- Python fix commit: c2d0ed6, b8fb934
- Package that proves it: M-12 (also M-08, M-09, M-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0037: All engines of a mixed run shared one timestamp

- Area: Measurement data
- What went wrong: all engines of one mixed run sharing one timestamp
- Correct behaviour: Each row records when its engine itself started (`ts`) and the start of the run it belonged to (`cycle_ts`), plus how long the engine ran.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:22; private data-quality list 2026-11-10, item D10
- Python fix commit: 1012b59
- Package that proves it: M-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0038: Generic web address given as the result link

- Area: Measurement data
- What went wrong: a generic web address given as the result link
- Correct behaviour: A result link is stored only when it points to this result at the third party; otherwise there is none.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:185; CHANGELOG.md@ad3758c:301; CHANGELOG.md@ad3758c:183; private data-quality list 2026-11-10, item D9
- Python fix commit: c2d0ed6
- Package that proves it: M-06 (also M-08)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0039: Diagnostic bundle concluded double NAT and LAN fault without proof

- Area: Measurement data
- What went wrong: the diagnostic bundle concluding "double NAT" from a PPP peer address and "LAN fault" from outage bursts
- Correct behaviour: The bundle concludes only what its data shows: "possibly double NAT" with the reason a traceroute cannot tell; a known PPP peer placeholder is called that; CGNAT is "consistent with", not "detected"; a LAN-fault verdict uses only the loss left once the box's own actions and recorded outages are taken out, and names possible causes instead of a certainty.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:172; CHANGELOG.md@ad3758c:173; CHANGELOG.md@ad3758c:174; private data-quality list 2026-11-10, item D7
- Python fix commit: 96832e9
- Package that proves it: R-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0040: Diagnostic bundle did not know its own version

- Area: Measurement data
- What went wrong: the diagnostic bundle not knowing its own version
- Correct behaviour: The version is the installed one, read from the installed version record, everywhere.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:617; CHANGELOG.md@ad3758c:178; CHANGELOG.md@ad3758c:23; private data-quality list 2026-11-10, item D7
- Python fix commit: 96832e9
- Package that proves it: R-04 (also C-05)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the register line 617 says "being fixed"; it was written before the fixes landed: CHANGELOG and the private data-quality list record them as live on 2026-11-10

### KF-0041: Result did not say whether the box was behind another router

- Area: Measurement setup
- What went wrong: a result did not say whether the box measured behind another router or directly on the ISP
- Correct behaviour: Every speed test, outage and ping minute states three facts: the box's mode (for `auto` the mode it resolved to), the link the measurement left through, and the path to the ISP (`direct_pppoe`, `direct_public`, `direct_cgnat`, `behind_router`), plus where the facts come from (`recorded`, `changed`, `derived`), with the same codes as the Python version; every surface shows it (roadmap 4.29, 450 and 455).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:618; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:448; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:450; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:455
- Python fix commit: ad3758c
- Package that proves it: C-11 (also M-01, M-02, M-04, R-01..R-05, S-02, U-02, S-10, S-11, S-12, S-13)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0042: Mode stored late and read after the test

- Area: Measurement setup
- What went wrong: the mode was stored only from 2026-11-10 noon on, read after the test, so a switch during it went unseen
- Correct behaviour: The setup is captured when the test starts and checked at its end; a mode switch during the test marks it `changed` (roadmap 4.29 line 450).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:618; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:450
- Python fix commit: ad3758c
- Package that proves it: C-11 (also M-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0043: Link and position never stored

- Area: Measurement setup
- What went wrong: the link and the position were never stored
- Correct behaviour: The link and the position are stored with every result; the position comes from the link's IPv4 address, never from the mode; a global IPv6 address proves nothing; VPN interfaces are never the measuring link and Tailscale's 100.64/10 is not the ISP's CGNAT; without an address to bind to the engines leave through the default route and that link is described; what cannot be told for sure stays unknown (roadmap 4.29 lines 450-453).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:618; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:450; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:451; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:452; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:453
- Python fix commit: ad3758c
- Package that proves it: C-11 (also N-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0044: Skipped rows had no context

- Area: Measurement setup
- What went wrong: skipped rows had no context at all
- Correct behaviour: Skipped and aborted rows carry the same facts and setup as measured ones, and no values (CHANGELOG 164; rules from fixes, roadmap 683; derived: roadmap 4.29 448 "every speed test").
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:618; CHANGELOG.md@ad3758c:164; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:683
- Python fix commit: ad3758c, da88eb1
- Package that proves it: M-04 (also C-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0045: Diagnostic bundle reported the configured mode

- Area: Measurement setup
- What went wrong: the diagnostic bundle reported the configured mode (`auto` read as "LAN probe behind a router" also when the box was the router)
- Correct behaviour: The bundle states the active mode with the configured one beside it.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:618; CHANGELOG.md@ad3758c:139
- Python fix commit: ad3758c
- Package that proves it: R-04 (also C-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0046: Outages and ping minutes had no setup

- Area: Measurement setup
- What went wrong: outages and ping minutes had no setup
- Correct behaviour: netd records a history row whenever the mode, the link, its kind or its position changes (and at stop); a ping minute or an outage is read against the setup of its time, and an outage takes the last record before it (roadmap 4.29 line 454); the native build reads the Python history and does not repeat the one-time reconstruction (line 456).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:618; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:454; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:456
- Python fix commit: ad3758c
- Package that proves it: M-02 (also C-11, M-01, C-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0047: Tests depended on the machine

- Area: Tests that depended on the machine
- What went wrong: tests depended on the live network mode, router state file, line traffic, existing tunnel interfaces, the wall or monotonic clock, a cache in `/tmp` and third-party websites
- Correct behaviour: Native tests are hermetic by design: fixed inputs, injected clock and system interfaces, throwaway files, no network, no host tools, no live state; they pass on a fresh machine, in any order, three runs in a row (work plan 2.4; roadmap 747).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:619; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:583; CHANGELOG.md@ad3758c:202; CHANGELOG.md@ad3758c:203; CHANGELOG.md@ad3758c:199; CHANGELOG.md@ad3758c:323; docs/TESTING.md@ad3758c:159; git log@ad3758c: a0e04ec, 1059f2e, 1814fd1 (box tests saw the real tunnel interfaces)
- Python fix commit: a0e04ec, 1059f2e, 1814fd1, 316fcd4
- Package that proves it: F-00 (also F-06: the fakes)
- Spec vector: none yet (F-07)
- Native test, Go: gate self-test cases R05, R06 and R30 of the F-00 brief section 5 (planned: hermetic test runs)
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0048: Web UI texts untranslated in every language

- Area: Translations
- What went wrong: About 950 visible Web UI texts (labels, buttons, hints, table headers, placeholders, status words, script messages) had no translation key and were English in every language, Hungarian included; the router settings panel the same.
- Correct behaviour: Every visible text has a key and exists in all 11 languages, and the build fails on visible text without a key (CHANGELOG 198) and on a built-in key missing in any language (roadmap 602). Owner, 2026-11-11: a key that is still missing at run time (for example from an add-on) shows the English text, visibly marked as untranslated; it is logged and listed in the self-test result and in the diagnostic bundle; it is never a silent fallback to any language (this agrees with roadmap 406 and 602; KF-0049 says the same).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:583; CHANGELOG.md@ad3758c:198; CHANGELOG.md@ad3758c:280
- Python fix commit: ecd72dd, 9d6bca4
- Package that proves it: C-06 (also U-01, U-02, P-07, R-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0049: Hungarian text leaking into other languages

- Area: Translations
- What went wrong: Hungarian appeared in other languages: five input hints showed the Hungarian "pl." ("e.g.") everywhere, six script messages fell back to Hungarian, the latency chart's labels were Hungarian in every language, and older builds had Hungarian text in other languages in many places.
- Correct behaviour: No text, report or chart has Hungarian (or English) as a silent default for another language, and every report and chart label exists in every language (roadmap 602); a build check fails on Hungarian outside the Hungarian files and on a built-in key missing in any language. Owner, 2026-11-11: a key that is still missing at run time (for example from an add-on) shows the English text, visibly marked as untranslated; it is logged and listed in the self-test result and in the diagnostic bundle; it is never a silent fallback to Hungarian, English or any other language (roadmap 406 and 602; KF-0048 says the same).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:602; CHANGELOG.md@ad3758c:198; CHANGELOG.md@ad3758c:137; CHANGELOG.md@ad3758c:413; docs/TESTING.md@ad3758c:296 (release checklist 13)
- Python fix commit: ecd72dd, ad3758c, 48fd816, 22f0664, ebb1099
- Package that proves it: C-06 (also U-01, R-03, S-11, P-07, R-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0050: Fixed list of only 8 Wi-Fi countries

- Area: Wi-Fi
- What went wrong: The Wi-Fi country choice offered only 8 countries.
- Correct behaviour: Every country of the Linux wireless regulatory database, each with its DFS rule set, names in the UI's language; an unknown code is refused by the API, and one in the settings falls back to automatic (CHANGELOG 31); the country comes from the location setup (roadmap 600).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:583; CHANGELOG.md@ad3758c:31; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:600
- Python fix commit: b714f45
- Package that proves it: N-07 (also C-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0051: Diagnostic report changed the radio's regulatory domain

- Area: Diagnostics
- What went wrong: Making the diagnostic bundle could set the Wi-Fi regulatory domain (with the country on auto, it set the GeoIP country).
- Correct behaviour: A report only reads; only the access point sets the domain when it starts.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:583; CHANGELOG.md@ad3758c:199
- Python fix commit: 316fcd4
- Package that proves it: R-04 (also N-07)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0052: Time spans measured with the wall clock

- Area: Clocks
- What went wrong: Durations, intervals, timeouts and round-trip times were taken from the wall clock in several components (engines, schedules, DynDNS, the DNS cache, GeoIP, netd intervals), so a clock step changed them.
- Correct behaviour: Every duration, interval and timeout uses a monotonic clock; the wall clock is used only for stored timestamps (roadmap 583, an example of Quality Bar 6; the components beyond that example derived: Quality Bar 6, roadmap 581, no fragile timing is carried over).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:583; private audit 2026-11-10, engines item 21 and netd item L2
- Python fix commit: not recorded in the source
- Package that proves it: C-05 (also M-01, M-04, M-09, N-16, N-12)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0053: Logic compared translated UI text

- Area: Translations
- What went wrong: Code compared visible translated text instead of a state (the iPerf3 server button checked for the word "Active").
- Correct behaviour: Logic compares states and codes, never display text (roadmap 583, CHANGELOG 198).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:583; CHANGELOG.md@ad3758c:198
- Python fix commit: ecd72dd
- Package that proves it: C-06 (also U-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0054: Bot read a non-compliant verdict as compliant

- Area: Telegram bot
- What went wrong: The bot showed a non-compliant SLA result as compliant because it matched the shorter word inside the longer one; verdict keys are Hungarian words (`nmb/bot/i18n.py@ad3758c:1423`).
- Correct behaviour: Verdicts are codes from one source; the text is looked up from the code and never parsed back (derived: roadmap 583, "logic that compared translated UI text").
- Source: CHANGELOG.md@ad3758c:295; nmb/bot/i18n.py@ad3758c:358; nmb/bot/i18n.py@ad3758c:1423
- Python fix commit: 6466908
- Package that proves it: S-11 (also R-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: the matching; the Hungarian state keys remain

### KF-0055: Setup printed a password

- Area: Installer
- What went wrong: The installer printed the generated Wi-Fi password into its log, and stored a Web UI password the owner set in a notes file.
- Correct behaviour: No password or token is printed, logged or stored outside its own protected place; only flags such as "set" are shown (release checklist rows 2 and 22).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:592; CHANGELOG.md@ad3758c:289; CHANGELOG.md@ad3758c:211; docs/TESTING.md@ad3758c:285; docs/TESTING.md@ad3758c:305
- Python fix commit: 29c2ec8, 22eb38f
- Package that proves it: P-11 (also P-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0056: Factory login shared by every box

- Area: Web UI login
- What went wrong: The Web UI ships with a factory login that every box shares.
- Correct behaviour: Every native image gets a unique password per device, generated at manufacture or first boot from the device's own random source (never from the MAC, serial or another readable value), stored only as a hash, given to the user once, changeable at any time; a factory reset makes a new unique one (roadmap 592).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:592; CHANGELOG.md@ad3758c:211
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: P-03 (also S-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: the installer warns while the factory login is set

### KF-0057: Charts default to Hungarian

- Area: Reports
- What went wrong: The chart functions default to `lang="hu"`.
- Correct behaviour: No language default in report code: the language is always passed; every chart label exists in every language (roadmap 602, 604). Owner, 2026-11-11: the build refuses a built-in text missing in any language; a text still missing at run time shows the English text visibly marked as untranslated, is logged and is listed in the self-test result and the diagnostic bundle, never a silent fallback to any language (KF-0048).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:604; nmb/report/svg_charts.py@ad3758c:142; nmb/report/svg_charts.py@ad3758c:157; nmb/report/svg_charts.py@ad3758c:251; private audit 2026-11-10, Hungary-bound defaults
- Python fix commit: not recorded in the source
- Package that proves it: R-03 (also C-06)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0058: Report label only in Hungarian and English

- Area: Reports
- What went wrong: One dossier label ("6 (SLA Alatt)") existed only in Hungarian and English; other languages got English.
- Correct behaviour: Every report label exists in every language (roadmap 602). Owner, 2026-11-11: the build refuses a built-in text missing in any language; a text still missing at run time shows the English text visibly marked as untranslated, is logged and is listed in the self-test result and the diagnostic bundle, never a silent fallback to any language (KF-0048).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:604; private audit 2026-11-10, Hungary-bound defaults; git log -S'SLA Alatt'@ad3758c: removed from the dossier in da88eb1
- Python fix commit: da88eb1
- Package that proves it: R-03 (also C-06)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: source analysis: the label is no longer in the dossier at ad3758c; roadmap 604 still lists it

### KF-0059: Comments and test data name the owner's ISP and places

- Area: Code base
- What went wrong: Comments, documents and test data of the Python version name the owner's ISP, devices and places.
- Correct behaviour: Nothing personal or site-specific in source, tests, comments, documentation examples or images; reserved values only; every build scans for them and a hit fails the build (roadmap 594-598).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:598; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:604; private audit 2026-11-10, personal data scan
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: F-00 (the build scan; every package's gate)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0060: Invalid setting or % sign stopped the services

- Area: Configuration
- What went wrong: A single invalid setting or a `%` sign in a value stopped the services.
- Correct behaviour: Settings still load; a bad value is logged and replaced by its default; `%` is plain text, also in passwords (roadmap 629, 640).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:629; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:640; CHANGELOG.md@ad3758c:403
- Python fix commit: a7d119d
- Package that proves it: C-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0061: Wrong web port or address locked users out

- Area: Web server
- What went wrong: An invalid web port or address locked users out of the dashboard.
- Correct behaviour: Services stay up and the dashboard stays reachable with invalid web port, web address or ping interval values (roadmap 630).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:630; CHANGELOG.md@ad3758c:404
- Python fix commit: a7d119d
- Package that proves it: S-01 (also C-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0062: API crashed on invalid request bodies

- Area: Web server
- What went wrong: API endpoints failed on empty, invalid or non-object request bodies; missing imports made code paths fail only when they ran.
- Correct behaviour: Every POST endpoint answers an empty, invalid or non-object body without a crash or a server error; destructive actions refuse to run (roadmap 631; with the native API contract a 4xx, work plan 7 point 1); no code path can fail on a missing name at run time: the build fails on it, with warnings as errors (roadmap 588).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:631; CHANGELOG.md@ad3758c:412; docs/NATIVE_WORK_PLAN.md@ad3758c:306
- Python fix commit: a7d119d
- Package that proves it: S-01 (also S-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0063: `--help` changed something or failed

- Area: Command line
- What went wrong: The Phase 0 check "`--help` on every command" found failures in the Python version (roadmap 622: each check found real failures); the sources do not say which; the 2.5.0 release fixed incorrect README commands.
- Correct behaviour: `--help` on every command prints usage and changes nothing (roadmap 632); the command reference is generated from the same definitions (roadmap 6B.7).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:622; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:632; CHANGELOG.md@ad3758c:412
- Python fix commit: a7d119d; note: incorrect README commands (CHANGELOG 412); the failures of the check itself not recorded
- Package that proves it: P-10 (also S-10, P-12)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: assumed from the fixed clean-install round (CHANGELOG 401-413, a7d119d); which failures the check found is not recorded, and that it passes now is not verified

### KF-0064: Mode switching lost the address or the route

- Area: Network modes
- What went wrong: The Phase 0 check of switching between `standalone`, `mode_a`, `mode_b` and `auto` found failures in the Python version (roadmap 622); the sources do not say which beyond the entries for single mode faults in this register.
- Correct behaviour: After every switch the box keeps an IPv4 address and a route and can switch back (roadmap 633).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:622; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:633
- Python fix commit: a7d119d; note: bridge mode requests an IPv4 address (CHANGELOG 405, KF-0066); other failures of the check not recorded
- Package that proves it: N-03 (also N-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: assumed from the fixed clean-install round (CHANGELOG 401-413, a7d119d); which failures the check found is not recorded, and that it passes now is not verified

### KF-0065: Restart lost settings or the last snapshot

- Area: Database
- What went wrong: The Phase 0 restart check found failures in the Python version (roadmap 622); the sources do not say which beyond the restore and snapshot entries of this register.
- Correct behaviour: After a restart the settings and the last snapshot survive; at most about 15 minutes of data is lost after a power cut (roadmap 634; TESTING 171: the snapshot is atomic).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:622; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:634; docs/TESTING.md@ad3758c:171; docs/TESTING.md@ad3758c:398
- Python fix commit: not recorded in the source
- Package that proves it: C-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed; note: assumed: which failures the check found is not recorded, and that it passes now is not verified; TESTING 171 records a passed power-loss run on a VM (databases and settings intact)

### KF-0066: Bridge mode without IPv4, measuring on the wrong device

- Area: Network modes
- What went wrong: Bridge mode did not request an IPv4 address, did not measure on the bridge and read forwarded traffic from the wrong port.
- Correct behaviour: In bridge mode IPv6 link-local addresses do not count as an address; measurements use the bridge; traffic counters come from the WAN port (roadmap 641).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:641; CHANGELOG.md@ad3758c:405
- Python fix commit: a7d119d
- Package that proves it: N-03 (also M-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0067: Factory reset left an unreadable config

- Area: Factory reset
- What went wrong: A factory reset left a config the services could not read.
- Correct behaviour: A factory reset needs explicit confirmation and leaves a readable config and a unique password (roadmap 635, 639).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:635; CHANGELOG.md@ad3758c:406
- Python fix commit: a7d119d
- Package that proves it: P-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0068: Installation left failed units

- Area: Installer
- What went wrong: Installation left failed units and `nmb.target` did not run.
- Correct behaviour: A clean install exits 0 with no failed units and the service target active, and the dashboard answers (roadmap 628).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:628; CHANGELOG.md@ad3758c:407
- Python fix commit: a7d119d
- Package that proves it: P-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0069: Update verification touched the live database

- Area: Updates
- What went wrong: Update verification ran with too little time, from the wrong place, and touched the live database.
- Correct behaviour: Update verification never touches the live database (roadmap 642) and runs from the unpacked update.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:642; CHANGELOG.md@ad3758c:409
- Python fix commit: 03a3e3b
- Package that proves it: P-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0070: Skipped speed tests shown as 0 Mbps

- Area: Measurement data
- What went wrong: Skipped speed tests appeared as 0 Mbps and replaced real results; failed runs showed 0.00 Mbps; an upload-only iPerf3 run showed 0.00 Mbps download.
- Correct behaviour: A skipped test has no values; it is shown as skipped, never as 0, and never replaces a real result; a failed run shows its reason (roadmap 683; CHANGELOG 127).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:683; CHANGELOG.md@ad3758c:410; CHANGELOG.md@ad3758c:127; CHANGELOG.md@ad3758c:302
- Python fix commit: 1d7133b, 34515a9
- Package that proves it: M-04 (also U-02, R-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0071: iPerf3 server check connected to the port

- Area: iPerf3
- What went wrong: The check whether the local iPerf3 server runs connected to its port every few seconds, which logged a failed test each time.
- Correct behaviour: Whether a local service runs is read from the socket table; it is never connected to (roadmap 678).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:678; CHANGELOG.md@ad3758c:411
- Python fix commit: not recorded in the source
- Package that proves it: M-09 (also P-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0072: Impossible Cloudflare results

- Area: Cloudflare engine
- What went wrong: Cloudflare produced impossible results (an upload timing fault gave about 80000 Mbps; a rate-limited download was taken as the result).
- Correct behaviour: A rate is computed from the bytes really moved over the measured time of that transfer; a refused or throttled transfer is an error, not a speed (derived: QB6, roadmap 581).
- Source: CHANGELOG.md@ad3758c:412; git log@ad3758c: c2fca70 (subject)
- Python fix commit: c2fca70, a3ad78a
- Package that proves it: M-06
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0073: Web UI layout broke on phones; theme did not persist; DynDNS card stale

- Area: Web UI
- What went wrong: Mobile layouts broke (sub-navigation, tables, charts, a one-letter-wide card title), the schedule line overflowed and showed its label twice, and the theme did not persist; after a save the DynDNS card kept showing "never" and the old error until a reload, and it named DuckDNS while dynv6 was in use.
- Correct behaviour: Every page works at phone width without horizontal scrolling, text wraps, and the look does not change by itself (roadmap 4.4A, 172: works on a phone screen; no horizontal scrolling and wrapped text derived from it: Quality Bar 6, roadmap 581, no wrong result is carried over); after a save a card fetches its live status again once the action has run, and it names the provider really in use (CHANGELOG 260).
- Source: CHANGELOG.md@ad3758c:413; CHANGELOG.md@ad3758c:306; CHANGELOG.md@ad3758c:260
- Python fix commit: 0615352, a25ae6e, cd69f23, 8aa0efc
- Package that proves it: U-01 (also U-02, N-16)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0074: Callers used functions whose arguments had changed

- Area: Telegram bot and repair
- What went wrong: Self-heal from Telegram and alerts called the repair function wrongly; nine bot commands (`/heal`, `/wifi`, `/wifi scan`, `/wifi connect`, `/tasks`, `/kill`, `/tailscale`, `/backup`, `/drivers`) called functions whose arguments or results had changed, while mocks that invented return values kept the tests green.
- Correct behaviour: Interfaces are typed and checked at build time; test fakes have the real interface's shape (TESTING 255: mocks must match the real code), so a changed call fails the build or the tests.
- Source: CHANGELOG.md@ad3758c:408; CHANGELOG.md@ad3758c:292; docs/TESTING.md@ad3758c:167; docs/TESTING.md@ad3758c:255
- Python fix commit: 03a3e3b, 42a0fd9
- Package that proves it: F-06 (also S-11, P-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0075: Destructive actions without confirmation

- Area: System
- What went wrong: Factory reset and log purge ran without explicit confirmation in the CLI and the web API.
- Correct behaviour: Destructive actions need confirmation: factory reset (`--yes` or `RESET`), data and log purge (`PURGE`) (roadmap 639). Owner, 2026-11-11: in the Telegram bot every command that changes the box needs a confirmation tap, and the destructive ones (factory reset, purge) also need the typed word, as in the Web UI; read-only commands run at once.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:639; CHANGELOG.md@ad3758c:417
- Python fix commit: a7d119d, 0b51025
- Package that proves it: P-03 (also S-02, S-10, S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0076: Shared factory Wi-Fi password

- Area: Wi-Fi
- What went wrong: Boxes shared one factory Wi-Fi password that was public in the repository; it stayed a fallback in the defaults, the example config and the restore path.
- Correct behaviour: Each box gets its own Wi-Fi password, also after a factory reset; the access point does not start with the factory, an empty or an out-of-range (8-63 characters) password, the UI refuses to save one, and the API reports only whether a password is set.
- Source: CHANGELOG.md@ad3758c:418; CHANGELOG.md@ad3758c:289; CHANGELOG.md@ad3758c:270; CHANGELOG.md@ad3758c:332; docs/TESTING.md@ad3758c:182
- Python fix commit: 29c2ec8, d6202e4, bbcf9b6
- Package that proves it: N-07 (also P-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0077: Broad sudo, no login rate limit or CSRF protection, unescaped values

- Area: Web server
- What went wrong: Early versions had broad sudo rights, no brute-force or CSRF protection, and external values rendered into the dashboard without escaping.
- Correct behaviour: Rate-limited login, CSRF protection, constant-time comparisons, every external value escaped, a strict Content Security Policy (roadmap 589).
- Source: CHANGELOG.md@ad3758c:419; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:589
- Python fix commit: 5383820, 78841b6, 95b1e4d
- Package that proves it: S-01 (also U-01, P-06)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0078: Personal data in the repository; subscriber data in the LLM bundle

- Area: Privacy
- What went wrong: Personal data was in the repository, and subscriber data was not masked in the LLM bundle.
- Correct behaviour: No personal or site data in the code base (roadmap 594-598); exports mask subscriber data (roadmap 4.24).
- Source: CHANGELOG.md@ad3758c:420; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:594
- Python fix commit: 701df66, 230597e, 21ffff5, 8e6f341
- Package that proves it: F-00 (also R-04, C-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0079: Root ran code from /tmp and writable folders

- Area: Privileged helper
- What went wrong: Self-repair and speed test code could run files from `/tmp` (a way for local users to gain root).
- Correct behaviour: Never run code from `/tmp` or from folders the service user can write to (roadmap 679).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:679; CHANGELOG.md@ad3758c:348
- Python fix commit: b4f472f
- Package that proves it: P-06 (also P-01, M-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0080: Ookla download installed without checksum

- Area: Ookla engine
- What went wrong: The Ookla download was installed without checking its checksum.
- Correct behaviour: Every downloaded binary is installed only when its checksum (native: its signature or pinned hash) matches (roadmap 679).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:679; CHANGELOG.md@ad3758c:358
- Python fix commit: ea65ff7
- Package that proves it: M-10 (also P-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0081: Sudo rules that gave root

- Area: Privileged helper
- What went wrong: The service user could run `ip` through sudo, which allowed root access; a sudo rule let it write the settings file through `tee`.
- Correct behaviour: No tool with an execution option and no file-writing tool in sudoers or in any privileged allow-list; NMB writes its own settings without such a rule (roadmap 680; CHANGELOG 344).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:680; CHANGELOG.md@ad3758c:349; CHANGELOG.md@ad3758c:344; private audit 2026-11-10, privileged-operations item 3
- Python fix commit: b4f472f, e6512aa
- Package that proves it: P-06
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0082: Helper accepted VPN configs with commands and option-like names

- Area: Privileged helper
- What went wrong: The helper accepted VPN configs with command lines and package names that look like options.
- Correct behaviour: Helper arguments are validated; VPN configs with command lines are refused; names starting with `-` are refused (roadmap 680).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:680; CHANGELOG.md@ad3758c:350
- Python fix commit: b4f472f
- Package that proves it: P-06 (also N-14)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0083: Web server and web terminal ran as root

- Area: Web server
- What went wrong: The web server ran as root; an early change gave the web terminal full root and passwordless sudo, and the CLI terminal and task manager acted as root.
- Correct behaviour: The web server never runs as root; it runs as its own user with only the capabilities it needs (roadmap 681, 589).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:681; CHANGELOG.md@ad3758c:357; CHANGELOG.md@ad3758c:128; git log@ad3758c: 022bb73 (subject)
- Python fix commit: ea65ff7, b4f472f
- Package that proves it: S-01 (also P-04, S-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0084: Local service ports accepted public addresses

- Area: Firewall
- What went wrong: The Web UI, DNS and iPerf3 ports accepted public source addresses.
- Correct behaviour: The local service ports (Web UI, DNS, iPerf3) answer only on the networks they serve (the LAN ports, the VPNs, Tailscale and an uplink network the owner marked as their own), on IPv4 and IPv6; on any other uplink they are closed whatever source address a packet carries (derived: roadmap 587). Roadmap 681 says these ports "must refuse public addresses on IPv4 and IPv6"; the owner's answer of 2026-11-11 below replaces that address-class wording, and roadmap 681 needs a correction (outside F-03). Owner, 2026-11-11: the role of a network the box is plugged into is never inferred from address classes; nothing is open from the uplink side (Web UI, SSH, DNS) until the owner marks that network as their own; the first-start setup offers the network it was done from and the box remembers it; an unknown or new network gets nothing until it is confirmed; the LAN ports, the VPNs and Tailscale are not affected.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:681; CHANGELOG.md@ad3758c:355
- Python fix commit: ea65ff7
- Package that proves it: N-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0085: Secrets copied to the boot partition

- Area: Configuration
- What went wrong: Passwords and tokens, and a chosen admin password, were copied to the boot partition.
- Correct behaviour: Secrets stay off the boot partition; the boot copy carries `__KEEP_CURRENT__` for them, and that mark keeps existing boot files valid (roadmap 682; CHANGELOG 129).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:682; CHANGELOG.md@ad3758c:356; CHANGELOG.md@ad3758c:129
- Python fix commit: ea65ff7
- Package that proves it: C-01 (also C-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0086: Two writers shared one temporary file and broke the settings

- Area: Configuration
- What went wrong: Two programs saving at once wrote the same temporary file and left a broken `nmb.conf` that no program could read.
- Correct behaviour: Settings are saved one at a time (a lock across threads and processes) through a temporary file of the writer's own (roadmap 684); release checklist row 12.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:684; CHANGELOG.md@ad3758c:298; docs/TESTING.md@ad3758c:178; docs/TESTING.md@ad3758c:295
- Python fix commit: 10545ab
- Package that proves it: C-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0087: Ookla CLI had no writable home

- Area: Ookla engine
- What went wrong: The Ookla CLI ran as the service user with no writable home and could not keep its settings, so every test started as a new client.
- Correct behaviour: Programs run as an unprivileged user get a writable home for the tools they call (roadmap 685).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:685; CHANGELOG.md@ad3758c:297
- Python fix commit: 002548c
- Package that proves it: M-10 (also P-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0088: `iw` not found outside PATH

- Area: Wi-Fi
- What went wrong: `iw` was not on the service user's path, so the Wi-Fi status lost its SSID and signal and the access-point check never saw an access point.
- Correct behaviour: System tools are looked for outside `PATH` too (`/usr/sbin`, `/sbin`) (roadmap 686); native code reads the kernel interfaces directly where it can (roadmap 4.19).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:686; CHANGELOG.md@ad3758c:299
- Python fix commit: 09db7e6
- Package that proves it: N-07 (also M-14, N-17)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0089: DHCP guard read offers on a normal UDP socket

- Area: DHCP guard
- What went wrong: A normal UDP socket did not receive another server's offer, so the rogue-DHCP check would have passed while another server answered.
- Correct behaviour: DHCP replies are read from a packet socket (roadmap 687).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:687; CHANGELOG.md@ad3758c:79; docs/TESTING.md@ad3758c:52
- Python fix commit: 8a6ff52
- Package that proves it: N-05 (also N-08)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0090: DHCP guard did not wait long enough

- Area: DHCP guard
- What went wrong: The guard did not wait long enough for an offer; dnsmasq and many routers ping an address for up to 3 s first.
- Correct behaviour: Wait 8 s with three requests for DHCP offers (roadmap 688).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:688; docs/TESTING.md@ad3758c:52
- Python fix commit: 8a6ff52
- Package that proves it: N-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0091: Supervised daemon told to log to stderr in the foreground

- Area: Services
- What went wrong: dnsmasq started with `--log-facility=-` stayed in the foreground, so the start command never returned and timed out.
- Correct behaviour: A daemon started by a supervisor is never told to log to stderr in the foreground (roadmap 689).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:689; docs/TESTING.md@ad3758c:52
- Python fix commit: 8a6ff52
- Package that proves it: N-05 (also P-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0092: Fallback results stored as Ookla

- Area: Ookla engine
- What went wrong: When Ookla refused or paused, a one-second Cloudflare download (later RTR or Cloudflare) was stored as an Ookla result under the chosen server's name, and the home page kept showing an old Ookla result.
- Correct behaviour: Store no other engine's result in Ookla's place; a fallback test is stored under its own engine; Ookla failures are reported as such (roadmap 690).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:690; CHANGELOG.md@ad3758c:325; CHANGELOG.md@ad3758c:307; docs/TESTING.md@ad3758c:172
- Python fix commit: 0847ad0, 8aa0efc
- Package that proves it: M-10 (also M-04, U-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0093: Ookla pause per server and only in memory

- Area: Ookla engine
- What went wrong: A rate limit quarantined one server after another and silently swapped the chosen server; the pause was kept only in memory, so every restart asked Ookla again early and made its block longer.
- Correct behaviour: A refusal pauses Ookla for the whole box (30 min doubling to 4 h), kept across restarts; the chosen server is kept, and a result from another server is flagged; the IP is never changed to get around a block (roadmap 690).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:690; CHANGELOG.md@ad3758c:326; CHANGELOG.md@ad3758c:308; docs/TESTING.md@ad3758c:397
- Python fix commit: 0847ad0, 8aa0efc
- Package that proves it: M-05 (also M-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0094: Scheduled Ookla more often than hourly

- Area: Ookla engine
- What went wrong: Schedules ran Ookla more often than hourly, which Ookla answers by blocking the public IP.
- Correct behaviour: Scheduled Ookla runs at most hourly however often the schedule runs (`ookla_interval_min` never below 60), with a warning when the schedule asks for more (roadmap 690, 4.25). Owner, 2026-11-11: a test started by a person (Web UI, CLI, bot) is not held back by the hourly schedule limit, but during a refusal pause it is refused with the time left.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:690; CHANGELOG.md@ad3758c:126; CHANGELOG.md@ad3758c:98; CHANGELOG.md@ad3758c:92; docs/TESTING.md@ad3758c:162
- Python fix commit: 12d5a3e, 8aa0efc
- Package that proves it: M-05 (also M-10, M-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0095: Ookla runs during a pause stored as failed tests

- Area: Ookla engine
- What went wrong: Scheduled runs during a rate-limit pause were stored as failed tests.
- Correct behaviour: Runs during a pause are stored as skipped with the pause as reason, and add no measured values (CHANGELOG 268; roadmap 683).
- Source: CHANGELOG.md@ad3758c:268; docs/TESTING.md@ad3758c:341
- Python fix commit: d6202e4
- Package that proves it: M-04 (also M-05, M-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0096: Clearing the Ookla pause on disk did not end it

- Area: Ookla engine
- What went wrong: Deleting the pause file alone did not end the pause until the web server restarted. At ad3758c the loader returns without a change when the file is missing, so the pause held in memory stays.
- Correct behaviour: When an in-memory pause is cleared on disk, the running process reloads it, or a command clears it (roadmap 691).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:691; nmb/speed/anti_ban.py@ad3758c:137-145; nmb/speed/anti_ban.py@ad3758c:112-121
- Python fix commit: not recorded in the source
- Package that proves it: M-05 (also M-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: source analysis: a missing file leaves the in-memory pause; a clearing function exists

### KF-0097: Access point in standalone cut the box off its Wi-Fi

- Area: Wi-Fi
- What went wrong: `wifi_enabled = true` turned the radio into an access point in standalone mode too, which cut the box off its Wi-Fi network.
- Correct behaviour: The access point runs only in router mode; in standalone and bridge mode the radio stays a client of the home network and is handed back and reconnected if an access point held it.
- Source: CHANGELOG.md@ad3758c:119
- Python fix commit: ec00296
- Package that proves it: N-07 (also N-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0098: Country-specific values built in

- Area: Location
- What went wrong: A built-in Ookla server list had wrong ids (one provider's id answered "No servers defined"), the Wi-Fi country was not automatic, ISP placeholders named Hungarian providers, and status codes were regulator-specific.
- Correct behaviour: No country built in: servers from the engine's live list near the box, the country from the location setup, country data in replaceable files, neutral codes (roadmap 599-603).
- Source: CHANGELOG.md@ad3758c:124; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:601; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:603; docs/TESTING.md@ad3758c:172
- Python fix commit: 8aa0efc
- Package that proves it: C-10 (also M-10, N-07, R-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0099: Unchecked claims of official status and conformity

- Area: Reports
- What went wrong: The bot called the box's own evaluation "official" or "regulatory"; the dossier claimed "RFC 2681 / BEREC compliant" with a fixed version number and that tests ran "strictly under verified low background traffic"; the Hungarian appendix cited a national act.
- Correct behaviour: Reports say what the box does and claim nothing that was not checked: NMB's own evaluation of its own measurements, the installed version, the traffic gate as it works (traffic it cannot see is not excluded); legal references come only from the box's country file (roadmap 601).
- Source: CHANGELOG.md@ad3758c:116; CHANGELOG.md@ad3758c:151; CHANGELOG.md@ad3758c:117
- Python fix commit: eb6eb93
- Package that proves it: R-03 (also S-11, C-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0100: Web UI header showed fixed or missing values

- Area: Web UI
- What went wrong: The header's "WAN Interface" showed `eth0` whatever the box used, "System Time" always `--`, and the network mode always "standalone": the page read fields the server never sent and fell back to fixed words.
- Correct behaviour: The header shows the mode netd runs (with PPPoE and auto marked), the current link and the path to the ISP from the recorded setup, in all languages; a field the server does not send is shown as unknown, never as a fixed word (derived: roadmap 4.29 453).
- Source: CHANGELOG.md@ad3758c:137; CHANGELOG.md@ad3758c:226
- Python fix commit: 0eddf2b, ad3758c
- Package that proves it: U-02 (also S-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0101: Bot showed raw codes

- Area: Telegram bot
- What went wrong: `/status` and `/mode` showed the raw mode code (`MODE_A`).
- Correct behaviour: Every code shown to a person is translated text from one source.
- Source: CHANGELOG.md@ad3758c:138
- Python fix commit: ad3758c
- Package that proves it: S-11 (also C-06)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0102: Bot rebuilt the diagnostic bundle on every request

- Area: Diagnostics
- What went wrong: The diagnostic bundle was rebuilt on every `/ai` request instead of using its cache (an argument was passed in the wrong place).
- Correct behaviour: A cached bundle is used unless the user asks for a fresh one.
- Source: CHANGELOG.md@ad3758c:138; private audit 2026-11-10, bot item L6
- Python fix commit: ad3758c
- Package that proves it: R-04 (also S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0103: Self-tests left their data behind

- Area: Self-test
- What went wrong: The live self-test left its test chats' language rows in the database.
- Correct behaviour: A self-test leaves nothing behind: every throwaway object it made is removed (TESTING 231).
- Source: CHANGELOG.md@ad3758c:138; docs/TESTING.md@ad3758c:231
- Python fix commit: ad3758c
- Package that proves it: P-07
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0104: Unknown values shown as 100 % or as total loss

- Area: Reports
- What went wrong: Without ping data the uptime showed 100 %; without judged tests the SLA verdict said "meets targets" with 100 % compliance; `/ping` showed 100 % loss for a target without samples and listed switched-off targets.
- Correct behaviour: Nothing measured is no verdict: unknown uptime and compliance show "-"/"--"/N/A on every surface; a target without samples shows "--"; only switched-on targets are listed; `/test` says when a guaranteed speed is set but that direction was not measured.
- Source: CHANGELOG.md@ad3758c:140; CHANGELOG.md@ad3758c:154; CHANGELOG.md@ad3758c:155
- Python fix commit: da88eb1, 428b608
- Package that proves it: R-01 (also S-11, R-03, R-05, S-02, S-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0105: Local gateway counted into the line's loss

- Area: Reports
- What went wrong: The local gateway counted into the line's loss, latency and uptime timeline, so a gateway that filters ping made the line look lossy.
- Correct behaviour: The gateway stays in the per-target table but never counts into the line's figures.
- Source: CHANGELOG.md@ad3758c:141
- Python fix commit: 428b608
- Package that proves it: R-01 (also M-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0106: Outage before the report window judged by its clamped start

- Area: Reports
- What went wrong: An outage that began before the report window was judged by the clamped start.
- Correct behaviour: An outage is judged by its real start (whether it lies in the box's own windows), then clamped for the figures.
- Source: CHANGELOG.md@ad3758c:142
- Python fix commit: 428b608
- Package that proves it: R-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0107: LAN tests counted as the ISP's line

- Area: Measurement data
- What went wrong: Tests against an iPerf3 server on the local network counted into the line's figures (dossier, charts, CSV shares, averages, SLA card, bot, bundle, collector), and such a server was named after the box's own town and ISP.
- Correct behaviour: A test against a LAN server is marked LAN everywhere (also older rows, by the server's local address), left out of the line's figures and averages, has no share of the offered speed and is named "LAN" as ISP.
- Source: CHANGELOG.md@ad3758c:143; CHANGELOG.md@ad3758c:157; CHANGELOG.md@ad3758c:168
- Python fix commit: b8fb934, da88eb1, 428b608
- Package that proves it: R-01 (also R-02, R-03, S-12, S-13, M-09, U-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0108: LLM bundle invented verdicts

- Area: Diagnostics
- What went wrong: The LLM bundle showed 0.00 % loss for a target without samples, called a gateway that does not answer pings "loss inside the local network", treated unset contracts as values, mixed LAN tests into the engine statistics and counted upload-only tests as failed.
- Correct behaviour: "No ping data" when there is none; the gateway-silent verdict only when internet targets have data; unset speeds "not set"; LAN tests apart; a test is judged only on the directions it measured.
- Source: CHANGELOG.md@ad3758c:144; CHANGELOG.md@ad3758c:160; private data-quality list 2026-11-10, second review
- Python fix commit: da88eb1, 428b608
- Package that proves it: R-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0109: RTR direction without bytes stored as 0 bit/s

- Area: RTR engine
- What went wrong: A direction without transferred bytes was stored as 0 bit/s; a run without any counted as success; the iPerf3 command record named directions that did not run.
- Correct behaviour: A direction without bytes is "not measured"; a run without any is an error; the stored command record names what really ran.
- Source: CHANGELOG.md@ad3758c:145; private audit 2026-11-10, engines item 24
- Python fix commit: 428b608
- Package that proves it: M-08 (also M-09)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0110: Made-up values in old rows

- Area: Database
- What went wrong: Values filled in instead of measured were stored in many rows (57,735 on one box); they were cleared once in three passes with every original kept in `speedtest_correction`.
- Correct behaviour: The native build never stores a made-up value; reading a Python database it keeps `speedtest_correction` and the cleared values as they are and never re-creates them (derived: roadmap 456 "old data", QB6 581).
- Source: CHANGELOG.md@ad3758c:181; CHANGELOG.md@ad3758c:146; CHANGELOG.md@ad3758c:162
- Python fix commit: c2d0ed6, da88eb1, 428b608
- Package that proves it: C-02 (also S-13)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0111: Per-minute ping CSV built in memory with an open read transaction

- Area: Exports
- What went wrong: The ping CSV was built in memory (up to 120 days, about a million rows), then read with one transaction open for a whole slow download and no timeout; a minute without pings had loss 0; the public path had no range limit.
- Correct behaviour: Exports are streamed in pages, each finished before it is written; a client that stops reading times out; a minute without pings has no loss figure; ranges are limited (31 days without login, 120 with).
- Source: CHANGELOG.md@ad3758c:147; CHANGELOG.md@ad3758c:153; CHANGELOG.md@ad3758c:167
- Python fix commit: eb6eb93, 428b608, da88eb1
- Package that proves it: R-02 (also S-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0112: Buttons pointing to parts of the UI that no longer exist

- Area: Web UI
- What went wrong: The default-password notice's button opened a tab that no longer existed; "purge all logs" failed on a missing function; "Stop iPerf Server" did nothing because the web server could not stop the root unit.
- Correct behaviour: Every button reaches an existing page or action, checked by a build test (CHANGELOG 150); actions go through the service control interface, never `pkill`.
- Source: CHANGELOG.md@ad3758c:150; CHANGELOG.md@ad3758c:198; CHANGELOG.md@ad3758c:303
- Python fix commit: dfba323, ecd72dd, 34515a9
- Package that proves it: U-02 (also U-01, S-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0113: Dossier escaped values twice

- Area: Reports
- What went wrong: The dossier escaped server names and outage causes twice (`&amp;amp;`).
- Correct behaviour: Every value is escaped exactly once, at output.
- Source: CHANGELOG.md@ad3758c:152
- Python fix commit: da88eb1
- Package that proves it: R-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0114: Invented statistics in the dossier and the snapshot

- Area: Reports
- What went wrong: The dossier's packet loss minimum (always 0.00 %) and maximum (twice the average) were invented; the developer snapshot read columns that do not exist and always showed 0 Mbit/s; the bundle's last test was not the newest successful one.
- Correct behaviour: Every figure comes from data; a statistic that cannot be computed is not shown (CHANGELOG 156).
- Source: CHANGELOG.md@ad3758c:156
- Python fix commit: da88eb1
- Package that proves it: R-03 (also C-03, R-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0115: ISP ticket e-mail described a fixed gate and threshold

- Area: E-mail
- What went wrong: The ISP ticket e-mail claimed "<500 kbps" and an "80 % SLA threshold" for every box.
- Correct behaviour: The e-mail describes the traffic gate and the required speed as they are set.
- Source: CHANGELOG.md@ad3758c:159
- Python fix commit: da88eb1
- Package that proves it: R-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0116: Wrong skip code

- Area: Status codes
- What went wrong: A skip because other devices' traffic is not visible was code 9 ("other traffic"); the bot used codes 14 and 15, which the dossier's legend uses for Wi-Fi measurements.
- Correct behaviour: One code list for every surface: not visible = 10, below the guaranteed speed = 6 with the direction; codes have a legend (spec/ codes, roadmap 6B.8 point 1).
- Source: CHANGELOG.md@ad3758c:161; CHANGELOG.md@ad3758c:170; CHANGELOG.md@ad3758c:188; private data-quality list 2026-11-10, item D6
- Python fix commit: da88eb1, b8fb934
- Package that proves it: R-01 (also S-11, R-02, F-07)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0117: Engine interval counted from the engine's own start

- Area: Scheduling
- What went wrong: With per-engine start times an engine's interval counted from its own start, so hourly iPerf3 skipped every other hour; aborted runs stored no facts; RTR fallback rows did not say where their address came from.
- Correct behaviour: An engine's interval counts from the start of the run it belonged to; aborted runs store their facts; every address states its source.
- Source: CHANGELOG.md@ad3758c:164
- Python fix commit: da88eb1
- Package that proves it: M-04 (also M-08)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0118: Silent ISP first hop and PPP peer pinged as targets

- Area: Ping probe
- What went wrong: An ISP first hop that answers no pings (a PPP peer placeholder, a filtering hop) was pinged as a target and made up "100 % loss"; once its lost pings aged out it was switched on again every day.
- Correct behaviour: A first hop that answered none of at least 2,000 pings in 24 h is not pinged (remembered 30 days); a PPP peer placeholder seen while the box dialled PPPoE is remembered 90 days and never taken as the first hop; the first hop can be switched off.
- Source: CHANGELOG.md@ad3758c:165; CHANGELOG.md@ad3758c:179; CHANGELOG.md@ad3758c:39; private data-quality list 2026-11-10, item D8
- Python fix commit: 96832e9, da88eb1
- Package that proves it: M-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0119: Collector token could write or take another box's data

- Area: Central collector
- What went wrong: Any valid token could write into another box's data and take, fake or refuse its remote commands; a re-sent row was ignored, so the collector kept made-up values.
- Correct behaviour: A box belongs to the token it first uploaded with until that token is revoked; uploads and polls check the binding; the box sends corrections for rows it sent before and the collector applies them to that box's rows only.
- Source: CHANGELOG.md@ad3758c:166; CHANGELOG.md@ad3758c:171; private audit 2026-11-10, bot and collector item H3
- Python fix commit: b8fb934, da88eb1
- Package that proves it: S-13 (also S-12)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0120: SLA ignored a set guaranteed speed

- Area: Reports
- What went wrong: A test was judged against 80 % of the offered speed even when the owner had set a guaranteed speed.
- Correct behaviour: A test meets the required speed when it reaches the guaranteed speed the owner set; without one, 80 % of the offered speed; a direction not measured is not failed.
- Source: CHANGELOG.md@ad3758c:169
- Python fix commit: b8fb934
- Package that proves it: R-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0121: Path MTU reported when nothing answered

- Area: Diagnostics
- What went wrong: The bundle reported a path MTU of 1480 bytes and "PPPoE confirmed" when nothing answered; the MSS advice did not depend on who clamps it.
- Correct behaviour: The path MTU is "unknown" when nothing answers; the advice depends on whether this box clamps the MSS.
- Source: CHANGELOG.md@ad3758c:175; private data-quality list 2026-11-10, item D7
- Python fix commit: 96832e9
- Package that proves it: R-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0122: MTR marked loaded only for the own process

- Area: Diagnostics
- What went wrong: The MTR counted as loaded only when a speed test ran in its own process; a 10-packet MTR was titled multi-sample.
- Correct behaviour: "Loaded" only when a speed test ran during it in any process on the box (the box-wide measurement lock); titles say what was measured.
- Source: CHANGELOG.md@ad3758c:176; private data-quality list 2026-11-10, item D7
- Python fix commit: 96832e9
- Package that proves it: R-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0123: Event severities that the schema refuses or never stores

- Area: Events
- What went wrong: The USB fallback events used severity `warning`, which the schema refuses, so they were never stored; the bundle looked for `warning`, `error` and `critical`, which are never stored, so it never listed a warning. At ad3758c the storage guard still writes `warning` (`nmb/common/storage_guard.py@ad3758c:250` against the CHECK in `nmb/db/schema.sql@ad3758c:102`), so its cleanup event is lost.
- Correct behaviour: Severities are one typed list (`debug`, `info`, `warn`, `crit`); a test inserts every event kind into the real schema (TESTING 180).
- Source: CHANGELOG.md@ad3758c:177; docs/TESTING.md@ad3758c:180; private audit 2026-11-10, reports item 8; nmb/common/storage_guard.py@ad3758c:250; nmb/db/schema.sql@ad3758c:102
- Python fix commit: c72fa91, 96832e9
- Package that proves it: C-07 (also C-08, R-04, F-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: partly fixed: the storage guard event is still refused, source analysis

### KF-0124: RTR invented ping, jitter and place

- Area: RTR engine
- What went wrong: RTR rows carried a fixed 0.5 ms jitter (hard-coded 500 µs), a jitter computed as median minus minimum, a ping minimum equal to the median, an invented 15 ms ping, estimated loaded latency and bufferbloat, a fixed place and colo, 0 ms when no ping came back, and the project's home page as the result link; the native client did not bind to the asked link and the row had no client IP.
- Correct behaviour: The real minimum and jitter (mean change between pings, from the server's timings) or empty; the measurement server RTR assigned is stored; the client binds to the asked link and the row has its client IP.
- Source: CHANGELOG.md@ad3758c:183; CHANGELOG.md@ad3758c:267; private data-quality list 2026-11-10, items D2 and D3; private audit 2026-11-10, engines item 9
- Python fix commit: d6202e4, c2d0ed6
- Package that proves it: M-08
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0125: Bufferbloat grade without loaded samples

- Area: Speed engines
- What went wrong: A bufferbloat grade (A+ with 0 ms) was given without any latency sample under load; a missing grade showed a made-up `A+`.
- Correct behaviour: A grade only from the directions really sampled under load; otherwise none (`--`).
- Source: CHANGELOG.md@ad3758c:184; CHANGELOG.md@ad3758c:301; CHANGELOG.md@ad3758c:305
- Python fix commit: c2d0ed6
- Package that proves it: M-06 (also M-07, M-10, U-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0126: Cloudflare client place shown as server place; single-sample jitter

- Area: Cloudflare engine
- What went wrong: The city and country Cloudflare reports (the box's own place) were stored as the server's place; one latency sample gave jitter 0; no sample under load gave the idle latency as loaded latency.
- Correct behaviour: Server place "Cloudflare <colo>"; jitter needs at least two samples; no loaded latency without a loaded sample.
- Source: CHANGELOG.md@ad3758c:185; CHANGELOG.md@ad3758c:146; CHANGELOG.md@ad3758c:162
- Python fix commit: c2d0ed6, 428b608
- Package that proves it: M-06
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0127: iPerf3 values in the wrong columns

- Area: iPerf3
- What went wrong: The block size (or MSS, or "Auto") was stored as the TCP window, a missing retransmit count as 0, the TCP RTT under upload load in the ping column, and the port as the colo.
- Correct behaviour: Each value in its own column or empty: the largest congestion window when iPerf3 reports it, retransmits only when reported, the loaded RTT as loaded latency.
- Source: CHANGELOG.md@ad3758c:186
- Python fix commit: c2d0ed6
- Package that proves it: M-09
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0128: Detailed CSV with Hungarian file name and wrong defaults

- Area: Exports
- What went wrong: The detailed CSV's file name was Hungarian in every language, Wi-Fi tests were not recognised from their interface, "below the guaranteed speed" ignored the upload, and the speed CSV wrote "Unknown ISP" or "ookla" for an unknown engine; iPerf3 rows were labelled OOKLA in the list.
- Correct behaviour: Neutral file names; Wi-Fi tests recognised from the interface they ran on; both directions judged; an unknown value is empty, never another engine's name.
- Source: CHANGELOG.md@ad3758c:188; CHANGELOG.md@ad3758c:302; private data-quality list 2026-11-10, item D6
- Python fix commit: c2d0ed6, 34515a9
- Package that proves it: R-02 (also U-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0129: Export cells could act as spreadsheet formulas

- Area: Exports
- What went wrong: Text from outside the box (server names, error texts, event details) in a CSV could act as a spreadsheet formula.
- Correct behaviour: Every text cell whose first character a spreadsheet reads as the start of a formula gets a leading quote: at least `=`, `+`, `-`, `@`, TAB and CR, as the Python fix does (`nmb/report/csv_export.py@ad3758c:68-72`); the missing-value mark `--` stays as it is; one test vector for each of these characters.
- Source: CHANGELOG.md@ad3758c:23; nmb/report/csv_export.py@ad3758c:68-72; private audit 2026-11-10, reports item 19
- Python fix commit: 1012b59
- Package that proves it: R-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0130: Router mode with a public address left services open

- Area: Firewall
- What went wrong: In router mode with a public address (PPPoE) only the Web UI was closed to the internet; other services of the box were reachable on IPv4 and IPv6.
- Correct behaviour: The WAN lets in only answers to the box's own traffic, ICMP, DHCP renewals and the VPN ports that are on or running (KF-0021); WAN to LAN only for answers and forwarded ports (CHANGELOG 347). Owner, 2026-11-11: the role of a network the box is plugged into is never inferred from address classes; nothing is open from the uplink side (Web UI, SSH, DNS) until the owner marks that network as their own; the first-start setup offers the network it was done from and the box remembers it; an unknown or new network gets nothing until it is confirmed; the LAN ports, the VPNs and Tailscale are not affected. On the uplink no source is trusted for its address alone, and reverse-path (anti-spoofing) filtering is on for IPv4 and IPv6 (derived: roadmap 587, nothing listens on the WAN side unless the owner turned it on); checked from another network on IPv4 and IPv6.
- Source: CHANGELOG.md@ad3758c:347; docs/TESTING.md@ad3758c:86; private log review 2026-11-09 evening, item 5
- Python fix commit: 2272614
- Package that proves it: N-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0131: Diagnostic bundle failed on one unreadable log

- Area: Diagnostics
- What went wrong: A log file the bundle may not read made the whole bundle fail.
- Correct behaviour: An unreadable part is skipped and named; the rest of the bundle is made.
- Source: CHANGELOG.md@ad3758c:205
- Python fix commit: bfbba44
- Package that proves it: R-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0132: Files the installer or README needs missing from the repository

- Area: Packaging
- What went wrong: The USB settings template and the boot drop-in were kept out of the repository by an ignore rule, so an update was rightly refused.
- Correct behaviour: A build check fails when a file the installer installs or the documentation names is missing (CHANGELOG 200).
- Source: CHANGELOG.md@ad3758c:200
- Python fix commit: b56b0c4
- Package that proves it: P-11 (also P-12)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0133: Boot waited for the network manager in router mode

- Area: Services
- What went wrong: `NetworkManager-wait-online` failed at every boot in router mode, so VPN servers started about 100 s late and the bot about 2 minutes late.
- Correct behaviour: Services do not wait for a network-online condition that cannot become true in the box's mode; the bot retries by itself (CHANGELOG 201, 214).
- Source: CHANGELOG.md@ad3758c:201; CHANGELOG.md@ad3758c:214; docs/TESTING.md@ad3758c:115
- Python fix commit: 71222cc, f2f57f6
- Package that proves it: P-10 (also S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0134: Update test run used the box's real lock and state

- Area: Updates
- What went wrong: The update's test run used the box's real measurement lock and DynDNS state, so an update was refused at random while a scheduled test ran.
- Correct behaviour: An update's tests get their own lock and state files; nothing they use is shared with the running box.
- Source: CHANGELOG.md@ad3758c:206; docs/TESTING.md@ad3758c:118
- Python fix commit: e59463f
- Package that proves it: P-02 (also F-06)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0135: Update reported as installed when it had only started

- Area: Updates
- What went wrong: `/heal apply` said "Update installed, the verification tests passed" the moment the update had only started.
- Correct behaviour: Status reports what happened so far: started, which version, where to follow it; "installed" only after the health check passed (porting plan 7D 624).
- Source: CHANGELOG.md@ad3758c:207
- Python fix commit: e59463f
- Package that proves it: P-02 (also S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0136: Measurement cycle failed while counting its data

- Area: Data caps
- What went wrong: A cycle could fail at its end while counting the data it used (the daily usage was written outside the guard).
- Correct behaviour: Counting data never fails a cycle; a counting error is logged and the measured result is kept. When the usage cannot be counted, scheduled tests on that metered link are treated as possibly over the cap (they wait), and the counting error is shown in the status (derived: roadmap 587, errors fail closed). Owner, 2026-11-11: on a metered link without a cap decision speed tests wait until the user sets a cap or says unlimited, scheduled ones and those started by a person alike (the person is told why it waits); ping keeps running (KF-0294).
- Source: CHANGELOG.md@ad3758c:208
- Python fix commit: 70c4b91
- Package that proves it: M-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0137: Wired route priority set by the unprivileged web server

- Area: Routes
- What went wrong: The wired route priority was set only when the web server started, as a user that may not change network profiles, and with a route of its own outside the network manager, which stayed after the cable moved.
- Correct behaviour: The network daemon sets the priority in the profile at each mode change and periodically outside router mode, and removes routes it added.
- Source: CHANGELOG.md@ad3758c:209; CHANGELOG.md@ad3758c:233; docs/TESTING.md@ad3758c:117; private log review 2026-11-09, still open
- Python fix commit: 5bb420e, b605844
- Package that proves it: N-01 (also N-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0138: DNS fallback acted on short failures and restored the wrong file

- Area: DNS resolver
- What went wrong: The DNS fallback replaced the resolvers with public ones on any short failure, even without IP connectivity, and nothing undid it; when the system rewrote `resolv.conf` the stale marker could later restore another network's file; an empty `resolv.conf` (the local resolver) was taken as broken, so DNS flapped every 16 minutes.
- Correct behaviour: The fallback acts only when IP works but DNS does not (three tries), keeps a fresh copy, drops its marker when the system replaced the file, tests an empty list like any other, and puts the normal DNS back after 15 minutes when it works.
- Source: CHANGELOG.md@ad3758c:254; CHANGELOG.md@ad3758c:210; CHANGELOG.md@ad3758c:224; private log review 2026-11-09, item 4
- Python fix commit: 51dfd1b, 5bb420e, f838b5d
- Package that proves it: N-01 (also N-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0139: IPv6 DNS servers pushed out IPv4 ones

- Area: DNS resolver
- What went wrong: With IPv6 DNS servers from the router but no IPv6 route, the three nameserver slots held only IPv6 servers and every lookup failed.
- Correct behaviour: IPv4 DNS first on wired and Wi-Fi profiles; a resolver is used only on an address family with a route.
- Source: CHANGELOG.md@ad3758c:255; private log review 2026-11-09, item 5
- Python fix commit: 51dfd1b
- Package that proves it: N-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0140: Logs filled the RAM log disk

- Area: Logging
- What went wrong: The developer log rotated at 10 MB with 5 copies on a 48 MB RAM disk; every service then failed to write its log and an install stopped halfway; the storage guard checked only the data disk.
- Correct behaviour: Logs never fill a disk: rotation fits the file system the log is on, the storage guard watches a separate log disk too (porting plan 7D 617).
- Source: CHANGELOG.md@ad3758c:218; CHANGELOG.md@ad3758c:212; CHANGELOG.md@ad3758c:215; docs/TESTING.md@ad3758c:111; docs/TESTING.md@ad3758c:114; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:617
- Python fix commit: 558996c, 22eb38f, f2f57f6
- Package that proves it: C-03 (also C-08)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0141: Installer stopped halfway on a full log disk

- Area: Installer
- What went wrong: `setup.sh` logged to a nearly full `/var/log` and stopped halfway, so the web server kept running old code.
- Correct behaviour: The installer makes room first and logs to persistent storage; a failed install changes nothing (derived: porting plan 7D 623).
- Source: CHANGELOG.md@ad3758c:211; docs/TESTING.md@ad3758c:116
- Python fix commit: 22eb38f
- Package that proves it: P-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0142: Documentation described commands, versions and features wrongly

- Area: Documentation
- What went wrong: The Telegram command list showed 5 languages and old `/mode` values; deploy steps used a short commit id that the installer drops, so the update check did not know the installed version; seven Web UI texts described older behaviour; the READMEs described features differently from how they worked (the bridge needing an IP address, the self-update needing a Git checkout).
- Correct behaviour: The manual and the command reference are generated from the commands themselves and checked by the build, and describe features as they work (roadmap 6B.7; CHANGELOG 132). A documented feature that does not exist is KF-0288.
- Source: CHANGELOG.md@ad3758c:213; CHANGELOG.md@ad3758c:296; CHANGELOG.md@ad3758c:132
- Python fix commit: 22eb38f
- Package that proves it: P-12 (also C-06)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0143: Log purge left the SD-card logs while saying they were deleted

- Area: System
- What went wrong: `/purge logs` left the warnings kept on the SD card while saying the logs were deleted.
- Correct behaviour: A purge clears every copy it names and reports what it cleared.
- Source: CHANGELOG.md@ad3758c:216; docs/TESTING.md@ad3758c:113
- Python fix commit: f2f57f6
- Package that proves it: C-08 (also S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0144: `/kill` sent another signal than the one named

- Area: Telegram bot
- What went wrong: `/kill <pid> <signal>` sent SIGTERM for a signal it does not allow while the reply named the signal asked for.
- Correct behaviour: A signal that is not allowed is refused.
- Source: CHANGELOG.md@ad3758c:217; docs/TESTING.md@ad3758c:112
- Python fix commit: f2f57f6
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0145: DynDNS state split between processes

- Area: DynDNS
- What went wrong: `/dyndns off` was saved but the web server's updater kept updating; `/dyndns on` started a second updater in the bot; the bot showed its own empty state; `/dyndns update` used the wrong provider.
- Correct behaviour: One updater follows the saved settings within 30 s and shares its state; the bot only saves the setting; updates use the configured provider.
- Source: CHANGELOG.md@ad3758c:219; docs/TESTING.md@ad3758c:110
- Python fix commit: b42e08b, 0d63d43
- Package that proves it: N-16 (also S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0146: Purge in one transaction filled the RAM disk

- Area: Database
- What went wrong: `/purge measurements` deleted everything in one transaction; its write-ahead log grew about as big as the database and the RAM disk filled ("database or disk is full").
- Correct behaviour: Big tables are deleted in slices with a checkpoint after each, so the log stays small.
- Source: CHANGELOG.md@ad3758c:220; docs/TESTING.md@ad3758c:110
- Python fix commit: 5625c4f
- Package that proves it: C-08 (also C-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0147: Update check offered older code

- Area: Updates
- What went wrong: The update check offered an update whenever the installed commit differed from the remote, also when the box ran newer code; applying it would have put older code back.
- Correct behaviour: An update is offered only when the remote is newer; without one, applying changes nothing (native: the signed version decides, roadmap 6B.3).
- Source: CHANGELOG.md@ad3758c:221; docs/TESTING.md@ad3758c:110
- Python fix commit: e7a4f04
- Package that proves it: P-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0148: `/wifi connect` used the access point's radio

- Area: Wi-Fi
- What went wrong: In router mode `/wifi connect` tried to join a network with the radio that is the access point.
- Correct behaviour: It refuses and says why (switch to standalone or turn the access point off).
- Source: CHANGELOG.md@ad3758c:222
- Python fix commit: e7a4f04
- Package that proves it: N-07 (also S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0149: Router rebuilt and PPPoE redialled when not needed

- Area: Network modes
- What went wrong: A mode setting that led to the same mode rebuilt the router and redialled PPPoE; automatic mode dialled twice at every start and reload; an access point setting rebuilt the whole router (about 20-40 s without internet each).
- Correct behaviour: The running mode is kept when the new setting leads to it; the auto decision is applied before the first setup; an access point change restarts only the access point.
- Source: CHANGELOG.md@ad3758c:223; CHANGELOG.md@ad3758c:227; CHANGELOG.md@ad3758c:228; docs/TESTING.md@ad3758c:106; docs/TESTING.md@ad3758c:107; docs/TESTING.md@ad3758c:108
- Python fix commit: f838b5d, 656a20a, fe24f3f
- Package that proves it: N-03 (also N-04, N-07)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0150: Outages left open by a stopped probe

- Area: Outages
- What went wrong: Outages left open by a probe that stopped counted on, overlapping ones added up (0 % availability), each restart opened a second row for the same outage, and the bot's outage alert stayed "firing" for two days.
- Correct behaviour: An open outage ends at the next outage or the first answered ping; overlaps count once; after a restart the newest carries on if the line is still down; the alert resolves.
- Source: CHANGELOG.md@ad3758c:225; CHANGELOG.md@ad3758c:229; private log review 2026-11-10 (night)
- Python fix commit: 70b3793, f838b5d
- Package that proves it: M-02 (also R-01, S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0151: iPerf3 rest warning flooded the log

- Area: iPerf3
- What went wrong: The "left out for up to 6 h" warning came every minute (about 17 an hour after a first fix).
- Correct behaviour: Logged once per rest period.
- Source: CHANGELOG.md@ad3758c:230; CHANGELOG.md@ad3758c:259; private log review 2026-11-09, item 7
- Python fix commit: 51dfd1b, 504bcd5
- Package that proves it: M-09
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0152: Ping probe bound once to a link without an address

- Area: Ping probe
- What went wrong: The probe bound its socket once at start; under PPPoE that link had no address, so every ping was lost for hours and each restart opened an outage that never closed, while the line worked.
- Correct behaviour: The probe follows the line (the router's WAN in router mode, never a link without an address) and rebinds within 30 s; a gateway target is switched off when there is no gateway.
- Source: CHANGELOG.md@ad3758c:231; docs/TESTING.md@ad3758c:103; private log review 2026-11-10 (night)
- Python fix commit: 752b22d
- Package that proves it: M-01 (also N-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0153: Network manager fought router mode for its ports

- Area: Network modes
- What went wrong: The network manager ran DHCP on the PPPoE port every 45 s and unbridged a LAN port.
- Correct behaviour: Ports used by router mode are released from the network manager while it runs and given back afterwards.
- Source: CHANGELOG.md@ad3758c:232; CHANGELOG.md@ad3758c:285; docs/TESTING.md@ad3758c:104
- Python fix commit: 78ecc2e
- Package that proves it: N-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0154: DynDNS sent the wrong address or waited too long

- Area: DynDNS
- What went wrong: With IPv6 on the box the IPv4 address was asked from dual-stack services and the IPv6 one sent as IPv4; after a WAN change the name pointed at the old line for up to the whole interval; turned on without a token it said nothing.
- Correct behaviour: The IPv4 address is asked over IPv4 only and sent explicitly; the route is checked every 30 s and a change updates at once; a missing token is logged and shown.
- Source: CHANGELOG.md@ad3758c:234; CHANGELOG.md@ad3758c:244; CHANGELOG.md@ad3758c:262; docs/TESTING.md@ad3758c:91; private log review 2026-11-09 evening, items 6 and 18
- Python fix commit: 3298900, 093dfa2, 8a08ca2
- Package that proves it: N-16
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0155: Kernel tuning re-applied on every settings save

- Area: System
- What went wrong: Every settings save re-applied the kernel tuning (about 8 s); dnsmasq's own pid file logged a capability warning at every start.
- Correct behaviour: Tuning is applied only when its file or the set of interfaces changed; no needless privileged operations.
- Source: CHANGELOG.md@ad3758c:235; private log review 2026-11-09 evening, item 17
- Python fix commit: 26ce021
- Package that proves it: N-03 (also C-09)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0156: Installer changed owners of runtime files

- Area: Installer
- What went wrong: The installer gave the pppd options file in the runtime folder to the service user, so the network daemon could not rewrite it and router mode on PPPoE did not start.
- Correct behaviour: Generated files are written new and moved into place; the installer does not change owners of runtime files.
- Source: CHANGELOG.md@ad3758c:236; private log review 2026-11-09 evening, item 10
- Python fix commit: da66ffb
- Package that proves it: N-05 (also P-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0157: Refused PPPoE login sent again every 8 s for ever

- Area: PPPoE
- What went wrong: A wrong login was sent every 8 s for ever, after which the ISP stopped answering for a while.
- Correct behaviour: Give up after 3 failures; dial again after 10, 20, then 30 minutes for a refused login (30 s for a line problem); the reason shows in the status.
- Source: CHANGELOG.md@ad3758c:237; docs/TESTING.md@ad3758c:331; private log review 2026-11-09 evening, item 2
- Python fix commit: 093dfa2
- Package that proves it: N-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0158: PPPoE took no IPv6 because forwarding was on

- Area: IPv6
- What went wrong: ppp0 took no IPv6 from the ISP because forwarding was on.
- Correct behaviour: The WAN accepts router adverts from the moment the link is created, also with forwarding on.
- Source: CHANGELOG.md@ad3758c:238; private log review 2026-11-09 evening, item 8
- Python fix commit: 00c1e50
- Package that proves it: N-11 (also N-05)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0159: Scheduled tests ran over the USB hotspot

- Area: Measurement link
- What went wrong: Scheduled tests ran over the USB hotspot (mobile data) while the network daemon restarted or the router's WAN was down (19 tests in an hour).
- Correct behaviour: A USB hotspot carries a test only when it is the chosen measuring link, the router's WAN, or the only internet in standalone or bridge mode; no test while the router's WAN is down.
- Source: CHANGELOG.md@ad3758c:239; docs/TESTING.md@ad3758c:93; private log review 2026-11-09 evening, items 14 and 20
- Python fix commit: f859f51, 19d9ba1
- Package that proves it: M-04 (also N-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0160: Tests bound to a link with only a link-local address

- Area: Measurement link
- What went wrong: In router mode on PPPoE scheduled tests bound to the PPPoE carrier port, which has only a link-local address, and failed.
- Correct behaviour: Router mode tests measure the router's WAN; a link-local address never counts as a test link.
- Source: CHANGELOG.md@ad3758c:240; docs/TESTING.md@ad3758c:93; private log review 2026-11-09 evening, item 13
- Python fix commit: f3863f1
- Package that proves it: M-04 (also N-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0161: Engine error ended the whole cycle without a row

- Area: Scheduling
- What went wrong: An engine that raised an error ended the whole scheduled cycle without a row (24 tracebacks in an hour); RTR control requests retried over an unbound connection, which would have measured another link.
- Correct behaviour: An engine error is stored as a failed test of that engine and the others run; no retry leaves the asked link.
- Source: CHANGELOG.md@ad3758c:241; private log review 2026-11-09 evening, item 7
- Python fix commit: 00c1e50
- Package that proves it: M-04 (also M-08)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0162: Access point flags not from the radio

- Area: Wi-Fi access point
- What went wrong: Fixed 802.11ac flags (STBC) the radio does not have made the access point fall back to 802.11n at 20 MHz; 40 MHz was asked on 2.4 GHz where the radio does 20 MHz only, and 802.11ac on 2.4 GHz; 160 MHz was claimed wrongly.
- Correct behaviour: HT and VHT flags and the width come from what the radio reports for each band, with the right HT40 direction.
- Source: CHANGELOG.md@ad3758c:242; CHANGELOG.md@ad3758c:256; private log review 2026-11-09, item 6; private log review 2026-11-09 evening, item 12
- Python fix commit: b71f8f6, 51dfd1b
- Package that proves it: N-07
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0163: USB did not take over when Wi-Fi had a route but no internet

- Area: USB tethering
- What went wrong: USB tethering did not take over when the Wi-Fi network had a default route but no internet.
- Correct behaviour: After about 90 s without internet on every other route the box uses USB until another route works.
- Source: CHANGELOG.md@ad3758c:243; private log review 2026-11-09 evening, item 3
- Python fix commit: 093dfa2
- Package that proves it: N-10
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0164: UPnP started before the WAN had an address

- Area: UPnP
- What went wrong: UPnP started before ppp0 had its address and answered every request with "network failure"; a rule format it could not parse.
- Correct behaviour: It starts once the WAN has an address and restarts when the address changes (CHANGELOG 245); the configuration and rules NMB writes for the UPnP daemon are in the format the installed daemon version accepts and NMB itself parses back, checked by a test against that version (private log review 2026-11-09 evening, item 11; the test against the installed version derived: Quality Bar 6, roadmap 584, right by design, not by luck).
- Source: CHANGELOG.md@ad3758c:245; private log review 2026-11-09 evening, item 11
- Python fix commit: da66ffb; note: the UPnP feature itself came in 00c1e50
- Package that proves it: N-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0165: Self-tests checked the wrong things or changed the box

- Area: Self-test
- What went wrong: The real-line PPPoE self-test pinged before the route was up and put the ISP's DNS into the box's resolver; the IPv6 group looked at the wrong port in router mode; the router group did not read the WAN name; the scenario's client counter used an invalid rule and counted the simulated client's lease; a DHCP script flushed IPv6 so a check passed by luck; reports were never rotated; a weak password was not a failure with the access point off; the snapshot tool said it had mirrored to the boot partition when the hourly limit had skipped it.
- Correct behaviour: Self-tests measure the real effect, change nothing outside their throwaway objects, and their reports are bounded (TESTING 231); a tool reports only what it did, so a skipped step is reported as skipped (CHANGELOG 250).
- Source: CHANGELOG.md@ad3758c:246; CHANGELOG.md@ad3758c:247; CHANGELOG.md@ad3758c:272; CHANGELOG.md@ad3758c:234; CHANGELOG.md@ad3758c:250; CHANGELOG.md@ad3758c:261; CHANGELOG.md@ad3758c:277; private log review 2026-11-09 evening, items 4, 16, 19
- Python fix commit: a7fddcc, 5509ab4, 386782a, 772d571, beb53cd
- Package that proves it: P-07
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0166: User-space PPPoE capped the line; no MSS clamping

- Area: PPPoE
- What went wrong: The user-space `pppoe` program topped out around 300-400 Mbit/s on a Pi 4; clients behind PPPoE hung on some sites.
- Correct behaviour: Kernel PPPoE (with the saved service name); router mode clamps the TCP MSS to the path MTU.
- Source: CHANGELOG.md@ad3758c:248
- Python fix commit: 23c98dd
- Package that proves it: N-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0167: Automatic mode stayed standalone on Wi-Fi or USB internet

- Area: Automatic mode
- What went wrong: Wi-Fi or USB internet kept the box in standalone when its cable was on the ONT; PPPoE discovery was not throttled.
- Correct behaviour: Only internet through the WAN port counts; a PPPoE server answering with a saved login switches to router mode at once; discovery at most every 30 s (roadmap 4.28).
- Source: CHANGELOG.md@ad3758c:249
- Python fix commit: 23c98dd
- Package that proves it: N-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0168: Bot hammered the Telegram API on errors

- Area: Telegram bot
- What went wrong: On an HTTP error (429, 502) or without network the bot asked again at once, about 8 times a second; a message refused with "retry after" was lost.
- Correct behaviour: It waits (Telegram's retry_after, or a growing pause) and sends such a message once more.
- Source: CHANGELOG.md@ad3758c:251; private log review 2026-11-09, item 1
- Python fix commit: 51dfd1b
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0169: Bot sandbox too narrow to save its settings

- Area: Sandboxing
- What went wrong: Only the live database was writable for the bot, so its DynDNS commands could not save, `/heal` could not keep its state, and Ookla started as a new client on every bot test.
- Correct behaviour: Each daemon's sandbox allows exactly the paths its functions write; a test runs each function under its profile (roadmap 6B.5).
- Source: CHANGELOG.md@ad3758c:252; private log review 2026-11-09, item 2
- Python fix commit: 51dfd1b
- Package that proves it: P-04 (also S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0170: Dead default route deletion failed while the log said removed

- Area: Routes
- What went wrong: `ip route del ... dead` failed, yet the log said "removed" every 5 seconds.
- Correct behaviour: Only the route's selector is passed and success is checked before it is logged (native: netlink, roadmap 4.2).
- Source: CHANGELOG.md@ad3758c:253; private log review 2026-11-09, item 3
- Python fix commit: 51dfd1b
- Package that proves it: N-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0171: RTR client lost its client ID on every test

- Area: RTR engine
- What went wrong: The native RTR client kept its ID in `$HOME`, which did not exist for the services.
- Correct behaviour: Persistent client state lives in the box's own data folder.
- Source: CHANGELOG.md@ad3758c:257; private log review 2026-11-09, item 8
- Python fix commit: 51dfd1b
- Package that proves it: M-08
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0172: Installer made the root-only secrets copy readable

- Area: Installer
- What went wrong: The installer made the root-only secrets copy readable by the service user.
- Correct behaviour: Root-only files stay root-only; an installer never widens a mode or owner recursively.
- Source: CHANGELOG.md@ad3758c:258; private log review 2026-11-09, item 9
- Python fix commit: 51dfd1b
- Package that proves it: P-11 (also C-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0173: Scheduler errors without a traceback

- Area: Logging
- What went wrong: Scheduler errors were logged without their traceback.
- Correct behaviour: Errors are logged with their context and stack (roadmap 745; porting plan 7D 615).
- Source: CHANGELOG.md@ad3758c:259; private log review 2026-11-09, item 10
- Python fix commit: 51dfd1b
- Package that proves it: C-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0174: Bot blocked during a speed test; showed other rows

- Area: Telegram bot
- What went wrong: A speed test blocked the bot for its whole duration; the result showed whatever row was newest; the start message promised 15-25 seconds.
- Correct behaviour: Tests run in their own task and a second one is refused politely; the result shows only this test's rows; no promised duration.
- Source: CHANGELOG.md@ad3758c:263
- Python fix commit: e0201ed
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0175: Measurement lock per process

- Area: Scheduling
- What went wrong: The lock was per process, so a Telegram test and a scheduled test could run at once and measure each other's traffic.
- Correct behaviour: One measurement at a time on the whole box: a scheduled test skips, a manual one waits up to 10 minutes.
- Source: CHANGELOG.md@ad3758c:264
- Python fix commit: e0201ed
- Package that proves it: M-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0176: PPPoE login on the pppd command line

- Area: PPPoE
- What went wrong: The PPPoE user name and password were on the pppd command line, readable with `ps`.
- Correct behaviour: They are in a private options file; no secret on any command line (roadmap 587).
- Source: CHANGELOG.md@ad3758c:265
- Python fix commit: 4885e96
- Package that proves it: N-05 (also C-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0177: Starting the access point flushed its own address

- Area: Wi-Fi access point
- What went wrong: Starting the access point flushed every address from the radio, its own included, so DHCP failed and devices got no address.
- Correct behaviour: Only a DHCP client's dynamic addresses and routes are removed; a self-test fails when an access point has no address or no DHCP.
- Source: CHANGELOG.md@ad3758c:266
- Python fix commit: 721f163
- Package that proves it: N-07
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0178: Failing iPerf3 server tried every minute; busy server added failed rows

- Area: iPerf3
- What went wrong: An extra server that failed 282 of 283 tests was tried every minute; a busy or unreachable server added failed rows to scheduled runs.
- Correct behaviour: A server whose last five tests failed is left out for 6 hours; a busy server adds no row.
- Source: CHANGELOG.md@ad3758c:269; CHANGELOG.md@ad3758c:304; docs/TESTING.md@ad3758c:340; CHANGELOG.md@ad3758c:121 (also: scheduled iPerf3 at most hourly by default, since public servers take one client at a time)
- Python fix commit: d6202e4, 34515a9
- Package that proves it: M-09
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0179: Router mode left LAN devices with old addresses

- Area: Network modes
- What went wrong: A PC on the LAN port kept an address from before and had no internet until its cable was replugged.
- Correct behaviour: After setting up, router mode drops the wired LAN ports' link for 4 seconds so devices ask again.
- Source: CHANGELOG.md@ad3758c:271; docs/TESTING.md@ad3758c:181
- Python fix commit: 772d571
- Package that proves it: N-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0180: Dead default routes beat working links

- Area: Routes
- What went wrong: The box's own static default route on a port whose cable moved stayed, beat Wi-Fi and USB, and the box lost its internet and remote access.
- Correct behaviour: Default routes through a port with no link or no address are removed; routes on a port without link are ignored; a self-test fails on such a route.
- Source: CHANGELOG.md@ad3758c:273; docs/TESTING.md@ad3758c:180; docs/TESTING.md@ad3758c:336
- Python fix commit: b39e954
- Package that proves it: N-01 (also N-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0181: Boot-partition copy replaced set secrets

- Area: Configuration
- What went wrong: After an update the editable boot copy counted as newer and its empty token replaced the real one; restoring from it wrote `__KEEP_CURRENT__` marks as values and brought back the factory Wi-Fi password.
- Correct behaviour: The boot copy is taken over only when it was really edited (CHANGELOG 275); an empty value or a keep mark (`__KEEP_CURRENT__`, which existing boot files keep using) never replaces a set secret, and a restore never writes a keep mark as a value or brings back a factory password; no secret is written to the boot partition (roadmap 682). How set secrets are kept and restored is designed by C-04 and P-03 under their security review.
- Source: CHANGELOG.md@ad3758c:275; CHANGELOG.md@ad3758c:274; CHANGELOG.md@ad3758c:332; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:682
- Python fix commit: beb53cd, b39e954, bbcf9b6
- Package that proves it: C-01 (also C-04, P-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0182: Network daemon loop went silent

- Area: Network daemon
- What went wrong: The network daemon's loop went silent for 8 minutes while a cable moved, and USB never took over.
- Correct behaviour: A loop part that takes longer than its limit is logged, and stacks can be dumped on request (porting plan 7D 615-618).
- Source: CHANGELOG.md@ad3758c:276
- Python fix commit: beb53cd
- Package that proves it: N-03 (also C-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0183: `--interface` option read but never used

- Area: Speed engines
- What went wrong: `nmb-speed --interface` was read but never used, so a test asked for USB ran on the default link.
- Correct behaviour: Every command option reaches the code it names; a test checks the traffic left over the asked link.
- Source: CHANGELOG.md@ad3758c:278; docs/TESTING.md@ad3758c:179
- Python fix commit: 924b3c9
- Package that proves it: M-04 (also S-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0184: Collector answered before reading the upload

- Area: Central collector
- What went wrong: A box with a wrong token and a large backlog got "Broken pipe" instead of the refusal, because the collector answered before reading the upload.
- Correct behaviour: The client gets the refusal reason (the server answers after the request, or the client reads the answer first).
- Source: CHANGELOG.md@ad3758c:279
- Python fix commit: aaf05cc
- Package that proves it: S-13 (also S-12)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0185: Router WAN not moved to USB with only a link-local address

- Area: USB tethering
- What went wrong: Router mode kept its WAN on a port that had only an IPv6 link-local address.
- Correct behaviour: A port with only a link-local address is not a WAN; router mode moves its WAN to USB.
- Source: CHANGELOG.md@ad3758c:281
- Python fix commit: 3cb3040
- Package that proves it: N-10
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0186: DHCP guard missed routers on bridged ports

- Area: DHCP guard
- What went wrong: The DISCOVER went through the IP stack with another interface's address as the source, and offers were read with an IPv4-only packet socket that gets a bridge port's frames too late.
- Correct behaviour: A complete frame from 0.0.0.0 and reading with `ETH_P_ALL`, plain and bridged.
- Source: CHANGELOG.md@ad3758c:282; docs/TESTING.md@ad3758c:181
- Python fix commit: c12a272
- Package that proves it: N-05 (also N-06, N-08)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0187: Router mode on USB lost the box's own internet

- Area: USB tethering
- What went wrong: The radio that became the access point kept its home Wi-Fi address and default route, and a stale route on a bridged LAN port counted as another internet, so the USB route was removed.
- Correct behaviour: The access point starts on a radio without client addresses; in router mode on USB routes through bridge ports or down links never count and default routes not through USB are removed.
- Source: CHANGELOG.md@ad3758c:283; CHANGELOG.md@ad3758c:284; docs/TESTING.md@ad3758c:181
- Python fix commit: c12a272, b5bca61
- Package that proves it: N-10 (also N-07)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0188: Generated config left by another process blocked a service

- Area: Wi-Fi access point
- What went wrong: The access point never started when the web server had left `hostapd.conf` in the runtime folder (the daemon could not change its mode).
- Correct behaviour: Generated configs replace the old file (new file, then rename).
- Source: CHANGELOG.md@ad3758c:286
- Python fix commit: deb268c
- Package that proves it: N-07
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0189: USB phone or hotspot used as a LAN port or as the wired line

- Area: USB tethering
- What went wrong: Router and bridge modes could use a USB phone, hotspot or modem as a LAN port and serve DHCP to it; one named like an Ethernet port was treated as the wired line, so tests over mobile data were stored as the wired line.
- Correct behaviour: Such interfaces are recognised by their driver and are never the wired line or a LAN port; USB Ethernet dongles stay wired (roadmap 4.26).
- Source: CHANGELOG.md@ad3758c:287; CHANGELOG.md@ad3758c:288; docs/TESTING.md@ad3758c:179
- Python fix commit: 53274f3, 68458cf
- Package that proves it: N-10 (also N-17)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0190: A save undid another service's save

- Area: Configuration
- What went wrong: Each service wrote back its whole in-memory copy of the settings, so an entered Telegram token could disappear.
- Correct behaviour: A save writes only what that service changed (a three-way merge); release checklist row 12.
- Source: CHANGELOG.md@ad3758c:290; docs/TESTING.md@ad3758c:178
- Python fix commit: b1e64b9
- Package that proves it: C-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0191: Feature switched on without its token stayed silent

- Area: Configuration
- What went wrong: Telegram enabled without a bot token gave no warning (and DynDNS without a token, see the DynDNS entry).
- Correct behaviour: A feature that is on but cannot work says why, in the status and the log.
- Source: CHANGELOG.md@ad3758c:291
- Python fix commit: b1e64b9
- Package that proves it: S-11 (also C-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0192: Service kept running old code after an upgrade

- Area: Installer
- What went wrong: The installer never restarted the Telegram bot, so it ran week-old code ("cannot import name ...").
- Correct behaviour: After an install every service runs the new code (each started after the install, TESTING 259 and checklist 21).
- Source: CHANGELOG.md@ad3758c:293; docs/TESTING.md@ad3758c:167; docs/TESTING.md@ad3758c:184; docs/TESTING.md@ad3758c:259
- Python fix commit: 42a0fd9
- Package that proves it: P-11 (also P-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0193: Error without text gave an empty message

- Area: Telegram bot
- What went wrong: A bot error with no text gave an empty message.
- Correct behaviour: An error always names its kind.
- Source: CHANGELOG.md@ad3758c:294
- Python fix commit: 6fb0c3f
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0194: Home card showed a failed run as a result

- Area: Web UI
- What went wrong: The home card showed a failed run ("-- Mbps", "Upload only", "A+").
- Correct behaviour: It shows the last real result, says when a newer attempt failed, when the next Ookla test can run, and when the result is from an earlier chosen server.
- Source: CHANGELOG.md@ad3758c:305
- Python fix commit: 34515a9
- Package that proves it: U-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0195: Test bound to Wi-Fi left over Ethernet

- Area: Measurement link
- What went wrong: A test bound to Wi-Fi left over Ethernet when the Wi-Fi was joined outside NMB and both links were on the same subnet.
- Correct behaviour: A test bound to a link leaves over that link (source rule kept also for a link joined outside NMB), checked by byte counters (release checklist 8).
- Source: CHANGELOG.md@ad3758c:300; docs/TESTING.md@ad3758c:291
- Python fix commit: 09db7e6
- Package that proves it: N-08 (also M-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0196: Status API failed after a failed test

- Area: Web server
- What went wrong: `/api/status` failed once the history held a failed speed test.
- Correct behaviour: A failed test is a normal row; status never fails on stored data.
- Source: CHANGELOG.md@ad3758c:309; docs/TESTING.md@ad3758c:172
- Python fix commit: 8aa0efc
- Package that proves it: S-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0197: Reinstall left removed files; RAM data not saved first

- Area: Installer
- What went wrong: Reinstalling copied over the program files, so removed files stayed, and did not save the measurements still in RAM.
- Correct behaviour: An install saves the RAM data first and replaces the program files as a whole (native: A/B or a second copy, porting plan 7D 623).
- Source: CHANGELOG.md@ad3758c:310
- Python fix commit: 8aa0efc
- Package that proves it: P-11 (also P-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0198: Restore lost data and left files with wrong modes

- Area: Backup and restore
- What went wrong: Restoring a backup could lose the measurements (the old `-wal` journal was replayed onto the new file); a restored config was overwritten by the older boot copy and readable only by root; an uploaded backup was readable by every user.
- Correct behaviour: Databases are restored through SQLite's backup API; the restored config wins and is readable by the services only; uploaded files are private; a truncated upload changes nothing (TESTING 170).
- Source: CHANGELOG.md@ad3758c:311; CHANGELOG.md@ad3758c:312; docs/TESTING.md@ad3758c:170
- Python fix commit: 8ae01b0, 8aa0efc
- Package that proves it: P-03 (also C-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0199: Helper upgrade never upgraded; helper path not found

- Area: Privileged helper
- What went wrong: `nmb-helper upgrade` ran only `apt-get update` (a second `exec` never ran) and ran `apt-get` on Alpine; the update check called the helper without its path, which fails where `sudo` does not search `/usr/local/bin`.
- Correct behaviour: Each helper action does what it says, for the package manager of the system it runs on, called by full path.
- Source: CHANGELOG.md@ad3758c:313; CHANGELOG.md@ad3758c:314; docs/TESTING.md@ad3758c:42
- Python fix commit: 12d5a3e
- Package that proves it: P-06 (also P-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0200: `--config FILE` saved to the default file

- Area: Configuration
- What went wrong: `nmb-web --config FILE` saved settings to `/etc/nmb/nmb.conf` instead of FILE.
- Correct behaviour: Every program reads and writes the configuration path it was given.
- Source: CHANGELOG.md@ad3758c:315
- Python fix commit: 12d5a3e
- Package that proves it: C-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0201: Installer failures on other systems

- Area: Installer
- What went wrong: The installer failed halfway on Debian derivatives without systemd, did not install a polkit rule older polkit reads, skipped EPEL on RHEL rebuilds, did not detect interface names other than `eth0`, failed on Alpine reinstalls, port 80 and the RTR build, and printed raw colour codes.
- Correct behaviour: Packages and images for each supported base are built and tested there (roadmap 6A; work plan P-11); an unsupported system is refused with a clear message.
- Source: CHANGELOG.md@ad3758c:316; CHANGELOG.md@ad3758c:317; CHANGELOG.md@ad3758c:333; CHANGELOG.md@ad3758c:334; CHANGELOG.md@ad3758c:335; CHANGELOG.md@ad3758c:329; docs/TESTING.md@ad3758c:34; docs/TESTING.md@ad3758c:36; docs/TESTING.md@ad3758c:39; docs/TESTING.md@ad3758c:42
- Python fix commit: a819282, 12d5a3e, ded1b58
- Package that proves it: P-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0202: OpenVPN server certificate without TLS server usage

- Area: VPN
- What went wrong: The server certificate had no TLS server usage, so every client was rejected.
- Correct behaviour: Server certificates carry TLS server usage; an old one is reissued from the same CA so issued profiles keep working.
- Source: CHANGELOG.md@ad3758c:318; docs/TESTING.md@ad3758c:149
- Python fix commit: 8ae01b0
- Package that proves it: N-14
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0203: HTTPS proxy slow on broken IPv6, port kept, loopback reachable

- Area: Proxy
- What went wrong: The proxy took 20 s per request on networks with a broken IPv6 route, kept its port after stop, and forwarded to the box's own loopback and link-local addresses.
- Correct behaviour: IPv6 and IPv4 tried side by side; stop frees the port; the box's own loopback and link-local addresses are refused.
- Source: CHANGELOG.md@ad3758c:319; docs/TESTING.md@ad3758c:150
- Python fix commit: 8ae01b0
- Package that proves it: N-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0204: Guest Wi-Fi client could not be controlled and reported wrong state

- Area: Guest Wi-Fi
- What went wrong: The service user could not control the network manager, reconnecting failed on a stale profile, errors hid the manager's message, and a disconnected client was reported as connected.
- Correct behaviour: Real state and the manager's own error are reported; a stale profile does not block reconnecting.
- Source: CHANGELOG.md@ad3758c:320; docs/TESTING.md@ad3758c:139; docs/TESTING.md@ad3758c:140
- Python fix commit: 8ae01b0
- Package that proves it: N-08
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0205: 5 GHz never started without the regulatory database

- Area: Wi-Fi access point
- What went wrong: 5 GHz Wi-Fi never started on systems without the regulatory database.
- Correct behaviour: The access point falls back to 2.4 GHz rather than none; the image carries the database.
- Source: CHANGELOG.md@ad3758c:321; docs/TESTING.md@ad3758c:127; docs/TESTING.md@ad3758c:128
- Python fix commit: 8ae01b0
- Package that proves it: N-07 (also P-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0206: DHCP renewal after a restart not answered at once

- Area: DHCP server
- What went wrong: After a box restart a client renewing an old lease was not answered at once.
- Correct behaviour: Router mode is authoritative on its LAN and answers a renewal at once.
- Source: CHANGELOG.md@ad3758c:322; docs/TESTING.md@ad3758c:70
- Python fix commit: not recorded in the source
- Package that proves it: N-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0207: Machine-readable output mixed with text

- Area: Command line
- What went wrong: `nmb-heal update --json` and `nmb-heal fix --json` did not print pure JSON.
- Correct behaviour: Machine-readable output is only that format; text goes elsewhere (roadmap 668).
- Source: CHANGELOG.md@ad3758c:323; docs/TESTING.md@ad3758c:159
- Python fix commit: 8ae01b0
- Package that proves it: P-01 (also S-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0208: Wrong version number reported

- Area: Version
- What went wrong: The software reported 2.4.0 instead of 2.5.0; the dossier showed a fixed version; the bundle said "unknown".
- Correct behaviour: The version comes from one installed version record everywhere (CHANGELOG 23).
- Source: CHANGELOG.md@ad3758c:324; CHANGELOG.md@ad3758c:23; CHANGELOG.md@ad3758c:151
- Python fix commit: not recorded in the source
- Package that proves it: C-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0209: Kernel tuning from the unprivileged web server

- Area: System
- What went wrong: The unprivileged web server ran the kernel tuning and filled its log with permission errors.
- Correct behaviour: Privileged operations run in the component that holds the privilege; others ask it.
- Source: CHANGELOG.md@ad3758c:328
- Python fix commit: ded1b58
- Package that proves it: N-03 (also P-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0210: Wireless router mode broke wired DHCP

- Area: Network modes
- What went wrong: A Wi-Fi failure stopped wired DHCP; hostapd did not step down when it rejected a setting; WPA3 lacked its cipher setting; network settings needed a restart once the web server was no longer root.
- Correct behaviour: Wi-Fi and LAN share a bridge; a Wi-Fi failure never stops wired service; the access point steps down to an older standard; settings take effect through the daemon that owns them.
- Source: CHANGELOG.md@ad3758c:331; CHANGELOG.md@ad3758c:338; CHANGELOG.md@ad3758c:339
- Python fix commit: 620e533
- Package that proves it: N-03 (also N-07)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0211: PPPoE did not redial; dnsmasq clashed with the DNS filter

- Area: PPPoE and DHCP
- What went wrong: PPPoE stayed offline after the line dropped; dnsmasq conflicted with the built-in DNS filter, which could leave clients without DHCP.
- Correct behaviour: PPPoE redials after a drop; DHCP and DNS do not fight over port 53.
- Source: CHANGELOG.md@ad3758c:336; CHANGELOG.md@ad3758c:337
- Python fix commit: bbcf9b6
- Package that proves it: N-05 (also N-12)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0212: Diagnostic snapshot flooded the journal and the boot drive

- Area: Logging
- What went wrong: The diagnostic snapshot copied thousands of lines into the system journal and was written to the boot drive at every run.
- Correct behaviour: Flash-safe logging: bounded volume, rare writes (porting plan 7D 617); the boot drive at most once an hour.
- Source: CHANGELOG.md@ad3758c:340
- Python fix commit: b4f472f
- Package that proves it: R-05 (also C-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0213: Language changed without login

- Area: Web server
- What went wrong: The language setting could be changed without logging in.
- Correct behaviour: Every state-changing request needs authentication (roadmap 587).
- Source: CHANGELOG.md@ad3758c:351
- Python fix commit: b4f472f
- Package that proves it: S-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0214: DNS filter control socket open to others

- Area: DNS ad-blocker
- What went wrong: The DNS filter's control socket was not limited to the box's own services.
- Correct behaviour: Control sockets are reachable only by the components that need them.
- Source: CHANGELOG.md@ad3758c:352
- Python fix commit: b4f472f
- Package that proves it: N-12
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0215: Control characters in API values reached service configs

- Area: Web server
- What went wrong: API values with control characters could add lines to service config files.
- Correct behaviour: Values with control characters are refused at the API (roadmap 587).
- Source: CHANGELOG.md@ad3758c:353
- Python fix commit: b4f472f
- Package that proves it: S-01 (also C-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0216: Update references read as options

- Area: Updates
- What went wrong: Update references that git could read as options were accepted.
- Correct behaviour: Version references are validated against a strict format; arguments are passed as a list (roadmap 587).
- Source: CHANGELOG.md@ad3758c:354
- Python fix commit: b4f472f
- Package that proves it: P-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0217: Tor service of the OONI package left on

- Area: OONI
- What went wrong: The Tor service that comes with the OONI package was enabled though NMB does not use it.
- Correct behaviour: Nothing runs that the box does not use (roadmap 589).
- Source: CHANGELOG.md@ad3758c:359
- Python fix commit: ea65ff7
- Package that proves it: M-13 (also P-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0218: PPPoE killed before it could hang up

- Area: PPPoE
- What went wrong: On a network daemon stop the stop timeout killed pppd before it sent PADT; the ISP kept the old session, refused the next login and freed it only after about 6 minutes.
- Correct behaviour: On stop the PPPoE session is hung up first (PADT sent), with time for it before the daemon is stopped.
- Source: docs/TESTING.md@ad3758c:94; docs/TESTING.md@ad3758c:396; private log review 2026-11-09 evening, item 15
- Python fix commit: f859f51
- Package that proves it: N-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0219: One bad line in the settings stops every service

- Area: Configuration
- What went wrong: One bad line in the settings file stops every service.
- Correct behaviour: Settings still load: a bad line or value is logged and replaced by its default, the rest of the file is used, and a save writes a file that parses (derived: roadmap 629).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:629; private audit 2026-11-10, reports item 7
- Python fix commit: ffceb45; note: one cause only
- Package that proves it: C-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: partly fixed

### KF-0220: Traffic gate treats unreadable counters as idle

- Area: Traffic gate
- What went wrong: When `/proc/net/dev` cannot be read the gate gets no counters and computes zero traffic, so it reports idle; a missing interface makes it sum all interfaces (`nmb/probe/traffic_gate.py@ad3758c:121-144`; `nmb/probe/traffic_gate.py@ad3758c:168-175`).
- Correct behaviour: Counters that cannot be read are treated as busy (roadmap 653); a missing interface is an error, never a sum of others.
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:653; private audit 2026-11-10, engines item 15; nmb/probe/traffic_gate.py@ad3758c:121-175
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: source analysis confirms

### KF-0221: Cron parser widens invalid fields to every minute

- Area: Scheduling
- What went wrong: An invalid field, `*/0` or `A-B/N` becomes every value (`nmb/speed/scheduler.py@ad3758c:61`), so a mistyped schedule runs every minute; day-of-month and day-of-week are ANDed.
- Correct behaviour: The API refuses an invalid schedule with a message and stores nothing; an invalid schedule read from the settings file is logged and replaced by its default (Phase 0, roadmap 629), never widened to every value; day-of-month and day-of-week follow cron's OR rule (derived: Quality Bar 6, roadmap 581, a wrong result is not carried over; the cron documentation, which states the OR rule, is the outside source the M-04 analyst reads and records).
- Source: private audit 2026-11-10, engines item 7; nmb/speed/scheduler.py@ad3758c:23-61
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: source analysis confirms the widening

### KF-0222: Ping round-trip times can be wrong; replies not the probe's own can count

- Area: Ping probe
- What went wrong: Round-trip times can be wrong, and replies that are not the probe's own can be counted.
- Correct behaviour: RTT from kernel or monotonic timestamps, never the wall clock (Quality Bar 6, roadmap 583); a reply is counted only when it answers one of the probe's own requests, and at most once, so a late reply is not a second sample (derived: Quality Bar 6, roadmap 583).
- Source: private audit 2026-11-10, engines item 11; source analysis
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0223: First-hop discovery blocks the probe loop

- Area: Ping probe
- What went wrong: Discovering the ISP's first hop blocks the probe loop for up to about 8 s and traces only the default route.
- Correct behaviour: Discovery runs beside the probe and never delays a probe tick; it traces the link the probe measures (derived: roadmap 4.1, 124-125, one event loop with non-blocking requests; roadmap 4.29, 450, the link the measurement left through).
- Source: private audit 2026-11-10, engines item 20
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0224: Partial or failed engine runs stored as successful

- Area: Speed engines
- What went wrong: Runs that ended early or failed in part were stored with status ok for several engines.
- Correct behaviour: A run is ok only when every direction it claims was measured; otherwise it is partial or failed with its reason (derived: QB6 581).
- Source: private audit 2026-11-10, engines item 8; CHANGELOG.md@ad3758c:145
- Python fix commit: 428b608; note: RTR only
- Package that proves it: M-04 (also M-06, M-07, M-08, M-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: fixed for RTR

### KF-0225: Failed Ookla or RTR run silently starts a second full test

- Area: Speed engines
- What went wrong: A failed Ookla or RTR run starts a second full test without saying so, with a misleading error.
- Correct behaviour: No hidden retry: a retry is visible in the row and counts against the data cap and Ookla's hourly limit (derived: roadmap 690).
- Source: private audit 2026-11-10, engines item 10
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-10 (also M-08, M-05)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0226: Ookla pause triggered by output that is not a rate limit

- Area: Ookla engine
- What went wrong: The Ookla pause can be triggered by output that is not Ookla's rate-limit answer.
- Correct behaviour: Only Ookla's own rate-limit answer pauses Ookla (derived: roadmap 690).
- Source: private audit 2026-11-10, engines item 26 (marked as suspicion); source analysis of `nmb/speed/anti_ban.py` at ad3758c
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-05 (also M-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: confirmed by source analysis

### KF-0227: iPerf3 rest bypassed with a single server

- Area: iPerf3
- What went wrong: With one configured server the rest rule is bypassed (the list falls back to its first entry).
- Correct behaviour: A rested server is not tested until its rest ends, also when it is the only one (CHANGELOG 269).
- Source: private audit 2026-11-10, engines item 12
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-09
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0228: iPerf3 upload from the sender side; UDP result is the offered rate

- Area: iPerf3
- What went wrong: The upload figure comes from the sender side, and a UDP test reports the offered 1 Mbit/s as its result.
- Correct behaviour: Throughput from the receiver side; a UDP result is the received rate with its loss (derived: Quality Bar 6, roadmap 581, a wrong result is not carried over, and roadmap 4.8, line 213, the native client measures throughput, jitter and packet loss; the iPerf3 documentation, which says which side's figures are the measured ones, is the outside source the M-09 analyst reads and records).
- Source: private audit 2026-11-10, engines item 13
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-09
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0229: Web UI iPerf3 test bypasses lock, gate, cap and binding

- Area: iPerf3
- What went wrong: An iPerf3 test started from the Web UI does not take the measurement lock, the traffic gate, the data cap or the link binding.
- Correct behaviour: Every test, from any surface, goes through the same runner: one at a time, the gate where it applies, the cap, the chosen link (roadmap 741, layers).
- Source: private audit 2026-11-10, engines item 5
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-04 (also M-09, S-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0230: Data accounting gaps

- Area: Data caps
- What went wrong: A missing interface counts 0 and then adds the lifetime counter; a counter reset gives a negative amount; the cap is checked only before a cycle; tests from the Web UI, the port check and traceroute are not counted.
- Correct behaviour: Usage is the sum of non-negative deltas of the measured interface; a reset starts a new delta; the cap is checked between engines; every test counts (CHANGELOG 35: counted from the data the tests themselves move on the measured interface, and tests started by hand count; the deltas, the reset and the check between engines derived: Quality Bar 6, roadmap 581, no wrong result is carried over). Owner, 2026-11-11: there is no default cap number; on a metered link without a cap decision speed tests wait until the user sets a cap or says unlimited, scheduled ones and those started by a person alike (the person is told why it waits); ping keeps running (KF-0294).
- Source: private audit 2026-11-10, engines item 14
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0231: Engines left an unbound or late-bound socket

- Area: Measurement link
- What went wrong: Some engines bind to the device only after connecting, or run unbound when the chosen link has no address, while the row names the asked interface.
- Correct behaviour: Bind before connect; without an address the engine leaves through the default route and the row describes that link (roadmap 452).
- Source: private audit 2026-11-10, engines item 4; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:452; CHANGELOG.md@ad3758c:241; CHANGELOG.md@ad3758c:267
- Python fix commit: d6202e4, 093dfa2, f3863f1; note: parts
- Package that proves it: M-04 (also M-06, M-07, M-08)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: partly fixed

### KF-0232: Port check reports throttling it cannot prove

- Area: Port check
- What went wrong: The port neutrality check can report "throttled" or "blocked" from a protocol mismatch, an invented baseline, timeouts and third-party targets.
- Correct behaviour: A port is reported only as what was measured against a server that speaks that protocol, with "unknown" when it cannot be told (derived: QB6 581).
- Source: private audit 2026-11-10, engines item 17
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0233: iPerf3 command-line tool reports success on failure

- Area: iPerf3
- What went wrong: `nmb-iperf` prints success on failure, its stop crashes, its status is always STOPPED and it reads wrong keys.
- Correct behaviour: Commands report what happened and exit with a matching code (roadmap 6B.7, exit codes).
- Source: private audit 2026-11-10, engines item 18
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-09 (also S-10, P-12)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0234: Scheduler start race

- Area: Scheduling
- What went wrong: Two scheduler starts can race.
- Correct behaviour: One scheduler per box (derived: roadmap 744, no hidden state).
- Source: private audit 2026-11-10, engines item 23
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0235: Service alert watches a unit that does not exist

- Area: Alerts
- What went wrong: The service alert includes a unit that does not exist (`nmb/bot/alerts.py@ad3758c:181`): it fires on every start, stays "firing", and real crashes are never alerted.
- Correct behaviour: Alerts watch the services that exist on that build; an alert resolves when its cause ends (derived: Quality Bar 6, roadmap 581: no wrong result or confusing message is carried over).
- Source: private audit 2026-11-10, bot item H6; nmb/bot/alerts.py@ad3758c:181
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: source analysis confirms

### KF-0236: Alerts marked sent when sending failed

- Area: Alerts
- What went wrong: An alert is marked as sent when sending failed, one failing check disables the rest, and that failure leaves its database connection open.
- Correct behaviour: An alert is marked sent only after the send succeeded; each check runs on its own; every database connection is closed on every path, also on an error (derived: roadmap 745).
- Source: private audit 2026-11-10, bot item M5
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0237: Bot and API report success on failure

- Area: Surfaces
- What went wrong: Several bot commands (`/adblock`, `/dyndns`, `/lang` and others) and API endpoints (stop a speed test, reload, apply settings, Telegram save) report success when the action failed or was only requested.
- Correct behaviour: A surface reports the real outcome; a request that was only queued says so (work plan 7 point 1, real status codes).
- Source: private audit 2026-11-10, bot item M8; private audit 2026-11-10, web item 13
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-02 (also S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0238: Long bot jobs block polling

- Area: Telegram bot
- What went wrong: Long jobs other than speed tests still block the bot's polling.
- Correct behaviour: Long jobs run in their own task; the bot keeps answering (CHANGELOG 263).
- Source: private audit 2026-11-10, bot item L1; CHANGELOG.md@ad3758c:263
- Python fix commit: e0201ed; note: speed tests only
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: fixed for speed tests

### KF-0239: English left in bot texts; collector pages only in two languages

- Area: Translations
- What went wrong: A few bot texts are English in every language, and the collector's pages exist only in English and Hungarian (Hungarian labels in the collector, `nmb/collector/server.py@ad3758c:488`).
- Correct behaviour: Every text in all 11 languages (roadmap 602). Owner, 2026-11-11: the build refuses a built-in text missing in any language; a text still missing at run time shows the English text visibly marked as untranslated, is logged and is listed in the self-test result and the diagnostic bundle, never a silent fallback to any language (KF-0048).
- Source: private audit 2026-11-10, bot item L7; nmb/collector/server.py@ad3758c:488
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-06 (also S-11, S-13)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0240: Collector client can stop after a bad reply

- Area: Collector client
- What went wrong: The collector client can stop after a bad reply and then sends nothing until it is restarted.
- Correct behaviour: A bad reply is logged and retried later; the client keeps running (roadmap 745).
- Source: private audit 2026-11-10, bot item M3
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-12
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0241: Ping-hour upload drops targets at a batch cut

- Area: Collector client
- What went wrong: When a batch cuts an hour, the hourly ping upload drops targets.
- Correct behaviour: Hours are uploaded whole; a cut hour is completed in the next batch (derived: roadmap 4.25, 391: hourly ping summaries are part of the collector protocol; Quality Bar 6, roadmap 581).
- Source: private audit 2026-11-10, bot item M7
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-12
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0242: Stale in-memory settings overwrite other writers; changes before validation

- Area: Configuration
- What went wrong: The web server's in-memory settings can overwrite what the bot wrote, handlers change the settings before they validate them, and an uploaded restore does not update the scheduler's copy.
- Correct behaviour: Validate first, then change; every component reloads after another one saved (CHANGELOG 290 merge rule).
- Source: private audit 2026-11-10, web item 9
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-01 (also S-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0243: Heavy work on every status poll

- Area: Web server
- What went wrong: Each 3-second status poll runs a 24-hour aggregate, chart drawing, queries, a live DNS test and subprocesses; a network timeout is set process-wide.
- Correct behaviour: Status is served from cached state updated by the components; one live stream instead of polling timers (roadmap 4.4A).
- Source: private audit 2026-11-10, web item 10
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-02 (also U-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0244: `wifi_psk_set` means "not weak"

- Area: Web server
- What went wrong: The API's `wifi_psk_set` is true when the password is not weak, not when one is set (`nmb/web/server.py@ad3758c:3834`).
- Correct behaviour: A field means what its name says: "set" and "weak" are separate facts.
- Source: private audit 2026-11-10, web item 11; nmb/web/server.py@ad3758c:3834
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: source analysis confirms

### KF-0245: Secret kept or replaced by its display text

- Area: Web server
- What went wrong: Whether a secret field was left unchanged is decided from its masked display text.
- Correct behaviour: An unchanged secret is a state of the form (the `__KEEP_CURRENT__` mark or an explicit flag), never a comparison with display text (roadmap 583, 682).
- Source: private audit 2026-11-10, web item 18
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-02 (also C-04)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0246: Mode switch leaves a half-built state

- Area: Network modes
- What went wrong: An exception during mode setup leaves a half-built state with no retry, and watcher failures are logged only at debug level.
- Correct behaviour: A mode change is applied completely or rolled back, with a timed confirm (porting plan 7D 626; roadmap 6B.5); failures are warnings.
- Source: private audit 2026-11-10, netd item M11
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: N-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0247: LAN subnet assumed to be /24

- Area: Network modes
- What went wrong: The LAN is assumed to be a /24 although the netmask is a setting.
- Correct behaviour: The configured netmask is used everywhere (derived: Quality Bar 8, roadmap 594: no value that has to be changed is written into the code).
- Source: private audit 2026-11-10, netd item L5
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: N-05 (also N-03)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0248: Storage guard may delete evidence while the card stays full

- Area: Storage guard
- What went wrong: The storage guard may delete days of the live database (down to 1-7 days, including long-kept evidence) while the card stays full because the snapshot needs space there (marked as a suspicion in the audit).
- Correct behaviour: Measurements are deleted last, a day at a time, keeping `keep_days` (roadmap 4.26, 403); the guard checks the space it really needs before deleting anything (derived: Quality Bar 6, roadmap 581: no fragile side effect; roadmap 403: measurements go last).
- Source: private audit 2026-11-10, reports item 8 (suspicion); docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:403
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-08
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: unconfirmed

### KF-0249: Scenario restore ignores failures and can bring back cleared secrets

- Area: Self-test
- What went wrong: A scenario's restore ignores API failures and deletes its snapshot; the 45-minute watchdog can fire mid-run, a stale snapshot leaves no watchdog, nothing restores after a reboot; the restore brings back secrets the owner cleared.
- Correct behaviour: A restore checks each step, keeps its snapshot until it succeeded, survives a reboot, and restores only what the scenario changed (TESTING 233 describes the scenario and what it changes; derived: Quality Bar 6, roadmap 581, no fragile side effect is carried over).
- Source: private audit 2026-11-10, reports items 9 and 10
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: P-07
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0250: Crash mirror to the boot drive without a rate limit

- Area: Crash evidence
- What went wrong: A crash loop writes the crash mirror to the boot drive every few seconds.
- Correct behaviour: Crash evidence is kept bounded and flash-safe (porting plan 7D 617-618).
- Source: private audit 2026-11-10, reports item 11
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0251: Boot drive and live database stressed during measurements

- Area: Flash and database
- What went wrong: Writability tests write files to the boot drive every 3 minutes, and a full integrity check of the live database runs twice every 3 minutes, plus network diagnostics, during measurements.
- Correct behaviour: Flash writes are rare and small; heavy checks do not run during a measurement (porting plan 7D 617).
- Source: private audit 2026-11-10, reports item 15
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-03 (also C-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0252: Database snapshot and restore edge cases

- Area: Database
- What went wrong: The snapshot opens its source read-write (it can create an empty database and back it up over the good one), falls back silently to another location where its copy does not last, and a restore aborts on "not a database".
- Correct behaviour: The source is opened read-only; there is no silent fallback; a bad file is refused with a message and nothing is overwritten (derived: roadmap 634).
- Source: private audit 2026-11-10, reports item 16
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0253: Stored timestamps and retention follow wall-clock steps

- Area: Database
- What went wrong: The minute rollup skips samples after a backward clock step, a forward step prunes data, and metrics count gaps as uptime.
- Correct behaviour: Retention and rollups survive clock steps; a gap is unknown, not uptime (derived: QB6 583).
- Source: private audit 2026-11-10, reports item 17
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-02 (also R-01)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0254: Developer log shared by several processes

- Area: Logging
- What went wrong: One rotating developer log is written by several processes (races, a fallback location).
- Correct behaviour: One writer per log, or a log service; no fallback location (porting plan 7D 615).
- Source: private audit 2026-11-10, reports item 21
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0255: Diagnostic bundles never pruned

- Area: Diagnostics
- What went wrong: Bundles are never pruned, use the same archive name for three different configs, and are written to the small log disk.
- Correct behaviour: Bundles are bounded in number and size, unique names, on the data disk (roadmap 4.26 storage guard).
- Source: private audit 2026-11-10, reports item 14
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: R-04 (also C-08)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0256: Rollback restores only the code

- Area: Updates
- What went wrong: A rollback restores only the code folders; changed units and rules stay; a health check counts HTTP 500 as healthy.
- Correct behaviour: An update is applied beside the running version and the whole previous version comes back on failure; healthy means services up, network reachable, Web UI answering correctly (porting plan 7D 623-624).
- Source: private audit 2026-11-10, privileged-operations item 15
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: P-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0257: Logs lost at reboot

- Area: Logging
- What went wrong: The developer log lives on a RAM disk and is lost at reboot; only a small part is kept on the card.
- Correct behaviour: Crash evidence and the last log lines survive a reboot (porting plan 7D 618).
- Source: private data-quality list 2026-11-10, missing items; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:618
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: partly kept on the card

### KF-0258: SD card writes of the developer log

- Area: Logging
- What went wrong: The developer log grows about 10 MB a day on the card.
- Correct behaviour: Flash-safe logging (porting plan 7D 617).
- Source: docs/TESTING.md@ad3758c:387; private log review 2026-11-09, still open
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0259: Reapplying a network profile reset open connections

- Area: Network
- What went wrong: Putting IPv4 DNS first reapplies the profile once, which reset an open connection of the bot.
- Correct behaviour: DNS order is set without resetting live connections (derived: Quality Bar 6, roadmap 581: no side effect is carried over).
- Source: docs/TESTING.md@ad3758c:389
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: N-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0260: pppd warning on every dial

- Area: PPPoE
- What went wrong: Every dial logs "can't open options file /root/.ppprc: Permission denied" (harmless).
- Correct behaviour: No warning that is expected: the native PPPoE path does not read a user options file it cannot open (derived: Quality Bar 6, roadmap 581: no confusing message is carried over).
- Source: docs/TESTING.md@ad3758c:390
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: N-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0261: Per-test facts still missing

- Area: Measurement data
- What went wrong: Rows do not store the engine or binary version, bytes moved, streams, the IP version used, other traffic seen during the test, the CPU peak or the box's id.
- Correct behaviour: Each row stores the facts an ISP or an authority needs: engine version, bytes, streams, IP version, gate reading, CPU peak, the box's id (private data-quality list 2026-11-10, missing items; derived: roadmap 4.29, 450, "stated, not inferred by the reader").
- Source: private data-quality list 2026-11-10, missing items and still open
- Python fix commit: 1012b59; note: duration, version, mode
- Package that proves it: M-04 (also C-02, R-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: partly fixed

### KF-0262: Public address without its age

- Area: Measurement data
- What went wrong: A row's public address does not say how old the lookup was.
- Correct behaviour: The address's source and age are stored (private data-quality list 2026-11-10, still open).
- Source: private data-quality list 2026-11-10, still open
- Python fix commit: b8fb934; note: the source
- Package that proves it: M-12
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: source stored, age not

### KF-0263: Collector keeps old placeholders until boxes send corrections

- Area: Central collector
- What went wrong: The collector still holds placeholder values that boxes uploaded before the corrections, until each box sends its corrections.
- Correct behaviour: The native collector applies corrections when it imports a Python collector's data (derived: roadmap 4.25, 392: a Python collector's data opens in the native one).
- Source: private data-quality list 2026-11-10, still open; CHANGELOG.md@ad3758c:171
- Python fix commit: b8fb934
- Package that proves it: S-13
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: until boxes send corrections

### KF-0264: Exports carry no integrity proof

- Area: Exports
- What went wrong: Nothing ties a CSV to a box and a time range in a way that can be checked; there is no signature over an export.
- Correct behaviour: Exports name the box and the range (CHANGELOG 23) and carry a signature made with the box's own key over the export, its box id and its range, which anyone can check with the box's public key; a bare checksum is not enough, because whoever changes the file can compute a new one (derived: roadmap 589 and 6B.3, 860: accepted only with a valid signature; work plan section 7 point 6, line 325: signing yes; private data-quality list, missing items).
- Source: private data-quality list 2026-11-10, missing items; CHANGELOG.md@ad3758c:23; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:589; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:860; docs/NATIVE_WORK_PLAN.md@ad3758c:325
- Python fix commit: 1012b59; note: range in the file name
- Package that proves it: R-02 (also R-03, P-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: partly fixed

### KF-0265: Test scripts outside the repository

- Area: Tests
- What went wrong: The VM scripts and bot check scripts used for the recorded runs are not part of the repository.
- Correct behaviour: Every test that proves a result is in the repository and runs from it (work plan 2.4, hermetic tests; roadmap 747).
- Source: docs/TESTING.md@ad3758c:399
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: F-06 (also P-13)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0266: SSH accepts passwords by default

- Area: Remote access
- What went wrong: The box's SSH accepts passwords; keys-only is an option, off by default.
- Correct behaviour: Developer access is SSH with keys only, never on the WAN, inside the owner-enabled maintenance mode (porting plan 7D 594).
- Source: private audit 2026-11-10, box from outside; CHANGELOG.md@ad3758c:27; docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md@ad3758c:594
- Python fix commit: b624b96; note: the option
- Package that proves it: P-05
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: an option exists, off by default

### KF-0267: Privilege boundary between the service user and root (category)

- Area: Privileged components and files
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Quality Bar 7 at the privilege boundary: everything a less privileged user can write, name or replace is outside input to a privileged component, validated by type, length and allowed values; no path leaves its directory; no command is built from it; errors fail closed (roadmap 587); privileged code never runs anything from `/tmp` or from a folder the service user can write (roadmap 679), and its helper arguments are validated (roadmap 680).
- Source: private audit 2026-11-10, privileged-operations items 4, 5, 15 and 17; reports item 1; netd items M8 and L8; engines item 19
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: P-06
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0268: Files from outside: uploads, restores and imports (category)

- Area: Backup, restore, USB import, uploads
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Every uploaded, restored or imported file is outside input (roadmap 587): validated by type, length and allowed values before it is used; no path leaves its directory; no command is built from it; errors fail closed; every parser that reads such a file is fuzz-tested (roadmap 588).
- Source: private audit 2026-11-10, privileged-operations items 6, 7 and 14; web items 3a and 15
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: P-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0269: Services with more privilege than they need (category)

- Area: Services and sandboxing
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Daemons run as their own unprivileged user with only the capabilities they need, under systemd or procd sandboxing (read-only system, private `/tmp`, no new privileges, a system call filter where the platform has it) (roadmap 589, 6B.5); the web server does not run as root (roadmap 681).
- Source: private audit 2026-11-10, privileged-operations items 10 and 11; engines item 3
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: P-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0270: Integrity of updates and installed programs (category)

- Area: Updates and installed programs
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Updates and firmware are accepted only with a valid signature (roadmap 589, 6B.3); the checksum of every downloaded binary is verified, and nothing is run from `/tmp` or from a folder the service user can write (roadmap 679); errors fail closed (roadmap 587).
- Source: private audit 2026-11-10, privileged-operations items 8, 9, 16 and 17
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: P-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0271: Transport security of the box's connections (category)

- Area: Outgoing and incoming connections
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Secrets never reach logs, URLs, command lines or the UI, and errors fail closed (roadmap 587); a connection that carries secrets, results or data the box acts on is protected and verified (derived: roadmap 587, 589); the external scan before each release checks TLS (roadmap 591).
- Source: private audit 2026-11-10, reports items 5 and 20; bot and collector items H4 and L3; engines items 6 and 22; netd items M1 and L4; web item 17
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-12
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0272: Values written into the configuration of system daemons (category)

- Area: Configuration generators
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Every value is validated by type, length and allowed values before it is used; no command is built from a value (no shell, arguments as a list); no path leaves its directory; errors fail closed (roadmap 587); every configuration parser and writer is fuzz-tested (roadmap 588).
- Source: private audit 2026-11-10, netd items M1, M6, M7 and L7; web item 12; reports item 7 (one part)
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0273: Web UI and API request handling (category)

- Area: Web server and API
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Every request is outside input, validated by type, length and allowed values; authentication and authorisation are checked on every endpoint; errors fail closed (roadmap 587); the Web UI has a strict Content Security Policy, CSRF protection, secure cookies or headers, rate-limited login and constant-time comparisons (roadmap 589); the HTTP parser is fuzz-tested (roadmap 588).
- Source: private audit 2026-11-10, web items 6, 8, 13 (one part), 14, 16, 17 and 19; engines item 25
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0274: Network exposure and packet filtering (category)

- Area: Firewall, listeners and packets from the network
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Nothing listens on the WAN side unless the owner turned it on; every packet from outside (DHCP, DNS, PPPoE, UPnP) is outside input, validated by type, length and allowed values; errors fail closed (roadmap 587); local service ports answer only on the networks they serve, on IPv4 and IPv6 (owner, 2026-11-11, trust of the uplink network, which replaces the address-class wording of roadmap 681); an external scan from the LAN and the WAN side before each release (roadmap 591).
- Source: private audit 2026-11-10, web item 5; netd items H1 (one part), M4, M5 and M12; engines item 3
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: N-02
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0275: Secrets and personal data (category)

- Area: Secrets and personal data
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Secrets never reach logs, URLs, command lines, crash reports, diagnostic bundles or the UI (roadmap 587); everything that leaves the box or lands on the boot partition is masked by the one shared list of secrets and personal details (roadmap 4.24, 682); the repository and the built image are scanned for secrets (roadmap 588), and the image is checked for leftover keys and test accounts before each release (roadmap 591).
- Source: private audit 2026-11-10, privileged-operations items 12, 13 and 18; reports items 4, 6, 11, 14 and 18; netd items L1 and L3; bot and collector items L5 and M8; web item 19; engines item 16
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-04
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0276: Central collector server and its client connection (category)

- Area: Central collector
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: The collector is a network-facing server, and the box's connection to it carries results and credentials: every point of Quality Bar 7 (roadmap 587-591) applies to the server and to the client as it does to the box's own Web UI (KF-0273); the protocol stays as roadmap 4.25 describes it (lines 391-392).
- Source: private audit 2026-11-10, bot and collector items H4, M1, M6, L3 and L4
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-13
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0277: Telegram bot commands (category)

- Area: Telegram bot
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Authentication and authorisation are checked on every bot command; every argument is outside input, validated by type, length and allowed values; secrets never reach the chat; errors fail closed (roadmap 587). Owner, 2026-11-11: every bot command that changes the box needs a confirmation tap, and the destructive ones (factory reset, purge) also need the typed word, as in the Web UI; read-only commands run at once.
- Source: private audit 2026-11-10, bot and collector items M2 (one part), M8 (one part) and L5 (one part)
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0278: Resource limits of network-facing services (category)

- Area: Network-facing services
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: A network-facing service stays available: what one client can make it use is bounded, and one bad client or one error never stops the service for the others (derived: roadmap 587, "errors fail closed", and the rule of KF-0024); every parser on such a service is fuzz-tested (roadmap 588).
- Source: private audit 2026-11-10, web items 4 and 15; bot and collector item M6; netd items M3, M10 and L4; engines item 22
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0279: Errors that fail open (category)

- Area: Error handling
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Errors fail closed (roadmap 587): when something a security decision depends on fails, the component stays in the safer state and says so; it never continues in a weaker one.
- Source: private audit 2026-11-10, reports item 13; web items 5 and 13 (one part); engines item 16
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0280: Network tools a user can start (category)

- Area: Measurement and diagnostic tools
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: A tool a user can start from the Web UI, the CLI or the bot is checked like any other action: its arguments are outside input, validated by type, length and allowed values; it is authorised for that person; nothing it opens listens on the WAN side unless the owner turned it on; secrets it sees never leave the box (roadmap 587).
- Source: private audit 2026-11-10, engines items 3, 16 and 25; web item 6 (one part)
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-11
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0281: Outside data the box downloads and acts on (category)

- Area: Downloaded lists and lookups
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Downloaded lists, directories and lookups are outside input: validated by type, length and allowed values before use; no command is built from them; errors fail closed (roadmap 587); their parsers are fuzz-tested (roadmap 588); transport as in KF-0271.
- Source: private audit 2026-11-10, engines items 6 and 22; netd items M1, L4, L6 and L7
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-12
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0282: Authentication and credentials on every surface (category)

- Area: Authentication
- What went wrong: Open security faults of this category exist in the Python version. Details are private until they are fixed (roadmap 608).
- Correct behaviour: Authentication and authorisation are checked on every endpoint, bot command and client connection, and credentials never reach logs, URLs or command lines (roadmap 587); rate-limited login and constant-time comparisons on every surface (roadmap 589); every device has its own login (roadmap 592).
- Source: private audit 2026-11-10, web item 17; bot and collector item M1; netd item M10; reports items 13 and 20
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-01
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: security, details private until fixed

### KF-0283: Licence file and notices of bundled code missing

- Area: Licences
- What went wrong: The README states the MIT licence and links a `LICENSE` file that does not exist in the repository; the bundled RTR measurement client in `components/open-rmbt/` (code from another project, changed for NMB) carries no licence file, no notice and no mark of the changes.
- Correct behaviour: Every component NMB bundles, builds in, downloads or uses is listed with its licence, origin and version; the full licence texts are shipped; code taken from another project keeps its licence file and notices, and files changed by NMB say so; a missing licence file or notice, an unknown licence or a component without an entry fails the build (roadmap 868-871, 6B.3A). The upstream licence of the RTR client (Apache-2.0 according to the project notes) is confirmed by P-08 from the upstream project.
- Source: README.md@ad3758c:5; README.md@ad3758c:485-487; git ls-tree ad3758c (no `LICENSE`, no licence or notice file anywhere, `components/open-rmbt/` included); `git grep -i licen ad3758c -- components` (no hit); docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:866-871
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: P-08 (also P-11, M-08)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0284: iPerf3 schedule settings had no effect

- Area: iPerf3 and schedule
- What went wrong: `iperf.schedule_cron` and `iperf.automated_enabled` exist in the settings and in the Web UI, but nothing schedules iPerf3 from them: the cron value is loaded and saved only "for older configs" (`nmb/config.py@ad3758c:611`), and `automated_enabled` is only stored and shown by the web server (`nmb/web/server.py@ad3758c:3881`, `4193-4194`). Scheduled iPerf3 runs come from `[speedtest] iperf3_interval_min` instead (CHANGELOG 90, 121).
- Correct behaviour: Both settings work: with `automated_enabled` set, iPerf3 runs on its own schedule through the common runner, under the same traffic gate, data caps and hourly limit as the other engines (work plan 7 point 2; roadmap 756); how it relates to `iperf3_interval_min` is written down by M-04. Every setting that exists has an effect, and a test fails when a setting is read by nothing (derived: roadmap 581, Quality Bar 6).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:756; docs/NATIVE_WORK_PLAN.md@ad3758c:311-312; nmb/config.py@ad3758c:611; nmb/web/server.py@ad3758c:3881; nmb/web/server.py@ad3758c:4193-4194; CHANGELOG.md@ad3758c:90; CHANGELOG.md@ad3758c:121
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-04 (also M-09, U-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0285: API reported failures as HTTP 200

- Area: Web server and API
- What went wrong: Many endpoints report a failure as HTTP 200 with `"status": "error"`, and the dashboard relies on it.
- Correct behaviour: Real HTTP status codes (2xx only for success, 4xx for a refused or invalid request, 5xx for a failure of the box) and one error body (`code`, a translation key, detail); the native Web UI and `nmb-cli` are written for this; the collector protocol and the `nmb-cli` JSON command names stay as they are; the difference is listed in the changelog under Changed (work plan 7 point 1).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:757; docs/NATIVE_WORK_PLAN.md@ad3758c:306-310
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-01 (also S-02, U-01, S-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0286: Cloudflare downloads drop the TLS stream on the Pi (cause unknown)

- Area: Cloudflare engine
- What went wrong: Cloudflare downloads sometimes drop a TLS stream on the Raspberry Pi (`RECORD_LAYER_FAILURE`). The cause is unknown; GRO, receive checksums, read size, the latency probe and the TLS version were ruled out, and results stayed in the same range.
- Correct behaviour: A Cloudflare download keeps its streams to the end on the same box; a stream that is dropped is reported as such in the row, never hidden in the figure (derived: Quality Bar 6, roadmap 581). M-06 measures how often streams drop with the native TLS stack against the Python engine on the Pi 4, the same day, and writes the cause into this entry when it is found (roadmap 755: a different TLS library will show whether the cause is OpenSSL).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:755
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-06
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: cause unknown

### KF-0287: `nmb-cli` actions and Telegram admin commands not translated

- Area: Translations
- What went wrong: The `nmb-cli` action names and the Telegram admin commands were never finished in every language (roadmap 579 lists them among the open items the Python version never finished; roadmap 674: translated action names not yet covered).
- Correct behaviour: Every `nmb-cli` action label and every bot admin command text exists in all 11 languages (roadmap 579, 602); the JSON command names and their output stay as they are (roadmap 668). Owner, 2026-11-11: the build refuses a built-in text missing in any language; a text still missing at run time shows the English text visibly marked as untranslated, is logged and is listed in the self-test result and the diagnostic bundle, never a silent fallback to any language (KF-0048).
- Source: docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:579; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:602; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:668; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:674
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-06 (also S-10, S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open

### KF-0288: Documented features that did not exist or worked differently

- Area: Documentation and alerts
- What went wrong: Latency alerts were documented but did not exist; the READMEs described features differently from how they worked (for example the bridge needing an IP address and the self-update needing a Git checkout).
- Correct behaviour: Every alert, command and feature the documentation names exists and works as described, with a test per documented alert; the manual and the command reference are generated from the code and checked by the build (roadmap 6B.7).
- Source: CHANGELOG.md@ad3758c:105; CHANGELOG.md@ad3758c:132
- Python fix commit: 686b8aa
- Package that proves it: S-11 (also P-12)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0289: Exports left out stored facts and had no range or target role

- Area: Exports
- What went wrong: The speed CSV left out facts stored in the database (the interface and local address used, retransmits, the TCP window); the ping export was a fixed 24 hours with no range parameter and did not name each target's role.
- Correct behaviour: Every stored fact of a row is exportable (new columns after the old ones, which keep their order); every export takes a range (`?from=`/`?to=` as a date, ISO time or epoch ms, or `?days=`/`?hours=`) and its file name says the range; the ping export names each target's role (CHANGELOG 23). Owner, 2026-11-11: the ping evidence is the minute summaries only; there is no raw ping-sample export.
- Source: private data-quality list 2026-11-10, missing items; CHANGELOG.md@ad3758c:23
- Python fix commit: 1012b59
- Package that proves it: R-02 (also S-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0290: Data written into the home folder of whoever runs the code

- Area: Files and folders
- What went wrong: When their data folder cannot be written, components fall back to a folder under the home folder of the user that runs them, for example crash reports and diagnostic bundles (`nmb/report/crash_trap.py@ad3758c:29`, `35`, `375`) and the diagnostic snapshot (`nmb/report/invasive_logger.py@ad3758c:46`). Data then lands outside the box's data folders, where the rest of NMB does not look, and a run on another machine (a test on the development PC) writes into that user's home folder.
- Correct behaviour: Every component takes its folders from one path table; when a folder cannot be written it reports the error and writes nowhere else (derived: roadmap 584, right by design, not because of a file on the developer's PC); tests point every path at throwaway folders and fail when anything is written outside them (KF-0047). The analyst of C-03 finds every such fallback in the Python code before the native paths are fixed.
- Source: nmb/report/crash_trap.py@ad3758c:29; nmb/report/crash_trap.py@ad3758c:35; nmb/report/crash_trap.py@ad3758c:375; nmb/report/invasive_logger.py@ad3758c:46; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:584; project memory note (a unit test wrote a file into the home folder), confirmed by this source analysis
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: C-03 (also R-05, P-01, P-03, F-06)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: source analysis

### KF-0291: Bot reload request can reach the wrong process or none

- Area: Telegram bot
- What went wrong: The rest of the reload fault whose token and recipient part is KF-0011: a reload request can go to the wrong process or to none.
- Correct behaviour: A reload is passed to the running service through its supervisor (systemd or procd) or the native IPC; a service that is not running is reported as not running and nothing is signalled; a reload can reach no other process (derived: Quality Bar 6, roadmap 581: no fragile side effects; KF-0011).
- Source: private audit 2026-11-10, bot item M4 (the parts left open) and fix progress (301e1bd fixed the token and recipient part only)
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: S-11 (also S-02, P-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: the token and recipient part is fixed (KF-0011)

### KF-0292: Ookla refused the box every few hours at an hourly schedule (cause not proven)

- Area: Ookla engine
- What went wrong: With scheduled Ookla runs at most hourly, Ookla still refused the box every few hours (`docs/TESTING.md@ad3758c:386` lists the times). One cause is fixed: Ookla ran as a new client each time, from the bot and from the service user (KF-0087, KF-0169); whether that ended the refusals is not proven (`CHANGELOG.md@ad3758c:297`).
- Correct behaviour: The box never adds to a block itself: one persistent Ookla client identity for every caller (schedule, Web UI, CLI, bot), scheduled runs at most hourly (KF-0094), the pause after a refusal kept across restarts and followed by every process (KF-0093, KF-0096), no hidden retries (KF-0225); M-05 and M-10 record the refusal rate at an hourly schedule against the Python engine on the Pi 4 on the same line, and a native rate above the Python one is a fault (derived: roadmap 690; Quality Bar 6, roadmap 581). Owner, 2026-11-11: a test started by a person (Web UI, CLI, bot) is not held back by the hourly schedule limit, but during a refusal pause it is refused with the time left.
- Source: docs/TESTING.md@ad3758c:386; CHANGELOG.md@ad3758c:297; private log review 2026-11-09, still open (Ookla rate limits)
- Python fix commit: 002548c, 51dfd1b; note: one cause only (KF-0087, KF-0169)
- Package that proves it: M-05 (also M-10)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open; note: cause not proven

### KF-0293: `nmb-cli` lacked actions for endpoints the Web UI calls

- Area: Terminal interface (`nmb-cli`)
- What went wrong: Endpoints the Web UI calls had no `nmb-cli` action: the two diagnostic downloads, the main Wi-Fi client port role, the disk cleanup settings and the Ookla terms. The automatic comparison did not catch them; they were found by comparing the endpoints the Web UI really calls by hand.
- Correct behaviour: The Web UI and `nmb-cli` take their menus from one source, and a test fails when a Web UI tab or API call has no CLI action (roadmap 666); the test takes the endpoints the Web UI really calls, so none is missed as these were (TESTING 183: these were found only by comparing those endpoints by hand; derived: Quality Bar 6, roadmap 584, right by design, not by luck).
- Source: docs/TESTING.md@ad3758c:183; docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:664-666; CHANGELOG.md@ad3758c:67
- Python fix commit: 6b1ad91, eab2be7
- Package that proves it: S-10 (also U-02)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0294: Metered links tested on a built-in cap or on none

- Area: Data caps
- What went wrong: On a metered link the scheduled speed tests ran without the user having chosen a cap: the USB link came with a built-in daily limit of 1024 MB, and the per-link data caps were 0 (no cap) by default.
- Correct behaviour: Owner, 2026-11-11: there is no default cap number; on a metered link scheduled speed tests wait until the user sets a cap or says unlimited; ping keeps running. A waiting test is recorded as skipped with its reason, never as 0 (roadmap 683), and the status says what the user has to choose (derived: Quality Bar 6, roadmap 581: no confusing message). Owner, 2026-11-11: a speed test started by a person (Web UI, CLI, bot) on such a link waits like a scheduled test until the user sets a cap or says unlimited, and the person who started it is told that it waits and why; ping keeps running. A cap that is reached is a different case and keeps CHANGELOG 35 (tests started by hand still run and count).
- Source: docs/TESTING.md@ad3758c:388; CHANGELOG.md@ad3758c:35; nmb/config.py@ad3758c:693; nmb/config.py@ad3758c:728-731 (source analysis)
- Python fix commit: not recorded in the source; note: not fixed
- Package that proves it: M-05 (also C-01, U-02, S-10, S-11)
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: open
