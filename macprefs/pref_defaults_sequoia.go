package macprefs

func sequoiaPrefDefaults() domainPrefDefaults {
	return domainPrefDefaults{
		"GlobalPreferences": []*PrefDefault{
			&PrefDefault{
				Name:     "AppleLanguage",
				Type:     StringType,
				WhatSets: InstallSets,
			},
			&PrefDefault{
				Name:     "AppleLocale",
				Type:     LocaleType,
				WhatSets: InstallSets,
			},
		},
		"com.apple.Dock": []*PrefDefault{
			&PrefDefault{
				Name:     "autohide",
				Type:     BoolType,
				WhatSets: MacOSSets,
				Value:    "false",
			},
		},
	}
}
