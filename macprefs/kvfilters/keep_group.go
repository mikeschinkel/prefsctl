package kvfilters

//=============================================

type KeepWhenGroupIsOneOf []string
type kd = KeepWhenGroupIsOneOf

var _ Filter = (*kd)(nil)

func (f kd) Values() []string {
	return f
}
func (kd) Effect() Effect {
	return Keep
}
func (kd) Target() Target {
	return Groups
}
func (kd) Comparison() Comparison {
	return Equals
}
func (kd) Matches() Matches {
	return Entirety
}
func (kd) IgnoreCase() bool {
	return false
}

//=============================================

type KeepWhenGroupIsOneOfIgnoringCase []string
type kdic = KeepWhenGroupIsOneOfIgnoringCase

var _ Filter = (*kdic)(nil)

func (f kdic) Values() []string {
	return f
}
func (kdic) Effect() Effect {
	return Keep
}
func (kdic) Target() Target {
	return Groups
}
func (kdic) Comparison() Comparison {
	return Equals
}
func (kdic) Matches() Matches {
	return Entirety
}
func (kdic) IgnoreCase() bool {
	return true
}
