package macprefs

import (
	"errors"
	"fmt"
	"os"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type GetDefaultsArgs struct {
	Dummy string
}

func GetDefaults(ctx Context, ptr Printer, args GetDefaultsArgs) (result cobrautil.Result, err error) {
	switch OutputFormat(*GlobalFlags.Output) {
	case YAMLFormat:
		result, err = getDefaultsYAML(ctx, ptr, args)
	case JSONFormat:
		result, err = GetDefaultsJSON(ctx, ptr, args)
	case GoFormat:
		result, err = getDefaultsGo(ctx, ptr, args)
	case TXTFormat:
		fallthrough
	default:
		result, err = getDefaultsText(ctx, ptr, args)
	}
	return result, err
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
	err = domains.RetrievePrefs()
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

func getDefaultsText(ctx Context, ptr Printer, args GetDefaultsArgs) (result cobrautil.Result, err error) {
	domains, err := retrieveDefaults(ctx, args)
	if err != nil {
		goto end
	}
	domains.Describe(os.Stdout)
	fmt.Println("\nUnsupported Types:")
	for ut := range unsupportedTypes {
		fmt.Printf("— %s\n", ut)
	}
	result = NewResult("Success")

end:
	return result, err
}

func getDefaultsGo(ctx Context, ptr Printer, args GetDefaultsArgs) (result cobrautil.Result, err error) {
	var tmpl *preftemplates.DefaultsGoTemplate
	var output string

	code, err := macosutils.VersionCode()
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
	ptr.Print(output)
	result = NewResult("Success")
end:
	if err != nil {
		result = NewResult("Failure!")
	}
	return result, err
}

func getDefaultsYAML(ctx Context, ptr Printer, args GetDefaultsArgs) (result cobrautil.Result, err error) {
	return result, err
}
func GetDefaultsJSON(ctx Context, ptr Printer, args GetDefaultsArgs) (result cobrautil.Result, err error) {
	return result, err
}
