package macprefs

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
)

type ApplyArgs struct {
	Filename Filename
}

func Apply(ctx Context, ptr Printer, args ApplyArgs) (result Result) {
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
		result = applyYAML(ctx, ptr, args)
	case JSONFormat:
		result = ApplyJSON(ctx, ptr, args)
	case GoFormat:
		fallthrough
	case TXTFormat:
		fallthrough
	default:
		result = Result{
			Err: errors.Join(
				ErrInvalidOutputFormat,
				fmt.Errorf("valid formats include: '%s', '%s'", YAMLFormat, JSONFormat),
				fmt.Errorf("%s=%s", logargs.SpecifiedFormat, format),
			),
		}
	}
	return result
}

func applyYAML(ctx Context, ptr Printer, args ApplyArgs) (result Result) {
	var domain string
	var prefs []*macosutil.Preference
	var success strings.Builder
	var specPrefs []preftemplates.YAMLPref
	var aaf func() error
	//var afterFunc func() error

	result = Result{}

	resource, err := preftemplates.LoadYAMLPrefsResource(string(args.Filename))
	if err != nil {
		result.Err = err
		goto end
	}
	specPrefs = resource.Spec.Prefs
	prefs = make([]*macosutil.Preference, len(specPrefs))
	domain = string(resource.MetaData.Domain)
	switch len(prefs) {
	case 0:
		result.Err = errors.Join(
			ErrNoPrefsFoundInResourceSpec,
			fmt.Errorf("%s=%s", logargs.Filename, args.Filename),
			err,
		)
		goto end
	case 1:
		prefs = []*macosutil.Preference{
			specPrefs[0].MacOSUtilPreference(),
		}
	default:
		success.WriteString(fmt.Sprintf("Prefs applied.\n\tDomain '%s':\n", domain))
		for i, pref := range specPrefs {
			prefs[i] = pref.MacOSUtilPreference()
			success.WriteString("\t— ")
			success.WriteString(renderPreference(prefs[i]))
			success.WriteByte('\n')
		}
		result.Success = success.String()
	}
	err = macosutil.ApplyPreferences(domain, prefs)
	if err != nil {
		result = Result{Err: err}
		goto end
	}
	aaf, err = GetAfterApplyFunc(DomainName(domain))
	if aaf == nil {
		goto end
	}
	err = aaf()
	if err != nil {
		goto end
	}
end:
	return result
}
func ApplyJSON(ctx Context, ptr Printer, args ApplyArgs) (result Result) {
	return Result{
		Err: errors.New("apply JSON output not implemented"),
	}
}

func renderPreference(p *macosutil.Preference) (value string) {
	value = p.Value
	if !p.Valid() {
		value = "<invalid>"
	}
	return fmt.Sprintf("%s: %s", p.Name, value)
}
