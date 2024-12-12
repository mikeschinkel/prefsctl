package prefdefaults

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macprefs"
)

var floatPattern = regexp.MustCompile(`^([-+])?((\d*)\.?(\d+)|(\d+)\.?(\d*))$`)
var ErrNotAFloat = fmt.Errorf("string value does not represent a float value")

func FloatPrecision(f string) (prec int, err error) {
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

func ParsePrefValue(value string) (kind reflect.Kind, label *kvfilters.Label) {
	switch value {
	case "":
		kind = reflect.Invalid
		label = &macprefs.UnknownType
	case "true", "false":
		kind = reflect.Bool
		label = &macprefs.BoolType
	case "0", "1":
		kind = reflect.Int64
		label = &macprefs.IntBoolType
	default:
		var err error
		var prec int
		n, _ := strconv.ParseInt(value, 10, 64)
		if value == strconv.FormatInt(n, 10) {
			kind = reflect.Int64
			label = &macprefs.IntType
			goto end
		}
		prec, err = FloatPrecision(value)
		if err != nil {
			kind = reflect.String
			label = &macprefs.StringType
			goto end
		}
		f, _ := strconv.ParseFloat(value, 64)
		if value == strconv.FormatFloat(f, 'f', prec, 64) {
			kind = reflect.Float64
			label = &macprefs.FloatType
			goto end
		}
		kind = reflect.String
		label = &macprefs.StringType
	}
end:
	return kind, label
}

func ParsePrefKind(kind reflect.Kind) (_ reflect.Kind, label *kvfilters.Label) {
	switch kind {
	case reflect.Int64:
		label = &macprefs.IntType
	case reflect.String:
		label = &macprefs.StringType
	case reflect.Bool:
		label = &macprefs.BoolType
	case reflect.Float64:
		label = &macprefs.FloatType
	default:
		label = &macprefs.UnknownType
	}
	return kind, label
}

func ParsePrefType(typ TypeName) (kind reflect.Kind, label *kvfilters.Label) {
	switch kvfilters.LabelValue(typ) {
	case StringType.Value:
		kind = reflect.String
		label = &macprefs.StringType
	case BoolType.Value:
		kind = reflect.Bool
		label = &macprefs.BoolType
	case IntBoolType.Value:
		kind = reflect.Int64
		label = &macprefs.IntBoolType
	case IntType.Value:
		kind = reflect.Int64
		label = &macprefs.IntType
	case FloatType.Value:
		kind = reflect.Float64
		label = &macprefs.FloatType
	default:
		kind = reflect.Invalid
		label = &macprefs.UnknownType
	}
	return kind, label
}

func GetPrefKind(typ TypeName, value string) (kind reflect.Kind) {
	if typ == "" {
		kind, _ = ParsePrefValue(value)
		goto end
	}
	kind, _ = ParsePrefType(typ)
end:
	return kind
}

func GetPrefKindAndTypeLabel(kind reflect.Kind, typ TypeName, value string) (_ reflect.Kind, label *kvfilters.Label) {
	// Type is manually set/managed so it gets preference
	if typ != "" && kvfilters.LabelValue(typ) != macprefs.UnknownType.Value {
		kind, label = ParsePrefType(typ)
		goto end
	}
	kind, label = ParsePrefValue(value)
	switch {
	case kind == reflect.Invalid:
	case label == IntBoolType:
	default:
		kind, label = ParsePrefKind(kind)
	}
end:
	return kind, label
}
