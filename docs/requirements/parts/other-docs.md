# Requirements register, part: the documents no other part read

Sources, all of the Python repository at commit `ad3758c`, each read in full with `git show ad3758c:<file>` in pages:

- `docs/NMHH_SZELESSAVNET_COMPARISON.md` (95 lines, read 1-95): the source of the REQ-NH rows.
- `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` (102 lines, read 1-102), compared line by line with the English file.
- `docs/DEV_SETUP_WINDOWS.hu.md` (220 lines, read 1-220), compared line by line with `docs/DEV_SETUP_WINDOWS.md` (214 lines,
  read 1-214; its requirements are the REQ-DV rows of `workplan-architecture-devsetup.md`).
- `docs/TESTING.hu.md` (399 lines, read 1-80, 81-138, 139-213, 214-298, 299-399), compared line by line with
  `docs/TESTING.md` (401 lines, read 1-80, 81-140, 141-215, 216-300, 301-401). This completes the comparison that
  `testing.md` made for lines 22-275 by structure only (its header).
- `CONTRIBUTING.md` and `CONTRIBUTING.hu.md` (55 lines each, read 1-55): already the sources of `contributing.md`
  (REQ-CT-0001 to REQ-CT-0026). Read again and compared line by line in this round: no requirement is missing and the
  Hungarian file has no content of its own (its line 10 says a test elsewhere is "very useful" where the English says
  "welcome"; same meaning). No REQ-CT id is added, so the next free one stays REQ-CT-0027, and their coverage stays in
  `contributing.md` (one table per source).
- `README.hu.md` (721 lines, read 1-120, 121-149, 150-260, 260-370, 371-470, 470-580, 581-721), compared line by line with
  `README.md` (487 lines, read 1-110, 111-200, 201-330, 331-487) and with the Hungarian table of `readme.md`: what that
  table already lists is not repeated; what it lacked is listed in *Content only in the Hungarian version* below, and
  the one new requirement (REQ-RE-0341) and the new context rows are added to `readme.md`, where `README.hu.md` is covered.

Quality Bar 9: the comparison is about one country's official measuring box. Nothing of it is built into the native
code: its statements become requirements of the country data files (package C-10), of the neutral international
default (a box with no country file, or a country file without an official tool), and of the packages that show what a
country file holds (P-08 disclaimer, R-03 dossier, P-12 manual). Quality Bar 8: the owner's experience, the screenshot's
values, the provider names of the owner's line and the official server's address are not copied; the public names of the
regulator and its programme are, because the Hungary file cannot be checked without them (register README, *Names and
site data*).

Package ids are those of `docs/NATIVE_WORK_PLAN.md@ad3758c:98-226`; the open series are written as in the other parts.
Ids REQ-NH-0001 to REQ-NH-0013 (prefix new in round 3). Ids are permanent; the next free one is REQ-NH-0014.

## Requirements

