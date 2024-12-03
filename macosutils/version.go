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
	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
)

var versionNumberCodeMap = map[NumericVersion]Code{
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

func VersionName() (name Name, _ error) {
	var ok bool

	code, err := VersionCode()
	if err != nil {
		goto end
	}
	name, ok = versionCodeNameMap[code]
	if !ok {
		panicf("Version code '%s' has no associated macOS version name. Did the programmer forget to add that into the code?")
	}
end:
	return name, err
}

func VersionCode() (code Code, _ error) {
	var ok bool
	var matches []string
	var n int
	var versionCode Code

	parseVersion := func(v string) (int, error) {
		n, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			err = errors.Join(ErrFailedParsingMajorVersion,
				fmt.Errorf("%v=%v", logargs.VersionLogArg, v),
			)
		}
		return int(n), err
	}

	v, err := Version()
	if err != nil {
		err = errors.Join(ErrFailedToGetMacOSVersionName, err)
		goto end
	}
	matches = semVerRegex.FindStringSubmatch(v)
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
		v = strconv.Itoa(n)
	default: // n==10
		// Handle versions 10.0 (Cheetah) through 10.15 (Catalina)
		n, err = parseVersion(matches[2])
		if err != nil {
			goto end
		}
		v = fmt.Sprintf("10.%d", n)
	}
	versionCode, ok = versionNumberCodeMap[NumericVersion(v)]
	if !ok {
		err = errors.Join(ErrUnrecognizedMacOSVersion,
			fmt.Errorf("%s=%s", logargs.VersionLogArg, v),
		)
		goto end
	}
	code = versionCode
end:
	return code, err
}

func Version() (v string, err error) {
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

	v = C.GoString(cVer)
end:
	return v, err
}
