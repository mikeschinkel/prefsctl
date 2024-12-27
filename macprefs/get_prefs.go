package macprefs

import (
	"errors"
	"fmt"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type (
	OSVersion = preftemplates.OSVersion
)

type GenerateArgs struct {
	Printer          Printer
	OmitEmpty        bool
	UseCurrent       bool
	IncludeUnchanged bool
}

func (args *GenerateArgs) PrinterOutput() string {
	return args.Printer.(fmt.Stringer).String()
}

func GetPrefs(ctx Context, args GenerateArgs) (err error) {
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

func retrievePrefsDomains(ctx Context, args GenerateArgs) (domains *PrefDomains, err error) {
	var nameFilters, valueFilters []kvfilters.Filter
	var filtered []kvfilters.Group

	toDomains := func(domains *PrefDomains, group []kvfilters.Group) []*PrefsDomain {
		return sliceconv.Func(group, func(g kvfilters.Group) *PrefsDomain {
			return g.(*PrefsDomain)
		})
	}

	domains, err = RetrievePrefDomains()
	if err != nil {
		goto end
	}
	err = domains.Initialize()
	if err != nil {
		goto end
	}
	err = domains.RetrievePrefs(RetrievePrefArgs{
		IgnoreMissingDomains: true,
	})
	if err != nil {
		goto end
	}
	nameFilters, err = QueryFiltersForTargets(kvfilters.Groups, kvfilters.Keys)
	if err != nil {
		goto end
	}

	filtered, err = kvfilters.Query(kvfilters.QueryArgs{
		Groups:    domains.ToFiltersGroups(),
		Filters:   nameFilters,
		Labels:    []*kvfilters.Label{&UserManaged},
		OmitEmpty: args.OmitEmpty,
	})
	if err != nil {
		err = errors.Join(ErrFailedToQueryGroups, err)
		goto end
	}

	domains.domains = toDomains(domains, filtered)

	err = domains.RetrievePrefValues()
	if err != nil {
		goto end
	}

	valueFilters, err = QueryFiltersForTargets(kvfilters.Values, kvfilters.KeyValues)
	if err != nil {
		goto end
	}

	filtered, err = kvfilters.Query(kvfilters.QueryArgs{
		Groups:      domains.ToFiltersGroups(),
		Filters:     valueFilters,
		Labels:      []*kvfilters.Label{&UserManaged},
		OmitEmpty:   args.OmitEmpty,
		OmitInvalid: true,
	})

	if err != nil {
		err = errors.Join(ErrFailedToQueryPrefState, err)
		goto end
	}

	domains.domains = toDomains(domains, filtered)
	domains.Sort()
end:
	return domains, err
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

func getPrefsYAML(ctx Context, args GenerateArgs) (err error) {
	var osVersion OSVersion
	var ptr Printer
	var resources []preftemplates.YAMLPrefsResource

	domains, err := retrievePrefsDomains(ctx, args)
	if err != nil {
		goto end
	}

	ptr = args.Printer
	osVersion = OSVersion(macosutil.MustGetVersionCode())
	resources = preftemplates.NewYAMLPrefsResources(newDomains(domains.domains))
	for _, resource := range resources {
		ptr.Println("---")
		resource.MetaData.OSVersion = osVersion
		ptr.Println(resource.YAML())
	}
end:
	return err
}

func getPrefsJSON(ctx Context, args GenerateArgs) (err error) {
	args.Printer.Print("JSON output not implemented")
	return err
}

func getPrefsText(ctx Context, args GenerateArgs) (err error) {
	args.Printer.Print("TXT output not implemented")
	return err
}

func getPrefsGo(ctx Context, args GenerateArgs) (err error) {
	args.Printer.Print("Go output not implemented")
	return err
}
