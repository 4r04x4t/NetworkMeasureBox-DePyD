# Work status of the native build (step 1: x86-64 and ARM)

Kept by the orchestrator (work plan section 2.3). States: waiting, brief, code, verify, accepted, blocked.
Packages and order: `docs/NATIVE_WORK_PLAN.md` in the Python repository, sections 3 and 4.
An accepted package is never redone; this file wins over anyone's memory.

Remote: `origin` = `NetworkMeasureBox-DePyD` on GitHub (public, empty until the first push), same account as the Python repository. Commits use the GitHub no-reply e-mail (owner, 2026-11-11).

## Foundation

| Id | Package | State | Commit | Gate | Open |
| :--- | :--- | :--- | :--- | :--- | :--- |
| F-00 | Skeleton | code (brief 6c3eb45 on pkg/F-00) | — | — | — |
| F-01 | Owner's open decisions | code (answered 2026-11-11; plan patch waits until F-02/F-03 finish reading) | — | — | — |
| F-02 | Requirements register | brief (register being extracted) | — | — | — |
| F-03 | Known Python faults | brief | — | — | — |
| F-04 | Contracts | waiting | — | — | — |
| F-05 | Phase 0 reference | waiting | — | — | — |
| F-06 | Test bed | waiting | — | — | — |
| F-07 | Shared specification and vectors | waiting | — | — | — |

## Images

| Id | Image | State | Commit | Open |
| :--- | :--- | :--- | :--- | :--- |
| I-01 | OpenWrt 25.12.6 x86-64 | verify (record b86aaf8 on pkg/I-01) | — | — |
| I-02 | OpenWrt 25.12.6 armsr/armv8 | brief | — | — |
| I-03 | OpenWrt 25.12.6 armsr/armv7 | brief | — | — |
| I-04 | OpenWrt 25.12.6 x86/generic | brief | — | — |
| I-05 | OpenWrt 25.12.6 Raspberry Pi 4 | brief | — | — |
| I-06 | Debian 13.7.0 amd64 and arm64 | brief | — | — |
| I-07 | Alpine 3.24.4 | brief | — | — |
| I-08 | Fedora IoT 44 | brief | — | — |
| I-09 | Comparison for 6A | waiting | — | — |

## Core, measuring, network, reports, surfaces, product

All waiting for the foundation (C-01..C-11, M-01..M-14, N-01..N-18, R-01..R-05, S/U, P-01..P-13).

## Log

- 2026-11-11: repository created (`git init`), first commit with `CHANGELOG.md`, `README.md`, this file.
- 2026-11-11: GitHub repository created by the owner; `origin` added, nothing pushed yet. Owner: commits use the GitHub no-reply e-mail; the first commits with the old address are rewritten before the first push.
- 2026-11-11: sudo on the PC and KVM made available by the owner; cpio, bsdtar, xorriso, ostree, shellcheck, squashfs-tools-ng installed.
- 2026-11-11: started in parallel: F-00 (implement), I-01 (verify), F-02, F-03, I-02..I-05.
- 2026-11-11: owner chose the disclaimer (warranty and liability exclusion plus measurement note); I-06..I-08 started one at a time (RAM).
