package pagarme

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gabstv/sqltypes"
)

// Address TBD
type Address struct {
	Object       sqltypes.NullString `json:"object,omitempty" form:"object,omitempty"`
	ID           int64               `json:"id,omitempty" form:"id,omitempty"`
	Street       sqltypes.NullString `json:"street" form:"street"`
	StreetNumber sqltypes.NullString `json:"street_number" form:"street_number"`
	// CEP. Para endereço brasileiro, deve conter uma numeração de 8 dígitos
	Zipcode sqltypes.NullString `json:"zipcode" form:"zipcode"`
	// Obrigatório. País. Duas letras minúsculas. Deve seguir o padão ISO 3166-1 alpha-2
	Country       CountryCode         `json:"country" form:"country"`
	State         sqltypes.NullString `json:"state" form:"state"`
	City          sqltypes.NullString `json:"city" form:"city"`
	Neighborhood  sqltypes.NullString `json:"neighborhood,omitempty" form:"neighborhood,omitempty"`
	Complementary sqltypes.NullString `json:"complementary,omitempty" form:"complementary,omitempty"`
}

// Document TBD
type Document struct {
	Object sqltypes.NullString `json:"object,omitempty" form:"object,omitempty"`
	ID     sqltypes.NullString `json:"id,omitempty" form:"id,omitempty"`
	Type   DocumentType        `json:"type,omitempty" form:"type,omitempty"`
	Number sqltypes.NullString `json:"number,omitempty" form:"number,omitempty"`
}

// SplitRule TBD
type SplitRule struct {
	Object              sqltypes.NullString `json:"object,omitempty" form:"object,omitempty"`
	ID                  sqltypes.NullString `json:"id,omitempty" form:"id,omitempty"`
	Liable              bool                `json:"liable" form:"liable"`
	ChargeProcessingFee bool                `json:"charge_processing_fee" form:"charge_processing_fee"`
	Percentage          Float64String       `json:"percentage,omitempty" form:"percentage,omitempty"`
	Amount              Int64String         `json:"amount,omitempty" form:"amount,omitempty"`
	ChargeRemainderFee  bool                `json:"charge_remainder_fee" form:"charge_remainder_fee"`
	RecipientID         sqltypes.NullString `json:"recipient_id" form:"recipient_id"`
	DateCreated         time.Time           `json:"date_created" form:"date_created"`
	DateUpdated         time.Time           `json:"date_updated" form:"date_updated"`
	BlockID             sqltypes.NullString `json:"block_id" form:"block_id"`
}

// Customer TODO: godoc
type Customer struct {
	Object       sqltypes.NullString   `json:"object,omitempty" form:"object,omitempty"`
	ID           int64                 `json:"id,omitempty" form:"id,omitempty"`
	ExternalID   sqltypes.NullString   `json:"external_id,omitempty" form:"external_id,omitempty"`
	Type         CustomerType          `json:"type" form:"type"`
	Name         sqltypes.NullString   `json:"name,omitempty" form:"name,omitempty"`
	Email        sqltypes.NullString   `json:"email,omitempty" form:"email,omitempty"`
	Country      CountryCode           `json:"country,omitempty" form:"country,omitempty"`
	Birthday     YYYYMMDD              `json:"birthday,omitempty" form:"birthday,omitempty"`
	Gender       sqltypes.NullString   `json:"gender,omitempty" form:"gender,omitempty"`
	Documents    []*Document           `json:"documents,omitempty" form:"documents,omitempty"`
	PhoneNumbers []sqltypes.NullString `json:"phone_numbers,omitempty" form:"phone_numbers,omitempty"`
}

