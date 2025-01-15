package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreFoundation
#import <CoreFoundation/CoreFoundation.h>

// CFPreferencesCopyApplicationList is deprecated but provides cleanest API for our use-case.
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

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

    // Release the original array
    CFRelease(appDomains);

    return allDomains;
}

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
	"unsafe"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logargs"
)

type PreferenceDomain string
type DomainName string

type RetrievalArgs struct {
	Domains     []PreferenceDomain
	domainsPtr  unsafe.Pointer
	domainIndex map[PreferenceDomain]NULL
}

func (args *RetrievalArgs) HasDomain(domain PreferenceDomain) (has bool) {
	if args.domainsPtr != unsafe.Pointer(&args.Domains) {
		args.domainIndex = make(map[PreferenceDomain]NULL, len(args.Domains))
		for _, d := range args.Domains {
			args.domainIndex[d] = NULL{}
		}
		args.domainsPtr = unsafe.Pointer(&args.Domains)
	}
	_, has = args.domainIndex[domain]
	return has
}

// GetPreferenceDomains returns a list of all preference domains available
// for the current user on macOS.
func (*macOSUtils) GetPreferenceDomains(args RetrievalArgs) (domains []PreferenceDomain, err error) {
	var count, maxCount int

	// Get the CFArray of preference domains
	domainsArray := C.GetPreferenceDomains()
	if domainsArray == 0 {
		err = ErrFailedToGetPrefDomains
		goto end
	}
	defer C.CFRelease(C.CFTypeRef(domainsArray))

	maxCount = int(C.GetArrayLen(domainsArray))
	if len(args.Domains) == 0 {
		// Get the count of available domains
		count = maxCount
	} else {
		// Get the count of requested domains
		count = len(args.Domains)
	}
	domains = make([]PreferenceDomain, 0, count)

	// Iterate through the array and convert each CFString to Go string
	for i := 0; i < maxCount; i++ {
		cDomainPtr := C.GetArrayValueAtIndex(domainsArray, C.int(i))
		if cDomainPtr == 0 {
			continue
		}

		cDomain := C.CFStringToCString(cDomainPtr)
		if cDomain == nil {
			continue
		}
		domain := PreferenceDomain(C.GoString(cDomain))
		if !args.HasDomain(domain) {
			continue
		}

		domains = append(domains, PreferenceDomain(C.GoString(cDomain)))

		C.free(unsafe.Pointer(cDomain))

		if len(domains) >= count {
			goto end
		}
	}
end:
	return domains, err
}

// GetPreferences returns all available preferences for this domain
func (*macOSUtils) GetPreferences(d PreferenceDomain) (prefs []*Preference, err error) {
	var numPrefs int

	// Convert domain name to CFString
	cfDomain := C.CreateCFString(C.CString(string(d)))
	// TODO Set type here instead of calling func
	var cfaPrefs = C.GetPrefNamesForDomain(cfDomain)
	if cfDomain == 0 {
		err = errutil.AnnotateErr(ErrFailedToCreateCFString, "%s=%s", logargs.PrefsDomain, d)
		goto end
	}
	defer C.CFRelease(C.CFTypeRef(cfDomain))

	// Get the array of keys
	cfaPrefs = C.GetPrefNamesForDomain(cfDomain)
	if cfaPrefs == 0 {
		err = errutil.AnnotateErr(ErrFailedToGetPrefDomain, "%s=%s", logargs.PrefsDomain, d)
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
	return prefs, err
}
