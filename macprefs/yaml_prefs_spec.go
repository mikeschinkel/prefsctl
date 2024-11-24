package macprefs

var _ YAMLSpec = (*YAMLPrefsSpec)(nil)

type YAMLPrefsSpec struct {
	Prefs []YAMLPref `yaml:"prefs"`
}

func (YAMLPrefsSpec) YAMLSpec() {}
