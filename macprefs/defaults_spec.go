package macprefs

import (
	"github.com/mikeschinkel/prefsctl/prefdefaults"
)

var _ Spec = (*YAMLDefaultsSpec)(nil)

type YAMLDefaultsSpec struct {
	Prefs []prefdefaults.PrefDefault `yaml:"preferences"`
}

func (YAMLDefaultsSpec) Spec() {}
