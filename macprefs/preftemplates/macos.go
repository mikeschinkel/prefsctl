package preftemplates

type MacOS struct {
	Name          OSVersion
	Domains       []*Domain
	ShowValueFunc func(*Default) bool
}

func NewMacOS(name OSVersion, prefs []*Pref) *MacOS {
	os := &MacOS{
		Name:    name,
		Domains: make([]*Domain, 0),
		ShowValueFunc: func(d *Default) bool {
			return true
		},
	}
	return os
}
