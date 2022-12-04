package main

import (
	"log"
	"net/http"

	"github.com/TAS2099/clean_architecture_todoexp/app/domain"
	"github.com/TAS2099/clean_architecture_todoexp/app/infra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (db *gorm.DB) {
	dsn := "user:password@tcp(localhost:3306)/task_db?charset=utf8mb4&parseTime=True&loc=Local"
	dbOpen := mysql.Open(dsn)
	config := &gorm.Config{}

	db, err := gorm.Open(dbOpen, config)
	if err != nil {
		panic(err)
	}

	db.Migrator().CreateTable(domain.Task{})

	return db
}

func main() {
	db := ConnectDB()
	server := infra.InitalizeServer(db)
	if err := server.Start("0.0.0.0:1323"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
