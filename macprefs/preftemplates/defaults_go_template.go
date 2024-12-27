package preftemplates

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed templates/defaults_go.template
var defaultsGoTemplateText TemplateText

var _ TemplateData = (*DefaultsGoTemplate)(nil)

type DefaultsGoTemplate struct {
	OSVersion     OSVersion
	Domains       []*Domain
	ShowValueFunc func(*Default) bool
}

func (t DefaultsGoTemplate) ShowPrefDefault(d *Default) bool {
	//return d.Value != "" && t.ShowValueFunc(d)
	return t.ShowValueFunc(d)
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
	return Generate("prefDefaults", defaultsGoTemplateText, func(tmpl *template.Template, buf *bytes.Buffer) error {
		return tmpl.Execute(buf, t)
	})
}
