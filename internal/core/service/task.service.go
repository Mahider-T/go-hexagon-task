package service

import (
	"go-hexagon-task/internal/core/domain"
	"go-hexagon-task/internal/core/port"
)

type TaskService struct {
	repo port.TaskRepository
}

func NewTaskService(repo port.TaskRepository) *TaskService {
	return &TaskService{
		repo,
	}
}

func (ts TaskService) AddTask(tsk *domain.Task) error {
	err := ts.repo.CreateTask(tsk)

	if err != nil {
		return err
	}

	return nil
}

func (ts TaskService) GetTask(id int) (*domain.Task, error) {

	tsk, err := ts.repo.GetTaskById(id)

	if err != nil {
		return nil, err
	}

	return tsk, nil
}

func (ts TaskService) UpdateTask(id int, tsk *domain.TaskStatus) error {

	err := ts.repo.UpdateTask(id, tsk)

	if err != nil {
		return err
	}

	return nil
}

func (ts TaskService) ListTask() ([]*domain.Task, error) {
	tasks, err := ts.repo.GetTasks()

	if err != nil {
		return nil, err
	}
	return tasks, nil
}
