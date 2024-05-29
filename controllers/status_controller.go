package controllers

import (
	"gin-todo-app/dto"
	"gin-todo-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IStatusController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type StatusController struct {
	service services.IStatusService
}

// Create implements IStatusController.
func (c *StatusController) Create(ctx *gin.Context) {
	var input dto.CreateStatusInput
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newStatus, err := c.service.CreateStatus(input)
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

func NewStatusController(service services.IStatusService) IStatusController {
	return &StatusController{
		service: service,
	}
}
