package db

import (
	"database/sql"

	"github.com/acgriswold/golang-task-manager/internal/tables"
)

type taskDb struct {
	db        *sql.DB
	directory string
}

func (t *taskDb) tableExists(name string) bool {
	if rows, err := t.db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='?'", name); err == nil {
		return rows.Next()
	}

	return false
}

func (t *taskDb) getTables() (*tables.Tables, error) {
	tasksTable := tables.TasksTable{}
	if !t.tableExists("tasks") {
		err := tasksTable.Create(t.db)

		if err != nil {
			return nil, err
		}
	}

	tables := tables.Tables{Db: t.db, Tasks: &tasksTable}

	return &tables, nil
}
