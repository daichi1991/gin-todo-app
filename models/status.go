package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
	Name   string `gorm:"not null,default:todo"`
}
