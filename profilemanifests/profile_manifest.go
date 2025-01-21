package profilemanifests

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/micromdm/plist"
	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/macosutil"
)

// ProfileManifest represents the top-level structure of a preference manifest file
//
//goland:noinspection SpellCheckingInspection
type ProfileManifest struct {
	Description      string        `plist:"pfm_description"`
	DescriptionRef   string        `plist:"pfm_description_reference"`
	DocumentationURL string        `plist:"pfm_documentation_url"`
	Domain           string        `plist:"pfm_domain"`
	FormatVersion    int           `plist:"pfm_format_version"`
	Interaction      string        `plist:"pfm_interaction"`
	LastModified     time.Time     `plist:"pfm_last_modified"`
	MacOSMin         string        `plist:"pfm_macos_min"`
	MacOSMax         string        `plist:"pfm_macos_max"`
	Platforms        []string      `plist:"pfm_platforms"`
	Subkeys          []ManifestKey `plist:"pfm_subkeys"`
	Targets          []string      `plist:"pfm_targets"`
	Title            string        `plist:"pfm_title"`
	Unique           bool          `plist:"pfm_unique"`
	Version          int           `plist:"pfm_version"`
}

// SaveProfileManifestFile saves a preference manifest to a file path
func (m *ProfileManifest) SaveProfileManifestFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		err = errors.Join(ErrFailedToCreateManifestFile, err)
		goto end
	}
	defer mustClose(file)
	err = m.SaveProfileManifest(file)
	if err != nil {
		err = errors.Join(ErrFailedToSaveManifestFile, err)
		goto end
	}
end:
	return err
}

// SaveProfileManifest writes a preference manifest to an io.Writer
func (m *ProfileManifest) SaveProfileManifest(w io.Writer) error {
	encoder := plist.NewEncoder(w)
	encoder.Indent("    ")
	err := encoder.Encode(m)
	if err != nil {
		err = errors.Join(ErrFailedToEncodeManifest, err)
		goto end
	}
end:
	return err
}

// GetKey searches for a key by name in the manifest's sub-keys
func (m *ProfileManifest) GetKey(name string) (mk *ManifestKey, err error) {
	for _, key := range m.Subkeys {
		if key.Name == name {
			mk = &key
			goto end
		}
	}
	err = errors.Join(
		ErrKeyNotFoundInManifest,
		fmt.Errorf("%s=%s", ManifestKeyLogArg, name),
	)
	goto end
end:
	return mk, err
}

// Validate performs basic validation of the manifest structure
func (m *ProfileManifest) Validate() (err error) {
	if m.Domain == "" {
		err = ErrMissingPreferenceDomain
		goto end
	}
	if m.FormatVersion <= 0 {
		err = errors.Join(
			ErrInvalidManifestFormatVersion,
			fmt.Errorf("%s=%d", FormatVersionLogArg, m.FormatVersion),
		)
		goto end
	}
end:
	return err
}

// Helper methods for common key types

// GetBoolKey retrieves a boolean key and its default value
func (m *ProfileManifest) GetBoolKey(name string) (key *ManifestKey, defaultValue bool, err error) {
	key, err = m.GetKey(name)
	if err != nil {
		goto end
	}
	if key.Type != "boolean" {
		err = errors.Join(
			ErrKeyNotBooleanType,
			fmt.Errorf("%s=%s", ManifestKeyLogArg, name),
			fmt.Errorf("%s=%s", KeyTypeLogArg, key.Type),
		)
		goto end
	}
	if key.Default != nil {
		defaultValue = key.Default.(bool)
	}
end:
	return key, defaultValue, err
}

// GetIntKey retrieves an integer key and its default value
func (m *ProfileManifest) GetIntKey(name string) (key *ManifestKey, defaultValue int, err error) {
	key, err = m.GetKey(name)
	if err != nil {
		goto end
	}
	if key.Type != "integer" {
		err = errors.Join(
			ErrKeyNotIntegerType,
			fmt.Errorf("%s=%s", ManifestKeyLogArg, name),
			fmt.Errorf("%s=%s", KeyTypeLogArg, key.Type),
		)
		goto end
	}
	if key.Default != nil {
		defaultValue = key.Default.(int)
	}
end:
	return key, defaultValue, err
}

// GetStringKey retrieves a string key and its default value
func (m *ProfileManifest) GetStringKey(name string) (key *ManifestKey, defaultValue string, err error) {
	key, err = m.GetKey(name)
	if err != nil {
		goto end
	}
	if key.Type != "string" {
		err = errors.Join(
			ErrKeyNotStringType,
			fmt.Errorf("%s=%s", ManifestKeyLogArg, name),
			fmt.Errorf("%s=%s", KeyTypeLogArg, key.Type),
		)
		goto end
	}
	if key.Default != nil {
		defaultValue = key.Default.(string)
	}
end:
	return key, defaultValue, err

}

