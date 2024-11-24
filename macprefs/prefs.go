package macprefs

import (
	"errors"
	"slices"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/logging"
)

type Map map[PrefId]*Pref

type PrefsState struct {
	Domains     []Domain
	Prefs       [][]*Pref
	Unsupported map[string]struct{}
	Map         Map
}

func (s *PrefsState) Query() (err error) {
	var errs errutil.MultiErr
	var domains []Domain
	var index int
	var prefs []*Pref

	domains, err = GetPrefDomains()
	if err != nil {
		err = errors.Join(ErrFailedToGetPrefDomains, err)
		goto end
	}

	for _, domain := range domains {
		if !domain.HasPrefix("com.apple.") {
			// TODO: Maybe we'll support more later domains later
			continue
		}
		prefs, err = domain.Prefs()
		if err != nil {
			errs.Add(errors.Join(
				ErrFailedToGetDomainPrefs,
				errutil.AnnotateErr(err, "%s=%s", logging.PrefsDomainLogArg, domain),
			))
			continue
		}
		header := false
		for id, pref := range prefs {
			if err = pref.Retrieve(); err != nil {
				if !errors.Is(err, ErrUnsupportedType) {
					errs.Add(errutil.AnnotateErr(err, "%s=%s", logging.PrefIdLogArg, id))
					continue
				}
				s.AddUnsupported(pref.Message())
				continue
			}
			if !header {
				index = s.AddDomain(domain)
				header = true
			}
			s.AddPrefAtIndex(index, pref)
		}
	}

	if errs.IsError() {
		err = errs.Err()
		goto end
	}

end:
	return err
}

func (s *PrefsState) AddUnsupported(msg string) {
	s.Unsupported[msg] = struct{}{}
}
func (s *PrefsState) AddDomain(d Domain) (index int) {
	index = len(s.Domains)
	s.Domains = append(s.Domains, d)
	return index
}
func (s *PrefsState) AddPrefAtIndex(index int, p *Pref) {
	switch {
	case index > len(s.Prefs):
		panicf("index %d out of range [0-%d] when adding macpress.PrefState preference", index, len(s.Prefs)-1)
	case index == len(s.Prefs):
		// Need to add a new pref slice for this domain
		s.Prefs = append(s.Prefs, make([]*Pref, 0))
	default:
		// Pref slice already added, carry on
	}
	s.Prefs[index] = append(s.Prefs[index], p)
}

func NewPrefsState() *PrefsState {
	return &PrefsState{
		Domains:     make([]Domain, 0),
		Prefs:       make([][]*Pref, 0),
		Unsupported: make(map[string]struct{}),
	}
}

func NewPrefsStateFromMap(prefsMap Map) *PrefsState {
	s := NewPrefsState()
	s.Map = prefsMap
	domainIndex := make(map[Domain]int, 0)
	for _, pref := range prefsMap {
		index, ok := domainIndex[pref.Domain]
		if ok {
			s.AddPrefAtIndex(index, pref)
			continue
		}
		index = s.AddDomain(pref.Domain)
		domainIndex[pref.Domain] = index
		s.AddPrefAtIndex(index, pref)
	}
	// TODO: Sort first by non-domain (e.g. GlobalPreferences) and then by domain
	// (e.g. com.apple.)
	slices.Sort(s.Domains)
	return s
}
