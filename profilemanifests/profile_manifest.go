package profilemanifests

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/micromdm/plist"
	"github.com/mikeschinkel/prefsctl/appinfo"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macpreflabels"
	"github.com/mikeschinkel/prefsctl/prefsyaml"
	"github.com/mikeschinkel/prefsctl/yamlutil"
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

func (pm *ProfileManifest) FilterableEntry() {}

// SaveProfileManifestFile saves a preference manifest to a file path
func (pm *ProfileManifest) SaveProfileManifestFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		err = errors.Join(ErrFailedToCreateManifestFile, err)
		goto end
	}
	defer mustClose(file)
	err = pm.SaveProfileManifest(file)
	if err != nil {
		err = errors.Join(ErrFailedToSaveManifestFile, err)
		goto end
	}
end:
	return err
}

// SaveProfileManifest writes a preference manifest to an io.Writer
func (pm *ProfileManifest) SaveProfileManifest(w io.Writer) error {
	encoder := plist.NewEncoder(w)
	encoder.Indent("    ")
	err := encoder.Encode(pm)
	if err != nil {
		err = errors.Join(ErrFailedToEncodeManifest, err)
		goto end
	}
end:
	return err
}

// GetKey searches for a key by name in the manifest's sub-keys
func (pm *ProfileManifest) GetKey(name string) (mk *ManifestKey, err error) {
	for _, key := range pm.Subkeys {
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
func (pm *ProfileManifest) Validate() (err error) {
	if pm.Domain == "" {
		err = ErrMissingPreferenceDomain
		goto end
	}
	if pm.FormatVersion <= 0 {
		err = errors.Join(
			ErrInvalidManifestFormatVersion,
			fmt.Errorf("%s=%d", FormatVersionLogArg, pm.FormatVersion),
		)
		goto end
	}
end:
	return err
}

// LoadProfileManifestFromFile loads a preference manifest from a file path
func LoadProfileManifestFromFile(path string) (pm *ProfileManifest, err error) {
	file, err := os.Open(path)
	if err != nil {
		err = errors.Join(ErrFailedToOpenManifestFile, err)
		goto end
	}
	defer mustClose(file)
	pm, err = ReadProfileManifest(file)
end:
	return pm, err
}

// ReadProfileManifest loads a preference manifest from an io.Reader
func ReadProfileManifest(r io.Reader) (_ *ProfileManifest, err error) {
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

// Options return ManifestKey.RangeList as slice of *Value
func (mk *ManifestKey) Options() (values []*Value) {
	values = make([]*Value, len(mk.RangeList))
	for i, item := range mk.RangeList {
		values[i] = NewValue(item.Value)
	}
	return values
}

// GetPrefsYAMLResource returns a prefsyaml.Resource filtered by an
// ObjectFilterFunc. By definition, this can only be a Defaults resource, it
// cannot be a Prefs resource.
func (pm *ProfileManifest) GetPrefsYAMLResource(includeFilter yamlutil.EntryFilterFunc) (yr prefsyaml.Resource, include bool) {
	yr = prefsyaml.NewResource(
		prefsyaml.KindName(macpreflabels.DefaultsKind),
		prefsyaml.DomainName(pm.Domain),
		prefsyaml.ResourceOpts{
			APIVersion:  appinfo.LatestAPIVersion,
			Description: prefsyaml.Description(pm.Description),
			Added:       macosutil.VersionNumber(pm.MacOSMin),
			Removed:     macosutil.VersionNumber(pm.MacOSMax),
			Process:     macosutil.GetProcessToKill(macosutil.DomainName(pm.Domain)),
		},
	)
	if !includeFilter(yr) {
		goto end
	}
	include = true
	for _, sk := range pm.Subkeys {
		var value *yamlutil.Value
		if sk.Default != nil {
			value = sk.Default.Value
		}
		def := prefsyaml.NewDefault(prefsyaml.PrefName(sk.Name), prefsyaml.DefaultOpts{
			Value:       value,
			Description: prefsyaml.Description(sk.Description),
			Options:     ExtractYAMLUtilValues(sk.Options()),
			Type:        prefsyaml.PrefType(sk.Type),
			Added:       macosutil.VersionNumber(sk.MacOSMin),
			Removed:     macosutil.VersionNumber(sk.MacOSMax),
			Labels:      nil, // TODO Find a way to derive labels
		})
		if !includeFilter(def) {
			continue
		}
		yr.Spec.Defaults = append(yr.Spec.Defaults, def)
	}
end:
	return yr, include
}

//func (pm *ProfileManifest) YAMLDocument() (yd yamlutil.Document, err error) {
//	var yml []byte
//	yml, err = yaml.Marshal(pm)
//	if err != nil {
//		goto end
//	}
//	yd = yamlutil.Document(yml)
//end:
//	return yd, err
//}
