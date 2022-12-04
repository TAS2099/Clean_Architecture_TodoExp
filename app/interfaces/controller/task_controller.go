package controller

import (
	"github.com/TAS2099/clean_architecture_todoexp/app/domain"
	"github.com/TAS2099/clean_architecture_todoexp/app/interfaces/database"
	"github.com/TAS2099/clean_architecture_todoexp/app/usecase"
	"github.com/labstack/echo"
)

type TaskController struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(mysqlHandler database.IMySQLHandler) *TaskController {
	return &TaskController{
		Interactor: usecase.TaskInteractor{
			Repository: &database.TaskRepository{
				IMySQLHandler: mysqlHandler,
			},
		},
	}
}

func (tc *TaskController) CreateTask(c echo.Context) *domain.Task {
	task := new(domain.Task)

	c.Bind(task)

	ct := tc.Interactor.CreateTask(task)

	return ct
}
