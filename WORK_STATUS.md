# Work status of the native build (step 1: x86-64 and ARM)

Kept by the orchestrator (work plan section 2.3). States: waiting, brief, code, verify, accepted, blocked.
Packages and order: `docs/NATIVE_WORK_PLAN.md` in the Python repository, sections 3 and 4.
An accepted package is never redone; this file wins over anyone's memory.

Remote: `NetworkMeasureBox-DePyD` on GitHub, same account as the Python repository. On 2026-11-11 it did not
exist yet ("Repository not found"); all work is local until the owner creates the empty public repository.

## Foundation

| Id | Package | State | Commit | Gate | Open |
| :--- | :--- | :--- | :--- | :--- | :--- |
| F-00 | Skeleton | brief | — | — | — |
| F-01 | Owner's open decisions | waiting | — | — | — |
| F-02 | Requirements register | waiting | — | — | — |
| F-03 | Known Python faults | waiting | — | — | — |
| F-04 | Contracts | waiting | — | — | — |
| F-05 | Phase 0 reference | waiting | — | — | — |
| F-06 | Test bed | waiting | — | — | — |
| F-07 | Shared specification and vectors | waiting | — | — | — |

## Images

| Id | Image | State | Commit | Open |
| :--- | :--- | :--- | :--- | :--- |
| I-01 | OpenWrt 25.12.6 x86-64 | brief | — | — |
| I-02 | OpenWrt 25.12.6 armsr/armv8 | waiting | — | — |
| I-03 | OpenWrt 25.12.6 armsr/armv7 | waiting | — | — |
| I-04 | OpenWrt 25.12.6 x86/generic | waiting | — | — |
| I-05 | OpenWrt 25.12.6 Raspberry Pi 4 | waiting | — | — |
| I-06 | Debian 13.7.0 amd64 and arm64 | waiting | — | — |
| I-07 | Alpine 3.24.4 | waiting | — | — |
| I-08 | Fedora IoT 44 | waiting | — | — |
| I-09 | Comparison for 6A | waiting | — | — |

## Core, measuring, network, reports, surfaces, product

All waiting for the foundation (C-01..C-11, M-01..M-14, N-01..N-18, R-01..R-05, S/U, P-01..P-13).

## Log

- 2026-11-11: repository created (`git init`), first commit with `CHANGELOG.md`, `README.md`, this file.
