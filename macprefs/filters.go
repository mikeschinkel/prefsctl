package macprefs

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type FilterType byte

const (
	OmitFilterType     FilterType = 1
	MustKeepFilterType FilterType = 2
)

type FilterTarget byte

// Numeric values of FilterString or less are for string filters (e.g. domains)
// Numeric values != FilterString  are for key-value filters (e.g. prefs)
const (
	FilterAll      FilterTarget = 1
	FilterString   FilterTarget = 2
	FilterKey      FilterTarget = 3
	FilterValue    FilterTarget = 4
	FilterKeyValue FilterTarget = 5
)

type KeyValueFilters []Filter

// MustKeepString returns true is a domain matches ANY MustKeep filters. Rarely
// use this as it overrides all Omit filters.
func (ff KeyValueFilters) MustKeepString(domain string) (keep bool) {
	for _, filter := range ff {
		fb := filter.Base()
		if fb.Target > FilterString {
			continue
		}
		if fb.Type != MustKeepFilterType {
			continue
		}
		if match1ParamFilter(filter, domain) {
			keep = true
			goto end
		}
	}
end:
	return keep
}

// MustKeepKeyValue returns true is a key or value as appropriate matches ANY MustKeep filters.
func (ff KeyValueFilters) MustKeepKeyValue(k, v string) (keep bool) {
	for _, filter := range ff {
		fb := filter.Base()
		if fb.Target == FilterString {
			continue
		}
		// fb.Target == FilterKeyValue OR fb.Target == FilterAll
		if fb.Type != MustKeepFilterType {
			continue
		}
		if matchKeyValue(filter, k, v) {
			keep = true
			goto end
		}
	}
end:
	return keep
}

// OmitDomain returns true is a domain matches ANY Omit filters.
func (ff KeyValueFilters) OmitDomain(domain string) (omit bool) {
	for _, filter := range ff {
		fb := filter.Base()
		if fb.Target > FilterString {
			continue
		}
		if fb.Type != OmitFilterType {
			continue
		}
		if match1ParamFilter(filter, domain) {
			omit = true
			goto end
		}
	}
end:
	return omit
}

// OmitKeyValue returns true is a key or value as appropriate matches ANY Omit filters.
func (ff KeyValueFilters) OmitKeyValue(k, v string) (omit bool) {
	for _, filter := range ff {
		fb := filter.Base()
		if fb.Target == FilterString {
			continue
		}
		// fb.Target == FilterKeyValue OR fb.Target == FilterAll
		if fb.Type != OmitFilterType {
			continue
		}
		if matchKeyValue(filter, k, v) {
			omit = true
			goto end
		}
	}
end:
	return omit
}

func matchKeyValue(f Filter, k, v string) (match bool) {
	fb := f.Base()
	var arg string
	switch fb.Target {
	case FilterKey:
		arg = k
	case FilterValue:
		arg = v
	case FilterKeyValue, FilterAll:
		arg = fmt.Sprintf("%s=%s", k, v)
	default:
		if match2ParamFilter(f, k, v) {
			match = true
		}
		goto end
	}
	if match1ParamFilter(f, arg) {
		match = true
		goto end
	}
end:
	return match
}

func match2ParamFilter(filter Filter, key, value string) (match bool) {
	switch t := filter.(type) {
	case KeyValueFuncFilter:
		match = t.Filter(key, value)
	default:
		match = match1ParamFilter(filter, fmt.Sprintf("%s=%s", key, value))
	}
	return match
}

func match1ParamFilter(filter Filter, s string) (match bool) {
	switch t := filter.(type) {
	case RegexFilter:
		match = t.Filter.MatchString(s)
	case EqualsFilter:
		match = s == t.Filter
	case DoesNotEqualFilter:
		match = s != t.Filter
	case HasPrefixFilter:
		match = strings.HasPrefix(s, t.Filter)
	case DoesNotHavePrefixFilter:
		match = !strings.HasPrefix(s, t.Filter)
	case HasSuffixFilter:
		match = strings.HasSuffix(s, t.Filter)
	case DoesNotHaveSuffixFilter:
		match = !strings.HasSuffix(s, t.Filter)
	case ContainsFilter:
		match = slices.Contains(t.Filter, s)
	case DoesNotContainFilter:
		match = !slices.Contains(t.Filter, s)
	case DomainFuncFilter:
		match = t.Filter(s)
	}
	return match
}

type Filter interface {
	IsFilter()
	Base() FilterBase
}
type FilterBase struct {
	Type   FilterType
	Target FilterTarget
}

func (FilterBase) IsFilter() {}
func (fb FilterBase) Base() FilterBase {
	return fb
}

var _ Filter = (*RegexFilter)(nil)

type RegexFilter struct {
	FilterBase
	Filter regexp.Regexp
}
type EqualsFilter struct {
	FilterBase
	Filter string
}
type DoesNotEqualFilter struct {
	FilterBase
	Filter string
}
type HasPrefixFilter struct {
	FilterBase
	Filter string
}
type DoesNotHavePrefixFilter struct {
	FilterBase
	Filter string
}
type HasSuffixFilter struct {
	FilterBase
	Filter string
}
type DoesNotHaveSuffixFilter struct {
	FilterBase
	Filter string
}
type ContainsFilter struct {
	FilterBase
	Filter []string
}
type DoesNotContainFilter struct {
	FilterBase
	Filter []string
}
type DomainFuncFilter struct {
	FilterBase
	Filter func(domain string) bool
}
type KeyValueFuncFilter struct {
	FilterBase
	Filter func(name, value string) bool
}
