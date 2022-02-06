package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/my/profile", myProfileHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/query", queryHandler)
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

func booksHandler(c *gin.Context)  {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H {
		"id" : id,
	})
}

func queryHandler(c *gin.Context)  {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H {
		"title" : title,
	})
}