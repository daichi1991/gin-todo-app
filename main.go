package main

import (
	"gin-todo-app/controllers"
	"gin-todo-app/infra"
	"gin-todo-app/middlewares"
	"gin-todo-app/repositories"
	"gin-todo-app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	statusRepository := repositories.NewStatusRepository(db)
	statusService := services.NewStatusRepository(statusRepository)
	statusController := controllers.NewStatusController(statusService)

	r := gin.Default()
	r.Use(cors.Default())
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")
	statusRouter := r.Group("/statuses")
	statusRouterWithAuth := r.Group("/statuses", middlewares.AuthMiddleware(authService))

	itemRouter.GET("", itemController.GetAll)
	itemRouterWithAuth.POST("", itemController.Create)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	statusRouter.GET("", statusController.FindAll)
	statusRouterWithAuth.POST("", statusController.Create)

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := setupRouter(db)

	r.Run("localhost:8080")
}
