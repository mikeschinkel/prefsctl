package preftemplates

type Domain struct {
	Name     DomainName
	Defaults []*Default
	MacOS    *MacOS
}
