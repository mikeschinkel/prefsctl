package macprefs

import (
	"os"
)

type defaultsMapFunc func() DomainPrefDefaults

var defaultsMapFuncs = make(map[MacOSName]defaultsMapFunc)

func RegisterDefaultsMapFunc(os MacOSName, f defaultsMapFunc) {
	defaultsMapFuncs[os] = f
}

// domainPrefDefaults will contain Pref Defaults by Domain after calling macprefs.Initialize()
var domainPrefDefaults DomainPrefDefaults
var prefDefaultsMap = make(map[PrefId]*PrefDefault)

func PrefDefaultsMap() map[PrefId]*PrefDefault {
	var osName MacOSName
	var err error
	var df defaultsMapFunc
	var ok bool

	if len(prefDefaultsMap) > 0 {
		goto end
	}
	osName, err = MacOSVersionName()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	df, ok = defaultsMapFuncs[osName]
	if !ok {
		slog.Error("Unsupported macOS version: %s (%s)", osName)
		os.Exit(2)
	}

	domainPrefDefaults = df()
	for domain, pds := range domainPrefDefaults {
		for _, pd := range pds {
			// Assign domain here vs. having to specify for every pref rather than just once
			// for each group of prefs
			pd.Domain = domain
			prefDefaultsMap[pd.Id()] = pd
		}
	}

end:
	return prefDefaultsMap
}
