package macosutils_test

import (
	"errors"
	"testing"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

var (
	RetrievePreferences = macosutils.RetrievePreferences
)

type (
	PreferenceDomain = macosutils.PreferenceDomain
	Preference       = macosutils.Preference
)

func Test_macOSUtils_RetrievePreferenceDomains(t *testing.T) {
	gotDomains, err := macosutils.RetrievePreferenceDomains()
	if err != nil {
		t.Errorf("RetrievePreferenceDomains() error = %v", err)
		return
	}
	if gotDomains == nil {
		t.Error("RetrievePreferenceDomains() returned nil")
	}
	if len(gotDomains) <= 500 {
		t.Errorf("RetrievePreferenceDomains() only returned %d domains; more than 500 expected", len(gotDomains))
	}
}

func Test_macOSUtils_RetrievePreferences(t *testing.T) {
	tests := []struct {
		name      string
		domain    string
		errWanted error
		wantPrefs []string
		wantMatch bool
	}{
		{
			name:      "Bad domain - baz.bar.foo",
			domain:    "baz.bar.foo",
			errWanted: macosutils.ErrFailedToGetPrefNames,
		},
		{
			name:   "com.apple.finder — Match all",
			domain: "com.apple.finder",
			wantPrefs: []string{
				"DisableAllAnimations",
				"NewWindowTarget",
				"PreviewPaneWidth",
				"QuitMenuItem",
				"ShowHardDrivesOnDesktop",
				"ShowExternalHardDrivesOnDesktop",
				"ShowPathbar",
				"ShowSidebar",
			},
			wantMatch: true,
		},
		{
			name:   "com.apple.finder — Don't match FooBarBaz",
			domain: "com.apple.finder",
			wantPrefs: []string{
				"FooBarBaz",
			},
			wantMatch: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefs, err := RetrievePreferences(PreferenceDomain(tt.domain))
			if errCheckFails(t, "RetrievePreferences()", tt.errWanted, err) {
				return
			}

			found, _ := chkPrefs(gotPrefs, tt.wantPrefs)

			switch {
			case tt.wantMatch:
				if stdlibex.AllInSlice(tt.wantPrefs, found) {
					return
				}
				names, _ := sliceconv.ToStringsFunc(gotPrefs, func(p *Preference) (bool, string, error) {
					return true, p.Name, nil
				})
				t.Errorf("RetrievePreferences() prefs wanted but not found:\n\tNot Found:%v",
					stdlibex.DiffSlices(tt.wantPrefs, names),
				)

			case !tt.wantMatch:
				if len(found) > 0 {
					t.Errorf("RetrievePreferences() prefs found but not wanted:\n\t%v", found)
				}
			}
		})
	}
}

func chkPrefs(prefs []*Preference, names []string) (found, missing []string) {
	found = make([]string, 0, len(names))
	missing = make([]string, 0, len(names))
	nameMap := sliceconv.ToIndexMap(names)
	for _, pref := range prefs {
		_, ok := nameMap[pref.Name]
		if ok {
			found = append(found, pref.Name)
		} else {
			missing = append(missing, pref.Name)
		}
	}
	return found, missing
}

func errCheckFails(t *testing.T, testName string, errWanted, errGot error) (failed bool) {
	switch {
	case errWanted == nil && errGot != nil:
		t.Errorf("%s: did not want error, got '%v'", testName, errutil.CleanErrString(errGot))
		failed = true
	case errWanted != nil:
		switch {
		case errGot == nil:
			t.Errorf("%s: want error, got nil", testName)
		case !errors.Is(errGot, errWanted):
			t.Errorf("%s: wanted error but wrong message:\n\tExpected: %s\n\tReceived: %s",
				testName,
				errutil.CleanErrString(errWanted),
				errutil.CleanErrString(errGot),
			)
		}
		failed = true
	}
	return failed
}
