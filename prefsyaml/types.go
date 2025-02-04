package prefsyaml

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type (
	PreferenceType = macosutil.PreferenceType
	ProcessName    = macosutil.ProcessName
)
type (
	OSVersion    string
	DomainName   string
	PrefName     string
	LabelName    string
	LabelValue   string
	OptionValue  string
	PrefType     string
	TemplateName string
	TemplateText string
	APIVersion   string
	KindName     string
)

type Printer interface {
	Print(i ...any)
	Println(i ...any)
	Printf(format string, i ...any)
}

type Identifier string
