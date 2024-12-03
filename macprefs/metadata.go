package macprefs

import (
	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/macprefs/kvfilters"
)

type Metadata struct {
	Name   macosutils.Name
	Domain DomainName
	Labels []kvfilters.Label
}
