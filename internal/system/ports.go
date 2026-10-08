package system

// RuntimePort defines how the system domain retrieves information about its executing environment.
type RuntimePort interface {
	GetGoVersion() string
	GetOSArch() string
	GetDependencies() []Dependency
}
