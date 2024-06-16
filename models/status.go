package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	UserID        uint   `gorm:"not null"`
	Name          string `gorm:"not null,default:todo"`
	DefaultStatus bool   `gorm:"not null,dafault:false"`
}
