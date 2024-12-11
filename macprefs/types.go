package macprefs

import (
	"context"
)

type (
	Context = context.Context
)

type YAMLSpec interface {
	YAMLSpec()
}
type Spec interface {
	Spec()
}

type PrefName string
type TypeName string

type PrefDefaultsMap map[PrefName]*PrefDefault
type DomainPrefDefaults map[DomainName]PrefDefaultsMap

type Filename string

type FilenamePtr *string
type OutputPtr *string

type YAMLFile struct {
	Preferences []YAMLDomain `yaml:"preferences"`
}

type Printer interface {
	Print(i ...any)
	Println(i ...any)
	Printf(format string, i ...any)
}
