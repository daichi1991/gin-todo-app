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
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")
	authRouterWithAuth := r.Group("/auth", middlewares.AuthMiddleware(authService))
	statusRouterWithAuth := r.Group("/statuses", middlewares.AuthMiddleware(authService))

	itemRouterWithAuth.GET("", itemController.FindAll)
	itemRouterWithAuth.GET(":id", itemController.FindByID)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT(":id", itemController.Update)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)
	authRouterWithAuth.PUT("/update", authController.Update)

	statusRouterWithAuth.GET("", statusController.FindAll)
	statusRouterWithAuth.POST("", statusController.Create)
	statusRouterWithAuth.PUT(":id", statusController.Update)

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := setupRouter(db)

	r.Run("localhost:8080")
}
