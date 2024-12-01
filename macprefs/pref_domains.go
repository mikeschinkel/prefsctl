package macprefs

import (
	"io"
	"slices"
	"strings"
	"sync"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/macprefs/filters"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type PrefDefaultsLookup map[PrefId]*PrefDefault

var prefDefaultsLookup = make(PrefDefaultsLookup)
var prefDefaultsLookupMutex sync.Mutex

func GetPrefDefault(id PrefId) *PrefDefault {
	pd, _ := prefDefaultsLookup[id]
	return pd
}

type PrefDomains []*PrefsDomain

func (dd *PrefDomains) Initialize() (err error) {
	var osCode macosutils.Code
	var dmf DefaultsMapFunc
	var errs errutil.MultiErr

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
			id := string(domain) + "/" + string(name)
			prefDefaultsLookup[PrefId(id)] = dv
		}
	}
	for _, domain := range *dd {
		err = domain.Initialize()
		if err != nil {
			errs.Add(err) // TODO: Annotate
		}
	}
end:
	return errs.Err()
}

func (dd *PrefDomains) Sort() {
	slices.SortFunc(*dd, func(a, b *PrefsDomain) int {
		return strings.Compare(
			strings.ToLower(a.String()),
			strings.ToLower(b.String()),
		)
	})
	for _, domain := range *dd {
		domain.SortPrefs()
	}
}

func (dd *PrefDomains) Describe(w io.Writer) {
	for _, domain := range *dd {
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

func (dd *PrefDomains) AsFilterGroups() (groups []filters.Group) {
	groups = make([]filters.Group, len(*dd))
	for i, domain := range *dd {
		groups[i] = domain
	}
	return groups
}

func GetPrefDomains() (pds PrefDomains, err error) {
	domains, err := macosutils.GetPreferenceDomains()
	if err != nil {
		goto end
	}
	pds = sliceconv.Func(domains, func(pd macosutils.PreferenceDomain) *PrefsDomain {
		return NewPrefsDomain(DomainName(pd))
	})
end:
	return pds, err
}

func GetDomainPrefs(domain DomainName) (pp Prefs, err error) {
	prefs, err := macosutils.GetDomainPreferences(macosutils.PreferenceDomain(domain))
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
