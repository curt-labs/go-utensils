package string_utils

import (
	"io"
	"bytes"
)

func EscapeQuotes(s string) string {
	var buf bytes.Buffer
	last := 0
	for ii, bb := range s {
		if bb == '\'' {
			io.WriteString(&buf, s[last:ii])
			io.WriteString(&buf, `''`)
			last = ii + 1
		}
	}
	io.WriteString(&buf, s[last:])
	return buf.String()
}
