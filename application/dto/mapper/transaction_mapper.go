package mapper

import (
	"strings"

	"com.github.gscampanario/mambu-assessment-test/application/dto/transaction"
	"com.github.gscampanario/mambu-assessment-test/domain/transaction/model"
)

func MapInsertTxnRequestToTransactionModel(dto transaction.InsertTransactionRequestDTO) model.Transaction {
	currency := strings.TrimSpace(dto.Currency)
	if currency == "" {
		currency = "EUR"
	}

	return model.Transaction{
		ID:           dto.IdempotencyUniqueKey,
		Amount:       dto.Amount,
		Currency:     currency,
		CreditorIBAN: dto.CreditorIBAN,
		CreditorName: dto.CreditorName,
		DebtorIBAN:   dto.DebtorIBAN,
		DebitorName:  dto.DebtorName,
	}
}
