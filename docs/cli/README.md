# docs/cli/

The command reference (roadmap 6B.7): one entry per command, with its options, output, exit codes, what it
changes and an example. From P-12 on the entries are generated from the commands' own definitions; until then
they are written by hand and each one is checked against the program by a test, so the reference and `--help`
cannot disagree.

| Command | Entry | Checked by |
| :--- | :--- | :--- |
| `nmbd` | [nmbd.md](nmbd.md) | `TestHelpMatchesReference` in `cmd/nmbd` |
