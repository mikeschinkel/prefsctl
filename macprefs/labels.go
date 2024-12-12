package macprefs

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
)

type Verified = preftemplates.Verified

const (
	InvalidLabel kvfilters.LabelName = "invalid"
	Sets         kvfilters.LabelName = "sets"
	Type         kvfilters.LabelName = "type"
	MacOS        kvfilters.LabelName = "macOS"
	Class        kvfilters.LabelName = "class"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UserManaged      = kvfilters.NewLabel(Class, "userManaged")
	SystemManaged    = kvfilters.NewLabel(Class, "systemManaged")
	AppManaged       = kvfilters.NewLabel(Class, "appManaged")
	RuntimeState     = kvfilters.NewLabel(Class, "runtimeState")
	VersionMigration = kvfilters.NewLabel(Class, "versionMigrate")
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UnknownSets = kvfilters.NewUnknownLabel(Sets, "unknownSets")
	DefaultsSet = kvfilters.NewLabel(Sets, "defaultsSet")
	SetupSets   = kvfilters.NewLabel(Sets, "setupSets")
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UnknownType = kvfilters.NewUnknownLabel(Type, "unknownType")
	StringType  = kvfilters.NewLabel(Type, "stringType")
	NumberType  = kvfilters.NewLabel(Type, "numberType")
	IntType     = kvfilters.NewLabel(Type, "intType")
	FloatType   = kvfilters.NewLabel(Type, "floatType")
	BoolType    = kvfilters.NewLabel(Type, "boolType")

	IntBoolType = kvfilters.NewLabel(Type, "intBoolType")

	LanguageType = kvfilters.NewLabel(Type, "languageType")
	LocaleType   = kvfilters.NewLabel(Type, "localeType")
	YesNoType    = kvfilters.NewLabel(Type, "yesNoType")
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
