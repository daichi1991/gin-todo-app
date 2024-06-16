package repositories

import (
	"gin-todo-app/models"

	"gorm.io/gorm"
)

type IStatusRepository interface {
	CreateStatus(newStatus models.Status) (*models.Status, error)
	FindAllStatus(userID uint) (*[]models.Status, error)
	FindDefaultStatus(userID uint) (*models.Status, error)
	FindStatusByID(statusID uint) (*models.Status, error)
	UpdateStatus(updateStatus models.Status) (*models.Status, error)
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
func (r *StatusRepository) FindAllStatus(userID uint) (*[]models.Status, error) {
	var statuses []models.Status
	result := r.db.Where("user_id = ?", userID).Find(&statuses)
	if result.Error != nil {
		return nil, result.Error
	}
	return &statuses, nil
}

func (r *StatusRepository) FindDefaultStatus(userID uint) (*models.Status, error) {
	status := models.Status{}
	result := r.db.Where("user_id = ? and default_status = true", userID).First(&status)
	if result.Error != nil {
		return nil, result.Error
	}
	return &status, nil
}

func (r *StatusRepository) FindStatusByID(statusID uint) (*models.Status, error) {
	status := models.Status{}
	result := r.db.Where("id = ?", statusID).First(&status)
	if result.Error != nil {
		return nil, result.Error
	}
	return &status, nil
}

func (r *StatusRepository) UpdateStatus(updateStatus models.Status) (*models.Status, error) {
	result := r.db.Save(&updateStatus)
	if result.Error != nil {
		return nil, result.Error
	}

	return &updateStatus, nil
}

func NewStatusRepository(db *gorm.DB) IStatusRepository {
	return &StatusRepository{
		db: db,
	}
}
