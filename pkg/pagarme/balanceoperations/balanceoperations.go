package balanceoperations

import (
	"fmt"
	"github.com/pedidopago/pagarme/internal/pkg/www"
	"github.com/pedidopago/pagarme/pkg/pagarme"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// API is the /1/balance/operations API
type API struct {
	Config *pagarme.Config
}

// New /1/balance/operations API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

func (api *API) Get(id string) (response *pagarme.Response, operation *pagarme.BalanceOperation, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, "/balance/operations/" + id, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.BalanceOperation)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal balance operations: " + rerr.Error())
			return
		} else {
			if rerr = www.Unmarshal(resp, result); rerr != nil {
				api.Config.Logger.Error("could not unmarshal balance operations: [Get]" + rerr.Error())
				return
			}
		}
	}

	operation = result
	response = www.Ok()
	return
}

type QueryInput struct {
	Count     int
	Page      int
	Status    *pagarme.BalanceOperationStatus
	StartDate *time.Time
	EndDate   *time.Time
	Extra map[string]string
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
	if in.Status != nil {
		vv.Set("status", string(*in.Status))
	}
	if in.StartDate != nil {
		vv.Set("start_date", strconv.FormatInt(in.StartDate.Unix() * 1000, 10))
	}
	if in.EndDate != nil {
		vv.Set("end_date", strconv.FormatInt(in.StartDate.Unix() * 1000, 10))
	}
	for k, v := range in.Extra {
		vv.Set(k, v)
	}
	return vv.Encode()
}

func (api *API) Query(input QueryInput) (response *pagarme.Response, operations []pagarme.BalanceOperation, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/balance/operations?%s", input.Export()), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.BalanceOperation, 0)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal balance operations: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal balance operations: [Query]" + rerr.Error())
			return
		}
	}

	operations = result
	response = www.Ok()
	return
}
