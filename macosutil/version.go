package macosutil

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

char* getMacOSVersion() {
    NSProcessInfo *processInfo = [NSProcessInfo processInfo];
    NSOperatingSystemVersion osVersion = [processInfo operatingSystemVersion];

    NSString *versionString = [NSString stringWithFormat:@"%ld.%ld.%ld",
        (long)osVersion.majorVersion,
        (long)osVersion.minorVersion,
        (long)osVersion.patchVersion];

    const char *cString = [versionString UTF8String];
    char *result = strdup(cString);
    return result;
}

 void freeMacOSVersion(char* str) {
     free(str);
 }
*/
import "C"
import (
	"cmp"
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logargs"
)

var versionNumberCodeMap = map[VersionNumber]Code{
	"15":    Sequoia,
	"14":    Sonoma,
	"13":    Ventura,
	"12":    Monterey,
	"11":    BigSur,
	"10.15": Catalina,
	"10.14": Mojave,
	"10.13": HighSierra,
	"10.12": Sierra,
	"10.11": ElCapitan,
	"10.10": Yosemite,
	"10.9":  Mavericks,
	"10.8":  MountainLion,
	"10.7":  Lion,
	"10.6":  SnowLeopard,
	"10.5":  Leopard,
	"10.4":  Tiger,
	"10.3":  Panther,
	"10.2":  Jaguar,
	"10.1":  Puma,
	"10.0":  Cheetah,
}

// MajorNeedsMinor is the version of macOS — "10" — where Apple released new
// named versions incrementing by 0.1, e.g. 10.0, 10.1, 10.2, etc. After this
// version Apple released new named versions incrementing macOS by 1.0, eg. 11,
// 12, 13, etc.
var MajorNeedsMinor = "10"

var versionCodeNameMap = map[Code]Name{
	Sequoia:      "Sequoia",
	Sonoma:       "Sonoma",
	Ventura:      "Ventura",
	Monterey:     "Monterey",
	BigSur:       "Big Sur",
	Catalina:     "Catalina",
	Mojave:       "Mojave",
	HighSierra:   "High Sierra",
	Sierra:       "Sierra",
	ElCapitan:    "El Capitan",
	Yosemite:     "Yosemite",
	Mavericks:    "Mavericks",
	MountainLion: "Mountain Lion",
	Lion:         "Lion",
	SnowLeopard:  "Snow Leopard",
	Leopard:      "Leopard",
	Tiger:        "Tiger",
	Panther:      "Panther",
	Jaguar:       "Jaguar",
	Puma:         "Puma",
	Cheetah:      "Cheetah",
}

var (
	Sequoia      Code = "sequoia"
	Sonoma       Code = "sonoma"
	Ventura      Code = "ventura"
	Monterey     Code = "monterey"
	BigSur       Code = "bigSur"
	Catalina     Code = "catalina"
	Mojave       Code = "mojave"
	HighSierra   Code = "highSierra"
	Sierra       Code = "sierra"
	ElCapitan    Code = "elCapitan"
	Yosemite     Code = "yosemite"
	Mavericks    Code = "mavericks"
	MountainLion Code = "mountainLion"
	Lion         Code = "lion"
	SnowLeopard  Code = "snowLeopard"
	Leopard      Code = "leopard"
	Tiger        Code = "tiger"
	Panther      Code = "panther"
	Jaguar       Code = "jaguar"
	Puma         Code = "puma"
	Cheetah      Code = "cheetah"
)

var semVerRegex = regexp.MustCompile(`(\d+)\.(\d+)\.\d+`)

func (m *macOSUtils) VersionName() (name Name, _ error) {
	var ok bool

	code, err := m.VersionCode()
	if err != nil {
		goto end
	}
	name, ok = versionCodeNameMap[code]
	if !ok {
		panicf("Version code '%s' has no associated MacOS version name. "+
			"Is this a new macOS Version which `prefsctl` does not yet support? "+
			"(Or if you are the developer of `prefsctl` have you not added that "+
			"version into the code yet?)", code)
	}
end:
	return name, err
}

func parseVersion(v string) (int, error) {
	n, err := strconv.ParseInt(v, 10, 0)
	if err != nil {
		err = errors.Join(ErrFailedParsingMajorVersion,
			fmt.Errorf("%v=%v", logargs.Version, v),
		)
	}
	return int(n), err
}

// code is cached by versionCode()
var code Code

