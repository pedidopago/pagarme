package www

import (
	"encoding/json"
	"net/http"

	"github.com/pedidopago/pagarme/pkg/pagarme"
)

func ExtractError(resp *http.Response) *pagarme.Response {
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	erbundle := &pagarme.Response{}
	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	if err := dec.Decode(erbundle); err != nil {
		return &pagarme.Response{
			Code: 999,
			Errors: []pagarme.Perror{
				pagarme.Perror{
					Message: err.Error(),
				},
			},
		}
	}
	erbundle.Code = pagarme.ResponseCode(resp.StatusCode)
	return erbundle
}

func Ok() *pagarme.Response {
	return &pagarme.Response{
		Code: pagarme.ResponseCodeOk,
	}
}
