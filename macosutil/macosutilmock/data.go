package macosutilmock

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type Data struct {
	Domains         []macosutil.PreferenceDomain
	DomainsErr      error
	DomainPrefs     map[macosutil.PreferenceDomain][]*macosutil.Preference
	DomainsPrefErrs map[macosutil.PreferenceDomain]error
}
