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

// isInteger checks if a string represents an integer value
func isInteger(str string) bool {
	_, err := strconv.ParseInt(str, 10, 64)
	return err == nil
}

// isFloat checks if a string represents a floating point value
func isFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil && !isInteger(str)
}

func cIsNullRef(p unsafe.Pointer) bool {
	return bool(C.isNullRef(p))
}

func cfReleaseCFTypeRef(ref C.CFTypeRef) {
	C.CFRelease(C.CFTypeRef(ref))
}
