package entity

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	PhotoURL  string    `gorm:"not null"`
	CreatedAt time.Time
}

func (p *Photo) BeforeCreate() error {
	p.ID = uuid.New()
	return nil
}
