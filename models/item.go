package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	StatusID    uint `gorm:"not null"`
	UserID      uint `gorm:"not null"`
}
