package db

import (
	"sync"

	"com.github.gscampanario/mambu-assessment-test/utils"
)

// IClientRepository defines the interface for client database operations
type IClientRepository interface {
	Insert(client interface{}) error
}

type ClientRepository struct{}

var (
	once sync.Once
	inst *ClientRepository
)

// NewClientRepository returns the same instance of ClientRepository per application
func NewClientRepository() *ClientRepository {
	if inst == nil {
		createClientRepository()
	}
	return inst
}

// createClientRepository initializes the singleton instance of ClientRepository
func createClientRepository() {
	once.Do(func() {
		inst = &ClientRepository{}
	})
}

// Insert inserts a new client into the database
func (r *ClientRepository) Insert(client interface{}) error {
	// Placeholder for actual database insertion logic
	logger := utils.GetLogger()
	logger.Info("Inserting client into the database")
	return nil
}
