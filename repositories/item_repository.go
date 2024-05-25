package repositories

import (
	"gin-todo-app/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	GetAll(userID uint) (*[]models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db}
}

func (r *ItemRepository) GetAll(userID uint) (*[]models.Item, error) {
	items := &[]models.Item{}
	if err := r.db.Where("user_id = ?", userID).Find(items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
