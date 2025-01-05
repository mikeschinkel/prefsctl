package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>

CFNumberRef createCFNumberFromLongLong(long long value) {
    CFNumberRef number = CFNumberCreate(NULL, kCFNumberLongLongType, &value);
    return number;
}

*/
import "C"
import (
	"unsafe"
)

type CFLongLong struct {
	cfLongLong C.CFNumberRef
}

func NewCFLongLong(ll int64) *CFLongLong {
	return &CFLongLong{
		cfLongLong: C.createCFNumberFromLongLong(C.longlong(ll)),
	}
}

func (ll *CFLongLong) CFLongLong() C.CFNumberRef {
	return ll.cfLongLong
}

func (ll *CFLongLong) IsNull() bool {
	return cIsNullRef(unsafe.Pointer(ll.cfLongLong))
}

func (ll *CFLongLong) Pointer() unsafe.Pointer {
	return unsafe.Pointer(ll.cfLongLong)
}

func (ll *CFLongLong) Release() {
	C.CFRelease(C.CFTypeRef(ll.cfLongLong))
	ll.cfLongLong = 0
}
