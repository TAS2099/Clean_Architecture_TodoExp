package database

import "github.com/TAS2099/clean_architecture_todoexp/app/domain"

type TaskRepository struct {
	IMySQLHandler
}

func (tr *TaskRepository) CreateTask(task *domain.Task) *domain.Task {
	return tr.AddTask(task)
}
