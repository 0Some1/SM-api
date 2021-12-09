package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	IsEntry     bool   `gorm:"default:true"`
	Executor    string `gorm:"default:true"`
	ProductID   uint   `gorm:"not null"`
	Description string
	Product     Product
}
