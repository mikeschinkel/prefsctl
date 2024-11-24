package macprefs

import (
	"gopkg.in/yaml.v3"
)

type YAMLMetadata struct {
	Domain string `yaml:"domain"`
	OSName string `yaml:"macos"`
}

func (m YAMLMetadata) YAML() string {
	b, err := yaml.Marshal(m)
	if err != nil {
		slog.Error("Failed to marshal YAML for macprefs.YAMLMetadata")
	}
	return string(b)
}
