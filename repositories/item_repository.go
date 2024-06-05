package repositories

import (
	"errors"
	"gin-todo-app/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll(userID uint) (*[]models.Item, error)
	FindByID(itemID uint, userID uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updatedItem models.Item) (*models.Item, error)
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) FindAll(userID uint) (*[]models.Item, error) {
	items := &[]models.Item{}
	result := r.db.Where("user_id = ?", userID).Find(items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (r *ItemRepository) FindByID(itemID uint, userID uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, "id = ? AND user_id = ?", itemID, userID)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (r *ItemRepository) Update(updatedItem models.Item) (*models.Item, error) {
	result := r.db.Save(&updatedItem)
	if result.Error != nil {
		return nil, result.Error
	}

	return &updatedItem, nil
}
