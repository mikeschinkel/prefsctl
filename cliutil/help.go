package cliutil

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Result interface {
	Result()
	fmt.Stringer
}

var _ Result = (*noResult)(nil)

type noResult struct{}

var NoResult *noResult

func (n *noResult) Result() {}
func (n *noResult) String() string {
	return ""
}

type Outcome struct {
	*cobra.Command
	Result Result
	Err    error
}

func NewErrorOutcome(err error) *Outcome {
	return &Outcome{
		Err: err,
	}
}

func NewCmdOutcome(err error, cmd *cobra.Command) *Outcome {
	return &Outcome{
		Command: cmd,
		Err:     err,
	}
}

func (r Outcome) WasError() bool {
	return r.Err != nil
}
