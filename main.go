package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanglean/UrlShortener/db"
	"github.com/jordanglean/UrlShortener/handlers"
)

func main() {

	// Init Database
	db.InitDB()

	router := gin.Default()

	{
		url := router.Group("/url")
		url.POST("/shorten", handlers.HandleURLShorten)
		url.GET("/:id", handlers.HandleURLRedirect)
	}
	{
		user := router.Group("/user")
		user.POST("/create", handlers.HandleCreateUser)
	}

	router.Run(":8080")
}
