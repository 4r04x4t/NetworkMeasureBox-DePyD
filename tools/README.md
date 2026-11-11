# tools/

The gate of the native build and its checkers. How to run them: [`docs/DEVELOPMENT.md`](../docs/DEVELOPMENT.md).

| File | Does |
| :--- | :--- |
| `gate.sh` | runs every step below and prints what ran where |
| `lib.sh` | common part of the scripts (sourced): refusals, toolchain, knobs, steps, offline and hermetic runs |
| `prepare.sh`, `deps.sh`, `build-matrix.sh`, `analyse.sh`, `test.sh`, `fuzz.sh`, `scan.sh` | the single steps |
| `gate-selftest.sh` | proves each check can fail, on throw-away copies |
| `cleancode/` | function, file and nesting limits (Go, standard library only) |
| `privscan/` | personal data and secrets scanner (Go, standard library only), with `allow.txt` |
| `depcheck/` | dependency and licence list checker (Go, standard library only) |
| `gosecvet/` | its own module: runs the gosec rules through gosec's go/analysis entry point |
