package macprefs

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
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
	DescrVerified   = NewLabel(Verification, "descrVerified")
	SetsVerified    = NewLabel(Verification, "setsVerified")
	ClassVerified   = NewLabel(Verification, "classVerified")
	OptionsVerified = NewLabel(Verification, "optionsVerified")

	DescrNotVerified = NewLabel(Verification, "descrNotVerified")
	TypeNotVerified  = NewLabel(Verification, "typeNotVerified")
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
	UnknownType = kvfilters.NewUnknownLabel(Type, kvfilters.LabelValue(macosutil.UnknownType))
	StringType  = kvfilters.NewUnknownLabel(Type, kvfilters.LabelValue(macosutil.StringType))
	NumberType  = kvfilters.NewUnknownLabel(Type, kvfilters.LabelValue(macosutil.NumberType))
	IntType     = kvfilters.NewUnknownLabel(Type, kvfilters.LabelValue(macosutil.IntType))
	FloatType   = kvfilters.NewUnknownLabel(Type, kvfilters.LabelValue(macosutil.FloatType))
	BoolType    = kvfilters.NewUnknownLabel(Type, kvfilters.LabelValue(macosutil.BoolType))

	IntBoolType = kvfilters.NewUnknownLabel(Type, kvfilters.LabelValue(macosutil.IntBoolType))

	LanguageType = kvfilters.NewUnknownLabel(Type, kvfilters.LabelValue(macosutil.LanguageType))
	LocaleType   = kvfilters.NewUnknownLabel(Type, kvfilters.LabelValue(macosutil.LocaleType))
)
var TypeLabelMap = map[macosutil.PreferenceType]*kvfilters.Label{
	macosutil.UnknownType:  &UnknownType,
	macosutil.StringType:   &StringType,
	macosutil.NumberType:   &NumberType,
	macosutil.IntType:      &IntType,
	macosutil.FloatType:    &FloatType,
	macosutil.BoolType:     &BoolType,
	macosutil.IntBoolType:  &IntBoolType,
	macosutil.LanguageType: &LanguageType,
	macosutil.LocaleType:   &LocaleType,
}

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
