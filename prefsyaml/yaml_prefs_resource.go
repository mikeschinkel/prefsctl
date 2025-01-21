package prefsyaml

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/mikeschinkel/prefsctl/logargs"
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

func NewYAMLPrefsResources(domains []*YAMLDomain) []YAMLPrefsResource {
	resources := make([]YAMLPrefsResource, len(domains))
	for i, domain := range domains {
		resources[i] = NewYAMLPrefsResource(domain, domain.Opts)
	}
	return resources
}

var osVersion OSVersion

type PrefsResourceOpts struct {
	KillOnApply macosutil.ProcessName
}

func NewYAMLPrefsResource(domain *YAMLDomain, opts PrefsResourceOpts) YAMLPrefsResource {
	prefs := make([]YAMLPref, len(domain.Prefs))
	for j, pref := range domain.Prefs {
		prefs[j] = YAMLPref{
			Name:  pref.Name,
			Value: YAMLValue{Value: pref.Value},
		}
	}
	return YAMLPrefsResource{
		KindName: "pref",
		MetaData: YAMLMetadata{
			Domain:      domain.Name,
			KillOnApply: opts.KillOnApply,
		},
		Spec: YAMLPrefSpec{
			Prefs: prefs,
		},
	}
}

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
			err = errors.Join(
				err,
				ErrFailedToDecodeYAML,
				fmt.Errorf("%s=%s", logargs.Filename, filename),
			)
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
