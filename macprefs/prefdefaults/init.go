package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/macprefs"
)

type dpd = macprefs.DomainPrefDefaults

var register = macprefs.RegisterDefaultsMapFunc

var (
	SequoiaLabel  = OSVersionLabel(macosutils.Sequoia)
	MontereyLabel = OSVersionLabel(macosutils.Monterey)
)

func OSVersionLabel(code macosutils.Code) macprefs.Label {
	return macprefs.Label{
		Name:  macprefs.LabelName(macprefs.MacOS),
		Value: macprefs.LabelValue(code),
	}
}

func init() {
	register(SequoiaLabel, func() dpd {
		return sequoiaPrefDefaults()
	})
	register(MontereyLabel, func() dpd {
		return montereyPrefDefaults()
	})
}
