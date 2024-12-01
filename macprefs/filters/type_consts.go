package filters

type Effect byte

const (
	InvalidEffect Effect = 0
	Omit                 = 'o'
	Keep                 = 'k'
)

type Comparison byte

const (
	InvalidComparison Comparison = 0
	Equals                       = '='
	NotEquals                    = '!'
	IsTrue                       = Equals
	IsFalse                      = NotEquals
	IsMatched                    = Equals
	NotMatched                   = NotEquals
)

type Target byte

func (t Target) String() string {
	switch t {
	case Groups:
		return "Groups"
	case Keys:
		return "Keys"
	case Values:
		return "Values"
	case KeyValues:
		return "KeyValues"
	case InvalidTarget:
		//fallthrough
	}
	return "Invalid target"
}

const (
	InvalidTarget Target = 0
	Groups               = 'g'
	Keys                 = 'k'
	Values               = 'v'
	KeyValues            = 'b' // b=both
)

var KeyValueOrKeyValue = []Target{Keys, Values, KeyValues}

type Matches byte

func (m Matches) String() string {
	switch m {
	case Prefix:
		return "Prefix"
	case Suffix:
		return "Suffix"
	case Entirety:
		return "Entirety"
	case Func:
		return "Func"
	case Regexp:
		return "Regexp"
	case Contains:
		return "Contains"
	case InvalidMatches:
		//fallthrough
	}
	return "Invalid match criteria"
}

const (
	InvalidMatches Matches = 0
	Prefix                 = 'p'
	Suffix                 = 's'
	Entirety               = 'e'
	Func                   = 'f'
	Regexp                 = 'r'
	Contains               = 'c'
)

type MatchFunc1 func(Code) bool
type MatchFunc2 func(Code, string) bool
