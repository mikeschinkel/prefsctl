package prefsyaml

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type YAMLPref struct {
	MetaData      *YAMLMetadata `yaml:"-"`
	Name          PrefName      `yaml:"name"`
	Type          PrefType      `yaml:"type"`
	Value         YAMLValue     `yaml:"value,omitempty"`
	Default       YAMLValue     `yaml:"default,omitempty"`
	Labels        []LabelValue  `yaml:"labels,omitempty"`
	SupportedIn   OSVersion     `yaml:"supported_in,omitempty"`
	UnsupportedIn OSVersion     `yaml:"unsupported_in,omitempty"`
}

func (yp *YAMLPref) PreferenceType() (pt macosutil.PreferenceType) {
	return macosutil.PreferenceType(yp.Type)
}

func (yp *YAMLPref) MacOSUtilPreference() (pref *macosutil.Preference) {
	return &macosutil.Preference{
		Domain: string(yp.MetaData.Domain),
		Name:   string(yp.Name),
		Value:  yp.Value.String(), // TODO: This will need to be revised when we support more than scalar types
	}
}
