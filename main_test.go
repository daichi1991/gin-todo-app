package main

import (
	"encoding/json"
	"fmt"
	"gin-todo-app/infra"
	"gin-todo-app/models"
	"gin-todo-app/repositories"
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
	statusRepository := repositories.NewStatusRepository(db)
	statusService := services.NewStatusRepository(statusRepository)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository, statusService)

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	statuses := []models.Status{
		{Name: "doing", UserID: 1, DefaultStatus: false},
		{Name: "pending", UserID: 2, DefaultStatus: false},
	}

	items := []models.Item{
		{Name: "テストタスク1", Description: "", UserID: 1, StatusID: 1},
		{Name: "テストタスク2", Description: "テスト2", UserID: 1, StatusID: 2},
		{Name: "テストタスク3", Description: "テスト3", UserID: 2, StatusID: 2},
	}

	for _, user := range users {
		err := authService.Signup(user.Email, user.Password)
		if err != nil {
			fmt.Println(err)
		}
	}
	for _, status := range statuses {
		tx := db.Create(&status)
		if tx.Error != nil {
			fmt.Println(tx.Error)
		}
	}
	for _, item := range items {
		tx := db.Create(&item)
		if tx.Error != nil {
			fmt.Println(tx.Error)
		}
	}
}

func setup() *gin.Engine {
	db := infra.TestDB()
	db.AutoMigrate(&models.Item{}, &models.User{}, &models.Status{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindByID(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/1", nil)

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)
	bearerToken := "Bearer " + *token
	req.Header.Set("Authorization", bearerToken)

	router.ServeHTTP(w, req)

	var res models.Item
	var response struct {
		Data models.Item `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)
	res = response.Data
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "テストタスク1", res.Name)
}

func TestFindAll(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)
	bearerToken := "Bearer " + *token
	req.Header.Set("Authorization", bearerToken)

	router.ServeHTTP(w, req)

	var res []models.Item
	var response struct {
		Data []models.Item `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)
	res = response.Data
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 2, len(res))
}
