package models

import "gorm.io/gorm"

type Policy struct {
	gorm.Model
	Role    string `gorm:"not null"`
	UserID  uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StoreID uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
