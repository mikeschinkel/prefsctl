package macprefs

type YAMLDomain struct {
	Name       string            `yaml:"-"`
	Properties map[string]string `yaml:",inline"`
}
