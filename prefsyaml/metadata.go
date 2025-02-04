package prefsyaml

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type Metadata struct {
	Domain      DomainName              `yaml:"domain"`
	Description string                  `yaml:"description,omitempty"`
	Process     macosutil.ProcessName   `yaml:"process,omitempty"`
	Added       macosutil.VersionNumber `yaml:"added,omitempty"`
	Removed     macosutil.VersionNumber `yaml:"removed,omitempty"`

	// only needed until we create an interface for Spec and can use polymorphism for
	// Prefs and Defaults in Spec.Items vs. the separate slices we currently have.
	KindName KindName `yaml:"-"`
}
