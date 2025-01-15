package macprefs

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
	"sync"

	"github.com/mikeschinkel/prefsctl/config"
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

var osCode macosutil.Code

func (pds *PrefDomains) DefaultsYAML(opts YAMLOpts) string {
	sb := strings.Builder{}
	for _, domain := range pds.domains {
		sb.WriteString(domain.DefaultsYAML(opts))
	}
	return sb.String()
}

func (pds *PrefDomains) UserManagedPrefDefaults() (pd []*PrefDefault) {
	pd = make([]*PrefDefault, 0)
	pds.Sort()
	for _, domain := range pds.domains {
		if domain.domain == ".GlobalPreferences_m" {
			continue
		}
		for _, pref := range domain.Prefs() {
			if !pref.HasLabel(&UserManaged) {
				continue
			}
			pd = append(pd, pref.PrefDefault)
		}
	}
	return pd
}

func (pds *PrefDomains) Domains() []*PrefsDomain {
	return pds.domains
}

func (pds *PrefDomains) DebugString() string {
	return fmt.Sprintf("Domains: %d, Initialized: %t, PrefsRetrieved: %t",
		len(pds.domains),
		pds.initialized,
		pds.prefsRetrieved,
	)
}

func NewPrefDomains(domains []*PrefsDomain) *PrefDomains {
	return &PrefDomains{
		domains: domains,
	}
}

