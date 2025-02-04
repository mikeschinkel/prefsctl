package yamlutil

import (
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/errutil"
)

type MultidocFile interface {
	fmt.Stringer
	Documents() []Document
	AddDocument(Document)
}

var _ MultidocFile = (*multidocFile)(nil)

type multidocFile struct {
	docs []Document
}

func NewMultidocFile() MultidocFile {
	return &multidocFile{
		docs: make([]Document, 0),
	}
}

func (f *multidocFile) AddDocument(yd Document) {
	f.docs = append(f.docs, yd)
}

func NewFile() MultidocFile {
	return &multidocFile{
		docs: make([]Document, 0),
	}
}

func (f *multidocFile) Documents() []Document {
	return f.docs
}

func (f *multidocFile) String() string {
	var sb strings.Builder
	for _, y := range f.docs {
		if y == "" {
			continue
		}
		sb.WriteString(string(y))
	}
	return sb.String()
}

func BuildMultidocFile(entryIterator EntriesIterator, filter EntryFilterFunc, getDocFromEntry GetDocFromEntryFunc) (yf MultidocFile, err error) {
	var errs errutil.MultiErr
	var entry FilterableEntry
	var doc Document

	yf = NewFile()
	for entry, err = range entryIterator.Seq2 {
		if err != nil {
			errs.Add(err)
			continue
		}
		doc, err = getDocFromEntry(entry, filter)
		if err != nil {
			errs.Add(err)
			continue
		}
		yf.AddDocument(doc)
	}
	return yf, errs.Err()
}
