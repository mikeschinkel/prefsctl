package macprefs

import (
	"os"

	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
)

func GetDefaults(ctx Context, ptr Printer, args QueryArgs) (err error) {
	if ptr == nil {
		ptr = StandardPrinter{}
	}

	// Defaults should include all defaults, including the unchanged ones.
	args.IncludeUnchanged = true

	format := *GlobalFlags.Output
	if format == "" {
		format = string(TXTFormat)
	}
	switch OutputFormat(format) {
	case YAMLFormat:
		err = getDefaultsYAML(ctx, ptr, args)
	case JSONFormat:
		err = GetDefaultsJSON(ctx, ptr, args)
	case GoFormat:
		err = getDefaultsGo(ctx, ptr, args)
	case TXTFormat:
		fallthrough
	default:
		err = getDefaultsText(ctx, ptr, args)
	}
	return err
}

func getDefaultsGo(ctx Context, ptr Printer, args QueryArgs) (err error) {
	var tmpl *preftemplates.DefaultsGoTemplate
	var output string

	code, err := macosutil.VersionCode()
	domains, err := QueryPrefDomains(ctx, args)
	if err != nil {
		goto end
	}
	tmpl = preftemplates.NewDefaultsGoTemplate(
		preftemplates.OSVersion(code),
		domains.TemplateDomains(TemplateDomainsArgs{
			UseCurrent: args.UseCurrent,
		}),
	)
	tmpl.ShowValueFunc = func(d *preftemplates.Default) bool {
		return d.Labels.HasLabel(&UserManaged)
	}
	output, err = tmpl.Generate()
	if err != nil {
		goto end
	}
	ptr.Print(output)
end:
	return err
}

func getDefaultsText(ctx Context, ptr Printer, args QueryArgs) (err error) {
	domains, err := QueryPrefDomains(ctx, args)
	if err != nil {
		goto end
	}
	domains.Describe(os.Stdout)
	ptr.Println("\nUnsupported Types:")
	for ut := range unsupportedTypes {
		ptr.Printf("— %s\n", ut)
	}
end:
	return err
}

func getDefaultsYAML(ctx Context, ptr Printer, args QueryArgs) (err error) {
	ptr.Print("YAML output not implemented")
	return err
}
func GetDefaultsJSON(ctx Context, ptr Printer, args QueryArgs) (err error) {
	ptr.Print("JSON output not implemented")
	return err
}
