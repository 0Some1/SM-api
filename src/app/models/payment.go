package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Person      string `gorm:"not null"`
	Amount      int64 `gorm:"not null"`
	Description string `gorm:"not null"`
	StoreID     uint

}
