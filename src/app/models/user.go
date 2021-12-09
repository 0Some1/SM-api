package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Username    string `gorm:"not null;unique"`
	PhoneNumber string `gorm:"not null;unique"`
	Email       string `gorm:"not null;unique"`
	StoreID     uint   `gorm:"not null"`
	Policies    []*Policy
}
