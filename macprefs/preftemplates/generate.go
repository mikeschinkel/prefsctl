package preftemplates

import (
	"bytes"
	"text/template"
)

type TemplateExecFunc func(tmpl *template.Template, buf *bytes.Buffer) error

func Generate(name TemplateName, text TemplateText, execFunc TemplateExecFunc) (output string, err error) {
	var tmpl *template.Template
	var buf bytes.Buffer

	// Parse and execute tmpl
	tmpl, err = template.New(string(name)).
		Funcs(templateFuncs).
		Parse(string(text))
	if err != nil {
		goto end
	}

	err = execFunc(tmpl, &buf)
	if err != nil {
		goto end
	}
end:
	return buf.String(), err
}
