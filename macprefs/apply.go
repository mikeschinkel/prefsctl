package macprefs

type ApplyArgs struct {
	Filename Filename
}

func Apply(ctx Context, ptr Printer, args ApplyArgs) (err error) {
	panic("No 'apply' command implemented (yet)")
	return err
}
