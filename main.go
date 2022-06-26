package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mcandeia/url-shortener/pkg/api"
	"github.com/mcandeia/url-shortener/pkg/shortener"
)

func main() {
	shortener.InitFactory()

	router := gin.Default()
	router.POST("/short", api.Shorten())
	router.GET("/:engine/:short", api.Redirect())

	err := router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	if err != nil {
		panic(err)
	}
}
