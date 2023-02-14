package transactions

import (
	"fmt"
	"net/http"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

// QueryOperations -> GET https://api.pagar.me/1/transactions/:tid/operations
//
//
func (api *API) QueryOperations(tid string) (response *pagarme.Response, operations []pagarme.TransactionOperation, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/transactions/%s/operations", tid), nil, nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.TransactionOperation, 0)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal refunds " + rerr.Error())
		return
	}

	operations = result
	response = www.Ok(resp)
	return
}
