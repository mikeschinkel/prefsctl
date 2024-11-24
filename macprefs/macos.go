package macprefs

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
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logging"
)

type MacOS struct {
	Name         MacOSName
	PrefDefaults []*PrefDefault
}
type MacOSName string

var macOSIndex = map[string]MacOSName{
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

const (
	Sequoia      MacOSName = "sequoia"
	Sonoma       MacOSName = "sonoma"
	Ventura      MacOSName = "ventura"
	Monterey     MacOSName = "monterey"
	BigSur       MacOSName = "bigSur"
	Catalina     MacOSName = "catalina"
	Mojave       MacOSName = "mojave"
	HighSierra   MacOSName = "highSierra"
	Sierra       MacOSName = "sierra"
	ElCapitan    MacOSName = "elCapitan"
	Yosemite     MacOSName = "yosemite"
	Mavericks    MacOSName = "mavericks"
	MountainLion MacOSName = "mountainLion"
	Lion         MacOSName = "lion"
	SnowLeopard  MacOSName = "snowLeopard"
	Leopard      MacOSName = "leopard"
	Tiger        MacOSName = "tiger"
	Panther      MacOSName = "panther"
	Jaguar       MacOSName = "jaguar"
	Puma         MacOSName = "puma"
	Cheetah      MacOSName = "cheetah"
)

var macOSVerRegex = regexp.MustCompile(`(\d+)\.(\d+)\.\d+`)

func MacOSVersionName() (name MacOSName, _ error) {
	var ok bool
	var matches []string
	var n int

	parseVersion := func(v string) (int, error) {
		n, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			err = errors.Join(ErrFailedParsingMajorVersion,
				fmt.Errorf("%v=%v", logging.VersionLogArg, v),
			)
		}
		return int(n), err
	}

	v, err := MacOSVersion()
	if err != nil {
		err = errors.Join(ErrFailedToGetMacOSVersionName, err)
		goto end
	}
	matches = macOSVerRegex.FindStringSubmatch(v)
	if matches == nil {
		err = errors.Join(ErrUnrecognizedMacOSVersionFormat,
			fmt.Errorf("%s=%s", logging.VersionLogArg, v),
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
			fmt.Errorf("%s=%s", logging.VersionLogArg, v),
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
	name, ok = macOSIndex[v]
	if !ok {
		err = errors.Join(ErrUnrecognizedMacOSVersion,
			fmt.Errorf("%s=%s", logging.VersionLogArg, v),
		)
		goto end
	}
end:
	return name, err
}

func MacOSVersion() (v string, err error) {
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

type data struct {
	Name    MacOSName
	Domains []Domain
}

const goMacOSPrefsFuncTemplate = `func {{.Name}}PrefDefaults() macprefs.DomainPrefDefaults {
	return macprefs.DomainPrefDefaults{
		{{- range $i, $domain := .Domains}}
		"{{$domain}}": []*macprefs.PrefDefault{
			{{- range $pref := index $.PrefDefaults $i}}
			{
				Name:     "{{$pref.Name}}",
				Type:     {{$pref.TypeName}},
				WhatSets: {{$pref.WhatSets}},
				{{- if $pref.ShowValue}}
				Value:    "{{$pref.Value}}",
				{{- end}}
			},
			{{- end}}
		},
		{{- end}}
	}
}`

// tmplPrefs represents the structure passed to the template
type tmplPrefs struct {
	Name         string
	Domains      []Domain
	PrefDefaults [][]tmplPref
}

//func newTmplPrefs(name string) *tmplPrefs {
//	seen := make(map[Domain]int)
//	var data tmplPrefs
//	data.Name = string(name)
//
//	// First pass: collect domains in order of appearance
//	for _, pref := range m.PrefDefaults {
//		if _, exists := seen[pref.Domain]; !exists {
//			seen[pref.Domain] = len(data.Domains)
//			data.Domains = append(data.Domains, pref.Domain)
//			data.PrefDefaults = append(data.PrefDefaults, make([]tmplPref, 0))
//		}
//	}
//
//	// Second pass: populate preferences
//	for _, pref := range m.PrefDefaults {
//		idx := seen[pref.Domain]
//
//		tPref := tmplPref{
//			Name:     pref.Name,
//			Value:    pref.Value,
//			TypeName: prefTypeString(pref.Type),
//			WhatSets: whatSetsString(pref.WhatSets),
//		}
//
//		data.PrefDefaults[idx] = append(data.PrefDefaults[idx], tPref)
//	}
//
//	return &tmplPrefs{Name: name}
//}

// tmplPref represents a preference formatted for template output
type tmplPref PrefDefault

// ShowValue returns true if WhatSets!="install" — This means it is either a true
// default or we do not yet know for sure, but we do not that it is not determine
// exclusively by the person installing the OS.
func (p tmplPref) ShowValue() bool {
	return p.WhatSets != InstallSets
}

// GoType converts a PrefType to its template string representation as a
// Go constant, e.g. "string" ->  "macprefs.StringType"
func (p tmplPref) GoType() string {
	return fmt.Sprintf("macprefs.%s%sType",
		strings.ToUpper(string(p.Type[:1])),
		string(p.Type[1:]),
	)
}

// GoWhatSets converts a WhatSetsType to its template string representation as a
// Go constant, e.g. "macos" -> "macprefs.MacOSSets" and "unsure" -> "macprefs.UnsureWhatSets"
func (p tmplPref) GoWhatSets() (what string) {
	switch p.WhatSets {
	case "macos":
		what = "MacOS"
	case "install":
		what = "Install"
	case "unsure":
		what = "UnsureWhat"
	default:
		what = "UnknownWhat"
	}
	return fmt.Sprintf("macprefs.%sSets", what)
}

// GoCode generates Go code for preference defaults
func (m *MacOS) GoCode() (string, error) {
	var tmpl = goMacOSPrefsFuncTemplate
	// Group preferences by domain while preserving order
	seen := make(map[Domain]int)
	var data tmplPrefs
	data.Name = string(m.Name)

	// First pass: collect domains in order of appearance
	for _, pref := range m.PrefDefaults {
		if _, exists := seen[pref.Domain]; !exists {
			seen[pref.Domain] = len(data.Domains)
			data.Domains = append(data.Domains, pref.Domain)
			data.PrefDefaults = append(data.PrefDefaults, make([]tmplPref, 0))
		}
	}

	// Second pass: populate preferences
	for _, pref := range m.PrefDefaults {
		idx := seen[pref.Domain]

		tPref := tmplPref{
			Name:     pref.Name,
			Value:    pref.Value,
			Type:     pref.Type,
			WhatSets: pref.WhatSets,
		}

		data.PrefDefaults[idx] = append(data.PrefDefaults[idx], tPref)
	}

	// Parse and execute template
	t, err := template.New("prefDefaults").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
