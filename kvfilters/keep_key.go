package kvfilters

//=============================================

type KeepWhenKeyIsOneOf []string
type kk = KeepWhenKeyIsOneOf

var _ Filter = (*kk)(nil)

func (f kk) Values() []string {
	return f
}
func (kk) Effect() Effect {
	return Keep
}
func (kk) Target() Target {
	return Keys
}
func (kk) Comparison() Comparison {
	return Equals
}
func (kk) Matches() Matches {
	return Entirety
}
func (kk) IgnoreCase() bool {
	return false
}

//=============================================

type KeepWhenKeyIsOneOfIgnoringCase []string
type kkic = KeepWhenKeyIsOneOfIgnoringCase

var _ Filter = (*kkic)(nil)

func (f kkic) Values() []string {
	return f
}
func (kkic) Effect() Effect {
	return Keep
}
func (kkic) Target() Target {
	return Keys
}
func (kkic) Comparison() Comparison {
	return Equals
}
func (kkic) Matches() Matches {
	return Entirety
}
func (kkic) IgnoreCase() bool {
	return true
}

//=============================================

type KeepWhenKeyPrefixIsOneOf []string
type kkp = KeepWhenKeyPrefixIsOneOf

var _ Filter = (*kkp)(nil)

func (f kkp) Values() []string {
	return f
}
func (kkp) Effect() Effect {
	return Keep
}
func (kkp) Target() Target {
	return Keys
}
func (kkp) Matches() Matches {
	return Prefix
}
func (kkp) Comparison() Comparison {
	return Equals
}
func (kkp) IgnoreCase() bool {
	return false
}

//=============================================
