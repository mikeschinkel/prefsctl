package applemdm

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	//"howett.net/plist"
	"github.com/micromdm/plist"
)

// PreferenceManifest represents the top-level structure of a preference manifest file
//
//goland:noinspection SpellCheckingInspection
type PreferenceManifest struct {
	Description      string        `plist:"pfm_description"`
	DescriptionRef   string        `plist:"pfm_description_reference"`
	DocumentationURL string        `plist:"pfm_documentation_url"`
	Domain           string        `plist:"pfm_domain"`
	FormatVersion    int           `plist:"pfm_format_version"`
	Interaction      string        `plist:"pfm_interaction"`
	LastModified     time.Time     `plist:"pfm_last_modified"`
	MacOSMin         string        `plist:"pfm_macos_min"`
	Platforms        []string      `plist:"pfm_platforms"`
	Subkeys          []ManifestKey `plist:"pfm_subkeys"`
	Targets          []string      `plist:"pfm_targets"`
	Title            string        `plist:"pfm_title"`
	Unique           bool          `plist:"pfm_unique"`
	Version          int           `plist:"pfm_version"`
}

// ManifestKey represents a key definition in the manifest
//
//goland:noinspection ALL
type ManifestKey struct {
	Default          interface{}   `plist:"pfm_default,omitempty"`
	Description      string        `plist:"pfm_description"`
	Name             string        `plist:"pfm_name"`
	Type             string        `plist:"pfm_type"`
	Title            string        `plist:"pfm_title"`
	RangeList        []interface{} `plist:"pfm_range_list,omitempty"`
	RangeListTitles  []string      `plist:"pfm_range_list_titles,omitempty"`
	Required         string        `plist:"pfm_require,omitempty"`
	MacOSMin         string        `plist:"pfm_macos_min,omitempty"`
	Hidden           string        `plist:"pfm_hidden,omitempty"`
	Subkeys          []ManifestKey `plist:"pfm_subkeys,omitempty"`
	ValueProcessor   string        `plist:"pfm_value_processor,omitempty"`
	Format           string        `plist:"pfm_format,omitempty"`
	RangeMin         int           `plist:"pfm_range_min,omitempty"`
	RangeMax         int           `plist:"pfm_range_max,omitempty"`
	AllowedFileTypes []string      `plist:"pfm_allowed_file_types,omitempty"`
}

// LoadPreferenceManifestFromFile loads a preference manifest from a file path
func LoadPreferenceManifestFromFile(path string) (m *PreferenceManifest, err error) {
	file, err := os.Open(path)
	if err != nil {
		err = errors.Join(ErrFailedToOpenManifestFile, err)
		goto end
	}
	defer mustClose(file)
	m, err = LoadPreferenceManifest(file)
end:
	return m, err
}

// LoadPreferenceManifest loads a preference manifest from an io.Reader
func LoadPreferenceManifest(r io.Reader) (_ *PreferenceManifest, err error) {
	decoder := plist.NewXMLDecoder(r)
	m := PreferenceManifest{}
	err = decoder.Decode(&m)
	if err != nil {
		err = errors.Join(ErrFailedToDecodeManifest, err)
		goto end
	}
end:
	return &m, err
}

// SavePreferenceManifestFile saves a preference manifest to a file path
func (m *PreferenceManifest) SavePreferenceManifestFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		err = errors.Join(ErrFailedToCreateManifestFile, err)
		goto end
	}
	defer mustClose(file)
	err = m.SavePreferenceManifest(file)
	if err != nil {
		err = errors.Join(ErrFailedToSaveManifestFile, err)
		goto end
	}
end:
	return err
}

// SavePreferenceManifest writes a preference manifest to an io.Writer
func (m *PreferenceManifest) SavePreferenceManifest(w io.Writer) error {
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
func (m *PreferenceManifest) GetKey(name string) (mk *ManifestKey, err error) {
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
func (m *PreferenceManifest) Validate() (err error) {
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
func (m *PreferenceManifest) GetBoolKey(name string) (key *ManifestKey, defaultValue bool, err error) {
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
func (m *PreferenceManifest) GetIntKey(name string) (key *ManifestKey, defaultValue int, err error) {
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
func (m *PreferenceManifest) GetStringKey(name string) (key *ManifestKey, defaultValue string, err error) {
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
