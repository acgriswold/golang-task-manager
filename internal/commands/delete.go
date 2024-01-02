package commands

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/acgriswold/golang-task-manager/internal/db"
)

var delete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDb()

		if err != nil {
			return err
		}

		defer t.Db.Close()

		id, err := strconv.Atoi(args[0])

		if err != nil {
			return err
		}

		return t.Tasks.Delete(t.Db, uint(id))
	},
}
