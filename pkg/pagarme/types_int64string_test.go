package pagarme

import (
	"testing"
)

func TestMarshalI64(t *testing.T) {
	v := Int64String(180)
	strb, _ := v.MarshalJSON()
	if string(strb) != "\"180\"" {
		t.Fatal(string(strb))
	}
	v = Int64String(1800)
	strb, _ = v.MarshalJSON()
	if string(strb) != "\"1800\"" {
		t.Fatal(string(strb))
	}
	v = Int64String(18000)
	strb, _ = v.MarshalJSON()
	if string(strb) != "\"18000\"" {
		t.Fatal(string(strb))
	}
	v = Int64String(18000)
	strb, _ = v.MarshalJSON()
	if string(strb) != "\"18000\"" {
		t.Fatal(string(strb))
	}
	v = Int64String(180000)
	strb, _ = v.MarshalJSON()
	if string(strb) != "\"180000\"" {
		t.Fatal(string(strb))
	}
	v = Int64String(1800000)
	strb, _ = v.MarshalJSON()
	if string(strb) != "\"1800000\"" {
		t.Fatal(string(strb))
	}
}

func TestUnmarshalI64(t *testing.T) {
	vv := Int64String(0)
	v := &vv
	if err := v.UnmarshalJSON([]byte("\"360\"")); err != nil {
		t.Fatal(err.Error())
	}
	if *v != 360 {
		t.Fatal(*v)
	}
	if err := v.UnmarshalJSON([]byte("362")); err != nil {
		t.Fatal(err.Error())
	}
	if *v != 362 {
		t.Fatal(*v)
	}
}
