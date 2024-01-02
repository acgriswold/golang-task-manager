package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"fmt"
	"log"
	"os"

	"github.com/acgriswold/golang-task-manager/internal/tables"
	gap "github.com/muesli/go-app-paths"
)

func _setupDbPath() string {
	scope := gap.NewScope(gap.User, "tasks")
	dirs, err := scope.DataDirs()

	if err != nil {
		log.Fatal(err)
	}

	var taskDir string
	if len(dirs) > 0 {
		taskDir = dirs[0]
	} else {
		taskDir, _ = os.UserHomeDir()
	}

	return taskDir
}

func _initTaskDirectory(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, 0o770)
			return err
		}

		return err
	}

	return nil
}

func OpenDb() (*tables.Tables, error) {
	path := _setupDbPath()

	if err := _initTaskDirectory(path); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/tasks.db", path))

	if err != nil {
		return nil, err
	}

	t := taskDb{db, path}

	return t.getTables()
}
