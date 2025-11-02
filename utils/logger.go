package utils

import (
	"sync"

	"com.github.gscampanario/mambu-assessment-test/config"
	"go.uber.org/zap"
)

var (
	once     sync.Once
	instance *zap.Logger
)

// GetLogger returns the same instance of a configured logger per application,
// creating a new one if not existent.
func GetLogger() *zap.Logger {
	once.Do(func() {
		cfg := config.Get()

		var err error
		if cfg.Env != "prod" {
			instance, err = zap.NewDevelopment()
		} else {
			instance, err = zap.NewProduction()
		}
		if err != nil {
			// returns a no-operation logger if production logger fails, avoiding unexpected panics
			instance = zap.NewNop()
		}
	})
	return instance
}

// FlushLogger flushes any pending log entries in memory.
// Always call it by defer to properly clean leftovers.
func FlushLogger() {
	if instance == nil {
		return
	}
	_ = instance.Sync()
}
