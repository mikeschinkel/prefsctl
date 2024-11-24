package macprefs

import (
	"fmt"
)

// PrefDefault represents a preference's default value
type PrefDefault struct {
	Domain Domain
	Name   string
	Value  string // raw string value for default
	Labels Labels
}

func (pd *PrefDefault) Id() PrefId {
	return GetPrefId(pd.Domain, pd.Name)
}
func (pd *PrefDefault) LogValue() any {
	return fmt.Sprintf("%s (default=%s,labels=[%s])",
		GetPrefId(pd.Domain, pd.Name),
		pd.Value,
		pd.Labels,
	)
}

func GetPrefId(domain Domain, name string) PrefId {
	return PrefId(string(domain) + "/" + name)
}
