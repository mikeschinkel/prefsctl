package macprefs

type YAMLResource struct {
	APIVersion string       `yaml:"apiVersion"`
	Kind       ResourceKind `yaml:"kind"`
	metadata   YAMLMetadata `yaml:"metadata"`
	Spec       YAMLSpec     `yaml:"spec"`
}
