package prefdefaults

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/mikeschinkel/prefsctl/macprefs/preftemplates"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	typeSuffix    = "Type"
	defaultsYAML  = "defaults.yaml"
	yamlFilesPath = "yaml/osversions"
)

var register = macprefs.RegisterPrefDefaultsFunc

var (
	SequoiaLabel  = OSVersionLabel(macosutil.Sequoia)
	MontereyLabel = OSVersionLabel(macosutil.Monterey)
)

func OSVersionLabel(code Code) *kvfilters.Label {
	return &kvfilters.Label{
		Name:  macprefs.MacOS,
		Value: kvfilters.LabelValue(code),
	}
}

type (
	Labels            = kvfilters.Labels
	YAMLPrefsResource = preftemplates.YAMLPrefsResource
	YAMLPrefSpec      = preftemplates.YAMLPrefSpec
	YAMLPref          = preftemplates.YAMLPref
	YAMLMetadata      = preftemplates.YAMLMetadata

	KindName   = preftemplates.KindName
	PrefName   = preftemplates.PrefName
	DomainName = preftemplates.DomainName
	OSVersion  = preftemplates.OSVersion
)

var NewLabels = kvfilters.NewLabels

var FinalizedLabels = NewLabels(
	Optional,
)
var FinalizedUserManaged = FinalizedLabels.SetLabel(UserManaged)
var FinalizedRuntimeState = FinalizedLabels.SetLabel(RuntimeState)

var (
	InvalidLabel = macprefs.InvalidLabel
	Sets         = macprefs.Sets
	Type         = macprefs.Type
	MacOS        = macprefs.MacOS
	Class        = macprefs.Class

	AppManaged       = &macprefs.AppManaged
	UserManaged      = &macprefs.UserManaged
	VersionMigration = &macprefs.VersionMigration
	RuntimeState     = &macprefs.RuntimeState
	SystemManaged    = &macprefs.SystemManaged
	Optional         = &macprefs.Optional
	Required         = &macprefs.Required
	UnknownType      = &macprefs.UnknownType
	StringType       = &macprefs.StringType
	BoolType         = &macprefs.BoolType
	IntBoolType      = &macprefs.IntBoolType
	IntType          = &macprefs.IntType
	FloatType        = &macprefs.FloatType
)

type OSPrefDefaults struct {
	Domains []Domain
}

//go:embed yaml/osversions/*.yaml
var osFiles embed.FS

func init() {
	register(SequoiaLabel, func(cfg config.Config) (macprefs.OSPrefDefaults, error) {
		return prepareMacprefsOSPrefDefaults(cfg, sequoiaPrefDefaults())
	})
	register(MontereyLabel, func(cfg config.Config) (macprefs.OSPrefDefaults, error) {
		return prepareMacprefsOSPrefDefaults(cfg, montereyPrefDefaults())
	})
}

var file *os.File

func convertToResources(domains []Domain) []YAMLPrefsResource {
	resources := make([]YAMLPrefsResource, len(domains))
	osVersion, err := macosutil.VersionCode()
	if err != nil {
		panic(err)
	}
	for i, domain := range domains {

		resource := YAMLPrefsResource{
			APIVersion: macprefs.LatestAPIVersion,
			KindName:   KindName(macprefs.DefaultsKind),
			MetaData: preftemplates.YAMLMetadata{
				Domain:    DomainName(domain.Domain),
				OSVersion: OSVersion(osVersion),
			},
			Spec: YAMLPrefSpec{
				Prefs: make([]YAMLPref, len(domain.Prefs)),
			},
		}
		for j, pref := range domain.Prefs {
			resource.Spec.Prefs[j] = YAMLPref{
				MetaData: &resource.MetaData,
				Name:     PrefName(pref.Name),
				Default:  pref.Default,
			}
		}
		resources[i] = resource
	}
	return resources
}

// prepareMacprefsOSPrefDefaults is designed to make manually
// managing the defaults easy via DomainDefaults but then allow the program to
// have what it needs in macprefs.OSPrefDefaults.
//
// The impetus for creating two data structures and converting was to allow
// `OSPrefDefaults` to have a `Labels` property and then allow
// `macprefs.PrefDomain` to support the `Labels` method required by the
// `kvfilters.KeyValue` interface.
func prepareMacprefsOSPrefDefaults(cfg config.Config, defaults OSPrefDefaults) (os macprefs.OSPrefDefaults, err error) {
	var resources []preftemplates.YAMLPrefsResource
	var domainMap map[DomainName]int
	var metaDataMap = make(map[DomainName]*YAMLMetadata)
	var osCode macosutil.Code

	// Load non-verified defaults from YAML file
	resources, err = loadPrefDefaultsResources(cfg)
	if err != nil {
		goto end
	}

	// Create a quick lookup map by domains
	domainMap = sliceconv.ToIndexMapFunc(resources, func(r YAMLPrefsResource) DomainName {
		return r.MetaData.Domain
	})

	// Setup slice to capture list of domains and their preference defaults
	os.Domains = make([]macprefs.PrefDomain, 0)

	// Add all domains from the YAML file
	for _, resource := range resources {
		os.Domains = append(os.Domains, newPrefDomainFromResource(resource))
		metaDataMap[resource.MetaData.Domain] = &resource.MetaData
	}
	osCode = macosutil.MustGetVersionCode()
	// Now add all verified domains from Go code.
	for _, domain := range defaults.Domains {
		domainName := DomainName(domain.Domain)
		domainIndex, ok := domainMap[domainName]
		if !ok {
			// If domain is NOT in YAML, add it
			os.Domains = append(os.Domains, newPrefDomainFromDomain(domain, &YAMLMetadata{
				Domain:    domainName,
				OSVersion: OSVersion(osCode),
			}))
			continue
		}
		// If domain IS in YAML, replace it
		prefDomain := newPrefDomainFromDomain(domain, metaDataMap[domainName])
		os.Domains[domainIndex] = prefDomain
	}
end:
	return os, err
}

