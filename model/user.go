package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	User      string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"unique;not null"`
	ApiKey    string `gorm:"not null;unique"`
	Token     string `gorm:"not null"` // TODO: define unique
	InviteBan bool   `gorm:"default:null"`
	Admin     bool
}
