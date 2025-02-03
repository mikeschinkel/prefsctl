package prefsyaml

//import (
//	"github.com/mikeschinkel/prefsctl/macosutil"
//)
//
//type Domain struct {
//	Name       DomainName            `yaml:"name"`
//	Description Description          `yaml:"description"`
//	Defaults   []*Default            `yaml:"defaults,omitempty"`
//	Prefs      []*PrefLite           `yaml:"prefs,omitempty"`
//	Process    macosutil.ProcessName `yaml:"process,omitempty"`
//	APIVersion APIVersion            `yaml:"api_version,omitempty"`
//	Added      OSVersion             `yaml:"added,omitempty"`
//	Removed    OSVersion             `yaml:"removed,omitempty"`
//}
//
//// FilterableEntry Implement profilemanifests.FilterableEntry interface
//func (yd *Domain) FilterableEntry() {}
//
//func (yd *Domain) AddDefault(def *Default) {
//	yd.Defaults = append(yd.Defaults, def)
//}
//
//type DomainOpts struct {
//	Description Description
//	Process    macosutil.ProcessName
//	APIVersion APIVersion
//	Added      OSVersion
//	Removed    OSVersion
//}
//
//func NewDomain(domain DomainName, opts DomainOpts) *Domain {
//	return &Domain{
//		Name:       domain,
//		Description: opts.Description,
//		Prefs:      make([]*PrefLite, 0),
//		Defaults:   make([]*Default, 0),
//		Process:    opts.Process,
//		APIVersion: opts.APIVersion,
//	}
//}
//
////func (yd *Domain) GetDocument(kind KindName) (yamlutil.Document, error) {
////	return yamlutil.BuildDocument(yd.GetResource(kind))
////}
////
////// GetResource
////func (yd *Domain) GetResource(kind KindName) (yr *Resource) {
////	yr = &Resource{
////		APIVersion: yd.APIVersion,
////		KindName:   kind,
////		MetaData: Metadata{
////			Domain:      yd.Name,
////			Process: yd.Process,
////			Added:       yd.Added,
////			Removed:     yd.Removed,
////		},
////		Spec: NewSpec(),
////	}
////	switch kind {
////	case KindName(macpreflabels.PrefsKind):
////		for _, pref := range yd.Prefs {
////			var typeName string
////			typeName, _ = strings.CutSuffix(string(pref.Name), "Type") //TODO: Isolate "Type" into a constant somewhere
////			yr.Spec.Prefs = append(yr.Spec.Prefs, Pref{
////				MetaData: &yr.MetaData,
////				Name:     pref.Name,
////				Type:     PrefType(typeName),
////				Default:  yamlutil.NewValue(pref.Value),
////				Labels:   pref.Labels,
////			})
////		}
////	case KindName(macpreflabels.DefaultsKind):
////		for _, def := range yd.Defaults {
////			var typeName string
////			typeName, _ = strings.CutSuffix(def.TypeName(), "Type") //TODO: Isolate "Type" into a constant somewhere
////			yr.Spec.Defaults = append(yr.Spec.Defaults, Default{
////				Name:   def.Name,
////				Type:   PrefType(typeName),
////				Value:  yamlutil.NewValue(def.Value),
////				Labels: def.Labels,
////			})
////		}
////	}
////	return yr
////}
