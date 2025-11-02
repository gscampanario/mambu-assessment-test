package model

import (
	"time"
)

// Document represents the root XML structure for a financial transaction
type Document struct {
	GrpHdr GrpHdrType `xml:"GrpHdr"`
	Cdtr   CdtrType   `xml:"Cdtr"`
	Dbtr   DbtrType   `xml:"Dbtr"`
	Amt    AmtType    `xml:"Amt"`
}

// GrpHdrType represents the group header information in the XML document
type GrpHdrType struct {
	MsgId   string    `xml:"MsgId"`
	CreDtTm time.Time `xml:"CreDtTm"`
}

// IdType represents an identification type, such as an IBAN
type IdType struct {
	IBAN string `xml:"IBAN"`
}

// CdtrAcctType represents the creditor account information
type CdtrAcctType struct {
	Id IdType `xml:"Id"`
}

// CdtrType represents the creditor information
type CdtrType struct {
	Nm       string       `xml:"Nm"`
	CdtrAcct CdtrAcctType `xml:"CdtrAcct"`
}

// DbtrType represents the debtor information
type DbtrType struct {
	Nm       string       `xml:"Nm"`
	CdtrAcct CdtrAcctType `xml:"CdtrAcct"`
}

// AmtType represents the amount information with currency
type AmtType struct {
	Value float64 `xml:",chardata"`
	Ccy   string  `xml:"Ccy,attr"`
}
