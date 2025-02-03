package prefsyaml

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

type Pref struct {
	MetaData    *Metadata       `yaml:"-"`
	Name        PrefName        `yaml:"name"`
	Description Description     `yaml:"descr"`
	Type        PrefType        `yaml:"type"`
	Value       *yamlutil.Value `yaml:"value,omitempty"`
	Default     *yamlutil.Value `yaml:"default,omitempty"`
	Labels      []*LabelValue   `yaml:"labels,omitempty"`
	Added       OSVersion       `yaml:"added,omitempty"`
	Removed     OSVersion       `yaml:"removed,omitempty"`
}

// FilterableEntry implement profilemanifests.FilterableEntry interface
func (Pref) FilterableEntry() {}

func (yp Pref) PreferenceType() (pt macosutil.PreferenceType) {
	return macosutil.PreferenceType(yp.Type)
}

func (yp Pref) MacOSUtilPreference() (pref *macosutil.Preference) {
	return &macosutil.Preference{
		Domain: string(yp.MetaData.Domain),
		Name:   string(yp.Name),
		Value:  yp.Value.String(), // TODO: This will need to be revised when we support more than scalar types
	}
}
