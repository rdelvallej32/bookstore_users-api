package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rdelvallej32/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	setRoutes()
	logger.Info("Starting application...")
	router.Run(":8080")
}
