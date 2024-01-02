package commands

import (
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

		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}

		if err := t.Tasks.Insert(t.Db, args[0], project); err != nil {
			return err
		}

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
