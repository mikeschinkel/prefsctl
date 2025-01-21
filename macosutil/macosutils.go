package macosutil

type MacOSUtil interface {
	GetPreferenceDomains(RetrievalArgs) (domains []PreferenceDomain, err error)
	GetPreferences(d PreferenceDomain) (prefs []*Preference, err error)
	GetPreference(domain string, name string) (dp *Preference, err error)
	GetVersionNumber() (v VersionNumber, err error)
	VersionName() (name Name, err error)
	GetVersionName(code Code) (name Name)
	GetVersionCode(num VersionNumber) (code Code)
	VersionCode() (code Code, err error)
	MustGetVersionNumber() VersionNumber
	MustGetVersionCode() Code
	ValidateVersionNumber(VersionNumber) bool
	ValidateVersionName(Name) bool
	ValidVersionNumbers() []VersionNumber
	ApplyPreferences(domain string, prefs []*Preference) (err error)
	SetPreference(domain, name, value string) (err error)
}

var instance MacOSUtil = New()

func New() MacOSUtil {
	return &macOSUtils{}
}

var _ MacOSUtil = (*macOSUtils)(nil)

type macOSUtils struct{}

func GetPreferenceDomains(args RetrievalArgs) (domains []PreferenceDomain, err error) {
	return instance.GetPreferenceDomains(args)
}

func GetPreferences(d PreferenceDomain) (prefs []*Preference, err error) {
	return instance.GetPreferences(d)
}

func GetPreference(domain string, name string) (dp *Preference, err error) {
	return instance.GetPreference(domain, name)
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

func MustGetVersionNumber() VersionNumber {
	return instance.MustGetVersionNumber()
}

func ApplyPreferences(domain string, prefs []*Preference) (err error) {
	return instance.ApplyPreferences(domain, prefs)
}

func SetPreference(domain, name, value string) (err error) {
	return instance.SetPreference(domain, name, value)
}
