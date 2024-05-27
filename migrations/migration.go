package main

import (
	"gin-todo-app/infra"
	"gin-todo-app/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}, &models.Status{}); err != nil {
		panic("Failed to migrate database")
	}
}
