package tables

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type TasksTable struct {
}

func (t *TasksTable) Create(db *sql.DB) error {
	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS 'tasks' ('id' INTEGER, 'name' TEXT NOT NULL, 'project' TEXT, 'status' Text, 'created' DATE, PRIMARY KEY('id' AUTOINCREMENT))",
	)

	return err
}

func (t *TasksTable) Insert(db *sql.DB, name, project string) error {
	_, err := db.Exec(
		"INSERT INTO 'tasks'(name, project, status, created) VALUES (?, ?, ?, ?)",
		name,
		project,
		"",
		time.Now(),
	)

	return err
}

func (t *TasksTable) Delete(db *sql.DB, id string) error {
	_, err := db.Exec(
		"DELETE FROM tasks WHERE id = ?",
		id,
	)

	return err
}

func (t *TasksTable) Update(db *sql.DB, task task) error {
	original, err := t.getTask(db, task.ID)

	if err != nil {
		return err
	}

	original.merge(task)
	_, err = db.Exec(
		"UPDATE tasks SET name = ?, project = ?, status = ?, WHERE id = ?",
		original.Name,
		original.Project,
		original.Status,
		original.ID,
	)

	return err
}

func (t *TasksTable) getAll(db *sql.DB) ([]task, error) {
	var tasks []task
	rows, err := db.Query("SELECT * from tasks")

	if err != nil {
		return tasks, fmt.Errorf("unable to get values: %w", err)
	}

	for rows.Next() {
		var task task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Created,
		)

		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	return tasks, err
}

func (t *TasksTable) getByStatus(db *sql.DB, status string) ([]task, error) {
	var tasks []task
	rows, err := db.Query("SELECT * from tasks WHERE status = ?", status)

	if err != nil {
		return tasks, fmt.Errorf("unable to get values: %w", err)
	}

	for rows.Next() {
		var task task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Created,
		)

		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	return tasks, err
}

func (t *TasksTable) getTask(db *sql.DB, id uint) (task, error) {
	var task task

	err := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).Scan(
		&task.ID,
		&task.Name,
		&task.Project,
		&task.Status,
		&task.Created,
	)

	return task, err
}

type task struct {
	ID      uint
	Name    string
	Project string
	Status  string
	Created time.Time
}

func (t task) FilterValue() string {
	return t.Name
}

func (t task) Title() string {
	return t.Name
}

func (t task) Description() string {
	return t.Project
}

func (original *task) merge(t task) {
	uValues := reflect.ValueOf(&t).Elem()
	oValues := reflect.ValueOf(original).Elem()

	for i := 0; i < uValues.NumField(); i++ {
		uField := uValues.Field(i).Interface()

		if oValues.CanSet() {
			if v, ok := uField.(int64); ok && uField != 0 {
				oValues.Field(i).SetInt(v)
			}

			if v, ok := uField.(string); ok && uField != "" {
				oValues.Field(i).SetString(v)
			}
		}
	}
}
