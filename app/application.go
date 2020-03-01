package app

import (
	"github.com/gin-gonic/gin"
	"github.com/meetmanok/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication(){
	mapsUrl()
	logger.Info("about to start the application...")
	router.Run()

}
