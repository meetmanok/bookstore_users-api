package app

import (
	"github.com/meetmanok/bookstore_users-api/controllers/ping"
	"github.com/meetmanok/bookstore_users-api/controllers/users"
)

func mapsUrl(){
	// Ping controller routes
	router.GET("/ping", ping.Ping)

	// Users controller routes
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/internal/users/search", users.Search)
}