| Id | Requirement | Source | Package | Also in | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| REQ-NH-0001 | Material that compares NMB with one country's official measuring tool (the tool's terms, its method, when to use which) is content of that country's data file and of that country's page of the manual, shown only on a box set to that country; none of it is in the code, the general Web UI or the general manual. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:3; docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:78 | C-10 | P-12, C-06 | QB 9 (REQ-RM-1104, REQ-RM-1109). The Python comparison is a document for one country; in the native build it becomes the Hungary file's (REQ-NH-0004). |
| REQ-NH-0002 | The disclaimer shown with the results (Web UI, dossier, bot, CLI) says, besides the open-source "as is" warranty and liability exclusion, that the results are indicative and not a certified measurement, and that an official complaint needs the certified measuring tool of the user's country; that tool is named only from the box's country data file, and without one the text names no tool and no regulator. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:14; docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:57; docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:91 | P-08 | C-10, R-03, C-06, U-01, S-10, S-11 | Owner decision of 2026-11-11 on the disclaimer wording (open in REQ-WP-0187). Refines REQ-RM-1344 and REQ-RE-0223; the dossier's own "no official claim" is REQ-RE-0151. |
| REQ-NH-0003 | A country data file's entry for the official measuring tool holds the regulator, the tool's name, a link to its official page (and to its terms where they are published), what its results may be used for (regulatory procedure, complaint, national statistics) and the date the entry was last checked; it links to the tool's terms (who may take part, duties, return of a device) and does not copy them, because the regulator's current text always applies. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:7-9; docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:14; docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:95 | C-10 | P-12 | Refines REQ-RM-1108 ("regulator and official measurement tool"). The comparison's sections on eligibility and contract duties (lines 20-36, 44) are such terms. |
| REQ-NH-0004 | The country data file for Hungary names the national regulator (NMHH) and its broadband measurement programme (Szélessáv Program, szelessav.net) as the official measuring tool, with the links of the programme's application and methodology pages, and says that the programme's results can serve the regulator's procedures while NMB's cannot. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:3; docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:7-9; docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:57 | C-10 | P-08, R-03 | Public regulator and programme names are copied (register README, QB 8 rule for public identifiers); the programme's terms are linked, not copied (REQ-NH-0003). |
| REQ-NH-0005 | Where the country data file names an official tool, the manual's page for that country says when to use which: the official tool for an official measurement or a complaint, NMB for an own, continuously running device with several engines and servers, and both together, NMB's results complementing the official ones without replacing them; the neutral default says the same without naming a tool. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:89-91 | P-12 | C-10, C-06 | |
| REQ-NH-0006 | Every statement of a country comparison carries its source: a document or web page with the date it was read, NMB's own documentation and test record, or the author's own experience marked as not confirmed by an official source; statements about NMB cite the native build's test record and are checked against it before each release. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:5-12 | P-12 | P-13, C-10 | Honest results (REQ-CT-0019). The owner's experience and screenshot are sources of the Python page only (source names a specific site value; not copied). |
| REQ-NH-0007 | The texts of a country data file (official tool description, when-to-use text, comparison) exist in every UI language, not only in the country's own language. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:3 | C-10 | C-06 | REQ-RM-1110 applied to country files; the Python comparison exists in English and Hungarian only. |
| REQ-NH-0008 | The country data file format has an optional link to the country's public broadband statistics or map (and to an EU-wide statistics project where one applies); the dossier and the manual show it only when the box's country file has one, and the neutral default has none. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:63 | C-10 | R-03, P-12 | The Hungarian regulator publishes aggregated national results for everyone. |
| REQ-NH-0009 | Upload data that the native build's own measuring clients generate is random (incompressible), so compression on the path cannot raise a measured speed. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:62 | M-04 | M-06, M-07, M-09 | The official method measures with random-content data blocks. If `NEW: M-17 Shared HTTP transfer core` is accepted, it owns this row. The analysts of M-06, M-07 and M-09 check what the Python engines send. |
| REQ-NH-0010 | The line settings keep the provider the contract is with and, optionally, the network operator the line runs on (a line resold over another company's network), and the dossier shows both, next to the provider name the measuring services report. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:48; docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md@ad3758c:54 | R-03 | C-01, M-12, U-02…U-0n (page: SLA) | An official register can name another provider than the one that serves the line, "especially when the subscription runs on another company's network" (Hungarian file only); the dossier must say which company is meant. Provider names not copied (source names a specific site value; not copied). |
| REQ-NH-0011 | The SLA dossier prints, or saves as PDF from the browser, with every chart and table readable (a print layout), so it can be handed in on paper. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:58 | R-03 | | The official programme offers its results in printable form. |
| REQ-NH-0012 | No country data file lists an official measuring tool's own servers as NMB test servers unless their operator publishes them for public use; the box measures only against servers meant for it. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:38; docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:40 | C-10 | M-04, M-09 | Respect for the services measured against (REQ-CT-0006, REQ-CT-0007). The official box tests against its operator's own server at the national internet exchange; its address is not copied. |
| REQ-NH-0013 | The Web UI and the manual say that Ookla and RTR-NetTest results are also stored by those services (they appear in their own and in international statistics), so the user knows that such a result leaves the box before choosing those engines. | docs/NMHH_SZELESSAVNET_COMPARISON.md@ad3758c:77 | P-12 | M-08, M-10, U-02…U-0n (page: Measurement and servers) | |

## Coverage

Each range includes the blank line(s) that end it. Every line of every source is in exactly one row of its table.

### docs/NMHH_SZELESSAVNET_COMPARISON.md (lines 1-95)

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-2 | heading | context: title |
| 3-4 | paragraph | REQ-NH-0001, REQ-NH-0004, REQ-NH-0007 |
| 5-6 | paragraph | REQ-NH-0006 |
| 7-9 | bullets | REQ-NH-0003, REQ-NH-0004, REQ-NH-0006 |
| 10-13 | bullets | REQ-NH-0006; the owner's experience and screenshot (source names a specific site value; not copied) |
| 14-15 | paragraph | REQ-NH-0002, REQ-NH-0003 |
| 16-17 | rule | context: separator |
| 18-19 | heading | context: section title |
| 20 | heading | context: section title |
| 21-23 | bullets | context: terms of the official tool (only in its country, only for residents); a country file links them, REQ-NH-0003 |
| 24 | heading | context: section title |
| 25-29 | bullets | context: eligibility and acceptance terms of the official tool; linked, not copied (REQ-NH-0003) |
| 30 | heading | context: section title |
| 31-36 | bullets | context: contract duties of the official tool (ownership, set-up, liability, return, notice); line 35's three-month condition is a reason to end the contract, not a quality threshold for NMB |
| 37 | heading | context: section title |
| 38 | bullet | REQ-NH-0012; the contrast (several engines and servers) is the README rows of M-04 |
| 39 | bullet | context: retention of the official system; NMB's own retention is C-02's |
| 40 | bullet | REQ-NH-0012; the server address and the example result (source names a specific site value; not copied) |
| 41-42 | bullet | context: lower precision of the official box's own tests (one transfer) |
| 43 | heading | context: section title |
| 44 | bullet | context: contract term (the software must not be touched) |
| 45-46 | bullet | context: closed source of the official box; the native build is open (REQ-WP-0188) |
| 47 | heading | context: section title |
| 48 | bullet | REQ-NH-0010 |
| 49 | bullet | context: the official box's basic web page (owner's experience) |
| 50-51 | bullet | context: about 1 Gbit/s and no link aggregation on the official box; NMB's speed claims are README and TESTING rows |
| 52-53 | heading | context: section title |
| 54 | bullet | context: hardware of the official box (screenshot values not copied) |
| 55 | bullet | context: tabs of the official box's local page; NMB's pages are the U series |
| 56 | bullet | context: stability (owner's experience) |
| 57 | bullet | REQ-NH-0002, REQ-NH-0004 |
| 58 | bullet | REQ-NH-0011 |
| 59 | bullet | context: the official box measures when the line is idle; NMB's traffic gate is REQ-RE-0087 and REQ-RE-0088 |
| 60 | bullet | context: LAN passthrough of the official box; NMB's is N-06 (the English "access point" is a mistranslation, see the Hungarian table, line 66) |
| 61 | bullet | context: who built the official system and where its data go |
| 62 | bullet | REQ-NH-0009 |
| 63 | bullet | REQ-NH-0008 |
| 64-65 | bullet | context: mobile measurement of the official system; the box itself is for wired lines |
| 66 | heading | context: section title |
| 67-68 | bullet | context: not tried on the official box (PPPoE, router functions) |
| 69-70 | heading | context: section title |
| 71-72 | table header | context: table header |
| 73 | table row | context: ownership (own device, no contract) |
| 74 | table row | context: engines and the mixed run, README rows of M-04 |
| 75 | table row | context: servers, README rows of M-06 to M-10 |
| 76 | table row | context: per-engine intervals and the Ookla limit, REQ-RE-0073 and REQ-RE-0095 |
| 77 | table row | REQ-NH-0013 |
| 78 | table row | REQ-NH-0001 |
| 79 | table row | context: modes and their test status, README and TESTING rows |
| 80 | table row | context: mobile internet over USB tethering, N-10 rows |
| 81 | table row | context: Wi-Fi client and access point, N-07 rows |
| 82 | table row | context: speed above 1 Gbit/s planned, link aggregation in the software, N-06 rows |
| 83 | table row | context: Web UI, Telegram and CLI |
| 84 | table row | context: open source; the native licence follows from its components (REQ-WP-0193) |
| 85-86 | table row | context: claim that open code and plain-text settings help working with an LLM; no box behaviour |
| 87-88 | heading | context: section title |
| 89 | bullet | REQ-NH-0005 |
| 90 | bullet | REQ-NH-0005 |
| 91-92 | bullet | REQ-NH-0002, REQ-NH-0005 |
| 93-94 | rule | context: separator |
| 95 | paragraph | REQ-NH-0003 |

### docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md (lines 1-102)

The Hungarian file translates the English one with the same sections and bullets; it has a blank line after every
sub-heading, so its line numbers run ahead (English line n is about Hungarian line n + 2 to n + 7).

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-2 | heading | context: translation of COMPARISON.md:1 |
| 3-4 | paragraph | context: translation of COMPARISON.md:3; says the purpose more fully ("so that whoever wants to measure their own line sees what they get with each"), same intent; links no English file |
| 5-22 | paragraph, bullets, headings | context: translation of COMPARISON.md:5-21; line 12 cites the Hungarian README and testing record |
| 23-24 | bullet | context: Hungarian only: names the contract section the residence condition is in, a more exact citation of the kind REQ-NH-0003 links to |
| 25-27 | heading, bullet | context: translation of COMPARISON.md:24-25 |
| 28 | bullet | context: Hungarian only: quotes the acceptance rule of annex 1 point 3 word for word and gives the aim ("for a representative sample") |
| 29 | bullet | context: translation of COMPARISON.md:27 |
| 30-31 | bullet | context: Hungarian only: "for now they focus on packages below 1000 Mb/s" (the English says this only at line 50) |
| 32-37 | heading, bullets | context: translation of COMPARISON.md:30-34 |
| 38-39 | bullet | context: Hungarian only: the condition reads "download and upload reach 90 % of the guaranteed bandwidth", and a closing inference ("so if your connection is good, they may ask for the box back"); contract terms, no NMB behaviour |
| 40-44 | heading, bullets | context: translation of COMPARISON.md:37-40 (server address and example result not copied) |
| 45-46 | bullet | context: Hungarian only: the participant's own test runs with one average-size data block "against the central measurement with several blocks", a detail of the official method |
| 47-53 | heading, bullets | context: translation of COMPARISON.md:43-47 |
| 54 | bullet | REQ-NH-0010 (Hungarian only: "especially when the subscription runs on another company's network"; provider names not copied) |
| 55 | bullet | context: translation of COMPARISON.md:49 |
| 56-57 | bullet | context: translation of COMPARISON.md:50, adding the term "LAN aggregation" |
| 58-64 | heading, bullets | context: translation of COMPARISON.md:52-58 |
| 65 | bullet | context: Hungarian only: the official box "checks the traffic before measuring" and measures when the user is not using the internet; NMB's traffic gate is REQ-RE-0087 |
| 66 | bullet | context: translation difference: "kiépítési végpont" (the line's termination point) where COMPARISON.md:60 says "access point"; the Hungarian reading is the right one (the official box sits between the line's termination and the user's devices) |
| 67-87 | bullets, headings, table | context: translation of COMPARISON.md:61-81 |
| 88 | table row | context: Hungarian only: "access point and client both tested live on 2.4 and 5 GHz", a claim about the Python version's tests (TESTING.md rows), not a native requirement |
| 89 | table row | context: translation of COMPARISON.md:82 |
| 90 | table row | context: translation of COMPARISON.md:83, adding "not only a web interface" |
| 91-102 | table rows, headings, bullets | context: translation of COMPARISON.md:84-95 |

### docs/DEV_SETUP_WINDOWS.hu.md (lines 1-220)

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-220 | whole file | context: line-by-line translation of `docs/DEV_SETUP_WINDOWS.md` (REQ-DV rows), compared line by line; no content of its own and no contradiction. Lines 92-94 split English line 92 into two bullets and lines 179-183 English line 177 into four, which moves the line numbers by up to 6; the links of lines 3, 190 and 207 point to the Hungarian files; placeholders are translated (`<doboz-tailscale-ip>`, `<sudo-jelszó>`) |

### docs/TESTING.hu.md (lines 1-399)

From line 7 on, Hungarian line n translates English line n + 2 (English line 7 has no counterpart, REQ-TS-0002), except
the Measurement rows Cloudflare and Fast.com, which are in the other order.

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-6 | heading, paragraphs | context: translation of TESTING.md:1-6; line 3 adds the link to the English file |
| 7-48 | paragraphs, tables | context: translation of TESTING.md:9-50 (line 15 has a doubled full stop) |
| 49 | table row | context: translation of TESTING.md:52; says that the kernel dropped the other DHCP server's replies (the English leaves out by what) |
| 50-52 | table rows | context: translation of TESTING.md:53, 54 and 51 |
| 53-164 | tables, headings | context: translation of TESTING.md:55-166 |
| 165 | table row | context: translation of TESTING.md:167 with one contradiction: the bot ran code "days older" where the English says "week-old code"; TESTING.md:184 (Hungarian line 182) says two days, so the English figure is the outlier; no requirement depends on it (REQ-TS-0133) |
| 166-290 | tables, headings, lists | context: translation of TESTING.md:168-292; line 221 says the USB profile never gets the default route, clearer than the English "profile never default" |
| 291 | table row | context: translation of TESTING.md:293 (checklist 10) with a mistranslation: "sáv" (band) for "banner" |
| 292-314 | tables, headings | context: translation of TESTING.md:294-316 |
| 315 | table row | context: translation of TESTING.md:317 with the same mistranslation (band for banner) |
| 316-399 | tables, headings, lists | context: translation of TESTING.md:318-401 |

## Content only in the Hungarian version

Every statement of a Hungarian file that its English source does not make, with what it became. Translation slips that
change no statement (a doubled full stop, a translated placeholder) are only in the coverage tables.

| File | Line | Statement only in Hungarian | Entry |
| :--- | :--- | :--- | :--- |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 3 | The page's aim: whoever wants to measure their own line sees what they get with each box. | context: same intent as English line 3 |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 23 | The residence condition is cited with its contract section. | context: more exact citation; REQ-NH-0003 links the terms |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 28 | The acceptance rule quoted word for word (annex 1 point 3), with the aim of a representative sample. | context: terms of the official tool, linked not copied (REQ-NH-0003) |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 30 | The programme focuses on packages below 1000 Mb/s for now. | context: the English says it at line 50 only |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 38 | The contract may be ended when download and upload reach 90 % of the guaranteed bandwidth; "so if your connection is good, they may ask for the box back". | context: contract term and an inference; no NMB behaviour |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 45 | The participant's own test uses one data block, the central measurement several. | context: method of the official tool |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 54 | A provider mismatch is likely "especially when the subscription runs on another company's network". | REQ-NH-0010 |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 56 | The term "LAN aggregation". | context: wording |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 65 | The official box checks the traffic before measuring. | context: NMB's traffic gate, REQ-RE-0087 |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 66 | The official box sits between the line's termination point and the user's devices (the English says "access point"). | context: the Hungarian is right; the English line 60 is a mistranslation of the Python page |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 88 | NMB's Wi-Fi access point and client were both tested live on 2.4 and 5 GHz. | context: claim about the Python version's tests |
| `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md` | 90 | NMB is controlled "not only through a web interface". | context: wording |
| `CONTRIBUTING.hu.md` | - | none | - |
| `docs/DEV_SETUP_WINDOWS.hu.md` | - | none (the structure differs at lines 92-94 and 179-183, the content does not) | - |
| `docs/TESTING.hu.md` | 49 | It was the kernel that dropped the other DHCP server's replies. | context: more precise wording of TESTING.md:52 |
| `docs/TESTING.hu.md` | 165 | The bot ran code "days older" (the English: "week-old"). | context: contradiction; TESTING.md:184 says two days, the English figure is the outlier |
| `docs/TESTING.hu.md` | 221 | The USB profile is never the default route. | context: clearer wording of TESTING.md:223 |
| `docs/TESTING.hu.md` | 291, 315 | "band" where the English says "banner" (device explorer). | context: mistranslation; the English is right |
| `README.hu.md` | 47 | NMB is a "professional" target software that "ends subjective disputes" with ISPs. | context: claim; README.md:22 and REQ-RE-0223 state what NMB is |
| `README.hu.md` | 90 | RTR-NetTest, Cloudflare and iPerf3 "run every time". | context: stale; the per-engine intervals of README.md:133 (REQ-RE-0095) win; added to `readme.md` |
| `README.hu.md` | 122 | The target device is "SBC, x86 mini PC, embedded router". | context: contradiction; router ports are planned only (REQ-RE-0002); added to `readme.md` |
| `README.hu.md` | 187 | A link points to a source file under a developer's home folder (`file://` path). | REQ-RE-0341 (the path is not copied); added to `readme.md` |
| `README.hu.md` | 268 | The `low` profile is "not tested on 512 MB". | context: stale like line 98; tested in a VM (REQ-RE-0209); noted in `readme.md` |
| `README.hu.md` | 420 | Flash wear is checked, besides free space (`statvfs()`). | context: claim; flash wear tracking is already REQ-RM-0128 (C-08); noted in `readme.md` |
| `README.hu.md` | 591 | `/vpn add` "returns the settings" (the English: "returns the QR code"). | context: same command; REQ-RE-0261 (the English) wins; added to `readme.md` |

The other Hungarian-only lines of `README.hu.md` are those that `readme.md` already lists (its Hungarian table:
REQ-RE-0312 to REQ-RE-0338 and the context rows there).

## Proposed package changes

None. Every row has a package of work plan section 3; REQ-NH-0009 moves to `NEW: M-17 Shared HTTP transfer core` if
that package is accepted (as REQ-RE-0193 in `readme.md`).

## Questions for the owner

None. The disclaimer (REQ-NH-0002) follows the owner's decision of 2026-11-11; everything country-specific follows
Quality Bar 9.
