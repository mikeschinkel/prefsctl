package kvfilters

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

func Codify(s string) Code {
	ss := strings.Split(s, " ")
	if len(ss[0]) != 0 && 'A' <= ss[0][0] && ss[0][0] <= 'Z' {
		ss[0] = strings.ToLower(string(ss[0][0])) + s[1:]
	}
	titler := cases.Title(language.English)
	for i := 1; i < len(ss); i++ {
		ss[i] = titler.String(ss[i])
	}
	return Code(strings.Join(ss, ""))
}
