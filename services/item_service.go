package services

import (
	"gin-todo-app/dto"
	"gin-todo-app/models"
	"gin-todo-app/repositories"
)

type IItemService interface {
	GetAll(userID uint) (*[]models.Item, error)
	Create(CreateItemInput dto.CreateItemInput, userID uint) (*models.Item, error)
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{
		repository: repository,
	}
}

func (s *ItemService) GetAll(userID uint) (*[]models.Item, error) {
	return s.repository.GetAll(userID)
}

func (s *ItemService) Create(CreateItemInput dto.CreateItemInput, userID uint) (*models.Item, error) {
	newItem := models.Item{
		Name:        CreateItemInput.Name,
		Description: CreateItemInput.Description,
		UserID:      userID,
	}
	return s.repository.Create(newItem)
}
