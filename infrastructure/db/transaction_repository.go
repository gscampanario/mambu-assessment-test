package db

import (
	"database/sql"

	"com.github.gscampanario/mambu-assessment-test/domain/transaction/model"
	_ "github.com/mattn/go-sqlite3"

	"com.github.gscampanario/mambu-assessment-test/infrastructure/db/schemas"
	"com.github.gscampanario/mambu-assessment-test/utils"
	"go.uber.org/zap"
)

// ITransactionRepository defines the interface for transaction database operations
type ITransactionRepository interface {
	Insert(id, status string) error
	GetById(id string) (model.Transaction, error)
	Update(id string, status string) error
}

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(dbFilePath string) (*TransactionRepository, error) {
	logger := utils.GetLogger()
	db, err := connect(dbFilePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schemas.CreateTableStmt)
	if err != nil {
		logger.Error("[TransactionRepository.NewTransactionRepository] Failed to load schemas into db", zap.Error(err))
		return nil, err
	}

	return &TransactionRepository{DB: db}, nil
}

func connect(dbFilePath string) (*sql.DB, error) {
	logger := utils.GetLogger()

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		logger.Error("[TransactionRepository.connect] Failed to connect to database", zap.Error(err))
		return nil, err
	}

	logger.Info("[TransactionRepository.connect] Connected to database")

	return db, nil
}

func (r *TransactionRepository) Insert(id, status string) error {
	logger := utils.GetLogger()

	_, err := r.DB.Exec(schemas.InsertTransactionStmt, id, status)
	if err != nil {
		logger.Error("[TransactionRepository.Insert] Failed to insert transaction", zap.Error(err))
	}

	return err
}

func (r *TransactionRepository) GetById(id string) (model.Transaction, error) {
	logger := utils.GetLogger()

	rows, err := r.DB.Query(schemas.GetTransactionByIDStmt, id)
	if err != nil {
		logger.Error("[TransactionRepository.GetById] Failed to query transaction by ID", zap.Error(err))
		return model.Transaction{}, err
	}
	defer rows.Close()

	txn := model.Transaction{}
	for rows.Next() {
		var id, status string
		if err := rows.Scan(&id, &status); err != nil {
			logger.Error("[TransactionRepository.GetById] Failed to scan transaction row", zap.Error(err))
			return model.Transaction{}, err
		}
		txn.ID = id
		txn.Status = status
	}

	return txn, nil
}

func (r *TransactionRepository) Update(id string, status string) error {
	logger := utils.GetLogger()

	_, err := r.DB.Exec(schemas.UpdateTransactionStatusByIdStmt, status, id)
	if err != nil {
		logger.Error("[TransactionRepository.Update] Failed to update transaction status", zap.Error(err))
	}

	return err
}
