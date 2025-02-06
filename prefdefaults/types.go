package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/types"
)

type (
	DomainName = macosutil.DomainName
	PrefName   = macosutil.PreferenceName
	TypeName   = macosutil.PreferenceType
)

type (
	Code           = types.Code
	PreferenceType = macosutil.PreferenceType
)

type QueryArgs struct {
	Domains []DomainName
}
