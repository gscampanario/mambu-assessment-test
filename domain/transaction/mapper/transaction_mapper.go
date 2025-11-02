package mapper

import (
	"time"

	"com.github.gscampanario/mambu-assessment-test/domain/transaction/model"
)

func MapTransactionToDocument(transaction model.Transaction) model.Document {
	return model.Document{
		GrpHdr: model.GrpHdrType{
			MsgId:   transaction.ID,
			CreDtTm: time.Now(),
		},
		Amt: model.AmtType{
			Value: transaction.Amount,
			Ccy:   transaction.Currency,
		},
		Cdtr: model.CdtrType{
			Nm: transaction.CreditorName,
			CdtrAcct: model.CdtrAcctType{
				Id: model.IdType{
					IBAN: transaction.CreditorIBAN,
				},
			},
		},
		Dbtr: model.DbtrType{
			Nm: transaction.DebitorName,
			CdtrAcct: model.CdtrAcctType{
				Id: model.IdType{
					IBAN: transaction.DebtorIBAN,
				},
			},
		},
	}
}
