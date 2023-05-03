package recipients

import (
	"fmt"
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

// NewRecipient consume a pagarme API to create a new recipient and return your informations
func (api *API) NewRecipient(recipient *pagarme.CreateRecipient) (*pagarme.Response, *pagarme.Recipient, error) {
	resp, err := api.Config.Do(http.MethodPost, "/recipients", nil, www.JSONReader(recipient))
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}
	result := &pagarme.Recipient{}

	if err := www.Unmarshal(api.Config, resp, result); err != nil {
		api.Config.Logger.Error("could not unmarshal recipient: " + err.Error())
		return nil, nil, err
	}

	return www.Ok(resp), result, nil
}

// UpdateRecipient consume a pagarme API to update a recipient and return its information
func (api *API) UpdateRecipient(recipientId string, updateRecipient *pagarme.UpdateRecipient) (response *pagarme.Response, recipient *pagarme.Recipient, rerr error) {
	resp, rerr := api.Config.Do(http.MethodPut, fmt.Sprintf("/recipients/%s", recipientId), nil, www.JSONReader(updateRecipient))
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Recipient)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal recipient: " + rerr.Error())
		return
	}

	recipient = result
	response = www.Ok(resp)
	return
}

func (api *API) GetRecipient(id string) (response *pagarme.Response, recipient *pagarme.Recipient, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/recipients/%s", id), nil, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Recipient)

	if rerr = www.Unmarshal(api.Config, resp, result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal recipient: " + rerr.Error())
		return
	}

	recipient = result
	response = www.Ok(resp)
	return
}

// GetBalance
//
// https://api.pagar.me/1/recipients/recipient_id/balance
func (api *API) GetBalance(recipientId string) (response *pagarme.Response, balance *pagarme.Balance, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/recipients/%s/balance", recipientId), nil, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Balance)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal balance: " + rerr.Error())
		return
	}

	balance = result
	response = www.Ok(resp)
	return
}

// GetBalanceOperations
//
// https://api.pagar.me/1/recipients/recipient_id/balance/operations
func (api *API) GetBalanceOperations(recipientId string, params *pagarme.QueryBuilder) (response *pagarme.Response, operations []pagarme.BalanceOperation, rerr error) {
	url := fmt.Sprintf("/recipients/%s/balance/operations", recipientId)
	resp, rerr := api.Config.Do(http.MethodGet, url, params.Values(), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.BalanceOperation, 0)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal balance operations: " + rerr.Error())
		return
	}

	operations = result
	response = www.Ok(resp)
	return
}

// DownloadBalanceOperations
//
// https://api.pagar.me/1/recipients/recipient_id/balance/operations.ext
func (api *API) DownloadBalanceOperations(recipientId string, ext string, params *pagarme.QueryBuilder) (response *http.Response, rerr error) {
	url := fmt.Sprintf("/recipients/%s/balance/operations.%s", recipientId, ext)
	response, rerr = api.Config.Do(http.MethodGet, url, params.Values(), nil)
	return
}
