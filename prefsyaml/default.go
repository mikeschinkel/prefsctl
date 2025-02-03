package prefsyaml

import (
	"strings"

	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

type Default struct {
	MetaData    *Metadata               `yaml:"-"`
	Name        PrefName                `yaml:"name"`
	Description Description             `yaml:"description,omitempty"`
	Value       *yamlutil.Value         `yaml:"default"`
	Options     []*yamlutil.Value       `yaml:"options,omitempty"`
	Labels      []*LabelValue           `yaml:"labels,omitempty"`
	Type        PrefType                `yaml:"type"`
	Added       macosutil.VersionNumber `yaml:"added,omitempty"`
	Removed     macosutil.VersionNumber `yaml:"removed,omitempty"`
}

func (Default) FilterableEntry() {}

type DefaultOpts struct {
	Description Description     `yaml:"description"`
	Value       *yamlutil.Value // raw string value for default
	Options     []*yamlutil.Value
	Labels      []*LabelValue
	Type        PrefType
	Added       macosutil.VersionNumber
	Removed     macosutil.VersionNumber
}

func NewDefault(name PrefName, opts DefaultOpts) Default {
	return Default{
		Name:        name,
		Description: opts.Description,
		Value:       yamlutil.NewValue(opts.Value),
		Options:     opts.Options,
		Labels:      opts.Labels,
		Type:        opts.Type,
		Added:       opts.Added,
		Removed:     opts.Removed,
	}
}

func (yd Default) TypeName() string {
	typ, _ := strings.CutSuffix(string(yd.Type), "Type")
	return typ
}
