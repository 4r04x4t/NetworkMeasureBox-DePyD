# nmbd

The native NetworkMeasureBox program. In this build (package F-00, the skeleton) it has no daemon and does
nothing besides printing its usage and its build information.

## Usage

The block below is the exact output of `nmbd --help`. The test `TestHelpMatchesReference` in `cmd/nmbd`
compares the two byte for byte, so this entry and the program cannot disagree.

<!-- nmbd-help-begin -->
```text
Usage: nmbd --help | --version

nmbd is the native NetworkMeasureBox program. This build is the skeleton:
there is no daemon yet and it does nothing besides printing this text and
its build information. It changes nothing and needs no root.

Options:
  -h, --help   print this text and exit
  --version    print the build information and exit

Exit status: 0 for --help and --version, 2 for any other use,
1 when the output cannot be written or the build information is missing.
```
<!-- nmbd-help-end -->

## Options

| Option | Does | Output | Exit code |
| :--- | :--- | :--- | :--- |
| `-h`, `--help` | prints the usage above | stdout | 0 |
| `--version` | prints the build information, one `key: value` line each | stdout | 0 |
| none, an unknown option, an extra argument | prints the reason and the usage | stderr | 2 |

The standard Go flag syntax applies, so `-help` and `-version` are accepted as well. When `--help` and
`--version` are both given, the usage is printed.

`--version` prints these keys in this order: `version`, `module`, `go`, `os`, `arch`, `goarm` (arm only),
`gomips` (mips and mipsle only), `cgo`, `revision`, `commit-time`, `modified`. The values are what the Go
toolchain recorded in the binary; a value it did not record is printed as `not recorded`. No version number is
invented: the version is the one Go stamps from the git state of the build.

## What it changes

Nothing: it opens no file or socket, reads no setting, writes only to stdout and stderr, and needs no root.

## Example

```sh
nmbd --version
```
