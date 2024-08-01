package storage

import (
	"database/sql"
	"fmt"
	"hw10/enteties"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func New(connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("openning database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return &Database{db: db}, nil
}

func (db *Database) InsertTask(task enteties.Task) error {
	query := `INSERT INTO tasks(title,descriptions,priority,status,createdAt,dueDate) 
			VALUES($1, $2, $3, $4, $5, $6)`
	_, err := db.db.Exec(query, task.Title, task.Description, task.Priority, task.Status, task.CreatedAt, task.DueDate)
	if err != nil {
		return fmt.Errorf("inserting task: %w", err)
	}

	return nil
}

func (db *Database) GetAllTasks() ([]enteties.Task, error) {
	query := `"SELECT id,title,descriptions,priority,status,createdAt,dueDate 
			   FROM tasks"`
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("selecting tasks: %w", err)
	}

	var tasks []enteties.Task

	for rows.Next() {
		var t enteties.Task

		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Priority, &t.Status, &t.CreatedAt, &t.DueDate)
		if err != nil {
			return nil, fmt.Errorf("scanning rows: %w", err)
		}
		tasks = append(tasks, t)

	}

	return tasks, nil
}
