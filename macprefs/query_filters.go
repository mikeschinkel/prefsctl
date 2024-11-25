package macprefs

const (
	FilterDomain    = FilterString
	FilterPrefKey   = FilterKey
	FilterPrefValue = FilterValue
	FilterPref      = FilterKeyValue
)

func QueryFilters() KeyValueFilters {
	return KeyValueFilters{
		DoesNotHavePrefixFilter{
			Filter: "com.apple.",
			FilterBase: FilterBase{
				Type:   OmitFilterType,
				Target: FilterDomain,
			},
		},
		HasPrefixFilter{
			Filter: "seed-numNotifications-",
			FilterBase: FilterBase{
				Type:   OmitFilterType,
				Target: FilterKey,
			},
		},
		EqualsFilter{
			Filter: "CKStartupTime",
			FilterBase: FilterBase{
				Type:   OmitFilterType,
				Target: FilterKey,
			},
		},
		//RegexFilter{},
		//EqualsFilter{},
		//HasPrefixFilter{},
		//HasSuffixFilter{},
		//ContainsFilter{},
		//NotContainsFilter{},
		//FuncFilter{},
	}
}
