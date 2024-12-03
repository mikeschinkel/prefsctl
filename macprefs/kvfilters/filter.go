package kvfilters

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
)

type Filter interface {
	Effect() Effect
	Target() Target
	Comparison() Comparison
	Matches() Matches
	IgnoreCase() bool
}
type ValuesGetter interface {
	Values() []string
}
type Funcs1Getter interface {
	Funcs1() []MatchFunc1
}
type Funcs2Getter interface {
	Funcs2() []MatchFunc2
}

// MustKeepGroup returns true is a value matches ANY Keep kvfilters. Rarely
// use this as it overrides all Omit kvfilters.
func MustKeepGroup(ff []Filter, group Name, targets ...Target) (keep bool, err error) {
	for _, filter := range ff {
		if filter.Effect() != Keep {
			continue
		}
		if !slices.Contains(targets, filter.Target()) {
			continue
		}
		keep, err = matchValueFilter(filter, string(group))
		if err != nil {
			goto end
		}
		if keep {
			goto end
		}
	}
end:
	return keep, err
}

// OmitGroup returns true if there is a match for ANY Omit kvfilters.
func OmitGroup(ff []Filter, group Name, targets ...Target) (omit bool, err error) {
	for _, filter := range ff {
		if filter.Effect() != Omit {
			continue
		}
		if !slices.Contains(targets, filter.Target()) {
			continue
		}
		omit, err = matchValueFilter(filter, string(group))
		if err != nil {
			goto end
		}
		if omit {
			goto end
		}
	}
end:
	return omit, err
}

//
//// KeepKeyValue1 returns true is a value matches ANY Keep kvfilters. Rarely
//// use this as it overrides all Omit kvfilters.
//func KeepKeyValue1(ff []Filter, keyValue Code, targets ...Target) (keep bool, err error) {
//	return KeepGroup(ff, keyValue, targets...)
//}
//
//// OmitKeyValue1 returns true if there is a match for ANY Omit kvfilters.
//func OmitKeyValue1(ff []Filter, keyValue Code, targets ...Target) (keep bool, err error) {
//	return KeepGroup(ff, keyValue, targets...)
//}

// MustKeepKeyValue2 returns true is a key or value as appropriate matches ANY Keep kvfilters.
func MustKeepKeyValue2(ff []Filter, key Code, value string, targets ...Target) (keep bool, err error) {
	for _, filter := range ff {
		if filter.Effect() != Keep {
			continue
		}
		if !slices.Contains(targets, filter.Target()) {
			continue
		}
		keep, err = matchKeyValue(filter, key, value)
		if err != nil {
			goto end
		}
		if keep {
			goto end
		}
	}
end:
	return keep, err
}

// OmitKeyValue2 returns true is a key or value as appropriate matches ANY Omit kvfilters.
func OmitKeyValue2(ff []Filter, key Code, value string, targets ...Target) (omit bool, err error) {
	for _, filter := range ff {
		if filter.Effect() != Omit {
			continue
		}
		if !slices.Contains(targets, filter.Target()) {
			continue
		}
		omit, err = matchKeyValue(filter, key, value)
		if err != nil {
			goto end
		}
		if omit {
			goto end
		}
	}
end:
	return omit, err
}

func matchValueFilter(filter Filter, value string) (match bool, err error) {
	if filter.IgnoreCase() {
		value = strings.ToLower(value)
	}
	matchCriteria := filter.Matches()
	switch matchCriteria {
	case Contains, Prefix, Suffix, Entirety, Regexp:
		getter, ok := filter.(ValuesGetter)
		if !ok {
			err = errors.Join(ErrFailedToAccessGroupKeyValues,
				fmt.Errorf("%s=kvfilters.ValuesGetter", logargs.TypeExpectedLogArg),
				err,
			)
			goto end
		}
		switch matchCriteria {
		case Contains:
			match = slices.Contains(getter.Values(), value)
		default:
			for _, specimen := range getter.Values() {
				//goland:noinspection GoDfaConstantCondition
				switch matchCriteria {
				case Prefix:
					match = strings.HasPrefix(value, specimen)
				case Suffix:
					match = strings.HasSuffix(value, specimen)
				case Entirety:
					match = value == specimen
				case Regexp:
					var re *regexp.Regexp
					reString := specimen
					re, err = regexp.Compile(reString)
					if err != nil {
						err = errors.Join(ErrInvalidRegExp,
							fmt.Errorf("%s=%s", logargs.RegexpLogArg, reString),
							err,
						)
						goto end
					}
					match = re.MatchString(value)
				}
				if match {
					break
				}
			}
		}
	case Func:
		getter, ok := filter.(Funcs1Getter)
		if !ok {
			err = errors.Join(ErrFailedTypeAssertion,
				fmt.Errorf("%s=kvfilters.Funcs1Getter", logargs.TypeExpectedLogArg),
				err,
			)
			goto end
		}
		for _, gf := range getter.Funcs1() {
			match = gf(Code(value))
			if match {
				goto end
			}
		}
	case InvalidMatches:
		//goland:noinspection GoDfaNilDereference
		err = errors.Join(ErrInvalidFilterMatchCriteria,
			fmt.Errorf("%s=%s", logargs.TypeExpectedLogArg, matchCriteria.String()),
			err,
		)
		goto end
	}

	if filter.Comparison() == NotEquals {
		match = !match
	}
end:
	return match, err
}

func matchKeyValueFilter(filter Filter, key Code, value string) (match bool, err error) {
	if filter.IgnoreCase() {
		key = Code(strings.ToLower(string(key)))
		value = strings.ToLower(value)
	}
	getter, ok := filter.(Funcs2Getter)
	if !ok {
		err = errors.Join(ErrFailedTypeAssertion,
			fmt.Errorf("%s=kvfilters.Funcs2Getter", logargs.TypeExpectedLogArg),
			err,
		)
	}
	for _, gf := range getter.Funcs2() {
		match = gf(key, value)
		if match {
			goto end
		}
	}
end:
	if filter.Comparison() == NotEquals {
		match = !match
	}
	return match, err
}

func matchKeyValue(f Filter, key Code, value string) (match bool, err error) {
	var arg Code
	switch f.Target() {
	case Keys:
		arg = key
	case Values:
		arg = Code(value)
	case KeyValues:
		arg = Code(fmt.Sprintf("%s=%s", key, value))
	default:
		match, err = matchKeyValueFilter(f, key, value)
		if err != nil {
			goto end
		}
		goto end
	}
	match, err = matchValueFilter(f, string(arg))
	if err != nil {
		goto end
	}
end:
	return match, err
}
