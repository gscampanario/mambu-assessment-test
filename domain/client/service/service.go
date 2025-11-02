package service

import (
	"context"

	"com.github.gscampanario/mambu-assessment-test/domain/client/model"
	"com.github.gscampanario/mambu-assessment-test/infrastructure/db"
	"com.github.gscampanario/mambu-assessment-test/utils"
	"go.uber.org/zap"
)

type IClientService interface {
	Insert(ctx context.Context, client model.Client) error
}

type ClientService struct {
	ClientRepository db.IClientRepository
}

// NewClientService creates a new instance of ClientService
func NewClientService() *ClientService {
	return &ClientService{
		ClientRepository: db.NewClientRepository(),
	}
}

func (s *ClientService) Insert(ctx context.Context, client model.Client) error {
	logger := utils.GetLogger()

	if err := s.ClientRepository.Insert(ctx, client); err != nil {
		logger.Error("Failed to insert new client", zap.Error(err))
		return err
	}

	logger.Info("Successfully inserted new client into database")
	return nil
}
