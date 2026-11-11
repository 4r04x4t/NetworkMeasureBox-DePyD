# tests/

Tests that span packages only: the Phase 0 golden files and comparison harness (F-05), headless-browser tests of
the Web UI and image tests. Unit tests and fuzz tests live next to their code (`*_test.go`), and the gate runs them
with `tools/test.sh` and `tools/fuzz.sh`. Empty in F-00.
