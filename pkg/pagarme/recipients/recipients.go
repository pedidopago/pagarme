package recipients

import (
	"fmt"
	"net/http"

	"github.com/pedidopago/pagarme/internal/pkg/www"
	"github.com/pedidopago/pagarme/pkg/pagarme"
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
	resp, err := api.Config.Do(http.MethodPost, "/recipients", www.JSONReader(recipient))
	if err != nil {
		return nil, nil, err
	}
	if werr := www.ExtractError(resp); werr != nil {
		return werr, nil, nil
	}
	result := &pagarme.Recipient{}

	if api.Config.Trace {
		if err := www.UnmarshalTrace(api.Config.Logger, resp, result); err != nil {
			api.Config.Logger.Error("could not unmarshal transaction: " + err.Error())
			return nil, nil, err
		}
	} else {
		if err := www.Unmarshal(resp, result); err != nil {
			api.Config.Logger.Error("could not unmarshal transaction [Put]: " + err.Error())
			return nil, nil, err
		}
	}
	return www.Ok(), result, nil
}

func (api *API) GetRecipient(id string) (response *pagarme.Response, recipient *pagarme.Recipient, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/recipients/%s", id), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Recipient)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal recipient: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal recipient: [GetRecipient]" + rerr.Error())
			return
		}
	}

	recipient = result
	response = www.Ok()
	return
}
