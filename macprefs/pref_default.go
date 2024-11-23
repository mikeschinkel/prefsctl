package macprefs

import (
	"fmt"
)

// YAMLPrefDefault represents a preference's default value
type YAMLPrefDefault struct {
	Name         string       `yaml:"name"`
	DefaultValue string       `yaml:"default"`  // raw string value for default
	WhatSets     WhatSetsType `yaml:"whatSets"` // determine what sets the value
	Kind         PrefType     `yaml:"kind"`     // type of the value
	Options      []string     `yaml:"options"`
}

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
