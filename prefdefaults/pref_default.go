package prefdefaults

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
)

type Lookup map[PrefId]*PrefDefault

var lookup = make(Lookup)
var lookupMutex sync.Mutex

func LookupPrefDefault(id PrefId) *PrefDefault {
	pd, _ := lookup[id]
	return pd
}

func SetPrefDefault(pd *PrefDefault) {
	lookupMutex.Lock()
	lookup[pd.Id()] = pd
	lookupMutex.Unlock()
}

var _ kvfilters.KeyValue = (*PrefDefault)(nil)

// PrefDefault represents a preference's default value
type PrefDefault struct {
	Domain       DomainName
	Name         PrefName
	DefaultValue string       // raw string value for default
	Kind         reflect.Kind // kind of the value
	typeName     TypeName
	labels       *kvfilters.Labels
	Added        macosutil.VersionNumber
	Removed      macosutil.VersionNumber
}

func (pd *PrefDefault) HasLabels() bool {
	return pd.labels.HasLabels()
}

func (pd *PrefDefault) Type() TypeName {
	return pd.typeName
}

type PrefDefaultOpts struct {
	Kind    reflect.Kind
	Added   macosutil.VersionNumber
	Removed macosutil.VersionNumber
}

func NewPrefDefault(domain DomainName, name PrefName, opts *PrefDefaultOpts) *PrefDefault {
	if opts == nil {
		opts = &PrefDefaultOpts{
			Kind: reflect.Invalid,
		}
	}
	return &PrefDefault{
		Domain: domain,
		Name:   name,
		Kind:   opts.Kind,
		labels: kvfilters.NewLabels(
			&macpreflabels.UnknownType,
			// 'defaults' is a reasonable default as the alternate is 'setup'
			// so we'll discover issues when comparing to current settings and
			// realize we need to manually change it.
			&macpreflabels.Optional,
			// 'userManaged' is a reasonable default as we'll discover issues
			// when comparing to current settings and realize we need to manually
			// change it.
			&macpreflabels.UserManaged,
		),
		typeName: TypeName(macpreflabels.UnknownType.Value),
		Added:    opts.Added,
		Removed:  opts.Removed,
	}
}
func GetPrefDefault(domain DomainName, name PrefName, opts *PrefDefaultOpts) (pd *PrefDefault) {
	pd = LookupPrefDefault(NewPrefId(domain, name))
	if pd == nil {
		pd = NewPrefDefault(domain, name, opts)
	}
	return pd
}

func (pd *PrefDefault) Labels() *kvfilters.Labels {
	return pd.labels
}

//	func (pd *PrefDefault) SetTypeName(name TypeName) {
//		pd.typeName = name
//	}
func (pd *PrefDefault) SetLabels(labels *kvfilters.Labels) {
	pd.labels = labels
	pd.labels.SyncMap()
	typeLabel := pd.GetNamedLabel(macpreflabels.Type)
	if typeLabel != nil {
		pd.typeName = TypeName(typeLabel.Value)
	}
}

func (pd *PrefDefault) HasLabel(label *kvfilters.Label) bool {
	return pd.labels.HasLabel(label)
}

func (pd *PrefDefault) GetNamedLabel(name kvfilters.LabelName) (label *kvfilters.Label) {
	return pd.labels.GetNamedLabel(name)
}

func (pd *PrefDefault) Valid() bool {
	//TODO implement me
	panic("implement me")
}

func (pd *PrefDefault) UserManaged() bool {
	return pd.HasLabel(&macpreflabels.UserManaged)
}

func (pd *PrefDefault) String() string {
	return fmt.Sprintf("%s/%s=%s", pd.Domain, pd.Name, pd.DefaultValue)
}

func (pd *PrefDefault) Id() PrefId {
	return NewPrefId(pd.Domain, pd.Name)
}

func (pd *PrefDefault) LogValue() any {
	return fmt.Sprintf("%s (default=%s,labels=[%s])",
		NewPrefId(pd.Domain, pd.Name),
		pd.DefaultValue,
		pd.labels,
	)
}

func (pd *PrefDefault) Key() kvfilters.Code {
	return kvfilters.Code(pd.Name)
}

func (pd *PrefDefault) Value() string {
	return pd.DefaultValue
}

func (pd *PrefDefault) ErrorInfo() error {
	return errors.Join(
		fmt.Errorf("%s=%s", logargs.Key, pd.Name),
		fmt.Errorf("%s=%s", logargs.DefaultValue, pd.DefaultValue),
	)
}

func (pd *PrefDefault) SetKey(key kvfilters.Code) {
	pd.Name = PrefName(key)
}

func (pd *PrefDefault) SetValue(value string) {
	pd.DefaultValue = value
}
