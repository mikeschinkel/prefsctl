package macprefs

// YAMLPrefDefault represents a preference's default value
type YAMLPrefDefault struct {
	Name         string       `yaml:"name"`
	DefaultValue string       `yaml:"default"` // raw string value for default
	Metadata     YAMLMetadata `yaml:"metadata"`
}
