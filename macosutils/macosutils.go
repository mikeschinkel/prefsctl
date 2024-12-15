package macosutils

type MacOSUtils interface {
	RetrievePreferenceDomains() (domains []PreferenceDomain, err error)
	RetrievePreferences(d PreferenceDomain) (prefs []*Preference, err error)
	RetrievePreference(domain string, name string) (dp *Preference, err error)
	GetVersionNumber() (v VersionNumber, err error)
	VersionName() (name Name, err error)
	VersionCode() (code Code, err error)
	MustGetVersionCode() Code
}

var instance MacOSUtils = New()

func New() MacOSUtils {
	return &macOSUtils{}
}

var _ MacOSUtils = (*macOSUtils)(nil)

type macOSUtils struct{}

func RetrievePreferenceDomains() (domains []PreferenceDomain, err error) {
	return instance.RetrievePreferenceDomains()
}

func RetrievePreferences(d PreferenceDomain) (prefs []*Preference, err error) {
	return instance.RetrievePreferences(d)
}

func RetrievePreference(domain string, name string) (dp *Preference, err error) {
	return instance.RetrievePreference(domain, name)
}

func GetVersionNumber() (v VersionNumber, err error) {
	return instance.GetVersionNumber()
}

func VersionName() (name Name, err error) {
	return instance.VersionName()
}

func VersionCode() (code Code, err error) {
	return instance.VersionCode()
}

func MustGetVersionCode() Code {
	return instance.MustGetVersionCode()
}
