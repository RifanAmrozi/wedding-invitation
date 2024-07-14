package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	PhotoID   uuid.UUID `gorm:"type:uuid;not null"`
	GuestName string    `gorm:"not null"`
	Comment   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
}

func (c *Comment) BeforeCreate() error {
	c.ID = uuid.New()
	return nil
}
