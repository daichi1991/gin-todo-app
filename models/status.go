package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Name    string `gorm:"not null,default:todo"`
	Default bool   `gorm:"not null,dafault:false"`
}
