package kvfilters

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/mikeschinkel/prefsctl/sliceconv"
)

type LabelName Name

func (name LabelName) String() string {
	return string(name)
}
func (value LabelValue) String() string {
	return string(value)
}

type LabelValue string

type Labels struct {
	labels      []*Label
	labelsMap   map[LabelName]*Label
	labelsIndex map[LabelName]int
}

func NewLabels(labels ...*Label) *Labels {
	ll := &Labels{
		labels: labels,
	}
	ll.SyncMap()
	return ll
}

func (ll *Labels) SyncMap() {
	ll.labelsMap = make(map[LabelName]*Label, len(ll.labels))
	ll.labelsIndex = make(map[LabelName]int, len(ll.labels))
	for index, label := range ll.labels {
		ll.labelsMap[label.Name] = label
		ll.labelsIndex[label.Name] = index
	}
}

func (ll *Labels) LabelPtrs() []*Label {
	labels := slices.Clone(ll.labels)
	slices.SortFunc(labels, func(a, b *Label) int {
		return cmp.Compare(a.Value, b.Value)
	})
	return labels
}

func (ll *Labels) Contains(labels ...*Label) (contains bool) {
	for _, label := range labels {
		if !slices.Contains(ll.labels, label) {
			continue
		}
		contains = true
		goto end
	}
end:
	return contains
}

func (ll *Labels) SetNamedLabel(label *Label) {
	ll.labelsMap[label.Name] = label
	ll.labels = make([]*Label, len(ll.labelsMap))
	ll.labelsIndex = make(map[LabelName]int, len(ll.labelsMap))
	index := 0
	for name, l := range ll.labelsMap {
		ll.labels[index] = l
		ll.labelsIndex[name] = index
		index++
	}
}

// Add appends one or more `*Label` objects to `.labels` with the assumption that
// each `label.Name` will be unique among labels, e.g. only one value for each of
// 'type', 'class' and 'sets'.
func (ll *Labels) Add(labels ...*Label) {
	for _, label := range labels {
		l := ll.GetNamedLabel(label.Name)
		if l == nil || l.IsUnknown {
			ll.SetNamedLabel(label)
			continue
		}
		if ll.HasNamedLabel(label) {
			continue
		}
		ll.labels = append(ll.labels, label)
	}
}

// HasNamedLabel returns true if `*Labels` has a label with the passed label name
func (ll *Labels) HasNamedLabel(label *Label) bool {
	return label == ll.GetNamedLabel(label.Name)
}

func (ll *Labels) GetNamedLabel(name LabelName) *Label {
	label, _ := ll.labelsMap[name]
	return label
}

func (ll *Labels) String() string {
	s, _ := sliceconv.ToStringsFunc(ll.labels, func(label *Label) (bool, string, error) {
		return true, label.String(), nil
	})
	return strings.Join(s, ",")
}

func (l Label) String() string {
	return fmt.Sprintf("%s=%s", l.Name, l.Value)
}

type Label struct {
	Name      LabelName
	Value     LabelValue
	IsUnknown bool
}

func NewLabel(name LabelName, value LabelValue) Label {
	return Label{Name: name, Value: value}
}
func NewUnknownLabel(name LabelName, value LabelValue) Label {
	l := NewLabel(name, value)
	l.IsUnknown = true
	return l
}
