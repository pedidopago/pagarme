package pagarme

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
)

type Float64String float64

// Scan implements the Scanner interface.
func (n *Float64String) Scan(value interface{}) error {
	if value == nil {
		*n = 0
		return nil
	}

	nf64 := sql.NullFloat64{}
	err := nf64.Scan(value)
	if err != nil {
		return err
	}

	*n = Float64String(nf64.Float64)
	return nil
}

// Value implements the driver Valuer interface.
func (n Float64String) Value() (driver.Value, error) {
	if n == 0 {
		return nil, nil
	}
	return float64(n), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (n *Float64String) UnmarshalJSON(v []byte) error {
	if v == nil {
		return nil
	}
	if len(v) == 0 {
		return nil
	}
	if string(v) == "null" {
		return nil
	}
	strof := string(v)
	if strof[0] == '"' && strof[len(strof)-1] == '"' {
		f64, err := strconv.ParseFloat(strof[1:len(strof)-1], 64)
		if err != nil {
			return err
		}
		*n = Float64String(f64)
		return nil
	}
	f64, err := strconv.ParseFloat(strof, 64)
	if err != nil {
		return err
	}
	*n = Float64String(f64)
	return nil
}

// MarshalJSON implements json.Marshaler
func (n *Float64String) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", float64(*n))), nil
}
