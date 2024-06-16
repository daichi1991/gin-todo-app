package main

import (
	"encoding/json"
	"gin-todo-app/infra"
	"gin-todo-app/models"
	"gin-todo-app/services"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "テストタスク1", Description: "", UserID: 1, StatusID: 1},
		{Name: "テストタスク2", Description: "テスト2", UserID: 1, StatusID: 2},
		{Name: "テストタスク3", Description: "テスト3", UserID: 2, StatusID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	statuses := []models.Status{
		{Name: "todo", UserID: 1, DefaultStatus: true},
		{Name: "doing", UserID: 1, DefaultStatus: false},
		{Name: "pending", UserID: 2, DefaultStatus: false},
	}

	for _, user := range users {
		db.Create(&user)
	}
	for _, status := range statuses {
		db.Create(&status)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{}, &models.Status{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindByID(t *testing.T) {
	router := setup()

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/1", nil)
	req.Header.Set("Authorization", "Bearer "+*token)

	router.ServeHTTP(w, req)

	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(res["data"]))
}
