package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	StatusID    uint   `gorm:"not null"`
	Status      Status `gorm:"foreignKey:StatusID"`
	UserID      uint   `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID"`
}
