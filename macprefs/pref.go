package macprefs

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/prefdefaults"
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
	*prefdefaults.PrefDefault
	value string // raw string value
	err   error  // last error encountered
}

func (p *Pref) SetValue(value string) {
	p.value = value
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

// Valid is needed to implement *kvfilters.KeyValue
func (p *Pref) Valid() bool {
	return true
}

// PrefArgs are used to pass to NewPref() to set initial struct properties
type PrefArgs struct {
	Domain      DomainName
	Name        PrefName
	Description string
	Value       string // raw string value
	Default     string // raw string value
	Labels      *kvfilters.Labels
	Added       macosutil.VersionNumber
	Removed     macosutil.VersionNumber
	Kind        reflect.Kind // kind of the value
	Type        TypeName
}

func (opts *PrefArgs) GetLabels() *kvfilters.Labels {
	labels := opts.Labels
	if labels == nil {
		labels = kvfilters.NewLabels()
	}
	return labels
}

// NewPref creates a new Pref instance
func NewPref(args PrefArgs) *Pref {
	dv := prefdefaults.NewPrefDefault(args.Domain, prefdefaults.PrefName(args.Name), &prefdefaults.PrefDefaultArgs{
		Type:        args.Type,
		Description: args.Description,
		Added:       args.Added,
		Removed:     args.Removed,
		Kind:        args.Kind,
		Default:     args.Default,
	})
	if args.Default != "" {
		dv.Default = args.Default
	}
	if !dv.HasLabels() {
		dv.SetLabels(args.GetLabels())
	}
	return &Pref{
		value:       args.Value,
		PrefDefault: dv,
	}
}

// NewPrefFromDefault creates a new Pref instance from a *PrefDefault instance
func NewPrefFromDefault(pd *prefdefaults.PrefDefault) *Pref {
	return &Pref{
		PrefDefault: pd,
	}
}

// Value returns the current value of the pref
func (p *Pref) Value() string {
	return p.value
}

// Default returns the default value of the pref, as best we can tell
func (p *Pref) Default() string {
	return p.PrefDefault.Default
}

// DefaultOrValue DefaultOrValue returns the default value unless it is empty
// then returns the current value
func (p *Pref) DefaultOrValue() string {
	if p.PrefDefault.Default != "" {
		return p.PrefDefault.Default
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
