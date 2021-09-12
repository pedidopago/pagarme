package payables

import (
	"net/http"

	"github.com/pedidopago/pagarme/internal/pkg/www"
	"github.com/pedidopago/pagarme/pkg/pagarme"
)

// API is the /1/payables API
type API struct {
	Config *pagarme.Config
}

// New /1/payables API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

// Query -> GET https://api.pagar.me/1/payables
//
//
func (api *API) Query(input *QueryInput) (*pagarme.Response, []pagarme.Payable, error) {
	urlpart := "/payables"
	if input != nil {
		urlpart += "?" + input.Build()
	}
	resp, err := api.Config.Do(http.MethodGet, urlpart, nil)
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}
	result := make([]pagarme.Payable, 0)
	//
	if api.Config.Trace {
		if err := www.UnmarshalTrace(api.Config.Logger, resp, &result); err != nil {
			api.Config.Logger.Error("could not unmarshal payables: " + err.Error())
			return nil, nil, err
		}
	} else {
		if err := www.Unmarshal(resp, &result); err != nil {
			api.Config.Logger.Error("could not unmarshal payables [Get]: " + err.Error())
			return nil, nil, err
		}
	}
	return www.Ok(), result, nil
}


// Get retrieves a payable by id
//
// GET https://api.pagar.me/1/payables/payable_id
func (api *API) Get(id string) (response *pagarme.Response,  payable *pagarme.Payable, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, "/payables/" + id, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Payable)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal payable: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal payable: [GetPayable]" + rerr.Error())
			return
		}
	}

	payable = result
	response = www.Ok()
	return
}

