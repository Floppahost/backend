package model

import "gorm.io/gorm"

type Uploads struct {
	gorm.Model
	Path         string `gorm:"unique"`
	UploadUrl    string `gorm:"unique"`
	FileUrl      string `gorm:"unique"`
	UserID       int    `gorm:"primarykey"`
	UploadID     string `gorm:"unique"`
	Name         string
	Title        string
	Description  string
	Author       string
	Color        string
	EmbedEnabled bool
	FileName     string
}
