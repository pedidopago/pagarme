package payables

import (
	"net/url"
	"time"

	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

// QueryInput is the input data of Query
type QueryInput struct {
	b *pagarme.QueryBuilder
}

func (qi *QueryInput) init() {
	if qi.b == nil {
		qi.b = &pagarme.QueryBuilder{}
	}
}

// CreatedAt -> Filtro pela data de criação do payable
func (qi *QueryInput) CreatedAt(op pagarme.QueryOp, t time.Time) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryTime{
		Name: "created_at",
		Op:   op,
		T:    t,
	})
	return qi
}

// PaymentDate -> Filtro pela data de pagamento do recebível
func (qi *QueryInput) PaymentDate(op pagarme.QueryOp, t time.Time) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryTime{
		Name: "payment_date",
		Op:   op,
		T:    t,
	})
	return qi
}

// Amount -> Filtro pelo valor do recebível
func (qi *QueryInput) Amount(op pagarme.QueryOp, v int) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryInt{
		Name: "amount",
		Op:   op,
		V:    v,
	})
	return qi
}

// RecipientID -> Filtro pelo ID do recebedor atrelado
func (qi *QueryInput) RecipientID(v string) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "recipient_id",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

// BulkAnticipationID -> Filtro pelo id da antecipação
func (qi *QueryInput) BulkAnticipationID(v string) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "bulk_anticipation_id",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

// Status -> Filtro pelo status do recebível. paid ou waiting_funds
func (qi *QueryInput) Status(v pagarme.PayableStatus) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "status",
		Op:   pagarme.QueryOpEquals,
		V:    string(v),
	})
	return qi
}

// Installment -> Filtro pelo installment do recebível - a qual parcela da transação o recebível se refere
func (qi *QueryInput) Installment(op pagarme.QueryOp, v int) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryInt{
		Name: "installment",
		Op:   op,
		V:    v,
	})
	return qi
}

// TransactionID -> Filtro pelo ID da transação referida pelo recebível
func (qi *QueryInput) TransactionID(v string) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "transaction_id",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

// SplitRuleID -> Filtro pelo ID da regra de split atrelada
func (qi *QueryInput) SplitRuleID(v string) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "split_rule_id",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

// Type -> Filtro pelo type do recebível. Pode ser chargeback, refund, chargeback_refund ou credit
func (qi *QueryInput) Type(v pagarme.PayableType) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "type",
		Op:   pagarme.QueryOpEquals,
		V:    string(v),
	})
	return qi
}

func (qi *QueryInput) LiquidationArrangementID(v string) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "liquidation_arrangement_id",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

// ID -> Filtro pelo ID do recebível
func (qi *QueryInput) ID(v string) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "id",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

// Count -> Parâmetro de quantos resultados devem ser retornados
func (qi *QueryInput) Count(v int) *QueryInput {
	qi.init()
	qi.b.Set(&pagarme.QueryInt{
		Name: "count",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

// GetCount -> retorna quantos resultados devem ser retornados
func (qi *QueryInput) GetCount() int {
	qi.init()
	qiface := qi.b.Get("count")
	if qiface == nil {
		return 0
	}
	if q, ok := qiface.(*pagarme.QueryInt); ok {
		return q.V
	}
	return 0
}

// Page -> Parâmetro de paginação: aplica um offset de page * count nos resultados
func (qi *QueryInput) Page(v int) *QueryInput {
	qi.init()
	qi.b.Set(&pagarme.QueryInt{
		Name: "page",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

// Build builds the payable query to a urlencoded format
func (qi *QueryInput) Build() string {
	qi.init()
	return qi.b.Build()
}

func (qi *QueryInput) Values() url.Values {
	qi.init()
	return qi.b.Values()
}
