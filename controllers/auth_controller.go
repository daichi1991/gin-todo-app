package controllers

import (
	"fmt"
	"gin-todo-app/dto"
	"gin-todo-app/models"
	"gin-todo-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Signup(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

// Signup implements IAuthController.
func (c *AuthController) Signup(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

// Login implements IAuthController.
func (c *AuthController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *AuthController) Update(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID := user.(*models.User).ID

	var input dto.UpdateInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)
	result, err := c.service.UpdateUser(input, userID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{
		service: service,
	}
}
