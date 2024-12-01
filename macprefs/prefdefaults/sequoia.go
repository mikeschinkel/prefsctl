package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs"
)

func sequoiaPrefDefaults() macprefs.DomainPrefDefaults {
	return montereyPrefDefaults()
}
