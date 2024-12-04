package macprefs

import (
	"fmt"
	"io"
	"slices"
	"strings"
	"sync"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/macprefs/kvfilters"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type PrefDefaultsLookup map[PrefId]*PrefDefault

var prefDefaultsLookup = make(PrefDefaultsLookup)
var prefDefaultsLookupMutex sync.Mutex

func LookupPrefDefault(id PrefId) *PrefDefault {
	pd, _ := prefDefaultsLookup[id]
	return pd
}

type PrefDomains struct {
	domains        []*PrefsDomain
	initialized    bool
	prefsRetrieved bool
}

func (dd *PrefDomains) UserManagedPrefDefaults() (pds []*PrefDefault) {
	pds = make([]*PrefDefault, 0)
	dd.Sort()
	for _, domain := range dd.domains {
		for _, pref := range domain.Prefs() {
			if !pref.Labels().Contains(pref.Labels()...) {
				continue
			}
			pds = append(pds, pref.PrefDefault)
		}
	}
	return pds
}

//	func (dd *PrefDomains) TemplatePrefs() (prefs []*preftemplates.Pref) {
//		prefs = make([]*preftemplates.Pref, 0)
//		dd.Sort()
//		for _, domain := range dd.domains {
//			for _, pref := range domain.Prefs() {
//				prefs = append(prefs, &preftemplates.Pref{
//					Domain:    preftemplates.DomainName(pref.Domain),
//					Name:      preftemplates.PrefName(pref.Name),
//					Default:   pref.DefaultValue,
//					Labels:    pref.Labels(),
//					NoDefault: pref.NoDefault,
//				})
//			}
//		}
//		return prefs
//	}
func (dd *PrefDomains) TemplateDomains() (domains []*preftemplates.Domain) {
	domains = make([]*preftemplates.Domain, len(dd.domains))
	dd.Sort()
	for i, domain := range dd.domains {
		d := &preftemplates.Domain{
			Name:     preftemplates.DomainName(domain.DomainName()),
			Defaults: nil,
			MacOS:    nil,
		}
		defaults := make([]*preftemplates.Default, len(domain.Prefs()))
		for j, pref := range domain.Prefs() {
			defaults[j] = &preftemplates.Default{
				Domain:    d,
				Name:      preftemplates.PrefName(pref.Name),
				Value:     pref.DefaultValue,
				Labels:    pref.Labels(),
				NoDefault: pref.NoDefault,
			}
		}
		d.Defaults = defaults
		domains[i] = d
	}
	return domains
}

func (dd *PrefDomains) Domains() []*PrefsDomain {
	return dd.domains
}

func (dd *PrefDomains) String() string {
	return fmt.Sprintf("Domains: %d, Initialized: %t, PrefsRetrieved: %t",
		len(dd.domains),
		dd.initialized,
		dd.prefsRetrieved,
	)
}

func NewPrefDomains(domains []*PrefsDomain) *PrefDomains {
	return &PrefDomains{
		domains: domains,
	}
}

func (dd *PrefDomains) Initialize() (err error) {
	var osCode macosutils.Code
	var dmf DefaultsMapFunc
	var errs errutil.MultiErr

	if dd.initialized {
		goto end
	}

	osCode, err = macosutils.VersionCode()
	if err != nil {
		errs.Add(err)
		goto end
	}
	dmf, err = GetDefaultsMapFunc(osCode)
	if err != nil {
		errs.Add(err)
		goto end
	}
	prefDefaultsLookupMutex.Lock()
	defer prefDefaultsLookupMutex.Unlock()

	for domain, dvs := range dmf() {
		for name, dv := range dvs {
			dv.Domain = domain
			dv.Name = name
			prefDefaultsLookup[NewPrefId(domain, name)] = dv
		}
	}
	dd.initialized = true
end:
	return errs.Err()
}

func (dd *PrefDomains) RetrievePrefs() (err error) {
	if !dd.initialized {
		panic("ERROR: Preferences domains must be initialized before calling " +
			"macprefs.PrefDomains.RetrievePrefs().",
		)
	}
	var errs errutil.MultiErr
	for _, domain := range dd.domains {
		err = domain.RetrievePrefs()
		if err != nil {
			errs.Add(err) // TODO: Annotate
		}
	}
	dd.prefsRetrieved = true
	return errs.Err()
}

func (dd *PrefDomains) RetrievePrefValues() (err error) {
	if !dd.prefsRetrieved {
		panicf("ERROR: Domain preferences must be retrieved with " +
			"macprefs.PrefDomains.RetrievePref() before calling " +
			"macprefs.PrefDomains.RetrievePrefValues()",
		)
	}
	var errs errutil.MultiErr
	for _, domain := range dd.domains {
		err = domain.RetrievePrefValues()
		if err != nil {
			errs.Add(err) // TODO: Annotate
		}
	}
	return errs.Err()
}

func (dd *PrefDomains) Sort() {
	slices.SortFunc(dd.domains, func(a, b *PrefsDomain) int {
		return strings.Compare(
			strings.ToLower(a.String()),
			strings.ToLower(b.String()),
		)
	})
	for _, domain := range dd.domains {
		domain.SortPrefs()
	}
}

func (dd *PrefDomains) Describe(w io.Writer) {
	for _, domain := range dd.domains {
		writeString(w, string(domain.Name()))
		writeByte(w, '\n')
		for _, pref := range domain.KeyValues() {
			writeString(w, "— ")
			writeString(w, string(pref.Key()))
			if pref.Value() == "" {
				writeByte(w, ':')
			} else {
				writeString(w, ": ")
				writeString(w, pref.Value())
			}
			writeByte(w, '\n')
		}
	}
}
func (dd *PrefDomains) ToFiltersGroups() (groups []kvfilters.Group) {
	groups = make([]kvfilters.Group, len(dd.domains))
	for i, domain := range dd.domains {
		groups[i] = domain
	}
	return groups
}

// RetrievePrefDomains retrieves the list of macOS preference domains available
// currently on the system via macOS.
func RetrievePrefDomains() (pds *PrefDomains, err error) {
	domains, err := macosutils.RetrievePreferenceDomains()
	if err != nil {
		goto end
	}
	pds = NewPrefDomains(sliceconv.Func(domains, func(pd macosutils.PreferenceDomain) *PrefsDomain {
		return NewPrefsDomain(DomainName(pd))
	}))
end:
	return pds, err
}

// RetrieveDomainPrefs retrieves the macOS preferences for the specified
// preference domain from macOS.
func RetrieveDomainPrefs(domain DomainName) (pp Prefs, err error) {
	prefs, err := macosutils.Retrieve(macosutils.PreferenceDomain(domain))
	if err != nil {
		goto end
	}
	pp = sliceconv.Func(prefs, func(p *macosutils.Preference) *Pref {
		return NewPref(PrefArgs{
			Domain: DomainName(p.Domain),
			Name:   PrefName(p.Name),
		})
	})
end:
	return pp, err
}
