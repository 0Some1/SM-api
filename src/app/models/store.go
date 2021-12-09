package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name         string `gorm:"not null"`
	StorePhone   string `gorm:"not null;unique"`
	Address      string `gorm:"not null"`
	ActivityType string `gorm:"not null"`
	Necessities  []*Necessity
	Payments     []*Payment
	Users        []*User
	Bills        []*Bill
	Cheques      []*Cheque
	Orders       []*Order
	Categories   []*Category
	Policies     []*Policy
}
