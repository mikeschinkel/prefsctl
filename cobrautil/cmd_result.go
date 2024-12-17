package cobrautil

import (
	"github.com/spf13/cobra"
)

var NoResult = NewCmdResult(nil, nil)

type CmdResult struct {
	cmd Cmd
	err error
}

func NewErrResult(err error) CmdResult {
	return CmdResult{
		err: err,
	}
}

func NewCmdResult(cmd Cmd, err error) CmdResult {
	return CmdResult{
		cmd: cmd,
		err: err,
	}
}

func (n CmdResult) IsErr() bool {
	return n.err != nil
}

func (n CmdResult) SetErr(err error) {
	n.err = err
}

func (n CmdResult) Err() error {
	return n.err
}

func (n CmdResult) Cmd() Cmd {
	return n.cmd
}

func (n CmdResult) Command() *cobra.Command {
	return n.cmd.Command()
}

func (n CmdResult) String() string {
	s := n.err.Error()
	if s == "" {
		s = "Success"
	}
	return s
}
func (n CmdResult) Error() string {
	return n.err.Error()
}
