package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Seller      string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	UnitPrice   string `gorm:"not null"`
	Description string `gorm:"not null"`
	StoreID     uint
}
