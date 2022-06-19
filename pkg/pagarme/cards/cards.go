package cards

import (
	"net/http"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

// API is the /1/cards API
type API struct {
	Config *pagarme.Config
}

// New /1/cards API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

// NewCard consume a pagarme API to create a new card and return your informations
func (api *API) NewCard(cr *pagarme.NCard) (*pagarme.Response, *pagarme.Card, error) {
	resp, err := api.Config.Do(http.MethodPost, "/cards", www.JSONReader(cr))
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}
	result := &pagarme.Card{}

	if err := www.Unmarshal(api.Config, resp, result); err != nil {
		api.Config.Logger.Error("could not unmarshal transaction: " + err.Error())
		return nil, nil, err
	}

	return www.Ok(resp), result, nil
}
