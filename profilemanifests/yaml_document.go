package profilemanifests

import (
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

func GetYAMLDocumentFromProfileManifest(entry yamlutil.FilterableEntry, filter yamlutil.EntryFilterFunc) (yd yamlutil.Document, err error) {
	var yr prefsyaml.Resource
	var isValid bool

	pm, ok := entry.(*ProfileManifest)
	if !ok {
		panicf("ERROR: Expected entry to be a *ProfileManifest; got '%T' instead.", entry)
	}
	yr, isValid = pm.GetPrefsYAMLResource(filter)
	if !isValid {
		goto end
	}
	yd, err = yr.YAMLDocument()
	if err != nil {
		goto end
	}
end:
	return yd, err
}
