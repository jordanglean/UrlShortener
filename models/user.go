package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID    `json:"id"`
	Username   string       `json:"username" binding:"required"`
	Email      string       `json:"email" binding:"required"`
	CreatedAt  time.Time    `json:"created_at"`
	ShortenURL []ShortenURL `json:"shorten_urls,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
