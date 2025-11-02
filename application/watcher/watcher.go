package watcher

import (
	"encoding/csv"
	"os"

	"com.github.gscampanario/mambu-assessment-test/config"
	"com.github.gscampanario/mambu-assessment-test/infrastructure/db"
	"com.github.gscampanario/mambu-assessment-test/utils"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

func WatchBankTxnFeedback() {
	logger := utils.GetLogger()
	cfg := config.Get()
	transactionRepository, err := db.NewTransactionRepository(cfg.Storage.DBFileLocation)
	if err != nil {
		logger.Error("[watcher.WatchBankTxnFeedback] Failed to create transaction repository", zap.Error(err))
		return
	}
	logger.Info("[watcher.WatchBankTxnFeedback] Start watching")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Error("[watcher.WatchBankTxnFeedback] Failed to create file watcher", zap.Error(err))
		return
	}
	defer watcher.Close()

	if err := watcher.Add(cfg.Storage.TxnFeedbackFolder); err != nil {
		logger.Error("[watcher.WatchBankTxnFeedback] Failed to add folder to watcher", zap.Error(err))
		return
	}

	done := make(chan bool)

	logger.Info("[watcher.WatchBankTxnFeedback] Start go routine")

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				logger.Info("[watcher.WatchBankTxnFeedback] File event detected", zap.String("event", event.String()))
				// Bitwise check for create event, recommended by fsnotify documentation
				if event.Op&fsnotify.Create == fsnotify.Create {
					logger.Info("[watcher.WatchBankTxnFeedback] File event detected", zap.String("file", event.Name))

					// Process the feedback file
					file, err := os.Open(event.Name)
					if err != nil {
						logger.Error("[watcher.WatchBankTxnFeedback] Failed to open feedback file", zap.Error(err))
						continue
					}
					defer file.Close()

					reader := csv.NewReader(file)
					records, err := reader.ReadAll()
					if err != nil {
						logger.Error("[watcher.WatchBankTxnFeedback] Failed to read feedback file", zap.Error(err))
						continue
					}

					for i, row := range records {
						if i == 0 {
							// Skip header row
							continue
						}
						id := row[0]
						status := row[1]

						txn, err := transactionRepository.GetById(id)
						if err != nil {
							logger.Error("[watcher.WatchBankTxnFeedback] Failed to get transaction by ID", zap.Error(err))
							continue
						}
						if txn.ID == "" {
							logger.Warn("[watcher.WatchBankTxnFeedback] Transaction not found", zap.String("id", id))
							continue
						}

						if err := transactionRepository.Update(id, status); err != nil {
							logger.Error("[watcher.WatchBankTxnFeedback] Failed to update transaction status", zap.Error(err))
							continue
						}
						logger.Info("[watcher.WatchBankTxnFeedback] Updated transaction status", zap.String("id", id), zap.String("status", status))
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Error("[watcher.WatchBankTxnFeedback] Watcher error", zap.Error(err))
			}
		}
	}()
	<-done // Blocks the goroutine while the app is running
}
