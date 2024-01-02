package tables

import "database/sql"

type Tables struct {
	Db    *sql.DB
	Tasks *TasksTable
}
