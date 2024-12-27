package preftemplates

type Pref struct {
	Domain DomainName
	Name   PrefName
	Value  string // raw string value for default
}
