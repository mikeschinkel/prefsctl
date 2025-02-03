package profilemanifests

import (
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

type Value struct {
	*yamlutil.Value `yaml:",inline"`
}

func NewValue(value any) *Value {
	return &Value{
		Value: yamlutil.NewValue(value),
	}
}

func ExtractYAMLUtilValues(values []*Value) (extracted []*yamlutil.Value) {
	extracted = make([]*yamlutil.Value, len(values))
	for i, value := range values {
		extracted[i] = value.Value
	}
	return extracted
}

func (v *Value) UnmarshalPlist(f func(interface{}) error) (err error) {
	var value any
	err = f(&value)
	if err != nil {
		goto end
	}
	v.Value = yamlutil.NewValue(value)
end:
	return err
}

//	switch args.PListKind {
//	case uint(plist.String):
//		var value string
//		err = f(&value)
//		if err != nil {
//			goto end
//		}
//		v.Value.Value = value
//	case uint(plist.Date):
//		var value time.Time
//		err = f(&value)
//		v.Value.Value = value
//	case uint(plist.Integer):
//		var value int
//		err = f(&value)
//		if err != nil {
//			goto end
//		}
//		v.Value.Value = value
//	case uint(plist.Real):
//		var value float64
//		err = f(&value)
//		if err != nil {
//			goto end
//		}
//		v.Value.Value = value
//	case uint(plist.Boolean):
//		var value bool
//		err = f(&value)
//		if err != nil {
//			goto end
//		}
//		v.Value.Value = value
//	case uint(plist.Array):
//		// THIS CURRENTLY ONLY HANDLES []string{}
//		// NEED A USE-CASE TO IDENTIFY HOW TO HANDLE OTHER Subtypes
//		value := make([]string, 0)
//		err = f(&value)
//		if err != nil {
//			goto end
//		}
//		v.Value.Value = value
//	case uint(plist.Dictionary):
//		if args.SubKeyCount == 0 {
//			v.Value.Value = make(map[string]any)
//			goto end
//		}
//		switch args.PListSubkind {
//		case uint(plist.String):
//			value := make(map[string]string)
//			err = f(&value)
//			if err != nil {
//				goto end
//			}
//			v.Value.Value = value
//		case uint(plist.Boolean):
//			value := make(map[string]bool)
//			err = f(&value)
//			if err != nil {
//				goto end
//			}
//			v.Value.Value = value
//		default:
//			panicf("PList SUBkind '%d' is not valid", args.PListSubkind)
//		}
//	case uint(plist.Data):
//		panicf("PList kind '%d' is not (yet) supported", args.PListKind)
//	case uint(plist.Invalid):
//		fallthrough
//	default:
//		panicf("PList kind '%d' is not valid", args.PListKind)
//	}
//end:
//	return err
//}
