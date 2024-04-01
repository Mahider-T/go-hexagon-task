package storage

import (
	"go-hexagon-task/internal/core/domain"
)

type TaskRepository struct {
	db *Database
}

func NewTaskRepository(db *Database) *TaskRepository {
	return &TaskRepository{
		db,
	}
}

func (ts TaskRepository) GetTaskById(id int) (*domain.Task, error) {
	stmt := `SELECT * FROM tasks WHERE id = $1`

	task := &domain.Task{}
	err := ts.db.db.QueryRow(stmt, id).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (ts TaskRepository) CreateTask(tsk *domain.Task) error {
	stmt := `INSERT INTO tasks(id, title, description, status, createdat) VALUES($1, $2, $3, $4, now())`

	_, err := ts.db.db.Exec(stmt, tsk.Id, tsk.Title, tsk.Description, tsk.Status)

	if err != nil {
		return err
	}

	return nil
}

func (ts TaskRepository) GetTasks() ([]*domain.Task, error) {

	stmt := `SELECT * FROM tasks`

	var tasks = []*domain.Task{}

	rows, err := ts.db.db.Query(stmt)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := &domain.Task{}

		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)

	}

	return tasks, nil
}

func (ts TaskRepository) UpdateTask(id int, sts *domain.TaskStatus) error {
	stmt := `UPDATE tasks SET status = $1 WHERE id = $2`

	_, err := ts.db.db.Exec(stmt, sts, id)

	if err != nil {
		return nil
	}

	return err
}
