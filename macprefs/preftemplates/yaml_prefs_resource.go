package preftemplates

import (
	_ "embed"
	"fmt"
	"io"
	"os"

	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/stdlibex"
	"gopkg.in/yaml.v3"
)

type YAMLPrefsResource struct {
	APIVersion APIVersion   `yaml:"apiVersion"`
	KindName   KindName     `yaml:"kind"`
	MetaData   YAMLMetadata `yaml:"metadata"`
	Spec       YAMLPrefSpec `yaml:"spec"`
}

type Identifier string

func (r *YAMLPrefsResource) Id() Identifier {
	// TODO: What would Kubernetes do to format this?
	return Identifier(fmt.Sprintf("%s:%s", r.KindName, r.MetaData.Domain))
}

func LoadYAMLPrefsResources(filename string) (resources []YAMLPrefsResource, err error) {
	var decoder *yaml.Decoder

	file, err := os.Open(filename)
	if err != nil {
		goto end
	}
	defer stdlibex.MustClose(file)

	resources = make([]YAMLPrefsResource, 0, 1)
	decoder = yaml.NewDecoder(file)
	for {
		var resource YAMLPrefsResource
		err = decoder.Decode(&resource)
		if err == io.EOF {
			err = nil
			goto end
		}
		if err != nil {
			goto end
		}
		for i := range resource.Spec.Prefs {
			resource.Spec.Prefs[i].MetaData = &resource.MetaData
		}
		resources = append(resources, resource)
	}
end:
	return resources, err
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

type YAMLPrefSpec struct {
	Prefs []YAMLPref `yaml:"preferences"`
}

func NewYAMLPrefSpec() YAMLPrefSpec {
	return YAMLPrefSpec{
		Prefs: make([]YAMLPref, 0),
	}
}

type YAMLPref struct {
	MetaData *YAMLMetadata `yaml:"-"`
	Name     PrefName      `yaml:"name"`
	Type     PrefType      `yaml:"type"`
	Value    string        `yaml:"value,omitempty"`
	Default  string        `yaml:"default,omitempty"`
	Labels   []LabelValue  `yaml:"labels,omitempty"`
}

func (yp YAMLPref) PreferenceType() (pt macosutil.PreferenceType) {
	return macosutil.PreferenceType(yp.Type)
}

func (yp YAMLPref) MacOSUtilPreference() (pref *macosutil.Preference) {
	return &macosutil.Preference{
		Domain: string(yp.MetaData.Domain),
		Name:   string(yp.Name),
		Value:  yp.Value,
	}
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
		Spec: YAMLPrefSpec{
			Prefs: prefs,
		},
	}
}
