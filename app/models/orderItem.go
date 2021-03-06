package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Quantity    int    `gorm:"not null"`
	UnitPrice   int64  `gorm:"not null"`
	Discount    int64  `gorm:"not null"`
	Description string `gorm:"not null"`
	OrderID     uint   `gorm:"not null"`
	ProductID   uint   `gorm:"not null"`
	Product     *Product
}
