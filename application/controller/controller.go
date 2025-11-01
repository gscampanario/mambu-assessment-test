package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// expose Inits all endpoints for the controller package
func Expose(ctx context.Context) {
	r := gin.Default()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "It works!"})
	})

	// TODO: Get logger level and server port from config file using viper.
	if err := r.Run(":8080"); err != nil {
		logger.Error("[controller.expose] Failed to run server", zap.Error(err))
		return
	}
}
