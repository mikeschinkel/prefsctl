package macprefs

import (
	"errors"
	"fmt"
	"io"
	"slices"

	"github.com/mikeschinkel/prefsctl/appinfo"
	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/gitutil"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/profilemanifests"
)

const (
	profileManifestsRepoURL = "https://github.com/ProfileManifests/ProfileManifests"
	gitRepoParentDir        = "/tmp/" + appinfo.Name

	MacOSPlatform = "macOS"
)

type UpdateArgs struct {
	Filename Filename
}

func Update(ctx Context, cfg config.Config, ptr Printer, args UpdateArgs) (result Result) {
	var pms *profilemanifests.ProfileManifests

	if ptr == nil {
		ptr = StandardPrinter{}
	}
	err := gitutil.EnsureGitRepo(profileManifestsRepoURL, gitRepoParentDir)
	if err != nil {
		goto end
	}
	ptr.Printf("\nProfileManifests Git repo (%s) updated in %s.", profileManifestsRepoURL, gitRepoParentDir)

	pms, err = LoadProfileManifests()
	if err != nil {
		goto end
	}
	ptr.Printf("\nProfileManifests loaded.")

	err = ProcessProfileManifests(pms)
	if err != nil {
		goto end
	}

	ptr.Printf("\nTODO: Update prefs")
	result = Result{Success: "Prefs updated."}
end:
	if err != nil {
		result = Result{Err: err}
	}
	return result
}

func LoadProfileManifests() (*profilemanifests.ProfileManifests, error) {
	pms := profilemanifests.New()
	err := pms.Load()
	return pms, err
}

func ProcessProfileManifests(pms *profilemanifests.ProfileManifests) error {
	var errs errutil.MultiErr
	for _, file := range pms.Files() {
		if file.IsDir() {
			continue
		}
		err := ProcessProfileManifest(file)
		if err != nil {
			errs.Add(err)
			goto end
		}
	}
end:
	return errs.Err()
}

func ProcessProfileManifest(file *profilemanifests.EmbeddedFile) (err error) {
	var pm *profilemanifests.ProfileManifest
	var r io.Reader
	var pd *DefaultsDomain

	r, err = file.Reader()
	if err != nil {
		goto end
	}
	pm, err = profilemanifests.LoadProfileManifest(r)
	if err != nil {
		err = errors.Join(
			ErrFailedToGetFileInfo,
			fmt.Errorf("%s=%s", logargs.ManifestFile, file),
			err,
		)
		goto end
	}

	if !slices.Contains(pm.Platforms, MacOSPlatform) {
		// Not for macOS
		goto end
	}
	pd = NewDefaultsDomainFromProfileManifest(pm)

	fmt.Printf("%v\n", pd)

end:
	return err
}

func NewDefaultsDomainFromProfileManifest(pm *profilemanifests.ProfileManifest) *DefaultsDomain {
	dn := macosutil.DomainName(pm.Domain)
	if len(dn) != 0 && dn[0] == '.' {
		dn = dn[1:]
	}
	pd := NewDefaultsDomain(dn, macosutil.GetProcessToKill(dn))
	for _, subKey := range pm.Subkeys {
		if subKey.IsPayloadKey() {
			continue
		}
		if !subKey.SupportsOSVersion() {
			continue
		}
		pd.AddDefault(NewPrefDefault(dn, PrefName(subKey.Name), &PrefDefaultOpts{
			Kind:          subKey.Kind(),
			SupportedIn:   OSVersion(subKey.MacOSMin),
			UnsupportedIn: "", // TODO Add support for this
		}))
	}
	return &pd
}
