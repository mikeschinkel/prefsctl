package prefsyaml

type YAMLPrefSpec struct {
	Prefs []YAMLPref `yaml:"preferences"`
}

func NewYAMLPrefSpec() YAMLPrefSpec {
	return YAMLPrefSpec{
		Prefs: make([]YAMLPref, 0),
	}
}
