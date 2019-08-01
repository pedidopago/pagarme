package pagarme

import (
	"bytes"
	"fmt"
	"time"
)

// Address TBD
type Address struct {
	Object       string `json:"object,omitempty" form:"object,omitempty"`
	ID           int64  `json:"id,omitempty" form:"id,omitempty"`
	Street       string `json:"street" form:"street"`
	StreetNumber string `json:"street_number" form:"street_number"`
	// CEP. Para endereço brasileiro, deve conter uma numeração de 8 dígitos
	Zipcode string `json:"zipcode" form:"zipcode"`
	// Obrigatório. País. Duas letras minúsculas. Deve seguir o padão ISO 3166-1 alpha-2
	Country       CountryCode `json:"country" form:"country"`
	State         string      `json:"state" form:"state"`
	City          string      `json:"city" form:"city"`
	Neighborhood  string      `json:"neighborhood,omitempty" form:"neighborhood,omitempty"`
	Complementary string      `json:"complementary,omitempty" form:"complementary,omitempty"`
}

// Document TBD
type Document struct {
	Object string       `json:"object,omitempty" form:"object,omitempty"`
	ID     string       `json:"id,omitempty" form:"id,omitempty"`
	Type   DocumentType `json:"type,omitempty" form:"type,omitempty"`
	Number string       `json:"number,omitempty" form:"number,omitempty"`
}

// SplitRule TBD
type SplitRule struct {
	Object              string        `json:"object,omitempty" form:"object,omitempty"`
	ID                  string        `json:"id,omitempty" form:"id,omitempty"`
	Liable              bool          `json:"liable" form:"liable"`
	ChargeProcessingFee bool          `json:"charge_processing_fee" form:"charge_processing_fee"`
	Percentage          Float64String `json:"percentage,omitempty" form:"percentage,omitempty"`
	Amount              string        `json:"amount,omitempty" form:"amount,omitempty"`
	ChargeRemainderFee  bool          `json:"charge_remainder_fee" form:"charge_remainder_fee"`
	RecipientID         string        `json:"recipient_id" form:"recipient_id"`
}

// Customer TODO: godoc
type Customer struct {
	Object       string       `json:"object,omitempty" form:"object,omitempty"`
	ID           int64        `json:"id,omitempty" form:"id,omitempty"`
	ExternalID   string       `json:"external_id,omitempty" form:"external_id,omitempty"`
	Type         CustomerType `json:"type" form:"type"`
	Name         string       `json:"name,omitempty" form:"name,omitempty"`
	Email        string       `json:"email,omitempty" form:"email,omitempty"`
	Country      CountryCode  `json:"country,omitempty" form:"country,omitempty"`
	Birthday     YYYYMMDD     `json:"birthday,omitempty" form:"birthday,omitempty"`
	Gender       string       `json:"gender,omitempty" form:"gender,omitempty"`
	Documents    []*Document  `json:"documents,omitempty" form:"documents,omitempty"`
	PhoneNumbers []string     `json:"phone_numbers,omitempty" form:"phone_numbers,omitempty"`
}

// Item TODO: godoc
type Item struct {
	ID        string `json:"id,omitempty" form:"id,omitempty"`
	Title     string `json:"title,omitempty" form:"title,omitempty"`
	UnitPrice int    `json:"unit_price,omitempty" form:"unit_price,omitempty"`
	Quantity  int    `json:"quantity,omitempty" form:"quantity,omitempty"`
	Tangible  bool   `json:"tangible,omitempty" form:"tangible,omitempty"`
}

type Billing struct {
	Address *Address `json:"address,omitempty" form:"address,omitempty"`
	Object  string   `json:"object,omitempty" form:"object,omitempty"`
	ID      int64    `json:"id,omitempty" form:"id,omitempty"`
	Name    string   `json:"name,omitempty" form:"name,omitempty"`
}

type Shipping struct {
	Address      *Address `json:"address,omitempty" form:"address,omitempty"`
	Object       string   `json:"object,omitempty" form:"object,omitempty"`
	ID           int64    `json:"id,omitempty" form:"id,omitempty"`
	Name         string   `json:"name,omitempty" form:"name,omitempty"`
	Fee          int      `json:"fee" form:"fee"`
	DeliveryDate YYYYMMDD `json:"delivery_date,omitempty" form:"delivery_date,omitempty"`
	Expedited    bool     `json:"expedited,omitempty" form:"expedited,omitempty"`
}

type Card struct {
	Object         string    `json:"object,omitempty" form:"object,omitempty"`
	ID             string    `json:"id,omitempty" form:"id,omitempty"`
	DateCreated    time.Time `json:"date_created,omitempty" form:"date_created,omitempty"`
	DateUpdated    time.Time `json:"date_updated,omitempty" form:"date_updated,omitempty"`
	Brand          string    `json:"brand,omitempty" form:"brand,omitempty"`
	HolderName     string    `json:"holder_name,omitempty" form:"holder_name,omitempty"`
	FirstDigits    string    `json:"first_digits,omitempty" form:"first_digits,omitempty"`
	LastDigits     string    `json:"last_digits,omitempty" form:"last_digits,omitempty"`
	Country        string    `json:"country,omitempty" form:"country,omitempty"`
	Fingerprint    string    `json:"fingerprint,omitempty" form:"fingerprint,omitempty"`
	Valid          bool      `json:"valid,omitempty" form:"valid,omitempty"`
	ExpirationDate MMYY      `json:"expiration_date,omitempty" form:"expiration_date,omitempty"`
}

type Response struct {
	Code   ResponseCode `json:"code,omitempty" form:"code,omitempty"`
	Errors []Perror     `json:"errors,omitempty" form:"errors,omitempty"`
	URL    string       `json:"url,omitempty" form:"url,omitempty"`
	Method string       `json:"method,omitempty" form:"method,omitempty"`
}

type Perror struct {
	Message       string `json:"message,omitempty" form:"message,omitempty"`
	ParameterName string `json:"parameter_name,omitempty" form:"parameter_name,omitempty"`
	Type          string `json:"type,omitempty" form:"type,omitempty"`
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
