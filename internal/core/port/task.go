package port

import (
	"go-hexagon-task/internal/core/domain"
)

type TaskRepository interface {
	GetTaskById(id string) (domain.Task, error)
	CreateTask(domain.Task) (domain.Task, error)
	GetTasks() ([]domain.Task, error)
	UpdateTask(id string, tsk domain.Task)
}

type TaskService interface {
	AddTask(tsk domain.Task) (domain.Task, error)
	GetTask(id string) (domain.Task, error)
	UpdateTask(id string, tsk domain.Task) (domain.Task, error)
	ListTask() ([]domain.Task, error)
}
