package pagarme

import (
	"encoding/json"
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
	PixExpirationDate    string                 `json:"pix_expiration_date,omitempty" form:"pix_expiration_date,omitempty"`
	PixAdditionalFields  []NameValue            `json:"pix_additional_fields,omitempty" form:"pix_additional_fields,omitempty"`
	SplitRules           []*SplitRule           `json:"split_rules,omitempty" form:"split_rules,omitempty"`
	BoletoRules          []BoletoRule           `json:"boleto_rules,omitempty" form:"boleto_rules,omitempty"`
	Customer             *Customer              `json:"customer" form:"customer"`
	Billing              *Billing               `json:"billing" form:"billing"`
	Shipping             *Shipping              `json:"shipping" form:"shipping"`
	Items                []*Item                `json:"items,omitempty" form:"items,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty" form:"metadata,omitempty"`
	Session              string                 `json:"session,omitempty" form:"session,omitempty"`
	LocalTime            *time.Time             `json:"local_time,omitempty" form:"local_time,omitempty"`
	ReferenceKey         string                 `json:"reference_key,omitempty" form:"reference_key,omitempty"`
	//
	// Data returned by pagar-me
	//
	Object               string         `json:"object,omitempty" form:"object,omitempty"`
	Status               TrStatus       `json:"status,omitempty" form:"status,omitempty"`
	RefuseReason         string         `json:"refuse_reason,omitempty" form:"refuse_reason,omitempty"`
	StatusReason         string         `json:"status_reason,omitempty" form:"status_reason,omitempty"`
	AcquirerResponseCode string         `json:"acquirer_response_code,omitempty" form:"acquirer_response_code,omitempty"`
	AcquirerName         string         `json:"acquirer_name,omitempty" form:"acquirer_name,omitempty"`
	AcquirerID           string         `json:"acquirer_id,omitempty" form:"acquirer_id,omitempty"`
	AuthorizationCode    string         `json:"authorization_code,omitempty" form:"authorization_code,omitempty"`
	TID                  int            `json:"tid,omitempty" form:"tid,omitempty"`
	NSU                  int            `json:"nsu,omitempty" form:"nsu,omitempty"`
	DateCreated          *time.Time     `json:"date_created,omitempty" form:"date_created,omitempty"`
	DateUpdated          *time.Time     `json:"date_updated,omitempty" form:"date_updated,omitempty"`
	AuthorizedAmount     int            `json:"authorized_amount,omitempty" form:"authorized_amount,omitempty"`
	PaidAmount           int            `json:"paid_amount,omitempty" form:"paid_amount,omitempty"`
	RefundedAmount       int            `json:"refunded_amount,omitempty" form:"refunded_amount,omitempty"`
	ID                   int64          `json:"id,omitempty" form:"id,omitempty"`
	Cost                 int            `json:"cost,omitempty" form:"cost,omitempty"`
	CardFirstDigits      string         `json:"card_first_digits,omitempty" form:"card_first_digits,omitempty"`
	CardLastDigits       string         `json:"card_last_digits,omitempty" form:"card_last_digits,omitempty"`
	CardBrand            string         `json:"card_brand,omitempty" form:"card_brand,omitempty"`
	Card                 *Card          `json:"card,omitempty" form:"card,omitempty"`
	CaptureMethod        string         `json:"capture_method,omitempty" form:"capture_method,omitempty"`
	CardPinMode          string         `json:"card_pin_mode,omitempty" form:"card_pin_mode,omitempty"`
	PixData              map[string]any `json:"pix_data,omitempty" form:"pix_data,omitempty"`
	PixQrCode            string         `json:"pix_qr_code,omitempty" form:"pix_qr_code,omitempty"`
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

func GetBoletoPdfUrl(boletoUrl string) string {
	if boletoUrl == "" {
		return ""
	}
	return boletoUrl + "?format=pdf"
}

func (t *Transaction) GetBoletoPdfUrl() string {
	return GetBoletoPdfUrl(t.BoletoURL)
}

type PixData struct {
	EndToEndId          string     `json:"end_to_end_id"`
	ExpirationDate      time.Time  `json:"expiration_date"`
	Payer               PixPayer   `json:"payer"`
	QrCode              string     `json:"qr_code"`
	RefundFailureDate   *time.Time `json:"refund_failure_date"`
	RefundFailureReason string     `json:"refund_failure_reason,omitempty"`
}

type PixPayer struct {
	BankAccount  PixPayerBankAccount `json:"bank_account"`
	Name         string              `json:"name"`
	Document     string              `json:"document"`
	DocumentType string              `json:"document_type"`
}

type PixPayerBankAccount struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	BranchCode    string `json:"branch_code"`
	Ispb          string `json:"ispb"`
}

