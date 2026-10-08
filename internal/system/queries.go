package system

import "context"

type GetVersionQuery struct {
	AppVersion string
	Commit     string
	Branch     string
}

type GetVersionQueryHandler struct {
	runtime RuntimePort
}

func NewGetVersionQueryHandler(runtime RuntimePort) *GetVersionQueryHandler {
	return &GetVersionQueryHandler{runtime: runtime}
}

func (h *GetVersionQueryHandler) Handle(ctx context.Context, query GetVersionQuery) (VersionInfo, error) {
	return VersionInfo{
		Version:   query.AppVersion,
		Commit:    query.Commit,
		Branch:    query.Branch,
		GoVersion: h.runtime.GetGoVersion(),
		OSArch:    h.runtime.GetOSArch(),
		Deps:      h.runtime.GetDependencies(),
	}, nil
}
