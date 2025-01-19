package prefsyaml

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mikeschinkel/prefsctl/errutil"
	"gopkg.in/yaml.v3"
)

type YAMLValue struct {
	Value any
}

func (v *YAMLValue) Empty() (empty bool) {
	switch t := v.Value.(type) {
	case []string:
		empty = len(t) == 0
	case string:
		empty = t == ""
	case bool:
		empty = !t
	case nil:
		empty = true
	case any:
		empty = true
	default:
		panic(fmt.Sprintf("Type '%T' for '%v' not yet implemented", v.Value, v.Value))
	}
	return empty
}

func (v *YAMLValue) String() (s string) {
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
	case nil:
		s = ""
	default:
		panic(fmt.Sprintf("Type '%T' for '%v' not yet implemented", v.Value, v.Value))
	}
	return s
}

// MarshalYAML allows Value to contain one string or a slice of string
func (v *YAMLValue) MarshalYAML() (out any, err error) {
	switch t := v.Value.(type) {
	case []string:
		out = t
	case string:
		out = t
	case nil:
		out = nil
	default:
		typ := fmt.Sprintf("%s", v.Value)
		print(typ)
		panic(fmt.Sprintf("Type '%T' for '%v' not yet implemented", v.Value, v.Value))
	}
	return out, err
}

// UnmarshalYAML unmarshalls a YAMLValue
func (v *YAMLValue) UnmarshalYAML(value *yaml.Node) (err error) {
	var tv YAMLValue
	tv, err = UnmarshalYAMLNode(value)
	if err != nil {
		goto end
	}
	*v = tv
end:
	return err
}

// UnmarshalYAMLNode unmarshalls a yaml.Node based on its Tag field
func UnmarshalYAMLNode(value *yaml.Node) (v YAMLValue, err error) {
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
	v = YAMLValue{Value: val}
end:
	return v, err
}

//// UnmarshalYAML allows Value to contain one string or a slice of string
//func (v *YAMLValue) UnmarshalYAML(value *yaml.Node) (err error) {
//	// Create a type alias to avoid recursive UnmarshalYAML calls
//	type YAMLValueAlias YAMLValue
//
//	// First unmarshal all the alias fields using the type alias
//	var alias YAMLValueAlias
//	alias.Value = make([]string, 0)
//	err = value.Decode(&alias)
//	if err != nil {
//		err = errors.Join(ErrFailedToDecodeYAML, err)
//		goto end
//	}
//	*yp = YAMLValue(alias)
//
//	// Find and specially handle the 'names' field
//	for i, node := range value.Content {
//		if i%2 != 0 {
//			continue
//		}
//		if node.Value != "value" {
//			continue
//		}
//		// Get the value node
//		valueNode := value.Content[i+1]
//
//		// Try as single value
//		var single string
//		err = valueNode.Decode(&single)
//		if err == nil {
//			yp.Value = []string{single}
//			goto end
//		}
//
//		// Try as slice
//		var slice []string
//		err = valueNode.Decode(&slice)
//		if err != nil {
//			goto end
//		}
//		yp.Value = slice
//		goto end
//	}
//end:
//	return err
//}
