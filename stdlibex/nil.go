package stdlibex

import (
	"reflect"
)

// ContainsNil returns true if an interface contains nil (vs. ==nil)
func ContainsNil(value any) (contains bool) {
	if value == nil {
		contains = true
		goto end
	}
	//goland:noinspection GoSwitchMissingCasesForIotaConsts
	switch reflect.TypeOf(value).Kind() {
	case reflect.Ptr:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Array:
		fallthrough
	case reflect.Chan:
		fallthrough
	case reflect.Slice:
		contains = reflect.ValueOf(value).IsNil()
	}
end:
	return contains
}
