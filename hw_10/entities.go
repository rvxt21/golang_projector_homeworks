package main

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	DueDate     time.Time
}

// type UpdateTask struct {
// 	Title       *string    `json:"title"`
// 	Description *string    `json:"description"`
// 	Status      *string    `json:"status"`
// 	DueDate     *time.Time `json:"due_date"`
// }
