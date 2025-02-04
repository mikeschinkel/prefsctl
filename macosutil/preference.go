package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework SystemConfiguration -framework CoreFoundation
#import <Foundation/Foundation.h>
#import <SystemConfiguration/SystemConfiguration.h>

extern CFTypeRef CFPreferencesGetValue(
    CFStringRef key,
    CFStringRef applicationID,
    CFStringRef userName,
    CFStringRef hostName
);

// Match Go's reflect.Kind relevant values
typedef enum {
	 KIND_INVALID = 0,
	 KIND_BOOL = 1,
	 KIND_INT = 6,
	 KIND_FLOAT = 14,
	 KIND_STRING = 24
} PreferenceKind;

// Error codes
int PREF_SUCCESS = 0;
int PREF_INVALID_INPUT = 1;
int PREF_INVALID_DOMAIN = 2;
int PREF_NOT_FOUND = 3;
int PREF_UNSUPPORTED_TYPE = 4;

typedef struct {
	 const char* value;     // UTF-8 string current value, or error message
	 const char* descr;     // UTF-8 string description of value
	 int error;             // 0 for success, error code otherwise
	 PreferenceKind kind;   // Type of the value (valid only if error == 0)
} PreferenceResult;

// Forward declare functions
PreferenceResult getPreferenceResult(const char* domain, const char* name);
void freePreferenceResult(PreferenceResult result);
PreferenceResult getGlobalPreferenceResult(const char* name);
PreferenceResult processPreferenceValue(id value);
bool isDomainValid(CFStringRef domain);

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

		// Check if domain exists
		if (!isDomainValid((__bridge CFStringRef)domainStr)) {
			result.value = strdup("invalid preference domain");
			result.error = PREF_INVALID_DOMAIN;
			return result;
		}

		// Get current value
		id value = CFBridgingRelease(CFPreferencesCopyAppValue(
			(__bridge CFStringRef)keyStr,
			(__bridge CFStringRef)domainStr
		));

		return processPreferenceValue(value);
	}
}

void freePreferenceResult(PreferenceResult result) {
	 if (result.value) free((void*)result.value);
   if (result.descr) free((void*)result.descr);
}

PreferenceResult getGlobalPreferenceResult(const char* name) {
	PreferenceResult result = {NULL, NULL, PREF_SUCCESS, KIND_INVALID};

	@autoreleasepool {
		if (!name) {
			result.value = strdup("invalid input parameters");
			result.error = PREF_INVALID_INPUT;
			return result;
		}

		NSString* keyStr = [NSString stringWithUTF8String:name];

		// Get current value using CFPreferencesCopyValue for global preferences
		id value = CFBridgingRelease(CFPreferencesCopyValue(
			(__bridge CFStringRef)keyStr,
			kCFPreferencesAnyApplication,
			kCFPreferencesCurrentUser,
			kCFPreferencesAnyHost
		));

		return processPreferenceValue(value);
	}
}

PreferenceResult processPreferenceValue(id value) {
	PreferenceResult result = {NULL, NULL, PREF_SUCCESS, KIND_INVALID};

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

  // Get the class name as a string for comparison
  NSString *className = NSStringFromClass([value class]);

  // Handle Dictionary Types
  if ([value isKindOfClass:[NSDictionary class]]) {
      if ([className isEqualToString:@"__NSDictionaryI"]) {
          // Standard immutable dictionary with entries
          // Note: Will need to handle nested key-value pairs
          result.value = strdup("__NSDictionaryI: Not implemented");
          result.descr = strdup([[value description] UTF8String]);
          result.error = PREF_UNSUPPORTED_TYPE;
          return result;
      }
      if ([className isEqualToString:@"__NSDictionary0"]) {
          // Empty immutable dictionary
          // Note: Should probably return an empty dictionary representation
          result.value = strdup("__NSDictionary0: Not implemented");
          result.descr = strdup([[value description] UTF8String]);
          result.error = PREF_UNSUPPORTED_TYPE;
          return result;
      }
      if ([className isEqualToString:@"__NSSingleEntryDictionaryI"]) {
          // Single entry immutable dictionary
          result.value = strdup("__NSSingleEntryDictionaryI: Not implemented");
          result.descr = strdup([[value description] UTF8String]);
          result.error = PREF_UNSUPPORTED_TYPE;
          return result;
      }
  }

  // Handle Array Types
  if ([value isKindOfClass:[NSArray class]]) {
      if ([className isEqualToString:@"__NSArray0"]) {
          // Empty immutable array
          // Note: Should probably return an empty array representation
          result.value = strdup("__NSArray0: Not implemented");
          result.descr = strdup([[value description] UTF8String]);
          result.error = PREF_UNSUPPORTED_TYPE;
          return result;
      }
      if ([className isEqualToString:@"__NSArrayM"]) {
          // Mutable array
          // Note: Contents could theoretically change during processing
          result.value = strdup("__NSArrayM: Not implemented");
          result.descr = strdup([[value description] UTF8String]);
          result.error = PREF_UNSUPPORTED_TYPE;
          return result;
      }
      if ([className isEqualToString:@"__NSCFArray"]) {
          // Core Foundation array bridge
          // Note: May need special handling for CF-bridged content
          result.value = strdup("__NSCFArray: Not implemented");
          result.descr = strdup([[value description] UTF8String]);
          result.error = PREF_UNSUPPORTED_TYPE;
          return result;
      }
  }

  // Handle Date Type
  if ([value isKindOfClass:[NSDate class]] && [className isEqualToString:@"__NSTaggedDate"]) {
      // Note: Consider timezone handling and date format standardization
      result.value = strdup("__NSTaggedDate: Not implemented");
      result.descr = strdup([[value description] UTF8String]);
      result.error = PREF_UNSUPPORTED_TYPE;
      return result;
  }

  // Handle Data Type
  if ([value isKindOfClass:[NSData class]] && [className isEqualToString:@"__NSCFData"]) {
      // Note: May need to handle binary data conversion
      result.value = strdup("__NSCFData: Not implemented");
      result.descr = strdup([[value description] UTF8String]);
      result.error = PREF_UNSUPPORTED_TYPE;
      return result;
  }

	// Handle unsupported types
	NSString *typeDesc = [value description];
	NSString *errorMsg = [NSString stringWithFormat:@"unsupported preference class: %@", className];

	result.value = strdup([errorMsg UTF8String]);
	result.descr = strdup([typeDesc UTF8String]);
	result.error = PREF_UNSUPPORTED_TYPE;
	return result;
}

