package kvfilters

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mikeschinkel/prefsctl/macosutils"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type LabelName macosutils.Name

func (name LabelName) String() string {
	return string(name)
}
func (value LabelValue) String() string {
	return string(value)
}

type LabelValue string

type Labels []*Label

func (ll Labels) Contains(labels ...*Label) (contains bool) {
	for _, label := range labels {
		if !slices.Contains(ll, label) {
			continue
		}
		contains = true
		goto end
	}
end:
	return contains
}

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
