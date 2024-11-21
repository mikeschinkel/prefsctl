package stdlibex

import (
	"fmt"
	"strings"
)

type String string

func (s String) String() string {
	return string(s)
}

// IndentLines each line of a string by `n` number of spaces
// Uses tab (`\t`) is n==0
func IndentLines(n int, s string) string {
	var indent string
	switch n {
	case 0:
		indent = "\t"
	default:
		indent = strings.Repeat(" ", n)
	}
	return strings.Replace("\n"+s, "\n", "\n"+indent, -1)[1:]
}

// AsciiToUpper uppercases ASCII strings in a manner slightly more performant
// than strings.ToUpper().
func AsciiToUpper(s any) string {
	var b []byte
	switch v := s.(type) {
	case []byte:
	case string:
		b = []byte(v)
	case fmt.Stringer:
		b = []byte(v.String())
	default:
		b = []byte(fmt.Sprintf("%v", s))
	}
	for i, c := range b {
		if c < 97 {
			continue
		}
		if c > 122 {
			continue
		}
		b[i] = c - 32
	}
	return string(b)
}

// AsciiToLower lowercases ASCII strings in a manner slightly more performant
// than strings.ToLower().
func AsciiToLower(s any) string {
	var b []byte
	switch v := s.(type) {
	case []byte:
	case string:
		b = []byte(v)
	case fmt.Stringer:
		b = []byte(v.String())
	default:
		b = []byte(fmt.Sprintf("%v", s))
	}
	for i, c := range b {
		if c < 65 {
			continue
		}
		if c > 90 {
			continue
		}
		b[i] = c + 32
	}
	return string(b)
}
