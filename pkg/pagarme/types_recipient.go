package pagarme

import "time"

// CreateRecipient struct
type CreateRecipient struct {
	TransferEnable                string               `json:"transfer_enable,omitempty"`
	TransferInterval              string               `json:"transfer_interval,omitempty"`
	TransferDay                   string               `json:"transfer_day,omitempty"`
	AutomaticAnticipationEnabled  bool                 `json:"automatic_anticipation_enabled,omitempty"`
	AnticipatableVolumePercentage int                  `json:"anticipatable_volume_percentage,omitempty"`
	BankAccID                     string               `json:"bank_account_id,omitempty"`
	BankAcc                       BankAccountRecipient `json:"bank_account,omitempty"`
	PostbackURL                   string               `json:"postback_url,omitempty"`
	RegisterInfo                  RegisterInformation  `json:"register_information,omitempty,omitempty"`
	Metadata                      interface{}          `json:"metadata,omitempty"`
}

// Recipient structure to define recipient
type Recipient struct {
	Object                        string      `json:"object,omitempty"`
	ID                            string      `json:"id,omitempty"`
	TransferEnabled               bool        `json:"transfer_enabled,omitempty"`
	LastTransfer                  interface{} `json:"last_transfer,omitempty"`
	TransferInterval              string      `json:"transfer_interval,omitempty"`
	TransferDay                   int         `json:"transfer_day,omitempty"`
	AutomaticAnticipationEnabled  bool        `json:"automatic_anticipation_enabled,omitempty"`
	AnticipatableVolumePercentage int         `json:"anticipatable_volume_percentage,omitempty"`
	DateCreated                   time.Time   `json:"date_created,omitempty"`
	DateUpdated                   time.Time   `json:"date_updated,omitempty"`
	PostbackURL                   string      `json:"postback_url,omitempty"`
	Status                        string      `json:"status,omitempty"`
	StatusReason                  interface{} `json:"status_reason,omitempty"`
	Metadata                      interface{} `json:"metadata,omitempty"`
	Bank                          BankAccount `json:"bank_account,omitempty"`
}

// BankAccountRecipient structure
type BankAccountRecipient struct {
	Object             string      `json:"object,omitempty"`
	ID                 int         `json:"id,omitempty"`
	BankCode           string      `json:"bank_code,omitempty"`
	Agencia            string      `json:"agencia,omitempty"`
	AgenciaDv          interface{} `json:"agencia_dv,omitempty"`
	Conta              string      `json:"conta,omitempty"`
	ContaDv            string      `json:"conta_dv,omitempty"`
	Type               string      `json:"type,omitempty"`
	DocumentType       string      `json:"document_type,omitempty"`
	DocumentNumber     string      `json:"document_number,omitempty"`
	LegalName          string      `json:"legal_name,omitempty"`
	ChargeTransferFees bool        `json:"charge_transfer_fees,omitempty"`
	DateCreated        time.Time   `json:"date_created,omitempty"`
}

// RegisterInformation struct
type RegisterInformation struct {
	Name           string             `json:"name,omitempty"`
	Type           string             `json:"type,omitempty"`
	DocumentNumber string             `json:"document_number,omitempty"`
	CompanyName    string             `json:"company_name,omitempty"`
	Email          string             `json:"email,omitempty"`
	SiteURL        string             `json:"site_url,omitempty"`
	PhoneNum       []PhoneNumbers     `json:"phone_numbers,omitempty"`
	ManagingPart   []ManagingPartners `json:"managing_partners,omitempty"`
}

// PhoneNumbers struct
type PhoneNumbers struct {
	Ddd    string `json:"ddd,omitempty"`
	Number string `json:"number,omitempty"`
	Type   string `json:"type,omitempty"`
}

// ManagingPartners struct
type ManagingPartners struct {
	Type           string `json:"type,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
	Email          string `json:"email,omitempty"`
	Name           string `json:"name,omitempty"`
}
