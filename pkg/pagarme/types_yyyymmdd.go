package pagarme

import (
	"time"
)

type YYYYMMDD string

func YYYYMMDDFromTime(t time.Time) YYYYMMDD {
	return YYYYMMDD(t.Format("2006-01-02"))
}

type MMYY string
