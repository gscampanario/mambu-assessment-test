package transaction

type InsertTransactionRequestDTO struct {
	DebtorIBAN           string  `json:"debtor_iban" binding:"required" validate:"required,matches(^[A-Z]{2}[0-9]{2}[A-Za-z0-9]{1,30}$)"`
	DebtorName           string  `json:"debtor_name" binding:"required" validate:"required,min=3,max=30"`
	CreditorIBAN         string  `json:"creditor_iban" binding:"required" validate:"required,matches(^[A-Z]{2}[0-9]{2}[A-Za-z0-9]{1,30}$)"`
	CreditorName         string  `json:"creditor_name" binding:"required" validate:"required,min=3,max=30"`
	Amount               float64 `json:"ammount" binding:"required" validate:"required,gt=0"`
	IdempotencyUniqueKey string  `json:"idempotency_unique_key" binding:"required" validate:"required,len=10"`
	Currency             string  `json:"currency"`
}
