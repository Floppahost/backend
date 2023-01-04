package model

import "gorm.io/gorm"

type Uploads struct {
	gorm.Model
	UploadID 	string
	Object 		string 
	UserID      int `gorm:"primarykey"`
	Name 		string
	Title       string
	Description string
	Author      string
	Color       string
	EmbedEnabled bool
	Domain		string
	FileName 	string
}
