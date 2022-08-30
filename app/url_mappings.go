package app

import "controllers"

//every GET request against "/ping" will be handled by Ping function
func mapUrls() {
	router.GET("/ping", controllers.Ping())
}
