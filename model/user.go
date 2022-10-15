package model

import "gorm.io/gorm"

type Users struct {
	User string `gorm:"index:idx_user,unique"`
	Password   string
	gorm.Model
}

