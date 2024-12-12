package macosutils

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
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"strconv"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logargs"
)

var versionNumberCodeMap = map[VersionNumber]VersionCode{
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

var versionCodeNameMap = map[VersionCode]VersionName{
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
	Sequoia      VersionCode = "sequoia"
	Sonoma       VersionCode = "sonoma"
	Ventura      VersionCode = "ventura"
	Monterey     VersionCode = "monterey"
	BigSur       VersionCode = "bigSur"
	Catalina     VersionCode = "catalina"
	Mojave       VersionCode = "mojave"
	HighSierra   VersionCode = "highSierra"
	Sierra       VersionCode = "sierra"
	ElCapitan    VersionCode = "elCapitan"
	Yosemite     VersionCode = "yosemite"
	Mavericks    VersionCode = "mavericks"
	MountainLion VersionCode = "mountainLion"
	Lion         VersionCode = "lion"
	SnowLeopard  VersionCode = "snowLeopard"
	Leopard      VersionCode = "leopard"
	Tiger        VersionCode = "tiger"
	Panther      VersionCode = "panther"
	Jaguar       VersionCode = "jaguar"
	Puma         VersionCode = "puma"
	Cheetah      VersionCode = "cheetah"
)

var semVerRegex = regexp.MustCompile(`(\d+)\.(\d+)\.\d+`)

func versionName() (name VersionName, _ error) {
	var ok bool

	code, err := versionCode()
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
			fmt.Errorf("%v=%v", logargs.VersionLogArg, v),
		)
	}
	return int(n), err
}

// code is cached by versionCode()
var code VersionCode

func mustGetVersionCode() VersionCode {
	if code == "" {
		panic("Must call VersionCode() in macosutils package before calling MustGetVersionCode()")
	}
	return code
}

func versionCode() (_ VersionCode, err error) {
	var ok bool
	var matches []string
	var n int
	var v VersionNumber

	if code != "" {
		goto end
	}

	v, err = getVersionNumber()
	if err != nil {
		err = errors.Join(ErrFailedToGetMacOSVersionName, err)
		goto end
	}
	matches = semVerRegex.FindStringSubmatch(string(v))
	if matches == nil {
		err = errors.Join(ErrUnrecognizedMacOSVersionFormat,
			fmt.Errorf("%s=%s", logargs.VersionLogArg, v),
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
			fmt.Errorf("%s=%s", logargs.VersionLogArg, v),
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
			fmt.Errorf("%s=%s", logargs.VersionLogArg, v),
		)
		goto end
	}
end:
	return code, err
}

func getVersionNumber() (v VersionNumber, err error) {
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
end:
	return v, err
}
