package repositories

import (
	"errors"
	"gin-todo-app/models"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user models.User) (*models.User, error)
	FindUser(email string) (*models.User, error)
	FindUserByID(userID uint) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(user models.User) (*models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *AuthRepository) FindUser(email string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *AuthRepository) FindUserByID(userID uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "ID = ?", userID)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *AuthRepository) UpdateUser(user models.User) (*models.User, error) {
	result := r.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
