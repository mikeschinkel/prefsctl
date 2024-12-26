package macprefstest

import (
	"fmt"

	"github.com/mikeschinkel/prefsctl/macosutil"
)

type (
	PreferenceDomain = macosutil.PreferenceDomain
	Preference       = macosutil.Preference
)

func DomainsForTest() []PreferenceDomain {
	return []PreferenceDomain{
		"com.apple.Accessibility",
		"com.apple.dock",
		"com.apple.finder",
		"com.apple.findmy",
		"com.apple.keyboard",
		"GlobalPreferences",
	}
}
func DomainPrefsForTest() map[PreferenceDomain][]*Preference {
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
func ExpectedOutputForTest() string {
	goCode := `package prefdefaults

//goland:noinspection SpellCheckingInspection
func %sPrefDefaults() DomainDefaults {
	return DomainDefaults{
		"com.apple.Accessibility": DomainPrefs{
			"InvertColorsEnabled": DomainPref{
				Type:     "intBool",
				Default:  "0",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"KeyRepeatDelay": DomainPref{
				Type:     "float",
				Default:  "0.25",
				Labels: NewLabels(
					DefaultsSet,
					TypeVerified,
					UserManaged,
				),
			},
			"KeyRepeatInterval": DomainPref{
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
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"region": DomainPref{
				Type:     "string",
				Default:  "US",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"showhidden": DomainPref{
				Type:     "bool",
				Default:  "true",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"tilesize": DomainPref{
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
				Type:     "bool",
				Default:  "false",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"ShowStatusBar": DomainPref{
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
				Type:     "string",
				Default:  "Maximize",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleMeasurementUnits": DomainPref{
				Type:     "string",
				Default:  "Inches",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"AppleTemperatureUnit": DomainPref{
				Type:     "string",
				Default:  "Fahrenheit",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.mouse.scaling": DomainPref{
				Type:     "int",
				Default:  "3",
				Labels: NewLabels(
					DefaultsSet,
					UserManaged,
				),
			},
			"com.apple.trackpad.scaling": DomainPref{
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
