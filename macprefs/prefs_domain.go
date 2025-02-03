package macprefs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/appinfo"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
	"github.com/mikeschinkel/prefsctl/prefdefaults"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

var _ kvfilters.Group = (*PrefsDomain)(nil)

// PrefsDomain represents a preference domain in macOS
type PrefsDomain struct {
	domain               DomainName
	invalid              bool
	prefs                Prefs
	prefsRetrieved       bool
	prefsValuesRetrieved bool
}
type YAMLOpts struct {
	UseValueForDefault bool
	APIVersion         OSVersion
}

func (pd *PrefsDomain) GetYAMLResource(kind prefsyaml.KindName, opts YAMLOpts) (yr *prefsyaml.Resource) {
	var labelValues []*prefsyaml.LabelValue
	var typeName string
	yr = &prefsyaml.Resource{
		APIVersion: appinfo.LatestAPIVersion,
		KindName:   kind,
		Spec:       prefsyaml.NewSpec(),
		MetaData: prefsyaml.Metadata{
			Domain: prefsyaml.DomainName(pd.domain),
		},
	}
	for _, pref := range pd.Prefs() {
		var val, def string
		if opts.UseValueForDefault {
			def = pref.Value()
		} else {
			val = pref.Value()
			def = pref.Default()
		}
		labels := pref.Labels()
		// Delete "type" from in YAML because as there is an explicit "type" field
		labels.DeleteNamedLabel(macpreflabels.Type)
		if labels.HasLabels() {
			labelValues = sliceconv.Func(labels.Labels(), func(l *kvfilters.Label) *prefsyaml.LabelValue {
				lv := prefsyaml.LabelValue(l.Value)
				return &lv
			})
		}
		typeName, _ = strings.CutSuffix(string(pref.TypeName()), "Type")
		yr.Spec.Prefs = append(yr.Spec.Prefs, prefsyaml.Pref{
			MetaData: &yr.MetaData,
			Name:     prefsyaml.PrefName(pref.Name),
			Type:     prefsyaml.PrefType(typeName),
			Value:    yamlutil.NewValue(val),
			Default:  yamlutil.NewValue(def),
			Labels:   labelValues,
		})
	}
	return yr
}

func (pd *PrefsDomain) GetYAMLDocument(kind prefsyaml.KindName, opts YAMLOpts) (yd yamlutil.Document, err error) {
	yr := pd.GetYAMLResource(kind, opts)
	yd, err = yamlutil.BuildDocument(yr)
	if err != nil {
		goto end
	}
end:
	return yd, err
}

func (pd *PrefsDomain) Valid() bool {
	// TODO: Change to determining invalid based on labels
	//       Need to understand what determines something is invalid before doing so, though
	return !pd.invalid
}

func (pd *PrefsDomain) ShallowCopy() kvfilters.Group {
	return &PrefsDomain{
		domain:               pd.domain,
		invalid:              pd.invalid,
		prefsRetrieved:       pd.prefsRetrieved,
		prefsValuesRetrieved: pd.prefsValuesRetrieved,
		prefs:                make(Prefs, 0),
	}
}

func (pd *PrefsDomain) SortPrefs() {
	pd.prefs.Sort()
}

var unsupportedTypes = make(map[string]struct{})

func (pd *PrefsDomain) RetrievePrefs() (err error) {
	pd.prefsRetrieved = true
	pd.prefs, err = RetrieveDomainPrefs(pd.DomainName())
	if err != nil {
		pd.invalid = true
		pd.prefs = make(Prefs, 0)
		goto end
	}
	for _, pref := range pd.prefs {
		pd := prefdefaults.LookupPrefDefault(pref.Id())
		if pd == nil {
			pd = prefdefaults.NewPrefDefault(pref.Domain, pref.Name, &prefdefaults.PrefDefaultOpts{
				Kind:    0, // TODO: Populate these opts
				Added:   "",
				Removed: "",
			})
		}
		pref.PrefDefault = pd
	}
end:
	return err
}

func (pd *PrefsDomain) RetrievePrefValues() (err error) {
	var errs errutil.MultiErr

	for _, pref := range pd.prefs {
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
	pd.prefsValuesRetrieved = true
	return errs.Err()
}

func NewPrefsDomain(domain DomainName) *PrefsDomain {
	return &PrefsDomain{
		domain: domain,
		prefs:  make([]*Pref, 0),
	}
}

func (pd *PrefsDomain) KeyValues() (kvs []kvfilters.KeyValue) {
	if !pd.prefsRetrieved {
		panicf("ERROR: Preferences domain '%s' must be retrieved before calling macprefs.PrefsDomain.KeyValues()",
			pd.Name(),
		)
	}
	kvs = make([]kvfilters.KeyValue, len(pd.prefs))
	for i, pref := range pd.prefs {
		kvs[i] = pref
	}
	return kvs
}

func (pd *PrefsDomain) DomainName() DomainName {
	return pd.domain
}
func (pd *PrefsDomain) Name() kvfilters.Name {
	return kvfilters.Name(pd.domain)
}

func (pd *PrefsDomain) Code() kvfilters.Code {
	return kvfilters.Codify(string(pd.domain))
}

func (pd *PrefsDomain) AddKeyValue(value kvfilters.KeyValue) {
	pref, ok := value.(*Pref)
	if !ok {
		panicf("Cannot add type '%T' with value %#v to macprefs.PrefsDomain %#v",
			value,
			value,
			pd,
		)
	}
	pd.prefs = append(pd.prefs, pref)
}

func (pd *PrefsDomain) LogArgs() []any {
	return []any{
		logargs.PrefsDomain,
		pd.Name(),
	}
}

func (pd *PrefsDomain) String() string {
	return string(pd.Name())
}
func (pd *PrefsDomain) ErrorInfo() error {
	return fmt.Errorf("%s=%s", logargs.PrefsDomain, pd)
}

func (pd *PrefsDomain) HasPrefix(prefix string) bool {
	return strings.HasPrefix(string(pd.Name()), prefix)
}

// Prefs returns all available preferences for this domain
func (pd *PrefsDomain) Prefs() []*Pref {
	if !pd.prefsRetrieved {
		panicf("ERROR: Preferences domain '%s' must be retrieved before calling macprefs.PrefsDomain.Prefs()",
			pd.Name(),
		)
	}
	return pd.prefs
}
