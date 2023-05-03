package bankaccounts

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
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
	Count          int
	Page           int
	Id             *string
	BankCode       *string
	Agencia        *string
	AgenciaDv      *string
	Conta          *string
	ContaDv        *string
	DocumentNumber *string
	LegalName      *string
	DateCreated    *time.Time
	Extra          map[string]string
}

func (in *QueryInput) Values() url.Values {
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
		vv.Set("date_created", strconv.FormatInt(v.Unix()*1000, 10))
	}
	for k, v := range in.Extra {
		vv.Set(k, v)
	}
	return vv
}

func (in *QueryInput) Export() string {
	return in.Values().Encode()
}

func (api *API) Query(input QueryInput) (response *pagarme.Response, accounts []pagarme.BankAccount, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, "/bank_accounts", input.Values(), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.BankAccount, 0)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal bank accounts: " + rerr.Error())
		return
	}

	accounts = result
	response = www.Ok(resp)
	return
}

func (api *API) Get(id int) (response *pagarme.Response, account *pagarme.BankAccount, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/bank_accounts/%d", id), nil, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.BankAccount)

	if rerr = www.Unmarshal(api.Config, resp, result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal bank account: " + rerr.Error())
		return
	}

	account = result
	response = www.Ok(resp)
	return
}

func (api *API) Create(createInput pagarme.BankAccount) (response *pagarme.Response, account *pagarme.BankAccount, rerr error) {
	resp, rerr := api.Config.Do(http.MethodPost, "/bank_accounts", nil, www.JSONReader(createInput))
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.BankAccount)

	if rerr = www.Unmarshal(api.Config, resp, result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal bank account: " + rerr.Error())
		return
	}

	account = result
	response = www.Ok(resp)
	return
}
