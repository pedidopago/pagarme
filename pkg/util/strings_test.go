package util

import "testing"

type padtest struct {
	Source   string
	Pad      string
	Length   int
	Left     bool
	Expected string
}

func TestPad(t *testing.T) {
	var tests = []padtest{
		padtest{"éã🧐", "😎🎻😎", 6, true, "😎🎻😎éã🧐"},
		padtest{"100", "0", 4, true, "0100"},
		padtest{"4994", "0", 8, false, "49940000"},
	}
	for _, v := range tests {
		var res string
		if v.Left {
			res = Lpad(v.Source, v.Pad, v.Length)
		} else {
			res = Rpad(v.Source, v.Pad, v.Length)
		}
		if res != v.Expected {
			t.Fatal("expected", v.Expected, "actual", res)
		}
	}
}
