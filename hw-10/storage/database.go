package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"hw10/enteties"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase(connStr string) (*Database, error) {
	Db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("openning database: %w", err)
	}
	if err := Db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return &Database{Db: Db}, nil
}

func (Db *Database) InsertTask(task enteties.Task) error {
	query := `INSERT INTO tasks(title,description,priority,status,createdAt,dueDate) 
			VALUES($1, $2, $3, $4, $5, $6)`
	_, err := Db.Db.Exec(query, task.Title, task.Description, task.Priority, task.Status, task.CreatedAt, task.DueDate)
	if err != nil {
		return fmt.Errorf("inserting task: %w", err)
	}

	return nil
}

func (Db *Database) GetAllTasks() ([]enteties.Task, error) {
	query := `SELECT * 
			   FROM tasks`
	rows, err := Db.Db.Query(query)
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

var ErrDeleteFailed error = errors.New("errot to delete task")

func (db *Database) DeleteTask(id int) (bool, error) {
	query := `DELETE FROM tasks WHERE id = $1`
	res, err := db.Db.Exec(query, id)
	if err != nil {
		log.Error().Err(err).Msgf("unable to delete category")
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, ErrDeleteFailed
	}

	return true, nil
}

func (db *Database) GetTaskByID(id int) (enteties.Task, bool, error) {
	query := `SELECT id, title, description, priority, status, createdAt, dueDate 
			  FROM tasks 
			  WHERE id = $1`
	var t enteties.Task
	err := db.Db.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Description, &t.Priority, &t.Status, &t.CreatedAt, &t.DueDate)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msgf("unable to get category")
			return t, false, nil
		}
		return t, false, err
	}
	return t, true, nil
}
