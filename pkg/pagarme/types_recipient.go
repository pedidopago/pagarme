package pagarme

import "time"

// CreateRecipient struct
type CreateRecipient struct {
	TransferEnable                string
	TransferInterval              string
	TransferDay                   string
	AutomaticAnticipationEnabled  bool `json:"automatic_anticipation_enabled"`
	AnticipatableVolumePercentage int  `json:"anticipatable_volume_percentage"`
	BankAccountID                 string
	BankAcc                       BankAccount
	PostbackURL                   string
	RegisterInfo                  RegisterInformation
	Metadata                      interface{}
}

// Recipient structure to define recipient
type Recipient struct {
	Object                        string      `json:"object"`
	ID                            string      `json:"id"`
	TransferEnabled               bool        `json:"transfer_enabled"`
	LastTransfer                  interface{} `json:"last_transfer"`
	TransferInterval              string      `json:"transfer_interval"`
	TransferDay                   int         `json:"transfer_day"`
	AutomaticAnticipationEnabled  bool        `json:"automatic_anticipation_enabled"`
	AnticipatableVolumePercentage int         `json:"anticipatable_volume_percentage"`
	DateCreated                   time.Time   `json:"date_created"`
	DateUpdated                   time.Time   `json:"date_updated"`
	PostbackURL                   string      `json:"postback_url"`
	Status                        string      `json:"status"`
	StatusReason                  interface{} `json:"status_reason"`
	Metadata                      interface{} `json:"metadata"`
	Bank                          BankAccount `json:"bank_account"`
}

// BankAccountRecipient structure
type BankAccountRecipient struct {
	Object             string      `json:"object"`
	ID                 int         `json:"id"`
	BankCode           string      `json:"bank_code"`
	Agencia            string      `json:"agencia"`
	AgenciaDv          interface{} `json:"agencia_dv"`
	Conta              string      `json:"conta"`
	ContaDv            string      `json:"conta_dv"`
	Type               string      `json:"type"`
	DocumentType       string      `json:"document_type"`
	DocumentNumber     string      `json:"document_number"`
	LegalName          string      `json:"legal_name"`
	ChargeTransferFees bool        `json:"charge_transfer_fees"`
	DateCreated        time.Time   `json:"date_created"`
}

// RegisterInformation struct
type RegisterInformation struct {
	Type           string `json:"type"`
	DocumentNumber string `json:"document_number"`
	CompanyName    string `json:"company_name"`
	Email          string `json:"email"`
	SiteURL        string `json:"site_url"`
	PhoneNumbers   []struct {
		Ddd    string `json:"ddd"`
		Number string `json:"number"`
		Type   string `json:"type"`
	} `json:"phone_numbers"`
	ManagingPartners []struct {
		Type           string `json:"type"`
		DocumentNumber string `json:"document_number"`
		Email          string `json:"email"`
		Name           string `json:"name"`
	} `json:"managing_partners"`
}
