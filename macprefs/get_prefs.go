package macprefs

import (
	"errors"

	"github.com/mikeschinkel/prefsctl/config"
	prefsyaml2 "github.com/mikeschinkel/prefsctl/prefsyaml"
)

type (
	OSVersion = prefsyaml2.OSVersion
)

func GetPrefs(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) (result Result) {
	if ptr == nil {
		ptr = StandardPrinter{}
	}

	format := *GlobalFlags.Output
	if format == "" {
		format = string(TXTFormat)
	}
	switch OutputFormat(format) {
	case YAMLFormat:
		result = getPrefsYAML(ctx, cfg, ptr, args)
	case JSONFormat:
		result = getPrefsJSON(ctx, cfg, ptr, args)
	case TXTFormat:
		fallthrough
	default:
		result = getPrefsText(ctx, cfg, ptr, args)
	}
	return result
}

func newDomains(domains []*PrefsDomain) []*prefsyaml2.YAMLDomain {
	dd := make([]*prefsyaml2.YAMLDomain, len(domains))
	for i, domain := range domains {
		prefs := make([]*prefsyaml2.YAMLPrefLite, len(domain.Prefs()))
		for j, pref := range domain.Prefs() {
			prefs[j] = &prefsyaml2.YAMLPrefLite{
				Domain: prefsyaml2.DomainName(pref.Domain),
				Name:   prefsyaml2.PrefName(pref.Name),
				Value:  pref.Value(),
			}
		}
		dd[i] = &prefsyaml2.YAMLDomain{
			Name:  prefsyaml2.DomainName(domain.Name()),
			Prefs: prefs,
		}
	}
	return dd
}

func getPrefsYAML(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) (result Result) {
	var resources []prefsyaml2.YAMLPrefsResource

	domains, err := QueryPrefDomains(ctx, cfg, args)
	if err != nil {
		goto end
	}

	resources = prefsyaml2.NewYAMLPrefsResources(newDomains(domains.domains))
	for _, resource := range resources {
		ptr.Println("---")
		resource.APIVersion = LatestAPIVersion
		ptr.Println(resource.YAML())
	}
end:
	return Result{Err: err}
}

func getPrefsJSON(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) (result Result) {
	return Result{
		Err: errors.New("get prefs --output=json not (yet?) implemented"),
	}
}

func getPrefsText(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) (result Result) {
	return Result{
		Err: errors.New("get prefs --output=txt not (yet?) implemented"),
	}
}
