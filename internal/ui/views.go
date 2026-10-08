package ui

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Logo = `
  ____                            __     __          _ _
 / ___|  ___  _   _ _ __ ___ ___  \ \   / /_ _ _   _| | |_
 \___ \ / _ \| | | | '__/ __/ _ \  \ \ / / _` + "`" + ` | | | | | __|
  ___) | (_) | |_| | | | (_|  __/   \ V / (_| | |_| | | |_
 |____/ \___/ \__,_|_|  \___\___|    \_/ \__,_|\__,_|_|\__|
`

func HandleHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Usage()
}

func HandleUsage(cmd *cobra.Command) error {
	fmt.Println(LogoStyle.Render(Logo))
	fmt.Println(HelpTitleStyle.Render(cmd.Short))
	if cmd.Long != "" {
		fmt.Println(HelpDescStyle.Render(cmd.Long))
	}

	fmt.Println(HelpSectionStyle.Render("USAGE"))
	if cmd.HasSubCommands() {
		fmt.Printf("  %s [command]\n", cmd.CommandPath())
	} else {
		fmt.Printf("  %s\n", cmd.UseLine())
	}

	if len(cmd.Commands()) > 0 {
		fmt.Println(HelpSectionStyle.Render("AVAILABLE COMMANDS"))
		for _, c := range cmd.Commands() {
			if !c.Hidden {
				fmt.Printf("  %-15s %s\n", c.Name(), c.Short)
			}
		}
	}

	if cmd.Flags().HasFlags() {
		fmt.Println(HelpSectionStyle.Render("FLAGS"))
		fmt.Println(HelpFlagStyle.Render(cmd.Flags().FlagUsages()))
	}

	if len(cmd.Commands()) > 0 {
		fmt.Println(HelpSectionStyle.Render("LEARN MORE"))
		fmt.Printf("  Use \"%s [command] --help\" for more information about a command.\n", cmd.CommandPath())
	}
	return nil
}
