package macprefs

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/prefdefaults"
)

type DefaultsDomain struct {
	Domain   DomainName
	Defaults []*prefdefaults.PrefDefault
	Process  macosutil.ProcessName
}

func (d *DefaultsDomain) AddDefault(def *prefdefaults.PrefDefault) {
	d.Defaults = append(d.Defaults, def)
}

func NewDefaultsDomain(domain DomainName, process macosutil.ProcessName) DefaultsDomain {
	return DefaultsDomain{
		Domain:   domain,
		Defaults: make([]*prefdefaults.PrefDefault, 0),
		Process:  process,
	}
}
