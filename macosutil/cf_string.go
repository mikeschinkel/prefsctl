package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>

CFStringRef createCFString(const char* str) {
    return CFStringCreateWithCString(NULL, str, kCFStringEncodingUTF8);
}

*/
import "C"
import (
	"unsafe"
)

type CFString struct {
	cfString C.CFStringRef
}

func NewCFString(s string) *CFString {
	return &CFString{
		cfString: C.createCFString(C.CString(s)),
	}
}

func (s *CFString) CFString() C.CFStringRef {
	return s.cfString
}

func (s *CFString) IsNull() bool {
	return cIsNullRef(unsafe.Pointer(s.cfString))
}

func (s *CFString) Pointer() unsafe.Pointer {
	return unsafe.Pointer(s.cfString)
}

func (s *CFString) Release() {
	C.CFRelease(C.CFTypeRef(s.cfString))
	s.cfString = 0
}
