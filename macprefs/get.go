package macprefs

import (
	"errors"
	"os"

	"github.com/mikeschinkel/prefsctl/cliutil"
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

func getText(ctx Context, ptr Printer, args GetArgs) (result cliutil.Result, err error) {
	state := NewPrefsState()
	err = state.Query(QueryArgs{
		Filters: QueryFilters(),
	})
	if err != nil {
		err = errors.Join(ErrFailedToQueryPrefState, err)
		goto end
	}
	state.Describe(os.Stdout)
	//result = NewResult(state)
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
