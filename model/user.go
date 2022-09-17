package model

import "gorm.io/gorm"

type BaseUser struct {
	Usuario string `gorm:"index:idx_user,unique"`
	Senha   string
	gorm.Model
}

