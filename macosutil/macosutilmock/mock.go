package macosutilmock

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

func MockMacOSUtil(data Data) {
	mock := NewMock()
	macosutil.SetInstance(mock)
	m := mock.(*MacOSUtilMock)
	m.SetPreferenceDomains(data.Domains, data.DomainsErr)
	m.SetPreferences(data.DomainPrefs, data.DomainsPrefErrs)
}

func NewMock() MacOSUtil {
	return &MacOSUtilMock{
		MacOSUtil:   macosutil.New(),
		Domains:     make([]PreferenceDomain, 0),
		DomainPrefs: make(map[PreferenceDomain][]*Preference),
		DomainErrs:  make(map[PreferenceDomain]error),
		Prefs:       make(map[Identifier]*Preference),
	}
}

var _ MacOSUtil = (*MacOSUtilMock)(nil)

var mock = NewMock()

func (mock *MacOSUtilMock) SetPreferenceDomains(domains []PreferenceDomain, err error) {
	mock.Domains = domains
	mock.DomainsErr = err
}

func (mock *MacOSUtilMock) RetrievePreferenceDomains(args RetrievalArgs) (domains []PreferenceDomain, err error) {
	if len(args.Domains) == 0 {
		domains = mock.Domains
		goto end
	}
	domains = filterPreferenceDomains(mock.Domains, filter{
		IncludeDomains: mock.Domains,
	})
end:
	return domains, mock.DomainsErr
}

func (mock *MacOSUtilMock) SetPreferences(prefs map[PreferenceDomain][]*Preference, errs map[PreferenceDomain]error) {
	mock.DomainPrefs = prefs
	mock.DomainErrs = errs
	for _, pp := range prefs {
		for _, p := range pp {
			mock.Prefs[p.Id()] = p
		}
	}
}
func (mock *MacOSUtilMock) GetPreferences(d PreferenceDomain) (prefs []*Preference, err error) {
	//return mock.MacOSUtil.GetPreferences(d)
	var ok bool

	prefs, ok = mock.DomainPrefs[d]
	if !ok {
		err = macosutil.ErrFailedToGetPrefDomains
		goto end
	}

	err, _ = mock.DomainErrs[d]

end:
	return prefs, err
}

func (mock *MacOSUtilMock) GetPreference(domain string, name string) (pref *Preference, err error) {
	//return mock.MacOSUtil.GetPreference(domain, name)
	var ok bool

	id := macosutil.PreferenceId(domain, name)
	err, ok = mock.DomainErrs[PreferenceDomain(domain)]
	if !ok {
		err = macosutil.ErrInvalidPreferenceDomain
		goto end
	}
	pref, ok = mock.Prefs[id]
	if !ok {
		err = macosutil.ErrPreferenceNotFound
		goto end
	}
end:
	//err = macosutil.ErrInvalidInput
	//err = macosutil.ErrUnsupportedType
	//err = macosutil.ErrUnknownPreferenceError
	return pref, err
}

type filter struct {
	IncludeDomains []PreferenceDomain
}

func filterPreferenceDomains(domains []PreferenceDomain, filter filter) []PreferenceDomain {
	index := make(map[PreferenceDomain]NULL, len(filter.IncludeDomains))
	for _, domain := range filter.IncludeDomains {
		index[domain] = NULL{}
	}
	dd := make([]PreferenceDomain, 0, len(filter.IncludeDomains))
	for _, domain := range domains {
		if _, ok := index[domain]; !ok {
			continue
		}
		dd = append(dd, domain)
	}
	return dd
}
