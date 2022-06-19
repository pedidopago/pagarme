package transfers

import (
	"fmt"
	"net/http"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

// API is the /1/transfers API
type API struct {
	Config *pagarme.Config
}

// New /1/transfers API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

// Query
//
// GET https://api.pagar.me/1/transfers
func (api *API) Query(params *pagarme.QueryBuilder) (response *pagarme.Response, transfers []pagarme.Transfer, rerr error) {
	url := "/transfers"
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
	result := make([]pagarme.Transfer, 0)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal transfers: " + rerr.Error())
		return
	}

	transfers = result
	response = www.Ok(resp)
	return
}

// Get
//
// https://api.pagar.me/1/transfers/transfer_id
func (api *API) Get(id int) (response *pagarme.Response, transfer *pagarme.Transfer, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/transfers/%d", id), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Transfer)

	if rerr = www.Unmarshal(api.Config, resp, result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal transfer: " + rerr.Error())
		return
	}

	transfer = result
	response = www.Ok(resp)
	return
}

type CreateInput struct {
	Amount      int                    `json:"amount"`
	RecipientId string                 `json:"recipient_id"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// Create
//
// POST https://api.pagar.me/1/transfers
func (api *API) Create(in CreateInput) (response *pagarme.Response, transfer *pagarme.Transfer, rerr error) {
	resp, rerr := api.Config.Do(http.MethodPost, "/transfers", www.JSONReader(in))
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Transfer)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal transfer: " + rerr.Error())
		return
	}

	transfer = result
	response = www.Ok(resp)
	return
}

func (api *API) Cancel(id string) (response *pagarme.Response, transfer *pagarme.Transfer, rerr error) {
	resp, rerr := api.Config.Do(http.MethodPost, fmt.Sprintf("/transfers/%s/cancel", id), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Transfer)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal transfer: " + rerr.Error())
		return
	}

	transfer = result
	response = www.Ok(resp)
	return
}
