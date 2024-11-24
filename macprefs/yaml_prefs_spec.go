package macprefs

var _ YAMLSpec = (*YAMLPrefsSpec)(nil)

type YAMLPrefsSpec struct {
	Prefs []*YAMLPref `yaml:"preferences"`
}

func (YAMLPrefsSpec) YAMLSpec() {}
