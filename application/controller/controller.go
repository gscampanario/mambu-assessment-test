package controller

import (
	"context"
	"fmt"
	"net/http"

	"com.github.gscampanario/mambu-assessment-test/application/dto/client"
	"com.github.gscampanario/mambu-assessment-test/application/dto/mapper"
	"com.github.gscampanario/mambu-assessment-test/config"
	"com.github.gscampanario/mambu-assessment-test/domain/client/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Expose Inits all endpoints for the controller package
func Expose(ctx context.Context) {
	cfg := config.Get()

	r := gin.Default()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	clientMapper := mapper.NewClientMapper()
	clientService := service.NewClientService()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "It works!"})
	})

	r.POST("/client", func(c *gin.Context) {
		var dto client.InsertClientRequestDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ct := clientMapper.MapInsertClientRequestDTOToClient(dto)
		if err := clientService.Insert(c.Request.Context(), ct); err != nil {
			// TODO: improve error handling with custom error types
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusCreated, gin.H{"message": "ok"})
	})

	if err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
		logger.Error("[controller.expose] Failed to run server", zap.Error(err))
		return
	}
}
