package pagarme

import (
	"testing"
	"time"
)

func TestMMYY(t *testing.T) {
	v := MMYY("1233")
	if v.Month() != time.Month(12) {
		t.Fail()
	}
	if v.Year() != 2033 {
		t.Fail()
	}
}
