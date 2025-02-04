package prefsyaml

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
)

type Spec struct {
	Prefs    []Pref    `yaml:"preferences,omitempty"`
	Defaults []Default `yaml:"defaults,omitempty"`
}

func NewSpec() Spec {
	return Spec{
		Prefs:    make([]Pref, 0),
		Defaults: make([]Default, 0),
	}
}

// Initialize propagates the metadata passed in to the appropriate slice in Spec;
// either each of Props or each of Defaults AND if sets the Kind value.
func (s Spec) Initialize(md *Metadata) {
	switch macpreflabels.ResourceKind(md.KindName) {
	case macpreflabels.PrefsKind:
		for i, pref := range s.Prefs {
			pref.Kind = macosutil.GetKind(macosutil.PreferenceType(pref.Type))
			pref.MetaData = md
			s.Prefs[i] = pref
		}
	case macpreflabels.DefaultsKind:
		for i, def := range s.Defaults {
			def.Kind = macosutil.GetKind(macosutil.PreferenceType(def.Type))
			def.MetaData = md
			s.Defaults[i] = def
		}
	}
}
