package prefdefaults

import (
	"os"
	"reflect"
	"strings"

	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/mikeschinkel/prefsctl/macprefs/kvfilters"
	"github.com/mikeschinkel/prefsctl/mapsutils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	typeSuffix = "Type"
)

var register = macprefs.RegisterDefaultsMapFunc

var (
	SequoiaLabel  = OSVersionLabel(macosutils.Sequoia)
	MontereyLabel = OSVersionLabel(macosutils.Monterey)
)

func OSVersionLabel(code macosutils.Code) *kvfilters.Label {
	return &kvfilters.Label{
		Name:  kvfilters.LabelName(macprefs.MacOS),
		Value: kvfilters.LabelValue(code),
	}
}

type Labels = kvfilters.Labels

var NewLabels = kvfilters.NewLabels

type Verified = macprefs.Verified

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
)

type DomainDefaults = map[string]DomainPrefs
type DomainPrefs = map[string]DomainPref
type DomainPref struct {
	Domain    string
	Name      string
	Type      string
	Default   string // raw string value for default
	Labels    *kvfilters.Labels
	Verified  Verified
	kind      reflect.Kind
	typeLabel *kvfilters.Label
}

func (dp DomainPref) Kind() reflect.Kind {
	if dp.kind != reflect.Invalid {
		goto end
	}
	dp.kind = GetPrefKind(TypeName(dp.Type), dp.Default)
end:
	return dp.kind
}

func (dp DomainPref) TypeName() (name TypeName) {
	if dp.Type == "" {
		name = TypeName(UnknownType.Value)
		goto end
	}
	name = TypeName(string(dp.Type) + typeSuffix)
end:
	return name
}

func (dp DomainPref) TypeLabel() (label *kvfilters.Label) {
	if dp.typeLabel != nil {
		goto end
	}
	_, dp.typeLabel = GetPrefKindAndTypeLabel(dp.Kind(), TypeName(dp.Type), dp.Default)
end:
	return dp.typeLabel
}

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
	for domain, defs := range mapsutils.KeysSorted(defaults) {
		pp := make(macprefs.PrefDefaultsMap, len(defs))
		for name, def := range mapsutils.KeysSorted(defs) {
			def.Domain = domain
			def.Name = name
			pp[macprefs.PrefName(name)] = getPrefDefaultFromDomainPref(def)
		}
		dpd[macprefs.DomainName(domain)] = pp
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
	var typeLabel *kvfilters.Label

	dn := macprefs.DomainName(def.Domain)
	pn := macprefs.PrefName(def.Name)
	pd = macprefs.GetPrefDefault(dn, pn)
	if pd == nil {
		pd = macprefs.NewPrefDefault(dn, pn)
	}
	if def.Default != "" {
		pd.DefaultValue = def.Default
		pd.Verified = def.Verified
		//pd.Kind = def.Kind()
	} else {
		p, err := macosutils.RetrievePreference(def.Domain, def.Name)
		if err == nil {
			pd.DefaultValue = p.Value
			pd.Kind = p.Kind
		}
		pd.Verified = Verified{}
	}
	pd.Kind, typeLabel = GetPrefKindAndTypeLabel(pd.Kind, def.TypeName(), pd.DefaultValue)

	def.Labels.Add(typeLabel)

	sets := pd.Labels().GetNamedLabel(macprefs.Sets)
	switch {
	case sets != nil:
		def.Labels.Add(sets)
	case def.Labels.GetNamedLabel(macprefs.Sets) == nil:
		def.Labels.Add(&macprefs.DefaultsSet)
	}

	class := pd.Labels().GetNamedLabel(macprefs.Class)
	switch {
	case class != nil:
		def.Labels.Add(class)
	case def.Labels.GetNamedLabel(macprefs.Class) == nil:
		def.Labels.Add(&macprefs.UserManaged)
	}

	pd.SetLabels(def.Labels)
	return pd
}
