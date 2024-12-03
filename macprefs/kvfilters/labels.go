package kvfilters

import (
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type LabelName macosutils.Name
type LabelValue string

type Labels []*Label

func (ll Labels) String() string {
	s, _ := sliceconv.ToStringsFunc(ll, func(label *Label) (bool, string, error) {
		return true, label.String(), nil
	})
	return strings.Join(s, ",")
}

func (l Label) String() string {
	return fmt.Sprintf("%s=%s", l.Name, l.Value)
}

type Label struct {
	Name  LabelName
	Value LabelValue
}
