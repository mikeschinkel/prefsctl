package macprefs

import (
	"fmt"
	"strconv"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

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

func (p *Pref) TypeName() TypeName {
	label := p.labels.GetNamedLabel(Type)
	if label == nil {
		label = &UnknownType
	}
	return TypeName(label.Value)
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

// PrefArgs are used to pass to NewPref() to set initial struct properties
type PrefArgs struct {
	Domain  DomainName
	Name    PrefName
	Value   string // raw string value
	Default string // raw string value
	Labels  *kvfilters.Labels
	//Kind    reflect.Kind // kind of the value
	Invalid bool
}

// NewPref creates a new Pref instance
func NewPref(args PrefArgs) *Pref {
	dv := GetPrefDefault(args.Domain, args.Name)
	if args.Default != "" {
		dv.DefaultValue = args.Default
	}
	if dv.labels == nil {
		dv.labels = args.Labels
	}
	return &Pref{
		value:       args.Value,
		invalid:     args.Invalid,
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
	mp, err := macosutil.RetrievePreference(string(p.Domain), string(p.Name))
	if err == nil {
		p.value = mp.Value
		p.Description = mp.Description
		p.PrefDefault = GetPrefDefault(p.Domain, p.Name)
		p.invalid = !mp.Valid()
	}
	return p.err
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
