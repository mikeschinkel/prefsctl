package macprefs

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type DefaultsDomain struct {
	Domain      DomainName
	Defaults    []*PrefDefault
	KillOnApply macosutil.ProcessName
}

func (d *DefaultsDomain) AddDefault(def *PrefDefault) {
	d.Defaults = append(d.Defaults, def)
}

func NewDefaultsDomain(domain DomainName, killOnApply macosutil.ProcessName) DefaultsDomain {
	return DefaultsDomain{
		Domain:      domain,
		Defaults:    make([]*PrefDefault, 0),
		KillOnApply: killOnApply,
	}
}
