package prefdefaults

import (
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mikeschinkel/prefsctl/appinfo"
	"github.com/mikeschinkel/prefsctl/config"
	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	typeSuffix = "Type"
)

func OSVersionLabel(code Code) *kvfilters.Label {
	return &kvfilters.Label{
		Name:  macpreflabels.MacOS,
		Value: kvfilters.LabelValue(code),
	}
}

type (
	Labels       = kvfilters.Labels
	YAMLResource = prefsyaml.Resource
	YAMLSpec     = prefsyaml.Spec
	YAMLPref     = prefsyaml.Pref
	YAMLMetadata = prefsyaml.Metadata

	KindName  = prefsyaml.KindName
	OSVersion = prefsyaml.OSVersion
)

var NewLabels = kvfilters.NewLabels

var FinalizedLabels = NewLabels(
	Optional,
)
var FinalizedUserManaged = FinalizedLabels.SetLabel(UserManaged)
var FinalizedRuntimeState = FinalizedLabels.SetLabel(RuntimeState)

var (
	InvalidLabel = macpreflabels.InvalidLabel
	Sets         = macpreflabels.Sets
	Type         = macpreflabels.Type
	MacOS        = macpreflabels.MacOS
	Class        = macpreflabels.Class

	AppManaged       = &macpreflabels.AppManaged
	UserManaged      = &macpreflabels.UserManaged
	VersionMigration = &macpreflabels.VersionMigration
	RuntimeState     = &macpreflabels.RuntimeState
	SystemManaged    = &macpreflabels.SystemManaged
	Optional         = &macpreflabels.Optional
	Required         = &macpreflabels.Required
	UnknownType      = &macpreflabels.UnknownType
	StringType       = &macpreflabels.StringType
	BoolType         = &macpreflabels.BoolType
	IntBoolType      = &macpreflabels.IntBoolType
	IntType          = &macpreflabels.IntType
	FloatType        = &macpreflabels.FloatType
)

type OSPrefDefaults struct {
	Domains []*Domain
}

//go:embed pref-defaults.yaml
var prefDefaultsYAML []byte

func init() {
	RegisterPrefDefaultsFunc(func(cfg config.Config) (OSPrefDefaults, error) {
		return loadPrefDefaultsFromYAML(cfg)
	})
}

func WritePrefsDefaultsFile(cfg config.Config, content string) (err error) {
	defaultsFile := cfg.OtherFilepath(appinfo.PrefDefaultsFile)
	// First write a revision file
	err = os.WriteFile(insertFilenameTimestamp(defaultsFile), []byte(content), os.ModePerm)
	if err != nil {
		goto end
	}
	// Then write the actual file to use
	err = os.WriteFile(defaultsFile, []byte(content), os.ModePerm)
end:
	return err
}

func insertFilenameTimestamp(filename string) string {
	dir, base := filepath.Split(filename)

	// Split the base filename into name and extension
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	// Create timestamp in desired format
	timestamp := time.Now().Format("2006-01-02-15_04-05-00")

	// Combine everything back together
	return filepath.Join(dir, name+"."+timestamp+ext)
}

// loadPrefDefaultsYAML loads defaults YAMLDocument file from ~/.config found in
// prefsctl/pref-defaults.yaml, and first creates it from embedded source if it
// does not already exist.
//
// TODO: Allow user to override this or any portion of this. Overrides should be
//
//	able to be 1.) global from a file in ~/.config, 2.) from a file in the
//	current dir, or 3. from a file specified via CLI flag.
func loadPrefDefaultsFromYAML(cfg config.Config) (osDefaults OSPrefDefaults, err error) {
	var resources []YAMLResource

	defaultsFile := cfg.OtherFilepath(appinfo.PrefDefaultsFile)

	exists, err := stdlibex.FileExists(defaultsFile)
	if err != nil {
		goto end
	}
	if !exists {
		err = WritePrefsDefaultsFile(cfg, string(prefDefaultsYAML))
	}
	if err != nil {
		goto end
	}
	resources, err = prefsyaml.LoadPrefsResources(defaultsFile)
	if err != nil {
		goto end
	}
	// Setup slice to capture list of domains and their preference defaults
	osDefaults.Domains = make([]*Domain, 0)

	// Add all domains from the YAMLDocument file
	for _, resource := range resources {
		osDefaults.Domains = append(
			osDefaults.Domains,
			newDomainFromYAMLResource(resource),
		)
	}
end:
	return osDefaults, err
}

// convertToYAMLDefault converts a `DomainPref` to a `*prefdefaults.PrefDefault` with
// caveats, the caveats being if no `.Default` value is set — e.g. `.Default==""`
// — then `PrefDefault.DefaultValue` will be set by the current value from macOS
// and `.Verified` will be set to `false`. Also Kind
func convertToPrefDefault(dn DomainName, yd *prefsyaml.Default) (pd *PrefDefault) {
	var prefType, labelType PreferenceType
	var pdLabels *kvfilters.Labels
	var typeLabel, setsLabel, classLabel *kvfilters.Label
	var unknownType macosutil.PreferenceType
	pn := PrefName(yd.Name)
	opts := PrefDefaultOpts{
		Kind:    0, // TODO fill these out,
		Added:   "",
		Removed: "",
	}
	pd = GetPrefDefault(dn, pn, &opts)
	if pd == nil {
		pd = NewPrefDefault(dn, pn, &opts)
	}
	if pd.DefaultValue == "" {
		pd.DefaultValue = yd.Value.String()
	} else {
		p, err := macosutil.GetPreference(string(dn), string(pn))
		if err != nil {
			pd = nil
			goto end
		}
		pd.DefaultValue = p.Value
		pd.Kind = p.Kind
	}
	panic("The following needed to be verified and likely fixed. The Type value is probably wrong)")
	pd.Kind, prefType = macosutil.GetPrefKindAndType(
		pd.Kind,
		macosutil.PreferenceType(yd.Type),
		pd.DefaultValue,
	)

	pdLabels = pd.Labels()
	for _, labelValue := range yd.Labels {
		label := macpreflabels.GetLabelByValue(macpreflabels.LabelValue(*labelValue))
		if !pdLabels.HasLabel(label) {
			*label = kvfilters.NewLabel(label.Name, label.Value)
		}
		pdLabels.SetLabel(label)
	}

	unknownType = macosutil.PreferenceType(macpreflabels.UnknownType.Value)
	labelType = macosutil.PreferenceType(pdLabels.GetNamedLabel(macpreflabels.Type).Value)
	if prefType != labelType && prefType != unknownType {
		typeLabel = macpreflabels.GetLabelByValue(kvfilters.LabelValue(prefType))
		pdLabels.SetLabel(typeLabel)
	}
	if setsLabel = pdLabels.GetNamedLabel(macpreflabels.Sets); setsLabel == nil {
		pdLabels.SetLabel(&macpreflabels.Optional)
	}
	if classLabel = pdLabels.GetNamedLabel(macpreflabels.Class); classLabel == nil {
		pdLabels.SetLabel(&macpreflabels.UserManaged)
	}

	pd.SetLabels(pdLabels)
end:
	return pd
}
