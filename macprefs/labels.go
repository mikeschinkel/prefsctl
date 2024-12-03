package macprefs

import (
	"github.com/mikeschinkel/prefsctl/macprefs/kvfilters"
)

const (
	InvalidLabel kvfilters.LabelName = "invalid"
	Sets         kvfilters.LabelName = "sets"
	Type         kvfilters.LabelName = "type"
	MacOS        kvfilters.LabelName = "macos"
	Class        kvfilters.LabelName = "class"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UserManaged      = kvfilters.Label{Class, "userManaged"}
	SystemManaged    = kvfilters.Label{Class, "systemManaged"}
	AppManaged       = kvfilters.Label{Class, "appManaged"}
	RuntimeState     = kvfilters.Label{Class, "runtimeState"}
	VersionMigration = kvfilters.Label{Class, "versionMigrate"}
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UnknownSets = kvfilters.Label{Sets, "unknown"}
	DefaultsSet = kvfilters.Label{Sets, "defaults"}
	SetupSets   = kvfilters.Label{Sets, "setup"}
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UnknownType  = kvfilters.Label{Type, "unknown"}
	StringType   = kvfilters.Label{Type, "string"}
	NumberType   = kvfilters.Label{Type, "number"}
	BoolType     = kvfilters.Label{Type, "bool"}
	LanguageType = kvfilters.Label{Type, "language"}
	LocaleType   = kvfilters.Label{Type, "locale"}
)

func GoVarName(value kvfilters.LabelValue) kvfilters.LabelName {
	name, ok := goVarMap[value]
	if !ok {
		name = InvalidLabel
	}
	return name
}

var goVarMap = map[kvfilters.LabelValue]kvfilters.LabelName{
	UnknownSets.Value: "UnknownSets",
	DefaultsSet.Value: "DefaultsSet",
	SetupSets.Value:   "SetupSets",

	StringType.Value:   "StringType",
	NumberType.Value:   "NumberType",
	BoolType.Value:     "BoolType",
	LanguageType.Value: "LanguageType",
	LocaleType.Value:   "LocaleType",
}
