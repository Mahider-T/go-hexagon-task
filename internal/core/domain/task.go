package domain

import "time"

type TaskStatus string

const (
	completed TaskStatus = "COMPLETED"
	underway  TaskStatus = "UNDERWAY"
	todo      TaskStatus = "TODO"
)

type Task struct {
	Id          int
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
