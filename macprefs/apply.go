package macprefs

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
)

type ApplyArgs struct {
	Filename Filename
}

func Apply(ctx Context, cfg config.Config, ptr Printer, args ApplyArgs) (result Result) {
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
		result = applyYAML(ctx, cfg, ptr, args)
	case JSONFormat:
		result = ApplyJSON(ctx, cfg, ptr, args)
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

func applyYAMLResource(ctx Context, cfg config.Config, resource prefsyaml.YAMLPrefsResource, args ApplyArgs) (result Result) {
	var success strings.Builder
	var aaf func() error
	var err error

	preparePref := func(yamlPref prefsyaml.YAMLPref) (pref *macosutil.Preference) {
		pref = yamlPref.MacOSUtilPreference()
		success.WriteString("\t— ")
		success.WriteString(renderPreference(pref))
		success.WriteByte('\n')
		return pref
	}

	specPrefs := resource.Spec.Prefs
	prefs := make([]*macosutil.Preference, len(specPrefs))
	domain := string(resource.MetaData.Domain)
	switch len(prefs) {
	case 0:
		result.Err = errors.Join(
			ErrNoPrefsFoundInResourceSpec,
			fmt.Errorf("%s=%s", logargs.Filename, args.Filename),
			err,
		)
		goto end
	case 1:
		success.WriteString(fmt.Sprintf("\tDomain '%s':\n", domain))
		prefs = []*macosutil.Preference{preparePref(specPrefs[0])}
		result.Success = success.String()
	default:
		success.WriteString(fmt.Sprintf("\tDomain '%s':\n", domain))
		for i, pref := range specPrefs {
			prefs[i] = preparePref(pref)
		}
		result.Success = success.String()
	}
	err = macosutil.ApplyPreferences(domain, prefs)
	if err != nil {
		result = Result{Err: err}
		goto end
	}
	aaf, err = GetAfterApplyFunc(cfg, DomainName(domain))
	if err != nil {
		result = Result{Err: err}
		goto end
	}
	if aaf == nil {
		goto end
	}
	err = aaf()
	if err != nil {
		result = Result{Err: err}
		goto end
	}
end:
	return result
}

func applyYAML(ctx Context, cfg config.Config, ptr Printer, args ApplyArgs) (result Result) {
	var errs errutil.MultiErr

	successes := []string{fmt.Sprintf("Prefs applied.\n")}

	resources, err := prefsyaml.LoadYAMLPrefsResources(string(args.Filename))
	if err != nil {
		errs.Add(err)
		goto end
	}
	for _, resource := range resources {
		itemResult := applyYAMLResource(ctx, cfg, resource, args)
		if itemResult.Err != nil {
			errs.Add(itemResult.Err)
			continue
		}
		successes = append(successes, itemResult.Success)
	}
end:
	if !errs.IsError() && len(successes) != 0 {
		result.Success = strings.Join(successes, "")
	}
	result.Err = errs.Err()
	return result
}

func ApplyJSON(ctx Context, cfg config.Config, ptr Printer, args ApplyArgs) (result Result) {
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
