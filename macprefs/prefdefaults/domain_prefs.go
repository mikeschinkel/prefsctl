package prefdefaults

import (
	"reflect"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type ErrorFunc func() error
type PrefsMap map[string]DomainPref
type Domain struct {
	AfterApplyFunc ErrorFunc
	Prefs          PrefsMap
}

type DomainPref struct {
	Domain    string
	Name      string
	Type      string
	Descr     string
	Requires  []string
	Options   []string
	Default   string   // raw string value for default
	Range     []string // raw string value for default
	Labels    *kvfilters.Labels
	kind      reflect.Kind
	typeLabel *kvfilters.Label
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
