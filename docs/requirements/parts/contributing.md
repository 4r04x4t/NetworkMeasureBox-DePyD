# Requirements register part: CONTRIBUTING.md (compared with CONTRIBUTING.hu.md)

Source: `CONTRIBUTING.md` of the Python repository at commit `ad3758c` (55 lines), read in full, and `CONTRIBUTING.hu.md`
(55 lines), compared line by line. Added in round 1 of the checks: the first merge left both files out (`python-files.md`
skipped them as "covered by their own parts", but no part covered them).
The file gives the rules for contributing to the Python version. Each rule is carried over as what the native build, or the
native repository's own contribution rules, must do; rules that only concern Python tooling say so in Notes.
Package ids are those of `docs/NATIVE_WORK_PLAN.md` section 3 and 5.4 at `ad3758c`; `NEW:` ids are those of
`PACKAGE_CHANGES.md` section B. Ids REQ-CT-0001 to REQ-CT-0026. Ids are permanent.

## Requirements

| Id | Requirement | Source | Package | Also in | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| REQ-CT-0001 | The native repository's contribution rules ask for an issue first, saying what is to change and why, for anything bigger than a small fix. | CONTRIBUTING.md@ad3758c:7 | F-00 | P-12 | The native repository is open (REQ-WP-0188); its own contribution file carries the rules of this part. |
| REQ-CT-0002 | A test run on other hardware, another line or another distribution records exactly what it ran on (line type, board, distribution). | CONTRIBUTING.md@ad3758c:10 | P-13 | F-06 | (source names a specific site value; not copied) Same rule as REQ-TS-0004. |
| REQ-CT-0003 | Every call of a system tool has full error handling, and the program keeps working, with the reason shown, when the tool is missing. | CONTRIBUTING.md@ad3758c:14 | F-04 | every package, N-17, P-06 | The Python half of the line (Python 3.10, standard library only, no pip packages) is not carried over; the native language rule is the work plan's (docs/NATIVE_WORK_PLAN.md@ad3758c:19-20). The native build calls fewer tools (roadmap section 7); the rule holds for those it keeps. |
| REQ-CT-0004 | Honest results: no result is stored under another engine's name, no value is filled in that was not measured, no physically impossible result is kept, and a failed run is stored as failed with its reason. | CONTRIBUTING.md@ad3758c:15 | M-04 | M-06, M-07, M-08, M-09, M-10, R-01, C-02, F-03 | |
| REQ-CT-0005 | Anything that hands out addresses, changes routes or sends traffic first checks that this is safe on the network the box is on (for example it never serves DHCP where another server answers) and undoes itself when it is switched off. | CONTRIBUTING.md@ad3758c:16 | N-03 | N-01, N-05, N-08, N-11, every package | The DHCP example is the Wi-Fi sharing guard (N-05, N-08). |
| REQ-CT-0006 | The limits of the services the box measures against are kept (Ookla at most once an hour). | CONTRIBUTING.md@ad3758c:17 | M-04 | M-10, M-05 | The per-engine minimum intervals (README.md@ad3758c:133). |
| REQ-CT-0007 | The box identifies itself honestly to the services it measures against (its own client name and version, no borrowed identity). | CONTRIBUTING.md@ad3758c:17 | C-05 | M-06, M-07, M-08, M-09, M-10, M-12, M-13 | C-05 holds the user agent. |
| REQ-CT-0008 | Every new setting has a default that is safe when the user never touches it, is read and written by the configuration code, and is documented. | CONTRIBUTING.md@ad3758c:18 | C-01 | P-12, every package | Python names `load_config` and `save_config`. |
| REQ-CT-0009 | Every new visible text goes through a translation key that exists in every language; no Hungarian text appears in another language. | CONTRIBUTING.md@ad3758c:19 | C-06 | U-01, every package | Python names `t('key', 'English fallback')`, `nmb/web/i18n.py` and `nmb/web/locales/`; the native build check is REQ-RM-0047. |
| REQ-CT-0010 | Every endpoint the Web UI calls is also an action of `nmb-cli`. | CONTRIBUTING.md@ad3758c:19 | S-10 | S-02…S-0n, U-02…U-0n | Python catalogue `nmb/tui/catalog.py`. |
| REQ-CT-0011 | Changes are small and readable: one topic per change, and comments say why, not what. | CONTRIBUTING.md@ad3758c:20 | F-00 | every package | The gate's review applies it. |
| REQ-CT-0012 | Every change says how it was tested; a change without a test record is not merged. | CONTRIBUTING.md@ad3758c:24 | F-00 | every package | The gate (skill `nmb-native-gate`). |
| REQ-CT-0013 | Unit tests run on every change; new behaviour gets new tests; unit tests fake the system and run without root, network or a box; a fake returns what the real function returns, never an invented result. | CONTRIBUTING.md@ad3758c:26-32 | F-06 | every package, F-00 | The Python commands (`unittest`, `node --check`, `bash -n`) and `autospec=True` are not carried over; the native commands are F-00's gate. |
| REQ-CT-0014 | A change that touches the network, the system or measuring is also tested on a real device installed the way users get it: the feature runs for real, its claim is measured (for network features the interface byte counters before and after), a veth pair and a network namespace are used instead of the port the tester is connected through, and the device is put back as it was. | CONTRIBUTING.md@ad3758c:33 | F-06 | P-13, every package | Python installs with `git archive` and `setup.sh`; the native way is P-11's packages. On the owner's Pi 4 the conditions of REQ-WP-0032 apply. |
| REQ-CT-0015 | Installer and distribution changes are tested in a VM: clean install, install again over it, reboot. | CONTRIBUTING.md@ad3758c:34 | P-11 | F-06 | |
| REQ-CT-0016 | Every test adds a row to the test record: where, the result, how. | CONTRIBUTING.md@ad3758c:35 | P-13 | every package | Python keeps the record in English and Hungarian; the language rule for native documents is open (QUESTIONS.md Q-09). Same rule as REQ-TS-0214. |
| REQ-CT-0017 | While testing, no token, password or key is read, printed or copied; only whether it is set is checked. | CONTRIBUTING.md@ad3758c:37 | F-06 | C-04, every package | |
| REQ-CT-0018 | After an install, every service is checked to have restarted (its start time later than the start of the install), not only to be running. | CONTRIBUTING.md@ad3758c:37 | P-11 | P-02, P-13 | Same rule as REQ-TS-0168; Python checks `ActiveEnterTimestamp`. |
| REQ-CT-0019 | The test record says what was not tested; "works" is never written without saying where. | CONTRIBUTING.md@ad3758c:39 | P-13 | every package | Same rule as REQ-TS-0214. |
| REQ-CT-0020 | The native repository has a pull request template that asks for the test record. | CONTRIBUTING.md@ad3758c:43 | F-00 | | Python template `.github/pull_request_template.md` (python-files.md). |
| REQ-CT-0021 | Every change a user can notice is entered in the changelog under *Unreleased* (Keep a Changelog: Added, Changed, Fixed, Removed). | CONTRIBUTING.md@ad3758c:44 | F-00 | every package | Same rule as REQ-RM-1230. |
| REQ-CT-0022 | When behaviour or settings change, the user documentation is updated in the same change. | CONTRIBUTING.md@ad3758c:45 | P-12 | every package | Python: the README in both languages; the native language rule is open (Q-09). Same rule as REQ-RM-1395. |
| REQ-CT-0023 | Commit messages are written `type(scope): what changed` (`feat`, `fix`, `docs`, `test`, `refactor`), with the reason in the body when it is not obvious. | CONTRIBUTING.md@ad3758c:46 | F-00 | | |
| REQ-CT-0024 | A change written partly by an AI assistant says so, and the person who submits it stays responsible for having read and tested it. | CONTRIBUTING.md@ad3758c:47 | F-00 | every package | The native agents' commits carry a co-author line. |
| REQ-CT-0025 | Security problems are reported privately to the maintainer, with the steps to reproduce, never in a public issue; the native repository says how. | CONTRIBUTING.md@ad3758c:51 | F-00 | F-03 | Same handling as REQ-RM-1025 (details private until fixed). |
| REQ-CT-0026 | The native repository states under which licence contributions are published; the Python version's MIT licence is not assumed for it. | CONTRIBUTING.md@ad3758c:55 | P-08 | F-00, F-01 | The native licence follows from its components (REQ-WP-0193); until it is fixed the repository says it is not decided (REQ-WP-0195). |

