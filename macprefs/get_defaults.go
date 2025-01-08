package macprefs

import (
	"errors"
	"os"

	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
)

func GetDefaults(ctx Context, ptr Printer, args QueryArgs) (result Result) {
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
		result = getDefaultsYAML(ctx, ptr, args)
	case JSONFormat:
		result = GetDefaultsJSON(ctx, ptr, args)
	case GoFormat:
		result = getDefaultsGo(ctx, ptr, args)
	case TXTFormat:
		fallthrough
	default:
		result = getDefaultsText(ctx, ptr, args)
	}
	return result
}

func getDefaultsGo(ctx Context, ptr Printer, args QueryArgs) Result {
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
	return Result{
		Success: "",
		Err:     err,
	}
}

func getDefaultsText(ctx Context, ptr Printer, args QueryArgs) Result {
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
	return Result{
		Success: "",
		Err:     err,
	}
}

func getDefaultsYAML(ctx Context, ptr Printer, args QueryArgs) Result {
	return Result{
		Err: errors.New("YAML output not implemented"),
	}
}
func GetDefaultsJSON(ctx Context, ptr Printer, args QueryArgs) Result {
	return Result{
		Err: errors.New("JSON output not implemented"),
	}
}
