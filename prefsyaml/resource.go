package prefsyaml

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/mikeschinkel/prefsctl/appinfo"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
	"github.com/mikeschinkel/prefsctl/stdlibex"
	"github.com/mikeschinkel/prefsctl/yamlutil"
	"gopkg.in/yaml.v3"
)

type Resource struct {
	APIVersion APIVersion `yaml:"apiVersion"`
	KindName   KindName   `yaml:"kind"`
	MetaData   Metadata   `yaml:"metadata"`
	Spec       Spec       `yaml:"spec"`
}

var osVersion OSVersion

type ResourceOpts struct {
	Description Description
	Process     ProcessName
	APIVersion  APIVersion
	Added       macosutil.VersionNumber
	Removed     macosutil.VersionNumber
}

func NewResource(kind KindName, domain DomainName, opts ResourceOpts) Resource {
	var spec Spec

	switch macpreflabels.ResourceKind(kind) {
	case macpreflabels.PrefsKind:
		spec = Spec{
			Prefs: make([]Pref, 0),
		}
	case macpreflabels.DefaultsKind:
		spec = Spec{
			Defaults: make([]Default, 0),
		}
	}
	if opts.APIVersion == "" {
		opts.APIVersion = appinfo.LatestAPIVersion
	}
	return Resource{
		APIVersion: opts.APIVersion,
		KindName:   KindName(kind),
		Spec:       spec,
		MetaData: Metadata{
			Domain:      domain,
			Process:     opts.Process,
			Description: opts.Description,
			Added:       opts.Added,
			Removed:     opts.Removed,
		},
	}
}

func (r Resource) FilterableEntry() {
}

func (r Resource) Id() Identifier {
	// TODO: What would Kubernetes do to format this?
	return Identifier(fmt.Sprintf("%s:%s", r.KindName, r.MetaData.Domain))
}

func LoadPrefsResources(filename string) (resources []Resource, err error) {
	var decoder *yaml.Decoder

	file, err := os.Open(filename)
	if err != nil {
		goto end
	}
	defer stdlibex.MustClose(file)

	resources = make([]Resource, 0, 1)
	decoder = yaml.NewDecoder(file)
	for {
		var resource Resource
		err = decoder.Decode(&resource)
		if err == io.EOF {
			err = nil
			goto end
		}
		if err != nil {
			err = errors.Join(
				err,
				ErrFailedToDecodeYAML,
				fmt.Errorf("%s=%s", errutil.FilenameErrKey, filename),
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

func (r Resource) YAMLDocument() (yd yamlutil.Document, err error) {
	var y bytes.Buffer
	enc := yaml.NewEncoder(&y)
	enc.SetIndent(2)
	err = enc.Encode(r)
	if err != nil {
		err = errors.Join(
			ErrFailedToMarshalToYAML,
			errutil.ErrArg(DocumentTypeErrKey, "prefsyaml.Resource"),
			errutil.ErrArg(ResourceIdErrKey, r.Id()),
			err,
		)
		goto end
	}
	yd = yamlutil.Document(y.String())
end:
	return yd, err
}
