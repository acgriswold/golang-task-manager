package commands

import (
	"fmt"

	"github.com/acgriswold/golang-task-manager/internal/db"
	"github.com/acgriswold/golang-task-manager/internal/tables"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"

	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "List all your tasks",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDb()

		if err != nil {
			return err
		}

		defer t.Db.Close()

		tasks, err := t.Tasks.GetAll(t.Db)
		if err != nil {
			return err
		}

		table := setupTable(tasks)

		fmt.Println(table.View())
		return nil
	},
}

func setupTable(tasks []tables.Task) table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 3},
		{Title: "Name", Width: 45},
		{Title: "Project", Width: 20},
		{Title: "Status", Width: 15},
		{Title: "Created At", Width: 10},
	}

	var rows []table.Row
	for _, task := range tasks {
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", task.ID),
			task.Name,
			task.Project,
			task.Status,
			task.Created.Format("2006-01-02"),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(len(tasks)),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	t.SetStyles(s)

	return t
}
