package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Goodbye represents a message with a string, a user UUID, and a timestamp.
type Goodbye struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Message   string    `gorm:"type:text;not null"`
	UserUuid  uuid.UUID `gorm:"type:uuid;not null"`
	Timestamp time.Time `gorm:"type:timestamp;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
