package pagarme

import (
	"testing"
)

func TestMarshalF64(t *testing.T) {
	v := Float64String(180)
	strb, _ := v.MarshalJSON()
	if string(strb) != "\"180\"" {
		t.Fatal(string(strb))
	}
}

func TestUnmarshalF64(t *testing.T) {
	vv := Float64String(0)
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
