package macprefs

import (
	"errors"

	"github.com/mikeschinkel/prefsctl/config"
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
		yrs[i] = domain.GetPrefsYAMLResource(kind, opts)
	}
	return yrs
}

func getPrefsYAML(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) Result {
	var domainsIterator yamlutil.EntriesIterator
	var yamlFile yamlutil.MultidocFile

	domains, err := QueryPrefDomains(ctx, cfg, args)
	if err != nil {
		goto end
	}
	domainsIterator = domains.DomainsIterator(PrefsKind, YAMLOpts{
		APIVersion: LatestAPIVersion,
	})
	yamlFile, err = yamlutil.BuildMultidocFile(
		domainsIterator,
		entryFilter,
		GetYAMLDocumentFromPrefsDomain,
	)
	if domainsIterator.Err != nil || err != nil {
		err = errors.Join(err, domainsIterator.Err)
		goto end
	}
	ptr.Print(yamlFile.String())
end:
	return Result{
		Success: "Defaults YAML generated.",
		Err:     err,
	}
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
