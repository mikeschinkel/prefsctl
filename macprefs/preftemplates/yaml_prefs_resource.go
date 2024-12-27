package preftemplates

import (
	_ "embed"
	"fmt"

	"github.com/mikeschinkel/prefsctl/macosutil"
	"gopkg.in/yaml.v3"
)

type YAMLPrefsResource struct {
	APIVersion APIVersion   `yaml:"apiVersion"`
	KindName   KindName     `yaml:"kind"`
	MetaData   YAMLMetadata `yaml:"metadata"`
	Spec       YAMLSpec     `yaml:"spec"`
}

type Identifier string

func (r *YAMLPrefsResource) Id() Identifier {
	// TODO: What would Kubernetes do to format this?
	return Identifier(fmt.Sprintf("%s:%s", r.KindName, r.MetaData.Domain))
}

func (r *YAMLPrefsResource) YAML() string {
	y, err := yaml.Marshal(r)
	if err != nil {
		panicf("Failed to marshal resource '%s' to YAML", r.Id())
	}
	return string(y)
}

type YAMLMetadata struct {
	Domain    DomainName `yaml:"domain"`
	OSVersion OSVersion  `yaml:"macos"`
}

type YAMLSpec struct {
	Prefs []YAMLPref `yaml:"preferences"`
}

type YAMLPref struct {
	Name  PrefName `yaml:"name"`
	Value string   `yaml:"value"`
}

func NewYAMLPrefsResources(domains []*Domain) []YAMLPrefsResource {
	resources := make([]YAMLPrefsResource, len(domains))
	for i, domain := range domains {
		resources[i] = NewYAMLPrefsResource(domain)
	}
	return resources
}

var osVersion OSVersion

func NewYAMLPrefsResource(domain *Domain) YAMLPrefsResource {
	if osVersion == "" {
		osVersion = OSVersion(macosutil.MustGetVersionCode())
	}
	prefs := make([]YAMLPref, len(domain.Prefs))
	for j, pref := range domain.Prefs {
		prefs[j] = YAMLPref{
			Name:  pref.Name,
			Value: pref.Value,
		}
	}
	return YAMLPrefsResource{
		KindName: "pref",
		MetaData: YAMLMetadata{
			Domain:    domain.Name,
			OSVersion: osVersion,
		},
		Spec: YAMLSpec{
			Prefs: prefs,
		},
	}
}
