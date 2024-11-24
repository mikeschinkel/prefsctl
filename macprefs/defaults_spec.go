package macprefs

var _ Spec = (*YAMLDefaultsSpec)(nil)

type YAMLDefaultsSpec struct {
	Prefs []PrefDefault `yaml:"preferences"`
}

func (YAMLDefaultsSpec) Spec() {}
