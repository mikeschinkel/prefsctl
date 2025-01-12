package prefdefaults

import (
	"reflect"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type ErrorFunc func() error
type PrefsMap map[string]DomainPref

type Domain struct {
	Domain         macosutil.DomainName `yaml:"Domain"`
	AfterApplyFunc ErrorFunc            `yaml:"-"`
	Prefs          []DomainPref         `yaml:"Prefs"`
}

type DomainPref struct {
	Domain    string            `yaml:"-"`
	Name      string            `yaml:"Name"`
	Type      string            `yaml:"Type"`
	Descr     string            `yaml:"-"`
	Requires  []string          `yaml:"-"`
	Options   []string          `yaml:"-"`
	Default   string            `yaml:"Default"` // raw string value for default
	Range     []string          `yaml:"-"`
	Labels    *kvfilters.Labels `yaml:"-"`
	kind      reflect.Kind      `yaml:"-"`
	typeLabel *kvfilters.Label  `yaml:"-"`
	Class     *kvfilters.Label  `yaml:"-"`
}

func (dp DomainPref) Kind() reflect.Kind {
	if dp.kind != reflect.Invalid {
		goto end
	}
	dp.kind = macosutil.GetPrefKind(PreferenceType(dp.Type), dp.Default)
end:
	return dp.kind
}

func (dp DomainPref) TypeName() (name TypeName) {
	if dp.Type == "" {
		name = TypeName(UnknownType.Value)
		goto end
	}
	name = TypeName(string(dp.Type) + typeSuffix)
end:
	return name
}

func (dp DomainPref) TypeLabel() (label *kvfilters.Label) {
	var prefType PreferenceType

	if dp.typeLabel != nil {
		goto end
	}
	_, prefType = macosutil.GetPrefKindAndType(dp.Kind(), PreferenceType(dp.Type), dp.Default)
	dp.typeLabel = GetTypeLabel(TypeName(prefType))
end:
	return dp.typeLabel
}
