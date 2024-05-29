package services

import (
	"gin-todo-app/dto"
	"gin-todo-app/models"
	"gin-todo-app/repositories"
)

type IStatusService interface {
	CreateStatus(CreateStatusInput dto.CreateStatusInput) (*models.Status, error)
	FindAllStatus() (*[]models.Status, error)
}

type StatusService struct {
	repository repositories.IStatusRepository
}

func (s *StatusService) CreateStatus(CreateStatusInput dto.CreateStatusInput) (*models.Status, error) {
	newStatus := models.Status{
		Name: CreateStatusInput.Name,
	}
	return s.repository.CreateStatus(newStatus)
}

func (s *StatusService) FindAllStatus() (*[]models.Status, error) {
	return s.repository.FindAllStatus()
}

func NewStatusRepository(repository repositories.IStatusRepository) IStatusService {
	return &StatusService{
		repository: repository,
	}
}
