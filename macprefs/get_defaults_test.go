package macprefs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/mikeschinkel/go-diffator"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macosutil/macosutilmock"
	"github.com/mikeschinkel/prefsctl/macprefs"
	_ "github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	GoFormat = macprefs.GoFormat
)

var (
	GlobalFlags = &macprefs.GlobalFlags
	GetDefaults = macprefs.GetDefaults
)

type (
	GetDefaultsArgs  = macprefs.GetDefaultsArgs
	OutputFormat     = macprefs.OutputFormat
	PreferenceDomain = macosutil.PreferenceDomain
	Preference       = macosutil.Preference
)

func TestGetDefaults(t *testing.T) {
	tests := []struct {
		name            string
		args            GetDefaultsArgs
		output          OutputFormat
		errWanted       error
		domains         []PreferenceDomain
		domainsErr      error
		domainPrefs     map[PreferenceDomain][]*Preference
		domainsPrefErrs map[PreferenceDomain]error
		expectedOutput  string
	}{
		{
			name: "",
			args: GetDefaultsArgs{
				OmitEmpty: false,
			},
			output:          GoFormat,
			errWanted:       nil,
			domains:         domainsForTest(),
			domainsErr:      nil,
			domainPrefs:     domainPrefsForTest(),
			domainsPrefErrs: nil,
			expectedOutput:  expectedOutputForTest(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			GlobalFlags.Output = stdlibex.Ptr(string(tt.output))
			GlobalFlags.Quiet = stdlibex.Ptr(true)
			tt.args.Printer = &errutil.BufferPrinter{}
			mock := macosutilmock.NewMock()
			macosutil.SetInstance(mock)
			m := mock.(*macosutilmock.MacOSUtilMock)
			m.SetPreferenceDomains(tt.domains, tt.domainsErr)
			m.SetPreferences(tt.domainPrefs, tt.domainsPrefErrs)
			err := GetDefaults(ctx, tt.args)
			if errutil.ErrorCheckFails(t, "GetDefaults", tt.errWanted, err) {
				return
			}
			expected := tt.expectedOutput
			received := tt.args.PrinterOutput()
			expected = errutil.StripInlineWhitespace(expected)
			received = errutil.StripInlineWhitespace(received)
			if expected != received {
				diff := diffator.CompareStrings(expected, received, nil)
				t.Errorf("Body <(expected/received)>:\n\t%s", diff)
				return
			}
		})
	}
}

func domainsForTest() []PreferenceDomain {
	return []PreferenceDomain{
		"com.apple.Accessibility",
		"com.apple.dock",
		"com.apple.finder",
		"com.apple.findmy",
		"com.apple.keyboard",
		"GlobalPreferences",
	}
}
func domainPrefsForTest() map[PreferenceDomain][]*Preference {
	return map[PreferenceDomain][]*Preference{
		"GlobalPreferences": {
			{
				Domain:      "GlobalPreferences",
				Name:        "AppleActionOnDoubleClick",
				Value:       "",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "GlobalPreferences",
				Name:        "AppleMeasurementUnits",
				Value:       "",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "GlobalPreferences",
				Name:        "com.apple.mouse.scaling",
				Value:       "",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "GlobalPreferences",
				Name:        "com.apple.trackpad.scaling",
				Value:       "",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "GlobalPreferences",
				Name:        "AppleTemperatureUnit",
				Value:       "",
				Description: "",
				Kind:        0,
			},
		},
		"com.apple.Accessibility": {
			{
				Domain:      "com.apple.Accessibility",
				Name:        "InvertColorsEnabled",
				Value:       "",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "com.apple.Accessibility",
				Name:        "KeyRepeatDelay",
				Value:       "0.25",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "com.apple.Accessibility",
				Name:        "KeyRepeatInterval",
				Value:       "",
				Description: "",
				Kind:        0,
			},
		},
		"com.apple.dock": {
			{
				Domain:      "com.apple.dock",
				Name:        "autohide",
				Value:       "",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "com.apple.dock",
				Name:        "tilesize",
				Value:       "",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "com.apple.dock",
				Name:        "region",
				Value:       "",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "com.apple.dock",
				Name:        "showhidden",
				Value:       "",
				Description: "",
				Kind:        0,
			},
		},
		"com.apple.finder": {
			{
				Domain:      "com.apple.finder",
				Name:        "ShowRemovableMediaOnDesktop",
				Value:       "",
				Description: "",
				Kind:        0,
			},
			{
				Domain:      "com.apple.finder",
				Name:        "ShowStatusBar",
				Value:       "",
				Description: "",
				Kind:        0,
			},
		},
		"com.apple.keyboard": {
			{
				Domain:      "com.apple.keyboard",
				Name:        "KeyboardWordOrSentenceTrackingForPFL",
				Value:       "",
				Description: "",
				Kind:        0,
			},
		},
	}
}

func expectedOutputForTest() string {
	goCode := `package prefdefaults

//goland:noinspection SpellCheckingInspection
func %sPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"com.apple.Accessibility": DomainPrefs{
      "InvertColorsEnabled": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeyRepeatDelay": DomainPref{
				Verified: Verified{},
				Type:     "float",
				Default:  "0.25",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeyRepeatInterval": DomainPref{
				Verified: Verified{},
				Type:     "unknown",
				Default:  "",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.dock": DomainPrefs{
			"autohide": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"region": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showhidden": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"tilesize": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "70",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.finder": DomainPrefs{
			"ShowRemovableMediaOnDesktop": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowStatusBar": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"com.apple.keyboard": DomainPrefs{
			"KeyboardWordOrSentenceTrackingForPFL": DomainPref{
				Verified: Verified{},
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
		"GlobalPreferences": DomainPrefs{
			"AppleActionOnDoubleClick": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Maximize",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleMeasurementUnits": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Inches",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleTemperatureUnit": DomainPref{
				Verified: Verified{},
				Type:     "string",
				Default:  "Fahrenheit",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.mouse.scaling": DomainPref{
				Verified: Verified{},
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.trackpad.scaling": DomainPref{
				Verified: Verified{},
				Type:     "intBool",
				Default:  "1",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
		},
	} 
}`
	versionCode, _ := macosutil.VersionCode()
	return fmt.Sprintf(goCode, versionCode)
}
