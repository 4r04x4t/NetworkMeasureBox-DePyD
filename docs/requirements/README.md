# Requirements register (package F-02)

Every requirement the native build of NetworkMeasureBox must meet, as one numbered list. Each requirement is one
checkable statement of what the native build must do, must not do, or must prove, with the source line it comes from
and the one package that owns it. Text with no requirement is listed as context, so every line of every source is
accounted for. Package F-02 of `docs/NATIVE_WORK_PLAN.md@ad3758c:108` and `:217-218`.

## Sources

All sources are files of the Python repository (`NetworkMeasureBox`) at commit **`ad3758c`**, read only. Every source
cell is written `<file>@ad3758c:<line>` or `<file>@ad3758c:<line>-<line>`. The line numbers belong to that commit and stay
valid when the Python files change later.

| Prefix | Source | Part | Ids |
| :--- | :--- | :--- | ---: |
| REQ-RM | `docs/DEPYTHONIZATION_ROADMAP.md` lines 1-507 (ids 0001-0219) and 508-1034 (ids 1001-1465) | `parts/roadmap-1-507.md`, `parts/roadmap-508-1034.md` | 684 |
| REQ-PP | `docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md` (porting plan) | `parts/porting-plan.md` | 198 |
| REQ-WP | `docs/NATIVE_WORK_PLAN.md` | `parts/workplan-architecture-devsetup.md` | 197 |
| REQ-AR | `docs/ARCHITECTURE.md` | `parts/workplan-architecture-devsetup.md` | 78 |
| REQ-DV | `docs/DEV_SETUP_WINDOWS.md` | `parts/workplan-architecture-devsetup.md` | 26 |
| REQ-TS | `docs/TESTING.md` (compared with `docs/TESTING.hu.md`) | `parts/testing.md` | 314 |
| REQ-RE | `README.md` and the content only in `README.hu.md` | `parts/readme.md` | 343 |
| REQ-CL | `CHANGELOG.md` | `parts/changelog.md` | 453 |
| REQ-CT | `CONTRIBUTING.md` (compared with `CONTRIBUTING.hu.md`) | `parts/contributing.md` | 26 |
| REQ-NH | `docs/NMHH_SZELESSAVNET_COMPARISON.md` (comparison with one country's official measuring box; QB 9: requirements of the country data files and the neutral default) | `parts/other-docs.md` | 13 |
| (none) | the Hungarian translations `docs/NMHH_SZELESSAVNET_OSSZEHASONLITAS.hu.md`, `docs/DEV_SETUP_WINDOWS.hu.md` and `docs/TESTING.hu.md`, each compared line by line with its English source (content only in Hungarian: a section of that part) | `parts/other-docs.md` | - |
| (none) | every other tracked file (code, tests, units, assets), file by file with line ranges | `parts/python-files.md` | - |
| | | **Total** | **2332** |

Every document of the commit (top-level `*.md` and `docs/`) is the source of a coverage table; `OUT_OF_SCOPE` in
`check.sh` is empty since round 3, when `parts/other-docs.md` covered the last four. A Hungarian translation is never
the source of a requirement on its own: where it says something its English source does not, the part of that source
(or `other-docs.md`) lists the statement and turns it into a requirement or a context entry.

`parts/python-files.md` has no requirement ids: it gives each line of each Python file to one package, and the analyst
of that package writes the requirements of its ranges in the brief (skill `nmb-native-brief`).

## Two rules of the register (decided by the orchestrator, round 3)

- **Exactly one package.** The work plan's "exactly one package" rule (`docs/NATIVE_WORK_PLAN.md`, F-02) is read as:
  every requirement, every Python file line range and every `TESTING.md` row has exactly one owning package. A roadmap
  section or a Python file may be split by line range between packages; each range then has its one owner.
- **Python files are mapped by structure.** `parts/python-files.md` maps the Python files to packages by their
  structure (module, top-level functions and classes, test classes); F-02 does not read the 2.5 MB of Python line by
  line. Each package's analyst reads its own files in full (skill `nmb-native-brief`) and asks the orchestrator to move
  a file or a range when the mapping is wrong.

## Files

| File | What it is |
| :--- | :--- |
| `parts/*.md` | The register itself. Each part has a *Requirements* table (Id, Requirement, Source, Package, Also in, Notes), a *Coverage* table (every source line in order: the ids it holds, or `context: <why no requirement>`), its proposed package changes and its questions. |
| `BY_PACKAGE.md` | For every package: the ids it owns and the ids it is also in. Generated from the parts; `check.sh` fails when they disagree. |
| `PACKAGE_CHANGES.md` | Proposed new, split, merged and moved packages, de-duplicated; and the mappings of every roadmap section (M1), every `TESTING.md` row (M2) and every Python file (M3) to one package. |
| `QUESTIONS.md` | The questions of all parts, de-duplicated and triaged (round 3): answered by existing decisions, developer's choices with their rule, and the few left for the owner. |
| `check.sh` | The register's check (POSIX `sh` and `awk`). |
| `check_test.sh` | Test of the check: the real register must pass, 30 planted faults in copies must each fail. |

## Names and site data (Quality Bar 8)

The parts copy no personal or site data: where a source names the owner's line, ISP, plan, test devices, places, accounts or
a measured site value, the row states the requirement generically and adds "(source names a specific site value; not
copied)". Public product targets that the plans themselves name as targets (for example the TL-WR1043ND v1, the first
router target of `docs/NATIVE_WORK_PLAN.md` package I-10, and the Raspberry Pi 4) and public protocol or vendor
identifiers (USB vendor IDs, socket paths, server names of public services) are copied, because a rule cannot be checked
without them. This rule was made one for all parts in round 1.

