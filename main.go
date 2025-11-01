package main

import (
	"context"

	"com.github.gscampanario/mambu-assessment-test/application/controller"
	"com.github.gscampanario/mambu-assessment-test/utils"
)

// main Starts the application
func main() {
	ctx := context.Background()

	logger := utils.GetLogger()
	// Ensure logger is flushed after ending application, leaving no buffered logs behind
	defer utils.FlushLogger()

	logger.Info("Application started")
	controller.Expose(ctx)
}
