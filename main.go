package main

import (
	"net/http"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"github.com/jordanglean/UrlShortener/db"
	"github.com/jordanglean/UrlShortener/models"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

var urlStore = make(map[string]string)

func main() {

	// Init Database
	db.InitDB()

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
	var shortenUrl models.ShortenURL

	err := c.ShouldBindJSON(&shortenUrl)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding request to json",
			"error":   err,
		})
		return
	}

	shortCode := models.GenerateURLCode(6)

	shortenUrl.ShortCode = shortCode
	shortenUrl.CreatedAt = time.Now()
	shortenUrl.ShortURL = "http://localhost:8080/url/" + shortCode

	db.DB.Create(&shortenUrl)

	c.JSON(http.StatusCreated, shortenUrl)
}
