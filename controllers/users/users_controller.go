package users

import (
	"github.com/gin-gonic/gin"
	"github.com/meetmanok/bookstore_users-api/domain/users"
	"github.com/meetmanok/bookstore_users-api/services"
	"github.com/meetmanok/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func getUserId(userId string) (int64, *errors.RestErr){
	userID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		err := errors.NewBadRequestErr("user id should be a number")
		return 0, err
	}

	return userID, nil
}

func GetUser(c *gin.Context) {
	userID, err := getUserId(c.Param("user_id"))

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	result, geterr := services.UsersService.GetUser(userID)

	if geterr != nil {
		c.JSON(geterr.Status, geterr)
		return
	}

	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func CreateUser(c *gin.Context) {
	var user users.User

	//Takes data from request
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	//Send it to service
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		//return any error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	//Return data
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	var user users.User

	//Takes data from request
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestErr("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user.Id = userID
	isPartial := c.Request.Method == http.MethodPatch

	//Send it to service
	result, saveErr := services.UsersService.UpdateUser(isPartial, user)
	if saveErr != nil {
		//return any error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	//Return data
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}


func DeleteUser(c *gin.Context) {
	userID, err := getUserId(c.Param("user_id"))

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}


func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}
