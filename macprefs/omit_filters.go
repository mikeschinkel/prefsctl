package macprefs

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
)

//=============================================

type omitWhenValueEqualsDefaultFilter []kvfilters.MatchFunc2
type oved = omitWhenValueEqualsDefaultFilter

var _ kvfilters.Filter = (*oved)(nil)

func (f oved) Funcs2() []kvfilters.MatchFunc2 {
	return f
}
func (oved) Effect() kvfilters.Effect {
	return kvfilters.Omit
}
func (oved) Target() kvfilters.Target {
	return kvfilters.KeyValues
}
func (oved) Comparison() kvfilters.Comparison {
	return kvfilters.IsTrue
}
func (oved) Matches() kvfilters.Matches {
	return kvfilters.Func
}
func (oved) IgnoreCase() bool {
	return false
}

func OmitValueEqualsDefaultFilter() kvfilters.Filter {
	return omitWhenValueEqualsDefaultFilter{
		func(kv kvfilters.KeyValue) bool {
			getter, ok := kv.(*Pref)
			if !ok {
				panicf("KeyValue does not have a Default() method: %s", kv.ErrorInfo())
			}
			return kv.Value() == getter.Default()
		},
	}
}
