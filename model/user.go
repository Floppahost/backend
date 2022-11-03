package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"unique;not null"`
	ApiKey    string `gorm:"not null;unique"`
	Token     string `gorm:"not null;unique"`
	InviteBan bool   `gorm:"default:null"`
	Admin     bool
}
