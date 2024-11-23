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

type PrefId string
type DomainPrefDefaults map[Domain][]*PrefDefault

type Filename string

type FilenamePtr *string
type OutputPtr *string

type YAMLFile struct {
	Preferences []YAMLDomain `yaml:"preferences"`
}

type YAMLDomain struct {
	Name       string            `yaml:"-"`
	Properties map[string]string `yaml:",inline"`
}
