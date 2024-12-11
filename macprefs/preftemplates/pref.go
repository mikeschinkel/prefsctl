package preftemplates

import (
	"github.com/mikeschinkel/prefsctl/macprefs/kvfilters"
)

type Pref struct {
	Domain   DomainName
	Name     PrefName
	Default  string // raw string value for default
	Labels   kvfilters.Labels
	Verified bool
}
