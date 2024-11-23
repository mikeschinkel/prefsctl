package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/macprefs"
)

func sequoiaPrefDefaults() macprefs.DomainPrefDefaults {
	return macprefs.DomainPrefDefaults{
		"GlobalPreferences": []*macprefs.PrefDefault{
			{
				Name:     "AppleLanguage",
				Type:     macprefs.StringType,
				WhatSets: macprefs.InstallSets,
			},
			{
				Name:     "AppleLocale",
				Type:     macprefs.LocaleType,
				WhatSets: macprefs.InstallSets,
			},
		},
		"com.apple.Dock": []*macprefs.PrefDefault{
			{
				Name:     "autohide",
				Type:     macprefs.BoolType,
				WhatSets: macprefs.MacOSSets,
				Value:    "false",
			},
		},
	}
}
