package controllers

import (
	"gin-todo-app/services"
	"net/http"

	"gin-todo-app/dto"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
}

type ItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &ItemController{
		service: service,
	}
}

func (c *ItemController) GetAll(ctx *gin.Context) {
	var userID uint = 1
	items, err := c.service.GetAll(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func (c *ItemController) Create(ctx *gin.Context) {
	var userID uint = 1
	var input dto.CreateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := c.service.Create(input, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}
