package prefsyaml

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
