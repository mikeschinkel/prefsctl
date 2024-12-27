package preftemplates

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/mikeschinkel/prefsctl/macosutil"
)

//go:embed templates/prefs_yaml.template
var prefsYAMLTemplateText TemplateText

var _ TemplateData = (*PrefsYAMLTemplate)(nil)

type PrefsYAMLTemplate struct {
	APIVersion APIVersion
	OSVersion  OSVersion
	Domains    []YAMLDomain
}

type YAMLDomain struct {
	KindName KindName
	MetaData YAMLMetadata
	Spec     YAMLSpec
}

type YAMLMetadata struct {
	Domain    DomainName
	OSVersion OSVersion
}

type YAMLSpec struct {
	Prefs []Pref
}

func NewYAMLDomains(domains []*Domain) []YAMLDomain {
	osVersion := OSVersion(macosutil.MustGetVersionCode())
	yamlDomains := make([]YAMLDomain, len(domains))
	for i, domain := range domains {
		prefs := make([]Pref, len(domain.Prefs))
		for j, pref := range domain.Prefs {
			prefs[j] = Pref{
				Domain:  pref.Domain,
				Name:    pref.Name,
				Default: pref.Default,
				Labels:  pref.Labels,
			}
		}
		yamlDomains[i] = YAMLDomain{
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
	return yamlDomains
}

type PrefsYAMLTemplateArgs struct {
	APIVersion APIVersion
	OSVersion  OSVersion
	Domains    []*Domain
}

func NewPrefsYAMLTemplate(args PrefsYAMLTemplateArgs) *PrefsYAMLTemplate {
	return &PrefsYAMLTemplate{
		APIVersion: args.APIVersion,
		OSVersion:  args.OSVersion,
		Domains:    NewYAMLDomains(args.Domains),
	}
}
func (t PrefsYAMLTemplate) Generate() (output string, err error) {
	return Generate("prefs", prefsYAMLTemplateText, func(tmpl *template.Template, buf *bytes.Buffer) error {
		return tmpl.Execute(buf, t)
	})
}
