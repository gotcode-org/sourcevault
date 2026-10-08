package ui

import (
	"fmt"
	"io"
	"text/tabwriter"

	"gotcode.org/sourcevault/internal/system"
)

func PrintVersion(w io.Writer, info system.VersionInfo) {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "Version\t%s\n", info.Version)
	fmt.Fprintf(tw, "Git Commit\t%s\n", info.Commit)
	fmt.Fprintf(tw, "Git Branch\t%s\n", info.Branch)
	fmt.Fprintf(tw, "Go Version\t%s\n", info.GoVersion)
	fmt.Fprintf(tw, "OS/Arch\t%s\n", info.OSArch)
	fmt.Fprintf(tw, "--------------------------------\t------------------------------------------\n")

	for _, dep := range info.Deps {
		fmt.Fprintf(tw, "%s\t%s\n", dep.Path, dep.Version)
	}

	_ = tw.Flush()
}
