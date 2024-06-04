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
	repository    repositories.IItemRepository
	statusService IStatusService
}

func NewItemService(repository repositories.IItemRepository, statusService IStatusService) IItemService {
	return &ItemService{
		repository:    repository,
		statusService: statusService,
	}
}

func (s *ItemService) GetAll(userID uint) (*[]models.Item, error) {
	return s.repository.GetAll(userID)
}

func (s *ItemService) Create(CreateItemInput dto.CreateItemInput, userID uint) (*models.Item, error) {
	defaultStatus, err := s.statusService.FindDefaultStatus(userID)
	if err != nil {
		return nil, err
	}

	newItem := models.Item{
		Name:        CreateItemInput.Name,
		Description: CreateItemInput.Description,
		UserID:      userID,
		StatusID:    defaultStatus.ID,
	}
	return s.repository.Create(newItem)
}
