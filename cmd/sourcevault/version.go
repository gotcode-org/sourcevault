package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"gotcode.org/sourcevault/internal/system"
	"gotcode.org/sourcevault/internal/ui"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information and dependencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			runtimeAdapter := system.NewGoRuntimeAdapter()
			queryHandler := system.NewGetVersionQueryHandler(runtimeAdapter)

			query := system.GetVersionQuery{
				AppVersion: Version,
				Commit:     Commit,
				Branch:     Branch,
			}

			info, err := queryHandler.Handle(cmd.Context(), query)
			if err != nil {
				return fmt.Errorf("fetching version info: %w", err)
			}

			ui.PrintVersion(cmd.OutOrStdout(), info)
			return nil
		},
	}
}
