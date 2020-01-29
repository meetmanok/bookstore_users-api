package users

import (
	"github.com/gin-gonic/gin"
	"github.com/meetmanok/bookstore_users-api/domain/users"
	"github.com/meetmanok/bookstore_users-api/services"
	"github.com/meetmanok/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
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
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		err := errors.NewBadRequestErr("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	result, geterr := services.GetUser(userID)

	if geterr != nil {
		c.JSON(geterr.Status, geterr)
		return
	}

	c.JSON(http.StatusOK, result)
}
