package model

// Transaction represents a financial transaction between a debtor and a creditor
type Transaction struct {
	DebtorIBAN   string  `json:"debtor_iban"`
	DebitorName  string  `json:"debtor_name"`
	CreditorIBAN string  `json:"creditor_iban"`
	CreditorName string  `json:"creditor_name"`
	Amount       float64 `json:"ammount"`
	ID           string  `json:"idempotency_unique_key"`
	Currency     string  `json:"currency"`
	Status       string  `json:"status"`
}
