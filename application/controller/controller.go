package controller

import (
	"context"
	"fmt"
	"net/http"

	"com.github.gscampanario/mambu-assessment-test/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Expose Inits all endpoints for the controller package
func Expose(ctx context.Context) {
	cfg := config.Get()

	r := gin.Default()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "It works!"})
	})

	// TODO: Get logger level and server port from config file using viper.
	if err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
		logger.Error("[controller.expose] Failed to run server", zap.Error(err))
		return
	}
}
