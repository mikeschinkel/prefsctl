package macprefs

import (
	"errors"
	"os"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type GetDefaultsArgs struct {
	Printer Printer
}

func GetDefaults(ctx Context, args GetDefaultsArgs) (err error) {
	if args.Printer == nil {
		args.Printer = StandardPrinter{}
	}

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

func retrieveDefaults(ctx Context, args GetDefaultsArgs) (domains *PrefDomains, err error) {
	var nameFilters, valueFilters []kvfilters.Filter
	var filtered []kvfilters.Group
	var pds []*PrefDefault

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
	pds = domains.UserManagedPrefDefaults()
	noop(pds)

	filtered, err = kvfilters.Query(kvfilters.QueryArgs{
		Filters: nameFilters,
		Groups:  domains.ToFiltersGroups(),
		Labels:  []*kvfilters.Label{&UserManaged},
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
		Filters:     valueFilters,
		Groups:      domains.ToFiltersGroups(),
		Labels:      []*kvfilters.Label{&UserManaged},
		OmitEmpty:   true,
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

func getDefaultsText(ctx Context, args GetDefaultsArgs) (err error) {
	domains, err := retrieveDefaults(ctx, args)
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

func getDefaultsGo(ctx Context, args GetDefaultsArgs) (err error) {
	var tmpl *preftemplates.DefaultsGoTemplate
	var output string

	code, err := macosutil.VersionCode()
	domains, err := retrieveDefaults(ctx, args)
	if err != nil {
		goto end
	}
	tmpl = preftemplates.NewDefaultsGoTemplate(
		preftemplates.OSVersion(code),
		domains.TemplateDomains(),
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

func getDefaultsYAML(ctx Context, args GetDefaultsArgs) (err error) {
	return err
}
func GetDefaultsJSON(ctx Context, args GetDefaultsArgs) (err error) {
	return err
}
