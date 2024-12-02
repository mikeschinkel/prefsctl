package macprefs

import (
	"errors"
	"fmt"
	"slices"

	"github.com/mikeschinkel/prefsctl/logging"
	"github.com/mikeschinkel/prefsctl/macprefs/filters"
)

var _ filters.KeyValue = (*PrefDefault)(nil)

// PrefDefault represents a preference's default value
type PrefDefault struct {
	Domain       DomainName
	Name         PrefName
	DefaultValue string // raw string value for default
	labels       filters.Labels
	NoDefault    bool
}

func (pd *PrefDefault) SetLabels(labels []*filters.Label) {
	pd.labels = labels
}
func (pd *PrefDefault) Labels() []*filters.Label {
	return pd.labels
}

func (pd *PrefDefault) Valid() bool {
	//TODO implement me
	panic("implement me")
}

func NewPrefDefault(domain DomainName, name PrefName) *PrefDefault {
	return &PrefDefault{
		Domain:    domain,
		Name:      name,
		labels:    make(filters.Labels, 0),
		NoDefault: true,
	}
}

func (pd *PrefDefault) UserManaged() bool {
	return slices.Contains([]*filters.Label(pd.labels), &filters.UserManaged)
}

func (pd *PrefDefault) String() string {
	return fmt.Sprintf("%s/%s=%s", pd.Domain, pd.Name, pd.DefaultValue)
}

func (pd *PrefDefault) Id() PrefId {
	return GetPrefId(pd.Domain, pd.Name)
}
func (pd *PrefDefault) LogValue() any {
	return fmt.Sprintf("%s (default=%s,labels=[%s])",
		GetPrefId(pd.Domain, pd.Name),
		pd.DefaultValue,
		pd.labels,
	)
}

func GetPrefId(domain DomainName, name PrefName) PrefId {
	return PrefId(string(domain) + "/" + string(name))
}

func (pd *PrefDefault) Key() filters.Code {
	return filters.Code(pd.Name)
}

func (pd *PrefDefault) Value() string {
	return pd.DefaultValue
}

func (pd *PrefDefault) ErrorInfo() error {
	return errors.Join(
		fmt.Errorf("%s=%s", logging.KeyLogArg, pd.Name),
		fmt.Errorf("%s=%s", logging.ValueLogArg, pd.Value),
	)
}
