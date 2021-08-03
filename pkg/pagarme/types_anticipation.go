package pagarme

import "time"

type CreateAnticipation struct {
	PaymentDate       int64                 `json:"payment_date"`
	Timeframe         AnticipationTimeframe `json:"timeframe"`
	RequestedAmount   int                   `json:"requested_amount"`
	Build             bool                  `json:"build"`
	AutomaticTransfer bool                  `json:"automatic_transfer"`
}

type Limit struct {
	Amount          int `json:"amount,omitempty"`
	AnticipationFee int `json:"anticipation_fee,omitempty"`
	Fee             int `json:"fee,omitempty"`
}

type Limits struct {
	Maximum Limit `json:"maximum,omitempty"`
	Minimum Limit `json:"minimum,omitempty"`
}

type Anticipation struct {
	ID              string                `json:"id,omitempty"`
	Status          AnticipationStatus    `json:"status,omitempty"`
	Timeframe       AnticipationTimeframe `json:"timeframe,omitempty"`
	PaymentDate     time.Time             `json:"payment_date,omitempty"`
	Amount          int                   `json:"amount,omitempty"`
	Fee             int                   `json:"fee,omitempty"`
	AnticipationFee int                   `json:"anticipation_fee,omitempty"`
}
