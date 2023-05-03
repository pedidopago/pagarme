package settlements

import (
	"net/http"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

// API is the /1/settlements API
type API struct {
	Config *pagarme.Config
}

// New /1/settlements API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

// Query -> GET https://api.pagar.me/1/settlements
//
//
func (api *API) Query(input *QueryInput) (*pagarme.Response, []pagarme.Settlement, error) {
	urlpart := "/settlements"
	resp, err := api.Config.Do(http.MethodGet, urlpart, input.Values(), nil)
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}
	var result struct {
		Settlements []pagarme.Settlement `json:"settlements"`
	}
	//
	if err := www.Unmarshal(api.Config, resp, &result); err != nil {
		api.Config.Logger.Error("could not unmarshal settlements: " + err.Error())
		return nil, nil, err
	}
	return www.Ok(resp), result.Settlements, nil
}
