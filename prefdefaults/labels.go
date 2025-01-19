package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macprefs"
)

func GetTypeLabel(typ TypeName) (label *kvfilters.Label) {
	label, ok := macprefs.TypeLabelMap[PreferenceType(typ)]
	if !ok {
		label = UnknownType
	}
	return label
}
