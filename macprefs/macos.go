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
	"slices"
	"strconv"
	"text/template"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logging"
)

type MacOS struct {
	Name         Name
	PrefDefaults []*PrefDefault
}
type Name string
type NumericVersion string

var macOSIndex = map[NumericVersion]Label{
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

var (
	Sequoia      = Label{OSVersion, "sequoia"}
	Sonoma       = Label{OSVersion, "sonoma"}
	Ventura      = Label{OSVersion, "ventura"}
	Monterey     = Label{OSVersion, "monterey"}
	BigSur       = Label{OSVersion, "bigSur"}
	Catalina     = Label{OSVersion, "catalina"}
	Mojave       = Label{OSVersion, "mojave"}
	HighSierra   = Label{OSVersion, "highSierra"}
	Sierra       = Label{OSVersion, "sierra"}
	ElCapitan    = Label{OSVersion, "elCapitan"}
	Yosemite     = Label{OSVersion, "yosemite"}
	Mavericks    = Label{OSVersion, "mavericks"}
	MountainLion = Label{OSVersion, "mountainLion"}
	Lion         = Label{OSVersion, "lion"}
	SnowLeopard  = Label{OSVersion, "snowLeopard"}
	Leopard      = Label{OSVersion, "leopard"}
	Tiger        = Label{OSVersion, "tiger"}
	Panther      = Label{OSVersion, "panther"}
	Jaguar       = Label{OSVersion, "jaguar"}
	Puma         = Label{OSVersion, "puma"}
	Cheetah      = Label{OSVersion, "cheetah"}
)

var macOSVerRegex = regexp.MustCompile(`(\d+)\.(\d+)\.\d+`)

func MacOSVersionName() (name Name, _ error) {
	var ok bool
	var matches []string
	var n int
	var label Label

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
	label, ok = macOSIndex[NumericVersion(v)]
	if !ok {
		err = errors.Join(ErrUnrecognizedMacOSVersion,
			fmt.Errorf("%s=%s", logging.VersionLogArg, v),
		)
		goto end
	}
	name = Name(label.Name)
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

// tmplPrefs represents the structure passed to the template
type tmplPrefs struct {
	Name         string
	Domains      []Domain
	PrefDefaults [][]tmplPref
}

// tmplPref represents a preference formatted for template output
type tmplPref PrefDefault

// ShowValue returns true if there is no label "install-sets" — This means it is
// either a true default or we do not yet know for sure, but we do not that it is
// not determine exclusively by the person installing the OS.
func (p tmplPref) ShowValue() bool {
	return !slices.ContainsFunc(p.Labels, func(label *Label) bool {
		return label.Value == InstallSets.Value
	})
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
			Name:   pref.Name,
			Value:  pref.Value,
			Labels: pref.Labels,
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

const goMacOSPrefsFuncTemplate = `func {{.Name}}PrefDefaults() macprefs.DomainPrefDefaults {
	return macprefs.DomainPrefDefaults{
		{{- range $i, $domain := .Domains}}
		"{{$domain}}": []*macprefs.PrefDefault{
			{{- range $pref := index $.PrefDefaults $i}}
			{
				Name:     "{{$pref.Name}}",
				Type:     {{$pref.TypeName}},
				{{- if $pref.ShowValue}}
				Value:    "{{$pref.Value}}",
				{{- end}}
				Labels: []macprefs.Label{
				{{- range $i, $label := .Labels}}
					{{$label.GoProperty}},
				{{- end}}
				},
			},
			{{- end}}
		},
		{{- end}}
	}
}`
