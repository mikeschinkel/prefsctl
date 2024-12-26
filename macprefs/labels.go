package macprefs

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
)

type (
	LabelName = kvfilters.LabelName
)

var (
	NewLabel = kvfilters.NewLabel
)

const (
	InvalidLabel LabelName = "invalid"
	Sets         LabelName = "sets"
	Type         LabelName = "type"
	MacOS        LabelName = "macOS"
	Class        LabelName = "class"
	Verification LabelName = "verification"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	TypeVerified    = NewLabel(Verification, "typeVerified")
	DefaultVerified = NewLabel(Verification, "defaultVerified")
	ValueVerified   = NewLabel(Verification, "valueVerified")
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UserManaged      = NewLabel(Class, "userManaged")
	SystemManaged    = NewLabel(Class, "systemManaged")
	AppManaged       = NewLabel(Class, "appManaged")
	RuntimeState     = NewLabel(Class, "runtimeState")
	VersionMigration = NewLabel(Class, "versionMigrate")
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UnknownSets = kvfilters.NewUnknownLabel(Sets, "unknownSets")
	DefaultsSet = NewLabel(Sets, "defaultsSet")
	SetupSets   = NewLabel(Sets, "setupSets")
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UnknownType = kvfilters.NewUnknownLabel(Type, "unknownType")
	StringType  = NewLabel(Type, "stringType")
	NumberType  = NewLabel(Type, "numberType")
	IntType     = NewLabel(Type, "intType")
	FloatType   = NewLabel(Type, "floatType")
	BoolType    = NewLabel(Type, "boolType")

	IntBoolType = NewLabel(Type, "intBoolType")

	LanguageType = NewLabel(Type, "languageType")
	LocaleType   = NewLabel(Type, "localeType")
	YesNoType    = NewLabel(Type, "yesNoType")
)

func GoVarName(value kvfilters.LabelValue) LabelName {
	name, ok := goVarMap[value]
	if !ok {
		name = InvalidLabel
	}
	return name
}

var goVarMap = map[kvfilters.LabelValue]LabelName{
	UnknownSets.Value: "UnknownSets",
	DefaultsSet.Value: "DefaultsSet",
	SetupSets.Value:   "SetupSets",

	StringType.Value:   "StringType",
	NumberType.Value:   "NumberType",
	BoolType.Value:     "BoolType",
	LanguageType.Value: "LanguageType",
	LocaleType.Value:   "LocaleType",
}
