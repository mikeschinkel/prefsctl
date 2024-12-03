package preftemplates

import (
	"bytes"
	"slices"
	"strings"
	"text/template"

	"github.com/mikeschinkel/prefsctl/macprefs/kvfilters"
)

const defaultsGoTemplateText = `func {{.Name}}PrefDefaults() DomainDefaults {
	return DomainDefaults{
		{{- range $defaults, $domain := .Domains}}
		"{{$domain}}": []DomainPrefs{
			{{- range $dflt := index $.Defaults $i}}
			{
				Name:      "{{$dflt.Name}}",
				Type:      "{{$dflt.TypeName}}",
				NoDefault: {{$dflt.NoDefault}},
				{{- if $dflt.ShowValue}}
				Default:    "{{$dflt.Value}}",
				{{- end}}
				Labels: Labels{
				{{- range $i, $label := .Labels}}
					&{{firstUp $label.Value}},
				{{- end}}
				},
			},
			{{- end}}
		},
		{{- end}}
	}
}`

type TypeName string
type OSVersion string
type DomainName string
type PrefName string
type Prefs struct {
	Domain    DomainName
	Name      PrefName
	Default   string // raw string value for default
	Labels    kvfilters.Labels
	NoDefault bool
}
type Default struct {
	Name      PrefName
	Value     string // raw string value for default
	Labels    kvfilters.Labels
	TypeName  TypeName
	NoDefault bool
	Domain    *Domain
}
type Domain struct {
	Domain   DomainName
	Defaults []*Default
	MacOS    *MacOS
}
type MacOS struct {
	Name          OSVersion
	Domains       []*Domain
	ShowValueFunc func(*Default) bool
}

func NewMacOS(name OSVersion, prefs []*Prefs) *MacOS {
	os := &MacOS{
		Name:    name,
		Domains: make([]*Domain, 0),
		ShowValueFunc: func(d *Default) bool {
			return true
		},
	}
	for _, pref := range prefs {
		f := func(d *Domain) bool {
			return d.Domain == pref.Domain
		}
		index := slices.IndexFunc(os.Domains, f)
		if index == -1 {
			os.Domains = append(os.Domains, &Domain{
				Domain:   pref.Domain,
				Defaults: make([]*Default, 0),
				MacOS:    os,
			})
		}
		os.Domains[index].Defaults = append(os.Domains[index].Defaults, &Default{
			Name:      pref.Name,
			Value:     pref.Default,
			Labels:    pref.Labels,
			NoDefault: pref.NoDefault,
			Domain:    os.Domains[index],
		})
	}
	slices.SortFunc(os.Domains, func(a, b *Domain) int {
		return strings.Compare(string(a.Domain), string(b.Domain))
	})
	for _, domain := range os.Domains {
		slices.SortFunc(domain.Defaults, func(a, b *Default) int {
			return strings.Compare(string(a.Name), string(b.Name))
		})
	}
	return os
}

// ShowValue returns true if there is no label "install-sets" — This means it is
// either a true default or we do not yet know for sure, but we do not that it is
// not determine exclusively by the person installing the OS.
//!slices.ContainsFunc(d.Labels, func(label *kvfilters.Label) bool {
//	return label.Value == macprefs.SetupSets.Value
//})

func (d *Default) ShowValue() bool {
	return d.Domain.MacOS.ShowValueFunc(d)
}

type Template interface {
	Generate() (string, error)
}

var _ Template = (*DefaultsGoTemplate)(nil)

type DefaultsGoTemplate struct {
	MacOS    *MacOS
	Defaults []*Default
}

var templateFuncs = template.FuncMap{
	"firstUp": func(s string) string {
		switch len(s) {
		case 0:
			return ""
		case 1:
			return strings.ToUpper(s)
		default:
			return strings.ToUpper(s[0:1]) + s[1:]
		}
	},
}

func (t DefaultsGoTemplate) Generate() (output string, err error) {
	var tmpl *template.Template
	var buf bytes.Buffer

	// Parse and execute tmpl
	tmpl, err = template.New("prefDefaults").
		Funcs(templateFuncs).
		Parse(defaultsGoTemplateText)
	if err != nil {
		goto end
	}

	err = tmpl.Execute(&buf, t.MacOS)
	if err != nil {
		goto end
	}
end:
	return buf.String(), nil
}
