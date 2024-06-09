package controllers

import (
	"gin-todo-app/dto"
	"gin-todo-app/models"
	"gin-todo-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IStatusController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type StatusController struct {
	service services.IStatusService
}

// Create implements IStatusController.
func (c *StatusController) Create(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userID := user.(*models.User).ID
	var input dto.CreateStatusInput
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newStatus, err := c.service.CreateStatus(input, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": newStatus})
}

// FindAll implements IStatusController.
func (c *StatusController) FindAll(ctx *gin.Context) {
	statuses, err := c.service.FindAllStatus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": statuses})
}

func (c *StatusController) Update(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userID := user.(*models.User).ID
	var input dto.UpdateStatusInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	statusID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, err := c.service.UpdateStatus(input, uint(statusID), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": status})
}

func NewStatusController(service services.IStatusService) IStatusController {
	return &StatusController{
		service: service,
	}
}
