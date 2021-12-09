package models

import "gorm.io/gorm"

type CategoryOption struct {
	gorm.Model
	Title                string `gorm:"not null"`
	Required             bool   `gorm:"default:false;"`
	Description          string
	CategoryID           uint `gorm:"not null"`
	CategoryOptionValues []*CategoryOptionValue
}
