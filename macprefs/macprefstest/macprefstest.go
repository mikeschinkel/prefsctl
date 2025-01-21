package macprefstest

import (
	"os"
	"testing"

	"github.com/mikeschinkel/prefsctl/macosutil"
)

const GetDefaultsTestOutputFile = "macprefstest/data/test-get-defaults.yaml"

func ExpectedOutputForTest(t *testing.T) string {
	content, _ := os.ReadFile(GetDefaultsTestOutputFile)
	return string(content)
}

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
