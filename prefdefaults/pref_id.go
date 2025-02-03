package prefdefaults

import (
	"fmt"
)

type PrefId string

func NewPrefId(domain DomainName, pref PrefName) PrefId {
	return PrefId(fmt.Sprintf("%s/%s", domain, pref))
}
