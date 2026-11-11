// Package buildinfo reports what the Go toolchain recorded about a build.
//
// Nothing is invented: a field the build did not record is reported as NotRecorded, never as a
// default. The version is the main module version Go stamps from the version control state.
package buildinfo

import (
	"errors"
	"runtime/debug"
)

// NotRecorded is reported for every field the build did not record.
const NotRecorded = "not recorded"

// ErrNoBuildInfo is returned when the binary carries no build information at all.
var ErrNoBuildInfo = errors.New("buildinfo: the binary carries no build information")

// ReadFunc returns the build information of the running binary; debug.ReadBuildInfo fits.
type ReadFunc func() (*debug.BuildInfo, bool)

// Field is one reported line: a stable key (data, never translated) and its value.
type Field struct {
	Key   string
	Value string
}

// reportedSetting maps a build setting to the key it is reported under.
type reportedSetting struct {
	key     string
	setting string
	arch    string // reported only for this architecture family; empty means always
}

// reportedSettings lists the build settings reported besides the module and Go versions, in this
// order. GOARM and GOMIPS are only reported for the architectures they belong to.
func reportedSettings() []reportedSetting {
	return []reportedSetting{
		{key: "os", setting: "GOOS"},
		{key: "arch", setting: "GOARCH"},
		{key: "goarm", setting: "GOARM", arch: "arm"},
		{key: "gomips", setting: "GOMIPS", arch: "mips"},
		{key: "cgo", setting: "CGO_ENABLED"},
		{key: "revision", setting: "vcs.revision"},
		{key: "commit-time", setting: "vcs.time"},
		{key: "modified", setting: "vcs.modified"},
	}
}

// Read returns the fields of the build, in a fixed order.
func Read(read ReadFunc) ([]Field, error) {
	info, ok := read()
	if !ok || info == nil {
		return nil, ErrNoBuildInfo
	}
	settings := make(map[string]string, len(info.Settings))
	for _, s := range info.Settings {
		settings[s.Key] = s.Value
	}
	fields := []Field{
		{Key: "version", Value: orNotRecorded(info.Main.Version)},
		{Key: "module", Value: orNotRecorded(info.Main.Path)},
		{Key: "go", Value: orNotRecorded(info.GoVersion)},
	}
	arch := settings["GOARCH"]
	for _, r := range reportedSettings() {
		if r.arch != "" && !archMatches(arch, r.arch) {
			continue
		}
		fields = append(fields, Field{Key: r.key, Value: orNotRecorded(settings[r.setting])})
	}
	return fields, nil
}

// archMatches reports whether arch is family itself or a variant of it (mips matches mipsle).
func archMatches(arch, family string) bool {
	switch family {
	case "mips":
		return arch == "mips" || arch == "mipsle"
	default:
		return arch == family
	}
}

func orNotRecorded(v string) string {
	if v == "" {
		return NotRecorded
	}
	return v
}
