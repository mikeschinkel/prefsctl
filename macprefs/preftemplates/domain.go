package preftemplates

type Domain struct {
	Name     DomainName
	Defaults []*Default
	Prefs    []*Pref
	MacOS    *MacOS
}
