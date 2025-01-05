package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation -framework Foundation
#include <CoreFoundation/CoreFoundation.h>

CFNumberRef createCFNumberFromDouble(double value) {
    CFNumberRef number = CFNumberCreate(NULL, kCFNumberDoubleType, &value);
    return number;
}
*/
import "C"
import (
	"unsafe"
)

type CFDouble struct {
	cfDouble C.CFNumberRef
}

func NewCFDouble(d float64) *CFDouble {
	return &CFDouble{
		cfDouble: C.createCFNumberFromDouble(C.double(d)),
	}
}

func (d *CFDouble) CFDouble() C.CFNumberRef {
	return d.cfDouble
}

func (d *CFDouble) IsNull() bool {
	return cIsNullRef(unsafe.Pointer(d.cfDouble))
}

func (d *CFDouble) Pointer() unsafe.Pointer {
	return unsafe.Pointer(d.cfDouble)
}

func (d *CFDouble) Release() {
	C.CFRelease(C.CFTypeRef(d.cfDouble))
	d.cfDouble = 0
}
