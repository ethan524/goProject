package goProject

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", Hello)

	_ = r.Run(":8000")
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  "jenkins",
		"title": "CICD",
	})
}
