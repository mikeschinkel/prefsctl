package macprefs

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation
#import <CoreFoundation/CoreFoundation.h>

// CFPreferencesCopyApplicationList is deprecated but provides cleanest API for our use-case.
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CFArrayRef GetPreferenceDomains() {
	return CFPreferencesCopyApplicationList(kCFPreferencesCurrentUser, kCFPreferencesAnyHost);
}

CFArrayRef GetKeysForDomain(CFStringRef domain) {
    return CFPreferencesCopyKeyList(
        domain,
        kCFPreferencesCurrentUser,
        kCFPreferencesAnyHost
    );
}

int GetArrayCount(CFArrayRef array) {
	 return (int)CFArrayGetCount(array);
}

CFStringRef GetArrayValueAtIndex(CFArrayRef array, int index) {
	 return (CFStringRef)CFArrayGetValueAtIndex(array, index);
}

char* CFStringToCString(CFStringRef str) {
	 if (str == NULL) return NULL;

	 CFIndex length = CFStringGetLength(str);
	 CFIndex maxSize = CFStringGetMaximumSizeForEncoding(length, kCFStringEncodingUTF8) + 1;
	 char *buffer = (char *)malloc(maxSize);

	 if (CFStringGetCString(str, buffer, maxSize, kCFStringEncodingUTF8)) {
			 return buffer;
	 }

	 free(buffer);
	 return NULL;
}

CFStringRef CreateCFString(const char* str) {
    return CFStringCreateWithCString(NULL, str, kCFStringEncodingUTF8);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Domain represents a preference domain in macOS
type Domain struct {
	Name string
}

// Prefs returns all available preferences for this domain
func (d *Domain) Prefs() ([]*Pref, error) {
	// Convert domain name to CFString
	cfDomain := C.CreateCFString(C.CString(d.Name))
	if cfDomain == 0 {
		return nil, fmt.Errorf("failed to create CFString from domain name")
	}
	defer C.CFRelease(C.CFTypeRef(cfDomain))

	// Get the array of keys
	keysArray := C.GetKeysForDomain(cfDomain)
	if keysArray == 0 {
		return nil, fmt.Errorf("failed to get preference keys for domain %s", d.Name)
	}
	defer C.CFRelease(C.CFTypeRef(keysArray))

	// Get the count of keys
	count := int(C.GetArrayCount(keysArray))
	prefs := make([]*Pref, 0, count)

	// Iterate through the array and create Pref objects
	for i := 0; i < count; i++ {
		cfStr := C.GetArrayValueAtIndex(keysArray, C.int(i))
		if cfStr == 0 {
			continue
		}

		cPref := C.CFStringToCString(cfStr)
		if cPref == nil {
			continue
		}
		name := C.GoString(cPref)
		// TODO Lookup PrefDefault here instead
		pref := &Pref{
			Name:  name,
			Value: "",
			Kind:  0,
			Default: &PrefDefault{
				domain: *d,
				Name:   name,
			},
		}
		prefs = append(prefs, pref)
		C.free(unsafe.Pointer(cPref))
	}

	return prefs, nil
}

// GetPreferenceDomains returns a list of all preference domains available
// for the current user on macOS.
func GetPreferenceDomains() ([]*Domain, error) {
	// Get the CFArray of preference domains
	domainsArray := C.GetPreferenceDomains()
	if domainsArray == 0 {
		return nil, fmt.Errorf("failed to get preference domains")
	}
	defer C.CFRelease(C.CFTypeRef(domainsArray))

	// Get the count of domains
	count := int(C.GetArrayCount(domainsArray))
	domains := make([]*Domain, 0, count)

	// Iterate through the array and convert each CFString to Go string
	for i := 0; i < count; i++ {
		domain := C.GetArrayValueAtIndex(domainsArray, C.int(i))
		if domain == 0 {
			continue
		}

		cDomain := C.CFStringToCString(domain)
		if cDomain == nil {
			continue
		}

		domains = append(domains, &Domain{
			Name: C.GoString(cDomain),
		})

		C.free(unsafe.Pointer(cDomain))
	}

	return domains, nil
}
