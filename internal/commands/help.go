package commands

import "github.com/spf13/cobra"

var Help = &cobra.Command{
	Use:   "tasks",
	Short: "A CLI task managemnet tool for ~staying~ on top of tasks",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	/**
	Used within main.go to setup entry point.
	Initialize and add all valid cobra commands to be used
	within the cli.
	*/

	Help.AddCommand(add)
}
