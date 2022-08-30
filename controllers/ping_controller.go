package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//function that signals to the cloud that the website is running
//when you receive a get request against a localhost, you return a json 200 OK, saying "pong"

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
