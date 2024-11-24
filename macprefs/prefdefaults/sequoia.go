package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs"
)

func sequoiaPrefDefaults() macprefs.DomainPrefDefaults {
	return macprefs.DomainPrefDefaults{
		"GlobalPreferences": []*macprefs.PrefDefault{
			{
				Name: "AppleLanguage",
				Type: macprefs.StringType,
				Labels: []macprefs.Label{
					macprefs.InstallSets,
				},
			},
			{
				Name: "AppleLocale",
				Type: macprefs.LocaleType,
				Labels: []macprefs.Label{
					macprefs.InstallSets,
				},
			},
		},
		"com.apple.Dock": []*macprefs.PrefDefault{
			{
				Name: "autohide",
				Type: macprefs.BoolType,
				Labels: []macprefs.Label{
					macprefs.MacOSSets,
				},
				Value: "false",
			},
		},
	}
}
