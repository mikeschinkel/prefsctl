package macosutil

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func ParsePrefValue(value string) (kind reflect.Kind, typ PreferenceType) {
	switch value {
	case "":
		kind = reflect.Invalid
		typ = UnknownType
	case "true", "false":
		kind = reflect.Bool
		typ = BoolType
	case "0", "1":
		kind = reflect.Int64
		typ = IntBoolType
	default:
		var err error
		var prec int
		n, _ := strconv.ParseInt(value, 10, 64)
		if value == strconv.FormatInt(n, 10) {
			kind = reflect.Int64
			typ = IntType
			goto end
		}
		prec, err = GetFloatPrecision(value)
		if err != nil {
			kind = reflect.String
			typ = StringType
			goto end
		}
		f, _ := strconv.ParseFloat(value, 64)
		if value == strconv.FormatFloat(f, 'f', prec, 64) {
			kind = reflect.Float64
			typ = FloatType
			goto end
		}
		kind = reflect.String
		typ = StringType
	}
end:
	return kind, typ
}

func ParsePrefKind(kind reflect.Kind) (_ reflect.Kind, typ PreferenceType) {
	switch kind {
	case reflect.Int64:
		typ = IntType
	case reflect.String:
		typ = StringType
	case reflect.Bool:
		typ = BoolType
	case reflect.Float64:
		typ = FloatType
	default:
		typ = UnknownType
	}
	return kind, typ
}

func ParsePrefType(pt PreferenceType) (kind reflect.Kind, typ PreferenceType) {
	typ = pt
	switch pt {
	case StringType:
		kind = reflect.String
	case BoolType:
		kind = reflect.Bool
	case IntBoolType:
		kind = reflect.Int64
	case IntType:
		kind = reflect.Int64
	case FloatType:
		kind = reflect.Float64
	default:
		kind = reflect.Invalid
	}
	return kind, typ
}

func GetPrefKind(typ PreferenceType, value string) (kind reflect.Kind) {
	if typ == "" {
		kind, _ = ParsePrefValue(value)
		goto end
	}
	kind, _ = ParsePrefType(typ)
end:
	return kind
}

func GetPrefKindAndType(kind reflect.Kind, typ PreferenceType, value string) (reflect.Kind, PreferenceType) {
	// Type is manually set/managed so it gets preference
	if typ != "" && typ != UnknownType {
		kind, typ = ParsePrefType(typ)
		goto end
	}
	kind, typ = ParsePrefValue(value)
	switch {
	case kind == reflect.Invalid:
	case typ == IntBoolType:
	default:
		kind, typ = ParsePrefKind(kind)
	}
end:
	return kind, typ
}

var floatPattern = regexp.MustCompile(`^([-+])?((\d*)\.?(\d+)|(\d+)\.?(\d*))$`)
var ErrNotAFloat = fmt.Errorf("string value does not represent a float value")

func GetFloatPrecision(f string) (prec int, err error) {
	var matches []string

	switch f {
	case "", ".", "+.", "-.", "-", "+":
		err = ErrNotAFloat
		goto end
	}
	matches = floatPattern.FindStringSubmatch(f)
	if len(matches) < 6 {
		err = ErrNotAFloat
		goto end
	}
	if !strings.Contains(f, ".") {
		goto end
	}
	prec = max(len(matches[4]), len(matches[6]))
end:
	if err != nil {
		err = errors.Join(err, fmt.Errorf("value=%s", f))
	}
	return prec, err
}
