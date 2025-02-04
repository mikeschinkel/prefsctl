package macosutil

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func GetType(kind reflect.Kind) (_ reflect.Kind, typ PreferenceType) {
	switch kind {
	case reflect.Int64:
		typ = IntegerType
	case reflect.String:
		typ = StringType
	case reflect.Bool:
		typ = BooleanType
	case reflect.Slice:
		typ = ArrayType
	default:
		typ = InvalidType
	}
	return kind, typ
}

func GetKind(typ PreferenceType) (kind reflect.Kind) {
	switch typ {
	case StringType:
		kind = reflect.String
	case BooleanType:
		kind = reflect.Bool
	case IntegerType:
		kind = reflect.Int64
	case ArrayType:
		kind = reflect.Slice
	default:
		kind = reflect.Invalid
	}
	return kind
}

var floatPattern = regexp.MustCompile(`^([-+])?((\d*)\.?(\d+)|(\d+)\.?(\d*))$`)
var ErrNotAFloat = fmt.Errorf("string value does not represent a float value")

func GetFloatPrecisionSAVE(f string) (prec int, err error) {
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
func GetFloatPrecision(f string) (prec int, err error) {
	var matches []string

	defer func() {
		if err != nil {
			err = errors.Join(err, fmt.Errorf("value=%s", f))
		}
	}()

	switch f {
	case "", ".", "+.", "-.", "-", "+":
		err = ErrNotAFloat
		return
	}
	matches = floatPattern.FindStringSubmatch(f)
	if len(matches) < 6 {
		err = ErrNotAFloat
		return
	}
	if !strings.Contains(f, ".") {
		return
	}
	prec = max(len(matches[4]), len(matches[6]))

	return prec, err
}
