package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "Implement me!",
		})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"status": "Implement me!",
	})
}