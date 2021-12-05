package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name         string `gorm:"not null"`
	PhoneNumber  string `gorm:"not null"`
	Address      string `gorm:"not null"`
	ManagerName  string `gorm:"not null"`
	ActivityType string `gorm:"not null"`
	Necessities  []*Necessity
	Payments     []*Payment
	Users        []*User
	Bills        []*Bill
	Cheques      []*Cheque
	Orders       []*Order
	Categories   []*Category
}
