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
	println(urlpart)
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
