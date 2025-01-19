package macprefs

import (
	"errors"
	"fmt"

	"github.com/mikeschinkel/prefsctl/applemdm"
	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logargs"
)

const (
	profileManifestsRepoURL = applemdm.ProfileManifestsRepoURL
	gitRepoParentDir        = applemdm.GitRepoParentDir
)

type UpdateArgs struct {
	Filename Filename
}

func Update(ctx Context, cfg config.Config, ptr Printer, args UpdateArgs) (result Result) {
	var pm *applemdm.ProfileManifests

	if ptr == nil {
		ptr = StandardPrinter{}
	}
	err := applemdm.EnsureGitRepo(profileManifestsRepoURL, gitRepoParentDir)
	if err != nil {
		goto end
	}
	ptr.Printf("\nProfileManifests Git repo (%s) updated in %s.", profileManifestsRepoURL, gitRepoParentDir)

	pm, err = LoadProfileManifests()
	if err != nil {
		goto end
	}
	ptr.Printf("\nProfileManifests loaded.")

	err = ProcessProfileManifests(pm)
	if err != nil {
		goto end
	}
	fmt.Printf("%v", pm)

	ptr.Printf("\nTODO: Update prefs")
	result = Result{Success: "Prefs updated."}
end:
	if err != nil {
		result = Result{Err: err}
	}
	return result
}

func ProcessProfileManifest(file *applemdm.EmbeddedFile) (pm *applemdm.PreferenceManifest, err error) {
	r, err := file.Reader()
	if err != nil {
		goto end
	}
	pm, err = applemdm.LoadPreferenceManifest(r)
	if err != nil {
		err = errors.Join(
			ErrFailedToGetFileInfo,
			fmt.Errorf("%s=%s", logargs.ManifestFile, file),
			err,
		)
		goto end
	}
	print(pm)
end:
	return pm, err
}

func ProcessProfileManifests(pm *applemdm.ProfileManifests) error {
	var errs errutil.MultiErr
	for _, file := range pm.Files() {
		if file.IsDir() {
			continue
		}
		pm, err := ProcessProfileManifest(file)
		if err != nil {
			errs.Add(err)
			goto end
		}
		print(pm)
	}
end:
	return errs.Err()
}

func LoadProfileManifests() (*applemdm.ProfileManifests, error) {
	pm := applemdm.NewProfileManifests()
	err := pm.Load()
	return pm, err
}
