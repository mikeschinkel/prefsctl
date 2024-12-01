package macprefs

import (
	"github.com/mikeschinkel/prefsctl/macosutils"
)

type Metadata struct {
	Name   macosutils.Name
	Domain DomainName
	Labels []Label
}
