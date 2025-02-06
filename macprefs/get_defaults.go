package macprefs

import (
	"errors"
	"fmt"
	"os"

	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/gitutil"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/profilemanifests"
	"github.com/mikeschinkel/prefsctl/slogutil"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

func GetYAMLDocumentFromPrefsDomain(entry yamlutil.FilterableEntry, filter yamlutil.EntryFilterFunc) (yd yamlutil.Document, err error) {
	yr, ok := entry.(*prefsyaml.Resource)
	if !ok {
		err = slogutil.PanicInTest(slog, "ERROR: Expected entry to be a *prefsyaml.Resource",
			logargs.EntryType,
			fmt.Sprintf("%T", entry),
		)
		goto end
	}
	yd, err = yr.YAMLDocument()
	if err != nil {
		goto end
	}
end:
	return yd, err
}

func GetDefaults(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) (result Result) {
	if ptr == nil {
		ptr = StandardPrinter{}
	}

	// Defaults should include all defaults, including the unchanged ones.
	args.IncludeUnchanged = true

	format := *GlobalFlags.Output
	if format == "" {
		format = string(TXTFormat)
	}
	switch OutputFormat(format) {
	case YAMLFormat:
		result = getDefaultsYAML(ctx, cfg, ptr, args)
	case JSONFormat:
		result = GetDefaultsJSON(ctx, cfg, ptr, args)
	case TXTFormat:
		fallthrough
	default:
		result = getDefaultsText(ctx, cfg, ptr, args)
	}
	return result
}

func getDefaultsYAML(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) Result {
	var pms *profilemanifests.ProfileManifests
	var repoDir string
	var pmsIterator yamlutil.EntriesIterator
	var yamlFile yamlutil.MultidocFile

	_, repoDir = gitutil.GetGitRepoPathAndDir(profileManifestsRepoURL, gitRepoParentDir)
	pms = profilemanifests.New(repoDir)
	err := pms.LoadFilesForDomains(args.Domains)
	if err != nil {
		goto end
	}
	slog.Info("ProfileManifests loaded.")

	pmsIterator = profilemanifests.Iterator(pms)
	yamlFile, err = yamlutil.BuildMultidocFile(
		pmsIterator,
		entryFilter,
		profilemanifests.GetYAMLDocumentFromProfileManifest,
	)
	if pmsIterator.Err != nil || err != nil {
		err = errors.Join(err, pmsIterator.Err)
		goto end
	}
	ptr.Print(yamlFile.String())
end:
	return Result{
		Success: "Defaults YAML generated.",
		Err:     err,
	}
}
func GetDefaultsJSON(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) Result {
	return Result{
		Err: errors.New("JSON output not yet implemented"),
	}
}

func getDefaultsText(ctx Context, cfg config.Config, ptr Printer, args QueryArgs) Result {
	domains, err := QueryPrefDomains(ctx, cfg, args)
	if err != nil {
		goto end
	}
	domains.Describe(os.Stdout)
	ptr.Println("\nUnsupported Types:")
	for ut := range unsupportedTypes {
		ptr.Printf("— %s\n", ut)
	}
end:
	return Result{
		Success: "",
		Err:     err,
	}
}
