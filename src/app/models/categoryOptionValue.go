package models

import "gorm.io/gorm"

type CategoryOptionValue struct {
	gorm.Model
	Value            string     `gorm:"not null"`
	CategoryOptionID uint       `gorm:"not null"`
	Products         []*Product `gorm:"many2many:products_CategoryOptionValues;"`
}
