package macosutil

type MacOSUtil interface {
	RetrievePreferenceDomains(RetrievalArgs) (domains []PreferenceDomain, err error)
	RetrievePreferences(d PreferenceDomain) (prefs []*Preference, err error)
	RetrievePreference(domain string, name string) (dp *Preference, err error)
	GetVersionNumber() (v VersionNumber, err error)
	VersionName() (name Name, err error)
	GetVersionName(code Code) (name Name)
	GetVersionCode(num VersionNumber) (code Code)
	VersionCode() (code Code, err error)
	MustGetVersionCode() Code
	ValidateVersionNumber(VersionNumber) bool
	ValidateVersionName(Name) bool
	ValidVersionNumbers() []VersionNumber
}

var instance MacOSUtil = New()

func New() MacOSUtil {
	return &macOSUtils{}
}

var _ MacOSUtil = (*macOSUtils)(nil)

type macOSUtils struct{}

func RetrievePreferenceDomains(args RetrievalArgs) (domains []PreferenceDomain, err error) {
	return instance.RetrievePreferenceDomains(args)
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
func GetVersionName(code Code) (name Name) {
	return instance.GetVersionName(code)
}
func GetVersionCode(num VersionNumber) (code Code) {
	return instance.GetVersionCode(num)
}
func ValidateVersionNumber(num VersionNumber) bool {
	return instance.ValidateVersionNumber(num)
}
func ValidateVersionName(name Name) bool {
	return instance.ValidateVersionName(name)
}
func ValidVersionNumbers() []VersionNumber {
	return instance.ValidVersionNumbers()
}

func VersionCode() (code Code, err error) {
	return instance.VersionCode()
}

func MustGetVersionCode() Code {
	return instance.MustGetVersionCode()
}
