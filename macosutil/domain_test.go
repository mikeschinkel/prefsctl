package macosutil_test

import (
	"testing"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

var (
	RetrievePreferences = macosutil.RetrievePreferences
)

type (
	PreferenceDomain = macosutil.PreferenceDomain
	Preference       = macosutil.Preference
)

func Test_macOSUtils_RetrievePreferenceDomains(t *testing.T) {
	gotDomains, err := macosutil.RetrievePreferenceDomains()
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
			errWanted: macosutil.ErrFailedToGetPrefDomains,
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
			if errutil.ErrorCheckFails(t, "RetrievePreferences()", tt.errWanted, err) {
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