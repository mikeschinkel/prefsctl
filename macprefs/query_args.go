package macprefs

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type QueryArgs struct {
	OmitEmpty bool
	//UseCurrent       bool
	IncludeUnchanged bool
	Domains          []DomainName
}

func (args *QueryArgs) ToMacOSUtilPreferenceDomains() (pds []macosutil.PreferenceDomain) {
	pds = make([]macosutil.PreferenceDomain, len(args.Domains))
	for i, domain := range args.Domains {
		pds[i] = macosutil.PreferenceDomain(domain)
	}
	return pds
}
