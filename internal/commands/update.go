package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/acgriswold/golang-task-manager/internal/db"
	"github.com/acgriswold/golang-task-manager/internal/tables"
	"github.com/spf13/cobra"
)

var update = &cobra.Command{
	Use:   "update",
	Short: "Update a task by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDb()

		if err != nil {
			return err
		}

		defer t.Db.Close()

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}

		status, err := cmd.Flags().GetInt("status")
		if err != nil {
			return err
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		var stat string
		switch status {
		case int(tables.InProgress):
			stat = tables.InProgress.String()
		case int(tables.Done):
			stat = tables.Done.String()
		default:
			stat = tables.Todo.String()
		}

		newTask := tables.Task{ID: uint(id), Name: name, Project: project, Status: stat, Created: time.Time{}}

		return t.Tasks.Update(t.Db, newTask)
	},
}

func init() {
	update.Flags().
		StringP(
			"project",
			"p",
			"",
			"specify a project for your task",
		)

	update.Flags().
		StringP(
			"name",
			"n",
			"",
			"specify the name of the task",
		)

	update.Flags().
		IntP(
			"status",
			"s",
			int(tables.Todo),
			fmt.Sprintf("specify the status of the task (i.e., %d=\"Todo\", %d=\"InProgress\", %d=\"Done\")", tables.Todo.Int(), tables.InProgress.Int(), tables.Done.Int()),
		)
}
