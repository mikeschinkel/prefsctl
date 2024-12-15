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
	Domain       DomainName
	Name         PrefName
	DefaultValue string       // raw string value for default
	Kind         reflect.Kind // kind of the value
	typeName     TypeName
	labels       *kvfilters.Labels
	Verified     Verified
}

func NewPrefDefault(domain DomainName, name PrefName) *PrefDefault {
	return &PrefDefault{
		Domain: domain,
		Name:   name,
		Kind:   reflect.Invalid,
		labels: kvfilters.NewLabels(
			&UnknownType,
			// 'defaults' is a reasonable default as the alternate is 'setup'
			// so we'll discover issues when comparing to current settings and
			// realize we need to manually change it.
			&DefaultsSet,
			// 'userManaged' is a reasonable default as we'll discover issues
			// when comparing to current settings and realize we need to manually
			// change it.
			&UserManaged,
		),
		Verified: Verified{},
		typeName: TypeName(UnknownType.Value),
	}
}
func GetPrefDefault(domain DomainName, name PrefName) (d *PrefDefault) {
	d = LookupPrefDefault(NewPrefId(domain, name))
	if d == nil {
		d = NewPrefDefault(domain, name)
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
		fmt.Errorf("%s=%s", logargs.KeyLogArg, pd.Name),
		fmt.Errorf("%s=%s", logargs.ValueLogArg, pd.Value()),
	)
}
