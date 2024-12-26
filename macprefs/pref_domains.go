package macprefs

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
	"sync"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
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
		if domain.domain == ".GlobalPreferences_m" {
			continue
		}
		for _, pref := range domain.Prefs() {
			if !pref.HasLabel(&UserManaged) {
				continue
			}
			pds = append(pds, pref.PrefDefault)
		}
	}
	return pds
}

type DefaultsTemplateDomainsArgs struct {
	UseCurrent bool
}

func (dd *PrefDomains) DefaultsTemplateDomains(args DefaultsTemplateDomainsArgs) (domains []*preftemplates.Domain) {
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
			value := pref.DefaultValue
			if args.UseCurrent {
				value = pref.Value()
			}
			defaults[j] = &preftemplates.Default{
				Domain: d,
				Name:   preftemplates.PrefName(pref.Name),
				Type:   pref.TypeName(),
				Value:  value,
				Labels: pref.Labels(),
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

func (dd *PrefDomains) DebugString() string {
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
	var osCode Code
	var dmf DefaultsMapFunc
	var errs errutil.MultiErr

	if dd.initialized {
		goto end
	}

	osCode, err = macosutil.VersionCode()
	if err != nil {
		errs.Add(err)
		goto end
	}
	dmf, err = GetDefaultsMapFunc(osCode)
	if err != nil {
		errs.Add(err)
		goto end
	}
	for domain, dvs := range dmf() {
		for name, dv := range dvs {
			dv.Domain = domain
			dv.Name = name
			SetPrefDefault(dv)
		}
	}
	dd.initialized = true
end:
	return errs.Err()
}

//func SetPrefDefault_NEW(pd *PrefDefault) {
//	var typeLabel *kvfilters.Label
//	dn := DomainName(pd.Domain)
//	pn := PrefName(pd.Name)
//	pd = GetPrefDefault(dn, pn)
//	if pd == nil {
//		pd = NewPrefDefault(dn, pn)
//	}
//	if def.Default != "" {
//		pd.DefaultValue = def.Default
//		pd.Verified = def.Verified
//		//pd.Kind = def.Kind()
//	} else {
//		p, err := macOSUtils.RetrievePreference(def.Domain, def.Name)
//		if err == nil {
//			pd.DefaultValue = p.Value
//			pd.Kind = p.Kind
//		}
//		pd.Verified = false
//	}
//	pd.Kind, typeLabel = GetPrefKindAndTypeLabel(pd.Kind, TypeName(def.Type), pd.DefaultValue)
//
//	//_, _ = fmt.Fprintf(file, "{\n")
//	//_, _ = fmt.Fprintf(file, "\t"+"id:        \"%s/%s\",\n", pd.Domain, pd.Name)
//	//_, _ = fmt.Fprintf(file, "\t"+"name:      \"xxx kind xxx type xxx value\",\n")
//	//_, _ = fmt.Fprintf(file, "\t"+"kind:      reflect.%s,\n", fixCase(pd.Kind.String()))
//	//_, _ = fmt.Fprintf(file, "\t"+"typ:       prefdefaults.%s.Value,\n", fixCase(string(typeLabel.Value)))
//	//_, _ = fmt.Fprintf(file, "\t"+"value:     \"%s\",\n", pd.DefaultValue)
//	//_, _ = fmt.Fprintf(file, "\t"+"wantKind:  reflect.%s,\n", fixCase(pd.Kind.String()))
//	//_, _ = fmt.Fprintf(file, "\t"+"wantLabel: prefdefaults.%s,\n", fixCase(string(typeLabel.Value)))
//	//_, _ = fmt.Fprintf(file, "},\n")
//
//	def.Labels.Add(typeLabel)
//	if !def.Labels.HasNamedLabel(&SetupSets) {
//		def.Labels.Add(&DefaultsSet)
//	}
//	if def.Labels.GetNamedLabel(Class) == nil {
//		def.Labels.Add(&UserManaged)
//	}
//	pd.SetLabels(def.Labels)
//	SetPrefDefault(pd)
//	return pd
//}

func SetPrefDefault(dv *PrefDefault) {
	prefDefaultsLookupMutex.Lock()
	prefDefaultsLookup[dv.Id()] = dv
	prefDefaultsLookupMutex.Unlock()
}

type RetrievePrefArgs struct {
	IgnoreMissingDomains bool
}

func (dd *PrefDomains) RetrievePrefs(args RetrievePrefArgs) (err error) {
	if !dd.initialized {
		panic("ERROR: Preferences domains must be initialized before calling " +
			"macprefs.PrefDomains.RetrievePrefs().",
		)
	}
	var errs errutil.MultiErr
	for _, domain := range dd.domains {
		err = domain.RetrievePrefs()
		if err == nil {
			continue
		}
		if args.IgnoreMissingDomains && errors.Is(err, macosutil.ErrFailedToGetPrefDomains) {
			continue
		}
		errs.Add(err) // TODO: Annotate
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
	domains, err := macosutil.RetrievePreferenceDomains()
	if err != nil {
		goto end
	}
	pds = NewPrefDomains(sliceconv.Func(domains, func(pd macosutil.PreferenceDomain) *PrefsDomain {
		return NewPrefsDomain(DomainName(pd))
	}))
end:
	return pds, err
}

// RetrieveDomainPrefs retrieves the macOS preferences for the specified
// preference domain from macOS.
func RetrieveDomainPrefs(domain DomainName) (pp Prefs, err error) {
	prefs, err := macosutil.RetrievePreferences(macosutil.PreferenceDomain(domain))
	if err != nil {
		goto end
	}
	pp = sliceconv.Func(prefs, func(p *macosutil.Preference) *Pref {
		return NewPref(PrefArgs{
			Domain: DomainName(p.Domain),
			Name:   PrefName(p.Name),
		})
	})
end:
	return pp, err
}
