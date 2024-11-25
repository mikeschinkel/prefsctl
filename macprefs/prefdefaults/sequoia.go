package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs"
)

func sequoiaPrefDefaults() macprefs.DomainPrefDefaults {
	return macprefs.DomainPrefDefaults{
		"GlobalPreferences": []*macprefs.PrefDefault{
			{
				Name: "AppleLanguage",
				Labels: macprefs.Labels{
					&macprefs.InstallSets,
					&macprefs.StringType,
				},
			},
			{
				Name: "AppleLocale",
				Labels: macprefs.Labels{
					&macprefs.InstallSets,
					&macprefs.StringType,
				},
			},
		},
		"com.apple.Dock": []*macprefs.PrefDefault{
			{
				Name:  "autohide",
				Value: "false",
				Labels: macprefs.Labels{
					&macprefs.MacOSSets,
					&macprefs.BoolType,
				},
			},
		},
	}
}
