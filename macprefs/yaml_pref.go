package macprefs

// YAMLPref represents a preference with its Domain and name
type YAMLPref struct {
	Name  string   `yaml:"name"`
	Value string   `yaml:"value"` // raw string value
	Kind  PrefType `yaml:"kind"`  // type of the value
}