func newPrefDomainFromResource(resource YAMLPrefsResource) macprefs.PrefDomain {
	domain := resource.MetaData.Domain
	pd := macprefs.NewPrefDomain(macprefs.DomainName(domain), nil)
	for _, def := range resource.Spec.Prefs {
		def.MetaData = &resource.MetaData
		osDefault := convertToMacOSPrefsDefault(&def)
		if osDefault == nil {
			continue
		}
		pd.Defaults = append(pd.Defaults, osDefault)
	}
	return pd
}

func newPrefDomainFromDomain(domain Domain, md *preftemplates.YAMLMetadata) macprefs.PrefDomain {
	pd := macprefs.NewPrefDomain(macprefs.DomainName(domain.Domain), nil)
	for _, def := range domain.Prefs {
		def.Domain = string(domain.Domain)
		pd.Defaults = append(pd.Defaults, convertToMacOSPrefsDefault(&YAMLPref{
			MetaData: md,
			Name:     PrefName(def.Name),
			Default:  def.Default,
			Labels:   nil,
		}))
	}
	return pd
}

// convertToMacOSPrefsDefault converts a `DomainPref` to a `*macprefs.PrefDefault` with
// caveats, the caveats being if no `.Default` value is set — e.g. `.Default==""`
// — then `PrefDefault.DefaultValue` will be set by the current value from macOS
// and `.Verified` will be set to `false`. Also Kind
func convertToMacOSPrefsDefault(yp *YAMLPref) (pd *macprefs.PrefDefault) {
	var prefType, labelType PreferenceType
	var pdLabels *kvfilters.Labels
	var typeLabel, setsLabel, classLabel *kvfilters.Label
	var unknownType macosutil.PreferenceType

	dn := macprefs.DomainName(yp.MetaData.Domain)
	pn := macprefs.PrefName(yp.Name)
	pd = macprefs.GetPrefDefault(dn, pn)
	if pd == nil {
		pd = macprefs.NewPrefDefault(dn, pn)
	}
	if yp.Default != "" {
		pd.DefaultValue = yp.Default
	} else {
		p, err := macosutil.GetPreference(string(dn), string(pn))
		if err != nil {
			pd = nil
			goto end
		}
		pd.DefaultValue = p.Value
		pd.Kind = p.Kind
	}
	pd.Kind, prefType = macosutil.GetPrefKindAndType(
		pd.Kind,
		yp.PreferenceType(),
		pd.DefaultValue,
	)

	pdLabels = pd.Labels()
	for _, labelValue := range yp.Labels {
		value := kvfilters.LabelValue(labelValue)
		label := macprefs.GetLabelByValue(value)
		if !pdLabels.HasLabel(label) {
			*label = kvfilters.NewLabel(label.Name, value)
		}
		pdLabels.SetLabel(label)
	}

	unknownType = macosutil.PreferenceType(macprefs.UnknownType.Value)
	labelType = macosutil.PreferenceType(pdLabels.GetNamedLabel(macprefs.Type).Value)
	if prefType != labelType && prefType != unknownType {
		typeLabel = macprefs.GetLabelByValue(kvfilters.LabelValue(prefType))
		pdLabels.SetLabel(typeLabel)
	}
	if setsLabel = pdLabels.GetNamedLabel(macprefs.Sets); setsLabel == nil {
		pdLabels.SetLabel(&macprefs.Optional)
	}
	if classLabel = pdLabels.GetNamedLabel(macprefs.Class); classLabel == nil {
		pdLabels.SetLabel(&macprefs.UserManaged)
	}

	pd.SetLabels(pdLabels)
end:
	return pd
}

func loadPrefDefaultsResources(cfg config.Config) (_ []YAMLPrefsResource, err error) {
	var content []byte
	var resources []YAMLPrefsResource

	dir := cfg.Dir()
	osCode := macosutil.MustGetVersionCode()
	defaultsFile := filepath.Join(dir, fmt.Sprintf("%s-%s", osCode, defaultsYAML))

	exists, err := stdlibex.FileExists(defaultsFile)
	if err != nil {
		goto end
	}
	if !exists {
		var yamlFilepath = filepath.Join(yamlFilesPath, fmt.Sprintf("%s.yaml", osCode))
		// If file does not exist in config directory
		content, err = osFiles.ReadFile(yamlFilepath)
		if err != nil {
			goto end
		}
		err = os.WriteFile(defaultsFile, content, os.ModePerm)
	} else {
		content, err = os.ReadFile(defaultsFile)
	}
	if err != nil {
		goto end
	}
	resources, err = preftemplates.LoadYAMLPrefsResources(defaultsFile)
end:
	return resources, err
}
