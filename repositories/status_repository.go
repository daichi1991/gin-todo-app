package repositories

import (
	"gin-todo-app/models"

	"gorm.io/gorm"
)

type IStatusRepository interface {
	CreateStatus(newStatus models.Status) (*models.Status, error)
	FindAllStatus() (*[]models.Status, error)
}

type StatusRepository struct {
	db *gorm.DB
}

// CreateStatus implements IStatusRepository.
func (r *StatusRepository) CreateStatus(newStatus models.Status) (*models.Status, error) {
	result := r.db.Create(&newStatus)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newStatus, nil
}

// FindAllStatus implements IStatusRepository.
func (r *StatusRepository) FindAllStatus() (*[]models.Status, error) {
	var statuses []models.Status
	result := r.db.Find(&statuses)
	if result.Error != nil {
		return nil, result.Error
	}
	return &statuses, nil
}

func NewStatusRepository(db *gorm.DB) IStatusRepository {
	return &StatusRepository{
		db: db,
	}
}
