package filters

//=============================================

type OmitWhenValueMatchedByOneOfRegexps []string
type ovr = OmitWhenValueMatchedByOneOfRegexps

var _ Filter = (*ovr)(nil)

func (f ovr) Values() []string {
	return f
}
func (ovr) Effect() Effect {
	return Omit
}
func (ovr) Target() Target {
	return Values
}
func (ovr) Comparison() Comparison {
	return IsMatched
}
func (ovr) Matches() Matches {
	return Regexp
}
func (ovr) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenValueMatchedByOneOfRegexpsIgnoringCase []string
type ovric = OmitWhenValueMatchedByOneOfRegexpsIgnoringCase

var _ Filter = (*ovric)(nil)

func (f ovric) Values() []string {
	return f
}
func (ovric) Effect() Effect {
	return Omit
}
func (ovric) Target() Target {
	return Values
}
func (ovric) Comparison() Comparison {
	return IsMatched
}
func (ovric) Matches() Matches {
	return Regexp
}
func (ovric) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenValueNotMatchedByOneOfRegexps []string
type ovrn = OmitWhenValueNotMatchedByOneOfRegexps

var _ Filter = (*ovrn)(nil)

func (f ovrn) Values() []string {
	return f
}
func (ovrn) Effect() Effect {
	return Omit
}
func (ovrn) Target() Target {
	return Values
}
func (ovrn) Comparison() Comparison {
	return NotMatched
}
func (ovrn) Matches() Matches {
	return Regexp
}
func (ovrn) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenValueNotMatchedByOneOfRegexpsIgnoringCase []string
type ovrnic = OmitWhenValueNotMatchedByOneOfRegexpsIgnoringCase

var _ Filter = (*ovrnic)(nil)

func (f ovrnic) Values() []string {
	return f
}
func (ovrnic) Effect() Effect {
	return Omit
}
func (ovrnic) Target() Target {
	return Values
}
func (ovrnic) Comparison() Comparison {
	return IsMatched
}
func (ovrnic) Matches() Matches {
	return Regexp
}
func (ovrnic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenValuePassedToFuncReturnsTrue []MatchFunc1
type ovft = OmitWhenValuePassedToFuncReturnsTrue

var _ Filter = (*ovft)(nil)

func (f ovft) Funcs1() []MatchFunc1 {
	return f
}
func (ovft) Effect() Effect {
	return Omit
}
func (ovft) Target() Target {
	return Values
}
func (ovft) Comparison() Comparison {
	return IsTrue
}
func (ovft) Matches() Matches {
	return Func
}
func (ovft) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenValuePassedToFuncReturnsTrueIgnoringCase []MatchFunc1
type ovftic = OmitWhenValuePassedToFuncReturnsTrueIgnoringCase

var _ Filter = (*ovftic)(nil)

func (f ovftic) Funcs1() []MatchFunc1 {
	return f
}
func (ovftic) Effect() Effect {
	return Omit
}
func (ovftic) Target() Target {
	return Values
}
func (ovftic) Comparison() Comparison {
	return IsTrue
}
func (ovftic) Matches() Matches {
	return Func
}
func (ovftic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenValuePassedToFuncReturnsFalse []MatchFunc1
type ovff = OmitWhenValuePassedToFuncReturnsFalse

var _ Filter = (*ovff)(nil)

func (f ovff) Funcs1() []MatchFunc1 {
	return f
}
func (ovff) Effect() Effect {
	return Omit
}
func (ovff) Target() Target {
	return Values
}
func (ovff) Comparison() Comparison {
	return IsFalse
}
func (ovff) Matches() Matches {
	return Func
}
func (ovff) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenValuePassedToFuncReturnsFalseIgnoringCase []MatchFunc1
type ovffic = OmitWhenValuePassedToFuncReturnsFalseIgnoringCase

var _ Filter = (*ovffic)(nil)

func (f ovffic) Funcs1() []MatchFunc1 {
	return f
}
func (ovffic) Effect() Effect {
	return Omit
}
func (ovffic) Target() Target {
	return Values
}
func (ovffic) Comparison() Comparison {
	return IsFalse
}
func (ovffic) Matches() Matches {
	return Func
}
func (ovffic) IgnoreCase() bool {
	return true
}
