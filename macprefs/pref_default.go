package macprefs

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/logargs"
)

var _ kvfilters.KeyValue = (*PrefDefault)(nil)

// PrefDefault represents a preference's default value
type PrefDefault struct {
	Domain        DomainName
	Name          PrefName
	DefaultValue  string       // raw string value for default
	Kind          reflect.Kind // kind of the value
	typeName      TypeName
	labels        *kvfilters.Labels
	SupportedIn   OSVersion
	UnsupportedIn OSVersion
}

func (pd *PrefDefault) SetKey(key kvfilters.Code) {
	pd.Name = PrefName(key)
}

func (pd *PrefDefault) SetValue(value string) {
	pd.DefaultValue = value
}

type PrefDefaultOpts struct {
	Kind          reflect.Kind
	SupportedIn   OSVersion
	UnsupportedIn OSVersion
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
			&UnknownType,
			// 'defaults' is a reasonable default as the alternate is 'setup'
			// so we'll discover issues when comparing to current settings and
			// realize we need to manually change it.
			&Optional,
			// 'userManaged' is a reasonable default as we'll discover issues
			// when comparing to current settings and realize we need to manually
			// change it.
			&UserManaged,
		),
		typeName:      TypeName(UnknownType.Value),
		SupportedIn:   opts.SupportedIn,
		UnsupportedIn: opts.UnsupportedIn,
	}
}
func GetPrefDefault(domain DomainName, name PrefName, opts *PrefDefaultOpts) (d *PrefDefault) {
	d = LookupPrefDefault(NewPrefId(domain, name))
	if d == nil {
		d = NewPrefDefault(domain, name, opts)
	}
	return d
}
func GetPrefId(domain DomainName, name PrefName) PrefId {
	return NewPrefId(domain, name)
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
	typeLabel := pd.GetNamedLabel(Type)
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
	return pd.HasLabel(&UserManaged)
}

func (pd *PrefDefault) String() string {
	return fmt.Sprintf("%s/%s=%s", pd.Domain, pd.Name, pd.DefaultValue)
}

func (pd *PrefDefault) Id() PrefId {
	return NewPrefIdFromDefault(pd)
}

func (pd *PrefDefault) LogValue() any {
	return fmt.Sprintf("%s (default=%s,labels=[%s])",
		GetPrefId(pd.Domain, pd.Name),
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
		fmt.Errorf("%s=%s", logargs.Value, pd.Value()),
	)
}
