package macprefs

import (
	"gopkg.in/yaml.v3"
)

type Resource struct {
	APIVersion string       `yaml:"apiVersion"`
	Kind       ResourceKind `yaml:"kind"`
	metadata   Metadata     `yaml:"metadata"`
	Spec       Spec         `yaml:"spec"`
}

func (m Resource) YAML() string {
	b, err := yaml.Marshal(m)
	if err != nil {
		slog.Error("Failed to marshal YAML for macprefs.YAMLResource")
	}
	return string(b)
}
