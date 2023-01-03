package model

import (
	"gorm.io/gorm"
)

type Embeds struct {
	gorm.Model
	UserID   int       `gorm:"primarykey"`
	Name 		string
	Title       string
	Description string
	Author      string
	Color       string
	Enabled 	bool
}

type EmbedStruct struct {
	Title 	string
	Author 	string
	Description string
	Name string
	Enabled bool
}