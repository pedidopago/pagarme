package transfers

import (
	"fmt"
	"github.com/pedidopago/pagarme/internal/pkg/www"
	"github.com/pedidopago/pagarme/pkg/pagarme"
	"net/http"
	"net/url"
	"strconv"
)

// API is the /1/transfers API
type API struct {
	Config *pagarme.Config
}

// New /1/transfers API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config:      cfg,
	}
}

type QueryInput struct {
	Count int
	Page int
	Filter string
	Value string
}

func (in *QueryInput) Export() string {
	vv := url.Values{}
	if in.Count != 0 {
		vv.Set("count", strconv.Itoa(in.Count))
	}
	if in.Page != 0 {
		vv.Set("page", strconv.Itoa(in.Page))
	} else {
		vv.Set("page", "1")
	}
	if in.Filter != "" {
		vv.Set(in.Filter, in.Value)
	}
	return vv.Encode()
}

// Query
//
// GET https://api.pagar.me/1/transfers
func (api *API) Query(input QueryInput) (response *pagarme.Response, transfers []pagarme.Transfer, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/transfers?%s", input.Export()), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.Transfer, 0)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal transfers: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal transfers: [Query]" + rerr.Error())
			return
		}
	}

	transfers = result
	response = www.Ok()
	return
}

type CreateInput struct {
	Amount int `json:"amount"`
	RecipientId string `json:"recipient_id"`
	Metadata map[string]interface{} `json:"metadata"`
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

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal transfer: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal transfer: [Create]" + rerr.Error())
			return
		}
	}

	transfer = result
	response = www.Ok()
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

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal transfer: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal transfer: [Cancel]" + rerr.Error())
			return
		}
	}

	transfer = result
	response = www.Ok()
	return
}
