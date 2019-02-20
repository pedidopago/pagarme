package pagarme

import (
	"strconv"
	"time"
)

type YYYYMMDD string

func YYYYMMDDFromTime(t time.Time) YYYYMMDD {
	return YYYYMMDD(t.Format("2006-01-02"))
}

// MMYY represents a two digit MM (month) and YY (last two digits of the year) combined
type MMYY string

// Month returns the month
func (s MMYY) Month() time.Month {
	if len(s) != 4 {
		return time.Month(0)
	}
	v, _ := strconv.Atoi(string(s[:2]))
	return time.Month(v)
}

// Year returns the year
func (s MMYY) Year(prefix ...string) int {
	if len(s) != 4 {
		return 0
	}
	part := string(s[2:])
	if len(prefix) > 0 {
		v, _ := strconv.Atoi(prefix[0] + part)
		return v
	}
	v, _ := strconv.Atoi("20" + part)
	return v
}
