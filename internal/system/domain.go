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
