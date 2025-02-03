package yamlutil

import (
	"reflect"
)

type chkType int8
type chkTypes []chkType

func (cts chkTypes) has(ct chkType) (has bool) {
	for _, t := range cts {
		if t == ct {
			return true
		}
	}
	return false
}

const (
	invalidChkType chkType = iota
	notNil
	notSet
	notZero
)

// chkValid checks a value to see if it contains a non-nil value, and optionally
// ensures it itself is not nil.
func chkValid(v *Value, checks ...chkType) {
	cts := chkTypes(checks)
	if cts.has(notNil) && v == nil {
		panic("The property of type YAML Value cannot be nil")
	}
	if cts.has(notSet) && v.Value == nil {
		panic("The contained value within the YAML Value not set, e.g. it is nil")
	}
	if cts.has(notZero) && reflect.ValueOf(v.Value).IsZero() {
		panicf("The contained value of type '%T' within the YAML Value has a zero value", v.Value)
	}
}
