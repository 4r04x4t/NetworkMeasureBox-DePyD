# Requirements register, part: work plan, architecture, Windows development setup

Sources (Python repository, commit `ad3758c`, read only):

- `docs/NATIVE_WORK_PLAN.md` (343 lines), read in pages 1-150, 151-300, 301-343 (all lines).
- `docs/ARCHITECTURE.md` (149 lines), read in one page 1-149 (all lines).
- `docs/DEV_SETUP_WINDOWS.md` (214 lines), read in pages 1-150, 151-214 (all lines).

Line numbers were checked with `awk` against the files at `ad3758c`. Package ids are those of `docs/NATIVE_WORK_PLAN.md@ad3758c:98-219` (section 3). `NEW: X2-01 Step-2 router tiers` is the package proposed in the roadmap part (`roadmap-1-507.md`), reused here so the register has one id for it.

Id ranges of this part: REQ-WP-0001 to REQ-WP-0197 (0196 and 0197 added in round 2), REQ-AR-0001 to REQ-AR-0080 (0079 and 0080 from the owner's answer to Q-15 of 2026-11-11), REQ-DV-0001 to REQ-DV-0026.

Many work-plan rows are rules for how the native work is done (roles, gate, status file). They are listed as requirements because the gate of section 2.4 checks them; their owner is F-00 (which builds the gate scripts and `WORK_STATUS.md`) unless a more specific package exists. The ARCHITECTURE rows describe the Python version; as requirements they state what the native build must keep, and where the source is out of date or faulty the Notes say so (such cases go to F-03, never into code).

## Requirements

| Id | Requirement | Source | Package | Also in | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| REQ-WP-0001 | The roadmap and the porting plan are the source of the requirements; where the work plan differs from them, they win and the work plan is corrected. | docs/NATIVE_WORK_PLAN.md@ad3758c:6-7 | F-02 | F-01 | The register flags each such conflict instead of choosing silently. |
| REQ-WP-0002 | The project is split into packages small enough for small agents. | docs/NATIVE_WORK_PLAN.md@ad3758c:9 | F-02 | all packages | Size rule: REQ-WP-0055. |
| REQ-WP-0003 | Every document is read with care and nothing is left out. | docs/NATIVE_WORK_PLAN.md@ad3758c:10 | F-02 | F-00 | Proved by the register's coverage tables and each brief's reading table (REQ-WP-0020). |
| REQ-WP-0004 | `linux/amd64` (PC) and `linux/arm64` (Raspberry Pi 4) are built and tested on every change. | docs/NATIVE_WORK_PLAN.md@ad3758c:17 | F-00 | P-13 | |
| REQ-WP-0005 | `linux/arm` v7 is also built and run under QEMU. | docs/NATIVE_WORK_PLAN.md@ad3758c:18 | F-00 | F-06 | |
| REQ-WP-0006 | `linux/mips` and `linux/mipsle` are cross-compiled only in step 1 (roadmap *Generic First* point 6). | docs/NATIVE_WORK_PLAN.md@ad3758c:18 | F-00 | NEW: X2-01 Step-2 router tiers | DEV_SETUP says they can also run under `qemu-*-static` (REQ-DV-0010); running them is allowed, not required. |
| REQ-WP-0007 | Step 1 is written in Go without cgo; C11/musl is for the router tiers and starts with step 2. | docs/NATIVE_WORK_PLAN.md@ad3758c:19-20 | F-00 | NEW: X2-01 Step-2 router tiers | |
| REQ-WP-0008 | Step 1 is done only when the whole native program (every component of roadmap section 5, not a subset) passes Quality Bar points 1 to 9 on the step-1 targets. | docs/NATIVE_WORK_PLAN.md@ad3758c:21-22 | P-13 | all packages | |
| REQ-WP-0009 | Nothing router-specific is written before step 1 is done. | docs/NATIVE_WORK_PLAN.md@ad3758c:22 | F-00 | NEW: X2-01 Step-2 router tiers | process rule: the orchestrator, when it picks packages. |
| REQ-WP-0010 | Every native file lives in `NetworkMeasureBox-DePyD/` (own repository, without exception), next to the Python folder; its first commit creates `CHANGELOG.md` before any code; F-00 makes the repository with that first commit. | docs/NATIVE_WORK_PLAN.md@ad3758c:23; docs/NATIVE_WORK_PLAN.md@ad3758c:333-336 | F-00 | | |
| REQ-WP-0011 | The orchestrator keeps `WORK_STATUS.md`, picks the next ready packages, starts agents, runs the gate, merges and goes on until step 1 is done; it writes no component code. | docs/NATIVE_WORK_PLAN.md@ad3758c:31 | F-00 | | Each role is a separate agent with its own context (line 27, REQ-WP-0196). |
| REQ-WP-0196 | Each role of work plan 2.1 (orchestrator, analyst, implementer, verifier, integrator, image analyst) is a separate agent with its own context. | docs/NATIVE_WORK_PLAN.md@ad3758c:27 | F-00 |  | process rule: checked by the orchestrator; no test expected. Added in round 2 (the 2.1 heading); the verifier's independence is also REQ-WP-0014. |
| REQ-WP-0012 | The analyst of a package reads its documents in full, the Python reference, the upstream sources and outside documentation, and writes the brief `docs/briefs/<id>.md`; it writes no code. | docs/NATIVE_WORK_PLAN.md@ad3758c:32 | F-00 | all packages | |
| REQ-WP-0013 | The implementer writes the code and tests of one package from its brief, inside the package's own folder; it changes no other package's files and no contract. | docs/NATIVE_WORK_PLAN.md@ad3758c:33 | F-00 | F-04 | |
| REQ-WP-0014 | The verifier is someone other than the author; it tries to break the package (tests, security list, known faults, behaviour against the reference) and reports; it does not fix. | docs/NATIVE_WORK_PLAN.md@ad3758c:34 | F-00 | F-03 | |
| REQ-WP-0015 | The integrator owns the shared files (`go.mod`, route table, `nmbd` wiring, changelog, manual index) and merges accepted packages; it never accepts a package the verifier has not passed. | docs/NATIVE_WORK_PLAN.md@ad3758c:35 | F-00 | P-10, P-12 | |
| REQ-WP-0016 | The image analyst takes one image apart and writes its record; it never flashes or boots anything on real hardware. | docs/NATIVE_WORK_PLAN.md@ad3758c:36 | F-00 | I-01…I-10 | process rule: checked by the orchestrator; no test expected. Shared image-method rule (round 2: owner moved from I-01 to F-00, which holds the roles and the gate); each I package proves it for its own image and F-00's gate checks it per image package. Booting I-05 on the Pi 4 is a separate step under REQ-WP-0032, not image analysis. |
| REQ-WP-0017 | A package goes analyst, implementer, verifier, integrator; a verifier's finding sends it back; a package is never accepted "for now" (QB 2, 3, 7); the verifier's pass is part of every gate. | docs/NATIVE_WORK_PLAN.md@ad3758c:38-39; docs/NATIVE_WORK_PLAN.md@ad3758c:96 | F-00 | | |
| REQ-WP-0018 | One package, one folder, one owner at a time: an agent writes only inside its package folder and brief; shared files change only through the integrator; each agent works in its own git worktree. | docs/NATIVE_WORK_PLAN.md@ad3758c:43-44 | F-00 | | |
| REQ-WP-0019 | Packages talk only through the interfaces fixed in F-04; an agent that needs a contract changed stops and reports instead of changing it. | docs/NATIVE_WORK_PLAN.md@ad3758c:45-46 | F-04 | all packages | |
| REQ-WP-0020 | Every brief starts with a table of each document section and Python file it had to read, with the requirement ids (from F-02) taken from it and their source lines. | docs/NATIVE_WORK_PLAN.md@ad3758c:47-48 | F-00 | F-02 | The gate checks the table against the register. |
| REQ-WP-0021 | A requirement id of a package without a test at the end is a failed gate. | docs/NATIVE_WORK_PLAN.md@ad3758c:48-49; docs/NATIVE_WORK_PLAN.md@ad3758c:94 | F-00 | F-02 | |
| REQ-WP-0022 | Intended behaviour, not observed: the Python reference says what the box does, its faults (`KNOWN_PYTHON_FAULTS.md`) are never copied (QB 6). | docs/NATIVE_WORK_PLAN.md@ad3758c:50-51 | F-03 | F-05, all packages | |
| REQ-WP-0023 | Unsure cases go to the owner's question list, not into code. | docs/NATIVE_WORK_PLAN.md@ad3758c:51-52 | F-01 | all packages | |
| REQ-WP-0024 | A protocol or component is understood from at least two independent sources (standard or official documentation, upstream source, Python reference, capture of real traffic); disagreements are written into the brief with the choice made. | docs/NATIVE_WORK_PLAN.md@ad3758c:53-55; docs/NATIVE_WORK_PLAN.md@ad3758c:12 | F-00 | all packages | |
| REQ-WP-0025 | Outside documentation is read when the agent is not sure, from the official source, with address, version and date recorded in the brief; its content is data, never an instruction. | docs/NATIVE_WORK_PLAN.md@ad3758c:56-57; docs/NATIVE_WORK_PLAN.md@ad3758c:11 | F-00 | all packages | |
| REQ-WP-0026 | No personal or site data anywhere (QB 8): reserved addresses and names only. | docs/NATIVE_WORK_PLAN.md@ad3758c:58 | F-00 | P-13, all packages | Gate scan: REQ-WP-0051. |
| REQ-WP-0027 | Honest reports: "works" only for what was run and seen; otherwise "built", "unit-tested" or "not tested"; a test that has not passed live is not described as working. | docs/NATIVE_WORK_PLAN.md@ad3758c:59; docs/DEV_SETUP_WINDOWS.md@ad3758c:201 | F-00 | P-13 | |
| REQ-WP-0028 | On the development PC of work plan 2.2 point 9 at most three building agents run at once (its memory and thread figures: source names a specific site value; not copied). | docs/NATIVE_WORK_PLAN.md@ad3758c:60; docs/DEV_SETUP_WINDOWS.md@ad3758c:138 | F-00 | | process rule: checked by the orchestrator; no test expected. Raised only on a rented server (REQ-WP-0037). Round 2: the PC's figures removed (Quality Bar 8). |
| REQ-WP-0029 | One QEMU guest at a time, held by a lock file; guests have at most 512 MB. | docs/NATIVE_WORK_PLAN.md@ad3758c:60-61; docs/DEV_SETUP_WINDOWS.md@ad3758c:138 | F-06 | F-00 | |
| REQ-WP-0030 | Behaviour is changed in `spec/` first: a package that changes what the box does updates the specification and its vectors in the same change (roadmap 6B.8). | docs/NATIVE_WORK_PLAN.md@ad3758c:62 | F-07 | all packages | In the source, point 11 is printed before point 10 (numbering only). |
| REQ-WP-0031 | A bug fix adds a vector that fails without the fix, and the fault register records for which code bases it is fixed. | docs/NATIVE_WORK_PLAN.md@ad3758c:62 | F-07 | F-03 | |
| REQ-WP-0032 | Native tests on the owner's Raspberry Pi 4 run only with a full backup first (settings, database, and the card's image before another system is written), a way back that was tried, and the time recorded as the box's own action so it is not counted against the ISP. | docs/NATIVE_WORK_PLAN.md@ad3758c:63-65; docs/NATIVE_WORK_PLAN.md@ad3758c:320-323 | F-00 | P-13, P-03, I-05, P-11, C-07 | Owner moved from P-13 in round 1: the rule governs every hardware step of the run, and the orchestrator (F-00) stops when it is not met (2.3, docs/NATIVE_WORK_PLAN.md@ad3758c:70-72); P-03 makes the backup, C-07 marks the time. |
| REQ-WP-0033 | ARM64 is proved first by the cross build, the tests under `qemu-aarch64-static` and the `armsr/armv8` QEMU guest, and on the Pi 4 after that. | docs/NATIVE_WORK_PLAN.md@ad3758c:65-66 | F-00 | F-06, I-02, P-13 | |
| REQ-WP-0034 | The orchestrator runs by itself in a loop: read `WORK_STATUS.md`, start every package whose prerequisites are accepted (within the limits of 2.2 point 9), gate, merge, update the status, changelog and manual, next. | docs/NATIVE_WORK_PLAN.md@ad3758c:70-71; docs/NATIVE_WORK_PLAN.md@ad3758c:12-13 | F-00 | P-12 | |
| REQ-WP-0035 | The orchestrator stops only for a question only the owner can answer, a real-hardware step whose backup or way back is not in place, or a gate that fails three times on the same package (then it reports what was tried); this holds wherever it runs. | docs/NATIVE_WORK_PLAN.md@ad3758c:71-73; docs/NATIVE_WORK_PLAN.md@ad3758c:85 | F-00 | F-01 | |
| REQ-WP-0036 | `WORK_STATUS.md` holds per package the state (waiting, brief, code, verify, accepted, blocked), the commit, the gate's result and open questions, so the work continues after any interruption, on any machine that has the repositories, without redoing accepted packages. | docs/NATIVE_WORK_PLAN.md@ad3758c:73-75; docs/NATIVE_WORK_PLAN.md@ad3758c:79 | F-00 | | |
| REQ-WP-0037 | If a server is rented: nested virtualisation (KVM) is checked before renting; with 16 GB RAM or more the limits of 2.2 point 9 are raised; an ARM64 server runs arm64 tests natively. | docs/NATIVE_WORK_PLAN.md@ad3758c:81 | F-00 | F-06 | process rule: checked by the orchestrator; no test expected. Conditional: development stays on the PC for now (REQ-WP-0042). |
| REQ-WP-0038 | If work moves to a server: both repositories come from their remotes, the image and firmware set is checked against its checksums after the copy, the tools of DEV_SETUP 5A are installed, and the agents' skills and notes are copied on purpose and kept in step (the parts that hold nothing private are better kept in the native repository). | docs/NATIVE_WORK_PLAN.md@ad3758c:82 | F-00 | | process rule: checked by the orchestrator; no test expected. Conditional. See REQ-DV-0022 for notes that must never be committed. |
| REQ-WP-0039 | Never on a rented server: the box's password and the owner's keys, the signing key, private logs and audit notes, copies of the box's real database, the private list of owner values used by the build scan; tests there use only the made-up data of the test bed. | docs/NATIVE_WORK_PLAN.md@ad3758c:83 | F-00 | F-06, P-02 | |
| REQ-WP-0040 | A server that must reach the box gets its own key that can be withdrawn at any time, only through the tunnel or Tailscale and only while maintenance mode is on (porting plan 7D); never the owner's key or password; steps on the owner's box stay with the owner's PC. | docs/NATIVE_WORK_PLAN.md@ad3758c:84 | P-05 | N-15, N-18 | Box side: access only while maintenance mode is on. |
| REQ-WP-0041 | A rented server has key-only SSH, no other service open to the internet, automatic security updates, and its disk wiped when the rental ends; the session runs in a persistent terminal followed from the owner's devices. | docs/NATIVE_WORK_PLAN.md@ad3758c:85 | F-00 | | process rule: checked by the orchestrator; no test expected. Conditional. |
| REQ-WP-0042 | Development stays on the one Windows PC (WSL2), with the limits of 2.2 point 9; the server points apply only if the owner rents a server later. | docs/NATIVE_WORK_PLAN.md@ad3758c:87 | F-00 | | process rule: checked by the orchestrator; no test expected. Owner decision 2026-11-10. |
| REQ-WP-0043 | The gate of every package builds for amd64, arm64, arm v7, mips and mipsle. | docs/NATIVE_WORK_PLAN.md@ad3758c:91 | F-00 | | |
| REQ-WP-0044 | The gate runs unit tests on amd64 and under QEMU user mode on arm64 and arm v7, with the race detector on amd64. | docs/NATIVE_WORK_PLAN.md@ad3758c:91-92 | F-00 | F-06 | |
| REQ-WP-0045 | The gate runs `gofmt`, `go vet`, `staticcheck`, `gosec`, `govulncheck` and complexity and size limits (clean code rules). | docs/NATIVE_WORK_PLAN.md@ad3758c:92-93 | F-00 | | Tool versions: REQ-DV-0009. |
| REQ-WP-0046 | Every parser has fuzz tests. | docs/NATIVE_WORK_PLAN.md@ad3758c:93 | F-00 | every package with a parser | |
| REQ-WP-0047 | Tests are hermetic only. | docs/NATIVE_WORK_PLAN.md@ad3758c:93 | F-06 | F-00 | |
| REQ-WP-0048 | Each known Python fault of a package has a test. | docs/NATIVE_WORK_PLAN.md@ad3758c:93-94 | F-03 | F-00 | |
| REQ-WP-0049 | The Phase 0 checks that apply to a package pass in its gate. | docs/NATIVE_WORK_PLAN.md@ad3758c:94 | F-05 | F-00 | |
| REQ-WP-0050 | Every package has a written security check (QB 7). | docs/NATIVE_WORK_PLAN.md@ad3758c:94-95 | F-00 | all packages | |
| REQ-WP-0051 | The gate scans for personal data and secrets. | docs/NATIVE_WORK_PLAN.md@ad3758c:94-95 | F-00 | P-13 | |
| REQ-WP-0052 | Every new dependency gets a licence entry. | docs/NATIVE_WORK_PLAN.md@ad3758c:95 | P-08 | F-00 | |
| REQ-WP-0053 | Size, memory and CPU figures are recorded for every package. | docs/NATIVE_WORK_PLAN.md@ad3758c:95 | F-00 | P-13 | |
| REQ-WP-0054 | Changelog, manual chapter and command reference are updated with every package (roadmap 6B.7). | docs/NATIVE_WORK_PLAN.md@ad3758c:95-96 | P-12 | F-00 | |
| REQ-WP-0055 | Size rule: a package is what one agent can read, write and test in one sitting (as a guide at most about 1,500 lines of Python reference, or one protocol); a bigger one is split by its analyst before coding. | docs/NATIVE_WORK_PLAN.md@ad3758c:100-101 | F-02 | all packages | |
| REQ-WP-0056 | The foundation packages come in order, before anything else. | docs/NATIVE_WORK_PLAN.md@ad3758c:103 | F-00 | | Section 4 lists only F-00 to F-06. Q-02 settled: F-07 is part of the foundation of 3.0; its first version (`spec/` layout, formats, vector runner) is accepted before Core starts and then grows with each package (rule 11 of 2.2). |
| REQ-WP-0057 | F-00 Skeleton: the `NetworkMeasureBox-DePyD/` layout, `CHANGELOG.md` (first commit), `KNOWN_PYTHON_FAULTS.md`, `WORK_STATUS.md`, Go module, build and gate scripts for the matrix of 2.4, and the toolchain version fixed (the roadmap names Go 1.23+; the version installed on the PC: source names a specific site value; not copied). | docs/NATIVE_WORK_PLAN.md@ad3758c:107 | F-00 | | |
| REQ-WP-0058 | F-01 Owner's open decisions: the list of section 7 answered and written into the roadmap. | docs/NATIVE_WORK_PLAN.md@ad3758c:108 | F-01 | | |
| REQ-WP-0059 | F-02 Requirements register: every requirement of both plans, `TESTING.md` (release checklist, To fix, Known limits), `ARCHITECTURE.md`, the READMEs and the changelog as one numbered list with source lines; each id assigned to a package; a second agent checks that no paragraph is without an id. | docs/NATIVE_WORK_PLAN.md@ad3758c:109 | F-02 | | This file is one part of it. |
| REQ-WP-0060 | F-03 Known Python faults: the roadmap's register plus the private audit and data-quality lists (open security faults by category only), each with the correct behaviour and the package that must prove it. | docs/NATIVE_WORK_PLAN.md@ad3758c:110 | F-03 | | |
| REQ-WP-0061 | F-04 Contracts: configuration format (`nmb.conf` compatible, `%` plain, `__KEEP_CURRENT__`), database schema and on-disk formats (same on every architecture), event and log record, platform interfaces (routes/links, Wi-Fi, firewall, service control, hardware), API contract and route table, translation keys, status file formats. | docs/NATIVE_WORK_PLAN.md@ad3758c:111 | F-04 | C-01, C-02, S-01 | |
| REQ-WP-0062 | F-05 Phase 0 reference: golden files and checks captured from the Python version, each reviewed against F-03 before use, and a comparison harness (Python and native, same input). | docs/NATIVE_WORK_PLAN.md@ad3758c:112 | F-05 | F-03 | |
| REQ-WP-0063 | F-06 Test bed: hermetic helpers (fake clock, network namespaces with virtual cables, `mac80211_hwsim` where the kernel has it, fake servers for every engine and protocol, the QEMU guest runner with its lock). | docs/NATIVE_WORK_PLAN.md@ad3758c:113 | F-06 | | |
| REQ-WP-0064 | F-07 Shared specification and conformance vectors: `spec/` (configuration, schema, API, CSV columns, codes, protocol details, limits, texts), `spec/vectors/` run by every code base, the fault register with a status per code base, the parity table per target (roadmap 6B.8); built in step 1 so the step-2 C build starts from it. | docs/NATIVE_WORK_PLAN.md@ad3758c:114 | F-07 | NEW: X2-01 Step-2 router tiers | |
| REQ-WP-0065 | C-01 Configuration: load, validate, defaults, atomic locked save (reference `nmb/config.py`; Phase 0, rules from fixes). | docs/NATIVE_WORK_PLAN.md@ad3758c:120 | C-01 | F-04 | |
| REQ-WP-0066 | C-02 Database: RAM live copy, snapshots, retention, migration from a Python database (reference `nmb/db/`; roadmap 4.5). | docs/NATIVE_WORK_PLAN.md@ad3758c:121 | C-02 | F-04 | |
| REQ-WP-0067 | C-03 Structured log, run-time debug level, crash evidence (reference `nmb/devlog.py`, `report/crash_trap.py`; 7D, 4.15). | docs/NATIVE_WORK_PLAN.md@ad3758c:122 | C-03 | | |
| REQ-WP-0068 | C-04 Secret masking (reference `common/redact.py`, `secret_keep.py`; 4.24). | docs/NATIVE_WORK_PLAN.md@ad3758c:123 | C-04 | | |
| REQ-WP-0069 | C-05 Addresses, version, user agent, monotonic time helpers (reference `common/netaddr.py`, `version.py`, `user_agent.py`; QB 6). | docs/NATIVE_WORK_PLAN.md@ad3758c:124 | C-05 | | |
| REQ-WP-0070 | C-06 Translations: keys, 11 languages, build checks (reference `web/i18n.py`, `bot/i18n.py`, the `locales/` folders; 4.26, 4.4A point 6). | docs/NATIVE_WORK_PLAN.md@ad3758c:125 | C-06 | | |
| REQ-WP-0071 | C-07 Events and own-action windows (reference `common/own_action.py`; register). | docs/NATIVE_WORK_PLAN.md@ad3758c:126 | C-07 | | |
| REQ-WP-0072 | C-11 Measurement setup: classification, history, the words in 11 languages (reference `common/net_setup.py`, `common/setup_texts.py`, `db/setup_backfill.py`; 4.29). | docs/NATIVE_WORK_PLAN.md@ad3758c:127 | C-11 | - | Also in: the packages named in REQ-WP-0134 (they show or store the setup). |
| REQ-WP-0073 | C-08 Storage guard, purge, memory and flash health (reference `common/storage_guard.py`, `purge.py`, `memory.py`; 4.23, 4.26). | docs/NATIVE_WORK_PLAN.md@ad3758c:128 | C-08 | | |
| REQ-WP-0074 | C-09 Tasks and load from `/proc`, IRQ and offload telemetry (reference `common/tasks.py`; 4.13, 4.22). | docs/NATIVE_WORK_PLAN.md@ad3758c:129 | C-09 | | |
| REQ-WP-0075 | C-10 Country data files and location (new; QB 9). | docs/NATIVE_WORK_PLAN.md@ad3758c:130 | C-10 | | |
| REQ-WP-0076 | M-01 ICMP probe, statistics, minute summaries (reference `probe/icmp.py`, `engine.py`; 4.1). | docs/NATIVE_WORK_PLAN.md@ad3758c:136 | M-01 | | `probe/engine.py` is also named for M-02 (see Proposed package changes). |
| REQ-WP-0077 | M-02 Outage and loss-pattern engine (reference `probe/engine.py`, `report/metrics.py`; 4.18). | docs/NATIVE_WORK_PLAN.md@ad3758c:137 | M-02 | | Shares two files with M-01 and R-01. |
| REQ-WP-0078 | M-03 Traffic gate and other devices' traffic (reference `probe/traffic_gate.py`, `cross_traffic.py`; Phase 0 table). | docs/NATIVE_WORK_PLAN.md@ad3758c:138 | M-03 | | |
| REQ-WP-0079 | M-04 Runner, schedule, per-engine intervals, cancel (reference `speed/runner.py`, `scheduler.py`, `cancel.py`; 4.25, 4.3). | docs/NATIVE_WORK_PLAN.md@ad3758c:139 | M-04 | | |
| REQ-WP-0080 | M-05 Data caps, budgets, Ookla pause (reference `speed/data_cap.py`, `data_budget.py`, `anti_ban.py`; rules from fixes). | docs/NATIVE_WORK_PLAN.md@ad3758c:140 | M-05 | | |
| REQ-WP-0081 | M-06 Cloudflare engine (reference `speed/cloudflare.py`; 4.3). | docs/NATIVE_WORK_PLAN.md@ad3758c:141 | M-06 | | |
| REQ-WP-0082 | M-07 Fast.com engine (reference `speed/fastcom.py`; 4.26). | docs/NATIVE_WORK_PLAN.md@ad3758c:142 | M-07 | | |
| REQ-WP-0083 | M-08 RTR (RMBT) engine (reference `speed/rtr.py`, `components/open-rmbt`; 4.6). | docs/NATIVE_WORK_PLAN.md@ad3758c:143 | M-08 | | |
| REQ-WP-0084 | M-09 iPerf3 client, server names, local server (reference `speed/iperf.py`, `iperf_directory.py`; 4.8). | docs/NATIVE_WORK_PLAN.md@ad3758c:144 | M-09 | | |
| REQ-WP-0085 | M-10 Ookla engine, under the decision of section 7 (reference `speed/ookla.py`; 4.3). | docs/NATIVE_WORK_PLAN.md@ad3758c:145 | M-10 | | |
| REQ-WP-0086 | M-11 Traceroute, port check, packet capture (reference `speed/traceroute.py`, `port_checker.py`, `sniffer/`; no roadmap section). | docs/NATIVE_WORK_PLAN.md@ad3758c:146 | M-11 | | |
| REQ-WP-0087 | M-12 DNS audit, STUN and NAT, public address (reference `netd/dns.py`, `netd/geoip.py`; 4.20, 4.21). | docs/NATIVE_WORK_PLAN.md@ad3758c:147 | M-12 | | Optional split proposed in the roadmap part. |
| REQ-WP-0088 | M-13 OONI (reference `probe/ooni.py`; 4.7). | docs/NATIVE_WORK_PLAN.md@ad3758c:148 | M-13 | | |
| REQ-WP-0089 | M-14 Wi-Fi radio telemetry (reference `netd/wifi.py`, status parts; 4.19). | docs/NATIVE_WORK_PLAN.md@ad3758c:149 | M-14 | N-07 | |
| REQ-WP-0090 | N-01 Links, addresses, routes, hot-plug over netlink (reference `netd/interfaces.py`, `detector.py`, `kernel.py`; 4.2). | docs/NATIVE_WORK_PLAN.md@ad3758c:155 | N-01 | | |
| REQ-WP-0091 | N-02 Firewall for every mode (reference `netd/firewall.py`, `os/nmb-guard.nft`; register, firewall). | docs/NATIVE_WORK_PLAN.md@ad3758c:156 | N-02 | | |
| REQ-WP-0092 | N-03 Modes and switching (reference `netd/modes.py`, `daemon.py`, `services.py`; Phase 0). | docs/NATIVE_WORK_PLAN.md@ad3758c:157 | N-03 | | |
| REQ-WP-0093 | N-04 Automatic mode and its probes, DHCP and PPPoE discovery (reference `netd/auto_mode.py`; 4.28). | docs/NATIVE_WORK_PLAN.md@ad3758c:158 | N-04 | | |
| REQ-WP-0094 | N-05 Built-in DHCP client, DHCP guard, PPPoE (reference `netd/modes.py`, `wifi_share.py` guard; 4.27, rules from fixes). | docs/NATIVE_WORK_PLAN.md@ad3758c:159 | N-05 | N-03, N-08 | Shares files with N-03 and N-08. |
| REQ-WP-0095 | N-06 LAN passthrough and link aggregation (reference `netd/lan_topology.py`; porting plan 7A). | docs/NATIVE_WORK_PLAN.md@ad3758c:160 | N-06 | | |
| REQ-WP-0096 | N-07 Wi-Fi client and access point control, regulatory domain (reference `netd/wifi.py`, `regdomains.py`; 4.14). | docs/NATIVE_WORK_PLAN.md@ad3758c:161 | N-07 | M-14 | |
| REQ-WP-0097 | N-08 Guest Wi-Fi, Wi-Fi uplink, sharing on a LAN port (reference `netd/guest_wifi.py`, `wifi_uplink.py`, `wifi_share.py`; 4.26). | docs/NATIVE_WORK_PLAN.md@ad3758c:162 | N-08 | N-05 | |
| REQ-WP-0098 | N-09 Captive portals (new; 6B.6). | docs/NATIVE_WORK_PLAN.md@ad3758c:163 | N-09 | | |
| REQ-WP-0099 | N-10 USB tethering and mobile uplink (reference `netd/usb_uplink.py`; 4.27). | docs/NATIVE_WORK_PLAN.md@ad3758c:164 | N-10 | | |
| REQ-WP-0100 | N-11 IPv6 on the LAN, port forwarding, UPnP on the WAN, proxy (reference `netd/ipv6_lan.py`, `port_forward.py`, `upnp_wan.py`, `proxy.py`; no roadmap section). | docs/NATIVE_WORK_PLAN.md@ad3758c:165 | N-11 | | Four modules; the analyst checks the size rule. |
| REQ-WP-0101 | N-12 DNS blocker (reference `netd/adblock.py`; 4.9). | docs/NATIVE_WORK_PLAN.md@ad3758c:166 | N-12 | | |
| REQ-WP-0102 | N-13 WireGuard and QR codes (reference `netd/wireguard.py`; 4.10). | docs/NATIVE_WORK_PLAN.md@ad3758c:167 | N-13 | | |
| REQ-WP-0103 | N-14 OpenVPN (reference `netd/openvpn.py`; 4.10). | docs/NATIVE_WORK_PLAN.md@ad3758c:168 | N-14 | | |
| REQ-WP-0104 | N-15 Tailscale (reference `netd/tailscale.py`; 4.10). | docs/NATIVE_WORK_PLAN.md@ad3758c:169 | N-15 | | |
| REQ-WP-0105 | N-18 Remote access without Tailscale on the device: WireGuard to a hub, subnet router on a Go box, reverse SSH (new; porting plan 7D). | docs/NATIVE_WORK_PLAN.md@ad3758c:170 | N-18 | P-05 | |
| REQ-WP-0106 | N-16 DynDNS (reference `netd/duckdns.py`; porting plan 7.10). | docs/NATIVE_WORK_PLAN.md@ad3758c:171 | N-16 | | |
| REQ-WP-0107 | N-17 Hardware and driver finder, system detection, device explorer (reference `netd/drivers.py`, `os_detector.py`, `explorer.py`; 4.11). | docs/NATIVE_WORK_PLAN.md@ad3758c:172 | N-17 | | |
| REQ-WP-0108 | R-01 Report figures, status codes, verdicts (reference `report/metrics.py`, `status_codes.py`; 4.17). | docs/NATIVE_WORK_PLAN.md@ad3758c:178 | R-01 | M-02 | |
| REQ-WP-0109 | R-02 CSV exports (reference `report/csv_export.py`; no roadmap section). | docs/NATIVE_WORK_PLAN.md@ad3758c:179 | R-02 | | |
| REQ-WP-0110 | R-03 SLA dossier and charts (reference `report/dossier.py`, `svg_charts.py`; 4.17). | docs/NATIVE_WORK_PLAN.md@ad3758c:180 | R-03 | | |
| REQ-WP-0111 | R-04 Diagnostic bundle for LLMs (reference `report/llm_bundle.py`; 4.24). | docs/NATIVE_WORK_PLAN.md@ad3758c:181 | R-04 | | |
| REQ-WP-0112 | R-05 E-mail, diagnostic logger (reference `report/smtp.py`, `invasive_logger.py`; 4.15). | docs/NATIVE_WORK_PLAN.md@ad3758c:182 | R-05 | | |
| REQ-WP-0113 | S-01 Web server: login, sessions, CSRF, route table, live stream (reference `web/server.py` transport parts; 4.4, QB 7). | docs/NATIVE_WORK_PLAN.md@ad3758c:188 | S-01 | | |
| REQ-WP-0114 | S-02 to S-0n: the API by area (status, speed, network, Wi-Fi, VPN, blocker, system, exports), one package per area (reference `web/server.py` handlers; 4.4, clean code). | docs/NATIVE_WORK_PLAN.md@ad3758c:189 | S-02…S-0n | S-01 | `web/server.py` is 4,612 lines; the split per area is fixed by the first S-02 analyst. |
| REQ-WP-0115 | U-01 Web UI shell, AMOLED black look, page addresses, size budgets (reference `web/assets/`; 4.4A, 6B.1). | docs/NATIVE_WORK_PLAN.md@ad3758c:190 | U-01 | | |
| REQ-WP-0116 | U-02 to U-0n: Web UI pages, one package per page (reference `web/assets/dashboard.*`; 4.4A). | docs/NATIVE_WORK_PLAN.md@ad3758c:191 | U-02…U-0n | U-01 | |
| REQ-WP-0117 | S-10 `nmb-cli`: JSON commands, menu, plain output (reference `nmb/tui/`; Phase 0, 6B.2). | docs/NATIVE_WORK_PLAN.md@ad3758c:192 | S-10 | | |
| REQ-WP-0118 | S-11 Telegram bot (reference `nmb/bot/`; 4.16, register bot). | docs/NATIVE_WORK_PLAN.md@ad3758c:193 | S-11 | | `nmb/bot/` is 4,132 lines: split proposed below. |
| REQ-WP-0119 | S-12 Collector client (reference `collector/client.py`; 4.25). | docs/NATIVE_WORK_PLAN.md@ad3758c:194 | S-12 | | |
| REQ-WP-0120 | S-13 Collector server (reference `collector/server.py`; 4.25). | docs/NATIVE_WORK_PLAN.md@ad3758c:195 | S-13 | | |
| REQ-WP-0121 | P-01 Self-healing watchdog (reference `common/heal.py`; 4.12). | docs/NATIVE_WORK_PLAN.md@ad3758c:201 | P-01 | | |
| REQ-WP-0122 | P-02 Updates: file upload and channel, signature, try and roll back (reference `common/updater.py`; 6B.3, 7D). | docs/NATIVE_WORK_PLAN.md@ad3758c:202 | P-02 | | |
| REQ-WP-0123 | P-03 Backup, restore, USB import, factory reset, unique password per device (reference `common/backup.py`, `usb_import.py`, `bin/nmb-factory-reset`; QB 7, Phase 0). | docs/NATIVE_WORK_PLAN.md@ad3758c:203 | P-03 | | |
| REQ-WP-0124 | P-04 Sandbox profiles and their tests (reference `systemd/` units; 6B.5). | docs/NATIVE_WORK_PLAN.md@ad3758c:204 | P-04 | | |
| REQ-WP-0125 | P-05 Maintenance mode and audit log (new; 7D). | docs/NATIVE_WORK_PLAN.md@ad3758c:205 | P-05 | | |
| REQ-WP-0126 | P-06 Privileged helper with allow-lists (reference `bin/nmb-helper`; register). | docs/NATIVE_WORK_PLAN.md@ad3758c:206 | P-06 | | |
| REQ-WP-0127 | P-07 Self-test and scenarios (reference `nmb/selftest/`; TESTING). | docs/NATIVE_WORK_PLAN.md@ad3758c:207 | P-07 | | |
| REQ-WP-0128 | P-08 Licences page and build check; disclaimer (reference `common/donate.py`; 6B.3A). | docs/NATIVE_WORK_PLAN.md@ad3758c:208 | P-08 | | |
| REQ-WP-0129 | P-09 Add-ons: themes and CLI colour schemes (new; 6B.1, 6B.2). | docs/NATIVE_WORK_PLAN.md@ad3758c:209 | P-09 | | |
| REQ-WP-0130 | P-10 `nmbd` single binary, multi-call names, service units (reference `bin/`, `systemd/`, `os/`; section 5). | docs/NATIVE_WORK_PLAN.md@ad3758c:210 | P-10 | | |
| REQ-WP-0131 | P-11 Packages and images for amd64 and arm64; evaluation of the bases (reference `setup.sh`; 6A). | docs/NATIVE_WORK_PLAN.md@ad3758c:211 | P-11 | I-09 | |
| REQ-WP-0132 | P-12 User manual and command reference, with their build checks (new; 6B.7). | docs/NATIVE_WORK_PLAN.md@ad3758c:212 | P-12 | | |
| REQ-WP-0133 | P-13 Release proof: release checklist, Python against native on the same box and day, outside scan (reference `docs/TESTING.md`; QB 1, 4, 5, 7). | docs/NATIVE_WORK_PLAN.md@ad3758c:213 | P-13 | | |
| REQ-WP-0134 | The measurement setup of roadmap 4.29 is shown or stored by M-01, M-02, M-04, N-01, N-03, R-01 to R-05, S-02, U-02, S-10, S-11 and S-12/S-13 besides C-11, and each of their briefs lists it as a requirement. | docs/NATIVE_WORK_PLAN.md@ad3758c:215 | C-11 | M-01, M-02, M-04, N-01, N-03, R-01, R-02, R-03, R-04, R-05, S-02, U-02, S-10, S-11, S-12, S-13 | |
| REQ-WP-0135 | F-02 may add, split or merge packages; every Python file, every roadmap section and every `TESTING.md` row ends up in exactly one package, and the register's check fails otherwise. | docs/NATIVE_WORK_PLAN.md@ad3758c:217-218 | F-02 | | Six Python files are named for two packages today (Proposed package changes, point 1). |
| REQ-WP-0136 | Order 1: F-00 to F-06 first, and in parallel the image records of section 5 (they need no native code). | docs/NATIVE_WORK_PLAN.md@ad3758c:222 | F-00 | I-01…I-09 | |
| REQ-WP-0137 | Order 2: Core (C) first, then Measuring (M), Network (N) and Reports (R) in parallel as their contracts allow. | docs/NATIVE_WORK_PLAN.md@ad3758c:223 | F-00 | | |
| REQ-WP-0138 | Order 3: Surfaces (S, U) on top of them; System and product (P) alongside. | docs/NATIVE_WORK_PLAN.md@ad3758c:224 | F-00 | | |
| REQ-WP-0139 | Order 4: P-10 to P-13 last; step 1 ends with P-13 accepted and the owner's acceptance of any remaining known limit. | docs/NATIVE_WORK_PLAN.md@ad3758c:225 | P-13 | F-00 | |
| REQ-WP-0140 | The image and firmware set (roadmap 6B.4) is used from its copy inside WSL2, outside every repository, and never committed; only records, lists, versions and checksums go into the repository, no file taken out of an image and nothing from a maker's closed firmware (porting plan 7H point 3). | docs/NATIVE_WORK_PLAN.md@ad3758c:229; docs/NATIVE_WORK_PLAN.md@ad3758c:278-279; docs/DEV_SETUP_WINDOWS.md@ad3758c:136 | F-00 | I-01…I-10 | DEV_SETUP line 136 still names the Windows Downloads folder; the work plan (newer) uses the WSL2 copy. |
| REQ-WP-0141 | Images are read only on the PC: no image is written to a device for analysis. | docs/NATIVE_WORK_PLAN.md@ad3758c:230 | F-00 | I-01…I-10 | process rule: checked by the orchestrator; no test expected. Shared image-method rule (round 2: owner moved from I-01 to F-00, which holds the roles and the gate); each I package proves it for its own image and F-00's gate checks it per image package. |
| REQ-WP-0142 | Unpacking is done on copies in a work folder outside the repository; the files are untrusted data, unpacked into fresh empty folders, without root and without loop mounts. | docs/NATIVE_WORK_PLAN.md@ad3758c:230-232 | F-00 | I-01…I-10 | Shared image-method rule (round 2: owner moved from I-01 to F-00, which holds the roles and the gate); each I package proves it for its own image and F-00's gate checks it per image package. The gate can check that no image file and no unpacked file is in the repository; the rest is checked by the orchestrator. |
| REQ-WP-0143 | Disk images (`*.img.gz`, `*.raw.xz`) are decompressed (a gzip warning from OpenWrt trailing metadata is expected), their partition table read (`sgdisk -p`, `fdisk -l`), each partition cut out by offset and read with `mtools` (FAT), `debugfs rdump` (ext4), `unsquashfs` (squashfs), other file systems by booting the image in QEMU. | docs/NATIVE_WORK_PLAN.md@ad3758c:238 | I-01 | I-02, I-03, I-05, I-08 | Method rule for disk images, owned like REQ-WP-0144 to -0146 by the image package of that kind: I-01's OpenWrt x86-64 image is a disk image (line 238); the other disk images prove it in their own packages (Also in). Round 2: owner kept, see the register's answer to the round-2 finding. |
| REQ-WP-0144 | Root file system archives (OpenWrt `rootfs.tar.gz`, Alpine Pi archive) are listed and unpacked directly; Alpine's `modloop` is read as a squashfs with the modules and firmware. | docs/NATIVE_WORK_PLAN.md@ad3758c:239 | I-07 | I-01, I-02, I-03, I-04 | |
| REQ-WP-0145 | Installer ISOs (Debian netinst, Alpine, Fedora IoT) are unpacked with `7z`; their package index and kernel package are read, and the installed system is judged from an image built or installed in QEMU, not from the installer. | docs/NATIVE_WORK_PLAN.md@ad3758c:240 | I-06 | I-07, I-08 | |
| REQ-WP-0146 | OpenWrt SDKs, ImageBuilders, toolchains and `kernel-debug` are unpacked and give the kernel `.config`, the package index, the `kmods` folder, `profiles.json`, compiler and C library versions. | docs/NATIVE_WORK_PLAN.md@ad3758c:241 | I-02 | I-01, I-03, I-04 | |
| REQ-WP-0147 | Router firmware (step 2 and later) is taken apart with `binwalk` and the maker's header (TP-Link v1 hardware id, kernel, root and boot lengths; ASUS `trx`); kernel and squashfs cut out by offset; old stock squashfs variants and UBI/JFFS2 need their own readers; stock firmware only for the rules of porting plan 7H. | docs/NATIVE_WORK_PLAN.md@ad3758c:242 | I-10 | P-02, NEW: X2-01 Step-2 router tiers | Step 2. |
| REQ-WP-0148 | Before image work starts, the PC gets `cpio`, `xorriso` or `bsdtar` and `ostree`. | docs/NATIVE_WORK_PLAN.md@ad3758c:244-245 | F-00 | I-01…I-08 | Development tools on the PC, like REQ-DV-0009 (F-00). Round 2: owner moved from I-01; the router-firmware readers moved to REQ-WP-0197 (I-10). |
| REQ-WP-0197 | Before router firmware is taken apart, the PC gets a UBI reader, a JFFS2 reader and a reader for vendor squashfs variants. | docs/NATIVE_WORK_PLAN.md@ad3758c:244-245 | I-10 | NEW: X2-01 Step-2 router tiers, P-14 | Split from REQ-WP-0148 in round 2: these readers serve router firmware (REQ-WP-0147), not the step-1 base images. |
| REQ-WP-0149 | One record per image in `NetworkMeasureBox-DePyD/docs/images/<image>.md`, with a summary into roadmap section 6A. | docs/NATIVE_WORK_PLAN.md@ad3758c:249 | F-00 | I-01…I-09 | Shared image-method rule (round 2: owner moved from I-01 to F-00, which holds the roles and the gate); each I package proves it for its own image and F-00's gate checks it per image package. |
| REQ-WP-0150 | Each image record holds: file name and SHA-256 checked against the publisher; partition and file system layout, boot method, update method (sysupgrade, A/B, ostree); kernel version and the options NMB needs (bridge, bonding, VLAN, nftables, WireGuard, PPP and PPPoE, `mac80211`/`cfg80211`, `mac80211_hwsim`, USB network and tethering drivers, namespaces, seccomp, Landlock, cgroups, watchdog); modules and firmware; init system and service manager; C library; package manager and package count; which 6A criterion-3 tools are present with versions; any Python (criterion 8); own network or firewall manager and how it is switched off (criterion 5); size on disk; and from the QEMU boot: boot time, free memory and storage, listening ports, running services. | docs/NATIVE_WORK_PLAN.md@ad3758c:249-257 | F-00 | I-01…I-08 | Shared image-method rule (round 2: owner moved from I-01 to F-00, which holds the roles and the gate); each I package proves it for its own image and F-00's gate checks it per image package. A gate script can check that every field is present in each record. |
| REQ-WP-0151 | Each record field is filled from the image with the command that produced it, or marked "not determined" with the reason; nothing is filled in from memory. | docs/NATIVE_WORK_PLAN.md@ad3758c:257-258 | F-00 | I-01…I-08 | Shared image-method rule (round 2: owner moved from I-01 to F-00, which holds the roles and the gate); each I package proves it for its own image and F-00's gate checks it per image package. |
| REQ-WP-0152 | The image records fill the 6A criteria table for choosing the base, in the roadmap's order (OpenWrt, Debian, Alpine, Fedora IoT). | docs/NATIVE_WORK_PLAN.md@ad3758c:262 | I-09 | P-11 | |
| REQ-WP-0153 | A capability table records which kernel features, tools and paths exist on which base. | docs/NATIVE_WORK_PLAN.md@ad3758c:263 | I-09 | F-04, N-17 | |
| REQ-WP-0154 | The platform interfaces (F-04) and the run-time detection are designed and tested against the capability table (missing tool, tool outside `PATH`, other init system), not against one developer machine. | docs/NATIVE_WORK_PLAN.md@ad3758c:263-265 | F-04 | N-17, F-06, P-10 | Owner moved from N-17 in round 1: the platform interfaces are F-04's contract; N-17's run-time detection is tested against the same table. |
| REQ-WP-0155 | OpenWrt x86-64 and `armsr` v8/v7 boot in QEMU as guests the native program is installed into and tested on (modes, firewall, DHCP, PPPoE between virtual ports); the root archives serve as quick user-mode test roots for arm64 and arm v7 binaries. | docs/NATIVE_WORK_PLAN.md@ad3758c:266-268 | F-06 | I-01, I-02, I-03, N-02, N-03, N-05 | |
| REQ-WP-0156 | The SDKs and toolchains are the build inputs for the later C11/musl builds and a first `nmbd` package; the ImageBuilder assembles an NMB image without rebuilding the base; the kernel configurations confirm that no own kernel is needed. | docs/NATIVE_WORK_PLAN.md@ad3758c:269-271 | P-11 | I-02, NEW: X2-01 Step-2 router tiers | |
| REQ-WP-0157 | The firmware and module lists of the images are data for the driver finder. | docs/NATIVE_WORK_PLAN.md@ad3758c:272 | N-17 | I-01…I-08 | |
| REQ-WP-0158 | The bases' default network and firewall configuration is the input for the coexistence design. | docs/NATIVE_WORK_PLAN.md@ad3758c:272-273 | N-02 | N-03, P-11 | |
| REQ-WP-0159 | The bases' package indexes feed the licence list (6B.3A). | docs/NATIVE_WORK_PLAN.md@ad3758c:273 | P-08 | I-01…I-08 | |
| REQ-WP-0160 | Firmware handling (6B.3): the router image header layouts define the parsers; their tests and fuzz seeds are made-up headers of the same layout, not the makers' files. | docs/NATIVE_WORK_PLAN.md@ad3758c:274-275 | P-02 | I-10 | |
| REQ-WP-0161 | From step 2, the stock firmware inventories per target give the missing parts (7H). | docs/NATIVE_WORK_PLAN.md@ad3758c:276 | I-10 | NEW: X2-01 Step-2 router tiers | Step 2. |
| REQ-WP-0162 | I-01: OpenWrt 25.12.6 x86-64 (squashfs and ext4); also the first QEMU guest of F-06. | docs/NATIVE_WORK_PLAN.md@ad3758c:285 | I-01 | F-06 | |
| REQ-WP-0163 | I-02: OpenWrt 25.12.6 `armsr/armv8` with SDK, ImageBuilder and toolchain; the ARM64 guest. | docs/NATIVE_WORK_PLAN.md@ad3758c:286 | I-02 | F-06 | |
| REQ-WP-0164 | I-03: OpenWrt 25.12.6 `armsr/armv7` with its tree; the 32-bit ARM guest. | docs/NATIVE_WORK_PLAN.md@ad3758c:287 | I-03 | F-06 | |
| REQ-WP-0165 | I-04: OpenWrt 25.12.6 `x86/generic` tree; 32-bit x86, cross-check of I-01. | docs/NATIVE_WORK_PLAN.md@ad3758c:288 | I-04 | | |
| REQ-WP-0166 | I-05: OpenWrt 25.12.6 Raspberry Pi 4, taken apart first; booted on the Pi 4 only under section 7 point 5, with the card image saved before. | docs/NATIVE_WORK_PLAN.md@ad3758c:289 | I-05 | P-13 | |
| REQ-WP-0167 | I-06: Debian 13.7.0 amd64 and arm64; package index and kernel package; minimal installed system in QEMU. | docs/NATIVE_WORK_PLAN.md@ad3758c:290 | I-06 | | |
| REQ-WP-0168 | I-07: Alpine 3.24.4 x86-64 and Pi arm64; `modloop`, `apks`, diskless mode. | docs/NATIVE_WORK_PLAN.md@ad3758c:291 | I-07 | | |
| REQ-WP-0169 | I-08: Fedora IoT 44 arm64 and x86-64; largest files, done after the first three bases as the roadmap orders. | docs/NATIVE_WORK_PLAN.md@ad3758c:292 | I-08 | | |
| REQ-WP-0170 | I-09: comparison of I-01 to I-08 and a proposal for 6A; the owner decides. | docs/NATIVE_WORK_PLAN.md@ad3758c:293 | I-09 | F-01 | |
| REQ-WP-0171 | I-10 and later: router firmware, TL-WR1043ND v1 first; step 2, not started in step 1. | docs/NATIVE_WORK_PLAN.md@ad3758c:294 | I-10 | NEW: X2-01 Step-2 router tiers | |
| REQ-WP-0172 | The owner sees `WORK_STATUS.md`, `CHANGELOG.md`, the briefs, the image records, and for each accepted package its evidence: tests, figures, faults found and fixed, and what is better than the Python version (QB 5). | docs/NATIVE_WORK_PLAN.md@ad3758c:298-300 | F-00 | P-13 | |
| REQ-WP-0173 | The native API answers with real HTTP status codes (2xx only for success, 4xx for a refused or invalid request, 5xx for a failure of the box) and one error body (`code`, a translation key, detail). | docs/NATIVE_WORK_PLAN.md@ad3758c:306-308 | F-04 | S-01, S-02…S-0n | Proposed by the developer; used unless the owner says otherwise. |
| REQ-WP-0174 | The native Web UI and `nmb-cli` are written for these status codes and this error body. | docs/NATIVE_WORK_PLAN.md@ad3758c:308-309 | S-10 | U-01, U-02…U-0n | |
| REQ-WP-0175 | The collector protocol stays as it is. | docs/NATIVE_WORK_PLAN.md@ad3758c:309 | S-12 | S-13 | |
| REQ-WP-0176 | The `nmb-cli` JSON command names stay as they are. | docs/NATIVE_WORK_PLAN.md@ad3758c:309 | S-10 | | |
| REQ-WP-0177 | The difference from the Python API is listed in the changelog under Changed. | docs/NATIVE_WORK_PLAN.md@ad3758c:310 | F-04 | P-12 | |
| REQ-WP-0178 | `iperf.schedule_cron` and `iperf.automated_enabled` work in the native build, under the same traffic gate, data caps and hourly limit as the other engines. | docs/NATIVE_WORK_PLAN.md@ad3758c:311-312 | M-04 | M-09, M-03, M-05 | |
| REQ-WP-0179 | M-10 starts with Ookla's own program in the strict sandbox, only after the user has accepted Ookla's terms, and never works around a refusal. | docs/NATIVE_WORK_PLAN.md@ad3758c:313-314 | M-10 | P-04, M-05 | |
| REQ-WP-0180 | A native Ookla engine replaces the program only if it gives the same results on the same line and Ookla's terms allow it. | docs/NATIVE_WORK_PLAN.md@ad3758c:315 | M-10 | P-13 | |
| REQ-WP-0181 | The analyst of M-10 reads Ookla's terms first and writes the finding into the brief. | docs/NATIVE_WORK_PLAN.md@ad3758c:315-316 | M-10 | | |
| REQ-WP-0182 | The base, the init system and the A/B update tool of the own distribution are chosen from I-09's comparison, after the images are analysed. | docs/NATIVE_WORK_PLAN.md@ad3758c:317-319 | I-09 | P-11, P-02 | Owner decision. |
| REQ-WP-0183 | After a test on the Pi 4, the Python version is put back until the native one has replaced it (QB 5). | docs/NATIVE_WORK_PLAN.md@ad3758c:323-324 | P-13 | P-11 | Native testing on the Pi 4 is allowed (line 320); conditions in REQ-WP-0032. |
| REQ-WP-0184 | Updates and images are accepted only with a valid signature. | docs/NATIVE_WORK_PLAN.md@ad3758c:325 | P-02 | P-11 | |
| REQ-WP-0185 | Signing key: an Ed25519 key pair made on the development PC; the private key outside every repository and image, a copy kept offline by the owner; only the public key in the build. | docs/NATIVE_WORK_PLAN.md@ad3758c:325-327 | P-02 | F-00 | Proposed in the source. Signing stays on the owner's PC (REQ-WP-0039). |
| REQ-WP-0186 | The release step signs; the box verifies the signature before anything is written. | docs/NATIVE_WORK_PLAN.md@ad3758c:327-328 | P-02 | | |
| REQ-WP-0187 | The disclaimer (6B.3A) gets its wording (open at `ad3758c`). | docs/NATIVE_WORK_PLAN.md@ad3758c:332 | P-08 | F-01 | Decided by the owner on 2026-11-11 (REQ-NH-0002; QUESTIONS.md, *Merged earlier*); F-01 writes it into the roadmap. Round 4: was "Open owner decision". |
| REQ-WP-0188 | The native repository is open (public), like the Python version. | docs/NATIVE_WORK_PLAN.md@ad3758c:334 | F-00 | P-08 | |
| REQ-WP-0189 | "For everyone": the native build's look pleases and gives a good first impression; it is not a closed product. | docs/NATIVE_WORK_PLAN.md@ad3758c:334-335 | U-01 | P-09, P-12 | Judged by the owner. |
| REQ-WP-0190 | The native repository is published on GitHub under the same account as the Python version, named `NetworkMeasureBox-DePyD`. | docs/NATIVE_WORK_PLAN.md@ad3758c:336-337 | F-00 | | |
| REQ-WP-0191 | The empty GitHub repository exists before the first push; before that push the first commit is checked for personal data and secrets (QB 8). | docs/NATIVE_WORK_PLAN.md@ad3758c:337-338 | F-00 | | |
| REQ-WP-0192 | F-00 starts a list of every dependency with its licence (6B.3A), and each analyst adds to it. | docs/NATIVE_WORK_PLAN.md@ad3758c:339-341 | F-00 | P-08, all packages | |
| REQ-WP-0193 | The licence of the whole is proposed from what the dependency licences require and allow (a built-in copyleft part decides more than a part that is only run). | docs/NATIVE_WORK_PLAN.md@ad3758c:341-342 | P-08 | | |
| REQ-WP-0194 | Until the licence is fixed, no dependency is taken whose licence would rule out an open licence. | docs/NATIVE_WORK_PLAN.md@ad3758c:342-343 | P-08 | F-00 | |
| REQ-WP-0195 | Until the licence is fixed, the repository says that the licence is not decided yet. | docs/NATIVE_WORK_PLAN.md@ad3758c:343 | F-00 | P-08 | |
| REQ-AR-0001 | The box sits inline between the ISP edge (ONT or modem) and the customer LAN, with a router mode (PPPoE, NAT, access point, DHCP/DNS) and a passive inline bridge mode. | docs/ARCHITECTURE.md@ad3758c:7-12; docs/ARCHITECTURE.md@ad3758c:22-26 | N-03 | N-05, N-06, N-07 | Diagram shows only "Mode A/B"; the Python version has more modes (automatic, standalone). N-03 takes the mode list from `netd/modes.py`. |
| REQ-AR-0002 | The network daemon records the active interface and the mode in the database. | docs/ARCHITECTURE.md@ad3758c:27 | N-03 | C-02, C-11 | |
| REQ-AR-0003 | Speed tests run only when the traffic gate reports the line idle. | docs/ARCHITECTURE.md@ad3758c:14; docs/ARCHITECTURE.md@ad3758c:29 | M-03 | M-04 | |
| REQ-AR-0004 | The speed service includes a port-throttle checker. | docs/ARCHITECTURE.md@ad3758c:15 | M-11 | | Diagram also names the Ookla CLI (see REQ-WP-0179). |
| REQ-AR-0005 | Reports are made from the database as HTML/SVG, CSV and SMTP e-mail. | docs/ARCHITECTURE.md@ad3758c:18; docs/ARCHITECTURE.md@ad3758c:30 | R-03 | R-02, R-05 | |
| REQ-AR-0006 | The bot's `/test` and `/report` start a speed test and a report. | docs/ARCHITECTURE.md@ad3758c:19; docs/ARCHITECTURE.md@ad3758c:31; docs/ARCHITECTURE.md@ad3758c:34 | S-11 | M-04, R-03 | |
| REQ-AR-0007 | The network daemon has root privileges for netlink, `pppd` and bridges and writes events. | docs/ARCHITECTURE.md@ad3758c:40 | P-04 | N-03, P-10 | Native: one binary with several units (P-10); each unit keeps its own privileges. |
| REQ-AR-0008 | The probe has `CAP_NET_RAW` only and writes samples. | docs/ARCHITECTURE.md@ad3758c:41 | P-04 | M-01 | Unprivileged ICMP may remove even this (roadmap part REQ-RM-0017). |
| REQ-AR-0009 | The speed service runs unprivileged and writes tests. | docs/ARCHITECTURE.md@ad3758c:42 | P-04 | M-04 | |
| REQ-AR-0010 | The report job runs unprivileged on a timer and does not write the database. | docs/ARCHITECTURE.md@ad3758c:43 | P-04 | R-03 | |
| REQ-AR-0011 | The Telegram bot runs unprivileged and writes alerts. | docs/ARCHITECTURE.md@ad3758c:44 | P-04 | S-11 | |
| REQ-AR-0012 | The web server runs as its own unprivileged user with only the right to bind its port (`CAP_NET_BIND_SERVICE`, or no capability with a socket passed by the service manager) and writes the configuration; network changes and other privileged tasks go to the network daemon or the privileged helper. The Python `CAP_NET_ADMIN` and `CAP_NET_RAW` are not carried over. | docs/ARCHITECTURE.md@ad3758c:45 | P-04 | S-01 | Q-13 settled (developer's choice: Quality Bar 7, only the capabilities a daemon needs; 6B.5). Restated in round 3. |
| REQ-AR-0013 | The snapshot job runs unprivileged on a timer and does not write the live database. | docs/ARCHITECTURE.md@ad3758c:46 | P-04 | C-02 | |
| REQ-AR-0014 | The collector server runs as its own user on the collector's server (an ISP or a specialist), never on the box, with its own database. | docs/ARCHITECTURE.md@ad3758c:47; docs/ARCHITECTURE.md@ad3758c:59 | S-13 | P-10 | |
| REQ-AR-0015 | Scheduled speed tests are started by an in-process scheduler driven by `schedule_cron` (hourly by default); there is no separate systemd timer for them. | docs/ARCHITECTURE.md@ad3758c:49 | M-04 | P-10 | |
| REQ-AR-0016 | The traffic gate reads `/proc/net/dev`, so it sees only traffic that passes through the box. | docs/ARCHITECTURE.md@ad3758c:14; docs/ARCHITECTURE.md@ad3758c:51 | M-03 | | |
| REQ-AR-0017 | The `experimental_stop` switch is off by default; when on, it separates other devices' traffic from the box's own using client-side ports or the router's UPnP counters. | docs/ARCHITECTURE.md@ad3758c:51 | M-03 | N-11 | |
| REQ-AR-0018 | With `experimental_stop` on, a scheduled test stops (cancel) when other devices' traffic appears, and does not run if that traffic cannot be seen. | docs/ARCHITECTURE.md@ad3758c:51 | M-03 | M-04 | |
| REQ-AR-0019 | The probe keeps measuring whether or not `experimental_stop` is on. | docs/ARCHITECTURE.md@ad3758c:51 | M-01 | M-03 | |
| REQ-AR-0020 | Root services: network daemon, DNS blocker, iPerf3 server, OONI, healer and diagnostic logger; the web server and speed service run as `nmb`. | docs/ARCHITECTURE.md@ad3758c:53 | P-04 | P-10 | Native sandbox profiles (6B.5) may lower these; never raise them. |
| REQ-AR-0021 | The web server manages services through the service-control rules (`os/90-nmb.rules`) and uses the privileged helper for other privileged tasks. | docs/ARCHITECTURE.md@ad3758c:53 | P-06 | S-01, P-04 | Round 2: "only" removed (not in the source). |
| REQ-AR-0022 | The host firewall (`nmb-guard`) blocks public addresses from the local service ports. | docs/ARCHITECTURE.md@ad3758c:53 | N-02 | | |
| REQ-AR-0023 | Single-writer contention is handled by WAL and `busy_timeout`. | docs/ARCHITECTURE.md@ad3758c:16; docs/ARCHITECTURE.md@ad3758c:53 | C-02 | | |
| REQ-AR-0024 | Probes batch-insert (one transaction per 10 s) so write-lock time stays tiny. | docs/ARCHITECTURE.md@ad3758c:53 | M-01 | C-02 | |
| REQ-AR-0025 | `nmb-cli` offers the Web UI menus over SSH and JSON commands for scripts. | docs/ARCHITECTURE.md@ad3758c:56 | S-10 | | |
| REQ-AR-0026 | `nmb-cli` takes its menu from the Web UI and calls the same REST API as the browser; it has no measurement logic of its own. | docs/ARCHITECTURE.md@ad3758c:56 | S-10 | U-01 | |
| REQ-AR-0027 | A test fails if a Web UI tab or API call is missing from the `nmb-cli` catalog. | docs/ARCHITECTURE.md@ad3758c:56 | S-10 | U-02…U-0n, S-02…S-0n | |
| REQ-AR-0028 | With `[collector]` on, the box sends from a background thread everything new since the last successful upload to `POST /api/ingest` with a bearer token: speed tests, hourly ping summaries per target, and its identity (a random `box_uid` from `/var/lib/nmb/box_uid`, box name, case number, ISP data). | docs/ARCHITECTURE.md@ad3758c:59 | S-12 | S-13 | |
| REQ-AR-0029 | With `share_diagnostics`, the box adds its events and, hourly, the LLM diagnostics bundle. | docs/ARCHITECTURE.md@ad3758c:59 | S-12 | R-04 | |
| REQ-AR-0030 | With `allow_remote_control`, the box asks `POST /api/poll` every minute for queued commands and reports their results on the next poll; only `speedtest`, `diagnostics` and `upload` run, and both sides refuse anything else. | docs/ARCHITECTURE.md@ad3758c:59 | S-12 | S-13 | |
| REQ-AR-0031 | The box only makes outgoing requests to the collector. | docs/ARCHITECTURE.md@ad3758c:59 | S-12 | | |
| REQ-AR-0032 | The collector keeps one SQLite file and stores tokens and its admin password as hashes in `collector.json` (mode 0600). | docs/ARCHITECTURE.md@ad3758c:59 | S-13 | | |
| REQ-AR-0033 | The collector serves search, box pages, diagnostics history, CSV and the command form behind HTTP Basic auth; admin forms must come from its own origin. | docs/ARCHITECTURE.md@ad3758c:59 | S-13 | | Python uses the standard library only. |
| REQ-AR-0034 | The installed system has no Git data; a self-update runs as root through the privileged helper, from its own copy of the program (the installer replaces the program files during the update). | docs/ARCHITECTURE.md@ad3758c:62 | P-02 | P-06 | |
| REQ-AR-0035 | The Python update downloads the archive of the newest commit and checks every path before unpacking it (Python reference). | docs/ARCHITECTURE.md@ad3758c:62 | P-02 |  | Python behaviour; the native update takes signed releases by file upload or channel instead (REQ-WP-0184, P-02), and keeps the path check. Restated in round 1. |
| REQ-AR-0036 | The update runs the new version's tests, backs up, installs, and checks that the web server answers; if not, it puts the previous program files back. | docs/ARCHITECTURE.md@ad3758c:62 | P-02 | P-03 | |
| REQ-AR-0037 | Update progress is written to `/var/lib/nmb/update_status.json`, which the Web UI and `nmb-cli update` poll, since the web server restarts during the install. | docs/ARCHITECTURE.md@ad3758c:62 | P-02 | U-02…U-0n, S-10 | |
| REQ-AR-0038 | The installed commit is recorded (`/opt/nmb/.source-commit`). | docs/ARCHITECTURE.md@ad3758c:62 | P-02 | C-05 | |
| REQ-AR-0039 | The live database lives in RAM (tmpfs); a consistent snapshot is copied to the SD card with SQLite's online backup API every 15 minutes and when the system shuts down. | docs/ARCHITECTURE.md@ad3758c:16-17; docs/ARCHITECTURE.md@ad3758c:33; docs/ARCHITECTURE.md@ad3758c:66-67; docs/ARCHITECTURE.md@ad3758c:69-71 | C-02 | P-10 | |
| REQ-AR-0040 | When the box boots, the database is restored from the SD snapshot into RAM. | docs/ARCHITECTURE.md@ad3758c:72 | C-02 | | |
| REQ-AR-0041 | Worst-case data loss on a power cut is one snapshot interval (default 15 min). | docs/ARCHITECTURE.md@ad3758c:75 | C-02 | | |
| REQ-AR-0042 | SD writes are about 96 sequential snapshot writes per day instead of thousands of small random page writes. | docs/ARCHITECTURE.md@ad3758c:76 | C-02 | C-08 | Measured, not assumed. |
| REQ-AR-0043 | `/var/log`, `/tmp` and `/var/tmp` are tmpfs; journald uses `Storage=volatile`. | docs/ARCHITECTURE.md@ad3758c:77 | P-11 | P-10 | |
| REQ-AR-0044 | Swap is disabled; the ext4 root uses `noatime,commit=600`. | docs/ARCHITECTURE.md@ad3758c:70; docs/ARCHITECTURE.md@ad3758c:78 | P-11 | | |
| REQ-AR-0045 | The program files are read-only at run time. | docs/ARCHITECTURE.md@ad3758c:83 | P-10 | P-04 | Native path is set by F-04/P-10. |
| REQ-AR-0046 | The configuration is the INI file `/etc/nmb/nmb.conf`, which the Web UI edits. | docs/ARCHITECTURE.md@ad3758c:87 | C-01 | F-04, S-02…S-0n | |
| REQ-AR-0047 | The hot database is in `/var/lib/nmb/live/` (tmpfs); snapshots and generated reports are in `/var/lib/nmb/persist/` and `persist/reports/` (SD). | docs/ARCHITECTURE.md@ad3758c:88-90 | C-02 | R-03 | |
| REQ-AR-0048 | Pidfiles, sockets and mode state are in `/run/nmb/` (tmpfs). | docs/ARCHITECTURE.md@ad3758c:91 | P-10 | N-03 | |
| REQ-AR-0049 | Logs are in `/var/log/nmb/` (tmpfs). | docs/ARCHITECTURE.md@ad3758c:92 | C-03 | | |
| REQ-AR-0050 | The Telegram bot runs alongside the web server, gives chat management and proactive incident alerts, and needs no external dependencies. | docs/ARCHITECTURE.md@ad3758c:112 | S-11 | | |
| REQ-AR-0051 | Bot diagnostics commands: `/status`, `/ping`, `/test [ookla\|cloudflare\|fastcom\|rtr\|iperf]`, `/report` (HTML dossier), `/ai` (Markdown bundle), `/mode`. | docs/ARCHITECTURE.md@ad3758c:115 | S-11 | R-03, R-04, M-04 | |
| REQ-AR-0052 | `/vpn` shows WireGuard and OpenVPN status and client counts. | docs/ARCHITECTURE.md@ad3758c:116-117 | S-11 | N-13, N-14 | |
| REQ-AR-0053 | `/vpn clients` lists connected peers with address, latest handshake and RX/TX traffic. | docs/ARCHITECTURE.md@ad3758c:118 | S-11 | N-13, N-14 | |
| REQ-AR-0054 | `/vpn add <name> [wg\|ovpn]` creates a client peer and sends its QR code image and `.conf` file when "secrets over Telegram" is on for it (REQ-AR-0079). | docs/ARCHITECTURE.md@ad3758c:119 | S-11 | N-13, N-14 | QR and profile hold the private key: Q-15 answered by the owner (2026-11-11): C. |
| REQ-AR-0079 | Secrets over Telegram, VPN profiles: a profile's QR code and `.conf` (they hold the client's private key) go through Telegram only when the user has switched "secrets over Telegram" on for it (per item; off by default). Then: the confirmation tap, a warning, the bot deletes its QR message after a few minutes, and the VPN peer can be withdrawn with one tap. | docs/ARCHITECTURE.md@ad3758c:119 | S-11 | N-13, N-14 | New: the owner's answer to Q-15 (2026-11-11). |
| REQ-AR-0055 | `/vpn qr <name>` makes and sends the QR code image. | docs/ARCHITECTURE.md@ad3758c:120 | S-11 | N-13 | |
| REQ-AR-0056 | `/vpn conf <name>` sends the configuration profile. | docs/ARCHITECTURE.md@ad3758c:121 | S-11 | N-13, N-14 | |
| REQ-AR-0057 | `/vpn start [wg\|ovpn]` and `/vpn stop [wg\|ovpn]` start and stop the VPN service. | docs/ARCHITECTURE.md@ad3758c:122 | S-11 | N-13, N-14 | |
| REQ-AR-0058 | `/vpn del <name>` revokes a peer. | docs/ARCHITECTURE.md@ad3758c:123 | S-11 | N-13, N-14 | |
| REQ-AR-0059 | `/adblock` shows queries today, blocked count, blocked share and list domain count. | docs/ARCHITECTURE.md@ad3758c:124-125 | S-11 | N-12 | |
| REQ-AR-0060 | `/adblock on` and `/adblock off` switch the DNS blocker. | docs/ARCHITECTURE.md@ad3758c:126 | S-11 | N-12 | |
| REQ-AR-0061 | `/adblock block <domain>` and `/adblock unblock <domain>` add a domain to the block or allow list. | docs/ARCHITECTURE.md@ad3758c:127 | S-11 | N-12 | |
| REQ-AR-0062 | `/adblock update` rebuilds the block list in the background. | docs/ARCHITECTURE.md@ad3758c:128 | S-11 | N-12 | |
| REQ-AR-0063 | `/adblock guide` gives the instructions for setting the box as DNS in the home router's DHCP. | docs/ARCHITECTURE.md@ad3758c:129 | S-11 | N-12 | |
| REQ-AR-0064 | `/heal audit` runs the component health audit. | docs/ARCHITECTURE.md@ad3758c:130-131 | S-11 | P-01 | |
| REQ-AR-0065 | `/heal fix [tier]` runs the repair escalation (tiers 1 to 3). | docs/ARCHITECTURE.md@ad3758c:132 | S-11 | P-01 | |
| REQ-AR-0066 | `/heal updates` checks for a verified stable release. | docs/ARCHITECTURE.md@ad3758c:133 | S-11 | P-02 | |
| REQ-AR-0067 | `/heal apply` installs an update atomically with a checkpoint before and automatic rollback. | docs/ARCHITECTURE.md@ad3758c:134 | S-11 | P-02 | Native: only signed updates (REQ-WP-0184). |
| REQ-AR-0068 | `/tasks` (or `/top`) shows tasks and load; `/kill <pid> [signal]` sends a signal. | docs/ARCHITECTURE.md@ad3758c:135 | S-11 | C-09 | Q-14 (developer's choice, round 4): the scope stays Python's effective one, NMB's own processes (README.md@ad3758c:422, :452; the Python bot as user `nmb` could signal no other); where such a process runs under another user than the bot, the privileged helper sends the signal after checking it against its allow-list of NMB's own units (never `sudo` or a shell). Every signal needs the confirmation tap of the owner's decision of 2026-11-11, KILL also the typed word; PID 1 and the bot stay protected (REQ-RE-0280). A wider scope would be an owner question. |
| REQ-AR-0069 | `/wifi` shows telemetry, `/wifi scan` nearby access points, `/wifi connect <ssid> [pw]` joins a network (a password in the chat only as REQ-AR-0080 allows), `/wifi saved` lists saved networks. | docs/ARCHITECTURE.md@ad3758c:136 | S-11 | M-14, N-07, N-08 | Heading says "Guest Wi-Fi" but no guest command is listed. Password in chat: Q-15 answered by the owner (2026-11-11): C. |
| REQ-AR-0080 | Secrets over Telegram, Wi-Fi passwords: `/wifi connect` takes a password in the chat only when the user has switched "secrets over Telegram" on for it (per item; off by default). Then: the confirmation tap, a warning, and the bot deletes the user's password message at once. | docs/ARCHITECTURE.md@ad3758c:136 | S-11 | N-07, N-08 | New: the owner's answer to Q-15 (2026-11-11). |
| REQ-AR-0070 | `/backup list`, `/backup create [label]` and `/backup send <filename>` (the archive sent in Telegram). | docs/ARCHITECTURE.md@ad3758c:137 | S-11 | P-03 | Out of date: Python sends the archive encrypted with the owner's passphrase, plain only with `plain` (`nmb/bot/commands.py@ad3758c:1627-1630`). The native build keeps the encrypted default. |
| REQ-AR-0071 | `/drivers` shows hardware and drivers; `/reboot now` reboots the box. | docs/ARCHITECTURE.md@ad3758c:138 | S-11 | N-17, P-10 | |
| REQ-AR-0072 | Alerts are stateful and deduplicated in the database (alert table). | docs/ARCHITECTURE.md@ad3758c:140-141 | S-11 | C-02 | |
| REQ-AR-0073 | Outage alert: fires when all probe targets become unreachable, alerts on start and on resolution with the total outage duration. | docs/ARCHITECTURE.md@ad3758c:142 | S-11 | M-02 | |
| REQ-AR-0074 | Packet-loss alert: fires when 5-minute rolling loss exceeds the configured threshold (example 10.0 %), resolves when normal again. | docs/ARCHITECTURE.md@ad3758c:143 | S-11 | M-01 | |
| REQ-AR-0075 | Service-failure alert: watches the core services (network daemon, web server, probe, scheduler); when one fails it alerts the admin and starts tier-1 self-healing. | docs/ARCHITECTURE.md@ad3758c:144 | S-11 | P-01 | Python fault: it watches a unit `nmb-scheduler` that does not exist (scheduler is inside `nmb-web`, line 49; `nmb/bot/alerts.py@ad3758c:181`, no such file in `systemd/`). For F-03; native watches its real units. |
| REQ-AR-0076 | Overheating alert: reads the SoC thermal sensor and warns above 82 °C. | docs/ARCHITECTURE.md@ad3758c:145 | S-11 | C-09, N-17 | Source path `/sys/class/thermal/thermal_zone0/temp` is Pi-specific; native finds the sensor (QB 9 location/platform aware). Python fires at `>= 82.0` (`nmb/bot/alerts.py@ad3758c:232`); boundary fixed in `spec/`. |
| REQ-AR-0077 | Low-disk alert: warns when free root disk space drops below 500 MB. | docs/ARCHITECTURE.md@ad3758c:146 | S-11 | C-08 | |
| REQ-AR-0078 | WAN failover alert: alerts when the primary WAN fails and traffic fails over to a secondary WAN or mobile uplink. | docs/ARCHITECTURE.md@ad3758c:147 | S-11 | N-10, N-03 | |
| REQ-DV-0001 | Development, builds and tests run inside WSL2 (Ubuntu), not in plain Windows; the build and gate scripts are written for a Linux shell. | docs/DEV_SETUP_WINDOWS.md@ad3758c:5; docs/DEV_SETUP_WINDOWS.md@ad3758c:9-11; docs/DEV_SETUP_WINDOWS.md@ad3758c:15; docs/DEV_SETUP_WINDOWS.md@ad3758c:207 | F-00 |  | Environment recorded per batch (Debian 13 box, Windows + WSL2 Ubuntu 24.04 PC). |
| REQ-DV-0002 | The repositories live in the WSL file system (not under `/mnt/c`), with `core.autocrlf false`, so line endings and file permissions stay correct. | docs/DEV_SETUP_WINDOWS.md@ad3758c:13; docs/DEV_SETUP_WINDOWS.md@ad3758c:83; docs/DEV_SETUP_WINDOWS.md@ad3758c:89; docs/DEV_SETUP_WINDOWS.md@ad3758c:210 | F-00 | | F-00 can add a `.gitattributes` check for LF. |
| REQ-DV-0003 | The box's Tailscale address and the owner's public addresses (home line, mobile) are never written into the repository. | docs/DEV_SETUP_WINDOWS.md@ad3758c:43; docs/DEV_SETUP_WINDOWS.md@ad3758c:198 | F-00 | P-13 | Part of the scan of REQ-WP-0051. |
| REQ-DV-0004 | Passwords are never put on a command line, in a script or in a log; `sudo` on the box gets it on standard input from a private file (umask 077) that is deleted afterwards; key login is preferred for agents. | docs/DEV_SETUP_WINDOWS.md@ad3758c:54; docs/DEV_SETUP_WINDOWS.md@ad3758c:73-79 | F-00 | P-11, P-13 | process rule: checked by the orchestrator; no test expected beyond the secret scan of REQ-RM-1076 in F-00's gate, which finds a password written into a script or log in the repository. |
| REQ-DV-0005 | A GitHub token is pasted at the prompt, never given as an argument; it has only the `repo` and `workflow` scopes; it is never written into a remote URL or `.git/config`. | docs/DEV_SETUP_WINDOWS.md@ad3758c:94-99 | F-00 |  | process rule: checked by the orchestrator; no test expected. |
| REQ-DV-0006 | All unit tests pass before every commit; they need no network and no root. | docs/DEV_SETUP_WINDOWS.md@ad3758c:107; docs/DEV_SETUP_WINDOWS.md@ad3758c:212 | F-00 | F-06 | |
| REQ-DV-0007 | The live checks (self-test) run on the box, not on the PC. | docs/DEV_SETUP_WINDOWS.md@ad3758c:107 | P-07 | | |
| REQ-DV-0008 | Tests stay correct when the wall clock jumps (fake clock, durations from monotonic time); for a test that fails now and then, the clock-jump check (must show 0 jumps in 30 s) comes first. | docs/DEV_SETUP_WINDOWS.md@ad3758c:109-119 | F-06 | C-05, N-16 | A Python DynDNS test failed about 1 run in 20 under wall-clock jumps (2026-11-10): candidate for F-03. |
| REQ-DV-0009 | The native toolchain of section 5A is installed on the PC (Go, C and cross compilers, QEMU system and user emulators with UEFI firmware, firmware and image tools, analysers), with the analysers pinned: `staticcheck` 2024.1.1, `gosec` v2.21.4, `govulncheck` v1.1.3, on `PATH` through `~/go/bin`. | docs/DEV_SETUP_WINDOWS.md@ad3758c:123-132; docs/DEV_SETUP_WINDOWS.md@ad3758c:134 | F-00 | | These versions match the Go version now installed on the PC (source names a specific site value; not copied); when F-00 fixes Go 1.23+ it may need newer pins. |
| REQ-DV-0010 | Go builds use `CGO_ENABLED=0` for `amd64`, `arm64`, `arm` v7, `mips` and `mipsle` (soft-float), each runnable under `qemu-*-static`. | docs/DEV_SETUP_WINDOWS.md@ad3758c:134 | F-00 | NEW: X2-01 Step-2 router tiers | |
| REQ-DV-0011 | C tooling for the router tiers: `musl-gcc -static`, cross GCC for AArch64, ARMv7, MIPS big and little endian, clang with AddressSanitizer and UndefinedBehaviorSanitizer, `cppcheck`, `clang-tidy`. | docs/DEV_SETUP_WINDOWS.md@ad3758c:134 | NEW: X2-01 Step-2 router tiers | F-07 | Step 2. |
| REQ-DV-0012 | The test suite is never run as root (or as the service user) on a working box; on the box tests run as `nobody` on a copy (as the updater does), and no test writes the real configuration. | docs/DEV_SETUP_WINDOWS.md@ad3758c:142-144 | P-02 | F-06, P-07 | Python fault: tests that save settings without a path overwrote the box's `nmb.conf` (2026-11-10); for F-03. |
| REQ-DV-0013 | A change is installed on the box the way users get it (the installer), never by copying single files, except a quick try that is reported as such. | docs/DEV_SETUP_WINDOWS.md@ad3758c:145-146 | P-11 | P-02, P-13 | |
| REQ-DV-0014 | The installed version is recorded as the full commit id (40 characters); a short id is dropped and the update check then cannot know the installed version. | docs/DEV_SETUP_WINDOWS.md@ad3758c:152; docs/DEV_SETUP_WINDOWS.md@ad3758c:155 | P-02 | C-05 | |
| REQ-DV-0015 | An install survives the restart of the network daemon and the loss of the remote session (in router mode PPPoE redials and SSH may stall about 20 s): it runs detached and is waited for. | docs/DEV_SETUP_WINDOWS.md@ad3758c:157; docs/DEV_SETUP_WINDOWS.md@ad3758c:164-165 | P-02 | N-03 | |
| REQ-DV-0016 | The install log is kept out of the `/var/log` RAM disk (48 MB on the box) on persistent storage; a full RAM disk must not stop an install halfway. | docs/DEV_SETUP_WINDOWS.md@ad3758c:157; docs/DEV_SETUP_WINDOWS.md@ad3758c:164; docs/DEV_SETUP_WINDOWS.md@ad3758c:166 | P-02 | C-08 | A full `/var/log` stopped a Python install halfway on 2026-11-10. |
| REQ-DV-0017 | Temporary copies of an install are removed from `/tmp` (a 256 MB tmpfs on the box; leftovers once filled it). | docs/DEV_SETUP_WINDOWS.md@ad3758c:166; docs/DEV_SETUP_WINDOWS.md@ad3758c:169 | P-02 | C-08 | |
| REQ-DV-0018 | After every install on the box the self-test runs and shows no failure. | docs/DEV_SETUP_WINDOWS.md@ad3758c:169-173; docs/DEV_SETUP_WINDOWS.md@ad3758c:213 | P-07 | P-02, P-13 | |
| REQ-DV-0019 | Tests that cut the box's own link (router mode, USB tethering, cable moves) run only through the guided scenarios, which run as a service on the box. | docs/DEV_SETUP_WINDOWS.md@ad3758c:177 | P-07 | N-03, N-10 | |
| REQ-DV-0020 | A guided scenario keeps the settings to restore, stops and restores when the remote link (Tailscale) is offline for 2 minutes, and restores after 45 minutes whatever happens. | docs/DEV_SETUP_WINDOWS.md@ad3758c:177 | P-07 | N-15, N-18 | |
| REQ-DV-0021 | A manual restore command for the scenarios exists (`nmb-selftest restore`). | docs/DEV_SETUP_WINDOWS.md@ad3758c:184 | P-07 | S-10 | |
| REQ-DV-0022 | The agent's notes (which may hold how to reach the box) are never put into the repository. | docs/DEV_SETUP_WINDOWS.md@ad3758c:188; docs/DEV_SETUP_WINDOWS.md@ad3758c:193 | F-00 |  | Work plan line 82 allows the parts without private data in the native repository; the gate scan decides per file. |
| REQ-DV-0023 | No secrets in the repository, in commits or in logs: tokens, passwords, keys, the Wi-Fi password. | docs/DEV_SETUP_WINDOWS.md@ad3758c:197 | F-00 | C-03, C-04, P-13 | |
| REQ-DV-0024 | Scheduled Ookla runs at most once an hour (a test a person starts obeys the refusal pause only, owner decision 2026-11-11, REQ-TS-0302); no code to dodge its rate limits. | docs/DEV_SETUP_WINDOWS.md@ad3758c:199 | M-05 | M-10 | Round 4: qualified by the owner's decision of 2026-11-11 (Q-28). |
| REQ-DV-0025 | Every user-visible change is in `CHANGELOG.md`. | docs/DEV_SETUP_WINDOWS.md@ad3758c:200 | P-12 | F-00 | |
| REQ-DV-0026 | Native documentation: the user manual and the command reference in every UI language from one source; developer documents (briefs, test record, contribution rules, development setup) in English as their one source. The Python pair English and Hungarian is not carried over. | docs/DEV_SETUP_WINDOWS.md@ad3758c:200 | P-12 | | Rule of the Python repository restated for the native build. Q-09: the manual's 11 languages are the owner's text (6B.7); the command reference's 11 languages (Quality Bar 9, no language is a default) and the English-only developer documents are developer's choices (round 4). Restated in round 3. |

## Coverage

Each range includes the blank line(s) that end it. Every line of each file is in exactly one row.

### docs/NATIVE_WORK_PLAN.md (1-343)

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-2 | heading | context: document title |
| 3-5 | paragraph | context: status (plan only, no native code) and purpose of the document |
| 6-8 | paragraph | REQ-WP-0001 |
| 9-14 | paragraph | REQ-WP-0002, REQ-WP-0003, REQ-WP-0025, REQ-WP-0024, REQ-WP-0034; line 13 (images and firmware taken apart, analysed and used): REQ-WP-0143 to REQ-WP-0161 (section 5); context: "efficient and as free of faults as possible" is a goal made checkable by the gate (REQ-WP-0043 to -0054) |
| 15-16 | heading | context: section 1 title |
| 17 | bullet | REQ-WP-0004 |
| 18 | bullet | REQ-WP-0005, REQ-WP-0006 |
| 19-20 | bullet | REQ-WP-0007 |
| 21-22 | bullet | REQ-WP-0008, REQ-WP-0009 |
| 23-24 | bullet | REQ-WP-0010 |
| 25-26 | heading | context: section 2 title |
| 27-28 | heading | REQ-WP-0196 (2.1 title: each role a separate agent with its own context) |
| 29-30 | table row | context: table header and alignment row |
| 31 | table row | REQ-WP-0011 |
| 32 | table row | REQ-WP-0012 |
| 33 | table row | REQ-WP-0013 |
| 34 | table row | REQ-WP-0014 |
| 35 | table row | REQ-WP-0015 |
| 36-37 | table row | REQ-WP-0016 |
| 38-40 | paragraph | REQ-WP-0017 |
| 41-42 | heading | context: 2.2 title |
| 43-44 | bullet | REQ-WP-0018 |
| 45-46 | bullet | REQ-WP-0019 |
| 47-49 | bullet | REQ-WP-0020, REQ-WP-0021 |
| 50-52 | bullet | REQ-WP-0022, REQ-WP-0023 |
| 53-55 | bullet | REQ-WP-0024 |
| 56-57 | bullet | REQ-WP-0025 |
| 58 | bullet | REQ-WP-0026 |
| 59 | bullet | REQ-WP-0027 |
| 60-61 | bullet | REQ-WP-0028, REQ-WP-0029 |
| 62 | bullet | REQ-WP-0030, REQ-WP-0031 |
| 63-67 | bullet | REQ-WP-0032, REQ-WP-0033 |
| 68-69 | heading | context: 2.3 title |
| 70-76 | paragraph | REQ-WP-0034, REQ-WP-0035, REQ-WP-0036 |
| 77-78 | heading | context: 2.3A title |
| 79-80 | paragraph | REQ-WP-0036 |
| 81 | bullet | REQ-WP-0037 |
| 82 | bullet | REQ-WP-0038 |
| 83 | bullet | REQ-WP-0039 |
| 84 | bullet | REQ-WP-0040 |
| 85 | bullet | REQ-WP-0041, REQ-WP-0035 |
| 86 | bullet | context: costs are the owner's to weigh; the plan does not assume a server |
| 87-88 | bullet | REQ-WP-0042 |
| 89-90 | heading | context: 2.4 title |
| 91-97 | paragraph | REQ-WP-0043, REQ-WP-0044, REQ-WP-0045, REQ-WP-0046, REQ-WP-0047, REQ-WP-0021, REQ-WP-0048, REQ-WP-0049, REQ-WP-0050, REQ-WP-0051, REQ-WP-0052, REQ-WP-0053, REQ-WP-0054, REQ-WP-0017 |
| 98-99 | heading | context: section 3 title |
| 100-102 | paragraph | REQ-WP-0055 |
| 103-104 | heading | REQ-WP-0056 |
| 105-106 | table row | context: table header and alignment row |
| 107 | table row | REQ-WP-0057 |
| 108 | table row | REQ-WP-0058 |
| 109 | table row | REQ-WP-0059 |
| 110 | table row | REQ-WP-0060 |
| 111 | table row | REQ-WP-0061 |
| 112 | table row | REQ-WP-0062 |
| 113 | table row | REQ-WP-0063 |
| 114-115 | table row | REQ-WP-0064 |
| 116-117 | heading | context: 3.1 title (Core runs in parallel after F; ordering in REQ-WP-0137) |
| 118-119 | table row | context: table header and alignment row |
| 120 | table row | REQ-WP-0065 |
| 121 | table row | REQ-WP-0066 |
| 122 | table row | REQ-WP-0067 |
| 123 | table row | REQ-WP-0068 |
| 124 | table row | REQ-WP-0069 |
| 125 | table row | REQ-WP-0070 |
| 126 | table row | REQ-WP-0071 |
| 127 | table row | REQ-WP-0072 |
| 128 | table row | REQ-WP-0073 |
| 129 | table row | REQ-WP-0074 |
| 130-131 | table row | REQ-WP-0075 |
| 132-133 | heading | context: 3.2 title |
| 134-135 | table row | context: table header and alignment row |
| 136 | table row | REQ-WP-0076 |
| 137 | table row | REQ-WP-0077 |
| 138 | table row | REQ-WP-0078 |
| 139 | table row | REQ-WP-0079 |
| 140 | table row | REQ-WP-0080 |
| 141 | table row | REQ-WP-0081 |
| 142 | table row | REQ-WP-0082 |
| 143 | table row | REQ-WP-0083 |
| 144 | table row | REQ-WP-0084 |
| 145 | table row | REQ-WP-0085 |
| 146 | table row | REQ-WP-0086 |
| 147 | table row | REQ-WP-0087 |
| 148 | table row | REQ-WP-0088 |
| 149-150 | table row | REQ-WP-0089 |
| 151-152 | heading | context: 3.3 title |
| 153-154 | table row | context: table header and alignment row |
| 155 | table row | REQ-WP-0090 |
| 156 | table row | REQ-WP-0091 |
| 157 | table row | REQ-WP-0092 |
| 158 | table row | REQ-WP-0093 |
| 159 | table row | REQ-WP-0094 |
| 160 | table row | REQ-WP-0095 |
| 161 | table row | REQ-WP-0096 |
| 162 | table row | REQ-WP-0097 |
| 163 | table row | REQ-WP-0098 |
| 164 | table row | REQ-WP-0099 |
| 165 | table row | REQ-WP-0100 |
| 166 | table row | REQ-WP-0101 |
| 167 | table row | REQ-WP-0102 |
| 168 | table row | REQ-WP-0103 |
| 169 | table row | REQ-WP-0104 |
| 170 | table row | REQ-WP-0105 |
| 171 | table row | REQ-WP-0106 |
| 172-173 | table row | REQ-WP-0107 |
| 174-175 | heading | context: 3.4 title |
| 176-177 | table row | context: table header and alignment row |
| 178 | table row | REQ-WP-0108 |
| 179 | table row | REQ-WP-0109 |
| 180 | table row | REQ-WP-0110 |
| 181 | table row | REQ-WP-0111 |
| 182-183 | table row | REQ-WP-0112 |
| 184-185 | heading | context: 3.5 title |
| 186-187 | table row | context: table header and alignment row |
| 188 | table row | REQ-WP-0113 |
| 189 | table row | REQ-WP-0114 |
| 190 | table row | REQ-WP-0115 |
| 191 | table row | REQ-WP-0116 |
| 192 | table row | REQ-WP-0117 |
| 193 | table row | REQ-WP-0118 |
| 194 | table row | REQ-WP-0119 |
| 195-196 | table row | REQ-WP-0120 |
| 197-198 | heading | context: 3.6 title |
| 199-200 | table row | context: table header and alignment row |
| 201 | table row | REQ-WP-0121 |
| 202 | table row | REQ-WP-0122 |
| 203 | table row | REQ-WP-0123 |
| 204 | table row | REQ-WP-0124 |
| 205 | table row | REQ-WP-0125 |
| 206 | table row | REQ-WP-0126 |
| 207 | table row | REQ-WP-0127 |
| 208 | table row | REQ-WP-0128 |
| 209 | table row | REQ-WP-0129 |
| 210 | table row | REQ-WP-0130 |
| 211 | table row | REQ-WP-0131 |
| 212 | table row | REQ-WP-0132 |
| 213-214 | table row | REQ-WP-0133 |
| 215-216 | paragraph | REQ-WP-0134 |
| 217-219 | paragraph | REQ-WP-0135 |
| 220-221 | heading | context: section 4 title |
| 222 | bullet | REQ-WP-0136 |
| 223 | bullet | REQ-WP-0137 |
| 224 | bullet | REQ-WP-0138 |
| 225-226 | bullet | REQ-WP-0139 |
| 227-228 | heading | context: section 5 title |
| 229-233 | paragraph | REQ-WP-0140, REQ-WP-0141, REQ-WP-0142 |
| 234-235 | heading | context: 5.1 title |
| 236-237 | table row | context: table header and alignment row |
| 238 | table row | REQ-WP-0143 |
| 239 | table row | REQ-WP-0144 |
| 240 | table row | REQ-WP-0145 |
| 241 | table row | REQ-WP-0146 |
| 242-243 | table row | REQ-WP-0147 |
| 244-246 | paragraph | REQ-WP-0148, REQ-WP-0197 |
| 247-248 | heading | context: 5.2 title |
| 249-259 | paragraph | REQ-WP-0149, REQ-WP-0150, REQ-WP-0151 |
| 260-261 | heading | context: 5.3 title |
| 262 | bullet | REQ-WP-0152 |
| 263-265 | bullet | REQ-WP-0153, REQ-WP-0154 |
| 266-268 | bullet | REQ-WP-0155 |
| 269-271 | bullet | REQ-WP-0156 |
| 272-273 | bullet | REQ-WP-0157, REQ-WP-0158, REQ-WP-0159 |
| 274-275 | bullet | REQ-WP-0160 |
| 276-277 | bullet | REQ-WP-0161 |
| 278-280 | paragraph | REQ-WP-0140 |
| 281-282 | heading | context: 5.4 title |
| 283-284 | table row | context: table header and alignment row |
| 285 | table row | REQ-WP-0162 |
| 286 | table row | REQ-WP-0163 |
| 287 | table row | REQ-WP-0164 |
| 288 | table row | REQ-WP-0165 |
| 289 | table row | REQ-WP-0166 |
| 290 | table row | REQ-WP-0167 |
| 291 | table row | REQ-WP-0168 |
| 292 | table row | REQ-WP-0169 |
| 293 | table row | REQ-WP-0170 |
| 294-295 | table row | REQ-WP-0171 |
| 296-297 | heading | context: section 6 title |
| 298-301 | paragraph | REQ-WP-0172 |
| 302-303 | heading | context: section 7 title |
| 304-305 | paragraph | context: date of the owner's answers |
| 306-310 | bullet | REQ-WP-0173, REQ-WP-0174, REQ-WP-0175, REQ-WP-0176, REQ-WP-0177 |
| 311-312 | bullet | REQ-WP-0178 |
| 313-316 | bullet | REQ-WP-0179, REQ-WP-0180, REQ-WP-0181 |
| 317-319 | bullet | REQ-WP-0182 |
| 320-324 | bullet | REQ-WP-0032, REQ-WP-0183 |
| 325-329 | bullet | REQ-WP-0184, REQ-WP-0185, REQ-WP-0186 |
| 330-331 | paragraph | context: "Still open:" label |
| 332 | bullet | REQ-WP-0187 |
| 333-338 | bullet | REQ-WP-0010, REQ-WP-0188, REQ-WP-0189, REQ-WP-0190, REQ-WP-0191 |
| 339-343 | bullet | REQ-WP-0192, REQ-WP-0193, REQ-WP-0194, REQ-WP-0195 |

### docs/ARCHITECTURE.md (1-149)

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-2 | heading | context: document title |
| 3-4 | heading | context: section 1 title |
| 5-36 | code block | REQ-AR-0001, REQ-AR-0002, REQ-AR-0003, REQ-AR-0004, REQ-AR-0005, REQ-AR-0006, REQ-AR-0016, REQ-AR-0023, REQ-AR-0039; context: lines 13, 20 name the probe and web server boxes (their requirements are in the rows below) |
| 37 | heading | context: "Process model" title |
| 38-39 | table row | context: table header and alignment row ("Phase" column is the Python build order, no requirement) |
| 40 | table row | REQ-AR-0007 |
| 41 | table row | REQ-AR-0008 |
| 42 | table row | REQ-AR-0009 |
| 43 | table row | REQ-AR-0010 |
| 44 | table row | REQ-AR-0011 |
| 45 | table row | REQ-AR-0012 |
| 46 | table row | REQ-AR-0013 |
| 47-48 | table row | REQ-AR-0014 |
| 49-50 | paragraph | REQ-AR-0015 |
| 51-52 | paragraph | REQ-AR-0016, REQ-AR-0017, REQ-AR-0018, REQ-AR-0019 |
| 53-54 | paragraph | REQ-AR-0020, REQ-AR-0021, REQ-AR-0022, REQ-AR-0023, REQ-AR-0024 |
| 55 | heading | context: "Terminal interface" title |
| 56-57 | paragraph | REQ-AR-0025, REQ-AR-0026, REQ-AR-0027 |
| 58 | heading | context: "Central collector" title |
| 59-60 | paragraph | REQ-AR-0014, REQ-AR-0028, REQ-AR-0029, REQ-AR-0030, REQ-AR-0031, REQ-AR-0032, REQ-AR-0033 |
| 61 | heading | context: "Updates without Git" title |
| 62-63 | paragraph | REQ-AR-0034, REQ-AR-0035, REQ-AR-0036, REQ-AR-0037, REQ-AR-0038 |
| 64-65 | heading | context: section 2 title |
| 66-68 | paragraph | REQ-AR-0039 |
| 69-74 | code block | REQ-AR-0039, REQ-AR-0040, REQ-AR-0044 |
| 75 | bullet | REQ-AR-0041 |
| 76 | bullet | REQ-AR-0042 |
| 77 | bullet | REQ-AR-0043 |
| 78-79 | bullet | REQ-AR-0044 |
| 80-81 | heading | context: section 3 title |
| 82 | code block | context: code fence |
| 83 | code block | REQ-AR-0045 |
| 84-86 | code block | context: Python package folders and entry-point wrappers (native layout is a single binary, P-10) |
| 87 | code block | REQ-AR-0046 |
| 88-90 | code block | REQ-AR-0047 |
| 91 | code block | REQ-AR-0048 |
| 92 | code block | REQ-AR-0049 |
| 93-94 | code block | context: code fence and blank line |
| 95-103 | paragraph, code block | context: layout of the Python repository; the native repository layout is set by F-00 (REQ-WP-0057) |
| 104-105 | heading | context: section 4 title |
| 106-107 | paragraph | context: points to the roadmap; its requirements are in the roadmap parts (REQ-RM-0007, -0008, -0010) |
| 108-109 | paragraph | context: current state (only the Pi 4 build exists, no router port) and pointer to the porting plan, whose requirements are in its own part; router models are step 2 (NEW: X2-01) |
| 110-111 | heading | context: section 5 title |
| 112-113 | paragraph | REQ-AR-0050 |
| 114 | heading | context: "Interactive Commands" title |
| 115 | bullet | REQ-AR-0051 |
| 116 | bullet | REQ-AR-0052 |
| 117 | bullet | REQ-AR-0052 |
| 118 | bullet | REQ-AR-0053 |
| 119 | bullet | REQ-AR-0054, REQ-AR-0079 |
| 120 | bullet | REQ-AR-0055 |
| 121 | bullet | REQ-AR-0056 |
| 122 | bullet | REQ-AR-0057 |
| 123 | bullet | REQ-AR-0058 |
| 124 | bullet | REQ-AR-0059 |
| 125 | bullet | REQ-AR-0059 |
| 126 | bullet | REQ-AR-0060 |
| 127 | bullet | REQ-AR-0061 |
| 128 | bullet | REQ-AR-0062 |
| 129 | bullet | REQ-AR-0063 |
| 130 | bullet | REQ-AR-0064 |
| 131 | bullet | REQ-AR-0064 |
| 132 | bullet | REQ-AR-0065 |
| 133 | bullet | REQ-AR-0066 |
| 134 | bullet | REQ-AR-0067 |
| 135 | bullet | REQ-AR-0068 |
| 136 | bullet | REQ-AR-0069, REQ-AR-0080 |
| 137 | bullet | REQ-AR-0070 |
| 138-139 | bullet | REQ-AR-0071 |
| 140 | heading | REQ-AR-0072 |
| 141 | paragraph | REQ-AR-0072 |
| 142 | bullet | REQ-AR-0073 |
| 143 | bullet | REQ-AR-0074 |
| 144 | bullet | REQ-AR-0075 |
| 145 | bullet | REQ-AR-0076 |
| 146 | bullet | REQ-AR-0077 |
| 147-149 | bullet | REQ-AR-0078 |

### docs/DEV_SETUP_WINDOWS.md (1-214)

| Lines | Kind | Ids or "context: <why no requirement>" |
| :--- | :--- | :--- |
| 1-2 | heading | context: document title |
| 3-4 | paragraph | context: purpose (move development to a Windows PC) and link to the Hungarian copy |
| 5-6 | paragraph | REQ-DV-0001; context: reaching the box over Tailscale is the developer's access path, no native behaviour |
| 7-8 | heading | context: "Why WSL2" title |
| 9-10 | paragraph | REQ-DV-0001 |
| 11 | bullet | REQ-DV-0001 (Python scripts are bash) |
| 12 | bullet | context: what the Python tests use (`fcntl`, `AF_PACKET`, namespaces); the native test bed is F-06 |
| 13-14 | bullet | REQ-DV-0002 |
| 15-16 | paragraph | REQ-DV-0001 |
| 17-18 | heading | context: section 1 title |
| 19-20 | paragraph | context: install instruction |
| 21-24 | code block | context: WSL install command |
| 25-26 | paragraph | context: install instruction |
| 27-31 | code block | context: Python development packages (native tools in REQ-DV-0009) |
| 32-33 | paragraph | context: the Python RTR client build on the PC; the box builds it during the Python installer (native RTR is M-08) |
| 34-35 | paragraph | context: the agent tool is installed inside WSL (agent tooling, no native behaviour) |
| 36-39 | code block | context: agent install command |
| 40-41 | heading | context: section 2 title |
| 42 | bullet | context: Tailscale client install on the PC (developer access) |
| 43 | bullet | REQ-DV-0003 |
| 44-45 | bullet | context: test step |
| 46-49 | code block | context: ping command with placeholder |
| 50-51 | paragraph | context: use the Tailscale address, not the LAN address, for box work (developer access) |
| 52-53 | heading | context: section 3 title |
| 54-55 | paragraph | REQ-DV-0004 |
| 56-61 | code block | context: SSH key commands with placeholders |
| 62-63 | paragraph | context: optional SSH config |
| 64-70 | code block | context: SSH config example |
| 71-72 | paragraph | context: SSH keep-alive explanation |
| 73-74 | paragraph | REQ-DV-0004 |
| 75-80 | code block | REQ-DV-0004 |
| 81-82 | heading | context: section 4 title |
| 83-84 | paragraph | REQ-DV-0002 |
| 85-91 | code block | REQ-DV-0002 (line 89); context: clone commands (source names a specific account; not copied) |
| 92-93 | paragraph | context: two ways of push access |
| 94-98 | code block | REQ-DV-0005 |
| 99-100 | paragraph | REQ-DV-0005 |
| 101-102 | heading | context: section 5 title |
| 103-106 | code block | context: Python test command |
| 107-108 | paragraph | REQ-DV-0006, REQ-DV-0007 |
| 109-110 | paragraph | REQ-DV-0008 |
| 111-118 | code block | REQ-DV-0008 |
| 119-120 | paragraph | REQ-DV-0008 |
| 121-122 | heading | context: 5A title |
| 123-124 | paragraph | REQ-DV-0009 |
| 125-133 | code block | REQ-DV-0009 |
| 134-135 | paragraph | REQ-DV-0009, REQ-DV-0010, REQ-DV-0011; context: "all were checked with a small Go and C program" records a past check |
| 136-137 | paragraph | REQ-WP-0140 (location superseded by the work plan's WSL2 copy) |
| 138-139 | paragraph | REQ-WP-0028, REQ-WP-0029; context: `.wslconfig` memory setting is a PC option |
| 140-141 | heading | context: section 6 title |
| 142-144 | paragraph | REQ-DV-0012 |
| 145-146 | paragraph | REQ-DV-0013 |
| 147-154 | code block | REQ-DV-0014 (line 152); context: Python deploy commands |
| 155-156 | paragraph | REQ-DV-0014 |
| 157-158 | paragraph | REQ-DV-0015, REQ-DV-0016 |
| 159-168 | code block | REQ-DV-0015, REQ-DV-0016, REQ-DV-0017 |
| 169-170 | paragraph | REQ-DV-0017, REQ-DV-0018 |
| 171-174 | code block | REQ-DV-0018 |
| 175-176 | heading | context: section 7 title |
| 177-178 | paragraph | REQ-DV-0019, REQ-DV-0020 |
| 179-183 | code block | context: scenario start and log commands (log under `/var/log/nmb`) |
| 184-185 | paragraph | REQ-DV-0021; context: pointer to TESTING.md sections 2a and 7 (covered by the TESTING part) |
| 186-187 | heading | context: section 8 title |
| 188-189 | paragraph | REQ-DV-0022 |
| 190 | bullet | context: how to copy the notes to a new machine |
| 191-192 | bullet | context: alternative: give the state in the first message |
| 193-194 | paragraph | REQ-DV-0022 |
| 195-196 | heading | context: section 9 title |
| 197 | bullet | REQ-DV-0023 |
| 198 | bullet | REQ-DV-0003 |
| 199 | bullet | REQ-DV-0024 |
| 200 | bullet | REQ-DV-0025, REQ-DV-0026 |
| 201-202 | bullet | REQ-WP-0027 |
| 203-204 | heading | context: checklist title |
| 205-206 | table row | context: table header and alignment row |
| 207 | table row | REQ-DV-0001 |
| 208 | table row | context: Tailscale reachability of the box (developer access) |
| 209 | table row | REQ-DV-0004 (key login) |
| 210 | table row | REQ-DV-0002 |
| 211 | table row | context: push access works |
| 212 | table row | REQ-DV-0006 |
| 213 | table row | REQ-DV-0018 |
| 214 | table row | context: agent notes copied or state given |

## Proposed package changes

1. **One owner per Python file (REQ-WP-0135).** Section 3 names six Python files for two packages each, which the register's "exactly one package" check would fail: `probe/engine.py` (M-01, M-02), `report/metrics.py` (M-02, R-01), `netd/modes.py` (N-03, N-05), `netd/wifi.py` (M-14 status parts, N-07), `netd/wifi_share.py` (N-05 guard, N-08), `web/server.py` (S-01 transport, S-02…S-0n handlers). Proposal: one owning package per file, the other package listed as "uses" with the parts it takes in its brief: `probe/engine.py` to M-01, `report/metrics.py` to R-01, `netd/modes.py` to N-03, `netd/wifi.py` to N-07, `netd/wifi_share.py` to N-08, `web/server.py` to S-01 (handlers split to S-02…S-0n by area, each area's lines named in its brief). Alternatively the register's check counts a file split by named line ranges as covered once. Section 3 also names four further files or folders for two packages (added in round 1): `systemd/` units (P-04 sandbox, P-10 service units), `bin/nmb-helper` (P-06, and P-10's `bin/`), `bin/nmb-factory-reset` (P-03, and P-10's `bin/`) and `os/nmb-guard.nft` (N-02, and P-10's `os/`). The register's M3 gives each `systemd/` unit to P-10 (P-04 takes the sandbox lines), `bin/nmb-helper` to P-06, `bin/nmb-factory-reset` to P-03 and `os/nmb-guard.nft` to N-02.
2. **Split S-11 into S-11a "Bot commands and menus" and S-11b "Alert monitor".** `nmb/bot/` is 4,132 lines (`commands.py` 1,725, `alerts.py` 310, the rest translations and client), well above the size rule (REQ-WP-0055). The alert monitor (REQ-AR-0072 to -0078) is a separate state machine with its own thresholds and dedup table. S-11a owns REQ-AR-0050 to -0071, S-11b owns REQ-AR-0072 to -0078.
3. **Optional split of F-00** into F-00a "Repository skeleton, toolchain, licence list" and F-00b "Gate scripts and work status". F-00 owns about 60 process and gate requirements of this part (REQ-WP-0004 to -0057 and others); F-00b would own the gate rows (REQ-WP-0017 to -0028, -0034 to -0036, -0043 to -0053, -0172) and the scans (REQ-DV-0003 to -0005, -0022, -0023).
4. **NEW: X2-01 Step-2 router tiers** (proposed in the roadmap part) owns REQ-DV-0011 here and is "Also in" for REQ-WP-0006, -0007, -0009, -0064, -0147, -0156, -0161, -0171 and REQ-DV-0010. No further new package.

