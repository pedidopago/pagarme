package pagarme

import (
	"time"

	"github.com/gabstv/sqltypes"
)

type Transaction struct {
	Amount               sqltypes.NullInt0      `json:"amount" form:"amount"`
	CardHash             sqltypes.NullString    `json:"card_hash,omitempty" form:"card_hash,omitempty"`
	CardID               sqltypes.NullString    `json:"card_id,omitempty" form:"card_id,omitempty"`
	CardHolderName       sqltypes.NullString    `json:"card_holder_name,omitempty" form:"card_holder_name,omitempty"`
	CardExpirationDate   MMYY                   `json:"card_expiration_date,omitempty" form:"card_expiration_date,omitempty"`
	CardNumber           sqltypes.NullString    `json:"card_number,omitempty" form:"card_number,omitempty"`
	CardCVV              sqltypes.NullString    `json:"card_cvv,omitempty" form:"card_cvv,omitempty"`
	PaymentMethod        PaymentMethod          `json:"payment_method,omitempty" form:"payment_method,omitempty"`
	PostbackURL          sqltypes.NullString    `json:"postback_url,omitempty" form:"postback_url,omitempty"`
	Async                bool                   `json:"async,omitempty" form:"async,omitempty"`
	Installments         sqltypes.NullInt0      `json:"installments,omitempty" form:"installments,omitempty"`
	BoletoExpirationDate sqltypes.NullString    `json:"boleto_expiration_date,omitempty" form:"boleto_expiration_date,omitempty"`
	SoftDescriptor       sqltypes.NullString    `json:"soft_descriptor,omitempty" form:"soft_descriptor,omitempty"`
	Capture              sqltypes.NullString    `json:"capture,omitempty" form:"capture,omitempty"`
	BoletoInstructions   sqltypes.NullString    `json:"boleto_instructions,omitempty" form:"boleto_instructions,omitempty"`
	SplitRules           []*SplitRule           `json:"split_rules,omitempty" form:"split_rules,omitempty"`
	Customer             *Customer              `json:"customer" form:"customer"`
	Billing              *Billing               `json:"billing" form:"billing"`
	Shipping             *Shipping              `json:"shipping" form:"shipping"`
	Items                []*Item                `json:"items,omitempty" form:"items,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty" form:"metadata,omitempty"`
	Session              sqltypes.NullString    `json:"session,omitempty" form:"session,omitempty"`
	LocalTime            *time.Time             `json:"local_time,omitempty" form:"local_time,omitempty"`
	//
	// Data returned by pagar-me
	//
	Object               sqltypes.NullString `json:"object,omitempty" form:"object,omitempty"`
	Status               TrStatus            `json:"status,omitempty" form:"status,omitempty"`
	RefuseReason         sqltypes.NullString `json:"refuse_reason,omitempty" form:"refuse_reason,omitempty"`
	StatusReason         sqltypes.NullString `json:"status_reason,omitempty" form:"status_reason,omitempty"`
	AcquirerResponseCode sqltypes.NullString `json:"acquirer_response_code,omitempty" form:"acquirer_response_code,omitempty"`
	AcquirerName         sqltypes.NullString `json:"acquirer_name,omitempty" form:"acquirer_name,omitempty"`
	AcquirerID           sqltypes.NullString `json:"acquirer_id,omitempty" form:"acquirer_id,omitempty"`
	AuthorizationCode    sqltypes.NullString `json:"authorization_code,omitempty" form:"authorization_code,omitempty"`
	TID                  sqltypes.NullInt0   `json:"tid,omitempty" form:"tid,omitempty"`
	NSU                  sqltypes.NullInt0   `json:"nsu,omitempty" form:"nsu,omitempty"`
	DateCreated          *time.Time          `json:"date_created,omitempty" form:"date_created,omitempty"`
	DateUpdated          *time.Time          `json:"date_updated,omitempty" form:"date_updated,omitempty"`
	AuthorizedAmount     sqltypes.NullInt0   `json:"authorized_amount,omitempty" form:"authorized_amount,omitempty"`
	PaidAmount           sqltypes.NullInt0   `json:"paid_amount,omitempty" form:"paid_amount,omitempty"`
	RefundedAmount       sqltypes.NullInt0   `json:"refunded_amount,omitempty" form:"refunded_amount,omitempty"`
	ID                   int64               `json:"id,omitempty" form:"id,omitempty"`
	Cost                 sqltypes.NullInt0   `json:"cost,omitempty" form:"cost,omitempty"`
	CardFirstDigits      sqltypes.NullString `json:"card_first_digits,omitempty" form:"card_first_digits,omitempty"`
	CardLastDigits       sqltypes.NullString `json:"card_last_digits,omitempty" form:"card_last_digits,omitempty"`
	CardBrand            sqltypes.NullString `json:"card_brand,omitempty" form:"card_brand,omitempty"`
	Card                 *Card               `json:"card,omitempty" form:"card,omitempty"`
	CaptureMethod        sqltypes.NullString `json:"capture_method,omitempty" form:"capture_method,omitempty"`
	CardPinMode          sqltypes.NullString `json:"card_pin_mode,omitempty" form:"card_pin_mode,omitempty"`
	// card_magstripe_fallback
	AntifraudScore float64             `json:"antifraud_score,omitempty" form:"antifraud_score,omitempty"`
	BoletoURL      sqltypes.NullString `json:"boleto_url,omitempty" form:"boleto_url,omitempty"`
	BoletoBarcode  sqltypes.NullString `json:"boleto_barcode,omitempty" form:"boleto_barcode,omitempty"`
	Referer        sqltypes.NullString `json:"referer,omitempty" form:"referer,omitempty"`
	IP             sqltypes.NullString `json:"ip,omitempty" form:"ip,omitempty"`
	SubscriptionID sqltypes.NullString `json:"subscription_id,omitempty" form:"subscription_id,omitempty"`
	// phone
	// address
	AntifraudMetadata map[string]interface{} `json:"antifraud_metadata,omitempty" form:"antifraud_metadata,omitempty"`
	// reference_key
	// device
	// local_transaction_id
	// fraud_covered
	// order_id
	RiskLevel sqltypes.NullString `json:"risk_level,omitempty" form:"risk_level,omitempty"`
	// receipt_url
	// payment
	// addition
	// discount
}

type BankAccount struct {
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
