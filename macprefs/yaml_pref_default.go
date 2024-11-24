package macprefs

// YAMLPrefDefault represents a preference's default value
type YAMLPrefDefault struct {
	Name         string       `yaml:"name"`
	DefaultValue string       `yaml:"default"`  // raw string value for default
	WhatSets     WhatSetsType `yaml:"whatSets"` // determine what sets the value
	Kind         PrefType     `yaml:"kind"`     // type of the value
	Options      []string     `yaml:"options"`
}
