package models

import (
	"gorm.io/gorm"
	"time"
)

type Cheque struct {
	gorm.Model
	Date        time.Time `gorm:"not null"`
	BankName    string    `gorm:"not null"`
	Payee       string    `gorm:"not null"`
	Amount      int64     `gorm:"not null"`
	StoreID     uint      `gorm:"not null"`
	Description string
}
