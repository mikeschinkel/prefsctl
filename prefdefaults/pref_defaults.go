package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

const prefDefaultsGoImport = "github.com/mikeschinkel/prefsctl/prefdefaults"

type Func func(cfg config.Config, domains []DomainName) (OSPrefDefaults, error)

var thisFunc Func

func RegisterPrefDefaultsFunc(f Func) {
	thisFunc = f
}
func GetDefaultsFunc() Func {
	return thisFunc
}
func GetAfterApplyFunc(cfg config.Config, domain DomainName) (f func() error, err error) {
	var defaults OSPrefDefaults

	defaultsFunc := GetDefaultsFunc()
	defaults, err = defaultsFunc(cfg, []DomainName{domain})
	if err != nil {
		goto end
	}
	for _, defaults := range defaults.Domains {
		if defaults.Domain != domain {
			continue
		}
		f = func() error {
			return macosutil.KillProcess(defaults.Process)
		}
		goto end
	}
end:
	return f, err
}
