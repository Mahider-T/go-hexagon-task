package port

import (
	"go-hexagon-task/internal/core/domain"
)

type TaskRepository interface {
	GetTaskById(id int) (*domain.Task, error)
	CreateTask(*domain.Task) error
	GetTasks() ([]*domain.Task, error)
	UpdateTask(id int, tsk *domain.TaskStatus) error
}

type TaskService interface {
	AddTask(tsk *domain.Task) error
	GetTask(id int) (*domain.Task, error)
	UpdateTask(id int, sts *domain.TaskStatus) error
	ListTask() ([]*domain.Task, error)
}
