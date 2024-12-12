package macprefs

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
)

type Metadata struct {
	Name   Name
	Domain DomainName
	Labels []kvfilters.Label
}
