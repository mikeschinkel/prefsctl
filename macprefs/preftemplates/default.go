package preftemplates

import (
	"github.com/mikeschinkel/prefsctl/macprefs/kvfilters"
)

type Default struct {
	Name      PrefName
	Value     string // raw string value for default
	Labels    kvfilters.Labels
	TypeName  TypeName
	NoDefault bool
	Domain    *Domain
}
