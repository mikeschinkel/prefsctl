package filesutil

import (
	"fmt"
	"io"
	"io/fs"
	"log/slog"
)

// ListFilesRecursive returns a slice of file entries from an fs.FS
func ListFilesRecursive(fileSys fs.FS, dir string) (files []FileEntry, err error) {
	err = fs.WalkDir(fileSys, dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			goto end
		}
		if entry.IsDir() {
			goto end
		}
		files = append(files, NewFileEntryWithFS(entry, path, fileSys))
	end:
		return err
	})
	return files, err
}

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}
func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}
