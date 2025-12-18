package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

type ShortenURL struct {
	OriginalURL string    `json:"original_url" binding:"required"`
	ShortCode   string    `json:"short_code"`
	ShortURL    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      string    `json:"user_id" binding:"required"`
}

func (s *ShortenURL) Save() {
	fmt.Println("URL: ", s)
}

func GetAllShortenURL() {

}

func GenerateURLCode(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return base64.RawURLEncoding.EncodeToString(bytes)[:length]
}
