package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
)

type Domain struct {
	Domain      macosutil.DomainName    `yaml:"domain"`
	Defaults    []*PrefDefault          `yaml:"defaults"`
	Description string                  `yaml:"description"`
	Process     macosutil.ProcessName   `yaml:"process.omitempty"`
	Added       macosutil.VersionNumber `yaml:"added.omitempty"`
	Removed     macosutil.VersionNumber `yaml:"removed.omitempty"`
}
type DomainOpts struct {
	Domain      macosutil.DomainName
	Description string
	Process     macosutil.ProcessName
	Added       macosutil.VersionNumber
	Removed     macosutil.VersionNumber
}

func NewDomain(opts DomainOpts) *Domain {
	return &Domain{
		Domain:      opts.Domain,
		Description: opts.Description,
		Process:     opts.Process,
		Added:       opts.Added,
		Removed:     opts.Removed,
		Defaults:    make([]*PrefDefault, 0),
	}
}

func newDomainFromYAMLResource(resource YAMLResource) *Domain {
	meta := resource.MetaData
	defaults := resource.Spec.Defaults
	domain := NewDomain(DomainOpts{
		Domain:      macosutil.DomainName(meta.Domain),
		Description: meta.Description,
		Process:     resource.MetaData.Process,
		Added:       meta.Added,
		Removed:     meta.Removed,
	})
	domain.Defaults = make([]*PrefDefault, len(defaults))
	for i, def := range resource.Spec.Defaults {
		pd := convertToPrefDefault(DomainName(meta.Domain), &def)
		if pd == nil {
			continue
		}
		domain.Defaults[i] = &PrefDefault{
			Domain:      domain.Domain,
			Name:        PrefName(def.Name),
			Description: def.Description,
			Default:     def.Value.String(),
			Kind:        def.Kind,
			Type:        TypeName(def.Type),
			labels:      GetLabelsByValues(def.Labels),
			Added:       def.Added,
			Removed:     def.Removed,
		}
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

func GetDomains(cfg config.Config) (domains []*Domain, err error) {
	f := GetDefaultsFunc()
	prefDefaults, err := f(cfg)
	if err != nil {
		goto end
	}
	domains = prefDefaults.Domains
end:
	return domains, err
}
