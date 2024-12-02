package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/mikeschinkel/prefsctl/macprefs/filters"
)

var register = macprefs.RegisterDefaultsMapFunc

var (
	SequoiaLabel  = OSVersionLabel(macosutils.Sequoia)
	MontereyLabel = OSVersionLabel(macosutils.Monterey)
)

func OSVersionLabel(code macosutils.Code) filters.Label {
	return filters.Label{
		Name:  filters.LabelName(filters.MacOS),
		Value: filters.LabelValue(code),
	}
}

type DomainDefaults = map[string]DomainPrefs
type DomainPrefs = map[string]DomainPref
type DomainPref struct {
	Domain       string
	Name         string
	DefaultValue string // raw string value for default
	Labels       filters.Labels
	NoDefault    bool
}

func init() {
	register(SequoiaLabel, func() macprefs.DomainPrefDefaults {
		return convertDomainDefaultsToMacprefsDomainPrefDefaults(sequoiaPrefDefaults())
	})
	register(MontereyLabel, func() macprefs.DomainPrefDefaults {
		return convertDomainDefaultsToMacprefsDomainPrefDefaults(montereyPrefDefaults())
	})
}

// convertDomainDefaultsToMacprefsDomainPrefDefaults is designed to make manually
// managing the defaults easy via DomainDefaults but then allow the program to
// have what it needs in macprefs.DomainPrefDefaults.
//
// The impetus for creating two data structures and converting was to allow
// `DomainDefaults` to have a `Labels` property and then allow
// `macprefs.PrefDefaults` to support the `Labels` method required by the
// `filters.KeyValue` interface.
func convertDomainDefaultsToMacprefsDomainPrefDefaults(defaults DomainDefaults) (dpd macprefs.DomainPrefDefaults) {
	dpd = make(macprefs.DomainPrefDefaults, len(defaults))
	for domain, prefs := range defaults {
		pp := make(macprefs.PrefDefaultsMap, len(prefs))
		for name, pref := range prefs {
			pd := &macprefs.PrefDefault{
				Domain:       macprefs.DomainName(domain),
				Name:         macprefs.PrefName(name),
				DefaultValue: pref.DefaultValue,
				NoDefault:    pref.NoDefault,
			}
			pd.SetLabels(pref.Labels)
			pp[macprefs.PrefName(name)] = pd
		}
		dpd[macprefs.DomainName(domain)] = pp
	}
	return dpd
}
