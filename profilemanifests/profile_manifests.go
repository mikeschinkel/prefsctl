package profilemanifests

import (
	"bytes"
	"embed"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/filesutil"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/prefdefaults"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

const (
	profileManifestsDir = "data/ProfileManifests/Manifests"
)

type (
	Content     = string
	EntryFiles  = []filesutil.FileEntry
	DirEntryMap map[string]EntryFiles
)

//go:embed data/ProfileManifests
var profileManifests embed.FS

type EmbeddedFS struct {
	embed.FS
	RootDir string
}

// GetEmbeddedFS returns the embedded file system object
func GetEmbeddedFS() EmbeddedFS {
	return EmbeddedFS{
		FS:      profileManifests,
		RootDir: profileManifestsDir,
	}
}

var (
	profileManifestFiles   EntryFiles
	profileManifestFileMap DirEntryMap
)

type ProfileManifests struct {
	repoDir string
	files   EntryFiles
	fileMap DirEntryMap
	pms     []*ProfileManifest
}

func (pms *ProfileManifests) GetFileEntries() []filesutil.FileEntry {
	return pms.files
}

func New(repoDir string) *ProfileManifests {
	return &ProfileManifests{
		repoDir: repoDir,
	}
}

func (pms *ProfileManifests) sortFiles() {
	slices.SortFunc(pms.files, func(a, b filesutil.FileEntry) int {
		switch {
		case a.Filepath() < b.Filepath():
			return -1
		case a.Filepath() > b.Filepath():
			return 1
		}
		return 0
	})
}

func (pms *ProfileManifests) Files() EntryFiles {
	pms.sortFiles()
	return pms.files
}

func (pms *ProfileManifests) LoadFiles() error {
	return pms.LoadFilesForDomains(nil)
}
func (pms *ProfileManifests) LoadFilesForDomains(domains []DomainName) error {
	files, err := filesutil.ListFilesRecursive(profileManifests, profileManifestsDir)
	if err != nil {
		goto end
	}
	pms.files = make(EntryFiles, 0, len(files))
	pms.fileMap = make(DirEntryMap, len(pms.files))
	for _, file := range files {
		fn := file.Name()
		domain := DomainName(strings.TrimSuffix(fn, filepath.Ext(fn)))
		if !prefdefaults.IncludeDomain(domains, domain) {
			continue
		}
		fp := file.Filepath()
		_, ok := pms.fileMap[fp]
		if ok {
			slog.Warn("Unexpected duplicated preference domain", logargs.PrefsDomain, filepath.Base(file.Name()))
		} else {
			pms.fileMap[fp] = make(EntryFiles, 0, 1)
		}
		pms.files = append(pms.files, file)
		pms.fileMap[fp] = append(pms.fileMap[fp], file)
	}
end:
	return err
}

//	func NewDefaultsDomainFromProfileManifest(pm *profilemanifests.ProfileManifest) *DefaultsDomain {
//		dn := macosutil.DomainName(pm.Domain)
//		if len(dn) != 0 && dn[0] == '.' {
//			dn = dn[1:]
//		}
//		pd := NewDefaultsDomain(dn, macosutil.GetProcessToKill(dn))
//		for _, subKey := range pm.Subkeys {
//			if subKey.IsPayloadKey() {
//				continue
//			}
//			if !subKey.SupportsOSVersion() {
//				continue
//			}
//			pd.AddDefault(NewPrefDefault(dn, PrefName(subKey.Name), &PrefDefaultOpts{
//				Kind:          subKey.Kind(),
//				SupportedIn:   OSVersion(subKey.MacOSMin),
//				UnsupportedIn: "", // TODO Add support for this
//			}))
//		}
//		return &pd
//	}

func Iterator(pms *ProfileManifests) (ei yamlutil.EntriesIterator) {
	var errs errutil.MultiErr
	ei = yamlutil.EntriesIterator{
		Seq2: func(yield func(entry yamlutil.FilterableEntry, _ error) (ok bool)) {
			for _, file := range pms.Files() {
				content, err := file.Content()
				if err != nil {
					errs.Add(err)
					continue
				}
				if !yield(ReadProfileManifest(bytes.NewReader(content))) {
					break
				}
			}
		},
	}
	if errs.IsError() {
		ei.Err = errs.Err()
	}
	return ei
}
