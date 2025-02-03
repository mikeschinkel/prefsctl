package prefsyaml

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type Metadata struct {
	Domain      DomainName              `yaml:"domain"`
	Description Description             `yaml:"description,omitempty"`
	Process     macosutil.ProcessName   `yaml:"process,omitempty"`
	Added       macosutil.VersionNumber `yaml:"added,omitempty"`
	Removed     macosutil.VersionNumber `yaml:"removed,omitempty"`
}
