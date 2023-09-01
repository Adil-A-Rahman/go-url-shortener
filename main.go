package main

import (
	"fmt"

	"github.com/Adil-A-Rahman/go-url-shortener/handler"
	"github.com/Adil-A-Rahman/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", handler.CreateShortUrl)

	r.GET("/:shortUrl", handler.HandleShortUrlRedirect)

	// Note: store initialization happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
