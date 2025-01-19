package prefsyaml

type YAMLDomain struct {
	Name     DomainName
	Defaults []*YAMLDefault
	Prefs    []*YAMLPrefLite
	Opts     PrefsResourceOpts
}
