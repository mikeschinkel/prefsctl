package macosutilsmock

import (
	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/types"
)

type MacOSUtils = macosutils.MacOSUtils

type Preference = macosutils.Preference
type PreferenceDomain = macosutils.PreferenceDomain
type VersionNumber = macosutils.VersionNumber

type VersionCode = types.Code
type VersionName = types.Name

func NewMock() MacOSUtils {
	return &macOSUtilsMock{}
}

var _ MacOSUtils = (*macOSUtilsMock)(nil)

type macOSUtilsMock struct{}

func (m *macOSUtilsMock) MustGetVersionCode() VersionCode {
	code, _ := m.VersionCode()
	return code
}

var values = struct {
	Domains        []PreferenceDomain
	DomainsErr     error
	DomainPrefs    map[PreferenceDomain][]*Preference
	DomainPrefsErr map[PreferenceDomain]error
	Prefs          map[string]*Preference
	PrefsErr       map[string]error
	Version        struct {
		Number    VersionNumber
		NumberErr error
		Code      VersionCode
		CodeErr   error
		Name      VersionName
		NameErr   error
	}
}{}

func SetPreferenceDomains(domains []PreferenceDomain, err error) {
	values.Domains = domains
	values.DomainsErr = err
}
func (*macOSUtilsMock) RetrievePreferenceDomains() (domains []PreferenceDomain, err error) {
	return values.Domains, values.DomainsErr
}

func SetPreferences(d PreferenceDomain, prefs []*Preference, err error) {
	values.DomainPrefs[d] = prefs
	values.DomainPrefsErr[d] = err
}
func (*macOSUtilsMock) RetrievePreferences(d PreferenceDomain) (prefs []*Preference, err error) {
	return values.DomainPrefs[d], values.DomainPrefsErr[d]
}
func prefId(domain string, name string) string {
	return domain + "|" + name
}
func SetPreference(domain string, name string, dp *Preference, err error) {
	id := prefId(domain, name)
	values.Prefs[id] = dp
	values.PrefsErr[id] = err
}

func (*macOSUtilsMock) RetrievePreference(domain string, name string) (dp *Preference, err error) {
	id := prefId(domain, name)
	return values.Prefs[id], values.PrefsErr[id]
}

func SetVersionNumber(number VersionNumber, err error) {
	values.Version.Number = number
	values.Version.NumberErr = err
}
func (*macOSUtilsMock) GetVersionNumber() (v VersionNumber, err error) {
	return values.Version.Number, values.Version.NumberErr
}
func SetVersionCode(code VersionCode, err error) {
	values.Version.Code = code
	values.Version.CodeErr = err
}
func (*macOSUtilsMock) VersionCode() (code VersionCode, err error) {
	return values.Version.Code, values.Version.CodeErr
}

func SetVersionName(name VersionName, err error) {
	values.Version.Name = name
	values.Version.NameErr = err
}
func (*macOSUtilsMock) VersionName() (name VersionName, err error) {
	return values.Version.Name, values.Version.NameErr
}
