package macprefs

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
	"strconv"
	"unsafe"
)

// YAMLPref represents a preference with its Domain and name
type YAMLPref struct {
	Name  string   `yaml:"name"`
	Value string   `yaml:"value"` // raw string value
	Kind  PrefType `yaml:"kind"`  // type of the value
}

// Pref represents a preference with its Domain and name
type Pref struct {
	*PrefDefault
	Value       string       // raw string value
	Kind        reflect.Kind // kind of the value
	err         error        // last error encountered
	Description string
}

// NewPref creates a new Pref instance
func NewPref(pd *PrefDefault) *Pref {
	return &Pref{
		PrefDefault: pd,
	}
}

// Retrieve fetches the preference value from the system
func (p *Pref) Retrieve() error {
	cDomain := C.CString(string(p.Domain))
	cKey := C.CString(p.Name)
	defer C.free(unsafe.Pointer(cDomain))
	defer C.free(unsafe.Pointer(cKey))

	result := C.getPreferenceResult(cDomain, cKey)
	defer C.freePreferenceResult(result)

	if result.error != C.PREF_SUCCESS {
		p.Value = C.GoString(result.value) // Save error message
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

	p.Value = C.GoString(result.value)
	p.Kind = reflect.Kind(result.kind)
	p.err = nil
	return nil
}

// IsDefault returns true if the property's current value is the same as its default value.
//func (p *Pref) IsDefault() (isDefault bool) {
//	if p.err != nil {
//		isDefault = false
//		goto end
//	}
//	switch p.Kind {
//	case reflect.String:
//		isDefault = p.Value == p.DefaultValue
//	case reflect.Int:
//		isDefault = p.Int() == p.DefaultInt()
//	case reflect.Bool:
//		isDefault = p.Value == p.DefaultValue
//		if !isDefault && p.DefaultValue == "" {
//			isDefault = true
//		}
//	}
//end:
//	return isDefault
//}

// String returns the raw string value regardless of type
func (p *Pref) String() string {
	return p.Value
}
func (p *Pref) Message() string {
	if p.err != nil {
		return p.Value
	}
	return ""
}

func (p *Pref) Bool() bool {
	return p.Value == "true"
}

// Int returns the value as an integer
func (p *Pref) Int() int {
	n, _ := strconv.Atoi(p.Value)
	return n
}

// Err returns the last error encountered
func (p *Pref) Err() error {
	return p.err
}
