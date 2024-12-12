package macosutils

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation
#import <CoreFoundation/CoreFoundation.h>

// CFPreferencesCopyApplicationList is deprecated but provides cleanest API for our use-case.
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

//CFArrayRef GetPreferenceDomains() {
//	return CFPreferencesCopyApplicationList(kCFPreferencesCurrentUser, kCFPreferencesAnyHost);
//}
CFArrayRef GetPreferenceDomains() {
    // Get the list of application domains
    CFArrayRef appDomains = CFPreferencesCopyApplicationList(kCFPreferencesCurrentUser, kCFPreferencesAnyHost);

    // Create a mutable array from the application domains
    CFMutableArrayRef allDomains = CFArrayCreateMutableCopy(kCFAllocatorDefault,
                                                           CFArrayGetCount(appDomains) + 1,
                                                           appDomains);

    // Add .GlobalPreferences
    CFStringRef globalDomain = CFSTR("GlobalPreferences");
    CFArrayAppendValue(allDomains, globalDomain);

    // Release the original array and the global domain string
    CFRelease(appDomains);
    CFRelease(globalDomain);

    return allDomains;
}

//CFArrayRef GetPrefNamesForDomain(CFStringRef domain) {
//    return CFPreferencesCopyKeyList(
//        domain,
//        kCFPreferencesCurrentUser,
//        kCFPreferencesAnyHost
//    );
//}
CFArrayRef GetPrefNamesForDomain(CFStringRef domain) {
    CFStringRef useDomain = domain;
    if (CFStringCompare(domain, CFSTR("GlobalPreferences"), 0) == kCFCompareEqualTo) {
        useDomain = kCFPreferencesAnyApplication;
    }
    return CFPreferencesCopyKeyList(
        useDomain,
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
	"unsafe"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logargs"
)

type PreferenceDomain string

// retrievePreferenceDomains returns a list of all preference domains available
// for the current user on macOS.
func retrievePreferenceDomains() (domains []PreferenceDomain, err error) {
	var count int

	// Get the CFArray of preference domains
	domainsArray := C.GetPreferenceDomains()
	if domainsArray == 0 {
		err = fmt.Errorf("failed to get preference domains")
		goto end
	}
	defer C.CFRelease(C.CFTypeRef(domainsArray))

	// Get the count of domains
	count = int(C.GetArrayLen(domainsArray))
	domains = make([]PreferenceDomain, 0, count)

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
		domains = append(domains, PreferenceDomain(C.GoString(cDomain)))

		C.free(unsafe.Pointer(cDomain))
	}
end:
	return domains, err
}

// retrievePreferences returns all available preferences for this domain
func retrievePreferences(d PreferenceDomain) (prefs []*Preference, err error) {
	var numPrefs int

	// Convert domain name to CFString
	cfDomain := C.CreateCFString(C.CString(string(d)))
	// TODO Set type here instead of calling func
	var cfaPrefs = C.GetPrefNamesForDomain(cfDomain)
	if cfDomain == 0 {
		err = errutil.AnnotateErr(ErrFailedToCreateCFString, "%s=%s", logargs.PrefsDomainLogArg, d)
		goto end
	}
	defer C.CFRelease(C.CFTypeRef(cfDomain))

	// Get the array of keys
	cfaPrefs = C.GetPrefNamesForDomain(cfDomain)
	if cfaPrefs == 0 {
		err = errutil.AnnotateErr(ErrFailedToGetPrefNames, "%s=%s", logargs.PrefsDomainLogArg, d)
		goto end
	}
	defer C.CFRelease(C.CFTypeRef(cfaPrefs))

	// Get the numPrefs of keys
	numPrefs = int(C.GetArrayLen(cfaPrefs))
	prefs = make([]*Preference, 0, numPrefs)

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
		prefs = append(prefs, &Preference{
			Domain: string(d),
			Name:   C.GoString(cPref),
		})
		C.free(unsafe.Pointer(cPref))
	}

end:
	return prefs, nil
}
