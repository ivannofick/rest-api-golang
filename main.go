package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"name" : "ivannofick",
			"bio" : "A Profesional Developer",
		})
	})

	router.GET("/my/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"name" : "ivannofick",
			"bio" : "A Profesional Developer",
			"bod" : "18-xx-xxxx",
		})
	})

	router.Run(":8100")
}