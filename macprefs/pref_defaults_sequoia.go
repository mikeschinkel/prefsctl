package macprefs

func sequoiaPrefDefaults() domainPrefDefaults {
	return domainPrefDefaults{
		"GlobalPreferences": []*PrefDefault{
			{
				Name:     "AppleLanguage",
				Type:     StringType,
				WhatSets: InstallSets,
			},
			{
				Name:     "AppleLocale",
				Type:     LocaleType,
				WhatSets: InstallSets,
			},
		},
		"com.apple.Dock": []*PrefDefault{
			{
				Name:     "autohide",
				Type:     BoolType,
				WhatSets: MacOSSets,
				Value:    "false",
			},
		},
	}
}
