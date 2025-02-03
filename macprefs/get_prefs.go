package macprefs

import (
	"errors"

	"github.com/mikeschinkel/prefsctl/appinfo"
	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

type (
	OSVersion  = prefsyaml.OSVersion
	APIVersion = prefsyaml.APIVersion
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

func newYAMLResources(kind prefsyaml.KindName, domains []*PrefsDomain, opts YAMLOpts) []*prefsyaml.Resource {
	yrs := make([]*prefsyaml.Resource, len(domains))
	for i, domain := range domains {
		yrs[i] = domain.GetYAMLResource(kind, opts)
	}
	return yrs
}

func getPrefsYAML(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) (result Result) {
	var resources []*prefsyaml.Resource
	var ymf yamlutil.MultidocFile
	var yd yamlutil.Document
	var errs errutil.MultiErr

	domains, err := QueryPrefDomains(ctx, cfg, args)
	if err != nil {
		goto end
	}

	resources = newYAMLResources(
		prefsyaml.KindName(macpreflabels.PrefsKind),
		domains.domains,
		YAMLOpts{
			UseValueForDefault: false,
			APIVersion:         appinfo.LatestAPIVersion,
		},
	)
	ymf = yamlutil.NewMultidocFile()
	for _, resource := range resources {
		yd, err = resource.YAMLDocument()
		if err != nil {
			errs.Add(err)
			goto end
		}
		ymf.AddDocument(yd)
	}
	ptr.Print(ymf.String())

end:
	return Result{Err: errs.Err()}
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
