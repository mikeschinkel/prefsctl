package macpreflabels

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
)

type (
	Label      = kvfilters.Label
	Labels     = kvfilters.Labels
	LabelName  = kvfilters.LabelName
	LabelValue = kvfilters.LabelValue
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
	UserManaged      = NewLabel(Class, "userManaged")
	SystemManaged    = NewLabel(Class, "systemManaged")
	AppManaged       = NewLabel(Class, "appManaged")
	RuntimeState     = NewLabel(Class, "runtimeState")
	VersionMigration = NewLabel(Class, "versionMigrate")
)

//goland:noinspection GoStructInitializationWithoutFieldNames
var (
	UnknownSets = kvfilters.NewUnknownLabel(Sets, "unknownSets")
	Optional    = NewLabel(Sets, "optional")
	Required    = NewLabel(Sets, "required")
)

var LabelMap = map[kvfilters.LabelValue]*kvfilters.Label{
	UserManaged.Value:      &UserManaged,
	SystemManaged.Value:    &SystemManaged,
	AppManaged.Value:       &AppManaged,
	RuntimeState.Value:     &RuntimeState,
	VersionMigration.Value: &VersionMigration,
	Optional.Value:         &Optional,
	Required.Value:         &Required,
}

func GetLabelByValue(value LabelValue) *Label {
	label, _ := LabelMap[value]
	return label
}
func GetLabelsByValues(values []LabelValue) *Labels {
	labels := make([]*Label, len(values))
	for i, value := range values {
		labels[i] = GetLabelByValue(value)
	}
	return kvfilters.NewLabels(labels...)
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
	Optional.Value:    "Optional",
	Required.Value:    "Required",
}
