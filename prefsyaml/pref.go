package prefsyaml

import (
	"reflect"

	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

type Pref struct {
	MetaData    *Metadata       `yaml:"-"`
	Name        PrefName        `yaml:"name"`
	Type        PrefType        `yaml:"type"`
	Description string          `yaml:"descr,omitempty"`
	Value       *yamlutil.Value `yaml:"value,omitempty"`
	Default     *yamlutil.Value `yaml:"default,omitempty"`
	Labels      []*LabelValue   `yaml:"labels,omitempty"`
	Added       OSVersion       `yaml:"added,omitempty"`
	Removed     OSVersion       `yaml:"removed,omitempty"`
	Kind        reflect.Kind    `yaml:"-"`
}

func (p *Pref) SetMetadata(md *Metadata) {
	p.MetaData = md
}

// FilterableEntry implement profilemanifests.FilterableEntry interface
func (*Pref) FilterableEntry() {}

func (p *Pref) PreferenceType() (pt macosutil.PreferenceType) {
	return macosutil.PreferenceType(p.Type)
}

func (p *Pref) MacOSUtilPreference() (pref *macosutil.Preference) {
	return &macosutil.Preference{
		Domain: string(p.MetaData.Domain),
		Name:   string(p.Name),
		Value:  p.Value.String(), // TODO: This will need to be revised when we support more than scalar types
	}
}
