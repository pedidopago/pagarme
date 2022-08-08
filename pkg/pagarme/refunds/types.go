package refunds

import (
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

// DateCreated -> Filtro pela data de criação do estorno
func (qi *QueryInput) DateCreated(op pagarme.QueryOp, t time.Time) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryTime{
		Name: "date_created",
		Op:   op,
		T:    t,
	})
	return qi
}

// DateUpdated -> Filtro pela data de atualização do estorno
func (qi *QueryInput) DateUpdated(op pagarme.QueryOp, t time.Time) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryTime{
		Name: "date_updated",
		Op:   op,
		T:    t,
	})
	return qi
}

// TransactionID -> Filtro pelo ID da transação referida pelo estorno
func (qi *QueryInput) TransactionID(v string) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "transaction_id",
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
