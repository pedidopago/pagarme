package pagarme

import "time"

// Payable -> Objeto-resposta Recebível
//
// Objeto contendo os dados de um recebível. O recebível (payable) é gerado automaticamente após uma transação ser paga.
// Para cada parcela de uma transação é gerado um recebível, que também pode ser dividido por recebedor (no caso de um split ter sido feito).
type Payable struct {
	Object                   string        `json:"object"`
	ID                       int           `json:"id"`
	Status                   PayableStatus `json:"status"`
	Amount                   int           `json:"amount"`
	Fee                      int           `json:"fee"`
	AnticipationFee          int           `json:"anticipation_fee"`
	FraudCoverageFee         int           `json:"fraud_coverage_fee"`
	Installment              int           `json:"installment"`
	TransactionID            int           `json:"transaction_id"`
	SplitRuleID              string        `json:"split_rule_id,omitempty"`
	BulkAnticipationID       string        `json:"bulk_anticipation_id,omitempty"`
	RecipientID              string        `json:"recipient_id,omitempty"`
	PaymentDate              time.Time     `json:"payment_date"`
	OriginalPaymentDate      time.Time     `json:"original_payment_date,omitempty"`
	Type                     PayableType   `json:"type"`
	PaymentMethod            PaymentMethod `json:"payment_method"`
	AccrualDate              time.Time     `json:"accrual_date,omitempty"`
	DateCreated              time.Time     `json:"date_created"`
	LiquidationArrangementID string        `json:"liquidation_arrangement_id,omitempty"`
}
