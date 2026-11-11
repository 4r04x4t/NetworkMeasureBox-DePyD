package buildinfo

import (
	"errors"
	"runtime/debug"
	"testing"
)

func fake(info *debug.BuildInfo, ok bool) ReadFunc {
	return func() (*debug.BuildInfo, bool) { return info, ok }
}

func settings(kv ...string) []debug.BuildSetting {
	var out []debug.BuildSetting
	for i := 0; i+1 < len(kv); i += 2 {
		out = append(out, debug.BuildSetting{Key: kv[i], Value: kv[i+1]})
	}
	return out
}

func asMap(t *testing.T, fields []Field) map[string]string {
	t.Helper()
	m := make(map[string]string, len(fields))
	for _, f := range fields {
		if _, dup := m[f.Key]; dup {
			t.Fatalf("key %q reported twice", f.Key)
		}
		m[f.Key] = f.Value
	}
	return m
}

// F00-R09: the version lines come from what the build recorded.
func TestReadFromFake(t *testing.T) {
	info := &debug.BuildInfo{
		GoVersion: "go1.27.4",
		Main:      debug.Module{Path: "nmb", Version: "v0.0.0-20261111000000-000000000000"},
		Settings: settings("GOOS", "linux", "GOARCH", "arm", "GOARM", "7", "CGO_ENABLED", "0",
			"vcs.revision", "0123456789abcdef", "vcs.time", "2026-11-11T00:00:00Z", "vcs.modified", "false"),
	}
	fields, err := Read(fake(info, true))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	want := []Field{
		{"version", "v0.0.0-20261111000000-000000000000"}, {"module", "nmb"}, {"go", "go1.27.4"},
		{"os", "linux"}, {"arch", "arm"}, {"goarm", "7"}, {"cgo", "0"},
		{"revision", "0123456789abcdef"}, {"commit-time", "2026-11-11T00:00:00Z"}, {"modified", "false"},
	}
	if len(fields) != len(want) {
		t.Fatalf("got %d fields %v, want %d", len(fields), fields, len(want))
	}
	for i := range want {
		if fields[i] != want[i] {
			t.Errorf("field %d = %v, want %v", i, fields[i], want[i])
		}
	}
}

// Honest data: a field the build did not record is reported as not recorded, never as a default.
func TestMissingFieldsNotRecorded(t *testing.T) {
	info := &debug.BuildInfo{Settings: settings("GOARCH", "mipsle")}
	fields, err := Read(fake(info, true))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	m := asMap(t, fields)
	for _, k := range []string{"version", "module", "go", "os", "gomips", "cgo", "revision", "commit-time", "modified"} {
		if m[k] != NotRecorded {
			t.Errorf("%s = %q, want %q", k, m[k], NotRecorded)
		}
	}
	if _, ok := m["goarm"]; ok {
		t.Errorf("goarm reported for mipsle")
	}
	if m["arch"] != "mipsle" {
		t.Errorf("arch = %q", m["arch"])
	}
}

func TestArchSpecificSettingsOnlyForTheirArch(t *testing.T) {
	info := &debug.BuildInfo{Settings: settings("GOARCH", "amd64", "GOARM", "7", "GOMIPS", "softfloat")}
	fields, err := Read(fake(info, true))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	m := asMap(t, fields)
	if _, ok := m["goarm"]; ok {
		t.Errorf("goarm reported for amd64")
	}
	if _, ok := m["gomips"]; ok {
		t.Errorf("gomips reported for amd64")
	}
}

func TestNoBuildInfo(t *testing.T) {
	for _, tc := range []struct {
		info *debug.BuildInfo
		ok   bool
	}{{nil, false}, {nil, true}, {&debug.BuildInfo{}, false}} {
		if _, err := Read(fake(tc.info, tc.ok)); !errors.Is(err, ErrNoBuildInfo) {
			t.Errorf("Read(%v, %v) error = %v, want ErrNoBuildInfo", tc.info, tc.ok, err)
		}
	}
}
