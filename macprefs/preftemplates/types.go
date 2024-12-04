package preftemplates

type TypeName string
type OSVersion string
type DomainName string
type PrefName string

type Template interface {
	Generate() (string, error)
}
