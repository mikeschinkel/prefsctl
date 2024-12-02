package macprefs

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logging"
	"github.com/mikeschinkel/prefsctl/macprefs/filters"
)

type DomainName string
type Prefs []*Pref

func (pp *Prefs) Sort() {
	slices.SortFunc(*pp, func(a, b *Pref) int {
		return strings.Compare(
			strings.ToLower(string(a.Key())),
			strings.ToLower(string(b.Key())),
		)
	})
}

var _ filters.Group = (*PrefsDomain)(nil)

// PrefsDomain represents a preference domain in macOS
type PrefsDomain struct {
	domain               DomainName
	prefs                Prefs
	prefsRetrieved       bool
	prefsValuesRetrieved bool
}

func (d *PrefsDomain) ShallowCopy() filters.Group {
	return &PrefsDomain{
		domain:               d.domain,
		prefsRetrieved:       d.prefsRetrieved,
		prefsValuesRetrieved: d.prefsValuesRetrieved,
		prefs:                make(Prefs, 0),
	}
}

func (d *PrefsDomain) SortPrefs() {
	d.prefs.Sort()
}

var unsupportedTypes = make(map[string]struct{})

func (d *PrefsDomain) RetrievePrefs() (err error) {
	d.prefs, err = RetrieveDomainPrefs(d.DomainName())
	for _, pref := range d.prefs {
		pd := GetPrefDefault(pref.Id())
		if pd == nil {
			pd = NewPrefDefault(d.domain, pref.Name)
		}
		pref.PrefDefault = pd
	}
	d.prefsRetrieved = true
	return err
}

func (d *PrefsDomain) RetrievePrefValues() (err error) {
	var errs errutil.MultiErr

	for _, pref := range d.prefs {
		err = pref.Retrieve()
		if err == nil {
			continue
		}
		if errors.Is(err, ErrUnsupportedType) {
			unsupportedType, _ := strings.CutPrefix(pref.value, "unsupported preference class: ")
			unsupportedTypes[unsupportedType] = struct{}{}
			continue
		}
		errs.Add(err) // TODO Annotate
	}
	d.prefsValuesRetrieved = true
	return errs.Err()
}

func NewPrefsDomainFromFiltersGroup(group filters.Group) *PrefsDomain {
	domain := DomainName(group.Name())
	return &PrefsDomain{
		domain:         domain,
		prefs:          NewDomainPrefsFromFiltersKeyValues(domain, group.KeyValues()),
		prefsRetrieved: false,
	}
}

func NewDomainPrefsFromFiltersKeyValues(domain DomainName, kvs []filters.KeyValue) (prefs []*Pref) {
	prefs = make([]*Pref, len(kvs))
	for i, kv := range kvs {
		prefs[i] = NewPref(PrefArgs{
			Domain:  domain,
			Name:    PrefName(kv.Key()),
			Value:   kv.Value(),
			Default: "",  // TODO
			Labels:  nil, // TODO
			Kind:    0,   // TODO
		})
	}
	return prefs
}

func NewPrefsDomain(domain DomainName) *PrefsDomain {
	return &PrefsDomain{
		domain: domain,
		prefs:  make([]*Pref, 0),
	}
}

func (d *PrefsDomain) KeyValues() (kvs []filters.KeyValue) {
	if !d.prefsRetrieved {
		panicf("ERROR: Preferences domain '%s' must be retrieved before calling macprefs.PrefsDomain.KeyValues()",
			d.Name(),
		)
	}
	kvs = make([]filters.KeyValue, len(d.prefs))
	for i, pref := range d.prefs {
		kvs[i] = pref
	}
	return kvs
}

func (d *PrefsDomain) DomainName() DomainName {
	return d.domain
}
func (d *PrefsDomain) Name() filters.Name {
	return filters.Name(d.domain)
}

func (d *PrefsDomain) Code() filters.Code {
	return filters.Codify(string(d.domain))
}

func (d *PrefsDomain) AddKeyValue(value filters.KeyValue) {
	pref, ok := value.(*Pref)
	if !ok {
		panicf("Cannot add type '%T' with value %#v to macprefs.PrefsDomain %#v",
			value,
			value,
			d,
		)
	}
	d.prefs = append(d.prefs, pref)
}

func (d *PrefsDomain) LogArgs() []any {
	return []any{
		logging.PrefsDomainLogArg,
		d.Name(),
	}
}

func (d *PrefsDomain) String() string {
	return string(d.Name())
}
func (d *PrefsDomain) ErrorInfo() error {
	return fmt.Errorf("%s=%s", logging.PrefsDomainLogArg, d)
}

func (d *PrefsDomain) HasPrefix(prefix string) bool {
	return strings.HasPrefix(string(d.Name()), prefix)
}

// Prefs returns all available preferences for this domain
func (d *PrefsDomain) Prefs() []*Pref {
	if !d.prefsRetrieved {
		panicf("ERROR: Preferences domain '%s' must be retrieved before calling macprefs.PrefsDomain.Prefs()",
			d.Name(),
		)
	}
	return d.prefs
}
