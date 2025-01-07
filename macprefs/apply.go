package macprefs

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mikeschinkel/prefsctl/logargs"
)

type ApplyArgs struct {
	Filename Filename
}

func Apply(ctx Context, ptr Printer, args ApplyArgs) (err error) {
	if ptr == nil {
		ptr = StandardPrinter{}
	}

	// Maybe get output format from command line
	format := *GlobalFlags.Output
	if format == "" {
		// But if not output format specified on command line, use file extension.
		format = strings.ToLower(filepath.Ext(string(args.Filename)))
		if format[0] == '.' {
			format = format[1:]
		}
		if format == "yml" {
			format = string(YAMLFormat)
		}
	}
	switch OutputFormat(format) {
	case YAMLFormat:
		err = applyYAML(ctx, ptr, args)
	case JSONFormat:
		err = ApplyJSON(ctx, ptr, args)
	case GoFormat:
		fallthrough
	case TXTFormat:
		fallthrough
	default:
		err = errors.Join(
			ErrInvalidOutputFormat,
			fmt.Errorf("valid formats include: '%s', '%s'", YAMLFormat, JSONFormat),
			fmt.Errorf("%s=%s", logargs.SpecifiedFormat, format),
		)
	}
	return err
}

func applyYAML(ctx Context, ptr Printer, args ApplyArgs) (err error) {
	return fmt.Errorf("apply YAML output not implemented")
}
func ApplyJSON(ctx Context, ptr Printer, args ApplyArgs) (err error) {
	return fmt.Errorf("apply JSON output not implemented")
}
