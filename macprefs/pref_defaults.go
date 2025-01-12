package macprefs

import (
	"errors"
	"fmt"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

const prefDefaultsGoImport = "github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"

type DefaultsMapFunc func() DomainPrefDefaults

type DefaultsMapFuncs map[Code]DefaultsMapFunc

var defaultsMapFuncs = make(DefaultsMapFuncs)

func RegisterDefaultsMapFunc(os *kvfilters.Label, f DefaultsMapFunc) {
	defaultsMapFuncs[Code(os.Value)] = f
}

func GetAfterApplyFunc(domain DomainName) (f func() error, _ error) {
	mapFunc, err := GetDefaultsMapFunc()
	if err != nil {
		goto end
	}
	for name, defaults := range mapFunc() {
		if name != domain {
			continue
		}
		f = defaults.AfterApplyFunc
		goto end
	}
end:
	return f, err
}
func GetDefaultsMapFunc() (f DefaultsMapFunc, _ error) {
	osCode, err := macosutil.VersionCode()
	if err != nil {
		goto end
	}
	f, err = GetDefaultsMapFuncForOS(osCode)
end:
	return f, err
}
func GetDefaultsMapFuncForOS(os Code) (f DefaultsMapFunc, err error) {
	var ok bool

	f, ok = defaultsMapFuncs[os]
	if !ok {
		err = errors.Join(ErrUnsupportedMacOSVersion,
			fmt.Errorf("%s=%s", logargs.OSName, os),
			fmt.Errorf(`(Did you forget import "%s" in main.go?)`, prefDefaultsGoImport),
		)
	}
	return f, err
}
