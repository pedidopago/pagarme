package transactions

import (
	"net/http"
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
	Amount         int                     `json:"amount,omitempty"`
	SplitRules     []*pagarme.SplitRule    `json:"split_rules,omitempty"`
	BankAccountID  string                  `json:"bank_account_id,omitempty"`
	BankCode       string                  `json:"bank_code,omitempty"`
	Agencia        string                  `json:"agencia,omitempty"`
	AgenciaDV      string                  `json:"agencia_dv,omitempty"`
	Conta          string                  `json:"conta,omitempty"`
	ContaDV        string                  `json:"conta_dv,omitempty"`
	DocumentNumber string                  `json:"document_number,omitempty"`
	LegalName      string                  `json:"legal_name,omitempty"`
	Async          bool                    `json:"async,omitempty"`
	Type           pagarme.BankAccountType `json:"type,omitempty"`
	Metadata       map[string]interface{}  `json:"metadata,omitempty"`
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
