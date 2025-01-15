package macprefs

import (
	"errors"
	"fmt"

	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

const prefDefaultsGoImport = "github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"

type PrefDefaultsFunc func(cfg config.Config) (OSPrefDefaults, error)

type PrefDefaultsFuncs map[Code]PrefDefaultsFunc

var osPrefDefaultsFuncs = make(PrefDefaultsFuncs)

func RegisterPrefDefaultsFunc(os *kvfilters.Label, f PrefDefaultsFunc) {
	osPrefDefaultsFuncs[Code(os.Value)] = f
}

func GetAfterApplyFunc(cfg config.Config, domain DomainName) (f func() error, _ error) {
	var defaults OSPrefDefaults

	defaultsFunc, err := GetDefaultsFunc()
	if err != nil {
		goto end
	}
	defaults, err = defaultsFunc(cfg)
	if err != nil {
		goto end
	}
	for _, defaults := range defaults.Domains {
		if defaults.Domain != domain {
			continue
		}
		f = defaults.AfterApplyFunc
		goto end
	}
end:
	return f, err
}
func GetDefaultsFunc() (f PrefDefaultsFunc, _ error) {
	osCode, err := macosutil.VersionCode()
	if err != nil {
		goto end
	}
	f, err = GetOSPrefDefaultsFunc(osCode)
end:
	return f, err
}
func GetOSPrefDefaultsFunc(os Code) (f PrefDefaultsFunc, err error) {
	var ok bool

	f, ok = osPrefDefaultsFuncs[os]
	if !ok {
		err = errors.Join(ErrUnsupportedMacOSVersion,
			fmt.Errorf("%s=%s", logargs.OSName, os),
			fmt.Errorf(`(Did you forget import "%s" in main.go?)`, prefDefaultsGoImport),
		)
	}
	return f, err
}
