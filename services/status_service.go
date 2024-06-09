package services

import (
	"errors"
	"gin-todo-app/dto"
	"gin-todo-app/models"
	"gin-todo-app/repositories"
)

type IStatusService interface {
	CreateStatus(CreateStatusInput dto.CreateStatusInput, userID uint) (*models.Status, error)
	CreateDefaultStatus(userID uint) (*models.Status, error)
	FindAllStatus() (*[]models.Status, error)
	FindDefaultStatus(userID uint) (*models.Status, error)
	UpdateStatus(UpdateStatusInput dto.UpdateStatusInput, itemID uint, userID uint) (*models.Status, error)
}

type StatusService struct {
	repository repositories.IStatusRepository
}

func (s *StatusService) CreateStatus(CreateStatusInput dto.CreateStatusInput, userID uint) (*models.Status, error) {
	newStatus := models.Status{
		UserID:  userID,
		Name:    CreateStatusInput.Name,
		Default: false,
	}
	return s.repository.CreateStatus(newStatus)
}

func (s *StatusService) CreateDefaultStatus(userID uint) (*models.Status, error) {
	defaultStatus := models.Status{
		UserID:  userID,
		Name:    "todo",
		Default: true,
	}

	return s.repository.CreateStatus(defaultStatus)
}

func (s *StatusService) FindAllStatus() (*[]models.Status, error) {
	return s.repository.FindAllStatus()
}

func (s *StatusService) FindDefaultStatus(userID uint) (*models.Status, error) {
	defaultStatus, err := s.repository.FindDefaultStatus(userID)
	if err != nil {
		return nil, err
	}
	return defaultStatus, nil
}

func (s *StatusService) UpdateStatus(UpdateStatusInput dto.UpdateStatusInput, statusID uint, userID uint) (*models.Status, error) {
	targetStatus, err := s.repository.FindStatusByID(statusID)
	if err != nil {
		return nil, err
	}
	if targetStatus.UserID != userID {
		return nil, errors.New("not permitted")
	}
	targetStatus.Name = UpdateStatusInput.Name
	_, err = s.repository.UpdateStatus(*targetStatus)
	if err != nil {
		return nil, err
	}
	return targetStatus, nil
}

func NewStatusRepository(repository repositories.IStatusRepository) IStatusService {
	return &StatusService{
		repository: repository,
	}
}
