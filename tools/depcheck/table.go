package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Column names of the dependency table, in this exact order.
func columns() []string {
	return []string{"Name", "Kind", "Module or package", "Version", "Licence (SPDX)", "Origin", "Use",
		"Licence file SHA-256", "Added by", "Checked"}
}

// Kinds of rows.
const (
	kindToolchain     = "toolchain"      // the Go toolchain and standard library
	kindToolchainPart = "toolchain-part" // a part shipped inside the toolchain (race runtime)
	kindModule        = "module"         // a module in the require list of a go.mod
	kindGateTool      = "gate-tool"      // an analyser installed by tools/prepare.sh
	kindSystem        = "system"         // a program of the development PC the gate runs
)

// Uses of a dependency.
const (
	useBuiltIn   = "built in"
	useTestOnly  = "test only"
	useBuildOnly = "build only"
)

// noHash is written instead of a licence-file hash where no licence file is checked.
const noHash = "-"

// Row is one dependency.
type Row struct {
	Line      int
	Name      string
	Kind      string
	Module    string
	Version   string
	Licence   string
	Origin    string
	Use       string
	Hash      string
	AddedBy   string
	Checked   string
	permitted bool // licence allowed for parts built into nmbd
}

// licences maps each known SPDX expression to whether it may be built into nmbd (question Q4:
// anything else stops until the owner decides).
func licences() map[string]bool {
	return map[string]bool{
		"MIT": true, "BSD-2-Clause": true, "BSD-3-Clause": true, "Apache-2.0": true, "ISC": true,
		"0BSD": true, "MPL-2.0": true,
		"GPL-2.0-only": false, "GPL-2.0-or-later": false, "GPL-3.0-only": false, "GPL-3.0-or-later": false,
		"GPL-3.0-or-later WITH GCC-exception-3.1": false, "LGPL-2.1-or-later": false,
		"Apache-2.0 WITH LLVM-exception": false, "LicenseRef-Public-Domain": false,
	}
}

type validators struct {
	modVersion, goVersion, sysVersion, path, hash, addedBy, origin *regexp.Regexp
}

func newValidators() validators {
	return validators{
		modVersion: regexp.MustCompile(`^v[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.-]+)?(\+incompatible)?$`),
		goVersion:  regexp.MustCompile(`^go1\.[0-9]+\.[0-9]+$`),
		sysVersion: regexp.MustCompile(`^[0-9][0-9A-Za-z.+~-]*$`),
		path:       regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._~/-]*$`),
		hash:       regexp.MustCompile(`^[0-9a-f]{64}$`),
		addedBy:    regexp.MustCompile(`^[A-Z]-[0-9]{2}$`),
		origin:     regexp.MustCompile(`^https://[A-Za-z0-9.-]+(/[^\s|]*)?$`),
	}
}

// ParseTable reads the one dependency table of the document. A missing or second table, an
// unknown column, a missing cell or a value of the wrong form is an error.
func ParseTable(doc string) ([]Row, error) {
	lines := strings.Split(doc, "\n")
	header := "| " + strings.Join(columns(), " | ") + " |"
	start := -1
	for i, l := range lines {
		if strings.TrimSpace(l) == header {
			if start >= 0 {
				return nil, fmt.Errorf("line %d: a second dependency table", i+1)
			}
			start = i
		}
	}
	if start < 0 || start+1 >= len(lines) || !isSeparator(lines[start+1]) {
		return nil, fmt.Errorf("no dependency table with the columns %q and a separator line", header)
	}
	v := newValidators()
	var rows []Row
	for i := start + 2; i < len(lines) && strings.HasPrefix(strings.TrimSpace(lines[i]), "|"); i++ {
		row, err := parseRow(lines[i], i+1, v)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("the dependency table has no rows")
	}
	return rows, nil
}

func isSeparator(line string) bool {
	cells, ok := splitCells(line)
	if !ok || len(cells) != len(columns()) {
		return false
	}
	for _, c := range cells {
		if strings.Trim(c, ":-") != "" || !strings.Contains(c, "---") {
			return false
		}
	}
	return true
}

// splitCells splits "| a | b |" into its trimmed cells.
func splitCells(line string) ([]string, bool) {
	line = strings.TrimSpace(line)
	if len(line) < 2 || line[0] != '|' || line[len(line)-1] != '|' {
		return nil, false
	}
	cells := strings.Split(line[1:len(line)-1], "|")
	for i, c := range cells {
		c = strings.TrimSpace(c)
		if len(c) >= 2 && c[0] == '`' && c[len(c)-1] == '`' {
			c = c[1 : len(c)-1]
		}
		cells[i] = c
	}
	return cells, true
}

func parseRow(line string, n int, v validators) (Row, error) {
	cells, ok := splitCells(line)
	if !ok || len(cells) != len(columns()) {
		return Row{}, fmt.Errorf("line %d: want %d cells", n, len(columns()))
	}
	r := Row{Line: n, Name: cells[0], Kind: cells[1], Module: cells[2], Version: cells[3], Licence: cells[4],
		Origin: cells[5], Use: cells[6], Hash: cells[7], AddedBy: cells[8], Checked: cells[9]}
	permitted, known := licences()[r.Licence]
	if !known {
		return r, fmt.Errorf("line %d: licence %q is not a known SPDX id of this list", n, r.Licence)
	}
	r.permitted = permitted
	if err := validateRow(r, v); err != nil {
		return r, fmt.Errorf("line %d: %w", n, err)
	}
	return r, nil
}

func validateRow(r Row, v validators) error {
	checks := []struct {
		ok   bool
		what string
	}{
		{r.Name != "" && len(r.Name) <= 80, "name empty or longer than 80"},
		{kindVersion(r.Kind, v) != nil, "unknown kind " + r.Kind},
		{r.Kind == kindSystem || v.path.MatchString(r.Module), "malformed module or package"},
		{kindVersion(r.Kind, v) == nil || kindVersion(r.Kind, v).MatchString(r.Version), "malformed version for kind " + r.Kind},
		{v.origin.MatchString(r.Origin), "origin must be an https address"},
		{r.Use == useBuiltIn || r.Use == useTestOnly || r.Use == useBuildOnly, "unknown use " + r.Use},
		{hashFormOK(r, v), "licence file SHA-256 must be 64 lower-case hex digits, or - for kinds system and toolchain-part"},
		{v.addedBy.MatchString(r.AddedBy), "added by must be a package id such as F-00"},
		{dateOK(r.Checked), "checked must be a date YYYY-MM-DD"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf("%s (row %q)", c.what, r.Name)
		}
	}
	return nil
}

func kindVersion(kind string, v validators) *regexp.Regexp {
	switch kind {
	case kindToolchain, kindToolchainPart:
		return v.goVersion
	case kindModule, kindGateTool:
		return v.modVersion
	case kindSystem:
		return v.sysVersion
	}
	return nil
}

func hashFormOK(r Row, v validators) bool {
	if r.Kind == kindSystem || r.Kind == kindToolchainPart {
		return r.Hash == noHash
	}
	return v.hash.MatchString(r.Hash)
}

func dateOK(s string) bool {
	_, err := time.Parse(time.DateOnly, s)
	return err == nil
}
