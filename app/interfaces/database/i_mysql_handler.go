package database

import "github.com/TAS2099/clean_architecture_todoexp/app/domain"

type IMySQLHandler interface {
	AddTask(task *domain.Task) *domain.Task
}
