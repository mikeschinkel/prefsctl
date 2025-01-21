package profilemanifests

import (
	"embed"
	"io"
	"io/fs"
	"path/filepath"
	"slices"

	"github.com/mikeschinkel/prefsctl/logargs"
)

type (
	EntryFiles  []*EmbeddedFile
	DirEntryMap map[string]EntryFiles
)

type EmbeddedFile struct {
	fs.DirEntry
	Filepath string
}

func NewEmbeddedFile(file fs.DirEntry, fp string) *EmbeddedFile {
	return &EmbeddedFile{
		DirEntry: file,
		Filepath: fp,
	}
}

func (f *EmbeddedFile) Content() ([]byte, error) {
	return profileManifests.ReadFile(f.Filepath)
}

func (f *EmbeddedFile) Reader() (r io.Reader, err error) {
	return profileManifests.Open(f.Filepath)
}

//go:embed data/ProfileManifests
var profileManifests embed.FS

const (
	profileManifestsDir = "data/ProfileManifests/Manifests"
)

var (
	profileManifestFiles   EntryFiles
	profileManifestFileMap DirEntryMap
)

type ProfileManifests struct {
	files   EntryFiles
	fileMap DirEntryMap
}

func New() *ProfileManifests {
	return &ProfileManifests{}
}

func (pm *ProfileManifests) sortFiles() {
	slices.SortFunc(pm.files, func(a, b *EmbeddedFile) int {
		switch {
		case a.Filepath < b.Filepath:
			return -1
		case a.Filepath > b.Filepath:
			return 1
		}
		return 0
	})
}

func (pm *ProfileManifests) Files() EntryFiles {
	pm.sortFiles()
	return pm.files
}

func (pm *ProfileManifests) Load() error {
	files, err := ListFilesRecursive(profileManifests, profileManifestsDir)
	if err != nil {
		goto end
	}
	pm.files = make(EntryFiles, 0, len(files))
	pm.fileMap = make(DirEntryMap, len(pm.files))
	for _, file := range files {
		name := file.Name()
		_, ok := pm.fileMap[name]
		if ok {
			slog.Warn("Unexpected duplicated preference domain", logargs.PrefsDomain, filepath.Base(file.Name()))
		} else {
			pm.fileMap[name] = make(EntryFiles, 0, 1)
		}
		pm.files = append(pm.files, file)
		pm.fileMap[name] = append(pm.fileMap[name], file)
	}
end:
	return err
}

// ListFilesRecursive returns a slice of file entries from an fs.FS
func ListFilesRecursive(fileSys fs.FS, dir string) (files EntryFiles, err error) {
	err = fs.WalkDir(fileSys, dir, func(path string, entry fs.DirEntry, err error) error {
		var file *EmbeddedFile
		if err != nil {
			goto end
		}
		if entry.IsDir() {
			goto end
		}
		file = NewEmbeddedFile(entry, path)
		files = append(files, file)
	end:
		return err
	})
	return files, err
}
