package infra

import (
	"net/http"

	"github.com/TAS2099/clean_architecture_todoexp/app/interfaces/controller"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func InitalizeServer(db *gorm.DB) *echo.Echo {

	e := echo.New()
	taskController := controller.NewTaskController(NewMySQLHandler(db))

	e.POST("/task", func(c echo.Context) error {
		task := taskController.CreateTask(c)
		return c.JSON(http.StatusCreated, task)
	})

	return e
}
