package macprefs

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type Prefs []*Pref

func (pp *Prefs) Sort() {
	slices.SortFunc(*pp, func(a, b *Pref) int {
		return strings.Compare(
			strings.ToLower(string(a.Key())),
			strings.ToLower(string(b.Key())),
		)
	})
}

var _ kvfilters.KeyValue = (*Pref)(nil)

// Pref represents a preference with its Domain and name
type Pref struct {
	*PrefDefault
	value       string // raw string value
	err         error  // last error encountered
	Description string
	invalid     bool
}

func (p *Pref) SetValue(value string) {
	p.value = value
}

func (p *Pref) TypeName() (typ TypeName) {
	//var label *kvfilters.Label
	var prefType macosutil.PreferenceType

	typ = p.PrefDefault.typeName
	if typ != "" && typ != TypeName(macosutil.UnknownType) {
		goto end
	}
	//label = p.labels.GetNamedLabel(Type)
	//if label != nil {
	//	typ = TypeName(label.Value)
	//	goto end
	//}
	_, prefType = macosutil.ParsePrefValue(p.DefaultOrValue())
	typ = TypeName(prefType)
end:
	return typ
}

func (p *Pref) HasLabel(label *kvfilters.Label) bool {
	return p.PrefDefault.Labels().HasLabel(label)
}
func (p *Pref) HasLabels(ll ...*kvfilters.Label) bool {
	labelPtrs := make([]*kvfilters.Label, len(ll))
	labels := p.PrefDefault.Labels()
	for i, label := range labels.LabelPtrs() {
		labelPtrs[i] = label
	}
	return labels.ContainsAny(labelPtrs...)
}

func (p *Pref) Valid() bool {
	return !p.invalid
}

// PrefOpts are used to pass to NewPref() to set initial struct properties
type PrefOpts struct {
	Domain  DomainName
	Name    PrefName
	Value   string // raw string value
	Default string // raw string value
	Labels  *kvfilters.Labels
	//Kind    reflect.Kind // kind of the value
	Invalid bool
}

// NewPref creates a new Pref instance
func NewPref(opts PrefOpts) *Pref {
	dv := GetPrefDefault(opts.Domain, opts.Name, nil)
	if opts.Default != "" {
		dv.DefaultValue = opts.Default
	}
	if dv.labels == nil {
		dv.labels = opts.Labels
	}
	return &Pref{
		value:       opts.Value,
		invalid:     opts.Invalid,
		PrefDefault: dv,
	}
}

// NewPrefFromDefault creates a new Pref instance from a *PrefDefault instance
func NewPrefFromDefault(pd *PrefDefault) *Pref {
	return &Pref{
		PrefDefault: pd,
	}
}

// Retrieve fetches the preference value from the system
func (p *Pref) Retrieve() error {
	mp, err := macosutil.GetPreference(string(p.Domain), string(p.Name))
	if errors.Is(err, macosutil.ErrUnsupportedType) {
		err = nil
		p.value = mp.Description
		p.err = fmt.Errorf(p.value)
		p.invalid = true
		goto end
	}
	if err != nil {
		err = errors.Join(ErrUnexpectedPreferenceType, err)
		goto end
	}
	p.value = mp.Value
	p.Description = mp.Description
	p.PrefDefault = GetPrefDefault(p.Domain, p.Name, &PrefDefaultOpts{
		Kind:          0, // TODO set these values
		SupportedIn:   "",
		UnsupportedIn: "",
	})
	p.invalid = !mp.Valid()
end:
	return err
}

// IsDefault returns true if the property's current value is the same as its default value.
//func (p *Pref) IsDefault() (isDefault bool) {
//	if p.err != nil {
//		isDefault = false
//		goto end
//	}
//	switch p.Kind {
//	case reflect.String:
//		isDefault = p.value == p.DefaultValue
//	case reflect.Int:
//		isDefault = p.Int() == p.DefaultInt()
//	case reflect.Bool:
//		isDefault = p.value == p.DefaultValue
//		if !isDefault && p.DefaultValue == "" {
//			isDefault = true
//		}
//	}
//end:
//	return isDefault
//}

// Value returns the current value of the pref
func (p *Pref) Value() string {
	return p.value
}

// Default returns the default value of the pref, as best we can tell
func (p *Pref) Default() string {
	return p.PrefDefault.DefaultValue
}

// DefaultOrValue DefaultOrValue returns the default value unless it is empty
// then returns the current value
func (p *Pref) DefaultOrValue() string {
	if p.PrefDefault.DefaultValue != "" {
		return p.PrefDefault.DefaultValue
	}
	return p.value
}

// String returns the raw string value regardless of type
func (p *Pref) String() string {
	return p.value
}

// LogValue returns a value for logging
func (p *Pref) LogValue() string {
	return fmt.Sprintf("%s/%s=%s", p.Domain, p.Name, p.value)
}
func (p *Pref) Message() string {
	if p.err != nil {
		return p.value
	}
	return ""
}

func (p *Pref) Bool() bool {
	return p.value == "true"
}

// Int returns the value as an integer
func (p *Pref) Int() int {
	n, _ := strconv.Atoi(p.value)
	return n
}

// Err returns the last error encountered
func (p *Pref) Err() error {
	return p.err
}

func (p *Pref) DebugString() string {
	if p.PrefDefault == nil {
		return "nil"
	}
	return fmt.Sprintf("%s=%s", p.Id(), p.value)
}
