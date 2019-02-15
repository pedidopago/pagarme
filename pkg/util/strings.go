package util

import (
	"bytes"
	"unicode/utf8"
)

// Lpad pads text to the left
func Lpad(str, pad string, length int) string {
	return padstr(str, pad, length, true)
}

// Rpad pads text to the right
func Rpad(str, pad string, length int) string {
	return padstr(str, pad, length, false)
}

func padstr(str, pad string, length int, dirleft bool) string {
	lenstr := utf8.RuneCountInString(str)
	if lenstr >= length {
		return str
	}
	lenpad := utf8.RuneCountInString(pad)
	if lenpad == 0 {
		return str
	}
	buf := bytes.Buffer{}
	if !dirleft {
		buf.WriteString(str)
	}
	for lenstr < length {
		buf.WriteString(pad)
		lenstr += lenpad
	}
	if dirleft {
		buf.WriteString(str)
	}
	return buf.String()
}