## Questions for the owner

*Triaged in round 3 (2026-11-11), answers and choices in `QUESTIONS.md`:* 3 → Q-14, 5 → Q-09 (answered; both moved to the developer's choices in round 4); 1 → Q-02, 2 → Q-13 (developer's choices); 4 → Q-15 (for the owner). Answered by the owner on 2026-11-11: Q-15 (REQ-AR-0054, -0069 restated, REQ-AR-0079 and -0080 added). The rows were updated where an answer or a choice settled them; the text below is the question as the part asked it.

1. **F-07 in the order (work plan lines 103 and 222):** section 3.0 says the foundation F-00 to F-07 comes in order before anything else, but section 4 step 1 lists only F-00 to F-06. Should F-07 (`spec/` and vectors) also be done before Core starts, as rule 11 of 2.2 ("behaviour is changed in `spec/` first") implies, or may it grow alongside Core?
2. **Capabilities of the web server (ARCHITECTURE line 45):** the Python web server holds `CAP_NET_ADMIN` and `CAP_NET_RAW` besides `CAP_NET_BIND_SERVICE`. Is that intended for the native build, or should the native web server have no network-admin rights and pass such tasks to the network daemon or the privileged helper (QB 7)?
3. **`/kill` from Telegram (ARCHITECTURE line 135):** the Python bot sends a signal to any PID above 1, with `sudo` if needed (`nmb/bot/commands.py@ad3758c:1423`, `nmb/common/tasks.py@ad3758c:444-474`), so it can also stop the SSH server or Tailscale. Intended, or should the native `/kill` be limited to NMB's own processes?
4. **Wi-Fi password in chat (ARCHITECTURE line 136):** `/wifi connect <ssid> [pw]` puts the Wi-Fi password in a Telegram message, which Telegram keeps (the same reason the backup is sent encrypted). Keep it as written, or should the native bot only point to the Web UI for networks with a password?
5. **Documentation languages (DEV_SETUP line 200):** the Python repository keeps documentation in English and Hungarian. Does that rule hold for the native repository too (manual and command reference of P-12), or does the native documentation follow the 11 UI languages, or English only?