func (*macOSUtils) MustGetVersionNumber() VersionNumber {
	if versionNumber == "" {
		panic("Must call .VersionNumber() in macosutil package before calling .MustGetVersionNumber()")
	}
	return versionNumber
}

func (*macOSUtils) MustGetVersionCode() Code {
	if code == "" {
		panic("Must call .VersionCode() in macosutil package before calling .MustGetVersionCode()")
	}
	return code
}

func (m *macOSUtils) VersionCode() (_ Code, err error) {
	var ok bool
	var matches []string
	var n int
	var v VersionNumber

	if code != "" {
		goto end
	}

	v, err = m.GetVersionNumber()
	if err != nil {
		err = errors.Join(ErrFailedToGetMacOSVersionName, err)
		goto end
	}
	matches = semVerRegex.FindStringSubmatch(string(v))
	if matches == nil {
		err = errors.Join(ErrUnrecognizedMacOSVersionFormat,
			fmt.Errorf("%s=%s", logargs.Version, v),
		)
		goto end
	}
	n, err = parseVersion(matches[1])
	if err != nil {
		goto end
	}
	switch {
	case n < 10:
		// We support nothing before OS-X, and mostly nothing before Sierra, really.
		err = errors.Join(ErrUnrecognizedMacOSVersion,
			fmt.Errorf("%s=%s", logargs.Version, v),
		)
	case n > 10:
		// Version of macOS Big Sur thru Sequoia are 11.0 thru 15.0
		// After Sequoia hopefully Apple follows the pattern. If not,
		// this will probably break and need revision.
		v = VersionNumber(strconv.Itoa(n))
	default: // n==10
		// Handle versions 10.0 (Cheetah) through 10.15 (Catalina)
		n, err = parseVersion(matches[2])
		if err != nil {
			goto end
		}
		v = VersionNumber(fmt.Sprintf("10.%d", n))
	}
	code, ok = versionNumberCodeMap[v]
	if !ok {
		err = errors.Join(ErrUnrecognizedMacOSVersion,
			fmt.Errorf("%s=%s", logargs.Version, v),
		)
		goto end
	}
end:
	return code, err
}

var versionNumber VersionNumber

func (m *macOSUtils) GetVersionNumber() (v VersionNumber, err error) {
	var cVer *C.char
	if runtime.GOOS != "darwin" {
		err = errutil.AnnotateErr(ErrNotMacOS, "os="+runtime.GOOS)
		goto end
	}

	cVer = C.getMacOSVersion()
	if cVer == nil {
		err = ErrFailedToGetMacOSVersion
		goto end
	}
	defer C.freeMacOSVersion(cVer)

	v = VersionNumber(C.GoString(cVer))
	versionNumber = v
end:
	return v, err
}

func (m *macOSUtils) ValidVersionNumbers() []VersionNumber {
	nums := make([]VersionNumber, 0, len(versionNumberCodeMap))
	for num := range versionNumberCodeMap {
		nums = append(nums, num)
	}
	slices.SortFunc(nums, func(vn1, vn2 VersionNumber) int {
		var f1, f2 float64
		var err1, err2 error
		f1, err1 = strconv.ParseFloat(string(vn1), 64)
		f2, err2 = strconv.ParseFloat(string(vn2), 64)
		switch {
		case err1 != nil && err2 != nil:
			return 0
		case err1 != nil:
			return 1
		case err2 != nil:
			return -1
		}
		return cmp.Compare(f1, f2)
	})
	return nums
}

// ValidateVersionNumber accepts a version number string and compares to keys in
// the map versionNumberCodeMap to see if they are a valid version.
func (m *macOSUtils) ValidateVersionNumber(num VersionNumber) (valid bool) {
	parts := strings.SplitN(string(num), ".", 2)
	ver := parts[0]
	if parts[0] == MajorNeedsMinor {
		ver = strings.Join(parts, ".")
	}
	_, valid = versionNumberCodeMap[VersionNumber(ver)]
	return valid
}

func (m *macOSUtils) ValidateVersionName(name Name) (valid bool) {
	for _, vn := range versionCodeNameMap {
		if vn == name {
			valid = true
			goto end
		}
	}
end:
	return valid
}

func (m *macOSUtils) GetVersionName(code Code) (name Name) {
	name, _ = versionCodeNameMap[code]
	return name
}

func (m *macOSUtils) GetVersionCode(num VersionNumber) (code Code) {
	code, _ = versionNumberCodeMap[num]
	return code
}
