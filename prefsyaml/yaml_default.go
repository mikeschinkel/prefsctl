package prefsyaml

import (
	"strings"

	"github.com/mikeschinkel/prefsctl/kvfilters"
)

type YAMLDefault struct {
	Domain *YAMLDomain
	Name   PrefName
	Value  string // raw string value for default
	Labels *kvfilters.Labels
	Type   PreferenceType
}

func (d YAMLDefault) TypeName() string {
	typ, _ := strings.CutSuffix(string(d.Type), "Type")
	return typ
}