func (t *Transaction) GetPixData() *PixData {
	if t.PixData == nil {
		return nil
	}
	var data PixData
	b, _ := json.Marshal(t.PixData)
	if json.Unmarshal(b, &data) != nil {
		return nil
	}
	return &data
}

type BankAccount struct {
	ID                 int32           `json:"id,omitempty"`
	BankAccountID      string          `json:"bank_account_id,omitempty"`
	BankCode           string          `json:"bank_code,omitempty"`
	Agencia            string          `json:"agencia,omitempty"`
	AgenciaDV          string          `json:"agencia_dv,omitempty"`
	Conta              string          `json:"conta,omitempty"`
	ContaDV            string          `json:"conta_dv,omitempty"`
	DocumentType       DocumentType    `json:"document_type,omitempty"`
	DocumentNumber     string          `json:"document_number,omitempty"`
	LegalName          string          `json:"legal_name,omitempty"`
	Type               BankAccountType `json:"type,omitempty"`
	ChargeTransferFees bool            `json:"charge_transfer_fees"`
	PixKey             string          `json:"pix_key,omitempty"`
	DateCreated        time.Time       `json:"date_created,omitempty"`
}

// CardHash wraps the response from new card hash public key request
type CardHash struct {
	ID          int64      `json:"id,omitempty" form:"id,omitempty"`
	PublicKey   string     `json:"public_key,omitempty" form:"public_key,omitempty"`
	IP          string     `json:"ip,omitempty" form:"ip,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty" form:"date_created,omitempty"`
}

// CardHashData contains the fields used to generate a card_hash
type CardHashData struct {
	CardNumber         string `json:"card_number,omitempty" form:"card_number,omitempty"`
	CardHolderName     string `json:"card_holder_name,omitempty" form:"card_holder_name,omitempty"`
	CardExpirationDate MMYY   `json:"card_expiration_date,omitempty" form:"card_expiration_date,omitempty"`
	CardCvv            string `json:"card_cvv,omitempty" form:"card_cvv,omitempty"`
}

type BoletoRule string

const (
	RuleStrictExpirationDate BoletoRule = "strict_expiration_date"
	RuleNoStrict             BoletoRule = "no_strict"
)

type NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TransactionOperation struct {
	Id                    string    `json:"id"`
	DateCreated           time.Time `json:"date_created"`
	DateUpdated           time.Time `json:"date_updated"`
	Status                string    `json:"status"`
	FailReason            string    `json:"fail_reason,omitempty"`
	Type                  string    `json:"type"`
	Rollbacked            bool      `json:"rollbacked"`
	Model                 string    `json:"model"`
	ModelId               string    `json:"model_id"`
	GroupId               string    `json:"group_id"`
	NextGroupId           string    `json:"next_group_id,omitempty"`
	RequestId             string    `json:"request_id"`
	StartedAt             int64     `json:"started_at"`
	EndedAt               int64     `json:"ended_at"`
	Processor             string    `json:"processor"`
	ProcessorResponseCode string    `json:"processor_response_code,omitempty"`
	Metadata              any       `json:"metadata"`
}
