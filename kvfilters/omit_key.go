package kvfilters

//=============================================

type OmitWhenKeyIsOneOf []string
type ok = OmitWhenKeyIsOneOf

var _ Filter = (*ok)(nil)

func (f ok) Values() []string {
	return f
}
func (ok) Effect() Effect {
	return Omit
}
func (ok) Target() Target {
	return Keys
}
func (ok) Matches() Matches {
	return Entirety
}
func (ok) Comparison() Comparison {
	return Equals
}
func (ok) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenKeyIsOneOfIgnoringCase []string
type okic = OmitWhenKeyIsOneOfIgnoringCase

var _ Filter = (*okic)(nil)

func (f okic) Values() []string {
	return f
}
func (okic) Effect() Effect {
	return Omit
}
func (okic) Target() Target {
	return Keys
}
func (okic) Matches() Matches {
	return Entirety
}
func (okic) Comparison() Comparison {
	return Equals
}
func (okic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenKeyPrefixIsOneOf []string
type okp = OmitWhenKeyPrefixIsOneOf

var _ Filter = (*okp)(nil)

func (f okp) Values() []string {
	return f
}
func (okp) Effect() Effect {
	return Omit
}
func (okp) Target() Target {
	return Keys
}
func (okp) Matches() Matches {
	return Prefix
}
func (okp) Comparison() Comparison {
	return Equals
}
func (okp) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenKeyPrefixIsOneOfIgnoringCase []string
type okpic = OmitWhenKeyPrefixIsOneOfIgnoringCase

var _ Filter = (*okpic)(nil)

func (f okpic) Values() []string {
	return f
}
func (okpic) Effect() Effect {
	return Omit
}
func (okpic) Target() Target {
	return Keys
}
func (okpic) Matches() Matches {
	return Prefix
}
func (okpic) Comparison() Comparison {
	return Equals
}
func (okpic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenKeyPrefixIsNotOneOf []string
type okpn = OmitWhenKeyPrefixIsNotOneOf

var _ Filter = (*okpn)(nil)

func (f okpn) Values() []string {
	return f
}
func (okpn) Effect() Effect {
	return Omit
}
func (okpn) Target() Target {
	return Keys
}
func (okpn) Matches() Matches {
	return Prefix
}
func (okpn) Comparison() Comparison {
	return NotEquals
}
func (okpn) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenKeyPrefixIsNotOneOfIgnoringCase []string
type okpnic = OmitWhenKeyPrefixIsNotOneOfIgnoringCase

var _ Filter = (*okpnic)(nil)

func (f okpnic) Values() []string {
	return f
}
func (okpnic) Effect() Effect {
	return Omit
}
func (okpnic) Target() Target {
	return Keys
}
func (okpnic) Matches() Matches {
	return Prefix
}
func (okpnic) Comparison() Comparison {
	return NotEquals
}
func (okpnic) IgnoreCase() bool {
	return true
}

//=============================================

var _ Filter = (*oks)(nil)

type OmitWhenKeySuffixIsOneOf []string
type oks = OmitWhenKeySuffixIsOneOf

func (f oks) Values() []string {
	return f
}
func (oks) Effect() Effect {
	return Omit
}
func (oks) Target() Target {
	return Keys
}
func (oks) Matches() Matches {
	return Suffix
}
func (oks) Comparison() Comparison {
	return Equals
}
func (oks) IgnoreCase() bool {
	return true
}

//=============================================

var _ Filter = (*oksic)(nil)

type OmitWhenKeySuffixIsOneOfIgnoringCase []string
type oksic = OmitWhenKeySuffixIsOneOfIgnoringCase

func (f oksic) Values() []string {
	return f
}
func (oksic) Effect() Effect {
	return Omit
}
func (oksic) Target() Target {
	return Keys
}
func (oksic) Matches() Matches {
	return Suffix
}
func (oksic) Comparison() Comparison {
	return Equals
}
func (oksic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenKeySuffixIsNotOneOf []string
type oksn = OmitWhenKeySuffixIsNotOneOf

var _ Filter = (*oksn)(nil)

func (f oksn) Values() []string {
	return f
}
func (oksn) Effect() Effect {
	return Omit
}
func (oksn) Target() Target {
	return Keys
}
func (oksn) Matches() Matches {
	return Suffix
}
func (oksn) Comparison() Comparison {
	return NotEquals
}
func (oksn) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenKeySuffixIsNotOneOfIgnoringCase []string
type oksnic = OmitWhenKeySuffixIsNotOneOfIgnoringCase

var _ Filter = (*oksnic)(nil)

func (f oksnic) Values() []string {
	return f
}
func (oksnic) Effect() Effect {
	return Omit
}
func (oksnic) Target() Target {
	return Keys
}
func (oksnic) Matches() Matches {
	return Suffix
}
func (oksnic) Comparison() Comparison {
	return NotEquals
}
func (oksnic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenKeyContainsOneOf []string
type okc = OmitWhenKeyContainsOneOf

var _ Filter = (*okc)(nil)

func (f okc) Values() []string {
	return f
}
func (okc) Effect() Effect {
	return Omit
}
func (okc) Target() Target {
	return Keys
}
func (okc) Matches() Matches {
	return Contains
}
func (okc) Comparison() Comparison {
	return Equals
}
func (okc) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenKeyContainsOneOfIgnoringCase []string
type okcic = OmitWhenKeyContainsOneOfIgnoringCase

var _ Filter = (*okcic)(nil)

func (f okcic) Values() []string {
	return f
}
func (okcic) Effect() Effect {
	return Omit
}
func (okcic) Target() Target {
	return Keys
}
func (okcic) Matches() Matches {
	return Contains
}
func (okcic) Comparison() Comparison {
	return Equals
}
func (okcic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenKeyMatchedByRegexp []string
type okr = OmitWhenKeyMatchedByRegexp

var _ Filter = (*okr)(nil)

func (f okr) Values() []string {
	return f
}
func (okr) Effect() Effect {
	return Omit
}
func (okr) Target() Target {
	return Keys
}
func (okr) Comparison() Comparison {
	return Equals
}
func (okr) Matches() Matches {
	return Prefix
}
func (okr) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenKeyMatchedByRegexpIgnoringCase []string
type okric = OmitWhenKeyMatchedByRegexpIgnoringCase

var _ Filter = (*okric)(nil)

func (f okric) Values() []string {
	return f
}
func (okric) Effect() Effect {
	return Omit
}
func (okric) Target() Target {
	return Keys
}
func (okric) Comparison() Comparison {
	return Equals
}
func (okric) Matches() Matches {
	return Regexp
}
func (okric) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenKeyPassedToFuncReturnsTrue []MatchFunc1
type okft = OmitWhenKeyPassedToFuncReturnsTrue

var _ Filter = (*okft)(nil)

func (f okft) Funcs1() []MatchFunc1 {
	return f
}
func (okft) Effect() Effect {
	return Omit
}
func (okft) Target() Target {
	return Keys
}
func (okft) Comparison() Comparison {
	return IsTrue
}
func (okft) Matches() Matches {
	return Func
}
func (okft) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenKeyPassedToFuncReturnsTrueIgnoringCase []MatchFunc1
type okftic = OmitWhenKeyPassedToFuncReturnsTrueIgnoringCase

var _ Filter = (*okftic)(nil)

func (f okftic) Funcs1() []MatchFunc1 {
	return f
}
func (okftic) Effect() Effect {
	return Omit
}
func (okftic) Target() Target {
	return Keys
}
func (okftic) Comparison() Comparison {
	return IsTrue
}
func (okftic) Matches() Matches {
	return Func
}
func (okftic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenKeyPassedToFuncReturnsFalse []MatchFunc1
type okff = OmitWhenKeyPassedToFuncReturnsFalse

var _ Filter = (*okff)(nil)

func (f okff) Funcs1() []MatchFunc1 {
	return f
}
func (okff) Effect() Effect {
	return Omit
}
func (okff) Target() Target {
	return Keys
}
func (okff) Comparison() Comparison {
	return IsFalse
}
func (okff) Matches() Matches {
	return Func
}
func (okff) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenKeyPassedToFuncReturnsFalseIgnoringCase []MatchFunc1
type okffic = OmitWhenKeyPassedToFuncReturnsFalseIgnoringCase

var _ Filter = (*okffic)(nil)

func (f okffic) Funcs1() []MatchFunc1 {
	return f
}
func (okffic) Effect() Effect {
	return Omit
}
func (okffic) Target() Target {
	return Keys
}
func (okffic) Comparison() Comparison {
	return IsFalse
}
func (okffic) Matches() Matches {
	return Func
}
func (okffic) IgnoreCase() bool {
	return true
}
