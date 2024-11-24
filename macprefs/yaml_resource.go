package macprefs

import (
	"gopkg.in/yaml.v3"
)

type YAMLResource struct {
	APIVersion string       `yaml:"apiVersion"`
	Kind       ResourceKind `yaml:"kind"`
	metadata   YAMLMetadata `yaml:"metadata"`
	Spec       YAMLSpec     `yaml:"spec"`
}

func (r YAMLResource) YAML() string {
	b, err := yaml.Marshal(r)
	if err != nil {
		slog.Error("Failed to marshal YAML for macprefs.YAMLResource")
	}
	return string(b)
}
