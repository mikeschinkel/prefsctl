package stdlibex

import (
	"fmt"
	r "reflect"
	"strings"
)

func Empty(value any) (empty bool) {
	switch x := value.(type) {
	case nil:
		return true
	case string:
		return strings.TrimSpace(x) == ""
	case *string:
		return strings.TrimSpace(*x) == ""
	case bool:
		return x == false
	case *bool:
		return *x == false
	case int, int64, int32, int16, int8:
		return x == 0
	case *int:
		return *x == 0
	case *int64:
		return *x == 0
	case *int32:
		return *x == 0
	case *int16:
		return *x == 0
	case *int8:
		return *x == 0
	case uint, uint64, uint32, uint16, uint8:
		return x == 0
	case *uint:
		return *x == 0
	case *uint64:
		return *x == 0
	case *uint32:
		return *x == 0
	case *uint16:
		return *x == 0
	case *uint8:
		return *x == 0
	case float32, float64:
		return x == 0
	case *float32:
		return *x == 0
	case *float64:
		return *x == 0
	case complex64, complex128:
		return x == 0
	case *complex64:
		return *x == 0
	case *complex128:
		return *x == 0
	case uintptr:
		return x == 0
	default:
		v := r.ValueOf(x)
		switch v.Kind() {
		case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
			return v.Int() == 0
		case r.String:
			return Empty(v.String())
		case r.Array:
			return v.Len() == 0
		case r.Slice, r.Ptr, r.Interface, r.Map, r.Chan, r.Func:
			return !v.IsNil()
		default:
			panic(fmt.Sprintf("unsupported type '%T' of value %v", x, x))
		}
	}
	return false
}
