package models

import "gorm.io/gorm"

type Necessity struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Supplier    string
	ProductID   uint
	Product     *Product
	StoreID     uint `gorm:"not null"`
}
