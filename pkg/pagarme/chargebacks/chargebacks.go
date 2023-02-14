package chargebacks

import (
	"net/http"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

// API is the /1/chargebacks API
type API struct {
	Config *pagarme.Config
}

// New /1/chargebacks API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

// Query -> GET https://api.pagar.me/1/chargebacks
//
//
func (api *API) Query(input *QueryInput) (response *pagarme.Response, chargebacks []pagarme.Chargeback, rerr error) {
	url := "/chargebacks"
	resp, rerr := api.Config.Do(http.MethodGet, url, input.Values(), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.Chargeback, 0)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal chargebacks: " + rerr.Error())
		return
	}

	chargebacks = result
	response = www.Ok(resp)
	return
}
