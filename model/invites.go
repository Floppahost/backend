package model

import "gorm.io/gorm"

type Invites struct {
	gorm.Model
	UserID int    `gorm:"primarykey"`
	Code   string `gorm:"unique"`
	UsedBy bool
}
