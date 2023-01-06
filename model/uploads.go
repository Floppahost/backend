package model

import "gorm.io/gorm"

type Uploads struct {
	gorm.Model
	UploadUrl string
	FileUrl string
	UploadID 	string
	Object 		string 
	UserID      int `gorm:"primarykey"`
	Name 		string
	Title       string
	Description string
	Author      string
	Color       string
	EmbedEnabled bool
	FileName 	string
}
