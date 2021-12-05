package models

import "gorm.io/gorm"

type CategoryOptionValue struct {
	gorm.Model
	Value            string `gorm:"not null"`
	CategoryOptionID uint
	Products         []*Product `gorm:"many2many:products_CategoryOptionValues;"`
}
