package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Greet function to called
func Greet(c *gin.Context) {
	c.String(http.StatusOK, "Hello Gophers")
}
