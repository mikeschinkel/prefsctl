package macosutils

import (
	"github.com/mikeschinkel/prefsctl/logutil"
	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
)

var slog = logutil.Logger().With(logargs.GoPackageLogArg, "macosutils")
