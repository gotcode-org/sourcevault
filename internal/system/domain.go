package system

type Dependency struct {
	Path    string
	Version string
}

type VersionInfo struct {
	Version   string
	Commit    string
	Branch    string
	GoVersion string
	OSArch    string
	Deps      []Dependency
}

var (
	// Injected by Makefile -ldflags
	AppVersion = "v0.0.0-dev"
	Commit     = "unknown"
	Branch     = "unknown"
)
