package preftemplates

import (
	"strings"
	"text/template"
)

var templateFuncs = template.FuncMap{
	"firstUp": func(s string) string {
		switch len(s) {
		case 0:
			return ""
		case 1:
			return strings.ToUpper(s)
		default:
			return strings.ToUpper(s[0:1]) + s[1:]
		}
	},
}
