package services

import (
	"gin-todo-app/models"
	"gin-todo-app/repositories"
)

type IStatusService interface {
	CreateStatus(name string) (*models.Status, error)
	FindAllStatus() (*[]models.Status, error)
}

type StatusService struct {
	repository repositories.IStatusRepository
}

func (s *StatusService) CreateStatus(name string) (*models.Status, error) {
	panic("")
}

func (s *StatusService) FindAllStatus() (*[]models.Status, error) {
	panic("")
}

func NewStatusRepository(repository repositories.IStatusRepository) IStatusService {
	return &StatusService{
		repository: repository,
	}
}
