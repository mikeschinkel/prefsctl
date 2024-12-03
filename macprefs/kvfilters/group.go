package kvfilters

import (
	"fmt"

	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
)

// Name for the group or key-value pair, or anything else that needs a name.
type Name string

// Code is the camel-cased version of name without spaces used for comparison,
// e.g. "userManaged" instead of "User Managed."
type Code string

type Group interface {
	Name() Name
	Code() Code
	KeyValues() []KeyValue
	AddKeyValue(KeyValue)
	LogArgs() []any
	ErrorInfo() error
	ShallowCopy() Group
	fmt.Stringer
}

var _ Group = (*group)(nil)

type group struct {
	name        Name
	code        Code
	keyValues   []KeyValue
	initialized bool
}

func (g *group) ShallowCopy() Group {
	return &group{
		name:        g.name,
		code:        g.code,
		keyValues:   make([]KeyValue, 0),
		initialized: false,
	}
}

func (g *group) String() string {
	if string(g.name) == string(g.code) {
		return string(g.name)
	}
	return fmt.Sprintf("%s [%s]", g.name, g.code)
}

func (g *group) Initialize() error {
	g.initialized = true
	return nil
}

// NewGroupFromGroup creates a new group from an existing one w/o copying its keyValues
func NewGroupFromGroup(g Group) Group {
	return &group{
		name:      g.Name(),
		code:      g.Code(),
		keyValues: make([]KeyValue, 0),
	}
}

func (g *group) AddKeyValue(kv KeyValue) {
	g.keyValues = append(g.keyValues, kv)
}

func (g *group) ErrorInfo() error {
	return fmt.Errorf("%s=%s", logargs.GroupNameLogArg, g.name)
}

func (g *group) LogArgs() []any {
	return []any{
		logargs.GroupNameLogArg,
		g.name,
	}
}

func (g *group) Name() Name {
	return g.name
}

func (g *group) Code() Code {
	return g.code
}