bool isDomainValid(CFStringRef domain) {
    CFArrayRef keyList = CFPreferencesCopyKeyList(domain, kCFPreferencesCurrentUser, kCFPreferencesAnyHost);
    if (keyList != NULL) {
        CFRelease(keyList);
        return true;
    }

    // Try with any user to catch system domains
    keyList = CFPreferencesCopyKeyList(domain, kCFPreferencesAnyUser, kCFPreferencesCurrentHost);
    if (keyList != NULL) {
        CFRelease(keyList);
        return true;
    }

    return false;
}

void setPreference(CFStringRef domain, CFStringRef key, CFPropertyListRef value) {
    CFPreferencesSetValue(key, value, domain, kCFPreferencesCurrentUser, kCFPreferencesAnyHost);
}

*/
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

const GlobalPreferencesDomain = ".GlobalPreferences"

type Preference struct {
	Domain      string
	Name        string
	Value       string
	Description string
	Kind        reflect.Kind
	invalid     bool
	err         error
}

func (p *Preference) Valid() bool {
	return !p.invalid
}
func (p *Preference) Id() Identifier {
	return PreferenceId(p.Domain, p.Name)
}

type ApplyOpts struct {
	SkipSync bool
}

// ApplyPreference applies a single preference for the specified domain
func (p *Preference) ApplyPreference(opts ApplyOpts) (err error) {

	// Set the preference
	err = SetPreference(p.Domain, p.Name, p.Value)
	if err != nil {
		goto end
	}
	if !opts.SkipSync {
		// Sync preferences to make them persistent
		p.SyncPreference()
	}

end:
	return err

}

func (p *Preference) SyncPreference() {
	SyncDomainPreferences(p.Domain)
}

type Identifier string

func PreferenceId(domain string, name string) Identifier {
	return Identifier(domain + "/" + name)
}

var preferenceCache = make(map[Identifier]*Preference)
var preferenceCacheMutex sync.Mutex

// GetPreference fetches the preference value from the system
func (*macOSUtils) GetPreference(domain string, name string) (dp *Preference, err error) {
	var ok bool
	var cDomain, cName *C.char
	var cResult C.PreferenceResult

	preferenceCacheMutex.Lock()
	prefId := PreferenceId(domain, name)
	if dp, ok = preferenceCache[prefId]; ok {
		goto end
	}
	cDomain = C.CString(domain)
	cName = C.CString(name)

	defer C.free(unsafe.Pointer(cDomain))
	defer C.free(unsafe.Pointer(cName))

	if domain == GlobalPreferencesDomain {
		cResult = C.getGlobalPreferenceResult(cName)
	} else {
		cResult = C.getPreferenceResult(cDomain, cName)
	}
	defer C.freePreferenceResult(cResult)

	dp = &Preference{
		Domain: domain,
		Name:   name,
	}

	if cResult.error == C.PREF_SUCCESS {
		dp.Kind = reflect.Kind(cResult.kind)
	} else {
		dp.Value = C.GoString(cResult.value) // Save error message
		dp.Description = C.GoString(cResult.descr)
		dp.invalid = true
		switch cResult.error {
		case C.PREF_INVALID_INPUT:
			err = ErrInvalidInput
		case C.PREF_INVALID_DOMAIN:
			err = ErrInvalidPreferenceDomain
		case C.PREF_NOT_FOUND:
			err = ErrPreferenceNotFound
		case C.PREF_UNSUPPORTED_TYPE:
			err = ErrUnsupportedType
			err = ErrUnsupportedType
		default:
			err = ErrUnknownPreferenceError
		}
		goto end
	}
	dp.Value = C.GoString(cResult.value)
	preferenceCache[prefId] = dp
end:
	preferenceCacheMutex.Unlock()
	return dp, err
}

func (*macOSUtils) SetPreference(domain, name, value string) (err error) {
	var cfDomain, cfName *CFString
	var cfValue *CFPropertyListRef

	cfDomain = NewCFString(domain)
	if cfDomain.IsNull() {
		err = errors.Join(ErrFailedToCreateCFString, fmt.Errorf("domain=%s", domain))
		goto end
	}
	defer cfDomain.Release()

	cfName = NewCFString(name)
	if cfName.IsNull() {
		err = errors.Join(ErrFailedToCreateCFString, fmt.Errorf("key=%s", name))
		goto end
	}
	defer cfName.Release()

	cfValue = NewCFPropertyListRef(value)
	if cfValue.IsNull() {
		err = errors.Join(ErrFailedToCreateCFString, fmt.Errorf("value=%s", value))
		goto end
	}
	defer cfValue.Release()

	C.setPreference(
		cfDomain.cfString,
		cfName.cfString,
		cfValue.cfValue,
	)

end:
	return err
}
