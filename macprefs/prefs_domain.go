package macprefs

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/mikeschinkel/prefsctl/yamlutil"
)

var _ kvfilters.Group = (*PrefsDomain)(nil)

// PrefsDomain represents a preference Domain in macOS
type PrefsDomain struct {
	Domain      DomainName
	Description string
	Process     ProcessName
	Added       VersionNumber
	Removed     VersionNumber
	prefs       Prefs
}

func (*PrefsDomain) FilterableEntry() {}

type YAMLOpts struct {
	APIVersion OSVersion
}

func (pd *PrefsDomain) GetPrefsYAMLResource(kind prefsyaml.KindName, opts YAMLOpts) (yr *prefsyaml.Resource) {
	var labelValues []*prefsyaml.LabelValue

	yr = &prefsyaml.Resource{
		APIVersion: APIVersion(opts.APIVersion),
		KindName:   kind,
		Spec:       prefsyaml.NewSpec(),
		MetaData: prefsyaml.Metadata{
			Domain: prefsyaml.DomainName(pd.Domain),
		},
	}
	for _, pref := range pd.Prefs() {
		var val, def string
		val = pref.Value()
		def = pref.Default()
		labels := pref.Labels()
		// Delete "type" from in YAML because as there is an explicit "type" field
		labels.DeleteNamedLabel(macpreflabels.Type)
		if labels.HasLabels() {
			labelValues = sliceconv.Func(labels.Labels(), func(l *kvfilters.Label) *prefsyaml.LabelValue {
				lv := prefsyaml.LabelValue(l.Value)
				return &lv
			})
		}
		yr.Spec.Prefs = append(yr.Spec.Prefs, prefsyaml.Pref{
			MetaData: &yr.MetaData,
			Name:     prefsyaml.PrefName(pref.Name),
			Type:     prefsyaml.PrefType(pref.Type),
			Value:    yamlutil.NewValue(val),
			Default:  yamlutil.NewValue(def),
			Labels:   labelValues,
			Kind:     reflect.Invalid, // TODO: Update this
		})
	}
	return yr
}

func (pd *PrefsDomain) GetYAMLDocument(kind prefsyaml.KindName, opts YAMLOpts) (yd yamlutil.Document, err error) {
	yr := pd.GetPrefsYAMLResource(kind, opts)
	yd, err = yamlutil.BuildDocument(yr)
	if err != nil {
		goto end
	}
end:
	return yd, err
}

// Valid is needed to implement kvfilters.Group
func (pd *PrefsDomain) Valid() bool {
	return true
}

// ShallowCopy is used to copy PrefsDomain without its children
func (pd *PrefsDomain) ShallowCopy() kvfilters.Group {
	return &PrefsDomain{
		Domain:      pd.Domain,
		Description: pd.Description,
		Process:     pd.Process,
		Added:       pd.Added,
		Removed:     pd.Removed,
		prefs:       make(Prefs, 0),
	}
}

func (pd *PrefsDomain) SortPrefs() {
	pd.prefs.Sort()
}

var unsupportedTypes = make(map[string]struct{})

type PrefsDomainArgs struct {
	Description string
	Process     ProcessName
	Added       VersionNumber
	Removed     VersionNumber
}

func NewPrefsDomain(domain DomainName, args PrefsDomainArgs) *PrefsDomain {
	return &PrefsDomain{
		Domain:      domain,
		Description: args.Description,
		Process:     args.Process,
		Added:       args.Added,
		Removed:     args.Removed,
		prefs:       make([]*Pref, 0),
	}
}

func (pd *PrefsDomain) KeyValues() (kvs []kvfilters.KeyValue) {
	kvs = make([]kvfilters.KeyValue, len(pd.prefs))
	for i, pref := range pd.prefs {
		kvs[i] = pref
	}
	return kvs
}

func (pd *PrefsDomain) Name() kvfilters.Name {
	return kvfilters.Name(pd.Domain)
}

func (pd *PrefsDomain) Code() kvfilters.Code {
	return kvfilters.Codify(string(pd.Domain))
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
	return pd.prefs
}
