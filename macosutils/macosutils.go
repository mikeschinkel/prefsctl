package macosutils

type MacOSUtils interface {
	RetrievePreferenceDomains() (domains []PreferenceDomain, err error)
	RetrievePreferences(d PreferenceDomain) (prefs []*Preference, err error)
	RetrievePreference(domain string, name string) (dp *Preference, err error)
	GetVersionNumber() (v VersionNumber, err error)
	VersionName() (name VersionName, err error)
	VersionCode() (code VersionCode, err error)
	MustGetVersionCode() VersionCode
}

func New() MacOSUtils {
	return &macOSUtils{}
}

var _ MacOSUtils = (*macOSUtils)(nil)

type macOSUtils struct{}

func (macOSUtils) RetrievePreferenceDomains() (domains []PreferenceDomain, err error) {
	return retrievePreferenceDomains()
}

func (macOSUtils) RetrievePreferences(d PreferenceDomain) (prefs []*Preference, err error) {
	return retrievePreferences(d)
}

func (macOSUtils) RetrievePreference(domain string, name string) (dp *Preference, err error) {
	return retrievePreference(domain, name)
}

func (macOSUtils) GetVersionNumber() (v VersionNumber, err error) {
	return getVersionNumber()
}

func (macOSUtils) VersionName() (name VersionName, err error) {
	return versionName()
}

func (macOSUtils) VersionCode() (code VersionCode, err error) {
	return versionCode()
}

func (u macOSUtils) MustGetVersionCode() VersionCode {
	return mustGetVersionCode()
}
