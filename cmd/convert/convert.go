package main

//import (
//	"os"
//	"strings"
//
//	"github.com/mikeschinkel/prefsctl/macosutil"
//	"github.com/mikeschinkel/prefsctl/macprefs"
//	"github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"
//	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
//	"gopkg.in/yaml.v3"
//)
//
//const (
//	rootDir    = "./macprefs/prefdefaults/yaml"
//	inputFile  = rootDir + "/monterey.yaml"
//	outputFile = rootDir + "/monterey-defaults.yaml"
//)
//
//type (
//	DomainName = preftemplates.DomainName
//	OSVersion  = preftemplates.OSVersion
//	KindName   = preftemplates.KindName
//	PrefName   = preftemplates.PrefName
//
//	YAMLPrefsResource = preftemplates.YAMLPrefsResource
//	YAMLPrefSpec      = preftemplates.YAMLPrefSpec
//	YAMLPref          = preftemplates.YAMLPref
//)
//
//func main() {
//	var domains struct {
//		Domains []prefdefaults.Domain `yaml:"Domains"`
//	}
//	b, err := os.ReadFile(inputFile)
//	if err != nil {
//		panic(err)
//	}
//	err = yaml.Unmarshal(b, &domains)
//	if err != nil {
//		panic(err)
//	}
//	sb := strings.Builder{}
//	resources := convertToResources(domains.Domains)
//	for _, resource := range resources {
//		b, err = yaml.Marshal(resource)
//		if err != nil {
//			panic(err)
//		}
//		sb.WriteString("\n---\n")
//		sb.WriteString(string(b))
//		sb.WriteByte('\n')
//	}
//	b = []byte(sb.String())
//	err = os.WriteFile(outputFile, b, os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func convertToResources(domains []prefdefaults.Domain) []YAMLPrefsResource {
//	resources := make([]YAMLPrefsResource, len(domains))
//	osVersion, err := macosutil.VersionCode()
//	if err != nil {
//		panic(err)
//	}
//	for i, domain := range domains {
//
//		resource := YAMLPrefsResource{
//			APIVersion: macprefs.LatestAPIVersion,
//			KindName:   KindName(macprefs.DefaultsKind),
//			MetaData: preftemplates.YAMLMetadata{
//				Domain:    DomainName(domain.Domain),
//				OSVersion: OSVersion(osVersion),
//			},
//			Spec: YAMLPrefSpec{
//				Prefs: make([]YAMLPref, len(domain.Prefs)),
//			},
//		}
//		for j, pref := range domain.Prefs {
//			resource.Spec.Prefs[j] = YAMLPref{
//				MetaData: &resource.MetaData,
//				Name:     PrefName(pref.Name),
//				Default:  pref.Default,
//			}
//		}
//		resources[i] = resource
//	}
//	return resources
//}
