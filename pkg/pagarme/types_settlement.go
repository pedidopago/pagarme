package pagarme

type Settlement struct {
	Id                       string            `json:"id"`
	Status                   SettlementStatus  `json:"status"`
	Amount                   string            `json:"amount"`
	Product                  SettlementProduct `json:"product"`
	CardBrand                string            `json:"card_brand,omitempty"`
	CompanyId                string            `json:"company_id"`
	PaymentDate              string            `json:"payment_date"`
	RecipientId              string            `json:"recipient_id"`
	DocumentType             string            `json:"document_type"`
	DocumentNumber           string            `json:"document_number"`
	ContractObligationId     string            `json:"contract_obligation_id,omitempty"`
	LiquidationArrangementId string            `json:"liquidation_arrangement_id"`
	LiquidationEngine        string            `json:"liquidation_engine,omitempty"`
	LiquidationType          LiquidationType   `json:"liquidation_type"`
	Ispb                     string            `json:"ispb"`
	TargetAccountIspb        string            `json:"target_account.ispb"`
}
