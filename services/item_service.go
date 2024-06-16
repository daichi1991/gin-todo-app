package services

import (
	"errors"
	"gin-todo-app/dto"
	"gin-todo-app/models"
	"gin-todo-app/repositories"
)

type IItemService interface {
	FindAll(userID uint) (*[]models.Item, error)
	FindByID(itemID uint, userID uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput, userID uint) (*models.Item, error)
	Update(updateItemInput dto.UpdateItemInput, itemID uint, userID uint) (*models.Item, error)
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

func (s *ItemService) FindAll(userID uint) (*[]models.Item, error) {
	return s.repository.FindAll(userID)
}

func (s *ItemService) FindByID(itemID uint, userID uint) (*models.Item, error) {
	return s.repository.FindByID(itemID, userID)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput, userID uint) (*models.Item, error) {
	defaultStatus, err := s.statusService.FindDefaultStatus(userID)
	if err != nil {
		return nil, err
	}

	newItem := models.Item{
		Name:        createItemInput.Name,
		Description: createItemInput.Description,
		UserID:      userID,
		StatusID:    defaultStatus.ID,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(updateItemInput dto.UpdateItemInput, itemID uint, userID uint) (*models.Item, error) {
	targetItem, err := s.repository.FindByID(itemID, userID)
	if err != nil {
		return nil, err
	}
	if targetItem.UserID != userID {
		return nil, errors.New("not parmitted")
	}
	targetItem.Name = updateItemInput.Name
	targetItem.Description = updateItemInput.Description
	targetItem.StatusID = updateItemInput.StatusID
	_, err = s.repository.Update(*targetItem)
	if err != nil {
		return nil, err
	}
	return targetItem, nil
}
