package kvfilters

//goland:noinspection SpellCheckingInspection
func SequoiaQueryFilters() []Filter {
	filters := []Filter{
		KeepWhenKeyIsOneOf{},
		OmitWhenGroupIsOneOf{},
		OmitWhenGroupSuffixIsOneOfIgnoringCase{},
		OmitWhenKeySuffixIsOneOfIgnoringCase{},
		OmitWhenKeyPrefixIsOneOf{},
		OmitWhenGroupContainsOneOf{},
		OmitWhenKeyContainsOneOfIgnoringCase{},
		OmitWhenGroupPrefixIsNotOneOf{},
		OmitWhenGroupIsOneOf{},
		OmitWhenKeyPrefixIsOneOf{},
		OmitWhenKeyIsOneOf{},
		OmitWhenKeyPrefixIsOneOf{},
		OmitWhenValueMatchedByOneOfRegexps{},
		OmitWhenValueMatchedByOneOfRegexpsIgnoringCase{},
		OmitWhenGroupPassedToFuncReturnsFalse{},
	}
	return append(
		MontereyQueryFilters(),
		filters...,
	)
}
