package macprefs

type GetArgs struct {
	Filename Filename
	Output   OutputFormat
}

func Get(ctx Context, ptr Printer, args GetArgs) (err error) {
	switch args.Output {
	case YAMLFormat:
		err = getYAML(ctx, ptr, args)
	case JSONFormat:
		err = GetJSON(ctx, ptr, args)
	case GoFormat:
		err = getGo(ctx, ptr, args)
	case TXTFormat:
		fallthrough
	default:
		err = getText(ctx, ptr, args)
	}
	return err
}

func getText(ctx Context, ptr Printer, args GetArgs) (err error) {
	return err
}

func getGo(ctx Context, ptr Printer, args GetArgs) (err error) {
	return err
}

func getYAML(ctx Context, ptr Printer, args GetArgs) (err error) {
	return err
}
func GetJSON(ctx Context, ptr Printer, args GetArgs) (err error) {
	return err
}
