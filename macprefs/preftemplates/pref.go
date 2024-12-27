package preftemplates

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
)

type Pref struct {
	Domain  DomainName
	Name    PrefName
	Default string // raw string value for default
	Labels  kvfilters.Labels
}
