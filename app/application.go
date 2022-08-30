package app

import "github.com/gin-gonic/gin"

//the only points we interact with the HTTP server is in the application layer and the controllers

// only point in the entire API where we interact with the router
var (
	router = gin.Default()
)

func startApplication() {
	//defining our maps then running our router
	mapUrls()
	router.Run(":8080")
}
