package macprefs

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"gopkg.in/yaml.v3"
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
}

func (pd *PrefsDomain) DefaultsYAML(opts YAMLOpts) string {
	var labelValues []prefsyaml.LabelValue
	var typeName string

	sb := strings.Builder{}
	resource := prefsyaml.YAMLPrefsResource{
		APIVersion: LatestAPIVersion,
		KindName:   "defaults",
		MetaData: prefsyaml.YAMLMetadata{
			Domain: prefsyaml.DomainName(pd.domain),
		},
		Spec: prefsyaml.NewYAMLPrefSpec(),
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
		labels.DeleteNamedLabel(Type)
		if labels.HasLabels() {
			labelValues = sliceconv.Func(labels.Labels(), func(l *kvfilters.Label) prefsyaml.LabelValue {
				return prefsyaml.LabelValue(l.Value)
			})
		}
		typeName, _ = strings.CutSuffix(string(pref.TypeName()), "Type")
		resource.Spec.Prefs = append(resource.Spec.Prefs, prefsyaml.YAMLPref{
			MetaData: &resource.MetaData,
			Name:     prefsyaml.PrefName(pref.Name),
			Type:     prefsyaml.PrefType(typeName),
			Value:    prefsyaml.YAMLValue{Value: val},
			Default:  prefsyaml.YAMLValue{Value: def},
			Labels:   labelValues,
		})
	}

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	err := enc.Encode(resource)
	if err != nil {
		// This should really never happen, right?
		panicf("ERROR: Failed to YAML encode a Go value of type '%T'; %s", resource, err.Error())
	}
	err = enc.Close()
	if err != nil {
		// This should really never happen, right?
		panicf("ERROR: Failed to close YAML encoder; %s", err.Error())
	}
	sb.WriteString("\n---\n")
	sb.WriteString(buf.String())
	sb.WriteByte('\n')
	return sb.String()
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
		pd := LookupPrefDefault(pref.Id())
		if pd == nil {
			pd = NewPrefDefault(pref.Domain, pref.Name, &PrefDefaultOpts{
				Kind:          0, // TODO: Populate these opts
				SupportedIn:   "",
				UnsupportedIn: "",
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
