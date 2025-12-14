package main

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

type ShortenResponse struct {
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:short_url`
}

var urlStore = make(map[string]string)

func main() {
	router := gin.Default()

	{
		url := router.Group("/url")
		url.POST("/shorten", handleShorten)
		url.GET("/:id", handleRedirect)
	}

	router.Run(":8080")
}

func handleRedirect(c *gin.Context) {
	code := c.Param("id")

	if url, exists := urlStore[code]; exists {
		c.Redirect(302, url)
		return
	}

	c.JSON(404, gin.H{
		"error": "not found",
	})
}

func handleShorten(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "url is required",
		})
		return
	}

	code := generateURLCode(6)
	urlStore[code] = req.URL

	c.JSON(200, ShortenResponse{
		OriginalURL: req.URL,
		ShortCode:   code,
		ShortURL:    "http://localhost:8080/url/" + code,
	})
}

func generateURLCode(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return base64.RawURLEncoding.EncodeToString(bytes)[:length]
}
