package system

import (
	"runtime"
	"runtime/debug"
)

type GoRuntimeAdapter struct{}

func NewGoRuntimeAdapter() *GoRuntimeAdapter {
	return &GoRuntimeAdapter{}
}

func (a *GoRuntimeAdapter) GetGoVersion() string {
	return runtime.Version()
}

func (a *GoRuntimeAdapter) GetOSArch() string {
	return runtime.GOOS + "/" + runtime.GOARCH
}

func (a *GoRuntimeAdapter) GetDependencies() []Dependency {
	var deps []Dependency
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range buildInfo.Deps {
			deps = append(deps, Dependency{
				Path:    dep.Path,
				Version: dep.Version,
			})
		}
	}
	return deps
}
