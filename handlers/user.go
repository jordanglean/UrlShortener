package handlers

import (
	"net/http"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"github.com/jordanglean/UrlShortener/models"

	"github.com/jordanglean/UrlShortener/db"
)

func HandleCreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil {
		logger.Error("Error binding request to user", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user.CreatedAt = time.Now()

	db.DB.Create(&user)

	c.JSON(http.StatusCreated, user)

}
