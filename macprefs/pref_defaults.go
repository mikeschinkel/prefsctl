package macprefs

import (
	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

const prefDefaultsGoImport = "github.com/mikeschinkel/prefsctl/prefdefaults"

type PrefDefaultsFunc func(cfg config.Config) (OSPrefDefaults, error)

var osPrefDefaultsFuncs PrefDefaultsFunc

func RegisterPrefDefaultsFunc(f PrefDefaultsFunc) {
	osPrefDefaultsFuncs = f
}
func GetDefaultsFunc() PrefDefaultsFunc {
	return osPrefDefaultsFuncs
}
func GetAfterApplyFunc(cfg config.Config, domain DomainName) (f func() error, err error) {
	var defaults OSPrefDefaults

	defaultsFunc := GetDefaultsFunc()
	defaults, err = defaultsFunc(cfg)
	if err != nil {
		goto end
	}
	for _, defaults := range defaults.Domains {
		if defaults.Domain != domain {
			continue
		}
		f = func() error {
			return macosutil.KillProcess(defaults.KillOnApply)
		}
		goto end
	}
end:
	return f, err
}
