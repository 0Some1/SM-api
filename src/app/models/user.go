package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string    `gorm:"not null" json:"name,omitempty"`
	Username    string    `gorm:"not null;unique" json:"username,omitempty"`
	PhoneNumber string    `gorm:"not null;unique" json:"phone_number,omitempty"`
	Email       string    `gorm:"not null;unique" json:"email,omitempty"`
	Password    string    `gorm:"not null" json:"password,omitempty"`
	Policies    []*Policy `json:"policies,omitempty"`
}
