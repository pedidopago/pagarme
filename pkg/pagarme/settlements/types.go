package settlements

import (
	"fmt"
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

// PaymentDateStart -> Filtro pela data de pagamento da liquidação
// deve ser fornecido juntamente com PaymentDateEnd
func (qi *QueryInput) PaymentDateStart(yyyy int, mm time.Month, dd int) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "payment_date_start",
		Op:   pagarme.QueryOpEquals,
		V:    fmt.Sprintf("%d-%02d-%02d", yyyy, mm, dd),
	})
	return qi
}

// PaymentDateEnd -> Filtro pela data de pagamento da liquidação
// deve ser fornecido juntamente com PaymentDateStart
func (qi *QueryInput) PaymentDateEnd(yyyy int, mm time.Month, dd int) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "payment_date_end",
		Op:   pagarme.QueryOpEquals,
		V:    fmt.Sprintf("%d-%02d-%02d", yyyy, mm, dd),
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

// Build builds the settlement query to a urlencoded format
func (qi *QueryInput) Build() string {
	qi.init()
	return qi.b.Build()
}

func (qi *QueryInput) Values() url.Values {
	qi.init()
	return qi.b.Values()
}
