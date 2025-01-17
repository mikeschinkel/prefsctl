package macprefs

import (
	"context"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/types"
)

type (
	Context   = context.Context
	CmdResult = cobrautil.CmdResult
)

type YAMLSpec interface {
	YAMLSpec()
}
type Spec interface {
	Spec()
}

type PrefName string
type TypeName = types.TypeName
type Name = types.Name
type Code = types.Code

type PrefDomain struct {
	Domain      DomainName
	Defaults    []*PrefDefault
	KillOnApply macosutil.ProcessName
}

func NewPrefDomain(domain DomainName, killOnApply macosutil.ProcessName) PrefDomain {
	return PrefDomain{
		Domain:      domain,
		Defaults:    make([]*PrefDefault, 0),
		KillOnApply: killOnApply,
	}
}

// OSPrefDefaults is used to load a list of all domains and its default
// preferences for a given macOS version
type OSPrefDefaults struct {
	Domains []PrefDomain
}

type Filename string

type BoolPtr *bool
type FilenamePtr *string
type OutputPtr *string

type Printer interface {
	Print(i ...any)
	Println(i ...any)
	Printf(format string, i ...any)
}
