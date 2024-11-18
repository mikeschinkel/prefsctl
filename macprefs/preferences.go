package macprefs

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

// Match Go's reflect.Kind relevant values
typedef enum {
	 KIND_INVALID = 0,
	 KIND_BOOL = 1,
	 KIND_INT = 2,
	 KIND_STRING = 24
} PreferenceKind;

// Error codes
int PREF_SUCCESS = 0;
int PREF_INVALID_INPUT = 1;
int PREF_NOT_FOUND = 2;
int PREF_UNSUPPORTED_TYPE = 3;

typedef struct {
	 const char* value;     // UTF-8 string value or error message
	 const char* descr;     // UTF-8 string description of value
	 int error;             // 0 for success, error code otherwise
	 PreferenceKind kind;   // Type of the value (valid only if error == 0)
} PreferenceResult;

PreferenceResult getPreferenceValue(const char* domain, const char* key);
void freePreferenceResult(PreferenceResult result);

PreferenceResult getPreferenceValue(const char* domain, const char* key) {
	 PreferenceResult result = {NULL, PREF_SUCCESS, KIND_INVALID};

	 @autoreleasepool {
			 if (!domain || !key) {
					 result.value = strdup("invalid input parameters");
					 result.error = PREF_INVALID_INPUT;
					 return result;
			 }

			 NSString* domainStr = [NSString stringWithUTF8String:domain];
			 NSString* keyStr = [NSString stringWithUTF8String:key];

			id value = CFBridgingRelease(CFPreferencesCopyAppValue(
				 (__bridge CFStringRef)keyStr,
				 (__bridge CFStringRef)domainStr
			));

			 if (!value) {
					 result.value = strdup("preference not found");
					 result.error = PREF_NOT_FOUND;
					 return result;
			 }

			 // Handle strings
			 if ([value isKindOfClass:[NSString class]]) {
					 result.value = strdup([value UTF8String]);
					 result.kind = KIND_STRING;
					 return result;
			 }

			 // Handle numbers (including booleans)
			 if ([value isKindOfClass:[NSNumber class]]) {
					 NSNumber* number = (NSNumber*)value;
					 const char* objCType = [number objCType];

					 // Check if it's a boolean
					 if (strcmp(objCType, @encode(BOOL)) == 0) {
							 result.value = strdup([number boolValue] ? "true" : "false");
							 result.kind = KIND_BOOL;
							 return result;
					 }

					 // Handle other numbers as integers
					 result.value = strdup([[number stringValue] UTF8String]);
					 result.kind = KIND_INT;
					 return result;
			 }

			 // Handle unsupported types
			//result.value = strdup("unsupported preference type");
			NSString *className = NSStringFromClass([value class]);
			NSString *typeDesc = [value description];
			NSString *errorMsg = [NSString stringWithFormat:@"unsupported preference class: %@",className];

			result.value = strdup([errorMsg UTF8String]);
			result.descr = strdup([typeDesc UTF8String]);
			result.error = PREF_UNSUPPORTED_TYPE;
			return result;
	 }
}

void freePreferenceResult(PreferenceResult result) {
	 if (result.value) free((void*)result.value);
}
*/
import "C"
import (
	"reflect"
	"strconv"
	"unsafe"
)

// Pref represents a preference with its Domain and key
type Pref struct {
	domain      *Domain
	Key         string
	Description string
	value       string       // raw string value
	kind        reflect.Kind // type of the value
	err         error        // last error encountered
}

// NewPref creates a new Pref instance
func NewPref(domain *Domain, key string) *Pref {
	return &Pref{
		Key:    key,
		domain: domain,
	}
}

// Retrieve fetches the preference value from the system
func (p *Pref) Retrieve() error {
	cDomain := C.CString(p.domain.Name)
	cKey := C.CString(p.Key)
	defer C.free(unsafe.Pointer(cDomain))
	defer C.free(unsafe.Pointer(cKey))

	result := C.getPreferenceValue(cDomain, cKey)
	defer C.freePreferenceResult(result)

	if result.error != C.PREF_SUCCESS {
		p.value = C.GoString(result.value) // Save error message
		p.Description = C.GoString(result.descr)
		switch result.error {
		case C.PREF_INVALID_INPUT:
			p.err = ErrInvalidInput
		case C.PREF_NOT_FOUND:
			p.err = ErrNotFound
		case C.PREF_UNSUPPORTED_TYPE:
			p.err = ErrUnsupportedType
		default:
			p.err = ErrUnknown
		}
		return p.err
	}

	p.value = C.GoString(result.value)
	p.kind = reflect.Kind(result.kind)
	p.err = nil
	return nil
}

// String returns the raw string value regardless of type
func (p *Pref) String() string {
	return p.value
}
func (p *Pref) Message() string {
	if p.err != nil {
		return p.value
	}
	return ""
}
func (p *Pref) Bool() bool {
	return p.value == "true"
}

// Int returns the value as an integer
func (p *Pref) Int() int {
	n, _ := strconv.Atoi(p.value)
	return n
}

// Err returns the last error encountered
func (p *Pref) Err() error {
	return p.err
}

// Kind returns the preference value type
func (p *Pref) Kind() reflect.Kind {
	return p.kind
}
