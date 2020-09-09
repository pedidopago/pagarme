package pagarme

import (
	"time"
)

type Transaction struct {
	Amount               int                    `json:"amount" form:"amount"`
	CardHash             string                 `json:"card_hash,omitempty" form:"card_hash,omitempty"`
	CardID               string                 `json:"card_id,omitempty" form:"card_id,omitempty"`
	CardHolderName       string                 `json:"card_holder_name,omitempty" form:"card_holder_name,omitempty"`
	CardExpirationDate   MMYY                   `json:"card_expiration_date,omitempty" form:"card_expiration_date,omitempty"`
	CardNumber           string                 `json:"card_number,omitempty" form:"card_number,omitempty"`
	CardCVV              string                 `json:"card_cvv,omitempty" form:"card_cvv,omitempty"`
	PaymentMethod        PaymentMethod          `json:"payment_method,omitempty" form:"payment_method,omitempty"`
	PostbackURL          string                 `json:"postback_url,omitempty" form:"postback_url,omitempty"`
	Async                bool                   `json:"async,omitempty" form:"async,omitempty"`
	Installments         int                    `json:"installments,omitempty" form:"installments,omitempty"`
	BoletoExpirationDate string                 `json:"boleto_expiration_date,omitempty" form:"boleto_expiration_date,omitempty"`
	SoftDescriptor       string                 `json:"soft_descriptor,omitempty" form:"soft_descriptor,omitempty"`
	Capture              string                 `json:"capture,omitempty" form:"capture,omitempty"`
	BoletoInstructions   string                 `json:"boleto_instructions,omitempty" form:"boleto_instructions,omitempty"`
	SplitRules           []*SplitRule           `json:"split_rules,omitempty" form:"split_rules,omitempty"`
	Customer             *Customer              `json:"customer" form:"customer"`
	Billing              *Billing               `json:"billing" form:"billing"`
	Shipping             *Shipping              `json:"shipping" form:"shipping"`
	Items                []*Item                `json:"items,omitempty" form:"items,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty" form:"metadata,omitempty"`
	Session              string                 `json:"session,omitempty" form:"session,omitempty"`
	LocalTime            *time.Time             `json:"local_time,omitempty" form:"local_time,omitempty"`
	//
	// Data returned by pagar-me
	//
	Object               string     `json:"object,omitempty" form:"object,omitempty"`
	Status               TrStatus   `json:"status,omitempty" form:"status,omitempty"`
	RefuseReason         string     `json:"refuse_reason,omitempty" form:"refuse_reason,omitempty"`
	StatusReason         string     `json:"status_reason,omitempty" form:"status_reason,omitempty"`
	AcquirerResponseCode string     `json:"acquirer_response_code,omitempty" form:"acquirer_response_code,omitempty"`
	AcquirerName         string     `json:"acquirer_name,omitempty" form:"acquirer_name,omitempty"`
	AcquirerID           string     `json:"acquirer_id,omitempty" form:"acquirer_id,omitempty"`
	AuthorizationCode    string     `json:"authorization_code,omitempty" form:"authorization_code,omitempty"`
	TID                  int        `json:"tid,omitempty" form:"tid,omitempty"`
	NSU                  int        `json:"nsu,omitempty" form:"nsu,omitempty"`
	DateCreated          *time.Time `json:"date_created,omitempty" form:"date_created,omitempty"`
	DateUpdated          *time.Time `json:"date_updated,omitempty" form:"date_updated,omitempty"`
	AuthorizedAmount     int        `json:"authorized_amount,omitempty" form:"authorized_amount,omitempty"`
	PaidAmount           int        `json:"paid_amount,omitempty" form:"paid_amount,omitempty"`
	RefundedAmount       int        `json:"refunded_amount,omitempty" form:"refunded_amount,omitempty"`
	ID                   int64      `json:"id,omitempty" form:"id,omitempty"`
	Cost                 int        `json:"cost,omitempty" form:"cost,omitempty"`
	CardFirstDigits      string     `json:"card_first_digits,omitempty" form:"card_first_digits,omitempty"`
	CardLastDigits       string     `json:"card_last_digits,omitempty" form:"card_last_digits,omitempty"`
	CardBrand            string     `json:"card_brand,omitempty" form:"card_brand,omitempty"`
	Card                 *Card      `json:"card,omitempty" form:"card,omitempty"`
	CaptureMethod        string     `json:"capture_method,omitempty" form:"capture_method,omitempty"`
	CardPinMode          string     `json:"card_pin_mode,omitempty" form:"card_pin_mode,omitempty"`
	// card_magstripe_fallback
	AntifraudScore float64 `json:"antifraud_score,omitempty" form:"antifraud_score,omitempty"`
	BoletoURL      string  `json:"boleto_url,omitempty" form:"boleto_url,omitempty"`
	BoletoBarcode  string  `json:"boleto_barcode,omitempty" form:"boleto_barcode,omitempty"`
	Referer        string  `json:"referer,omitempty" form:"referer,omitempty"`
	IP             string  `json:"ip,omitempty" form:"ip,omitempty"`
	SubscriptionID string  `json:"subscription_id,omitempty" form:"subscription_id,omitempty"`
	// phone
	// address
	AntifraudMetadata map[string]interface{} `json:"antifraud_metadata,omitempty" form:"antifraud_metadata,omitempty"`
	// reference_key
	// device
	// local_transaction_id
	// fraud_covered
	// order_id
	RiskLevel string `json:"risk_level,omitempty" form:"risk_level,omitempty"`
	// receipt_url
	// payment
	// addition
	// discount
}

type BankAccount struct {
	ID             int32           `json:"id,omitempty"`
	BankAccountID  string          `json:"bank_account_id,omitempty"`
	BankCode       string          `json:"bank_code,omitempty"`
	Agencia        string          `json:"agencia,omitempty"`
	AgenciaDV      string          `json:"agencia_dv,omitempty"`
	Conta          string          `json:"conta,omitempty"`
	ContaDV        string          `json:"conta_dv,omitempty"`
	DocumentNumber string          `json:"document_number,omitempty"`
	LegalName      string          `json:"legal_name,omitempty"`
	Type           BankAccountType `json:"type,omitempty"`
}
