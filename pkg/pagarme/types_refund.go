package pagarme

import "time"

type Refund struct {
	Object               string         `json:"object"`
	Id                   string         `json:"id"`
	Amount               int            `json:"amount"`
	Fee                  int            `json:"fee"`
	FraudCoverageFee     int            `json:"fraud_coverage_fee"`
	Type                 string         `json:"type"`
	Status               string         `json:"status"`
	ChargeFeeRecipientId string         `json:"charge_fee_recipient_id"`
	BankAccountId        int            `json:"bank_account_id"`
	TransactionId        int            `json:"transaction_id"`
	DateCreated          time.Time      `json:"date_created"`
	DateUpdated          time.Time      `json:"date_updated"`
	Metadata             map[string]any `json:"metadata"`
}

type RefundStatus string

const (
	RefundStatusRefunded      RefundStatus = "refunded"
	RefundStatusPendingRefund RefundStatus = "pending_refund"
)
