package fsutil

import (
	"embed"
	"path/filepath"
)

func GetEmbeddedFilenames(fileSys embed.FS, rootPath string) (filenames []string, err error) {
	entries, err := fileSys.ReadDir(rootPath)
	if err != nil {
		goto end
	}
	for _, entry := range entries {
		filenames = append(filenames, filepath.Join(rootPath, entry.Name()))
	}

end:
	return filenames, err
}
