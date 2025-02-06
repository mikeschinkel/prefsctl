package macprefs

import (
	"github.com/mikeschinkel/prefsctl/appinfo"
)

const (
	appKey         = "prefs"
	apiVersion     = "v1beta1"
	fullApiVersion = appKey + "/" + apiVersion

	LatestAPIVersion = appinfo.LatestAPIVersion
)
