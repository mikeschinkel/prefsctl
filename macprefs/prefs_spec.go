package macprefs

var _ Spec = (*PrefsSpec)(nil)

type PrefsSpec struct {
	Prefs []*YAMLPref `yaml:"preferences"`
}

func (PrefsSpec) Spec() {}
