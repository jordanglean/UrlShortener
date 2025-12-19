package handlers

import (
	"net/http"
	"time"

	"github.com/bytedance/gopkg/util/logger"

	"github.com/gin-gonic/gin"
	"github.com/jordanglean/UrlShortener/models"

	"github.com/jordanglean/UrlShortener/db"
)

func HandleURLRedirect(c *gin.Context) {
	shortCode := c.Param("id")

	var shortenUrlData models.ShortenURL
	result := db.DB.First(&shortenUrlData, "short_code = ?", shortCode)

	if result.Error != nil {
		logger.Debug("Error fetching shorten url", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error fetching shorten url",
			"error":   result.Error,
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, shortenUrlData.OriginalURL)
}

func HandleURLShorten(c *gin.Context) {
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

func HandleGetURLByUserID(c *gin.Context) {

	var query struct {
		UserID string `form:"userId" binding:"required"`
	}

	err := c.ShouldBindQuery(&query)

	if err != nil {
		logger.Error("Error binding query")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId query parameter is required",
		})
	}

	var urls []models.ShortenURL

	result := db.DB.Where("user_id = ?", query.UserID).Find(&urls)

	if result.Error != nil {
		logger.Debug("Can't find url for the user")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "error fetching urls",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"urls":  urls,
		"count": len(urls),
	})

}
