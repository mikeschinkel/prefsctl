package macprefs

import (
	"os"
	"sync"

	"github.com/mikeschinkel/prefsctl/logging"
)

const prefDefaultsGoImport = "github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"

type defaultsMapFunc func() DomainPrefDefaults

var defaultsMapFuncs = make(map[Name]defaultsMapFunc)

func RegisterDefaultsMapFunc(os Label, f defaultsMapFunc) {
	defaultsMapFuncs[Name(os.Name)] = f
}

// domainPrefDefaults will contain Pref Defaults by Domain after calling macprefs.Initialize()
var domainPrefDefaults DomainPrefDefaults
var prefDefaultsMap = make(map[PrefId]*PrefDefault)
var prefDefaultsMapMutex sync.Mutex

func PrefDefaultsMap() map[PrefId]*PrefDefault {
	var osName Name
	var err error
	var df defaultsMapFunc
	var ok bool

	if len(prefDefaultsMap) > 0 {
		goto end
	}
	prefDefaultsMapMutex.Lock()

	osName, err = MacOSVersionName()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	df, ok = defaultsMapFuncs[osName]
	if !ok {
		slog.Error("Unsupported macOS version", logging.OSNameLogArg, osName)
		slog.Error(`Did you forget "import _ "%s" in main.go?`, prefDefaultsGoImport)
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
