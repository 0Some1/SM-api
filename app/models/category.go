package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name            string `gorm:"not null"`
	NodeId          uint64 `gorm:"index;not null"`
	StoreID         uint   `gorm:"not null"`
	CategoryOptions []*CategoryOption
	Products        []*Product
}
