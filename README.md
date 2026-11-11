# NetworkMeasureBox-DePyD

The native (Python-free) build of NetworkMeasureBox: one `nmbd` program in Go (step 1: x86-64 and ARM),
later a C11 code base for small routers, both held to one shared specification in `spec/`.

Status: **in development, nothing usable yet.** The plans this work follows are kept in the Python
version's repository (`docs/NATIVE_WORK_PLAN.md`, `docs/DEPYTHONIZATION_ROADMAP.md`,
`docs/ROUTER_PORTING_AND_DEPYTHONIZATION_PLAN.md`). Progress per work package: [`WORK_STATUS.md`](WORK_STATUS.md).

## Licence

**Not decided yet.** The licence of the whole is fixed once its components are known, from what their
licences require and allow (work plan section 7, point 9). Until then no dependency is taken whose licence
would rule out an open licence; every dependency is listed with its licence as it is added.
