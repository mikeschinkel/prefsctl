package macprefs

import (
	"fmt"
)

type PrefId string

func NewPrefId(domain DomainName, pref PrefName) PrefId {
	return PrefId(fmt.Sprintf("%s/%s", domain, pref))
}
func NewPrefIdFromPref(pref *Pref) PrefId {
	return NewPrefId(pref.Domain, pref.Name)
}
func NewPrefIdFromDefault(d *PrefDefault) PrefId {
	return NewPrefId(d.Domain, d.Name)
}
