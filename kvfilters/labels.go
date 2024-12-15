package kvfilters

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

type LabelName Name

func (name LabelName) String() string {
	return string(name)
}

type LabelValue string

func (value LabelValue) String() string {
	return string(value)
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

func (l Label) String() string {
	return fmt.Sprintf("%s=%s", l.Name, l.Value)
}

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

func (ll *Labels) Equivalent(labels *Labels) (eq bool) {
	if len(ll.labels) != len(labels.labels) {
		goto end
	}
	for _, label := range ll.labels {
		l, eq := labels.labelsMap[label.Name]
		if !eq {
			goto end
		}
		if l != label {
			goto end
		}
	}
	eq = true
end:
	return eq
}

func (ll *Labels) ContainsAny(labels ...*Label) (contains bool) {
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

func (ll *Labels) Valid() (ok bool) {
	var label *Label
	var index int

	if len(ll.labels) != len(ll.labelsMap) {
		goto end
	}
	if len(ll.labels) != len(ll.labelsIndex) {
		goto end
	}
	for _, l := range ll.labels {
		label, ok = ll.labelsMap[l.Name]
		if !ok {
			goto end
		}
		if l != label {
			goto end
		}
		index, ok = ll.labelsIndex[l.Name]
		if !ok {
			goto end
		}
		if l != ll.labels[index] {
			goto end
		}
	}
	ok = true
end:
	return ok
}

func (ll *Labels) SetLabel(label *Label) {
	count := len(ll.labelsMap)
	ll.labelsMap[label.Name] = label
	if count == len(ll.labelsMap) {
		// No new name was added, just update the label in the slice
		ll.labels[ll.labelsIndex[label.Name]] = label
		goto end
	}
	// A new name was added, update the index then add the label to the slice
	ll.labelsIndex[label.Name] = len(ll.labels)
	ll.labels = append(ll.labels, label)
end:
}

// HasLabel returns true if `*Labels` has the passed label
func (ll *Labels) HasLabel(label *Label) bool {
	return label == ll.GetNamedLabel(label.Name)
}

// HasNamedLabel returns true if `*Labels` has a label with the passed label name
func (ll *Labels) HasNamedLabel(name LabelName) bool {
	return ll.GetNamedLabel(name) != nil
}

// DeleteNamedLabel removes the label that has the passed label name.
func (ll *Labels) DeleteNamedLabel(name LabelName) {
	delete(ll.labelsMap, name)
	if index, ok := ll.labelsIndex[name]; ok {
		stdlibex.RemoveElements(ll.labels, index, 1)
		delete(ll.labelsIndex, name)
	}
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