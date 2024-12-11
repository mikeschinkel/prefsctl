package macprefs

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macprefs/kvfilters"
	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
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

var _ kvfilters.Group = (*PrefsDomain)(nil)

// PrefsDomain represents a preference domain in macOS
type PrefsDomain struct {
	domain               DomainName
	prefs                Prefs
	prefsRetrieved       bool
	prefsValuesRetrieved bool
}

//func (d *PrefsDomain) TemplateDefaults() (defaults []*preftemplates.Default) {
//	defaults = make([]*preftemplates.Default, 0)
//	for _, pref := range d.Prefs() {
//		defaults = append(defaults, &preftemplates.Default{
//			Domain:    d,
//			Name:      preftemplates.PrefName(pref.Name),
//			Value:     pref.DefaultValue,
//			Labels:    pref.Labels(),
//			Verified:  pref.Verified,
//		})
//	}
//	return defaults
//}

func (d *PrefsDomain) ShallowCopy() kvfilters.Group {
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
		pd := LookupPrefDefault(pref.Id())
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

func NewPrefsDomain(domain DomainName) *PrefsDomain {
	return &PrefsDomain{
		domain: domain,
		prefs:  make([]*Pref, 0),
	}
}

func (d *PrefsDomain) KeyValues() (kvs []kvfilters.KeyValue) {
	if !d.prefsRetrieved {
		panicf("ERROR: Preferences domain '%s' must be retrieved before calling macprefs.PrefsDomain.KeyValues()",
			d.Name(),
		)
	}
	kvs = make([]kvfilters.KeyValue, len(d.prefs))
	for i, pref := range d.prefs {
		kvs[i] = pref
	}
	return kvs
}

func (d *PrefsDomain) DomainName() DomainName {
	return d.domain
}
func (d *PrefsDomain) Name() kvfilters.Name {
	return kvfilters.Name(d.domain)
}

func (d *PrefsDomain) Code() kvfilters.Code {
	return kvfilters.Codify(string(d.domain))
}

func (d *PrefsDomain) AddKeyValue(value kvfilters.KeyValue) {
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
		logargs.PrefsDomainLogArg,
		d.Name(),
	}
}

func (d *PrefsDomain) String() string {
	return string(d.Name())
}
func (d *PrefsDomain) ErrorInfo() error {
	return fmt.Errorf("%s=%s", logargs.PrefsDomainLogArg, d)
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
