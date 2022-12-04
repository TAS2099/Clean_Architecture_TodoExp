package usecase

import "github.com/TAS2099/clean_architecture_todoexp/app/domain"

type TaskRepository interface {
	CreateTask(task *domain.Task) *domain.Task
}
