package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name                 string `gorm:"not null"`
	Price                string `gorm:"not null"`
	Quantity             int    `gorm:"not null"`
	MinQuantity          int    `gorm:"not null"`
	CategoryID           uint
	Reports              []*Report
	CategoryOptionValues []*CategoryOptionValue `gorm:"many2many:products_CategoryOptionValues;"`
}
