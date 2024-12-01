package macprefs

import (
	"errors"
	"fmt"
	"os"

	"github.com/mikeschinkel/prefsctl/cliutil"
	"github.com/mikeschinkel/prefsctl/macprefs/filters"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type GetArgs struct {
	Filename Filename
	Output   OutputFormat
}

func Get(ctx Context, ptr Printer, args GetArgs) (result cliutil.Result, err error) {
	switch args.Output {
	case YAMLFormat:
		result, err = getYAML(ctx, ptr, args)
	case JSONFormat:
		result, err = GetJSON(ctx, ptr, args)
	case GoFormat:
		result, err = getGo(ctx, ptr, args)
	case TxtFormat:
		fallthrough
	default:
		result, err = getText(ctx, ptr, args)
	}
	return result, err
}

func toDomains(group []filters.Group) PrefDomains {
	return sliceconv.Func(group, func(g filters.Group) *PrefsDomain {
		return g.(*PrefsDomain)
	})
}

func getText(ctx Context, ptr Printer, args GetArgs) (result cliutil.Result, err error) {
	var groupFilters, allFilters []filters.Filter
	var domains PrefDomains
	var filtered []filters.Group

	domains, err = GetPrefDomains()
	if err != nil {
		goto end
	}

	groupFilters, err = filters.GroupsQueryFilters()
	if err != nil {
		goto end
	}

	filtered, err = filters.QueryGroups(filters.QueryArgs{
		Filters: groupFilters,
		Groups:  domains.AsFilterGroups(),
	})
	if err != nil {
		err = errors.Join(ErrFailedToQueryGroups, err)
		goto end
	}

	domains = toDomains(filtered)
	err = domains.Initialize()
	if err != nil {
		goto end
	}

	allFilters, err = filters.QueryFilters()
	if err != nil {
		goto end
	}

	filtered, err = filters.Query(filters.QueryArgs{
		Filters: allFilters,
		Groups:  domains.AsFilterGroups(),
	})
	if err != nil {
		err = errors.Join(ErrFailedToQueryPrefState, err)
		goto end
	}

	//domains, err = PrefDomainsFromFiltersGroups(filtered)
	//if err != nil {
	//	goto end
	//}
	domains = toDomains(filtered)
	domains.Sort()
	domains.Describe(os.Stdout)
	fmt.Println("\nUnsupported Types:")
	for ut := range unsupportedTypes {
		fmt.Printf("— %s\n", ut)
	}
	result = NewResult("Success")

end:
	return result, err
}

func getGo(ctx Context, ptr Printer, args GetArgs) (result cliutil.Result, err error) {
	return result, err
}

func getYAML(ctx Context, ptr Printer, args GetArgs) (result cliutil.Result, err error) {
	return result, err
}
func GetJSON(ctx Context, ptr Printer, args GetArgs) (result cliutil.Result, err error) {
	return result, err
}
