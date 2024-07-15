package models

import (
	"time"

	"gorm.io/gorm"
)

type Gameplay struct {
	gorm.Model
	Score  int  `json:"score"`
	UserID uint `json:"user_id"`
	User   User `json:"-"`
}

type Score struct {
	gorm.Model
	Score  int        `json:"score"`
	Month  *time.Time `json:"month"`
	UserID uint       `json:"user_id"`
	User   User       `json:"-"`
}