// Item TODO: godoc
type Item struct {
	ID        sqltypes.NullString `json:"id,omitempty" form:"id,omitempty"`
	Object    sqltypes.NullString `json:"object,omitempty" form:"object,omitempty"`
	Title     sqltypes.NullString `json:"title,omitempty" form:"title,omitempty"`
	UnitPrice int                 `json:"unit_price,omitempty" form:"unit_price,omitempty"`
	Quantity  int                 `json:"quantity,omitempty" form:"quantity,omitempty"`
	Category  sqltypes.NullString `json:"category,omitempty" form:"category,omitempty"`
	Tangible  bool                `json:"tangible,omitempty" form:"tangible,omitempty"`
	Venue     sqltypes.NullString `json:"venue,omitempty" form:"venue,omitempty"`
	Date      sqltypes.NullString `json:"date,omitempty" form:"date,omitempty"`
}

type Billing struct {
	Address *Address            `json:"address,omitempty" form:"address,omitempty"`
	Object  sqltypes.NullString `json:"object,omitempty" form:"object,omitempty"`
	ID      int64               `json:"id,omitempty" form:"id,omitempty"`
	Name    sqltypes.NullString `json:"name,omitempty" form:"name,omitempty"`
}

type Shipping struct {
	Address      *Address            `json:"address,omitempty" form:"address,omitempty"`
	Object       sqltypes.NullString `json:"object,omitempty" form:"object,omitempty"`
	ID           int64               `json:"id,omitempty" form:"id,omitempty"`
	Name         sqltypes.NullString `json:"name,omitempty" form:"name,omitempty"`
	Fee          int                 `json:"fee" form:"fee"`
	DeliveryDate YYYYMMDD            `json:"delivery_date,omitempty" form:"delivery_date,omitempty"`
	Expedited    bool                `json:"expedited,omitempty" form:"expedited,omitempty"`
}

type Card struct {
	Object         sqltypes.NullString `json:"object,omitempty" form:"object,omitempty"`
	ID             sqltypes.NullString `json:"id,omitempty" form:"id,omitempty"`
	DateCreated    time.Time           `json:"date_created,omitempty" form:"date_created,omitempty"`
	DateUpdated    time.Time           `json:"date_updated,omitempty" form:"date_updated,omitempty"`
	Brand          sqltypes.NullString `json:"brand,omitempty" form:"brand,omitempty"`
	HolderName     sqltypes.NullString `json:"holder_name,omitempty" form:"holder_name,omitempty"`
	FirstDigits    sqltypes.NullString `json:"first_digits,omitempty" form:"first_digits,omitempty"`
	LastDigits     sqltypes.NullString `json:"last_digits,omitempty" form:"last_digits,omitempty"`
	Country        sqltypes.NullString `json:"country,omitempty" form:"country,omitempty"`
	Fingerprint    sqltypes.NullString `json:"fingerprint,omitempty" form:"fingerprint,omitempty"`
	Valid          bool                `json:"valid,omitempty" form:"valid,omitempty"`
	ExpirationDate MMYY                `json:"expiration_date,omitempty" form:"expiration_date,omitempty"`
}

type NCard struct {
	CardNumber     sqltypes.NullString `json:"card_number"`
	HolderName     sqltypes.NullString `json:"card_holder_name"`
	ExpirationDate sqltypes.NullString `json:"card_expiration_date"`
	CVV            sqltypes.NullString `json:"card_cvv"`
}

type Response struct {
	Code   ResponseCode        `json:"code,omitempty" form:"code,omitempty"`
	Errors []Perror            `json:"errors,omitempty" form:"errors,omitempty"`
	URL    sqltypes.NullString `json:"url,omitempty" form:"url,omitempty"`
	Method sqltypes.NullString `json:"method,omitempty" form:"method,omitempty"`
}

type Perror struct {
	Message       sqltypes.NullString `json:"message,omitempty" form:"message,omitempty"`
	ParameterName sqltypes.NullString `json:"parameter_name,omitempty" form:"parameter_name,omitempty"`
	Type          sqltypes.NullString `json:"type,omitempty" form:"type,omitempty"`
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

type UnixMS int64

func UnixMSFromTime(t time.Time) UnixMS {
	return UnixMS(t.UnixNano() / 1000000)
}
