package macprefs

import (
	"context"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/types"
)

type (
	PrefName    string
	Filename    string
	BoolPtr     *bool
	FilenamePtr *string
	OutputPtr   *string
)

type (
	Context    = context.Context
	CmdResult  = cobrautil.CmdResult
	DomainName = macosutil.DomainName
	TypeName   = types.TypeName
	Name       = types.Name
	Code       = types.Code
)

type YAMLSpec interface {
	YAMLSpec()
}

type Spec interface {
	Spec()
}

type Printer interface {
	Print(i ...any)
	Println(i ...any)
	Printf(format string, i ...any)
}
