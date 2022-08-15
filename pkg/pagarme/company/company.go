package company

import (
	"net/http"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

// API is the /1/company API
type API struct {
	Config *pagarme.Config
}

// New /1/company API
func New(cfg *pagarme.Config) *API {
	cfg.SkipUnmarshal = true
	return &API{
		Config: cfg,
	}
}

// Get -> GET https://api.pagar.me/1/company
//
//
func (api *API) Get() (response *pagarme.Response, rerr error) {
	url := "/company"
	resp, rerr := api.Config.Do(http.MethodGet, url, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}

	response = www.Ok(resp)
	return
}
