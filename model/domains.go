package model

import (
	"gorm.io/gorm"
)

type Domains struct {
	gorm.Model
	Domain   string `gorm:"unique;not null"`
	Wildcard bool
	ByUID    int `gorm:"unique;not null"`
}
