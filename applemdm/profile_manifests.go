package applemdm

import (
	"embed"
	"io/fs"
	"path/filepath"

	"github.com/mikeschinkel/prefsctl/logargs"
)

type (
	DirEntryMap map[string]fs.DirEntry
)

//go:embed data/ProfileManifests/Manifests
var profileManifests embed.FS

const (
	profileManifestsDir = "data/ProfileManifests/Manifests"
)

var (
	profileManifestFiles   []fs.DirEntry
	profileManifestFileMap DirEntryMap
)

type ProfileManifests struct {
	files   []fs.DirEntry
	fileMap DirEntryMap
}

func (pm *ProfileManifests) Files() []fs.DirEntry {
	return pm.files
}

func (pm *ProfileManifests) Load() (err error) {
	pm.files, err = ListFilesRecursive(profileManifests, profileManifestsDir)
	if err != nil {
		goto end
	}
	pm.fileMap = make(DirEntryMap, len(pm.files))
	for _, file := range pm.files {
		name := file.Name()
		_, ok := pm.fileMap[name]
		if ok {
			slog.Error("Unexpected duplicated domain", logargs.PrefsDomain, filepath.Base(file.Name()))
		}
		pm.fileMap[name] = file

	}
end:
	return err
}

// ListFilesRecursive returns a slice of file entries from an fs.FS
func ListFilesRecursive(fileSys fs.FS, dir string) (files []fs.DirEntry, err error) {
	err = fs.WalkDir(fileSys, dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			goto end
		}
		if entry.IsDir() {
			goto end
		}
		files = append(files, entry)
	end:
		return err
	})
	return files, err
}
