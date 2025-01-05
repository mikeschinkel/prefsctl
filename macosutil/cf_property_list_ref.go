package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import (
	"unsafe"
)

type CFPropertyListRef struct {
	cfValue C.CFPropertyListRef
}

func NewCFPropertyListRef(value string) *CFPropertyListRef {
	var cfValue C.CFPropertyListRef
	var intVal int64
	var floatVal float64
	var ok bool

	if value == "true" || value == "false" {
		cfValue = C.CFPropertyListRef(NewCFBoolean(value == "true").cfBoolean)
		goto end
	}

	intVal, ok = parseInt64(value)
	if ok {
		cfValue = C.CFPropertyListRef(NewCFLongLong(intVal).cfLongLong)
		goto end
	}

	floatVal, ok = parseFloat64(value)
	if ok {
		cfValue = C.CFPropertyListRef(NewCFDouble(floatVal).cfDouble)
		goto end
	}

	// Assume string
	cfValue = C.CFPropertyListRef(NewCFString(value).cfString)

end:
	return &CFPropertyListRef{
		cfValue: cfValue,
	}
}

func (ref *CFPropertyListRef) IsNull() bool {
	return cIsNullRef(unsafe.Pointer(ref.cfValue))
}

func (ref *CFPropertyListRef) Release() {
	C.CFRelease(C.CFTypeRef(ref.cfValue))
	ref.cfValue = 0
}
