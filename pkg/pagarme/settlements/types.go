package settlements

import (
	"fmt"
	"net/url"
	"strconv"
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

// PaymentDate -> Filtro pela data de pagamento da liquidação (não considera o campo das horas)
func (qi *QueryInput) PaymentDate(start, end time.Time) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "payment_date_start",
		Op:   pagarme.QueryOpEquals,
		V:    fmt.Sprintf("%d-%02d-%02d", start.Year(), start.Month(), start.Day()),
	})
	qi.b.Add(&pagarme.QueryString{
		Name: "payment_date_end",
		Op:   pagarme.QueryOpEquals,
		V:    fmt.Sprintf("%d-%02d-%02d", end.Year(), end.Month(), end.Day()),
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

func (qi *QueryInput) SettlementID(v string) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "settlement_id",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

func (qi *QueryInput) GetIspb(v bool) *QueryInput {
	qi.init()
	qi.b.Add(&pagarme.QueryString{
		Name: "get_ispb",
		Op:   pagarme.QueryOpEquals,
		V:    strconv.FormatBool(v),
	})
	return qi
}

// Count -> Parâmetro de quantos resultados devem ser retornados
func (qi *QueryInput) Count(v int) *QueryInput {
	qi.init()
	qi.b.Set(&pagarme.QueryInt{
		Name: "limit",
		Op:   pagarme.QueryOpEquals,
		V:    v,
	})
	return qi
}

// GetCount -> retorna quantos resultados devem ser retornados
func (qi *QueryInput) GetCount() int {
	qi.init()
	qiface := qi.b.Get("limit")
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

func (qi *QueryInput) GetPage() int {
	qi.init()
	qiface := qi.b.Get("page")
	if qiface == nil {
		return 0
	}
	if q, ok := qiface.(*pagarme.QueryInt); ok {
		return q.V
	}
	return 0
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
