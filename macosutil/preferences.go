package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation -framework Foundation
#include <CoreFoundation/CoreFoundation.h>
#include <stddef.h>

void syncPreferences(CFStringRef domain) {
    CFPreferencesSynchronize(domain, kCFPreferencesCurrentUser, kCFPreferencesAnyHost);
}
*/
import "C"
import (
	"errors"
	"fmt"

	"github.com/mikeschinkel/prefsctl/errutil"
)

// ApplyPreferences applies the given preferences to the specified domain
func (*macOSUtils) ApplyPreferences(domain string, prefs []Preference) (err error) {
	var errs errutil.MultiErr
	opts := ApplyOpts{
		SkipSync: true,
	}
	for _, pref := range prefs {
		err = pref.ApplyPreference(opts)
		if err != nil {
			errs.Add(errors.Join(ErrInvalidPreferenceDomain,
				fmt.Errorf("%s=%s", pref.Name, pref.Value),
				err))
		}
	}
	SyncDomainPreferences(domain)

	return errs.Err()
}

func SyncDomainPreferences(domain string) {
	cfDomain := NewCFString(domain)
	defer cfDomain.Release()
	C.syncPreferences(cfDomain.cfString)
}
