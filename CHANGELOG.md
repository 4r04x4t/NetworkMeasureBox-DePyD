# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Native rewrite project and this changelog.
- `nmbd --help` and `nmbd --version`: the first native command. There is no daemon yet; `--help` prints the usage and changes nothing.
- One source tree built for Linux on x86-64, 64-bit ARM, 32-bit ARM v7 and MIPS (big- and little-endian), with Go 1.27.4 and without cgo. Tests run on x86-64 and, under QEMU, on both ARM targets; every check is in `tools/gate.sh`.
- A scan of the sources, the built programs and the whole history for personal data and secrets, which fails the build on a hit (`tools/privscan`).
- `docs/DEPENDENCIES.md`: every dependency with its licence and version. A dependency that is not listed fails the build. The licence of the whole project is not decided yet.
- `KNOWN_PYTHON_FAULTS.md`: 294 faults of the Python version, each with the correct behaviour that the native build must have and the package that has to prove it. Open security holes are listed by category only.
- `docs/requirements/`: 2,340 requirements taken from the plans and the documents of the Python version, each with its source line and the work package that owns it.
- The check of the requirements register and its test now run in `tools/gate.sh` (`tools/register.sh`), so a register that disagrees with itself or with its sources fails the build.

### Fixed

- `tools/gate.sh` failed on `main`: shellcheck found unquoted variables and an `A && B || C` test in the register's check scripts. Both scripts are fixed and can now be called from any folder.
