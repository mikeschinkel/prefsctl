package macprefs

import (
	"fmt"

	"github.com/mikeschinkel/prefsctl/cliutil"
)

var _ cliutil.Result = (*Result)(nil)

func NewResult(a any) *Result {
	return &Result{Output: a}
}

type Result struct {
	Output any
}

func (r *Result) Result() {}

func (r *Result) String() string {
	switch t := r.Output.(type) {
	case string:
		return t
	case fmt.Stringer:
		return t.String()
	}
	return fmt.Sprintf("%v", r.Output)
}
