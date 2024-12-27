package preftemplates

import (
	"github.com/mikeschinkel/prefsctl/types"
)

type TypeName = types.TypeName

type (
	OSVersion    string
	DomainName   string
	PrefName     string
	TemplateName string
	TemplateText string
	APIVersion   string
	KindName     string
)

type TemplateData interface{}

type Printer interface {
	Print(i ...any)
	Println(i ...any)
	Printf(format string, i ...any)
}
