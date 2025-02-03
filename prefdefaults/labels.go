package prefdefaults

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
)

func GetTypeLabel(typ TypeName) (label *kvfilters.Label) {
	label, ok := macpreflabels.TypeLabelMap[PreferenceType(typ)]
	if !ok {
		label = UnknownType
	}
	return label
}
