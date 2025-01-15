package prefdefaults

//goland:noinspection SpellCheckingInspection
func montereyPrefDefaults() OSPrefDefaults {
	return OSPrefDefaults{
		Domains: []Domain{
			comAppleDock(),
		},
	}
}
