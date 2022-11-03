package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	User     string `gorm:"unique,not null"`
	Email    string `gorm:"unique,not null"`
	Password string `gorm:"unique,not null"`
	Token    string `gorm:"unique,not null"`
}
