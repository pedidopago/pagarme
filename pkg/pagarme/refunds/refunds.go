package refunds

import (
	"net/http"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

// API is the /1/payables API
type API struct {
	Config *pagarme.Config
}

// New /1/refunds API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

// Query -> GET https://api.pagar.me/1/refunds
//
//
func (api *API) Query(input *QueryInput) (response *pagarme.Response, refunds []pagarme.Refund, rerr error) {
	url := "/refunds"
	if input != nil {
		url += "?" + input.Build()
	}
	resp, rerr := api.Config.Do(http.MethodGet, url, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.Refund, 0)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal refunds " + rerr.Error())
		return
	}

	refunds = result
	response = www.Ok(resp)
	return
}
