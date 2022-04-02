package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		acceptContent := c.Request.Header.Get("Accept")
		if strings.Contains(acceptContent, "text/html") {
			c.String(http.StatusNotFound, "404页面")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    http.StatusNotFound,
				"error_message": "找不到url",
			})
		}
	})

	r.Run(":8080")
}
