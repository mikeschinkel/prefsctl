package prefdefaults

import (
	"os"
	"strings"

	"github.com/mikeschinkel/prefsctl/kvfilters"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/mikeschinkel/prefsctl/maputil"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	typeSuffix = "Type"
)

var register = macprefs.RegisterDefaultsMapFunc

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

type Labels = kvfilters.Labels

var NewLabels = kvfilters.NewLabels

var FinalizedLabels = NewLabels(
	DefaultsSet,
	TypeVerified,
	ClassVerified,
	SetsVerified,
	DefaultVerified,
	DescrVerified,
)
var FinalizedUserManaged = FinalizedLabels.SetLabel(UserManaged)
var FinalizedRuntimeState = FinalizedLabels.SetLabel(RuntimeState)
var FinalizedUserManagedWithOptions = FinalizedUserManaged.SetLabel(OptionsVerified)

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
	DefaultsSet      = &macprefs.DefaultsSet
	SetupSets        = &macprefs.SetupSets
	UnknownType      = &macprefs.UnknownType
	StringType       = &macprefs.StringType
	BoolType         = &macprefs.BoolType
	IntBoolType      = &macprefs.IntBoolType
	IntType          = &macprefs.IntType
	FloatType        = &macprefs.FloatType

	TypeVerified    = &macprefs.TypeVerified
	DefaultVerified = &macprefs.DefaultVerified
	DescrVerified   = &macprefs.DescrVerified
	SetsVerified    = &macprefs.SetsVerified
	ClassVerified   = &macprefs.ClassVerified
	OptionsVerified = &macprefs.OptionsVerified

	TypeNotVerified  = &macprefs.TypeNotVerified
	DescrNotVerified = &macprefs.DescrNotVerified
)

type DomainDefaults = map[string]Domain

func init() {
	register(SequoiaLabel, func() macprefs.DomainPrefDefaults {
		return convertDomainDefaultsToMacprefsDomainPrefDefaults(sequoiaPrefDefaults())
	})
	register(MontereyLabel, func() macprefs.DomainPrefDefaults {
		return convertDomainDefaultsToMacprefsDomainPrefDefaults(montereyPrefDefaults())
	})
}

var file *os.File

// convertDomainDefaultsToMacprefsDomainPrefDefaults is designed to make manually
// managing the defaults easy via DomainDefaults but then allow the program to
// have what it needs in macprefs.DomainPrefDefaults.
//
// The impetus for creating two data structures and converting was to allow
// `DomainDefaults` to have a `Labels` property and then allow
// `macprefs.PrefDefaults` to support the `Labels` method required by the
// `kvfilters.KeyValue` interface.
func convertDomainDefaultsToMacprefsDomainPrefDefaults(defaults DomainDefaults) (dpd macprefs.DomainPrefDefaults) {
	dpd = make(macprefs.DomainPrefDefaults, len(defaults))
	for domain, dd := range maputil.SortedKeysIterator(defaults) {
		pd := macprefs.PrefDefaults{
			Domain:         macprefs.DomainName(domain),
			AfterApplyFunc: dd.AfterApplyFunc,
			DefaultsMap:    make(macprefs.PrefDefaultsMap, len(dd.Prefs)),
		}
		for name, def := range maputil.SortedKeysIterator(dd.Prefs) {
			def.Domain = domain
			def.Name = name
			pd.DefaultsMap[macprefs.PrefName(name)] = getPrefDefaultFromDomainPref(def)
		}
		dpd[macprefs.DomainName(domain)] = pd
	}
	return dpd
}
func fixCase(s string) string {
	titler := cases.Title(language.English)
	return strings.Replace(titler.String(s), "type", "Type", 1)
}

// getPrefDefaultFromDomainPref converts a `prefdefaults.DomainPref` to a
// `*macprefs.PrefDefault` with caveats, the caveats being if no `.Default` value
// is set — e.g. `.Default==""` — then `PrefDefault.DefaultValue` will be set by
// the current value from macOS and `.Verified` will be set to `false`. Also Kind
func getPrefDefaultFromDomainPref(def DomainPref) (pd *macprefs.PrefDefault) {
	var prefType PreferenceType

	dn := macprefs.DomainName(def.Domain)
	pn := macprefs.PrefName(def.Name)
	pd = macprefs.GetPrefDefault(dn, pn)
	if pd == nil {
		pd = macprefs.NewPrefDefault(dn, pn)
	}
	if def.Default != "" {
		pd.DefaultValue = def.Default
	} else {
		p, err := macosutil.GetPreference(def.Domain, def.Name)
		if err == nil {
			pd.DefaultValue = p.Value
			pd.Kind = p.Kind
		}
	}
	pd.Kind, prefType = macosutil.GetPrefKindAndType(
		pd.Kind,
		macosutil.PreferenceType(def.TypeName()),
		pd.DefaultValue,
	)

	def.Labels.SetLabel(GetTypeLabel(TypeName(prefType)))

	pdLabels := pd.Labels()
	sets := pdLabels.GetNamedLabel(macprefs.Sets)
	switch {
	case sets != nil:
		def.Labels.SetLabel(sets)
	case def.Labels.GetNamedLabel(macprefs.Sets) == nil:
		def.Labels.SetLabel(&macprefs.DefaultsSet)
	}

	class := pdLabels.GetNamedLabel(macprefs.Class)
	switch {
	case class != nil:
		def.Labels.SetLabel(class)
	case def.Labels.GetNamedLabel(macprefs.Class) == nil:
		def.Labels.SetLabel(&macprefs.UserManaged)
	}

	pd.SetLabels(def.Labels)
	return pd
}
