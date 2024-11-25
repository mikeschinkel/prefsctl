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

CFArrayRef GetPrefNamesForDomain(CFStringRef domain) {
    return CFPreferencesCopyKeyList(
        domain,
        kCFPreferencesCurrentUser,
        kCFPreferencesAnyHost
    );
}

int GetArrayLen(CFArrayRef array) {
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
	"strings"
	"unsafe"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logging"
)

// Domain represents a preference domain in macOS
type Domain string

func (d Domain) String() string {
	return string(d)
}
func (d Domain) HasPrefix(prefix string) bool {
	return strings.HasPrefix(string(d), prefix)
}

// Prefs returns all available preferences for this domain
func (d Domain) Prefs() (prefs []*Pref, err error) {
	return GetDomainPrefs(d)
}

// GetPrefDomains returns a list of all preference domains available
// for the current user on macOS.
func GetPrefDomains() ([]Domain, error) {
	// Get the CFArray of preference domains
	domainsArray := C.GetPreferenceDomains()
	if domainsArray == 0 {
		return nil, fmt.Errorf("failed to get preference domains")
	}
	defer C.CFRelease(C.CFTypeRef(domainsArray))

	// Get the count of domains
	count := int(C.GetArrayLen(domainsArray))
	domains := make([]Domain, 0, count)

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

		domains = append(domains, Domain(C.GoString(cDomain)))

		C.free(unsafe.Pointer(cDomain))
	}

	return domains, nil
}

// GetDomainPrefs returns all available preferences for this domain
func GetDomainPrefs(d Domain) (prefs []*Pref, err error) {
	var numPrefs int

	// Convert domain name to CFString
	cfDomain := C.CreateCFString(C.CString(string(d)))
	// TODO Set type here instead of calling func
	var cfaPrefs = C.GetPrefNamesForDomain(cfDomain)
	if cfDomain == 0 {
		err = errutil.AnnotateErr(ErrFailedToCreateCFString, "%s=%s", logging.PrefsDomainLogArg, d)
		goto end
	}
	defer C.CFRelease(C.CFTypeRef(cfDomain))

	// Get the array of keys
	cfaPrefs = C.GetPrefNamesForDomain(cfDomain)
	if cfaPrefs == 0 {
		err = errutil.AnnotateErr(ErrFailedToGetPrefNames, "%s=%s", logging.PrefsDomainLogArg, d)
		goto end
	}
	defer C.CFRelease(C.CFTypeRef(cfaPrefs))

	// Get the numPrefs of keys
	numPrefs = int(C.GetArrayLen(cfaPrefs))
	prefs = make([]*Pref, 0, numPrefs)

	// Iterate through the array and create Pref objects
	for i := 0; i < numPrefs; i++ {
		cfStr := C.GetArrayValueAtIndex(cfaPrefs, C.int(i))
		if cfStr == 0 {
			continue
		}

		cPref := C.CFStringToCString(cfStr)
		if cPref == nil {
			continue
		}
		name := C.GoString(cPref)
		pref := NewPref(PrefArgs{
			Domain:  d,
			Name:    name,
			Value:   "",
			Default: "",
			Labels:  Labels{},
			Kind:    0,
		})
		prefs = append(prefs, pref)
		C.free(unsafe.Pointer(cPref))
	}

end:
	return prefs, nil
}
