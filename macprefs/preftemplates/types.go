package preftemplates

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type PreferenceType = macosutil.PreferenceType

type (
	OSVersion    string
	DomainName   string
	PrefName     string
	LabelName    string
	LabelValue   string
	PrefType     string
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
