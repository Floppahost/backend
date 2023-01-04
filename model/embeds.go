package model

import (
	"gorm.io/gorm"
)

type Embeds struct {
	gorm.Model
	UserID   	int    `gorm:"primarykey"`
	Name 		string `gorm:"default:Floppa!"`
	Title       string `gorm:"default:I am using Floppa.host!"`
	Description string `gorm:"default:Floppa.host is good!"`
	Author      string `gorm:"default:Floppa"`
	Color       string `gorm:"default:random"`
	Enabled 	bool   `gorm:"default:true"`
	Domain 	  	string `gorm:"default:floppa.host"`
}

type EmbedStruct struct {
	Title 	string
	Author 	string
	Description string
	Name string
	Enabled bool
}