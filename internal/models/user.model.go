package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	// ID     uint   `gorm:"primary_key"`
	Name           string `gorm:"type:varchar(255);not null"`
	PersonalNumber string `gorm:"uniqueIndex;not null"`
	ConfigJSON     string `gorm:"type:varchar(255)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
