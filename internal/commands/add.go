package commands

import (
	"fmt"

	"github.com/acgriswold/golang-task-manager/internal/db"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add TASK",
	Short: "Add a new task with an optional project name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDb()

		if err != nil {
			return err
		}

		defer t.Db.Close()

		fmt.Printf("Adding task \"%s\"", args[0])
		fmt.Println()

		return nil
	},
}

func init() {
	add.Flags().StringP(
		"project",
		"p",
		"",
		"specify a project for your task",
	)
}
