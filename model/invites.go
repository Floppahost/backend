package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invites struct {
	gorm.Model
	UserID   int       `gorm:"primarykey"`
	Code     uuid.UUID `gorm:"unique;type:uuid;default:uuid_generate_v4()"`
	UsedByID int       `gorm:"default:null"`
}
