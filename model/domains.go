package model

import (
	"gorm.io/gorm"
)

type Domains struct {
	gorm.Model
	Domain 		string
	Wildcard    bool
	ByUID		int
}
