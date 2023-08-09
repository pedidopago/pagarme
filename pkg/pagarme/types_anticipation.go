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
	Object          string                `json:"object,omitempty"`
	ID              string                `json:"id,omitempty"`
	Status          AnticipationStatus    `json:"status,omitempty"`
	Amount          int                   `json:"amount,omitempty"`
	Fee             int                   `json:"fee,omitempty"`
	AnticipationFee int                   `json:"anticipation_fee,omitempty"`
	Type            AnticipationType      `json:"type,omitempty"`
	Timeframe       AnticipationTimeframe `json:"timeframe,omitempty"`
	PaymentDate     time.Time             `json:"payment_date,omitempty"`
	DateCreated     time.Time             `json:"date_created,omitempty"`
	DateUpdated     time.Time             `json:"date_updated,omitempty"`
}

type AnticipationSimulation struct {
	Amount             int                   `json:"amount"`
	Fee                int                   `json:"fee"`
	FraudCoverageFee   int                   `json:"fraudCoverageFee"`
	AnticipationAmount int                   `json:"anticipationAmount"`
	AnticipationFee    int                   `json:"anticipationFee"`
	Timeframe          AnticipationTimeframe `json:"timeframe"`
	PaymentDate        time.Time             `json:"paymentDate"`
	StartIntervalDate  time.Time             `json:"startIntervalDate"`
	EndIntervalDate    time.Time             `json:"endIntervalDate"`
}
