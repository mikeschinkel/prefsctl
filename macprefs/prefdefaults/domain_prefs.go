package prefdefaults

import (
	"reflect"

	"github.com/mikeschinkel/prefsctl/kvfilters"
)

type (
	DomainPrefs = map[string]DomainPref
)

type DomainPref struct {
	Domain    string
	Name      string
	Type      string
	Descr     string
	Options   []string
	Default   string // raw string value for default
	Labels    *kvfilters.Labels
	kind      reflect.Kind
	typeLabel *kvfilters.Label
}

func (dp DomainPref) Kind() reflect.Kind {
	if dp.kind != reflect.Invalid {
		goto end
	}
	dp.kind = GetPrefKind(TypeName(dp.Type), dp.Default)
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
	if dp.typeLabel != nil {
		goto end
	}
	_, dp.typeLabel = GetPrefKindAndTypeLabel(dp.Kind(), TypeName(dp.Type), dp.Default)
end:
	return dp.typeLabel
}
