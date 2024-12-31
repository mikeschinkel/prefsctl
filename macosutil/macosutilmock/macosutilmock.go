package macosutilmock

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/types"
)

type MacOSUtil = macosutil.MacOSUtil

type RetrievalArgs = macosutil.RetrievalArgs
type DomainName = macosutil.DomainName
type Preference = macosutil.Preference
type PreferenceDomain = macosutil.PreferenceDomain
type VersionNumber = macosutil.VersionNumber
type Identifier = macosutil.Identifier
type NULL = struct{}

type VersionCode = types.Code
type VersionName = types.Name

var _ MacOSUtil = (*MacOSUtilMock)(nil)

type MacOSUtilMock struct {
	macosutil.MacOSUtil
	Domains     []PreferenceDomain
	DomainsErr  error
	DomainPrefs map[PreferenceDomain][]*Preference
	DomainErrs  map[PreferenceDomain]error
	Prefs       map[Identifier]*Preference
	Version     struct {
		Number    VersionNumber
		NumberErr error
		Code      VersionCode
		CodeErr   error
		Name      VersionName
		NameErr   error
	}
}
