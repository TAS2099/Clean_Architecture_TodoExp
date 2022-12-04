package infra

import (
	"github.com/TAS2099/clean_architecture_todoexp/app/domain"
	"github.com/TAS2099/clean_architecture_todoexp/app/interfaces/database"
	"gorm.io/gorm"
)

type MySQLHandler struct {
	db *gorm.DB
}

func NewMySQLHandler(db *gorm.DB) database.IMySQLHandler {

	mysqlHandler := new(MySQLHandler)
	mysqlHandler.db = db

	return mysqlHandler
}

func (mh *MySQLHandler) AddTask(task *domain.Task) *domain.Task {
	if err := mh.db.Create(task).Error; err != nil {
		return nil
	}
	return task
}
