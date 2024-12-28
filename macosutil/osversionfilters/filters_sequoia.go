package osversionfilters

import (
	"github.com/mikeschinkel/prefsctl/kvfilters"
)

//goland:noinspection SpellCheckingInspection
func SequoiaQueryFilters() []kvfilters.Filter {
	filters := []kvfilters.Filter{
		kvfilters.KeepWhenKeyIsOneOf{},
		kvfilters.OmitWhenGroupIsOneOf{},
		kvfilters.OmitWhenGroupSuffixIsOneOfIgnoringCase{},
		kvfilters.OmitWhenKeySuffixIsOneOfIgnoringCase{},
		kvfilters.OmitWhenKeyPrefixIsOneOf{},
		kvfilters.OmitWhenGroupContainsOneOf{},
		kvfilters.OmitWhenKeyContainsOneOfIgnoringCase{},
		kvfilters.OmitWhenGroupPrefixIsNotOneOf{},
		kvfilters.OmitWhenGroupIsOneOf{},
		kvfilters.OmitWhenKeyPrefixIsOneOf{},
		kvfilters.OmitWhenKeyIsOneOf{},
		kvfilters.OmitWhenKeyPrefixIsOneOf{},
		kvfilters.OmitWhenValueMatchedByOneOfRegexps{},
		kvfilters.OmitWhenValueMatchedByOneOfRegexpsIgnoringCase{},
		kvfilters.OmitWhenGroupPassedToFuncReturnsFalse{},
	}
	return append(
		MontereyQueryFilters(),
		filters...,
	)
}
