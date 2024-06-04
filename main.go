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
	statusRepository := repositories.NewStatusRepository(db)
	statusService := services.NewStatusRepository(statusRepository)
	statusController := controllers.NewStatusController(statusService)

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository, statusService)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository, statusService)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")
	authRouterWithAuth := r.Group("/auth", middlewares.AuthMiddleware(authService))
	statusRouter := r.Group("/statuses")
	statusRouterWithAuth := r.Group("/statuses", middlewares.AuthMiddleware(authService))

	itemRouter.GET("", itemController.GetAll)
	itemRouterWithAuth.POST("", itemController.Create)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)
	authRouterWithAuth.PUT("/update", authController.Update)

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
