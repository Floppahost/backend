package model

import "gorm.io/gorm"

type Uploads struct {
	gorm.Model
	UserID      int `gorm:"primarykey"`
	Embed       bool
	Title       string
	Description string
	Author      string
	Color       string
}
