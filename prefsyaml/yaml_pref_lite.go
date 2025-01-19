package prefsyaml

type YAMLPrefLite struct {
	Domain DomainName
	Name   PrefName
	Value  string // raw string value for default
}
