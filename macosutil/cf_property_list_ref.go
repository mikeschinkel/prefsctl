package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation -framework Foundation
#include <CoreFoundation/CoreFoundation.h>
#include <stddef.h>

CFBooleanRef createCFBoolean(bool value) {
    return value ? kCFBooleanTrue : kCFBooleanFalse;
}

CFNumberRef createCFNumberFromLongLong(long long value) {
    CFNumberRef number = CFNumberCreate(NULL, kCFNumberLongLongType, &value);
    return number;
}

CFNumberRef createCFNumberFromDouble(double value) {
    CFNumberRef number = CFNumberCreate(NULL, kCFNumberDoubleType, &value);
    return number;
}
*/
import "C"
import (
	"strconv"
	"unsafe"
)

type CFPropertyListRef struct {
	cfValue C.CFPropertyListRef
}

func NewCFPropertyListRef(value string) *CFPropertyListRef {
	var cfValue C.CFPropertyListRef
	switch {
	case value == "true" || value == "false":
		boolVal := value == "true"
		cfValue = C.CFPropertyListRef(C.createCFBoolean(C.bool(boolVal)))

	case isInteger(value):
		val, _ := strconv.ParseInt(value, 10, 64)
		cfValue = C.CFPropertyListRef(C.createCFNumberFromLongLong(C.longlong(val)))

	case isFloat(value):
		val, _ := strconv.ParseFloat(value, 64)
		cfValue = C.CFPropertyListRef(C.createCFNumberFromDouble(C.double(val)))

	default:
		cfValue = C.CFPropertyListRef(NewCFString(value).cfString)
	}
	return &CFPropertyListRef{cfValue: cfValue}
}

func (ref *CFPropertyListRef) IsNull() bool {
	return cIsNullRef(unsafe.Pointer(ref.cfValue))
}

func (ref *CFPropertyListRef) Release() {
	C.CFRelease(C.CFTypeRef(ref.cfValue))
	ref.cfValue = 0
}
