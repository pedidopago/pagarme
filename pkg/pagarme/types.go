package pagarme

import (
	"bytes"
	"fmt"
	"time"
)

type Address struct {
	Object       string `json:"object,omitempty"`
	ID           int64  `json:"id,omitempty"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
	// CEP. Para endereço brasileiro, deve conter uma numeração de 8 dígitos
	Zipcode string `json:"zipcode"`
	// Obrigatório. País. Duas letras minúsculas. Deve seguir o padão ISO 3166-1 alpha-2
	Country       CountryCode `json:"country"`
	State         string      `json:"state"`
	City          string      `json:"city"`
	Neighborhood  string      `json:"neighborhood,omitempty"`
	Complementary string      `json:"complementary,omitempty"`
}

type Document struct {
	Object string       `json:"object,omitempty"`
	ID     string       `json:"id,omitempty"`
	Type   DocumentType `json:"type,omitempty"`
	Number string       `json:"number,omitempty"`
}

type SplitRule struct {
	Liable              bool   `json:"liable"`
	ChargeProcessingFee bool   `json:"charge_processing_fee"`
	Percentage          string `json:"percentage,omitempty"`
	Amount              string `json:"amount,omitempty"`
	ChargeRemainderFee  bool   `json:"charge_remainder_fee"`
	RecipientID         string `json:"recipient_id"`
}

type Customer struct {
	Object       string       `json:"object,omitempty"`
	ID           int64        `json:"id,omitempty"`
	ExternalID   string       `json:"external_id,omitempty"`
	Type         CustomerType `json:"type"`
	Name         string       `json:"name,omitempty"`
	Email        string       `json:"email,omitempty"`
	Country      CountryCode  `json:"country,omitempty"`
	Birthday     YYYYMMDD     `json:"birthday,omitempty"`
	Gender       string       `json:"gender,omitempty"`
	Documents    []*Document  `json:"documents,omitempty"`
	PhoneNumbers []string     `json:"phone_numbers,omitempty"`
}

type Item struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	UnitPrice int    `json:"unit_price,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
	Tangible  bool   `json:"tangible,omitempty"`
}

type Billing struct {
	Address *Address `json:"address,omitempty"`
	Object  string   `json:"object,omitempty"`
	ID      int64    `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
}

type Shipping struct {
	Address      *Address `json:"address,omitempty"`
	Object       string   `json:"object,omitempty"`
	ID           int64    `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Fee          int      `json:"fee"`
	DeliveryDate YYYYMMDD `json:"delivery_date,omitempty"`
	Expedited    bool     `json:"expedited,omitempty"`
}

type Card struct {
	Object         string    `json:"object,omitempty"`
	ID             string    `json:"id,omitempty"`
	DateCreated    time.Time `json:"date_created,omitempty"`
	DateUpdated    time.Time `json:"date_updated,omitempty"`
	Brand          string    `json:"brand,omitempty"`
	HolderName     string    `json:"holder_name,omitempty"`
	FirstDigits    string    `json:"first_digits,omitempty"`
	LastDigits     string    `json:"last_digits,omitempty"`
	Country        string    `json:"country,omitempty"`
	Fingerprint    string    `json:"fingerprint,omitempty"`
	Valid          bool      `json:"valid,omitempty"`
	ExpirationDate MMYY      `json:"expiration_date,omitempty"`
}

type Response struct {
	Code   ResponseCode `json:"code,omitempty"`
	Errors []Perror     `json:"errors,omitempty"`
	URL    string       `json:"url,omitempty"`
	Method string       `json:"method,omitempty"`
}

type Perror struct {
	Message       string `json:"message,omitempty"`
	ParameterName string `json:"parameter_name,omitempty"`
	Type          string `json:"type,omitempty"`
}

func (r *Response) String() string {
	erbuf := new(bytes.Buffer)
	erbuf.WriteString("[")
	for k, v := range r.Errors {
		if k != 0 {
			erbuf.WriteString(",")
		}
		erbuf.WriteString(fmt.Sprintf("{type:%v,parameter_name:%v,message:%v}", v.Type, v.ParameterName, v.Message))
	}
	erbuf.WriteString("]")
	return fmt.Sprintf("code: %v, errors: %v", r.Code, erbuf.String())
}
