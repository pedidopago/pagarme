package transactions

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pedidopago/pagarme/internal/pkg/www"
	"github.com/pedidopago/pagarme/pkg/pagarme"
)

// API is the /1/transactions API
type API struct {
	Config *pagarme.Config
}

// New /1/transactions API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

// Put creates a new transaction
//
// POST https://api.pagar.me/1/transactions
func (api *API) Put(tr *pagarme.Transaction) (*pagarme.Response, *pagarme.Transaction, error) {
	resp, err := api.Config.Do(http.MethodPost, "/transactions", www.JSONReader(tr))
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}
	result := &pagarme.Transaction{}

	if api.Config.Trace {
		if err := www.UnmarshalTrace(api.Config.Logger, resp, result); err != nil {
			api.Config.Logger.Error("could not unmarshal transaction: " + err.Error())
			return nil, nil, err
		}
	} else {
		if err := www.Unmarshal(resp, result); err != nil {
			api.Config.Logger.Error("could not unmarshal transaction [Put]: " + err.Error())
			return nil, nil, err
		}
	}
	return www.Ok(), result, nil
}

// Get a transaction by ID
//
// GET https://api.pagar.me/1/transactions/id
func (api *API) Get(tid string) (*pagarme.Response, *pagarme.Transaction, error) {
	resp, err := api.Config.Do(http.MethodGet, "/transactions/"+tid, nil)
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}
	result := &pagarme.Transaction{}
	if err := www.Unmarshal(resp, result); err != nil {
		api.Config.Logger.Error("could not unmarshal transaction [Get]: " + err.Error())
		return nil, nil, err
	}

	return www.Ok(), result, nil
}

type QueryInput struct {
	Count           int
	Page            int
	Filter          string
	Value           string
	Metadata        map[string]string
	DateCreatedFrom string // Unix timestamp WITH MILLISECONDS
	DateCreatedTo   string // Unix timestamp WITH MILLISECONDS
}

func (qi *QueryInput) Export() string {
	vv := url.Values{}
	for k, v := range qi.Metadata {
		fmt.Println("K", k, "V", v)
		vv.Set("metadata["+k+"]", v)
	}
	if qi.Filter != "" {
		vv.Set(qi.Filter, qi.Value)
	}
	if qi.Count != 0 {
		vv.Set("count", strconv.Itoa(qi.Count))
	} else {
		vv.Set("count", "10")
	}
	if qi.Page != 0 {
		vv.Set("page", strconv.Itoa(qi.Page))
	} else {
		vv.Set("page", "1")
	}
	if qi.DateCreatedFrom != "" && qi.DateCreatedTo != "" {
		vv.Add("date_created", ">="+qi.DateCreatedFrom)
		vv.Add("date_created", "<="+qi.DateCreatedTo)
	}
	vvs := vv.Encode() //strings.Replace(vv.Encode(), "%2E", ".", -1)
	return vvs
}

// Query transactions
func (api *API) Query(input QueryInput) (*pagarme.Response, []pagarme.Transaction, error) {
	resp, err := api.Config.Do(http.MethodGet, "/transactions?"+input.Export(), nil)
	if api.Config.Trace {
		api.Config.Logger.Info("/transactions?" + input.Export())
	}
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}
	result := make([]pagarme.Transaction, 0)
	if err := www.Unmarshal(resp, &result); err != nil {
		api.Config.Logger.Error("could not unmarshal transactions [Query]: " + err.Error())
		return nil, nil, err
	}

	return www.Ok(), result, nil
}

// PutDevBillet (Testando pagamento de Boletos)
//
// Usado apenas em ambiente de Teste para simular o pagamento de um Boleto.
//
// PUT https://api.pagar.me/1/transactions/transaction_id
func (api *API) PutDevBillet(tid string, status pagarme.TrStatus) (*pagarme.Response, *pagarme.Transaction, error) {
	resp, err := api.Config.Do(http.MethodPut, "/transactions/"+tid, www.JSONReader(map[string]interface{}{
		"status": string(status),
	}))
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}
	result := &pagarme.Transaction{}
	if err := www.Unmarshal(resp, result); err != nil {
		api.Config.Logger.Error("could not unmarshal transaction [Get]: " + err.Error())
		return nil, nil, err
	}

	return www.Ok(), result, nil
}

// https://www.febraban.org.br/associados/utilitarios/bancos.asp?msg=&id_assunto=84&id_pasta=0&tipo=

// RefundInput is the input data for Refund
type RefundInput struct {
	// O estorno parcial obedece as mesmas regras de um estorno total, e usa o parâmetro amount como
	// referência para o valor a ser estornado. É bom observar que o status da transação vai permanecer
	// paid até que o valor total da transação tenha sido estornado.
	Amount      int                    `json:"amount,omitempty"`
	SplitRules  []*pagarme.SplitRule   `json:"split_rules,omitempty"`
	Async       bool                   `json:"async,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	BankAccount *pagarme.BankAccount   `json:"bank_account,omitempty"`
}

// Refund refunds a transaction
//
// POST https://api.pagar.me/1/transactions/transaction_id/refund
func (api *API) Refund(id int64, input *RefundInput) (*pagarme.Response, *pagarme.Transaction, error) {
	resp, err := api.Config.Do(http.MethodPost, "/transactions/"+strconv.Itoa(int(id))+"/refund", www.JSONReader(input))
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}

	result := &pagarme.Transaction{}

	if api.Config.Trace {
		if err := www.UnmarshalTrace(api.Config.Logger, resp, result); err != nil {
			api.Config.Logger.Error("could not unmarshal transaction: " + err.Error())
			return nil, nil, err
		}
	} else {
		if err := www.Unmarshal(resp, result); err != nil {
			api.Config.Logger.Error("could not unmarshal transaction: " + err.Error())
			return nil, nil, err
		}
	}
	return www.Ok(), result, nil
}