// ManifestKey represents a key definition in the manifest
//
//goland:noinspection ALL
type ManifestKey struct {
	Manifest         *ProfileManifest `plist:"-"`
	Default          interface{}      `plist:"pfm_default,omitempty"`
	Description      string           `plist:"pfm_description"`
	Name             string           `plist:"pfm_name"`
	Type             string           `plist:"pfm_type"`
	Title            string           `plist:"pfm_title"`
	RangeList        []interface{}    `plist:"pfm_range_list,omitempty"`
	RangeListTitles  []string         `plist:"pfm_range_list_titles,omitempty"`
	Required         string           `plist:"pfm_require,omitempty"`
	MacOSMin         string           `plist:"pfm_macos_min,omitempty"`
	MacOSMax         string           `plist:"pfm_macos_max,omitempty"`
	Hidden           string           `plist:"pfm_hidden,omitempty"`
	Subkeys          []ManifestKey    `plist:"pfm_subkeys,omitempty"`
	ValueProcessor   string           `plist:"pfm_value_processor,omitempty"`
	Format           string           `plist:"pfm_format,omitempty"`
	RangeMinAny      any              `plist:"pfm_range_min,omitempty"`
	RangeMaxAny      any              `plist:"pfm_range_max,omitempty"`
	AllowedFileTypes []string         `plist:"pfm_allowed_file_types,omitempty"`
}

// RangeMax returns contained value of RangeMaxAny as a string
func (k *ManifestKey) RangeMax() string {
	return fmt.Sprintf("%v", k.RangeMaxAny)
}

// RangeMin returns contained value of RangeMinAny as a string
func (k *ManifestKey) RangeMin() string {
	return fmt.Sprintf("%v", k.RangeMinAny)
}

// Kind returns reflect.Kind corresponding to the type of the value k.Default
func (k *ManifestKey) Kind() (rk reflect.Kind) {
	switch k.Type {
	case "string":
		rk = reflect.String
	case "boolean":
		rk = reflect.Bool
	case "integer":
		rk = reflect.Int64
	default:
		rk = reflect.Invalid
		slog.Error(
			"default type not handled in profilemanifests.ManifestKey.typeKind()",
			k.LogArgs()...,
		)
		goto end
	}
	if rk != k.typeKind() {
		slog.Warn("Mismatch in type of Default and declared type", k.LogArgs()...)
	}
end:
	return rk
}

// LogArgs returns array of log args for the manifest key
func (k *ManifestKey) LogArgs() []any {
	return []any{
		logargs.PrefsDomain, k.Domain(),
		logargs.PrefName, k.Name,
		logargs.DeclaredType, k.Type,
		logargs.DefaultType, fmt.Sprintf("%T", k.Default),
	}
}

// Domain returns the domain for the manifest key
func (k *ManifestKey) Domain() macosutil.DomainName {
	if k.Manifest == nil {
		return "<manifest_not_set>"
	}
	return macosutil.DomainName(k.Manifest.Domain)
}

// typeKind returns reflect.Kind corresponding to the value of k.Type field
func (k *ManifestKey) typeKind() (rk reflect.Kind) {
	rk = reflect.Invalid
	if k.Default == nil {
		slog.Error("default value not set", k.LogArgs()...)
		goto end
	}
	switch k.Default.(type) {
	case string:
		rk = reflect.String
	case bool:
		rk = reflect.Bool
	case int:
		rk = reflect.Int64
	default:
		slog.Error("default type not handled in profilemanifests.ManifestKey.Kind()",
			logargs.PrefsDomain, k.Manifest.Domain,
			logargs.PrefType, fmt.Sprintf("%T", k.Default),
		)
		goto end
	}
end:
	return rk
}

// SupportsOSVersion tests to see if the key is supported by the current macOS
// version, assuming that MacOSMin is not empty, otherwise assume supported.
func (k *ManifestKey) SupportsOSVersion() (is bool) {
	var vn macosutil.VersionNumber
	var err error

	if k.MacOSMin == "" && k.MacOSMax == "" {
		// Since we can't tell when it was supported assume it is
		is = true
		goto end
	}
	vn, err = macosutil.GetVersionNumber()
	if err != nil {
		panic(err.Error())
	}
	is = supportsOSVersion(string(vn), k.MacOSMin, k.MacOSMax)
end:
	return is
}

func parseMajorVersion(v string) (int, error) {
	parts := strings.Split(v, ".")
	if len(parts) == 0 {
		return 0, fmt.Errorf("invalid version format")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	return major, nil
}

func (k *ManifestKey) IsPayloadKey() (is bool) {
	if len(k.Name) < 10 { // len("PayloadType")==10
		goto end
	}
	if k.Name[:7] != "Payload" {
		goto end
	}
	switch k.Name[7:] {
	case "Description", "DisplayName", "Identifier", "Type", "UUID", "Version":
		is = true
		goto end
	}
end:
	return is
}

// LoadProfileManifestFromFile loads a preference manifest from a file path
func LoadProfileManifestFromFile(path string) (m *ProfileManifest, err error) {
	file, err := os.Open(path)
	if err != nil {
		err = errors.Join(ErrFailedToOpenManifestFile, err)
		goto end
	}
	defer mustClose(file)
	m, err = LoadProfileManifest(file)
end:
	return m, err
}

// LoadProfileManifest loads a preference manifest from an io.Reader
func LoadProfileManifest(r io.Reader) (_ *ProfileManifest, err error) {
	decoder := plist.NewXMLDecoder(r)
	pm := ProfileManifest{}
	err = decoder.Decode(&pm)
	if err != nil {
		err = errors.Join(ErrFailedToDecodeManifest, err)
		goto end
	}
	for i, subKey := range pm.Subkeys {
		subKey.Manifest = &pm
		pm.Subkeys[i] = subKey
	}
end:
	return &pm, err
}
