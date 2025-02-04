package prefsyaml

import (
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
)

const (
	DocumentTypeErrKey errutil.ErrKey = "document_type"
	ResourceIdErrKey   errutil.ErrKey = "resource_id"
)
const (
	PrefsKind    = KindName(macpreflabels.PrefsKind)
	DefaultsKind = KindName(macpreflabels.DefaultsKind)
)
