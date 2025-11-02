package client

import (
	"context"

	"com.github.gscampanario/mambu-assessment-test/infrastructure/db"
	"com.github.gscampanario/mambu-assessment-test/utils"
	"go.uber.org/zap"
)

type Client struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (c *Client) Insert(ctx context.Context) error {
	logger := utils.GetLogger()
	dbRepository := db.NewClientRepository()

	// Placeholder for insert logic
	if err := dbRepository.Insert(c); err != nil {
		logger.Error("Failed to insert new client", zap.Error(err))
	}

	logger.Info("Successfully inserted new client into database")
	return nil
}
