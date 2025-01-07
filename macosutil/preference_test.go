package macosutil_test

import (
	"reflect"
	"testing"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

var (
	ApplyPreferences = macosutil.ApplyPreferences
	ErrorCheckFails  = errutil.ErrorCheckFails
	GetPreference    = macosutil.GetPreference
)

func Test_GetPreference(t *testing.T) {
	tests := []struct {
		name      string
		domain    string
		prop      string
		wantKind  reflect.Kind
		errWanted error
	}{
		{
			name:     "com.apple.finder ShowSidebar",
			domain:   "com.apple.finder",
			prop:     "ShowSidebar",
			wantKind: reflect.Bool,
		},
		{
			name:      "not.a.valid.domain NotValid",
			domain:    "not.a.valid.domain",
			prop:      "NotValid",
			errWanted: macosutil.ErrInvalidPreferenceDomain,
		},
		{
			name:      "com.apple.finder InvalidProperty",
			domain:    "com.apple.finder",
			prop:      "InvalidProperty",
			errWanted: macosutil.ErrPreferenceNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPref, err := GetPreference(tt.domain, tt.prop)
			if ErrorCheckFails(t, "GetPreference()", tt.errWanted, err) {
				return
			}
			if tt.wantKind != gotPref.Kind {
				t.Errorf("GetPreference(): wanted kind '%s', got '%s'",
					tt.wantKind.String(),
					gotPref.Kind.String(),
				)
				if tt.wantKind == reflect.Invalid {
					t.Error("Did you maybe forget to set .kind for this test?")
				}
			}
		})
	}
}

type pair [2]string

func (p pair) toggle(value string) string {
	if value == p[0] {
		return p[1]
	}
	return p[0]
}

func Test_ApplyPreferences(t *testing.T) {
	tests := []struct {
		domain  string
		name    string
		values  pair
		wantErr bool
	}{
		{
			domain: "com.apple.dock",
			name:   "tilesize",
			values: pair{"70", "50"},
		},
		{
			domain: "com.apple.dock",
			name:   "largesize",
			values: pair{"75", "128"},
		},
		{
			domain: "com.apple.dock",
			name:   "minimize-to-application",
			values: pair{"true", "false"},
		},
		{
			domain: "com.apple.dock",
			name:   "autohide-delay",
			values: pair{"0", "1"},
		},
		{
			domain: "com.apple.dock",
			name:   "magnification",
			values: pair{"true", "false"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.domain+"_"+tt.name, func(t *testing.T) {
			pref, err := GetPreference(tt.domain, tt.name)
			if err != nil {
				t.Errorf("ApplyPreferences() Initial Get error: %s", err.Error())
				return
			}
			value := pref.Value
			expected := tt.values.toggle(value)
			pref.Value = expected
			prefs := []*Preference{pref}
			if err := ApplyPreferences(tt.domain, prefs); (err != nil) != tt.wantErr {
				t.Errorf("ApplyPreferences() SET error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil {
				return
			}
			pref, err = GetPreference(tt.domain, tt.name)
			if err != nil {
				t.Errorf("ApplyPreferences() Subsequent Get error: %s", err.Error())
				return
			}
			if pref.Value != expected {
				t.Errorf("ApplyPreferences() value error: expected after Apply(): %s, got instead: %s", expected, pref.Value)
				return
			}
			prefs = []*Preference{pref}
			pref.Value = value
			if err := ApplyPreferences(tt.domain, prefs); (err != nil) != tt.wantErr {
				t.Errorf("ApplyPreferences() RESET error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
