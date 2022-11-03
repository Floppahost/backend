package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invites struct {
	gorm.Model
	UserID   int       `gorm:"primarykey"`
	Code     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UsedBy   string    `gorm:"default:null"`
	UsedByID int       `gorm:"default:null"`
}
