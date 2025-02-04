package prefdefaults

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
)

var _ kvfilters.KeyValue = (*PrefDefault)(nil)

// PrefDefault represents a preference's default value
type PrefDefault struct {
	Domain      DomainName
	Name        PrefName
	Type        TypeName
	Description string
	Default     string // raw string value for default
	labels      *kvfilters.Labels
	Added       macosutil.VersionNumber
	Removed     macosutil.VersionNumber
	Kind        reflect.Kind // kind of the value
}

func (pd *PrefDefault) HasLabels() bool {
	return pd.labels.HasLabels()
}

type PrefDefaultArgs struct {
	Type        macosutil.PreferenceType
	Description string
	Default     string // raw string value for default
	Added       macosutil.VersionNumber
	Removed     macosutil.VersionNumber
	Kind        reflect.Kind
}

func NewPrefDefault(domain DomainName, name PrefName, args *PrefDefaultArgs) *PrefDefault {
	if args == nil {
		args = &PrefDefaultArgs{
			Kind: reflect.Invalid,
		}
	}
	return &PrefDefault{
		Domain: domain,
		Name:   name,
		Kind:   args.Kind,
		labels: kvfilters.NewLabels(
		// TODO: Decide if we need any labels at all
		//&macpreflabels.Optional,
		//// 'userManaged' is a reasonable default as we'll discover issues
		//// when comparing to current settings and realize we need to manually
		//// change it.
		//&macpreflabels.UserManaged,
		),
		Type:        args.Type,
		Description: args.Description,
		Added:       args.Added,
		Removed:     args.Removed,
	}
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
		pd.Type = TypeName(typeLabel.Value)
	}
}

func (pd *PrefDefault) HasLabel(label *kvfilters.Label) bool {
	return pd.labels.HasLabel(label)
}

func (pd *PrefDefault) GetNamedLabel(name kvfilters.LabelName) (label *kvfilters.Label) {
	return pd.labels.GetNamedLabel(name)
}

func (pd *PrefDefault) Valid() bool {
	return true
}

func (pd *PrefDefault) UserManaged() bool {
	return pd.HasLabel(&macpreflabels.UserManaged)
}

func (pd *PrefDefault) String() string {
	return fmt.Sprintf("%s/%s=%s", pd.Domain, pd.Name, pd.Default)
}

func (pd *PrefDefault) Id() PrefId {
	return NewPrefId(pd.Domain, pd.Name)
}

func (pd *PrefDefault) LogValue() any {
	return fmt.Sprintf("%s (default=%s,labels=[%s])",
		NewPrefId(pd.Domain, pd.Name),
		pd.Default,
		pd.labels,
	)
}

func (pd *PrefDefault) Key() kvfilters.Code {
	return kvfilters.Code(pd.Name)
}

func (pd *PrefDefault) Value() string {
	return pd.Default
}

func (pd *PrefDefault) ErrorInfo() error {
	return errors.Join(
		fmt.Errorf("%s=%s", logargs.Key, pd.Name),
		fmt.Errorf("%s=%s", logargs.DefaultValue, pd.Default),
	)
}

func (pd *PrefDefault) SetKey(key kvfilters.Code) {
	pd.Name = PrefName(key)
}

func (pd *PrefDefault) SetValue(value string) {
	pd.Default = value
}
