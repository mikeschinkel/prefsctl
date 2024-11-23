package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs"
)

func montereyPrefDefaults() macprefs.DomainPrefDefaults {
	return sequoiaPrefDefaults()
}
