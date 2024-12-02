package macprefs

import (
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
	case TXTFormat:
		fallthrough
	default:
		result, err = getText(ctx, ptr, args)
	}
	return result, err
}

func getText(ctx Context, ptr Printer, args GetArgs) (result cliutil.Result, err error) {
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
