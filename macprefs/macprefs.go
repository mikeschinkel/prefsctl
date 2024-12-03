package macprefs

import (
	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
	"github.com/mikeschinkel/prefsctl/slogutil"
)

var slog = slogutil.Logger().With(logargs.GoPackageLogArg, "macprefs")
