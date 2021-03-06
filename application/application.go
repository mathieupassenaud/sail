package application

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdApplicationList)
	Cmd.AddCommand(cmdApplicationShow)

	cmdApplicationDomain.AddCommand(cmdApplicationDomainList)
	cmdApplicationDomain.AddCommand(cmdApplicationDomainAttach)
	cmdApplicationDomain.AddCommand(cmdApplicationDomainDetach)

	Cmd.AddCommand(cmdApplicationDomain)
}

// Cmd application
var Cmd = &cobra.Command{
	Use:     "application",
	Short:   "Application commands: sail application --help",
	Long:    `Application commands: sail application <command>`,
	Aliases: []string{"a", "app", "apps", "applications"},
}
