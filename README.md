# NetworkMeasureBox-DePyD

The native (Python-free) build of NetworkMeasureBox: one `nmbd` program in Go (step 1: x86-64 and ARM),
later a C11 code base for small routers, both held to one shared specification in `spec/`.

Status: **in development, nothing usable yet.** The plans this work follows are kept in the Python
version's repository (`docs/NATIVE_WORK_PLAN.md`, `docs/DEPYTHONIZATION_ROADMAP.md`,
`docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md`). Progress per work package: [`WORK_STATUS.md`](WORK_STATUS.md).

## Licence

**MIT** (owner, 2026-11-11), the same licence as the Python version: see [`LICENSE`](LICENSE). The choice was made
from the licences of every planned component ([`docs/licensing/`](docs/licensing/PROPOSAL.md)): nothing built into
`nmbd` may be under a copyleft licence. Programs that the box only runs (for example `ooniprobe`, `pppd`, `nft`,
`dnsmasq`) keep their own licences, and the images carry their licence texts and source offers.
