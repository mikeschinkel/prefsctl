package stdlibex

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var defaultLanguage language.Tag
var titler cases.Caser

func Title(s string) string {
	if defaultLanguage != DefaultLanguage {
		defaultLanguage = DefaultLanguage
		titler = cases.Title(defaultLanguage)
	}
	return titler.String(s)
}