## Coverage

Each range includes the blank line(s) that end it. Every line of both files is in exactly one row of its table.

### CONTRIBUTING.md (lines 1-55)

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-2 | heading | context: title |
| 3-4 | paragraph | context: purpose (a wrong number is worse than a missing feature); the rule itself is line 15 (REQ-CT-0004); link to the Hungarian file |
| 5-6 | heading | context: section title |
| 7 | bullet | REQ-CT-0001 |
| 8 | bullet | context: pointer to `docs/TESTING.md` (its own part, `testing.md`) |
| 9 | bullet | context: pointer to `docs/DEV_SETUP_WINDOWS.md` (REQ-DV rows) |
| 10-11 | bullet | REQ-CT-0002 |
| 12-13 | heading | context: section title |
| 14 | bullet | REQ-CT-0003 |
| 15 | bullet | REQ-CT-0004 |
| 16 | bullet | REQ-CT-0005 |
| 17 | bullet | REQ-CT-0006, REQ-CT-0007 |
| 18 | bullet | REQ-CT-0008 |
| 19 | bullet | REQ-CT-0009, REQ-CT-0010 |
| 20-21 | bullet | REQ-CT-0011 |
| 22-23 | heading | context: section title |
| 24-25 | paragraph | REQ-CT-0012 |
| 26-32 | numbered item with code block | REQ-CT-0013 |
| 33 | numbered item | REQ-CT-0014 |
| 34 | numbered item | REQ-CT-0015 |
| 35-36 | numbered item | REQ-CT-0016 |
| 37-38 | paragraph | REQ-CT-0017, REQ-CT-0018 |
| 39-40 | paragraph | REQ-CT-0019 |
| 41-42 | heading | context: section title |
| 43 | bullet | REQ-CT-0020 |
| 44 | bullet | REQ-CT-0021 |
| 45 | bullet | REQ-CT-0022 |
| 46 | bullet | REQ-CT-0023 |
| 47-48 | bullet | REQ-CT-0024 |
| 49-50 | heading | context: section title |
| 51-52 | paragraph | REQ-CT-0025 |
| 53-54 | heading | context: section title |
| 55 | paragraph | REQ-CT-0026 |

### CONTRIBUTING.hu.md (lines 1-55)

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-55 | whole file | context: line-by-line translation of `CONTRIBUTING.md`, compared line by line; no content of its own and no contradiction. Only the links differ: line 8 points to `docs/TESTING.hu.md` and line 9 to `docs/DEV_SETUP_WINDOWS.hu.md` |

## Proposed package changes

None. Every row has a package of work plan section 3.

## Questions for the owner

None new. The language of the native repository's contribution rules and test record follows the answer to Q-09
(documentation languages) in `QUESTIONS.md`.
