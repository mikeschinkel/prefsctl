package preftemplates

import (
	"bytes"
	"strings"
	"text/template"
)

const defaultsGoTemplateName = "prefDefaults"
const defaultsGoTemplateText = `package prefdefaults

//goland:noinspection SpellCheckingInspection
func {{.OSVersion}}PrefDefaults() DomainDefaults {
	return DomainDefaults{
		{{- range $i, $domain := .Domains}}
		"{{$domain.Name}}": DomainPrefs{
			{{- range $i, $dflt := $domain.Defaults}}
				"{{$dflt.Name}}": DomainPref{
					Type:      "{{$dflt.TypeName}}",
					NoDefault: {{$dflt.NoDefault}},
					{{- if $.ShowPrefDefault $dflt }}
					Default:   "{{$dflt.Value}}",
					{{- end}}
					Labels: Labels{
					{{- range $i, $label := .Labels}}
						{{firstUp $label.Value.String}},
					{{- end}}
					},
				},
			{{- end}}
		},
		{{- end}}
	}	
}`

var _ Template = (*DefaultsGoTemplate)(nil)

type DefaultsGoTemplate struct {
	OSVersion     OSVersion
	Domains       []*Domain
	ShowValueFunc func(*Default) bool
}

func (t DefaultsGoTemplate) ShowPrefDefault(d *Default) bool {
	return d.Value != "" && t.ShowValueFunc(d)
}

func NewDefaultsGoTemplate(name OSVersion, domains []*Domain) *DefaultsGoTemplate {
	return &DefaultsGoTemplate{
		OSVersion: OSVersion(strings.ToLower(string(name))),
		Domains:   domains,
		ShowValueFunc: func(d *Default) bool {
			return true
		},
	}
}

func (t DefaultsGoTemplate) Generate() (output string, err error) {
	var tmpl *template.Template
	var buf bytes.Buffer

	// Parse and execute tmpl
	tmpl, err = template.New(defaultsGoTemplateName).
		Funcs(templateFuncs).
		Parse(defaultsGoTemplateText)
	if err != nil {
		goto end
	}

	err = tmpl.Execute(&buf, t)
	if err != nil {
		goto end
	}
end:
	return buf.String(), nil
}
