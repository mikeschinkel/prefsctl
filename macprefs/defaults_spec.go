package macprefs

var _ Spec = (*DefaultsSpec)(nil)

type DefaultsSpec struct {
	Prefs []PrefDefault `yaml:"preferences"`
}

func (DefaultsSpec) Spec() {}
