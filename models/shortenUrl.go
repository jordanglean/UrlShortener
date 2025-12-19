package models

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
	"time"
)

type ShortenURL struct {
	OriginalURL string    `json:"original_url" binding:"required"`
	ShortCode   string    `json:"short_code"`
	ShortURL    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	User        User      `json:"user" gorm:"foreignKey:UserID"  binding:"-"`
}

func GenerateURLCode(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return base64.RawURLEncoding.EncodeToString(bytes)[:length]
}
