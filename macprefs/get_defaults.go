package macprefs

import (
	"os"

	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
)

func GetDefaults(ctx Context, args GenerateArgs) (err error) {
	if args.Printer == nil {
		args.Printer = StandardPrinter{}
	}

	// Defaults should include all defaults, including the unchanged ones.
	args.IncludeUnchanged = true

	switch OutputFormat(*GlobalFlags.Output) {
	case YAMLFormat:
		err = getDefaultsYAML(ctx, args)
	case JSONFormat:
		err = GetDefaultsJSON(ctx, args)
	case GoFormat:
		err = getDefaultsGo(ctx, args)
	case TXTFormat:
		fallthrough
	default:
		err = getDefaultsText(ctx, args)
	}
	return err
}

func getDefaultsGo(ctx Context, args GenerateArgs) (err error) {
	var tmpl *preftemplates.DefaultsGoTemplate
	var output string

	code, err := macosutil.VersionCode()
	domains, err := retrievePrefDomains(ctx, args)
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
	args.Printer.Print(output)
end:
	return err
}

func getDefaultsText(ctx Context, args GenerateArgs) (err error) {
	domains, err := retrievePrefDomains(ctx, args)
	if err != nil {
		goto end
	}
	domains.Describe(os.Stdout)
	args.Printer.Println("\nUnsupported Types:")
	for ut := range unsupportedTypes {
		args.Printer.Printf("— %s\n", ut)
	}
end:
	return err
}

func getDefaultsYAML(ctx Context, args GenerateArgs) (err error) {
	args.Printer.Print("YAML output not implemented")
	return err
}
func GetDefaultsJSON(ctx Context, args GenerateArgs) (err error) {
	args.Printer.Print("JSON output not implemented")
	return err
}
