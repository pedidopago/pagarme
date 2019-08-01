package pagarme

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
)

type Int64String int64

// Scan implements the Scanner interface.
func (n *Int64String) Scan(value interface{}) error {
	if value == nil {
		*n = 0
		return nil
	}

	n64 := sql.NullInt64{}
	err := n64.Scan(value)
	if err != nil {
		return err
	}

	*n = Int64String(n64.Int64)
	return nil
}

// Value implements the driver Valuer interface.
func (n Int64String) Value() (driver.Value, error) {
	if n == 0 {
		return nil, nil
	}
	return float64(n), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (n *Int64String) UnmarshalJSON(v []byte) error {
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
		i64, err := strconv.ParseInt(strof[1:len(strof)-1], 10, 64)
		if err != nil {
			return err
		}
		*n = Int64String(i64)
		return nil
	}
	i64, err := strconv.ParseInt(strof, 10, 64)
	if err != nil {
		return err
	}
	*n = Int64String(i64)
	return nil
}

// MarshalJSON implements json.Marshaler
func (n *Int64String) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", float64(*n))), nil
}
