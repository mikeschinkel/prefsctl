package errutil

import (
	"regexp"
)

var anyWhitespaceRegexp = regexp.MustCompile(`[ \t]+`)
var eolWhitespaceRegexp = regexp.MustCompile(`[ \t]+\n`)

func StripInlineWhitespace(s string) string {
	s = anyWhitespaceRegexp.ReplaceAllString(s, " ")
	return eolWhitespaceRegexp.ReplaceAllString(s, "\n")
}
