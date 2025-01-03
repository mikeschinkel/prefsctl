package macosutil_test

import (
	"reflect"
	"testing"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

var (
	RetrievePreference = macosutil.RetrievePreference
	ApplyPreferences   = macosutil.ApplyPreferences
	ErrorCheckFails    = errutil.ErrorCheckFails
)

func Test_RetrievePreference(t *testing.T) {
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
			gotPref, err := RetrievePreference(tt.domain, tt.prop)
			if ErrorCheckFails(t, "RetrievePreference()", tt.errWanted, err) {
				return
			}
			if tt.wantKind != gotPref.Kind {
				t.Errorf("RetrievePreference(): wanted kind '%s', got '%s'",
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

func Test_ApplyPreferences(t *testing.T) {
	tests := []struct {
		domain  string
		name    string
		value   string
		wantErr bool
	}{
		{
			domain: "com.apple.dock",
			name:   "autohide-delay",
			value:  "0",
		},
		{
			domain: "com.apple.dock",
			name:   "largesize",
			value:  "75",
		},
		{
			domain: "com.apple.dock",
			name:   "magnification",
			value:  "true",
		},
		{
			domain: "com.apple.dock",
			name:   "minimize-to-application",
			value:  "true",
		},
		{
			domain: "com.apple.dock",
			name:   "tilesize",
			value:  "70",
		},
	}
	for _, tt := range tests {
		t.Run(tt.domain+"_"+tt.name, func(t *testing.T) {
			pref := []Preference{
				{
					Domain: tt.domain,
					Name:   tt.name,
					Value:  tt.value,
				},
			}
			if err := ApplyPreferences(tt.domain, pref); (err != nil) != tt.wantErr {
				t.Errorf("ApplyPreferences() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
