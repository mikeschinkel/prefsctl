package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs"
)

type dpd = macprefs.DomainPrefDefaults

var register = macprefs.RegisterDefaultsMapFunc

func init() {
	register(macprefs.Sequoia, func() dpd {
		return sequoiaPrefDefaults()
	})
	register(macprefs.Monterey, func() dpd {
		return montereyPrefDefaults()
	})
}
