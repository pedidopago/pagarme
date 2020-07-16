package recipients

import (
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
