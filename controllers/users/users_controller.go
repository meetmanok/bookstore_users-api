package users

import (
	"github.com/gin-gonic/gin"
	"github.com/meetmanok/bookstore_users-api/domain/users"
	"github.com/meetmanok/bookstore_users-api/services"
	"github.com/meetmanok/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	//Takes data from request
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	//Send it to service
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//return any error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	//Return data
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"status": "Implement me!",
	})
}
