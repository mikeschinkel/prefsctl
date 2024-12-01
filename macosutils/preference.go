package macosutils

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework SystemConfiguration -framework CoreFoundation
#import <Foundation/Foundation.h>
#import <SystemConfiguration/SystemConfiguration.h>

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
	 const char* value;     // UTF-8 string current value, or error message
	 const char* descr;     // UTF-8 string description of value
	 int error;             // 0 for success, error code otherwise
	 PreferenceKind kind;   // Type of the value (valid only if error == 0)
} PreferenceResult;

PreferenceResult getPreferenceResult(const char* domain, const char* name);
void freePreferenceResult(PreferenceResult result);

PreferenceResult getPreferenceResult(const char* domain, const char* name) {
	 PreferenceResult result = {NULL, NULL, PREF_SUCCESS, KIND_INVALID};

	 @autoreleasepool {
			 if (!domain || !name) {
					 result.value = strdup("invalid input parameters");
					 result.error = PREF_INVALID_INPUT;
					 return result;
			 }

			 NSString* domainStr = [NSString stringWithUTF8String:domain];
			 NSString* keyStr = [NSString stringWithUTF8String:name];

			// Get current value
			id value = CFBridgingRelease(CFPreferencesCopyAppValue(
				 (__bridge CFStringRef)keyStr,
				 (__bridge CFStringRef)domainStr
			));

			// Try different methods to get default value
			id defaultValue = NULL;

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
	"unsafe"
)

type Preference struct {
	Domain      string
	Name        string
	Value       string
	Description string
	Kind        reflect.Kind
	invalid     bool
}

func (p Preference) Valid() bool {
	return !p.invalid
}

// RetrievePreference fetches the preference value from the system
func RetrievePreference(domain string, name string) (dp *Preference, err error) {
	cDomain := C.CString(domain)
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cDomain))
	defer C.free(unsafe.Pointer(cName))

	result := C.getPreferenceResult(cDomain, cName)
	defer C.freePreferenceResult(result)

	dp = &Preference{
		Domain: domain,
		Name:   name,
	}

	if result.error == C.PREF_SUCCESS {
		dp.Kind = reflect.Kind(result.kind)
	} else {
		dp.Value = C.GoString(result.value) // Save error message
		dp.Description = C.GoString(result.descr)
		dp.invalid = true
		switch result.error {
		case C.PREF_INVALID_INPUT:
			err = ErrInvalidInput
		case C.PREF_NOT_FOUND:
			err = ErrNotFound
		case C.PREF_UNSUPPORTED_TYPE:
			err = ErrUnsupportedType
		default:
			err = ErrUnknownPreferenceError
		}
		goto end
	}
	dp.Value = C.GoString(result.value)

end:
	return dp, err
}
