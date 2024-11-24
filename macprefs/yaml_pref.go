package macprefs

// YAMLPref represents a preference with its Domain and name
type YAMLPref struct {
	Name     string       `yaml:"name"`
	Value    string       `yaml:"value"`    // raw string value
	Metadata YAMLMetadata `yaml:"metadata"` // type of the value
}
