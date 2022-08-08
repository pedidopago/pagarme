package anticipations

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

// API is the /1/recipients/:recipient_id/bulk_anticipations API
type API struct {
	Config      *pagarme.Config
	RecipientID string
}

// New /1/recipients/:recipient_id/bulk_anticipations API
func New(cfg *pagarme.Config, recipientID string) *API {
	return &API{
		Config:      cfg,
		RecipientID: recipientID,
	}
}

// NewAnticipation creates a new bulk_anticipation
//
// POST https://api.pagar.me/1/recipients/recipient_id/bulk_anticipations
func (api *API) NewAnticipation(req *pagarme.CreateAnticipation) (response *pagarme.Response, anticipation *pagarme.Anticipation, rerr error) {
	resp, rerr := api.Config.Do(http.MethodPost, fmt.Sprintf("/recipients/%s/bulk_anticipations", api.RecipientID), www.JSONReader(req))
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Anticipation)

	if rerr = www.Unmarshal(api.Config, resp, result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal anticipation: " + rerr.Error())
		return
	}

	anticipation = result
	response = www.Ok(resp)
	return
}

type GetLimitsInput struct {
	Timeframe   pagarme.AnticipationTimeframe
	PaymentDate int64
}

func (in *GetLimitsInput) Export() string {
	vv := url.Values{}
	vv.Set("timeframe", string(in.Timeframe))
	vv.Set("payment_date", strconv.FormatInt(in.PaymentDate, 10))
	return vv.Encode()
}

// GetLimits returns the limits a new anticipation can be created with
//
// GET https://api.pagar.me/1/recipients/recipient_id/bulk_anticipations/limits
func (api *API) GetLimits(input GetLimitsInput) (response *pagarme.Response, limits *pagarme.Limits, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/recipients/%s/bulk_anticipations/limits?%s", api.RecipientID, input.Export()), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Limits)

	if rerr = www.Unmarshal(api.Config, resp, result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal limits: " + rerr.Error())
		return
	}

	limits = result
	response = www.Ok(resp)
	return
}

// ConfirmNewAnticipation confirms the created anticipation
//
// POST https://api.pagar.me/1/recipients/recipient_id/bulk_anticipations/bulk_anticipation_id/confirm
func (api *API) ConfirmNewAnticipation(bulkAnticipationId string) (response *pagarme.Response, anticipation *pagarme.Anticipation, rerr error) {
	resp, rerr := api.Config.Do(http.MethodPost, fmt.Sprintf("/recipients/%s/bulk_anticipations/%s/confirm", api.RecipientID, bulkAnticipationId), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Anticipation)

	if rerr = www.Unmarshal(api.Config, resp, result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal anticipation: " + rerr.Error())
		return
	}

	anticipation = result
	response = www.Ok(resp)
	return
}

// CancelPendingAnticipation
//
// POST https://api.pagar.me/1/recipients/recipient_id/bulk_anticipations/bulk_anticipation_id/cancel
func (api *API) CancelPendingAnticipation(bulkAnticipationId string) (response *pagarme.Response, anticipation *pagarme.Anticipation, rerr error) {
	resp, rerr := api.Config.Do(http.MethodPost, fmt.Sprintf("/recipients/%s/bulk_anticipations/%s/cancel", api.RecipientID, bulkAnticipationId), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := new(pagarme.Anticipation)

	if rerr = www.Unmarshal(api.Config, resp, result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal anticipation: " + rerr.Error())
		return
	}

	anticipation = result
	response = www.Ok(resp)
	return
}

type QueryInput struct {
	b      *pagarme.QueryBuilder
	Count  int
	Page   int
	Filter string
	Value  string
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

// Delete
//
// DELETE https://api.pagar.me/1/recipients/recipient_id/bulk_anticipations/bulk_anticipation_id
func (api *API) Delete(bulkAnticipationId string) (response *pagarme.Response, rerr error) {
	resp, rerr := api.Config.Do(http.MethodDelete, fmt.Sprintf("/recipients/%s/bulk_anticipations/%s", api.RecipientID, bulkAnticipationId), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}

	response = www.Ok(resp)
	return
}

// Query
//
// GET https://api.pagar.me/1/recipients/recipient_id/bulk_anticipations/
func (api *API) Query(input QueryInput) (response *pagarme.Response, anticipations []pagarme.Anticipation, rerr error) {
	resp, rerr := api.Config.Do(http.MethodGet, fmt.Sprintf("/recipients/%s/bulk_anticipations/?%s", api.RecipientID, input.Export()), nil)
	if rerr != nil {
		return
	}
	if response = www.ExtractError(resp); response != nil {
		return
	}
	result := make([]pagarme.Anticipation, 0)

	if rerr = www.Unmarshal(api.Config, resp, &result); rerr != nil {
		api.Config.Logger.Error("could not unmarshal anticipations: " + rerr.Error())
		return
	}

	anticipations = result
	response = www.Ok(resp)
	return
}
