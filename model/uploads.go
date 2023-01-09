package model

import "gorm.io/gorm"

type Uploads struct {
	gorm.Model
	UploadUrl string   `gorm:"unique"`
	FileUrl string	   `gorm:"unique"`
	UploadID 	string `gorm:"unique"`
	Object 		string  `gorm:"unique"`
	UserID      int 	`gorm:"primarykey"`
	Name 		string
	Title       string
	Description string
	Author      string
	Color       string
	EmbedEnabled bool
	FileName 	string 
	Path		string `gorm:"unique;not null"`
}
