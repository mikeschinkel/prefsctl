package kvfilters

import (
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/slogutil"
)

var slog = slogutil.Logger().With(logargs.GoPackageLogArg, "kvfilters")
