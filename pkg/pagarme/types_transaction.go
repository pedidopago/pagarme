package pagarme

import (
	"time"
)

type Transaction struct {
	Amount               int                    `json:"amount"`
	CardHash             string                 `json:"card_hash,omitempty"`
	CardID               string                 `json:"card_id,omitempty"`
	CardHolderName       string                 `json:"card_holder_name,omitempty"`
	CardExpirationDate   MMYY                   `json:"card_expiration_date,omitempty"`
	CardNumber           string                 `json:"card_number,omitempty"`
	CardCVV              string                 `json:"card_cvv,omitempty"`
	PaymentMethod        PaymentMethod          `json:"payment_method,omitempty"`
	PostbackURL          string                 `json:"postback_url,omitempty"`
	Async                bool                   `json:"async,omitempty"`
	Installments         int                    `json:"installments,omitempty"`
	BoletoExpirationDate string                 `json:"boleto_expiration_date,omitempty"`
	SoftDescriptor       string                 `json:"soft_descriptor,omitempty"`
	Capture              string                 `json:"capture,omitempty"`
	BoletoInstructions   string                 `json:"boleto_instructions,omitempty"`
	SplitRules           []*SplitRule           `json:"split_rules,omitempty"`
	Customer             *Customer              `json:"customer"`
	Billing              *Billing               `json:"billing"`
	Shipping             *Shipping              `json:"shipping"`
	Items                []*Item                `json:"items,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	Session              string                 `json:"session,omitempty"`
	LocalTime            *time.Time             `json:"local_time,omitempty"`
	//
	// Data returned by pagar-me
	//
	Object               string     `json:"object,omitempty"`
	Status               TrStatus   `json:"status,omitempty"`
	RefuseReason         string     `json:"refuse_reason,omitempty"`
	StatusReason         string     `json:"status_reason,omitempty"`
	AcquirerResponseCode string     `json:"acquirer_response_code,omitempty"`
	AcquirerName         string     `json:"acquirer_name,omitempty"`
	AcquirerID           string     `json:"acquirer_id,omitempty"`
	AuthorizationCode    string     `json:"authorization_code,omitempty"`
	TID                  int        `json:"tid,omitempty"`
	NSU                  int        `json:"nsu,omitempty"`
	DateCreated          *time.Time `json:"date_created,omitempty"`
	DateUpdated          *time.Time `json:"date_updated,omitempty"`
	AuthorizedAmount     int        `json:"authorized_amount,omitempty"`
	PaidAmount           int        `json:"paid_amount,omitempty"`
	RefundedAmount       int        `json:"refunded_amount,omitempty"`
	ID                   int64      `json:"id,omitempty"`
	Cost                 int        `json:"cost,omitempty"`
	CardFirstDigits      string     `json:"card_first_digits,omitempty"`
	CardLastDigits       string     `json:"card_last_digits,omitempty"`
	CardBrand            string     `json:"card_brand,omitempty"`
	Card                 *Card      `json:"card,omitempty"`
	CaptureMethod        string     `json:"capture_method,omitempty"`
	// card_pin_mode
	// card_magstripe_fallback
	// antifraud_score
	BoletoURL     string `json:"boleto_url,omitempty"`
	BoletoBarcode string `json:"boleto_barcode,omitempty"`
	Referer       string `json:"referer,omitempty"`
	IP            string `json:"ip,omitempty"`
	// subscription_id
	// phone
	// address
	// antifraud_metadata
	// reference_key
	// device
	// local_transaction_id
	// fraud_covered
	// order_id
	RiskLevel string `json:"risk_level,omitempty"`
	// receipt_url
	// payment
	// addition
	// discount
}
