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
	domains []*PrefsDomain
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
		if domain.Domain == ".GlobalPreferences_m" {
			// This appears to be a duplicate, so we are ignoring it
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
	return fmt.Sprintf("Domains: %d", len(pds.domains))
}

func NewPrefDomains(domains []*PrefsDomain) *PrefDomains {
	return &PrefDomains{
		domains: domains,
	}
}

type RetrievePrefArgs struct {
	IgnoreMissingDomains bool
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
func GetPrefDomains(cfg config.Config, args QueryArgs) (pds *PrefDomains, err error) {
	domains, err := prefdefaults.GetDomains(cfg) // TODO: Limit based on args.Domains, if applicable
	if err != nil {
		goto end
	}
	pds = NewPrefDomains(make([]*PrefsDomain, len(domains)))
	for i, domain := range domains {
		pd := NewPrefsDomain(domain.Domain, PrefsDomainArgs{
			Description: domain.Description,
			Process:     domain.Process,
			Added:       domain.Added,
			Removed:     domain.Removed,
		})
		pd.prefs = make(Prefs, len(domain.Defaults))
		for j, dd := range domain.Defaults {
			pd.prefs[j] = NewPref(PrefArgs{
				Domain:      dd.Domain,
				Name:        dd.Name,
				Description: dd.Description,
				Value:       dd.Value(),
				Default:     dd.Default,
				Labels:      dd.Labels(),
				Kind:        dd.Kind,
				Type:        dd.Type,
			})
		}
		pds.domains[i] = pd
	}
end:
	return pds, err
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

	domains, err = GetPrefDomains(cfg, args)
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
		OmitEmpty: args.OmitEmpty,
	})
	if err != nil {
		err = errors.Join(ErrFailedToQueryGroups, err)
		goto end
	}

	domains.domains = toDomains(domains, filtered)

	valueFilters, err = QueryFiltersForTargets(kvfilters.Values, kvfilters.KeyValues)
	if err != nil {
		goto end
	}
	if !args.IncludeUnchanged {
		valueFilters = append(valueFilters, OmitValueEqualsDefaultFilter())
	}

	filtered, err = kvfilters.Query(kvfilters.QueryArgs{
		Groups:    domains.ToFiltersGroups(),
		Filters:   valueFilters,
		OmitEmpty: args.OmitEmpty,
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
