package profilemanifests

import (
	"reflect"
)

type TypeName string

// ManifestKey represents a key definition in the manifest
//
//goland:noinspection ALL
type ManifestKey struct {
	Manifest         *ProfileManifest `yaml:"-"`
	Default          *Value           `plist:"pfm_default,omitempty"`
	Description      string           `plist:"pfm_description"`
	Name             string           `plist:"pfm_name"`
	Type             TypeName         `plist:"pfm_type"`
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

const (
	StringType     TypeName = "string"
	BooleanType    TypeName = "boolean"
	IntegerType    TypeName = "integer"
	ArrayType      TypeName = "array"
	DataType       TypeName = "data"
	RealType       TypeName = "real"
	DateType       TypeName = "date"
	DictionaryType TypeName = "dictionary"
)

func (mk *ManifestKey) Kind() (kind reflect.Kind) {
	switch mk.Type {
	case StringType:
		kind = reflect.String
	case BooleanType:
		kind = reflect.Bool
	case IntegerType:
		kind = reflect.Int64
	case DataType:
		kind = reflect.Interface
	case DateType:
		kind = reflect.Struct
	case RealType:
		kind = reflect.Float64
	case ArrayType:
		kind = reflect.Slice
	case DictionaryType:
		kind = reflect.Map
	default:
		panicf("Type not yet handled: %s", mk.Type)
	}
	return kind
}
