package macprefs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/appinfo"
	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/gitutil"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/prefdefaults"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/profilemanifests"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

const (
	profileManifestsRepoURL = "https://github.com/ProfileManifests/ProfileManifests"
	gitRepoParentDir        = "/tmp/" + appinfo.Name
)

type UpdateArgs struct {
	Filename Filename
}

func Update(ctx Context, cfg config.Config, ptr Printer, args UpdateArgs) (result Result) {
	var pms *profilemanifests.ProfileManifests
	var yamlFile yamlutil.MultidocFile
	var repoDir string
	var pmsIterator yamlutil.EntriesIterator

	if ptr == nil {
		ptr = StandardPrinter{}
	}
	repoDir, err := gitutil.EnsureGitRepo(profileManifestsRepoURL, gitRepoParentDir)
	if err != nil {
		goto end
	}
	ptr.Printf("\nProfileManifests Git repo (%s) updated in %s.", profileManifestsRepoURL, gitRepoParentDir)

	pms = profilemanifests.New(repoDir)
	err = pms.LoadFiles()
	if err != nil {
		goto end
	}
	ptr.Printf("\nProfileManifests loaded.")

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
	err = prefdefaults.WritePrefsDefaultsFile(cfg, yamlFile.String())
	if err != nil {
		goto end
	}
	result = Result{
		Success: fmt.Sprintf("Pref defaults updated in %s", cfg.OtherFilepath(appinfo.PrefDefaultsFile)),
	}
end:
	if err != nil {
		result = Result{Err: err}
	}
	return result
}

var osVersionNumber macosutil.VersionNumber

func entryFilter(entry yamlutil.FilterableEntry) (include bool) {
	if osVersionNumber == "" {
		var err error
		osVersionNumber, err = macosutil.GetVersionNumber()
		if err != nil {
			panicf("ERROR: Failed to get OS version number")
		}
	}
	include = true
	switch t := entry.(type) {
	case prefsyaml.Resource:
		switch {
		case t.MetaData.Domain == ".GlobalPreferences":
			// Do nothing. Include it.
		case !strings.HasPrefix(string(t.MetaData.Domain), "com.apple"):
			include = false
			goto end
		}
		if !osSupported(osVersionNumber, t.MetaData.Added, t.MetaData.Removed) {
			include = false
			goto end
		}
	case *prefsyaml.Default:
		if !osSupported(osVersionNumber, t.Added, t.Removed) {
			include = false
			goto end
		}
		if !strings.HasPrefix(string(t.Name), "Payload") {
			goto end
		}
		switch t.Name[len("Payload"):] {
		case "Description", "DisplayName", "Content":
			fallthrough
		case "Type", "UUID", "Version", "Organization":
			fallthrough
		case "Identifier", "Identification":
			fallthrough
		case "CertificateUUID", "CertificateFileName":
			include = false
		}
		goto end
	default:
		panicf("entryFilter: unknown entry type %T", entry)
	}
end:
	return include
}

func osSupported(version, added, removed macosutil.VersionNumber) bool {
	return prefsyaml.SupportsOSVersion(
		prefsyaml.OSVersion(version),
		prefsyaml.OSVersion(added),
		prefsyaml.OSVersion(removed),
	)
}
