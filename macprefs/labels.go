package macprefs

import (
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type LabelName Name
type LabelValue string

const (
	InvalidLabel LabelName = "invalid-label"
	WhatSetsName LabelName = "what-sets"
	PrefType     LabelName = "pref-type"
	OSVersion    LabelName = "os-version"
)

type Labels []*Label

func (ll Labels) String() string {
	s, _ := sliceconv.ToStringsFunc(ll, func(label *Label) (bool, string, error) {
		return true, label.String(), nil
	})
	return strings.Join(s, ",")
}

func (l Label) String() string {
	return fmt.Sprintf("%s=%s", l.Name, l.Value)
}

type Label struct {
	Name  LabelName
	Value LabelValue
}

func (l Label) GoProperty() string {
	return fmt.Sprintf("macprefs.%s: %s", l.Name, l.Value)
}

var (
	UnknownSets = Label{WhatSetsName, "unknown-sets"}
	MacOSSets   = Label{WhatSetsName, "macos-sets"}
	InstallSets = Label{WhatSetsName, "install-sets"}
)

var (
	StringType   = Label{PrefType, "string"}
	NumberType   = Label{PrefType, "number"}
	BoolType     = Label{PrefType, "bool"}
	LanguageType = Label{PrefType, "language"}
	LocaleType   = Label{PrefType, "locale"}
)

func GoVarName(value LabelValue) LabelName {
	name, ok := goVarMap[value]
	if !ok {
		name = InvalidLabel
	}
	return name
}

var goVarMap = map[LabelValue]LabelName{
	UnknownSets.Value: "UnknownSets",
	MacOSSets.Value:   "MacOSSets",
	InstallSets.Value: "InstallSets",

	StringType.Value:   "StringType",
	NumberType.Value:   "NumberType",
	BoolType.Value:     "BoolType",
	LanguageType.Value: "LanguageType",
	LocaleType.Value:   "LocaleType",
}
