package pagarme

import "time"

type Transfer struct {
	Object               string                 `json:"object,omitempty"`
	Id                   int                    `json:"id,omitempty"`
	Amount               int                    `json:"amount,omitempty"`
	Type                 TransferType           `json:"type,omitempty"`
	Status               TransferStatus         `json:"status,omitempty"`
	SourceType           string                 `json:"source_type,omitempty"`
	SourceId             string                 `json:"source_id,omitempty"`
	TargetType           string                 `json:"target_type,omitempty"`
	TargetId             string                 `json:"target_id,omitempty"`
	Fee                  int                    `json:"fee,omitempty"`
	FundingDate          time.Time              `json:"funding_date,omitempty"`
	FundingEstimatedDate time.Time              `json:"funding_estimated_date,omitempty"`
	TransactionId        int                    `json:"transaction_id,omitempty"`
	DateCreated          time.Time              `json:"date_created,omitempty"`
	DateUpdated          time.Time              `json:"date_updated,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	BankAccount          BankAccount            `json:"bank_account"`
}