## Ids are permanent

- An id is given once and never reused, renumbered or given to another requirement, also when its row is withdrawn.
- A row that is wrong is corrected in place (text, source, package); a row that is no longer wanted keeps its id and
  says *withdrawn* with the reason in Notes: the Requirement cell starts with `*Withdrawn.*`, the Notes cell with
  `withdrawn: <reason>` (for example `withdrawn: superseded by REQ-WP-0179, -0180`). New requirements of a source get the
  next free number of its prefix (REQ-RM-0220 for roadmap lines 1-507, REQ-RM-1466 for 508-1034, REQ-PP-0199, REQ-WP-0198,
  REQ-RE-0344, REQ-CL-0454, REQ-CT-0027, REQ-NH-0014, and so on).
- Requirements that a package analyst finds in a Python file are written in the brief with the file's line range as their
  source; if the register should hold them, they get ids of a new prefix agreed with F-02 (not one of the ten above).
- Proposed package ids (section B of `PACKAGE_CHANGES.md`) follow the same rule: a rejected id is not reused.

## Owner and Also in

- **Package** is the one package that must build the requirement and prove it with a test (work plan 2.4: a requirement
  id of a package without a test at the end is a failed gate). A withdrawn row keeps its package for the trace; the brief
  lists it as withdrawn and the gate expects no test for it (`BY_PACKAGE.md`, *Flags*, lists them).
- A **process rule** is a rule for how the work is done (roles, the development PC, a rented server, tokens and passwords,
  how images are handled) that no test of the package's code can prove. Its Notes cell starts with `process rule:` and
  names who checks it: the orchestrator, or a gate script where one is possible. The owning package lists it in its brief
  and the gate expects no test for it (`BY_PACKAGE.md`, *Flags*, lists them). A brief that finds another owned row of this
  kind reports it to F-02, which adds the flag (added in round 2).
- Rows that concern only the C code base or only router targets (step 2, `docs/NATIVE_WORK_PLAN.md@ad3758c:19-22`) are
  owned by `NEW: X2-01`, with the step-1 package under *Also in*; a row that mixes both is split.
- **Also in** lists packages that must keep to the requirement, show or store what it produces, or test their side of it;
  `every package` or `all packages` means a rule for all of them (listed under `ALL` in `BY_PACKAGE.md`).
- The open series are written `S-02…S-0n (API: <area>)` and `U-02…U-0n (page: <page>)` until their numbers are fixed
  (`PACKAGE_CHANGES.md` section A). Proposed packages are written `NEW: <id> <name>`.

## How a package brief cites the register

The brief of a package (skill `nmb-native-brief`) starts with its trace table:

1. Take the package's section in `BY_PACKAGE.md`: every id under **Owns** must appear in the brief with the test that
   will prove it; every id under **Also in** is listed with what the package does for it.
2. Cite each id with its source, as the part gives it, for example:
   `REQ-RM-0005 (docs/DEPYTHONIZATION_ROADMAP.md@ad3758c:16): probe timing ... - test: <name>`.
3. Add the package's ranges from `parts/python-files.md` (M3 of `PACKAGE_CHANGES.md` gives each file's main package)
   and the roadmap sections and `TESTING.md` rows mapped to it (M1, M2). Most ranges were assigned from the files'
   structure, not read line by line (that part's *Not done as asked* note): the analyst confirms every range boundary
   and Notes cell of its files while reading them, and reports each moved boundary to F-02.
4. If the analyst finds a row that belongs to another package, or a requirement the register lacks, the brief says so
   and F-02 changes the part (and regenerates `BY_PACKAGE.md`); the brief never renumbers an id.

The gate (skill `nmb-native-gate`) checks that every owned id has a test.

## Running the check

From the worktree root:

```sh
sh docs/requirements/check.sh          # exit 0: register consistent; 1: a fault (each printed as FAIL); 2: cannot read
sh docs/requirements/check_test.sh     # exit 0: the check passes the real register and fails every planted fault
```

`NMB_SRC` points to the Python repository (default `$HOME/src/NetworkMeasureBox`); the check reads commit `ad3758c` from it
with `git show`, `git grep` and `git ls-tree`, so it does not depend on the working copy. It checks: ids defined once and
well formed; six cells per requirement row; a package on every row, known to `BY_PACKAGE.md`; every coverage table without
gap or overlap from the first to the last line of its source range (the Hungarian README table: in order, no overlap);
every id in a coverage row and every id of a coverage row defined; every line a requirement cites from its own part's
source in a coverage row that names the id (added in round 1); every Python file of the commit (without `docs/` and the
top-level `*.md`) in `parts/python-files.md` with each line exactly once; `BY_PACKAGE.md` equal to the parts; every work
plan package present; and M1-M3 of `PACKAGE_CHANGES.md` mapping each roadmap heading, `TESTING.md` row and Python file
exactly once to a known package, with each M2 row's owners column and package recomputed from `parts/testing.md` and each M1 row's
*Rows by owner* from the roadmap parts (both added in round 2); and every document of the commit (top-level `*.md`, `docs/`) either the source of a
coverage table or named in `OUT_OF_SCOPE` (empty since round 3).

`BY_PACKAGE.md` and sections M1-M3 of `PACKAGE_CHANGES.md` were generated by the F-02 merge step from the parts. The
generator is not in this repository yet (it was a one-off script of the merge, and again of rounds 1, 2 and 4; a POSIX or Go version
belongs with F-00's tools). Until then, after a change to a part, edit `BY_PACKAGE.md` and M1-M3 to match; `check.sh` names every place that
disagrees.
