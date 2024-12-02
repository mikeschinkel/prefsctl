package macprefs

import (
	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/macprefs/filters"
)

type Metadata struct {
	Name   macosutils.Name
	Domain DomainName
	Labels []filters.Label
}
