package prefdefaults

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/mikeschinkel/prefsctl/appinfo"
	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	prefsyaml2 "github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	typeSuffix = "Type"
)

var register = macprefs.RegisterPrefDefaultsFunc

func OSVersionLabel(code Code) *kvfilters.Label {
	return &kvfilters.Label{
		Name:  macprefs.MacOS,
		Value: kvfilters.LabelValue(code),
	}
}

type (
	Labels            = kvfilters.Labels
	YAMLPrefsResource = prefsyaml2.YAMLPrefsResource
	YAMLPrefSpec      = prefsyaml2.YAMLPrefSpec
	YAMLPref          = prefsyaml2.YAMLPref
	YAMLMetadata      = prefsyaml2.YAMLMetadata

	KindName   = prefsyaml2.KindName
	PrefName   = prefsyaml2.PrefName
	DomainName = prefsyaml2.DomainName
	OSVersion  = prefsyaml2.OSVersion
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

//go:embed pref-defaults.yaml
var prefDefaultsYAML []byte

func init() {
	register(func(cfg config.Config) (macprefs.OSPrefDefaults, error) {
		return loadPrefDefaultsYAML(cfg)
	})
}

// loadPrefDefaultsYAML loads defaults YAML file from ~/.config found in
// prefsctl/pref-defaults.yaml, and first creates it from embedded source if it
// does not already exist.
//
// TODO: Allow user to override this or any portion of this. Overrides should be
//
//	able to be 1.) global from a file in ~/.config, 2.) from a file in the
//	current dir, or 3. from a file specified via CLI flag.
func loadPrefDefaultsYAML(cfg config.Config) (osDefaults macprefs.OSPrefDefaults, err error) {
	var resources []YAMLPrefsResource

	defaultsFile := filepath.Join(cfg.Dir(), appinfo.PrefDefaultsFile)

	exists, err := stdlibex.FileExists(defaultsFile)
	if err != nil {
		goto end
	}
	if !exists {
		err = os.WriteFile(defaultsFile, prefDefaultsYAML, os.ModePerm)
	}
	if err != nil {
		goto end
	}
	resources, err = prefsyaml2.LoadYAMLPrefsResources(defaultsFile)
	if err != nil {
		goto end
	}
	// Setup slice to capture list of domains and their preference defaults
	osDefaults.Domains = make([]macprefs.PrefDomain, 0)

	// Add all domains from the YAML file
	for _, resource := range resources {
		osDefaults.Domains = append(
			osDefaults.Domains,
			newPrefDomain(resource),
		)
	}
end:
	return osDefaults, err
}

func newPrefDomain(resource YAMLPrefsResource) macprefs.PrefDomain {
	domain := resource.MetaData.Domain
	pd := macprefs.NewPrefDomain(macprefs.DomainName(domain), resource.MetaData.KillOnApply)
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
	opts := macprefs.PrefDefaultOpts{
		Kind:          0, // TODO fill these out,
		SupportedIn:   "",
		UnsupportedIn: "",
	}
	pd = macprefs.GetPrefDefault(dn, pn, &opts)
	if pd == nil {
		pd = macprefs.NewPrefDefault(dn, pn, &opts)
	}
	if !yp.Default.Empty() {
		pd.DefaultValue = yp.Default.String()
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
