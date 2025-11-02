package controller

import (
	"context"
	"fmt"
	"net/http"

	"com.github.gscampanario/mambu-assessment-test/application/controller/middleware"
	"com.github.gscampanario/mambu-assessment-test/application/dto/mapper"
	"com.github.gscampanario/mambu-assessment-test/application/dto/transaction"
	"com.github.gscampanario/mambu-assessment-test/config"
	txnService "com.github.gscampanario/mambu-assessment-test/domain/transaction/service"
	"com.github.gscampanario/mambu-assessment-test/infrastructure/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Expose Inits all endpoints for the controller package
func Expose(ctx context.Context) {
	cfg := config.Get()

	r := gin.Default()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	transactionRepository, err := db.NewTransactionRepository(cfg.Storage.DBFileLocation)
	if err != nil {
		logger.Error("[controller.Expose] Failed to create transaction repository", zap.Error(err))
		panic(err)
	}
	defer transactionRepository.DB.Close()

	transactionService := txnService.NewTransactionService(transactionRepository, cfg.Storage.BankFolder)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "It works!"})
	})

	authorized := r.Group("/transaction", middleware.BasicAuthMiddleware(cfg.Server.Auth.Username, cfg.Server.Auth.Password))
	authorized.POST("/", func(c *gin.Context) {
		var dto transaction.InsertTransactionRequestDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tx := mapper.MapInsertTxnRequestToTransactionModel(dto)
		if err := transactionService.Insert(tx); err != nil {
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
