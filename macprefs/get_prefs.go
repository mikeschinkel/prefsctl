package macprefs

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
)

type (
	OSVersion = preftemplates.OSVersion
)

func GetPrefs(ctx Context, args QueryArgs) (err error) {
	if args.Printer == nil {
		args.Printer = StandardPrinter{}
	}

	switch OutputFormat(*GlobalFlags.Output) {
	case YAMLFormat:
		err = getPrefsYAML(ctx, args)
	case JSONFormat:
		err = getPrefsJSON(ctx, args)
	case GoFormat:
		err = getPrefsGo(ctx, args)
	case TXTFormat:
		fallthrough
	default:
		err = getPrefsText(ctx, args)
	}
	return err
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

func getPrefsYAML(ctx Context, args QueryArgs) (err error) {
	var osVersion OSVersion
	var ptr Printer
	var resources []preftemplates.YAMLPrefsResource

	domains, err := QueryPrefDomains(ctx, args)
	if err != nil {
		goto end
	}

	ptr = args.Printer
	osVersion = OSVersion(macosutil.MustGetVersionCode())
	resources = preftemplates.NewYAMLPrefsResources(newDomains(domains.domains))
	for _, resource := range resources {
		ptr.Println("---")
		resource.APIVersion = LatestAPIVersion
		resource.MetaData.OSVersion = osVersion
		ptr.Println(resource.YAML())
	}
end:
	return err
}

func getPrefsJSON(ctx Context, args QueryArgs) (err error) {
	args.Printer.Print("JSON output not implemented")
	return err
}

func getPrefsText(ctx Context, args QueryArgs) (err error) {
	args.Printer.Print("TXT output not implemented")
	return err
}

func getPrefsGo(ctx Context, args QueryArgs) (err error) {
	args.Printer.Print("Go output not implemented")
	return err
}
