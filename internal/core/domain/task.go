package domain

import "time"

type TaskStatus string

const (
	completed TaskStatus = "COMPLETED"
	underway  TaskStatus = "UNDERWAY"
	todo      TaskStatus = "TODO"
)

type Task struct {
	id          string
	title       string
	description string
	status      TaskStatus
	createdAt   time.Time
	updatedAt   time.Time
}
