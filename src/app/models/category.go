package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name            string `gorm:"not null"`
	NodeId          uint64 `gorm:"index;unique;not null"`
	StoreID         uint
	CategoryOptions []*CategoryOption
	Products        []*Product
}
