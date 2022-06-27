package main

import (
	dapr "github.com/dapr/go-sdk/client"
	"github.com/gin-gonic/gin"
	"github.com/mcandeia/url-shortener/pkg/api"
	"github.com/mcandeia/url-shortener/pkg/shortener"
	"github.com/mcandeia/url-shortener/pkg/state"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	kvState := state.NewKV(client)

	shortener.InitFactory(kvState)

	router := gin.Default()
	router.POST("/short", api.Shorten())
	router.GET("/:engine/:short", api.Redirect())

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	if err := router.Run(); err != nil {
		panic(err)
	}
}
