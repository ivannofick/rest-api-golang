package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)

	router.GET("/my/profile", myProfileHandler)

	router.Run()
}

func rootHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H {
		"name" : "ivannofick",
		"bio" : "A Profesional Developer",
	})
}

func myProfileHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H {
		"name" : "ivannofick",
		"bio" : "A Profesional Developer",
		"bod" : "18-xx-xxxx",
	})
}