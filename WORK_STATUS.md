# Work status of the native build (step 1: x86-64 and ARM)

Kept by the orchestrator (work plan section 2.3). States: waiting, brief, code, verify, accepted, blocked.
Packages and order: `docs/NATIVE_WORK_PLAN.md` in the Python repository, sections 3 and 4.
An accepted package is never redone; this file wins over anyone's memory.

Remote: `origin` = `NetworkMeasureBox-DePyD` on GitHub (public, empty until the first push), same account as the Python repository. Commits use the GitHub no-reply e-mail (owner, 2026-11-11).

## Foundation

| Id | Package | State | Commit | Gate | Open |
| :--- | :--- | :--- | :--- | :--- | :--- |
| F-00 | Skeleton | accepted | merge 20a0409 | gate rc 0 (orchestrator, 6m14s); verifier pass; privscan rounds 2-3 pass | scanner known limits in docs/DEVELOPMENT.md |
| F-01 | Owner's open decisions | accepted | Python repo bb0a553 (pushed to its origin main) | verifier pass, minors fixed | — |
| F-02 | Requirements register | accepted | merge 71d3b80 | check.sh rc 0 (2340 requirements); independent checks pass | no open question |
| F-03 | Known Python faults | accepted | merge aafe010 | verifier pass (294 entries, twelve-field check) | no open question |
| F-04 | Contracts | brief (six part briefs a-f + coordinator) | — | — | — |
| F-05 | Phase 0 reference | brief | — | — | — |
| F-06 | Test bed | brief | — | — | — |
| F-07 | Shared specification and vectors | waiting | — | — | — |

## Images

| Id | Image | State | Commit | Open |
| :--- | :--- | :--- | :--- | :--- |
| I-01 | OpenWrt 25.12.6 x86-64 | accepted | merge 83a14fd | — |
| I-02 | OpenWrt 25.12.6 armsr/armv8 | accepted | merge fed2004 | — |
| I-03 | OpenWrt 25.12.6 armsr/armv7 | accepted | merge c653884 | — |
| I-04 | OpenWrt 25.12.6 x86/generic | verify (one open major: boot time of run 3) | — | — |
| I-05 | OpenWrt 25.12.6 Raspberry Pi 4 | accepted (taken apart only; boot on the Pi 4 is a later step) | merge 996f6e9 | — |
| I-06 | Debian 13.7.0 amd64 and arm64 | accepted | merge 57a6e1b | — |
| I-07 | Alpine 3.24.4 | brief (being analysed) | — | — |
| I-08 | Fedora IoT 44 | waiting (after I-07) | — | — |
| I-09 | Comparison for 6A | waiting | — | — |

## Core, measuring, network, reports, surfaces, product

All waiting for the foundation (C-01..C-11, M-01..M-14, N-01..N-18, R-01..R-05, S/U, P-01..P-13).

## Log

- 2026-11-11: repository created (`git init`), first commit with `CHANGELOG.md`, `README.md`, this file.
- 2026-11-11: GitHub repository created by the owner; `origin` added, nothing pushed yet. Owner: commits use the GitHub no-reply e-mail; the first commits with the old address are rewritten before the first push.
- 2026-11-11: sudo on the PC and KVM made available by the owner; cpio, bsdtar, xorriso, ostree, shellcheck, squashfs-tools-ng installed.
- 2026-11-11: started in parallel: F-00 (implement), I-01 (verify), F-02, F-03, I-02..I-05.
- 2026-11-11: owner chose the disclaimer (warranty and liability exclusion plus measurement note); I-06..I-08 started one at a time (RAM).
- 2026-11-11: every commit rewritten to the no-reply e-mail (backup bundle kept outside the repository); old commit ids in earlier notes no longer exist.
- 2026-11-11: owner answered the six F-03 questions (uplink trust per network, Ookla manual runs obey the pause only, VPN mismatch shown and asked, no default data cap, minute summaries only, bot confirms every change).
- 2026-11-11: F-02 round 3, F-03 round 4, I-07/I-08 and image minor fixes running; F-04, F-05, F-06 briefs started early (briefs only; code waits for F-00..F-03).
- 2026-11-11: owner answered F-03 Q6/Q7 and F-02 Q-01, Q-05, Q-12, Q-15, Q-16, Q-17 (migration component by component; all Python distributions; terminal kept hardened; bot secrets off by default; MACs pseudonymised for outside LLMs; maintenance needs an off-device log). F-01 started.
- 2026-11-11: F-01 accepted: every owner decision and the image findings written into the shared plans (Python repo 492fc46 + bb0a553, pushed). The Python folder ~/src/NetworkMeasureBox stays at ad3758c until the briefs that cite it finish; then fast-forward.
- 2026-11-11: F-00, F-03, F-02 merged into main; changelog updated.
- 2026-11-11: image records I-01, I-02, I-03, I-05, I-06 accepted and merged; their findings are in roadmap 6A (Python repo bb0a553). Gate on main: shellcheck fails on the register scripts; fix in progress before the first push.
