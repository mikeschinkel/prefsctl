package prefdefaults

import (
	"reflect"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
)

type Domain struct {
	Domain   macosutil.DomainName    `yaml:"domain"`
	Defaults []*PrefDefault          `yaml:"defaults"`
	Process  macosutil.ProcessName   `yaml:"process.omitempty"`
	Added    macosutil.VersionNumber `yaml:"added.omitempty"`
	Removed  macosutil.VersionNumber `yaml:"removed.omitempty"`
}

func newDomainFromYAMLResource(resource YAMLResource) *Domain {
	meta := resource.MetaData
	defaults := resource.Spec.Defaults
	domain := &Domain{
		Domain:   macosutil.DomainName(meta.Domain),
		Process:  resource.MetaData.Process,
		Defaults: make([]*PrefDefault, len(defaults)),
		Added:    meta.Added,
		Removed:  meta.Removed,
	}
	for _, def := range resource.Spec.Defaults {
		pd := convertToPrefDefault(DomainName(meta.Domain), &def)
		if pd == nil {
			continue
		}
		//kind := macosutil.GetPrefKind(def.Type, def.Value.String())
		//kind, typ := macosutil.GetPrefKindAndType(kind, def.Type, def.Value.String())
		kind := reflect.Invalid
		typ := TypeName("")
		domain.Defaults = append(domain.Defaults, &PrefDefault{
			Domain:       domain.Domain,
			Name:         PrefName(def.Name),
			DefaultValue: def.Value.String(),
			Kind:         kind,
			typeName:     typ,
			labels:       GetLabelsByValues(def.Labels),
			Added:        def.Added,
			Removed:      def.Removed,
		})
	}
	return domain
}

func GetLabelsByValues(values []*prefsyaml.LabelValue) *Labels {
	labels := make([]*kvfilters.Label, len(values))
	for i, value := range values {
		labels[i] = macpreflabels.GetLabelByValue(
			macpreflabels.LabelValue(*value),
		)
	}
	return kvfilters.NewLabels(labels...)
}
