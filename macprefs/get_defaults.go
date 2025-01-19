package macprefs

import (
	"errors"
	"os"

	"github.com/mikeschinkel/prefsctl/config"
)

func GetDefaults(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) (result Result) {
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
		result = getDefaultsYAML(ctx, cfg, ptr, args)
	case JSONFormat:
		result = GetDefaultsJSON(ctx, cfg, ptr, args)
	case TXTFormat:
		fallthrough
	default:
		result = getDefaultsText(ctx, cfg, ptr, args)
	}
	return result
}

func getDefaultsYAML(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) Result {
	//var tmpl *prefsyaml.DefaultsGoTemplate
	//var output string
	//
	//code, err := macosutil.VersionCode()
	domains, err := QueryPrefDomains(ctx, cfg, args)
	if err != nil {
		goto end
	}
	ptr.Print(domains.DefaultsYAML(YAMLOpts{
		UseValueForDefault: args.UseCurrent,
	}))
end:
	return Result{
		Success: "Defaults YAML generated.",
		Err:     err,
	}
}
func GetDefaultsJSON(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) Result {
	return Result{
		Err: errors.New("JSON output not implemented"),
	}
}

func getDefaultsText(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) Result {
	domains, err := QueryPrefDomains(ctx, cfg, args)
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
