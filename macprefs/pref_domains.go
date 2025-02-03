package macprefs

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
	"github.com/mikeschinkel/prefsctl/prefdefaults"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

type (
	YAMLDocument = yamlutil.Document
)

const (
	DefaultsKind = prefsyaml.KindName(macpreflabels.DefaultsKind)
)

type PrefDomains struct {
	domains        []*PrefsDomain
	initialized    bool
	prefsRetrieved bool
}

var osCode macosutil.Code

func (pds *PrefDomains) GetDefaultsYAMLDocument(opts YAMLOpts) (ymf yamlutil.MultidocFile, _ error) {
	var errs errutil.MultiErr
	ymf = yamlutil.NewMultidocFile()
	for _, domain := range pds.domains {
		yd, err := domain.GetYAMLDocument(DefaultsKind, opts)
		if err != nil {
			goto end
		}
		ymf.AddDocument(yd)
	}
end:
	return ymf, errs.Err()
}

func (pds *PrefDomains) UserManagedPrefDefaults() (pd []*prefdefaults.PrefDefault) {
	pd = make([]*prefdefaults.PrefDefault, 0)
	pds.Sort()
	for _, domain := range pds.domains {
		if domain.domain == ".GlobalPreferences_m" {
			continue
		}
		for _, pref := range domain.Prefs() {
			if !pref.HasLabel(&macpreflabels.UserManaged) {
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
	var df prefdefaults.Func
	var errs errutil.MultiErr
	var osPrefs prefdefaults.OSPrefDefaults

	if pds.initialized {
		goto end
	}

	df = prefdefaults.GetDefaultsFunc()
	osPrefs, err = df(cfg)
	if err != nil {
		errs.Add(err)
		goto end
	}
	for _, dd := range osPrefs.Domains {
		for _, dv := range dd.Defaults {
			prefdefaults.SetPrefDefault(dv)
		}
	}
	pds.initialized = true
end:
	return errs.Err()
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
		return NewPref(PrefOpts{
			Domain: DomainName(p.Domain),
			Name:   PrefName(p.Name),
		})
	})
end:
	return pp, err
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
		Labels:    []*kvfilters.Label{&macpreflabels.UserManaged},
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
		Labels:      []*kvfilters.Label{&macpreflabels.UserManaged},
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
