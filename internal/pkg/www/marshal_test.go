package www

import (
	"io/ioutil"
	"testing"
)

func TestJSONReader(t *testing.T) {
	var testobj = struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{"john", 23}
	rdr := JSONReader(testobj)
	bs, err := ioutil.ReadAll(rdr)
	if err != nil {
		t.Fatal(err.Error())
	}
	if string(bs) != "{\"name\":\"john\",\"age\":23}\n" {
		t.Fatal(string(bs), "!=", "{\"name\":\"john\",\"age\":23}")
	}
}
