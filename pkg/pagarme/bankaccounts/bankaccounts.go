package bankaccounts

import (
	"fmt"
	"github.com/pedidopago/pagarme/internal/pkg/www"
	"github.com/pedidopago/pagarme/pkg/pagarme"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// API is the /1/bank_accounts API
type API struct {
	Config *pagarme.Config
}

// New /1/bank_accounts API
func New(cfg *pagarme.Config) *API {
	return &API{
		Config: cfg,
	}
}

type QueryInput struct {
	Count     int
	Page      int
	Id        *string
	BankCode  *string
	Agencia *string
	AgenciaDv *string
	Conta *string
	ContaDv *string
	DocumentNumber *string
	LegalName *string
	DateCreated *time.Time
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
	if v := in.Id; v != nil {
		vv.Set("id", *v)
	}
	if v := in.BankCode; v != nil {
		vv.Set("bank_code", *v)
	}
	if v := in.Agencia; v != nil {
		vv.Set("agencia", *v)
	}
	if v := in.AgenciaDv; v != nil {
		vv.Set("agencia_dv", *v)
	}
	if v := in.Conta; v != nil {
		vv.Set("conta", *v)
	}
	if v := in.ContaDv; v != nil {
		vv.Set("conta_dv", *v)
	}
	if v := in.DocumentNumber; v != nil {
		vv.Set("document_number", *v)
	}
	if v := in.LegalName; v != nil {
		vv.Set("legal_name", *v)
	}
	if v := in.DateCreated; v != nil {
		vv.Set("date_created", strconv.FormatInt(v.Unix() * 1000, 10))
	}
	for k, v := range in.Extra {
		vv.Set(k, v)
	}
	return vv.Encode()
}

func (api *API) Query(input QueryInput) (response *pagarme.Response, accounts []pagarme.BankAccount, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/bank_accounts?%s", input.Export()), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.BankAccount, 0)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal bank accounts: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, &result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal bank accounts: [Query]" + rerr.Error())
			return
		}
	}

	accounts = result
	response = www.Ok()
	return
}

func (api *API) Get(id int) (response *pagarme.Response, account *pagarme.BankAccount, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/bank_accounts/%d", id), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.BankAccount)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal bank account: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal bank account: [Get]" + rerr.Error())
			return
		}
	}

	account = result
	response = www.Ok()
	return
}

func (api *API) Create(createInput pagarme.BankAccount) (response *pagarme.Response, account *pagarme.BankAccount, rerr error) {
	resp, rerr := api.Config.Do(http.MethodPost, "/bank_accounts", www.JSONReader(createInput))
	if rerr != nil {
		return
	}
	result := new(pagarme.BankAccount)

	if api.Config.Trace {
		if rerr = www.UnmarshalTrace(api.Config.Logger, resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal bank account: " + rerr.Error())
			return
		}
	} else {
		if rerr = www.Unmarshal(resp, result); rerr != nil {
			api.Config.Logger.Error("could not unmarshal bank account: [Create]" + rerr.Error())
			return
		}
	}

	account = result
	response = www.Ok()
	return
}