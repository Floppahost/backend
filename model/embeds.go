package model

import (
	"gorm.io/gorm"
)

type Embeds struct {
	gorm.Model
	UserID      int    `gorm:"unique;primarykey"`
	SiteName    string `gorm:"default:Floppa!"`
	SiteNameUrl string `gorm:"default:https://floppa.host"`
	Title       string `gorm:"default:I am using Floppa.host!"`
	Description string `gorm:"default:Floppa.host is good!"`
	Author      string `gorm:"default:Floppa"`
	AuthorUrl   string `gorm:"default:https://floppa.host"`
	Color       string `gorm:"default:random"`
	Domain      string `gorm:"default:floppa.host"`
	Path        string
	Path_Mode   string `gorm:"not null;default:invisible"`
	Path_Amount int    `gorm:"not null;default:5"`
}

type EmbedStruct struct {
	SiteName    string
	SiteNameUrl string
	Title       string
	Description string
	Author      string
	AuthorUrl   string
	Color       string
}
