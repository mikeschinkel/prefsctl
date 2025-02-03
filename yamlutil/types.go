package yamlutil

import (
	"iter"
)

type FilterableEntry interface {
	FilterableEntry()
}
type EntriesIterator struct {
	iter.Seq2[FilterableEntry, error]
	Err error
}

type EntryFilterFunc func(FilterableEntry) bool

type GetDocFromEntryFunc func(FilterableEntry, EntryFilterFunc) (Document, error)

type MultidocFileizer interface {
	YAMLFile() MultidocFile
}

type Documentizer interface {
	YAMLDocument() Document
}

type MultidocFileBuilder interface {
	BuildYAMLFile() (MultidocFile, error)
}

type DocumentBuilder interface {
	BuildYAMLDocument() (Document, error)
}
