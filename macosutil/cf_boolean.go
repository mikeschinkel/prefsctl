package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>

CFBooleanRef createCFBoolean(bool value) {
    return value ? kCFBooleanTrue : kCFBooleanFalse;
}

*/
import "C"
import (
	"unsafe"
)

type CFBoolean struct {
	cfBoolean C.CFBooleanRef
}

func NewCFBoolean(b bool) *CFBoolean {
	return &CFBoolean{
		cfBoolean: C.createCFBoolean(C.bool(b)),
	}
}

func (s *CFBoolean) CFBoolean() C.CFBooleanRef {
	return s.cfBoolean
}

func (s *CFBoolean) IsNull() bool {
	return cIsNullRef(unsafe.Pointer(s.cfBoolean))
}

func (s *CFBoolean) Pointer() unsafe.Pointer {
	return unsafe.Pointer(s.cfBoolean)
}

func (s *CFBoolean) Release() {
	C.CFRelease(C.CFTypeRef(s.cfBoolean))
	s.cfBoolean = 0
}
