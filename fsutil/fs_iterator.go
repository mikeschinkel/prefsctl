package fsutil

import (
	"io/fs"
	"iter"

	"github.com/mikeschinkel/prefsctl/errutil"
)

// FSIterator returns a slice of filesutil.FileEntry
type FSIterator struct {
	FS        fs.FS
	Filenames []string
	Errs      errutil.MultiErr
}

func NewFSIterator(FS fs.FS, filenames []string) FSIterator {
	return FSIterator{
		FS:        FS,
		Filenames: filenames,
	}
}

func (i FSIterator) All() iter.Seq2[fs.File, error] {
	return func(yield func(file fs.File, _ error) bool) {
		i.Errs.Clear()
		for _, filename := range i.Filenames {
			f, err := i.FS.Open(filename)
			if err != nil {
				i.Errs.Add(err)
				continue
			}
			if !yield(f, err) {
				break
			}
			mustClose(f)
		}
	}
}
