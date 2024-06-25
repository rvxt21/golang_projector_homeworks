package main

import (
	"errors"
	"time"
)

type Priority string

const (
	ExtremelyPriority    Priority = "Extremely"
	HighPriority         Priority = "High"
	MediumProirity       Priority = "Medium"
	LowPriority          Priority = "Low"
	NotNecessaryPriority Priority = "Not Necessery"
	NotSpecifiedPriority Priority = ""
)

func (p Priority) IsValid() error {
	if p == "" {
		return nil
	}
	switch p {
	case ExtremelyPriority, HighPriority, MediumProirity, LowPriority, NotNecessaryPriority, NotSpecifiedPriority:
		return nil
	default:
		return errors.New("invalid priority")
	}
}

type Status string

const (
	ToDoStatus         Status = "ToDo"
	InProgressStatus   Status = "In progress"
	DoneStatus         Status = "Done"
	ClosedStatus       Status = "Closed"
	NotSpecifiedStatus Status = ""
)

func (s Status) IsValid() error {
	switch s {
	case ToDoStatus, InProgressStatus, DoneStatus, ClosedStatus, NotSpecifiedStatus:
		return nil
	default:
		return errors.New("invalid status")
	}

}

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    Priority
	Status      Status
	CreatedAt   time.Time
	DueDate     time.Time
}

type Role string

const (
	StudentRole     Role = "Student"
	WorkerRole      Role = "Worker"
	SchoolchildRole Role = "Schoolchild"
)

func (r Role) IsValid() error {
	switch r {
	case StudentRole, WorkerRole, SchoolchildRole:
		return nil
	default:
		return errors.New("invalid user role")
	}
}

type User struct {
	Username string
	Password string
	ID       int
	Name     string
	Age      uint8
	Role     Role
}
