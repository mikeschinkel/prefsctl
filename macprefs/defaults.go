package macprefs

import (
	"os"
	"reflect"
)

const (
	appKey     = "macprefs"
	version    = "v1beta1"
	apiVersion = appKey + "/" + version
)

type prefDefaults []*PrefDefault
type domainPrefDefaults map[string]prefDefaults
type prefsMap map[string]*Pref

var PrefDefaults domainPrefDefaults
var Prefs prefsMap

func init() {

	verName, err := MacOSVersionName()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	// We only need to load Pref Defaults for one macOS version per run.
	switch verName {
	case Sequoia:
		PrefDefaults = sequoiaPrefDefaults()
	case Monterey:
		PrefDefaults = montereyPrefDefaults()
	default:
		slog.Error("Unsupported macOS version: %s (%s)", verName)
		os.Exit(2)
	}

	Prefs = make(prefsMap)
	for domain, defaults := range PrefDefaults {
		for i, prefDefault := range defaults {
			prefDefault.domain = Domain{
				Name: domain,
			}
			// Keep things D.R.Y.; assign domain now vs. having to specify when hardcoding
			// the pref default.
			PrefDefaults[domain][i] = prefDefault

			// Create lookup for pref default
			Prefs[domain+"/"+prefDefault.Name] = &Pref{
				Domain:  domain,
				Name:    prefDefault.Name,
				Kind:    reflectKindFromPrefType(prefDefault.Type),
				Default: prefDefault,
			}
		}
	}
	print()
}
func reflectKindFromPrefType(t PrefType) (rk reflect.Kind) {
	switch t {
	case NumberType:
		rk = reflect.Int // TODO: What about floating point?!?
	case BoolType:
		rk = reflect.Bool
	case StringType, LanguageType, LocaleType:
		rk = reflect.String
	}
	return rk
}
