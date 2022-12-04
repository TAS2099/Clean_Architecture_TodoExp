package usecase

import "github.com/TAS2099/clean_architecture_todoexp/app/domain"

type TaskInteractor struct {
	Repository TaskRepository
}

func (ti *TaskInteractor) CreateTask(task *domain.Task) *domain.Task {
	return ti.Repository.CreateTask(task)
}
