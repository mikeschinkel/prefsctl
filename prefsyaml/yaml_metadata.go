package prefsyaml

import (
	"github.com/mikeschinkel/prefsctl/macosutil"
)

type YAMLMetadata struct {
	Domain      DomainName            `yaml:"domain"`
	KillOnApply macosutil.ProcessName `yaml:"kill_on_apply,omitempty"`
}
