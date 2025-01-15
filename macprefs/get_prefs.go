package macprefs

import (
	"errors"

	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
)

type (
	OSVersion = preftemplates.OSVersion
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

func newDomains(domains []*PrefsDomain) []*preftemplates.Domain {
	dd := make([]*preftemplates.Domain, len(domains))
	for i, domain := range domains {
		prefs := make([]*preftemplates.Pref, len(domain.Prefs()))
		for j, pref := range domain.Prefs() {
			prefs[j] = &preftemplates.Pref{
				Domain: preftemplates.DomainName(pref.Domain),
				Name:   preftemplates.PrefName(pref.Name),
				Value:  pref.Value(),
			}
		}
		dd[i] = &preftemplates.Domain{
			Name:  preftemplates.DomainName(domain.Name()),
			Prefs: prefs,
			MacOS: nil,
		}
	}
	return dd
}

func getPrefsYAML(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) (result Result) {
	var osVersion OSVersion
	var resources []preftemplates.YAMLPrefsResource

	domains, err := QueryPrefDomains(ctx, cfg, args)
	if err != nil {
		goto end
	}

	osVersion = OSVersion(macosutil.MustGetVersionCode())
	resources = preftemplates.NewYAMLPrefsResources(newDomains(domains.domains))
	for _, resource := range resources {
		ptr.Println("---")
		resource.APIVersion = LatestAPIVersion
		resource.MetaData.OSVersion = osVersion
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

func getPrefsGo(ctx Context, ptr Printer, args QueryArgs) (result Result) {
	return Result{
		Err: errors.New("get prefs --output=go not (yet?) implemented"),
	}
}
