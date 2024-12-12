package preftemplates

import (
	"github.com/mikeschinkel/prefsctl/types"
)

type TypeName = types.TypeName

type OSVersion string
type DomainName string
type PrefName string

type Template interface {
	Generate() (string, error)
}
