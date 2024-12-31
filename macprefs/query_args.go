package macprefs

import (
	"fmt"

	"github.com/mikeschinkel/prefsctl/macosutil"
)

type QueryArgs struct {
	Printer          Printer
	OmitEmpty        bool
	UseCurrent       bool
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

func (args *QueryArgs) PrinterOutput() string {
	return args.Printer.(fmt.Stringer).String()
}
