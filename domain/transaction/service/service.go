package service

import (
	"encoding/xml"
	"fmt"
	"os"

	"com.github.gscampanario/mambu-assessment-test/domain/transaction/mapper"
	"com.github.gscampanario/mambu-assessment-test/domain/transaction/model"
	"com.github.gscampanario/mambu-assessment-test/infrastructure/db"
	"com.github.gscampanario/mambu-assessment-test/utils"
	"go.uber.org/zap"
)

// ITransactionService defines the interface for transaction service operations
type ITransactionService interface {
	Insert(transaction model.Transaction) error
}

// TransactionService provides methods to manage transactions
type TransactionService struct {
	TransactionRepository db.ITransactionRepository
	BankFolder            string
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(repository db.ITransactionRepository, bankFolder string) *TransactionService {
	return &TransactionService{
		TransactionRepository: repository,
		BankFolder:            bankFolder,
	}
}

// Insert inserts a new transaction and generates an XML document
func (s *TransactionService) Insert(transaction model.Transaction) error {
	if err := s.TransactionRepository.Insert(transaction.ID, "PENDING"); err != nil {
		return err
	}

	xmlDoc := mapper.MapTransactionToDocument(transaction)
	return s.generateXMLDocument(xmlDoc)
}

// generateXMLDocument creates an XML document from the provided Document model
func (s *TransactionService) generateXMLDocument(document model.Document) error {
	logger := utils.GetLogger()

	xmlData, err := xml.MarshalIndent(document, "", "  ")
	if err != nil {
		logger.Error("[TransactionService.generateXMLDocument] Failed to marshal XML document", zap.Error(err))
		return err
	}

	filename := fmt.Sprintf("%s_documents.xml", document.GrpHdr.MsgId)
	file, err := os.Create(s.BankFolder + "/" + filename)
	if err != nil {
		logger.Error("[TransactionService.generateXMLDocument] Failed to create XML document", zap.Error(err))
		return err
	}
	defer file.Close()

	file.WriteString(xml.Header)
	file.Write(xmlData)

	logger.Info("[TransactionService.generateXMLDocument] Created XML document")

	return nil
}
