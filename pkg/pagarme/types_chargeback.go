package pagarme

import "time"

type Chargeback struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Installment int       `json:"installment"`
	Amount      int       `json:"amount"`
	ReasonCode  string    `json:"reason_code"`
	CardBrand   string    `json:"card_brand"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	AccrualDate time.Time `json:"accrual_date"`
	Status      string    `json:"status"`
	Cycle       int       `json:"cycle"`
}

type ChargebackStatus string

const (
	ChargebackStatusPresented   ChargebackStatus = "presented"
	ChargebackStatusRepresented ChargebackStatus = "represented"
)
