package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	IsEntry     bool   `gorm:"default:true"`
	Executor    string `gorm:"default:true"`
	Description string
	ProductID   uint
	Product     Product
}
