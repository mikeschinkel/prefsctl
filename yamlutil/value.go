package yamlutil

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logargs"
	"gopkg.in/yaml.v3"
)

type Value struct {
	Value any
}

func NewValue(value any) *Value {
	return &Value{Value: value}
}

//func (v *Value) Empty() (empty bool) {
//	switch t := v.Value.(type) {
//	case []string:
//		empty = len(t) == 0
//	case string:
//		empty = t == ""
//	case bool:
//		empty = !t
//	case nil:
//		empty = true
//	case any:
//		empty = true
//	default:
//		panic(fmt.Sprintf("Type '%T' for '%v' not yet implemented", v.Value, v.Value))
//	}
//	return empty
//}

// LogArgs returns array of log args for the Value
func (v *Value) LogArgs() []any {
	return []any{
		logargs.ContainedType, fmt.Sprintf("%T", v.Value),
		logargs.ContainedValue, v.Value,
	}
}

// Kind returns reflect.Kind corresponding to the value of .Value field
func (v *Value) Kind() (rk reflect.Kind, err error) {
	chkValid(v, notNil)
	rk = reflect.Invalid
	switch v.Value.(type) {
	case string:
		rk = reflect.String
	case bool:
		rk = reflect.Bool
	case int:
		rk = reflect.Int64
	default:
		panic(fmt.Sprintf("Type '%T' for '%v' not yet implemented", v.Value, v.Value))
	}
	return rk, err
}

func (v *Value) Float64() (f float64) {
	chkValid(v, notNil, notSet)
	switch t := v.Value.(type) {
	case float64:
		f = t
	case float32:
		f = float64(t)
	default:
		panicf("Invalid type '%T' for '%v'; expected 'float64'", v.Value, v.Value)
	}
	return f
}

func (v *Value) Type() (s string) {
	return fmt.Sprint("%T", v.Value)
}

func (v *Value) String() (s string) {
	chkValid(v, notNil, notSet)
	switch t := v.Value.(type) {
	case []string:
		s = strings.Join(t, ",")
	case string:
		s = t
	case int:
		s = strconv.Itoa(t)
	case int64:
		s = strconv.FormatInt(t, 10)
	case bool:
		switch t {
		case true:
			s = "true"
		case false:
			s = "false"
		}
	case any:
		s = ""
	//case nil:
	//	s = ""
	case *Value:
		s = fmt.Sprintf("%v", t)
	default:
		s = fmt.Sprintf("%v", v.Value)
	}
	return s
}

// MarshalYAML allows Value to contain one string or a slice of string
func (v *Value) MarshalYAML() (out any, err error) {
	switch t := v.Value.(type) {
	case []string:
		out = t
	case string:
		out = t
	case nil:
		out = ""
	case *Value:
		if t == nil || t.Value == "<nil>" {
			out = ""
			goto end
		}
		out = fmt.Sprintf("%s", t)
	default:
		panic(fmt.Sprintf("Type '%T' for '%v' not yet implemented", v.Value, v.Value))
	}
end:
	return out, err
}

// UnmarshalYAML unmarshalls a Value
func (v *Value) UnmarshalYAML(value *yaml.Node) (err error) {
	var tv *Value
	tv, err = UnmarshalYAMLNode(value)
	if err != nil {
		goto end
	}
	*v = *tv
end:
	return err
}

// UnmarshalYAMLNode unmarshalls a yaml.Node based on its Tag field
func UnmarshalYAMLNode(value *yaml.Node) (v *Value, err error) {
	var val any
	var errs errutil.MultiErr

	switch value.Tag {
	case "!!str":
		val = value.Value
	case "!!bool":
		switch value.Value {
		case "true":
			val = true
		case "false":
			val = false
		}
	case "!!int":
		val, err = strconv.ParseInt(value.Value, 10, 64)
		if err != nil {
			goto end
		}
		switch {
		case math.MaxInt == math.MaxInt64:
			// If int and int64 are the same, just make it int
			val = int(val.(int64))
		case val.(int64) <= math.MaxInt:
			// If int and int64 are different but value fits in int, make it int
			val = int(val.(int64))
		}
	case "!!seq":
		slice := make([]any, len(value.Content))
		for i, node := range value.Content {
			slice[i], err = UnmarshalYAMLNode(node)
			if err != nil {
				errs.Add()
			}
		}
		err = errs.Err()
		val = slice
	default:
		panic(fmt.Sprintf("Tag '%s' not yet implemented", value.Tag))
	}
	v = NewValue(val)
end:
	return v, err
}
