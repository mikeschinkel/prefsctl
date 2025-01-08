package macprefs

import (
	"context"

	"github.com/mikeschinkel/prefsctl/cobrautil"
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

type PrefDefaultsMap map[PrefName]*PrefDefault
type PrefDefaults struct {
	Domain         DomainName
	AfterApplyFunc func() error
	DefaultsMap    PrefDefaultsMap
}
type DomainPrefDefaults map[DomainName]PrefDefaults

type Filename string

type BoolPtr *bool
type FilenamePtr *string
type OutputPtr *string

type Printer interface {
	Print(i ...any)
	Println(i ...any)
	Printf(format string, i ...any)
}
