package app

import (
	"github.com/RBrothersBSI/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)

func StartApplication(){
	mapUrls()
	logger.Info("Starting application...")
	router.Run(":8100")
}