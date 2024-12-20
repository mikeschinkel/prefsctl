package macosutil_test

import (
	"reflect"
	"testing"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

var (
	RetrievePreference = macosutil.RetrievePreference
	ErrorCheckFails    = errutil.ErrorCheckFails
)

func Test_macOSUtils_RetrievePreference(t *testing.T) {
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
