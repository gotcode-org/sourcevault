package system

import "context"

type GetVersionQuery struct{}

type GetVersionQueryHandler struct {
	runtime RuntimePort
}

func NewGetVersionQueryHandler(runtime RuntimePort) *GetVersionQueryHandler {
	return &GetVersionQueryHandler{runtime: runtime}
}

func (h *GetVersionQueryHandler) Handle(ctx context.Context, query GetVersionQuery) (VersionInfo, error) {
	return VersionInfo{
		Version:   AppVersion,
		Commit:    Commit,
		Branch:    Branch,
		GoVersion: h.runtime.GetGoVersion(),
		OSArch:    h.runtime.GetOSArch(),
		Deps:      h.runtime.GetDependencies(),
	}, nil
}
