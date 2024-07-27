package storage

import (
	"database/sql"
	"fmt"
	"hw10/enteties"
	"time"

	_ "github.com/lib/pq"
)

type DBStorage struct {
	db *sql.DB
}

func New(connStr string) (*DBStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("openning database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return &DBStorage{db: db}, nil
}

func (db *DBStorage) InsertTask(id int, title string, descr string, priority string, status string, createdAt time.Time, dueTo time.Time) error {
	_, err := db.db.Exec("INSERT INTO tasks(id,title,descriptions,priority,status,createdAt,dueDate) VALUES($1, $2, $3, $4, $5, $6, $7)", id, title, descr, priority, status, createdAt, dueTo)
	if err != nil {
		return fmt.Errorf("inserting task: %w", err)
	}

	return nil
}

func (db *DBStorage) GetAllTasks() ([]enteties.Task, error) {
	rows, err := db.db.Query("SELECT id,title,descriptions,priority,status,createdAt,dueDate FROM tasks")
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
