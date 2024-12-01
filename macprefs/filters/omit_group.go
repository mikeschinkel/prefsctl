package filters

//=============================================

type OmitWhenGroupIsOneOf []string
type og = OmitWhenGroupIsOneOf

var _ Filter = (*og)(nil)

func (f og) Values() []string {
	return f
}
func (og) Effect() Effect {
	return Omit
}
func (og) Target() Target {
	return Groups
}
func (og) Comparison() Comparison {
	return Equals
}
func (og) Matches() Matches {
	return Entirety
}
func (og) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenGroupIsNotOneOf []string
type ogn = OmitWhenGroupIsNotOneOf

var _ Filter = (*ogn)(nil)

func (f ogn) Values() []string {
	return f
}
func (ogn) Effect() Effect {
	return Omit
}
func (ogn) Target() Target {
	return Groups
}
func (ogn) Comparison() Comparison {
	return NotEquals
}
func (ogn) Matches() Matches {
	return Entirety
}
func (ogn) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenGroupContainsOneOf []string

type ogc = OmitWhenGroupContainsOneOf

var _ Filter = (*ogc)(nil)

func (f ogc) Values() []string {
	return f
}
func (ogc) Effect() Effect {
	return Omit
}
func (ogc) Target() Target {
	return Groups
}
func (ogc) Comparison() Comparison {
	return Equals
}
func (ogc) Matches() Matches {
	return Contains
}
func (ogc) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenGroupContainsOneOfIgnoringCase []string

type ogcic = OmitWhenGroupContainsOneOfIgnoringCase

var _ Filter = (*ogcic)(nil)

func (f ogcic) Values() []string {
	return f
}
func (ogcic) Effect() Effect {
	return Omit
}
func (ogcic) Target() Target {
	return Groups
}
func (ogcic) Comparison() Comparison {
	return Equals
}
func (ogcic) Matches() Matches {
	return Contains
}
func (ogcic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenGroupPrefixIsOneOf []string
type ogp = OmitWhenGroupPrefixIsOneOf

var _ Filter = (*ogp)(nil)

func (f ogp) Values() []string {
	return f
}
func (ogp) Effect() Effect {
	return Omit
}
func (ogp) Target() Target {
	return Groups
}
func (ogp) Comparison() Comparison {
	return Equals
}
func (ogp) Matches() Matches {
	return Prefix
}
func (ogp) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenGroupPrefixIsOneOfIgnoringCase []string
type ogpic = OmitWhenGroupPrefixIsOneOfIgnoringCase

var _ Filter = (*ogpic)(nil)

func (f ogpic) Values() []string {
	return f
}
func (ogpic) Effect() Effect {
	return Omit
}
func (ogpic) Target() Target {
	return Groups
}
func (ogpic) Comparison() Comparison {
	return Equals
}
func (ogpic) Matches() Matches {
	return Prefix
}
func (ogpic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenGroupPrefixIsNotOneOf []string
type ogpn = OmitWhenGroupPrefixIsNotOneOf

var _ Filter = (*ogpn)(nil)

func (f ogpn) Values() []string {
	return f
}
func (ogpn) Effect() Effect {
	return Omit
}
func (ogpn) Target() Target {
	return Groups
}
func (ogpn) Comparison() Comparison {
	return NotEquals
}
func (ogpn) Matches() Matches {
	return Prefix
}
func (ogpn) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenGroupPrefixIsNotOneOfIgnoringCase []string
type ogpnic = OmitWhenGroupPrefixIsNotOneOfIgnoringCase

var _ Filter = (*ogpnic)(nil)

func (f ogpnic) Values() []string {
	return f
}
func (ogpnic) Effect() Effect {
	return Omit
}
func (ogpnic) Target() Target {
	return Groups
}
func (ogpnic) Comparison() Comparison {
	return NotEquals
}
func (ogpnic) Matches() Matches {
	return Prefix
}
func (ogpnic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenGroupSuffixIsOneOf []string
type ogs = OmitWhenGroupSuffixIsOneOf

var _ Filter = (*ogs)(nil)

func (f ogs) Values() []string {
	return f
}
func (ogs) Effect() Effect {
	return Omit
}
func (ogs) Target() Target {
	return Groups
}
func (ogs) Comparison() Comparison {
	return Equals
}
func (ogs) Matches() Matches {
	return Suffix
}
func (ogs) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenGroupSuffixIsOneOfIgnoringCase []string
type ogsic = OmitWhenGroupSuffixIsOneOfIgnoringCase

var _ Filter = (*ogsic)(nil)

func (f ogsic) Values() []string {
	return f
}
func (ogsic) Effect() Effect {
	return Omit
}
func (ogsic) Target() Target {
	return Groups
}
func (ogsic) Comparison() Comparison {
	return Equals
}
func (ogsic) Matches() Matches {
	return Suffix
}
func (ogsic) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenGroupPassedToFuncReturnsTrue []MatchFunc1
type ogft1 = OmitWhenGroupPassedToFuncReturnsTrue

var _ Filter = (*ogft1)(nil)

func (ogft1) Effect() Effect {
	return Omit
}
func (ogft1) Target() Target {
	return Groups
}
func (ogft1) Comparison() Comparison {
	return IsTrue
}
func (ogft1) Matches() Matches {
	return Func
}
func (ogft1) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenGroupPassedToFuncReturnsTrueIgnoringCase []MatchFunc1
type ogftic1 = OmitWhenGroupPassedToFuncReturnsTrueIgnoringCase

var _ Filter = (*ogftic1)(nil)

func (ogftic1) Effect() Effect {
	return Omit
}
func (ogftic1) Target() Target {
	return Groups
}
func (ogftic1) Comparison() Comparison {
	return IsTrue
}
func (ogftic1) Matches() Matches {
	return Func
}
func (ogftic1) IgnoreCase() bool {
	return true
}

//=============================================

type OmitWhenGroupPassedToFuncReturnsFalse []MatchFunc1
type ogff1 = OmitWhenGroupPassedToFuncReturnsFalse

var _ Filter = (*ogff1)(nil)

func (f ogff1) MatchFuncs1(index int) MatchFunc1 {
	return f[index]
}
func (ogff1) Effect() Effect {
	return Omit
}
func (ogff1) Target() Target {
	return Groups
}
func (ogff1) Comparison() Comparison {
	return IsFalse
}
func (ogff1) Matches() Matches {
	return Func
}
func (ogff1) IgnoreCase() bool {
	return false
}

//=============================================

type OmitWhenGroupPassedToFuncReturnsFalseIgnoringCase []MatchFunc1
type ogffic1 = OmitWhenGroupPassedToFuncReturnsFalseIgnoringCase

var _ Filter = (*ogffic1)(nil)

func (ogffic1) Effect() Effect {
	return Omit
}
func (ogffic1) Target() Target {
	return Groups
}
func (ogffic1) Comparison() Comparison {
	return IsFalse
}
func (ogffic1) Matches() Matches {
	return Func
}
func (ogffic1) IgnoreCase() bool {
	return true
}
