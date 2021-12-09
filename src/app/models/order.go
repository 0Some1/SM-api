package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	customer            string    `gorm:"not null"`
	BuyDate             time.Time `gorm:"not null"`
	DeliveryDate        time.Time `gorm:"not null"`
	TotalPrice          int64     `gorm:"not null"`
	Pledge              int64     `gorm:"not null"`
	CustomerPhoneNumber string    `gorm:"not null"`
	StoreID             uint      `gorm:"not null"`
	Description         string
	OrderItems          []*OrderItem
	ProductID           uint
	Product             Product
}
