package macosutil

/*
 #cgo CFLAGS: -x objective-c
 #cgo LDFLAGS: -framework CoreFoundation -framework Foundation
 #include <CoreFoundation/CoreFoundation.h>
 #include <stddef.h>

 // Helper function for null checks
 bool isNullRef(void* ref) {
     return ref == NULL;
 }
*/
import "C"
import (
	"fmt"
	"strconv"
	"unsafe"
)

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

// parseInt64 checks if a string represents an integer value
func parseInt64(str string) (int64, bool) {
	value, err := strconv.ParseInt(str, 10, 64)
	return value, err == nil
}

// parseFloat64 checks if a string represents a floating point value
func parseFloat64(str string) (float64, bool) {
	value, err := strconv.ParseFloat(str, 64)
	return value, err == nil
}

func cIsNullRef(p unsafe.Pointer) bool {
	return bool(C.isNullRef(p))
}

func cfReleaseCFTypeRef(ref C.CFTypeRef) {
	C.CFRelease(C.CFTypeRef(ref))
}
