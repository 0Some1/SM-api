package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name         string       `gorm:"not null" json:"name,omitempty"`
	StorePhone   string       `gorm:"not null;unique" json:"store_phone,omitempty"`
	Address      string       `gorm:"not null" json:"address,omitempty"`
	ActivityType string       `gorm:"not null" json:"activity_type,omitempty"`
	Necessities  []*Necessity `json:"necessities,omitempty"`
	Payments     []*Payment   `json:"payments,omitempty"`
	Bills        []*Bill      `json:"bills,omitempty"`
	Cheques      []*Cheque    `json:"cheques,omitempty"`
	Orders       []*Order     `json:"orders,omitempty"`
	Categories   []*Category  `json:"categories,omitempty"`
	Policies     []*Policy    `json:"policies,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
