package api

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

//StartApplication to start the application
func StartApplication() {
	mapUrls()
	router.Run(":9999")
}
