package filesutil

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/mikeschinkel/prefsctl/errutil"
)

type FileEntry interface {
	Name() string
	Filepath() string
	Content() ([]byte, error)
	Reader() (io.Reader, error)
	DirEntry() fs.DirEntry
}

var _ FileEntry = (*file)(nil)

type file struct {
	dirEntry fs.DirEntry
	filepath string
	fs       fs.FS
}

func (f *file) Name() string {
	return f.dirEntry.Name()
}

func (f *file) DirEntry() fs.DirEntry {
	return f.dirEntry
}

func NewFileEntryWithFS(dirEntry fs.DirEntry, fp string, fileSys fs.FS) FileEntry {
	fileEntry := NewFileEntry(dirEntry, fp)
	fileEntry.(*file).fs = fileSys
	return fileEntry
}

func NewFileEntry(dirEntry fs.DirEntry, fp string) FileEntry {
	if dirEntry == nil {
		panic(fmt.Sprintf("ERROR: fs.DirEntry cannot be nil for file %s", fp))
	}
	if fp == "" {
		panic("ERROR: filepath for FileEntry cannot be empty")
	}
	if dirEntry.IsDir() {
		panic(fmt.Sprintf("ERROR: fs.DirEntry represents a directory not a file: %s", fp))
	}
	return &file{
		dirEntry: dirEntry,
		filepath: fp,
	}
}

func (f *file) ErrArgs() error {
	errs := []error{
		errutil.ErrArg(errutil.FilepathErrKey, f.filepath),
	}
	fi, err := f.dirEntry.Info()
	if err != nil {
		errs = append(errs, err, errutil.ErrGettingFileInfo)
		goto end
	}
	errs = append(errs,
		errutil.ErrArg(errutil.FileSizeErrKey, fi.Size()),
		errutil.ErrArg(errutil.FileModeErrKey, fi.Mode()),
		errutil.ErrArg(errutil.FileModTimeErrKey, fi.ModTime()),
	)
end:
	return errors.Join(errs...)
}

func (f *file) Content() (b []byte, err error) {
	var fsFile fs.File
	var fi fs.FileInfo
	var n int
	if f.fs == nil {
		b, err = os.ReadFile(f.filepath)
		if err != nil {
			err = errors.Join(err, errutil.ErrReadingFile, f.ErrArgs())
		}
		goto end
	}
	fsFile, err = f.fs.Open(f.filepath)
	if err != nil {
		err = errors.Join(err, errutil.ErrOpeningFile, f.ErrArgs())
		goto end
	}
	fi, err = fsFile.Stat()
	if err != nil {
		err = errors.Join(err, errutil.ErrGettingFileInfo, f.ErrArgs())
		goto end
	}
	b = make([]byte, fi.Size())
	defer mustClose(fsFile)
	n, err = fsFile.Read(b)
	if int64(n) != fi.Size() {
		err = errors.Join(
			errutil.ErrBytesReadNotSameAsFileSize,
			errutil.ErrArg(errutil.BytesReadErrKey, n),
			f.ErrArgs(),
		)
	}
end:
	return b, err
}

func (f *file) Reader() (r io.Reader, err error) {
	if f.fs != nil {
		return f.fs.Open(f.filepath)
	}
	return os.Open(f.filepath)
}

func (f *file) Filepath() string {
	return f.filepath
}
