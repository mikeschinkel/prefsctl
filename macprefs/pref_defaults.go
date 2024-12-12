package macprefs

import (
	"errors"
	"fmt"

	"github.com/mikeschinkel/prefsctl/kvfilters"
)

const prefDefaultsGoImport = "github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"

type DefaultsMapFunc func() DomainPrefDefaults

type DefaultsMapFuncs map[Code]DefaultsMapFunc

var defaultsMapFuncs = make(DefaultsMapFuncs)

func RegisterDefaultsMapFunc(os *kvfilters.Label, f DefaultsMapFunc) {
	defaultsMapFuncs[Code(os.Value)] = f
}
func GetDefaultsMapFunc(os Code) (f DefaultsMapFunc, err error) {
	var ok bool
	f, ok = defaultsMapFuncs[os]
	if !ok {
		err = errors.Join(ErrUnsupportedMacOSVersion,
			fmt.Errorf(`(Did you forget import "%s" in main.go?)`, prefDefaultsGoImport),
		)
	}
	return f, err
}
