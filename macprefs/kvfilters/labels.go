package kvfilters

import (
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type LabelName macosutils.Name
type LabelValue string

const (
	InvalidLabel LabelName = "invalid"
	Sets         LabelName = "sets"
	Type         LabelName = "type"
	MacOS        LabelName = "macos"
	Class        LabelName = "class"
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
	UserManaged      = Label{Class, "userManaged"}
	SystemManaged    = Label{Class, "systemManaged"}
	AppManaged       = Label{Class, "appManaged"}
	RuntimeState     = Label{Class, "runtimeState"}
	VersionMigration = Label{Class, "versionMigrate"}
)

var (
	UnknownSets = Label{Sets, "unknown"}
	DefaultsSet = Label{Sets, "defaults"}
	SetupSets   = Label{Sets, "setup"}
)

var (
	UnknownType  = Label{Type, "unknown"}
	StringType   = Label{Type, "string"}
	NumberType   = Label{Type, "number"}
	BoolType     = Label{Type, "bool"}
	LanguageType = Label{Type, "language"}
	LocaleType   = Label{Type, "locale"}
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
	DefaultsSet.Value: "DefaultsSet",
	SetupSets.Value:   "SetupSets",

	StringType.Value:   "StringType",
	NumberType.Value:   "NumberType",
	BoolType.Value:     "BoolType",
	LanguageType.Value: "LanguageType",
	LocaleType.Value:   "LocaleType",
}
