package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gotcode.org/sourcevault/internal/ui"
)

func main() {
	rootCmd := &cobra.Command{
		Use:           "sourcevaultd",
		Short:         "SourceVault Background Daemon",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCmd.SetHelpFunc(ui.HandleHelp)
	rootCmd.SetUsageFunc(ui.HandleUsage)

	rootCmd.AddCommand(versionCmd())

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
