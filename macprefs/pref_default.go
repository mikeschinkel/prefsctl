package macprefs

import (
	"fmt"
)

// PrefDefault represents a preference's default value
type PrefDefault struct {
	Domain   Domain
	Name     string
	Value    string       // raw string value for default
	Type     PrefType     // type of the value
	WhatSets WhatSetsType // determine what sets the value
}

func (pd *PrefDefault) Id() PrefId {
	return GetPrefId(pd.Domain, pd.Name)
}
func (pd *PrefDefault) LogValue() any {
	return fmt.Sprintf("%s [type=%s,default=%s,setby=%s]",
		GetPrefId(pd.Domain, pd.Name),
		pd.Type,
		pd.Value,
		pd.WhatSets,
	)
}

func GetPrefId(domain Domain, name string) PrefId {
	return PrefId(string(domain) + "/" + name)
}