func (pds *PrefDomains) Initialize(cfg config.Config) (err error) {
	var df PrefDefaultsFunc
	var errs errutil.MultiErr
	var osPrefs OSPrefDefaults

	if pds.initialized {
		goto end
	}

	df, err = GetDefaultsFunc()
	if err != nil {
		errs.Add(err)
		goto end
	}
	osPrefs, err = df(cfg)
	if err != nil {
		errs.Add(err)
		goto end
	}
	for _, dd := range osPrefs.Domains {
		for _, dv := range dd.Defaults {
			SetPrefDefault(dv)
		}
	}
	pds.initialized = true
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
//		p, err := macOSUtils.GetPreference(def.Domain, def.Name)
//		if err == nil {
//			pd.DefaultValue = p.Value
//			pd.Kind = p.Kind
//		}
//		pd.Verified = false
//	}
//	pd.Kind, typeLabel = GetPrefKindAndType(pd.Kind, TypeName(def.Type), pd.DefaultValue)
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

func (pds *PrefDomains) RetrievePrefs(args RetrievePrefArgs) (err error) {
	if !pds.initialized {
		panic("ERROR: Preferences domains must be initialized before calling " +
			"macprefs.PrefDomains.RetrievePrefs().", // TODO Explain how to initialize them
		)
	}
	var errs errutil.MultiErr
	for _, domain := range pds.domains {
		err = domain.RetrievePrefs()
		if err == nil {
			continue
		}
		if args.IgnoreMissingDomains && errors.Is(err, macosutil.ErrFailedToGetPrefDomain) {
			continue
		}
		errs.Add(err) // TODO: Annotate
	}
	pds.prefsRetrieved = true
	return errs.Err()
}

func (pds *PrefDomains) RetrievePrefValues() (err error) {
	if !pds.prefsRetrieved {
		panicf("ERROR: Domain preferences must be retrieved with " +
			"macprefs.PrefDomains.RetrievePref() before calling " +
			"macprefs.PrefDomains.RetrievePrefValues()",
		)
	}
	var errs errutil.MultiErr
	for _, domain := range pds.domains {
		err = domain.RetrievePrefValues()
		if err != nil {
			errs.Add(err) // TODO: Annotate
		}
	}
	return errs.Err()
}

func (pds *PrefDomains) Sort() {
	slices.SortFunc(pds.domains, func(a, b *PrefsDomain) int {
		return strings.Compare(
			strings.ToLower(a.String()),
			strings.ToLower(b.String()),
		)
	})
	for _, domain := range pds.domains {
		domain.SortPrefs()
	}
}

func (pds *PrefDomains) Describe(w io.Writer) {
	for _, domain := range pds.domains {
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
func (pds *PrefDomains) ToFiltersGroups() (groups []kvfilters.Group) {
	groups = make([]kvfilters.Group, len(pds.domains))
	for i, domain := range pds.domains {
		groups[i] = domain
	}
	return groups
}

// GetPrefDomains retrieves the list of macOS preference domains available
// currently on the system via macOS.
func GetPrefDomains(args QueryArgs) (pds *PrefDomains, err error) {
	domains, err := macosutil.GetPreferenceDomains(macosutil.RetrievalArgs{
		Domains: args.ToMacOSUtilPreferenceDomains(),
	})
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
	prefs, err := macosutil.GetPreferences(macosutil.PreferenceDomain(domain))
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

type TemplateDomainsArgs struct {
	UseCurrent bool
}

func (pds *PrefDomains) TemplateDomains(args TemplateDomainsArgs) (domains []*preftemplates.Domain) {
	domains = make([]*preftemplates.Domain, len(pds.domains))
	pds.Sort()
	for i, domain := range pds.domains {
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
				Type:   preftemplates.PreferenceType(pref.typeName),
				Value:  value,
				Labels: pref.Labels(),
			}
		}
		d.Defaults = defaults
		domains[i] = d
	}
	return domains
}

// QueryPrefDomains queries for a set of preference domains based on the QueryArg provided
func QueryPrefDomains(ctx Context, cfg config.Config, args QueryArgs) (domains *PrefDomains, err error) {
	var nameFilters, valueFilters []kvfilters.Filter
	var filtered []kvfilters.Group

	toDomains := func(domains *PrefDomains, group []kvfilters.Group) []*PrefsDomain {
		return sliceconv.Func(group, func(g kvfilters.Group) *PrefsDomain {
			return g.(*PrefsDomain)
		})
	}

	domains, err = GetPrefDomains(args)
	if err != nil {
		goto end
	}
	err = domains.Initialize(cfg)
	if err != nil {
		goto end
	}

	// Retrieve all the prefs for all domains, including all their pref defaults
	err = domains.RetrievePrefs(RetrievePrefArgs{
		IgnoreMissingDomains: true,
	})
	if err != nil {
		goto end
	}
	nameFilters, err = QueryFiltersForTargets(kvfilters.Groups, kvfilters.Keys)
	if err != nil {
		goto end
	}

	filtered, err = kvfilters.Query(kvfilters.QueryArgs{
		Groups:    domains.ToFiltersGroups(),
		Filters:   nameFilters,
		Labels:    []*kvfilters.Label{&UserManaged},
		OmitEmpty: args.OmitEmpty,
	})
	if err != nil {
		err = errors.Join(ErrFailedToQueryGroups, err)
		goto end
	}

	domains.domains = toDomains(domains, filtered)

	// Not retrieve all the current values of the prefs for all domains
	err = domains.RetrievePrefValues()
	if err != nil {
		goto end
	}

	valueFilters, err = QueryFiltersForTargets(kvfilters.Values, kvfilters.KeyValues)
	if err != nil {
		goto end
	}
	if !args.IncludeUnchanged {
		valueFilters = append(valueFilters, OmitValueEqualsDefaultFilter())
	}

	filtered, err = kvfilters.Query(kvfilters.QueryArgs{
		Groups:      domains.ToFiltersGroups(),
		Filters:     valueFilters,
		Labels:      []*kvfilters.Label{&UserManaged},
		OmitEmpty:   args.OmitEmpty,
		OmitInvalid: true,
	})

	if err != nil {
		err = errors.Join(ErrFailedToQueryPrefState, err)
		goto end
	}

	domains.domains = toDomains(domains, filtered)
	domains.Sort()
end:
	return domains, err
}
