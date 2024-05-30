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

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) GetAll(userID uint) (*[]models.Item, error) {
	items := &[]models.Item{}
	result := r.db.Where("user_id = ?", userID).Find(items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}
