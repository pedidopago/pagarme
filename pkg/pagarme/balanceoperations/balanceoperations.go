package balanceoperations

import (
	"github.com/pedidopago/pagarme/internal/pkg/www"
	"github.com/pedidopago/pagarme/pkg/pagarme"
	"net/http"
)

// API is the /1/balance/operations API
type API struct {
	Config *pagarme.Config
}

// New /1/balance/operations API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

func (api *API) Get(id string) (response *pagarme.Response, operation *pagarme.BalanceOperation, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, "/balance/operations/" + id, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.BalanceOperation)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal balance operations: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal balance operations: [Get]" + rerr.Error())
			return
		}
	}

	operation = result
	response = www.Ok()
	return
}

// Query
//
// GET https://api.pagar.me/1/balance/operations
func (api *API) Query(params *pagarme.QueryBuilder) (response *pagarme.Response, operations []pagarme.BalanceOperation, rerr error) {
	url := "/balance/operations"
	if params != nil {
		url += "?" + params.Build()
	}
	resp, rerr := api.Config.Do(http.MethodGet, url, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.BalanceOperation, 0)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal balance operations: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal balance operations: [Query]" + rerr.Error())
			return
		}
	}

	operations = result
	response = www.Ok()
	return
}
