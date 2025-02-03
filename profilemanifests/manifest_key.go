package profilemanifests

// ManifestKey represents a key definition in the manifest
//
//goland:noinspection ALL
type ManifestKey struct {
	Manifest         *ProfileManifest `yaml:"-"`
	Default          *Value           `plist:"pfm_default,omitempty"`
	Description      string           `plist:"pfm_description"`
	Name             string           `plist:"pfm_name"`
	Type             string           `plist:"pfm_type"`
	Title            string           `plist:"pfm_title"`
	RangeList        []*Value         `plist:"pfm_range_list,omitempty"`
	RangeListTitles  []string         `plist:"pfm_range_list_titles,omitempty"`
	Required         string           `plist:"pfm_require,omitempty"`
	MacOSMin         string           `plist:"pfm_macos_min,omitempty"`
	MacOSMax         string           `plist:"pfm_macos_max,omitempty"`
	Hidden           string           `plist:"pfm_hidden,omitempty"`
	Subkeys          []ManifestKey    `plist:"pfm_subkeys,omitempty"`
	ValueProcessor   string           `plist:"pfm_value_processor,omitempty"`
	Format           string           `plist:"pfm_format,omitempty"`
	RangeMin         *Value           `plist:"pfm_range_min,omitempty"`
	RangeMax         *Value           `plist:"pfm_range_max,omitempty"`
	AllowedFileTypes []string         `plist:"pfm_allowed_file_types,omitempty"`
}

//// Kind returns reflect.Kind corresponding to the type of the value k.Default
//func (k *ManifestKey) Kind() (rk reflect.Kind) {
//	switch k.Type {
//	case "string":
//		rk = reflect.String
//	case "boolean":
//		rk = reflect.Bool
//	case "integer":
//		rk = reflect.Int64
//	default:
//		rk = reflect.Invalid
//		slog.Error(
//			"default type not handled in profilemanifests.ManifestKey.typeKind()",
//			k.LogArgs()...,
//		)
//		goto end
//	}
//	kind,err := k.Default.Kind()
//	if rk !=  {
//		slog.Warn("Mismatch in type of Default and declared type", k.LogArgs()...)
//	}
//end:
//	return rk
//}

//// LogArgs returns array of log args for the manifest key
//func (k *ManifestKey) LogArgs() []any {
//	return []any{
//		logargs.PrefsDomain, k.Domain(),
//		logargs.PrefName, k.Name,
//		logargs.DeclaredType, k.Type,
//		logargs.DefaultType, fmt.Sprintf("%T", k.Default),
//	}
//}

//// Domain returns the domain for the manifest key
//func (k *ManifestKey) Domain() macosutil.DomainName {
//	if k.Manifest == nil {
//		return "<manifest_not_set>"
//	}
//	return macosutil.DomainName(k.Manifest.Domain)
//}

//// typeKind returns reflect.Kind corresponding to the value of k.Type field
//func (k *ManifestKey) typeKind() (rk reflect.Kind) {
//	rk = reflect.Invalid
//	if k.Default == nil {
//		slog.Error("default value not set", k.LogArgs()...)
//		goto end
//	}
//	switch k.Default.(type) {
//	case string:
//		rk = reflect.String
//	case bool:
//		rk = reflect.Bool
//	case int:
//		rk = reflect.Int64
//	default:
//		slog.Error("default type not handled in profilemanifests.ManifestKey.Kind()",
//			logargs.PrefsDomain, k.Manifest.Domain,
//			logargs.PrefType, fmt.Sprintf("%T", k.Default),
//		)
//		goto end
//	}
//end:
//	return rk
//}
