package model

import "gorm.io/gorm"

type Uploads struct {
	gorm.Model
	Path        string `gorm:"unique"`
	UploadUrl   string `gorm:"unique"`
	FileUrl     string `gorm:"unique"`
	UserID      int    `gorm:"primarykey"`
	UploadID    string `gorm:"unique"`
	SiteName    string
	SiteNameUrl string
	Title       string
	Description string
	Author      string
	AuthorUrl   string
	Color       string
	FileName    string
	MimeType    string
}
